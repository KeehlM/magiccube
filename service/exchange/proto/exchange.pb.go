// Code generated by protoc-gen-go. DO NOT EDIT.
// source: exchange.proto

/*
Package exchange is a generated protocol buffer package.

It is generated from these files:
	exchange.proto

It has these top-level messages:
	ConsumerBuyRequest
	ConsumerBuyResponse
	QueryTxRequest
	QueryTxResponse
	QueryTxData
	TxRow
	TransferRequest
	TransferResponse
	QueryTransferRequest
	QueryTransferResponse
*/
package exchange

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConsumerBuyRequest struct {
	PostBody string `protobuf:"bytes,1,opt,name=postBody" json:"postBody"`
}

func (m *ConsumerBuyRequest) Reset()                    { *m = ConsumerBuyRequest{} }
func (m *ConsumerBuyRequest) String() string            { return proto.CompactTextString(m) }
func (*ConsumerBuyRequest) ProtoMessage()               {}
func (*ConsumerBuyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ConsumerBuyRequest) GetPostBody() string {
	if m != nil {
		return m.PostBody
	}
	return ""
}

type ConsumerBuyResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Data string `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *ConsumerBuyResponse) Reset()                    { *m = ConsumerBuyResponse{} }
func (m *ConsumerBuyResponse) String() string            { return proto.CompactTextString(m) }
func (*ConsumerBuyResponse) ProtoMessage()               {}
func (*ConsumerBuyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ConsumerBuyResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ConsumerBuyResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *ConsumerBuyResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type QueryTxRequest struct {
	PageSize uint64 `protobuf:"varint,1,opt,name=pageSize" json:"pageSize"`
	PageNum  uint64 `protobuf:"varint,2,opt,name=pageNum" json:"pageNum"`
	Username string `protobuf:"bytes,3,opt,name=username" json:"username"`
}

func (m *QueryTxRequest) Reset()                    { *m = QueryTxRequest{} }
func (m *QueryTxRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryTxRequest) ProtoMessage()               {}
func (*QueryTxRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *QueryTxRequest) GetPageSize() uint64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *QueryTxRequest) GetPageNum() uint64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *QueryTxRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type QueryTxResponse struct {
	Code uint32       `protobuf:"varint,1,opt,name=code" json:"code"`
	Data *QueryTxData `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string       `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *QueryTxResponse) Reset()                    { *m = QueryTxResponse{} }
func (m *QueryTxResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryTxResponse) ProtoMessage()               {}
func (*QueryTxResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *QueryTxResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *QueryTxResponse) GetData() *QueryTxData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *QueryTxResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type QueryTxData struct {
	PageNum  uint64   `protobuf:"varint,1,opt,name=pageNum" json:"pageNum"`
	RowCount uint64   `protobuf:"varint,2,opt,name=rowCount" json:"rowCount"`
	Row      []*TxRow `protobuf:"bytes,3,rep,name=row" json:"row"`
}

func (m *QueryTxData) Reset()                    { *m = QueryTxData{} }
func (m *QueryTxData) String() string            { return proto.CompactTextString(m) }
func (*QueryTxData) ProtoMessage()               {}
func (*QueryTxData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *QueryTxData) GetPageNum() uint64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *QueryTxData) GetRowCount() uint64 {
	if m != nil {
		return m.RowCount
	}
	return 0
}

func (m *QueryTxData) GetRow() []*TxRow {
	if m != nil {
		return m.Row
	}
	return nil
}

type TxRow struct {
	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId" json:"transaction_id"`
	From          string `protobuf:"bytes,2,opt,name=from" json:"from"`
	To            string `protobuf:"bytes,3,opt,name=to" json:"to"`
	Price         uint64 `protobuf:"varint,4,opt,name=price" json:"price"`
	Type          string `protobuf:"bytes,5,opt,name=type" json:"type"`
	Date          string `protobuf:"bytes,6,opt,name=date" json:"date"`
	BlockId       uint64 `protobuf:"varint,7,opt,name=block_id,json=blockId" json:"block_id"`
}

func (m *TxRow) Reset()                    { *m = TxRow{} }
func (m *TxRow) String() string            { return proto.CompactTextString(m) }
func (*TxRow) ProtoMessage()               {}
func (*TxRow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TxRow) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *TxRow) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *TxRow) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *TxRow) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *TxRow) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TxRow) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *TxRow) GetBlockId() uint64 {
	if m != nil {
		return m.BlockId
	}
	return 0
}

type TransferRequest struct {
	PostBody string `protobuf:"bytes,1,opt,name=postBody" json:"postBody"`
}

func (m *TransferRequest) Reset()                    { *m = TransferRequest{} }
func (m *TransferRequest) String() string            { return proto.CompactTextString(m) }
func (*TransferRequest) ProtoMessage()               {}
func (*TransferRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TransferRequest) GetPostBody() string {
	if m != nil {
		return m.PostBody
	}
	return ""
}

type TransferResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Data string `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *TransferResponse) Reset()                    { *m = TransferResponse{} }
func (m *TransferResponse) String() string            { return proto.CompactTextString(m) }
func (*TransferResponse) ProtoMessage()               {}
func (*TransferResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TransferResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TransferResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *TransferResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type QueryTransferRequest struct {
	Username  string `protobuf:"bytes,1,opt,name=username" json:"username"`
	SessionId string `protobuf:"bytes,2,opt,name=session_id,json=sessionId" json:"session_id"`
	Random    string `protobuf:"bytes,3,opt,name=random" json:"random"`
	Signature string `protobuf:"bytes,4,opt,name=signature" json:"signature"`
}

func (m *QueryTransferRequest) Reset()                    { *m = QueryTransferRequest{} }
func (m *QueryTransferRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryTransferRequest) ProtoMessage()               {}
func (*QueryTransferRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *QueryTransferRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *QueryTransferRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *QueryTransferRequest) GetRandom() string {
	if m != nil {
		return m.Random
	}
	return ""
}

func (m *QueryTransferRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type QueryTransferResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Data string `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *QueryTransferResponse) Reset()                    { *m = QueryTransferResponse{} }
func (m *QueryTransferResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryTransferResponse) ProtoMessage()               {}
func (*QueryTransferResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *QueryTransferResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *QueryTransferResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *QueryTransferResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ConsumerBuyRequest)(nil), "ConsumerBuyRequest")
	proto.RegisterType((*ConsumerBuyResponse)(nil), "ConsumerBuyResponse")
	proto.RegisterType((*QueryTxRequest)(nil), "QueryTxRequest")
	proto.RegisterType((*QueryTxResponse)(nil), "QueryTxResponse")
	proto.RegisterType((*QueryTxData)(nil), "QueryTxData")
	proto.RegisterType((*TxRow)(nil), "TxRow")
	proto.RegisterType((*TransferRequest)(nil), "TransferRequest")
	proto.RegisterType((*TransferResponse)(nil), "TransferResponse")
	proto.RegisterType((*QueryTransferRequest)(nil), "QueryTransferRequest")
	proto.RegisterType((*QueryTransferResponse)(nil), "QueryTransferResponse")
}

func init() { proto.RegisterFile("exchange.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc7, 0x37, 0x9b, 0x7e, 0x4e, 0xd9, 0xb6, 0x9a, 0x2d, 0xc8, 0x54, 0x20, 0x55, 0x96, 0x90,
	0xf6, 0x42, 0x84, 0xca, 0x8d, 0xe3, 0x2e, 0x1c, 0x2a, 0x21, 0xd0, 0x86, 0x5e, 0x38, 0x20, 0xe4,
	0x26, 0xde, 0x12, 0x41, 0xec, 0x60, 0x3b, 0xb4, 0xe5, 0x05, 0x78, 0x18, 0x5e, 0x12, 0xd9, 0x71,
	0x43, 0xba, 0x05, 0x84, 0xb4, 0xb7, 0xf9, 0x8f, 0x3c, 0x33, 0x3f, 0x7b, 0x66, 0x0c, 0x43, 0xbe,
	0x4d, 0x3e, 0x31, 0xb1, 0xe6, 0x51, 0xa1, 0xa4, 0x91, 0xf4, 0x19, 0xe0, 0x95, 0x14, 0xba, 0xcc,
	0xb9, 0xba, 0x2c, 0x77, 0x31, 0xff, 0x5a, 0x72, 0x6d, 0x70, 0x0a, 0xbd, 0x42, 0x6a, 0x73, 0x29,
	0xd3, 0x1d, 0x09, 0x66, 0xc1, 0x45, 0x3f, 0xae, 0x35, 0x7d, 0x0b, 0xe7, 0x07, 0x11, 0xba, 0x90,
	0x42, 0x73, 0x44, 0x68, 0x25, 0x32, 0xe5, 0xee, 0xf8, 0x59, 0xec, 0x6c, 0xeb, 0x4b, 0x99, 0x61,
	0xe4, 0xd4, 0xa5, 0x70, 0x36, 0x8e, 0x21, 0xcc, 0xf5, 0x9a, 0x84, 0xce, 0x65, 0x4d, 0xba, 0x82,
	0xe1, 0x75, 0xc9, 0xd5, 0x6e, 0xb9, 0x6d, 0x96, 0x67, 0x6b, 0xfe, 0x2e, 0xfb, 0x5e, 0xe5, 0x6b,
	0xc5, 0xb5, 0x46, 0x02, 0x5d, 0x6b, 0xbf, 0x29, 0x73, 0x97, 0xb6, 0x15, 0xef, 0xa5, 0x8d, 0x2a,
	0x35, 0x57, 0x82, 0xe5, 0xdc, 0xa7, 0xaf, 0x35, 0x7d, 0x0f, 0xa3, 0xba, 0xc6, 0x3f, 0x80, 0x67,
	0x0d, 0xe0, 0xc1, 0xfc, 0x5e, 0xe4, 0x63, 0x5e, 0x32, 0xc3, 0xfe, 0x8a, 0xff, 0x01, 0x06, 0x8d,
	0x63, 0x4d, 0xbe, 0xe0, 0x88, 0x4f, 0xc9, 0xcd, 0x95, 0x2c, 0x85, 0xf1, 0xe8, 0xb5, 0x46, 0x02,
	0xa1, 0x92, 0x1b, 0x12, 0xce, 0xc2, 0x8b, 0xc1, 0xbc, 0x13, 0x2d, 0xb7, 0xb1, 0xdc, 0xc4, 0xd6,
	0x45, 0x7f, 0x06, 0xd0, 0x76, 0x12, 0x9f, 0xc0, 0xd0, 0x28, 0x26, 0x34, 0x4b, 0x4c, 0x26, 0xc5,
	0xc7, 0x2c, 0xf5, 0xad, 0x39, 0x6b, 0x78, 0x17, 0xa9, 0xbd, 0xd7, 0x8d, 0x92, 0xf9, 0xfe, 0xd1,
	0xad, 0x8d, 0x43, 0x38, 0x35, 0xd2, 0x43, 0x9f, 0x1a, 0x89, 0x13, 0x68, 0x17, 0x2a, 0x4b, 0x38,
	0x69, 0x39, 0x8e, 0x4a, 0xd8, 0x48, 0xb3, 0x2b, 0x38, 0x69, 0x57, 0x91, 0xd6, 0xf6, 0x2d, 0xe4,
	0xa4, 0x53, 0xb7, 0x90, 0xe3, 0x43, 0xe8, 0xad, 0xbe, 0xc8, 0xe4, 0xb3, 0x45, 0xe8, 0x56, 0x77,
	0x74, 0x7a, 0x91, 0xd2, 0xa7, 0x30, 0x5a, 0x5a, 0x9a, 0x1b, 0xae, 0xfe, 0x67, 0x96, 0x5e, 0xc3,
	0xf8, 0xf7, 0xf1, 0x3b, 0x0f, 0xd2, 0x8f, 0x00, 0x26, 0x55, 0x2b, 0x8e, 0x11, 0xea, 0xc9, 0x08,
	0x0e, 0x27, 0x03, 0x1f, 0x03, 0x68, 0xae, 0xb5, 0x7f, 0xd1, 0xaa, 0x40, 0xdf, 0x7b, 0x16, 0x29,
	0x3e, 0x80, 0x8e, 0x62, 0x22, 0x95, 0xb9, 0x2f, 0xe4, 0x15, 0x3e, 0x82, 0xbe, 0xce, 0xd6, 0x82,
	0x99, 0x52, 0x55, 0xaf, 0x68, 0xa3, 0xf6, 0x0e, 0x7a, 0x0d, 0xf7, 0x6f, 0x81, 0xdc, 0xf5, 0x72,
	0xf3, 0x6f, 0xd0, 0x7b, 0xe5, 0x57, 0x17, 0x5f, 0xc0, 0xa0, 0xb1, 0x82, 0x78, 0x1e, 0x1d, 0xaf,
	0xf0, 0x74, 0x12, 0xfd, 0x61, 0x4b, 0xe9, 0x09, 0x46, 0xd0, 0xf5, 0xe3, 0x8a, 0xa3, 0xe8, 0x70,
	0xef, 0xa6, 0xe3, 0xe8, 0xd6, 0x92, 0xd0, 0x93, 0x55, 0xc7, 0xfd, 0x13, 0xcf, 0x7f, 0x05, 0x00,
	0x00, 0xff, 0xff, 0xac, 0x02, 0x7f, 0x11, 0x39, 0x04, 0x00, 0x00,
}
