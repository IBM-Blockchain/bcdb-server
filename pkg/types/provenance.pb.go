// Code generated by protoc-gen-go. DO NOT EDIT.
// source: provenance.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TxProof struct {
	Header               *BlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Path                 [][]byte     `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TxProof) Reset()         { *m = TxProof{} }
func (m *TxProof) String() string { return proto.CompactTextString(m) }
func (*TxProof) ProtoMessage()    {}
func (*TxProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd9885392b6ac305, []int{0}
}

func (m *TxProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxProof.Unmarshal(m, b)
}
func (m *TxProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxProof.Marshal(b, m, deterministic)
}
func (m *TxProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxProof.Merge(m, src)
}
func (m *TxProof) XXX_Size() int {
	return xxx_messageInfo_TxProof.Size(m)
}
func (m *TxProof) XXX_DiscardUnknown() {
	xxx_messageInfo_TxProof.DiscardUnknown(m)
}

var xxx_messageInfo_TxProof proto.InternalMessageInfo

func (m *TxProof) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *TxProof) GetPath() [][]byte {
	if m != nil {
		return m.Path
	}
	return nil
}

type BlockProof struct {
	BlockNumber          uint64         `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	Path                 []*BlockHeader `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BlockProof) Reset()         { *m = BlockProof{} }
func (m *BlockProof) String() string { return proto.CompactTextString(m) }
func (*BlockProof) ProtoMessage()    {}
func (*BlockProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd9885392b6ac305, []int{1}
}

func (m *BlockProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockProof.Unmarshal(m, b)
}
func (m *BlockProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockProof.Marshal(b, m, deterministic)
}
func (m *BlockProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockProof.Merge(m, src)
}
func (m *BlockProof) XXX_Size() int {
	return xxx_messageInfo_BlockProof.Size(m)
}
func (m *BlockProof) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockProof.DiscardUnknown(m)
}

var xxx_messageInfo_BlockProof proto.InternalMessageInfo

func (m *BlockProof) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *BlockProof) GetPath() []*BlockHeader {
	if m != nil {
		return m.Path
	}
	return nil
}

type GetBlockQuery struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	BlockNumber          uint64   `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockQuery) Reset()         { *m = GetBlockQuery{} }
func (m *GetBlockQuery) String() string { return proto.CompactTextString(m) }
func (*GetBlockQuery) ProtoMessage()    {}
func (*GetBlockQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd9885392b6ac305, []int{2}
}

func (m *GetBlockQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockQuery.Unmarshal(m, b)
}
func (m *GetBlockQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockQuery.Marshal(b, m, deterministic)
}
func (m *GetBlockQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockQuery.Merge(m, src)
}
func (m *GetBlockQuery) XXX_Size() int {
	return xxx_messageInfo_GetBlockQuery.Size(m)
}
func (m *GetBlockQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockQuery.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockQuery proto.InternalMessageInfo

func (m *GetBlockQuery) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *GetBlockQuery) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

type GetBlockQueryEnvelope struct {
	Payload              *GetBlockQuery `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature            []byte         `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetBlockQueryEnvelope) Reset()         { *m = GetBlockQueryEnvelope{} }
func (m *GetBlockQueryEnvelope) String() string { return proto.CompactTextString(m) }
func (*GetBlockQueryEnvelope) ProtoMessage()    {}
func (*GetBlockQueryEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd9885392b6ac305, []int{3}
}

func (m *GetBlockQueryEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockQueryEnvelope.Unmarshal(m, b)
}
func (m *GetBlockQueryEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockQueryEnvelope.Marshal(b, m, deterministic)
}
func (m *GetBlockQueryEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockQueryEnvelope.Merge(m, src)
}
func (m *GetBlockQueryEnvelope) XXX_Size() int {
	return xxx_messageInfo_GetBlockQueryEnvelope.Size(m)
}
func (m *GetBlockQueryEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockQueryEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockQueryEnvelope proto.InternalMessageInfo

func (m *GetBlockQueryEnvelope) GetPayload() *GetBlockQuery {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *GetBlockQueryEnvelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type GetBlockResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	BlockHeader          *BlockHeader    `protobuf:"bytes,2,opt,name=blockHeader,proto3" json:"blockHeader,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetBlockResponse) Reset()         { *m = GetBlockResponse{} }
func (m *GetBlockResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlockResponse) ProtoMessage()    {}
func (*GetBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd9885392b6ac305, []int{4}
}

func (m *GetBlockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockResponse.Unmarshal(m, b)
}
func (m *GetBlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockResponse.Marshal(b, m, deterministic)
}
func (m *GetBlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockResponse.Merge(m, src)
}
func (m *GetBlockResponse) XXX_Size() int {
	return xxx_messageInfo_GetBlockResponse.Size(m)
}
func (m *GetBlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockResponse proto.InternalMessageInfo

func (m *GetBlockResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetBlockResponse) GetBlockHeader() *BlockHeader {
	if m != nil {
		return m.BlockHeader
	}
	return nil
}

type GetBlockResponseEnvelope struct {
	Payload              *GetBlockResponse `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature            []byte            `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetBlockResponseEnvelope) Reset()         { *m = GetBlockResponseEnvelope{} }
func (m *GetBlockResponseEnvelope) String() string { return proto.CompactTextString(m) }
func (*GetBlockResponseEnvelope) ProtoMessage()    {}
func (*GetBlockResponseEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd9885392b6ac305, []int{5}
}

func (m *GetBlockResponseEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockResponseEnvelope.Unmarshal(m, b)
}
func (m *GetBlockResponseEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockResponseEnvelope.Marshal(b, m, deterministic)
}
func (m *GetBlockResponseEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockResponseEnvelope.Merge(m, src)
}
func (m *GetBlockResponseEnvelope) XXX_Size() int {
	return xxx_messageInfo_GetBlockResponseEnvelope.Size(m)
}
func (m *GetBlockResponseEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockResponseEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockResponseEnvelope proto.InternalMessageInfo

func (m *GetBlockResponseEnvelope) GetPayload() *GetBlockResponse {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *GetBlockResponseEnvelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type GetLedgerPathQuery struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	StartBlockNumber     uint64   `protobuf:"varint,2,opt,name=start_block_number,json=startBlockNumber,proto3" json:"start_block_number,omitempty"`
	EndBlockNumber       uint64   `protobuf:"varint,3,opt,name=end_block_number,json=endBlockNumber,proto3" json:"end_block_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLedgerPathQuery) Reset()         { *m = GetLedgerPathQuery{} }
func (m *GetLedgerPathQuery) String() string { return proto.CompactTextString(m) }
func (*GetLedgerPathQuery) ProtoMessage()    {}
func (*GetLedgerPathQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd9885392b6ac305, []int{6}
}

func (m *GetLedgerPathQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLedgerPathQuery.Unmarshal(m, b)
}
func (m *GetLedgerPathQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLedgerPathQuery.Marshal(b, m, deterministic)
}
func (m *GetLedgerPathQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLedgerPathQuery.Merge(m, src)
}
func (m *GetLedgerPathQuery) XXX_Size() int {
	return xxx_messageInfo_GetLedgerPathQuery.Size(m)
}
func (m *GetLedgerPathQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLedgerPathQuery.DiscardUnknown(m)
}

var xxx_messageInfo_GetLedgerPathQuery proto.InternalMessageInfo

func (m *GetLedgerPathQuery) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *GetLedgerPathQuery) GetStartBlockNumber() uint64 {
	if m != nil {
		return m.StartBlockNumber
	}
	return 0
}

func (m *GetLedgerPathQuery) GetEndBlockNumber() uint64 {
	if m != nil {
		return m.EndBlockNumber
	}
	return 0
}

type GetLedgerPathQueryEnvelope struct {
	Payload              *GetLedgerPathQuery `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature            []byte              `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetLedgerPathQueryEnvelope) Reset()         { *m = GetLedgerPathQueryEnvelope{} }
func (m *GetLedgerPathQueryEnvelope) String() string { return proto.CompactTextString(m) }
func (*GetLedgerPathQueryEnvelope) ProtoMessage()    {}
func (*GetLedgerPathQueryEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd9885392b6ac305, []int{7}
}

func (m *GetLedgerPathQueryEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLedgerPathQueryEnvelope.Unmarshal(m, b)
}
func (m *GetLedgerPathQueryEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLedgerPathQueryEnvelope.Marshal(b, m, deterministic)
}
func (m *GetLedgerPathQueryEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLedgerPathQueryEnvelope.Merge(m, src)
}
func (m *GetLedgerPathQueryEnvelope) XXX_Size() int {
	return xxx_messageInfo_GetLedgerPathQueryEnvelope.Size(m)
}
func (m *GetLedgerPathQueryEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLedgerPathQueryEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_GetLedgerPathQueryEnvelope proto.InternalMessageInfo

func (m *GetLedgerPathQueryEnvelope) GetPayload() *GetLedgerPathQuery {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *GetLedgerPathQueryEnvelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type GetLedgerPathResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	BlockHeaders         []*BlockHeader  `protobuf:"bytes,2,rep,name=blockHeaders,proto3" json:"blockHeaders,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetLedgerPathResponse) Reset()         { *m = GetLedgerPathResponse{} }
func (m *GetLedgerPathResponse) String() string { return proto.CompactTextString(m) }
func (*GetLedgerPathResponse) ProtoMessage()    {}
func (*GetLedgerPathResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd9885392b6ac305, []int{8}
}

func (m *GetLedgerPathResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLedgerPathResponse.Unmarshal(m, b)
}
func (m *GetLedgerPathResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLedgerPathResponse.Marshal(b, m, deterministic)
}
func (m *GetLedgerPathResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLedgerPathResponse.Merge(m, src)
}
func (m *GetLedgerPathResponse) XXX_Size() int {
	return xxx_messageInfo_GetLedgerPathResponse.Size(m)
}
func (m *GetLedgerPathResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLedgerPathResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLedgerPathResponse proto.InternalMessageInfo

func (m *GetLedgerPathResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetLedgerPathResponse) GetBlockHeaders() []*BlockHeader {
	if m != nil {
		return m.BlockHeaders
	}
	return nil
}

type GetLedgerPathResponseEnvelope struct {
	Payload              *GetLedgerPathResponse `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature            []byte                 `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetLedgerPathResponseEnvelope) Reset()         { *m = GetLedgerPathResponseEnvelope{} }
func (m *GetLedgerPathResponseEnvelope) String() string { return proto.CompactTextString(m) }
func (*GetLedgerPathResponseEnvelope) ProtoMessage()    {}
func (*GetLedgerPathResponseEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd9885392b6ac305, []int{9}
}

func (m *GetLedgerPathResponseEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLedgerPathResponseEnvelope.Unmarshal(m, b)
}
func (m *GetLedgerPathResponseEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLedgerPathResponseEnvelope.Marshal(b, m, deterministic)
}
func (m *GetLedgerPathResponseEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLedgerPathResponseEnvelope.Merge(m, src)
}
func (m *GetLedgerPathResponseEnvelope) XXX_Size() int {
	return xxx_messageInfo_GetLedgerPathResponseEnvelope.Size(m)
}
func (m *GetLedgerPathResponseEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLedgerPathResponseEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_GetLedgerPathResponseEnvelope proto.InternalMessageInfo

func (m *GetLedgerPathResponseEnvelope) GetPayload() *GetLedgerPathResponse {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *GetLedgerPathResponseEnvelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*TxProof)(nil), "types.TxProof")
	proto.RegisterType((*BlockProof)(nil), "types.BlockProof")
	proto.RegisterType((*GetBlockQuery)(nil), "types.GetBlockQuery")
	proto.RegisterType((*GetBlockQueryEnvelope)(nil), "types.GetBlockQueryEnvelope")
	proto.RegisterType((*GetBlockResponse)(nil), "types.GetBlockResponse")
	proto.RegisterType((*GetBlockResponseEnvelope)(nil), "types.GetBlockResponseEnvelope")
	proto.RegisterType((*GetLedgerPathQuery)(nil), "types.GetLedgerPathQuery")
	proto.RegisterType((*GetLedgerPathQueryEnvelope)(nil), "types.GetLedgerPathQueryEnvelope")
	proto.RegisterType((*GetLedgerPathResponse)(nil), "types.GetLedgerPathResponse")
	proto.RegisterType((*GetLedgerPathResponseEnvelope)(nil), "types.GetLedgerPathResponseEnvelope")
}

func init() { proto.RegisterFile("provenance.proto", fileDescriptor_fd9885392b6ac305) }

var fileDescriptor_fd9885392b6ac305 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x56, 0xb7, 0xd1, 0x69, 0xb7, 0x82, 0x2a, 0x8b, 0x41, 0x28, 0x43, 0x1a, 0x79, 0x40, 0xd1,
	0x04, 0xa9, 0x60, 0x68, 0x3f, 0xa0, 0x02, 0x8d, 0x21, 0x84, 0x46, 0x84, 0x84, 0xc4, 0x4b, 0xe5,
	0x24, 0x47, 0x13, 0xad, 0xb5, 0x83, 0xed, 0x0c, 0xfa, 0xc2, 0x13, 0x3f, 0x1c, 0xed, 0xe2, 0xb4,
	0x89, 0x1b, 0x4d, 0x63, 0x6f, 0xf5, 0xe7, 0xef, 0xbe, 0xef, 0xbe, 0x3b, 0x37, 0x30, 0x2c, 0x94,
	0xbc, 0x42, 0xc1, 0x45, 0x82, 0x61, 0xa1, 0xa4, 0x91, 0xec, 0x9e, 0x59, 0x16, 0xa8, 0x47, 0x4f,
	0xe3, 0xb9, 0x4c, 0x2e, 0xa7, 0x5c, 0xa4, 0x53, 0xa3, 0xb8, 0xd0, 0x3c, 0x31, 0xb9, 0x14, 0x15,
	0x67, 0xe4, 0xfd, 0x2c, 0x51, 0x2d, 0xe9, 0x52, 0xa1, 0x2e, 0xa4, 0xd0, 0xb6, 0xda, 0x3f, 0x87,
	0xdd, 0xaf, 0xbf, 0x2f, 0x94, 0x94, 0x3f, 0xd8, 0x31, 0xf4, 0x33, 0xe4, 0x29, 0x2a, 0xaf, 0x77,
	0xd4, 0x0b, 0xf6, 0xdf, 0xb0, 0x90, 0x94, 0xc3, 0xc9, 0xb5, 0xf0, 0x07, 0xba, 0x89, 0x2c, 0x83,
	0x31, 0xd8, 0x29, 0xb8, 0xc9, 0xbc, 0xad, 0xa3, 0xed, 0x60, 0x10, 0xd1, 0x6f, 0xff, 0x1b, 0x00,
	0x51, 0x2b, 0xb5, 0xe7, 0x30, 0xa8, 0x3a, 0x12, 0xe5, 0x22, 0xb6, 0x9a, 0x3b, 0xd1, 0x3e, 0x61,
	0x9f, 0x09, 0x62, 0x2f, 0x1a, 0x22, 0xdd, 0x76, 0x95, 0xf0, 0x47, 0xb8, 0x7f, 0x86, 0x86, 0xf0,
	0x2f, 0xd7, 0x39, 0xd8, 0x23, 0xe8, 0x97, 0x1a, 0xd5, 0xf9, 0x3b, 0x52, 0xdd, 0x8b, 0xec, 0x69,
	0xc3, 0x73, 0x6b, 0xc3, 0xd3, 0x47, 0x38, 0x68, 0x69, 0xbd, 0x17, 0x57, 0x38, 0x97, 0x05, 0xb2,
	0x10, 0x76, 0x0b, 0xbe, 0x9c, 0x4b, 0x9e, 0xda, 0xf8, 0x0f, 0x6d, 0x3f, 0x2d, 0x7a, 0x54, 0x93,
	0xd8, 0x21, 0xec, 0xe9, 0x7c, 0x26, 0xb8, 0x29, 0x15, 0x92, 0xd1, 0x20, 0x5a, 0x03, 0xfe, 0x2f,
	0x18, 0xd6, 0x75, 0x91, 0x1d, 0x38, 0x7b, 0xe5, 0xcc, 0xf7, 0xc0, 0x1a, 0xd4, 0x04, 0x67, 0xc4,
	0x6f, 0xa1, 0x6a, 0xbc, 0x82, 0xc9, 0xa2, 0x7b, 0x48, 0x4d, 0x9a, 0x7f, 0x09, 0x9e, 0x6b, 0xbc,
	0x8a, 0xf8, 0xda, 0x8d, 0xf8, 0xd8, 0x89, 0x58, 0x57, 0xdc, 0x36, 0xe5, 0xdf, 0x1e, 0xb0, 0x33,
	0x34, 0x9f, 0x30, 0x9d, 0xa1, 0xba, 0xe0, 0x26, 0xbb, 0x79, 0x3d, 0x2f, 0x81, 0x69, 0xc3, 0x95,
	0x99, 0x76, 0x2c, 0x69, 0x48, 0x37, 0x93, 0xc6, 0xeb, 0x08, 0x60, 0x88, 0x22, 0x6d, 0x73, 0xb7,
	0x89, 0xfb, 0x00, 0x45, 0xda, 0x60, 0xfa, 0x12, 0x46, 0x9b, 0x5d, 0xac, 0x52, 0x9f, 0xb8, 0xa9,
	0x9f, 0xac, 0x53, 0x3b, 0x35, 0xb7, 0xcd, 0xfd, 0x87, 0x1e, 0xd1, 0xba, 0xf8, 0xae, 0x2b, 0x3e,
	0xb5, 0xef, 0xb5, 0x82, 0xf5, 0x0d, 0x7f, 0x84, 0x16, 0xcf, 0x2f, 0xe1, 0x59, 0xa7, 0xff, 0x2a,
	0xf3, 0xa9, 0x9b, 0xf9, 0xb0, 0x2b, 0xf3, 0x7f, 0xae, 0x7b, 0x72, 0xfc, 0x3d, 0x98, 0xe5, 0x26,
	0x2b, 0xe3, 0x30, 0x8f, 0x17, 0x61, 0x22, 0x17, 0x63, 0xea, 0x2a, 0xc9, 0x78, 0x2e, 0xd2, 0x78,
	0x4c, 0x1f, 0x14, 0x3d, 0x26, 0xa3, 0xb8, 0x4f, 0xa7, 0x93, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x3f, 0xcf, 0xfe, 0x64, 0xb0, 0x04, 0x00, 0x00,
}