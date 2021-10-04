// Copyright IBM Corp. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package httphandler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/hyperledger-labs/orion-server/internal/bcdb"
	"github.com/hyperledger-labs/orion-server/internal/errors"
	"github.com/hyperledger-labs/orion-server/internal/httputils"
	"github.com/hyperledger-labs/orion-server/pkg/constants"
	"github.com/hyperledger-labs/orion-server/pkg/cryptoservice"
	"github.com/hyperledger-labs/orion-server/pkg/logger"
	"github.com/hyperledger-labs/orion-server/pkg/types"
	"github.com/gorilla/mux"
)

// dataRequestHandler handles query and transaction associated
// the user's data/state
type dataRequestHandler struct {
	db          bcdb.DB
	sigVerifier *cryptoservice.SignatureVerifier
	router      *mux.Router
	txHandler   *txHandler
	logger      *logger.SugarLogger
}

// NewDataRequestHandler returns handler capable to serve incoming data requests
func NewDataRequestHandler(db bcdb.DB, logger *logger.SugarLogger) http.Handler {
	handler := &dataRequestHandler{
		db:          db,
		sigVerifier: cryptoservice.NewVerifier(db, logger),
		router:      mux.NewRouter(),
		txHandler: &txHandler{
			db: db,
		},
		logger: logger,
	}

	handler.router.HandleFunc(constants.GetData, handler.dataQuery).Methods(http.MethodGet)
	handler.router.HandleFunc(constants.PostDataTx, handler.dataTransaction).Methods(http.MethodPost)
	handler.router.HandleFunc(constants.PostDataQuery, handler.dataJSONQuery).Methods(http.MethodPost)

	return handler
}

func (d *dataRequestHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	d.router.ServeHTTP(response, request)
}

func (d *dataRequestHandler) dataQuery(response http.ResponseWriter, request *http.Request) {
	payload, respondedErr := extractVerifiedQueryPayload(response, request, constants.GetData, d.sigVerifier)
	if respondedErr {
		return
	}
	query := payload.(*types.GetDataQuery)

	if !d.db.IsDBExists(query.DbName) {
		httputils.SendHTTPResponse(response, http.StatusBadRequest, &types.HttpResponseErr{
			ErrMsg: "error db '" + query.DbName + "' doesn't exist",
		})
		return
	}

	data, err := d.db.GetData(query.DbName, query.UserId, query.Key)
	if err != nil {
		var status int

		switch err.(type) {
		case *errors.PermissionErr:
			status = http.StatusForbidden
		default:
			status = http.StatusInternalServerError
		}

		httputils.SendHTTPResponse(
			response,
			status,
			&types.HttpResponseErr{
				ErrMsg: "error while processing '" + request.Method + " " + request.URL.String() + "' because " + err.Error(),
			})
		return
	}

	httputils.SendHTTPResponse(response, http.StatusOK, data)
}

func (d *dataRequestHandler) dataTransaction(response http.ResponseWriter, request *http.Request) {
	timeout, err := validateAndParseTxPostHeader(&request.Header)
	if err != nil {
		httputils.SendHTTPResponse(response, http.StatusBadRequest, &types.HttpResponseErr{ErrMsg: err.Error()})
		return
	}

	requestData := json.NewDecoder(request.Body)
	requestData.DisallowUnknownFields()

	txEnv := &types.DataTxEnvelope{}
	if err := requestData.Decode(txEnv); err != nil {
		httputils.SendHTTPResponse(response, http.StatusBadRequest, &types.HttpResponseErr{ErrMsg: err.Error()})
		return
	}

	if txEnv.Payload == nil {
		httputils.SendHTTPResponse(response, http.StatusBadRequest,
			&types.HttpResponseErr{ErrMsg: fmt.Sprintf("missing transaction envelope payload (%T)", txEnv.Payload)})
		return
	}

	if len(txEnv.Payload.MustSignUserIds) == 0 {
		httputils.SendHTTPResponse(response, http.StatusBadRequest,
			&types.HttpResponseErr{ErrMsg: fmt.Sprintf("missing UserID in transaction envelope payload (%T)", txEnv.Payload)})
		return
	}

	var notSigned []string
	for _, user := range txEnv.Payload.MustSignUserIds {
		if user == "" {
			httputils.SendHTTPResponse(response, http.StatusBadRequest,
				&types.HttpResponseErr{ErrMsg: "an empty UserID in MustSignUserIDs list present in the transaction envelope"})
			return
		}

		if _, ok := txEnv.Signatures[user]; !ok {
			notSigned = append(notSigned, user)
		}
	}
	if len(notSigned) > 0 {
		sort.Strings(notSigned)
		httputils.SendHTTPResponse(response, http.StatusBadRequest,
			&types.HttpResponseErr{ErrMsg: "users [" + strings.Join(notSigned, ",") + "] in the must sign list have not signed the transaction"})
		return
	}

	for _, userID := range txEnv.Payload.MustSignUserIds {
		if err, code := VerifyRequestSignature(d.sigVerifier, userID, txEnv.Signatures[userID], txEnv.Payload); err != nil {
			httputils.SendHTTPResponse(response, code, &types.HttpResponseErr{ErrMsg: err.Error()})
			return
		}
	}

	d.txHandler.handleTransaction(response, request, txEnv, timeout)
}

func (d *dataRequestHandler) dataJSONQuery(response http.ResponseWriter, request *http.Request) {
	payload, respondedErr := extractVerifiedQueryPayload(response, request, constants.PostDataQuery, d.sigVerifier)
	if respondedErr {
		return
	}
	query := payload.(*types.DataJSONQuery)

	if !d.db.IsDBExists(query.DbName) {
		httputils.SendHTTPResponse(response, http.StatusBadRequest, &types.HttpResponseErr{
			ErrMsg: "'" + query.DbName + "' does not exist",
		})
		return
	}

	parent := request.Context()
	data, err := d.db.DataQuery(parent, query.DbName, query.UserId, []byte(query.Query))

	select {
	case <-parent.Done():
		if parent.Err() == context.DeadlineExceeded {
			d.logger.Debug("request has been timeout")
			httputils.SendHTTPResponse(response, http.StatusRequestTimeout, nil)
			return
		}

		d.logger.Debug("http client context has been cancelled")
	default:
		if err != nil {
			var status int

			switch err.(type) {
			case *errors.PermissionErr:
				status = http.StatusForbidden
			default:
				status = http.StatusInternalServerError
			}

			httputils.SendHTTPResponse(
				response,
				status,
				&types.HttpResponseErr{
					ErrMsg: "error while processing '" + request.Method + " " + request.URL.String() + "' because " + err.Error(),
				})
			return
		}

		httputils.SendHTTPResponse(response, http.StatusOK, data)
	}
}
