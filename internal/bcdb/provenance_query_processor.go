package bcdb

import (
	"github.ibm.com/blockchaindb/server/internal/provenance"
	"github.ibm.com/blockchaindb/server/pkg/crypto"
	"github.ibm.com/blockchaindb/server/pkg/cryptoservice"
	"github.ibm.com/blockchaindb/server/pkg/logger"
	"github.ibm.com/blockchaindb/server/pkg/types"
)

type provenanceQueryProcessor struct {
	nodeID          string
	signer          crypto.Signer
	provenanceStore *provenance.Store
	logger          *logger.SugarLogger
}

type provenanceQueryProcessorConfig struct {
	nodeID          string
	signer          crypto.Signer
	provenanceStore *provenance.Store
	logger          *logger.SugarLogger
}

func newProvenanceQueryProcessor(conf *provenanceQueryProcessorConfig) *provenanceQueryProcessor {
	return &provenanceQueryProcessor{
		nodeID:          conf.nodeID,
		signer:          conf.signer,
		provenanceStore: conf.provenanceStore,
		logger:          conf.logger,
	}
}

// GetValues returns all values associated with a given key
func (p *provenanceQueryProcessor) GetValues(dbName, key string) (*types.GetHistoricalDataResponseEnvelope, error) {
	values, err := p.provenanceStore.GetValues(dbName, key)
	if err != nil {
		return nil, err
	}

	return p.composeHistoricalDataResponseEnvelope(values)
}

// GetValueAt returns the value of a given key at a particular version
func (p *provenanceQueryProcessor) GetValueAt(dbName, key string, version *types.Version) (*types.GetHistoricalDataResponseEnvelope, error) {
	value, err := p.provenanceStore.GetValueAt(dbName, key, version)
	if err != nil {
		return nil, err
	}

	return p.composeHistoricalDataResponseEnvelope([]*types.ValueWithMetadata{value})
}

// GetPreviousValues returns previous values of a given key and a version. The number of records returned would be limited
// by the limit parameters.
func (p *provenanceQueryProcessor) GetPreviousValues(dbName, key string, version *types.Version) (*types.GetHistoricalDataResponseEnvelope, error) {
	values, err := p.provenanceStore.GetPreviousValues(dbName, key, version, -1)
	if err != nil {
		return nil, err
	}

	return p.composeHistoricalDataResponseEnvelope(values)
}

// GetNextValues returns next values of a given key and a version. The number of records returned would be limited
// by the limit parameters.
func (p *provenanceQueryProcessor) GetNextValues(dbName, key string, version *types.Version) (*types.GetHistoricalDataResponseEnvelope, error) {
	values, err := p.provenanceStore.GetNextValues(dbName, key, version, -1)
	if err != nil {
		return nil, err
	}

	return p.composeHistoricalDataResponseEnvelope(values)
}

// GetValuesReadByUser returns all values read by a given user
func (p *provenanceQueryProcessor) GetValuesReadByUser(userID string) (*types.GetDataReadByResponseEnvelope, error) {
	kvs, err := p.provenanceStore.GetValuesReadByUser(userID)
	if err != nil {
		return nil, err
	}

	response := &types.GetDataReadByResponseEnvelope{
		Payload: &types.GetDataReadByResponse{
			Header: &types.ResponseHeader{
				NodeID: p.nodeID,
			},
			KVs: kvs,
		},
	}

	response.Signature, err = cryptoservice.SignQueryResponse(p.signer, response.Payload)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// GetValuesReadByUser returns all values read by a given user
func (p *provenanceQueryProcessor) GetValuesWrittenByUser(userID string) (*types.GetDataWrittenByResponseEnvelope, error) {
	kvs, err := p.provenanceStore.GetValuesWrittenByUser(userID)
	if err != nil {
		return nil, err
	}

	response := &types.GetDataWrittenByResponseEnvelope{
		Payload: &types.GetDataWrittenByResponse{
			Header: &types.ResponseHeader{
				NodeID: p.nodeID,
			},
			KVs: kvs,
		},
	}

	response.Signature, err = cryptoservice.SignQueryResponse(p.signer, response.Payload)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// GetReaders returns all userIDs who have accessed a given key as well as the access frequency
func (p *provenanceQueryProcessor) GetReaders(dbName, key string) (*types.GetDataReadersResponseEnvelope, error) {
	users, err := p.provenanceStore.GetReaders(dbName, key)
	if err != nil {
		return nil, err
	}

	response := &types.GetDataReadersResponseEnvelope{
		Payload: &types.GetDataReadersResponse{
			Header: &types.ResponseHeader{
				NodeID: p.nodeID,
			},
			ReadBy: users,
		},
	}

	response.Signature, err = cryptoservice.SignQueryResponse(p.signer, response.Payload)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// GetReaders returns all userIDs who have accessed a given key as well as the access frequency
func (p *provenanceQueryProcessor) GetWriters(dbName, key string) (*types.GetDataWritersResponseEnvelope, error) {
	users, err := p.provenanceStore.GetWriters(dbName, key)
	if err != nil {
		return nil, err
	}

	response := &types.GetDataWritersResponseEnvelope{
		Payload: &types.GetDataWritersResponse{
			Header: &types.ResponseHeader{
				NodeID: p.nodeID,
			},
			WrittenBy: users,
		},
	}

	response.Signature, err = cryptoservice.SignQueryResponse(p.signer, response.Payload)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// GetTxIDsSubmittedByUser returns all ids of all transactions submitted by a given user
func (p *provenanceQueryProcessor) GetTxIDsSubmittedByUser(userID string) (*types.GetTxIDsSubmittedByResponseEnvelope, error) {
	txIDs, err := p.provenanceStore.GetTxIDsSubmittedByUser(userID)
	if err != nil {
		return nil, err
	}

	response := &types.GetTxIDsSubmittedByResponseEnvelope{
		Payload: &types.GetTxIDsSubmittedByResponse{
			Header: &types.ResponseHeader{
				NodeID: p.nodeID,
			},
			TxIDs: txIDs,
		},
	}

	response.Signature, err = cryptoservice.SignQueryResponse(p.signer, response.Payload)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *provenanceQueryProcessor) composeHistoricalDataResponseEnvelope(values []*types.ValueWithMetadata) (*types.GetHistoricalDataResponseEnvelope, error) {
	response := &types.GetHistoricalDataResponseEnvelope{
		Payload: &types.GetHistoricalDataResponse{
			Header: &types.ResponseHeader{
				NodeID: p.nodeID,
			},
			Values: values,
		},
	}

	var err error
	response.Signature, err = cryptoservice.SignQueryResponse(p.signer, response.Payload)
	if err != nil {
		return nil, err
	}
	return response, nil
}