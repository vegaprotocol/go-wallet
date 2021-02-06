// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vega.proto

package proto

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

// Represents a transaction to be sent to Vega.
type Transaction struct {
	// One of the set of Vega commands (proto marshalled).
	InputData []byte `protobuf:"bytes,1,opt,name=input_data,json=inputData,proto3" json:"input_data,omitempty"`
	// A random number used to provided uniqueness and prevents
	// against replay attack.
	Nonce uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// The block height associated to the transaction.
	// This should always be current height of the node at the time of sending the Tx.
	// BlockHeight is used as a mechanism for replay protection.
	BlockHeight uint64 `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// The sender of the transaction.
	// Any of the following would be valid:
	//
	// Types that are valid to be assigned to From:
	//	*Transaction_Address
	//	*Transaction_PubKey
	From                 isTransaction_From `protobuf_oneof:"from"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb6b8173ee11af27, []int{0}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetInputData() []byte {
	if m != nil {
		return m.InputData
	}
	return nil
}

func (m *Transaction) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Transaction) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

type isTransaction_From interface {
	isTransaction_From()
}

type Transaction_Address struct {
	Address []byte `protobuf:"bytes,1001,opt,name=address,proto3,oneof"`
}

type Transaction_PubKey struct {
	PubKey []byte `protobuf:"bytes,1002,opt,name=pubKey,proto3,oneof"`
}

func (*Transaction_Address) isTransaction_From() {}

func (*Transaction_PubKey) isTransaction_From() {}

func (m *Transaction) GetFrom() isTransaction_From {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Transaction) GetAddress() []byte {
	if x, ok := m.GetFrom().(*Transaction_Address); ok {
		return x.Address
	}
	return nil
}

func (m *Transaction) GetPubKey() []byte {
	if x, ok := m.GetFrom().(*Transaction_PubKey); ok {
		return x.PubKey
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Transaction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Transaction_Address)(nil),
		(*Transaction_PubKey)(nil),
	}
}

// A signature to be authenticate a transaction
// and to be verified by the vega network
type Signature struct {
	// The bytes of the signature
	Sig []byte `protobuf:"bytes,1,opt,name=sig,proto3" json:"sig,omitempty"`
	// The algorithm used to create the signature
	Algo string `protobuf:"bytes,2,opt,name=algo,proto3" json:"algo,omitempty"`
	// The version of the signature used to create the signature
	Version              uint64   `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb6b8173ee11af27, []int{1}
}

func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (m *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(m, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *Signature) GetAlgo() string {
	if m != nil {
		return m.Algo
	}
	return ""
}

func (m *Signature) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

// A bundle of a transaction and it's signature.
type SignedBundle struct {
	// Transaction payload (proto marshalled).
	Tx []byte `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	// The signature authenticating the transaction.
	Sig                  *Signature `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SignedBundle) Reset()         { *m = SignedBundle{} }
func (m *SignedBundle) String() string { return proto.CompactTextString(m) }
func (*SignedBundle) ProtoMessage()    {}
func (*SignedBundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb6b8173ee11af27, []int{2}
}

func (m *SignedBundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedBundle.Unmarshal(m, b)
}
func (m *SignedBundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedBundle.Marshal(b, m, deterministic)
}
func (m *SignedBundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedBundle.Merge(m, src)
}
func (m *SignedBundle) XXX_Size() int {
	return xxx_messageInfo_SignedBundle.Size(m)
}
func (m *SignedBundle) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedBundle.DiscardUnknown(m)
}

var xxx_messageInfo_SignedBundle proto.InternalMessageInfo

func (m *SignedBundle) GetTx() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *SignedBundle) GetSig() *Signature {
	if m != nil {
		return m.Sig
	}
	return nil
}

// Represents Vega domain specific error information over gRPC/Protobuf.
type ErrorDetail struct {
	// A Vega API domain specific unique error code, useful for client side mappings. e.g. 10004
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// A message that describes the error in more detail, should describe the problem encountered.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Any inner error information that could add more context, or be helpful for error reporting.
	Inner                string   `protobuf:"bytes,3,opt,name=inner,proto3" json:"inner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorDetail) Reset()         { *m = ErrorDetail{} }
func (m *ErrorDetail) String() string { return proto.CompactTextString(m) }
func (*ErrorDetail) ProtoMessage()    {}
func (*ErrorDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb6b8173ee11af27, []int{3}
}

func (m *ErrorDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorDetail.Unmarshal(m, b)
}
func (m *ErrorDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorDetail.Marshal(b, m, deterministic)
}
func (m *ErrorDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorDetail.Merge(m, src)
}
func (m *ErrorDetail) XXX_Size() int {
	return xxx_messageInfo_ErrorDetail.Size(m)
}
func (m *ErrorDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorDetail proto.InternalMessageInfo

func (m *ErrorDetail) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ErrorDetail) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ErrorDetail) GetInner() string {
	if m != nil {
		return m.Inner
	}
	return ""
}

func init() {
	proto.RegisterType((*Transaction)(nil), "vega.Transaction")
	proto.RegisterType((*Signature)(nil), "vega.Signature")
	proto.RegisterType((*SignedBundle)(nil), "vega.SignedBundle")
	proto.RegisterType((*ErrorDetail)(nil), "vega.ErrorDetail")
}

func init() { proto.RegisterFile("proto/vega.proto", fileDescriptor_bb6b8173ee11af27) }

var fileDescriptor_bb6b8173ee11af27 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x51, 0xc1, 0x6a, 0x2a, 0x31,
	0x14, 0x7d, 0xa3, 0xa3, 0x32, 0x77, 0xe4, 0x3d, 0x09, 0x6f, 0x31, 0x8f, 0x47, 0x41, 0x87, 0x52,
	0xdc, 0x54, 0xa1, 0xfd, 0x82, 0x8a, 0x05, 0xc1, 0x55, 0xa7, 0x5d, 0x75, 0x23, 0x71, 0xe6, 0x36,
	0x86, 0xc6, 0x44, 0x92, 0x8c, 0xb5, 0x1f, 0xd4, 0x8f, 0x6b, 0xbf, 0xa2, 0xe4, 0x3a, 0xd3, 0x55,
	0xce, 0x39, 0x97, 0x9c, 0x73, 0xb8, 0x17, 0x46, 0x07, 0x6b, 0xbc, 0x99, 0x1f, 0x51, 0xf0, 0x19,
	0x41, 0x16, 0x07, 0x9c, 0x7f, 0x44, 0x90, 0x3e, 0x59, 0xae, 0x1d, 0x2f, 0xbd, 0x34, 0x9a, 0x5d,
	0x00, 0x48, 0x7d, 0xa8, 0xfd, 0xa6, 0xe2, 0x9e, 0x67, 0xd1, 0x38, 0x9a, 0x0e, 0x8b, 0x84, 0x94,
	0x25, 0xf7, 0x9c, 0xfd, 0x85, 0x9e, 0x36, 0xba, 0xc4, 0xac, 0x33, 0x8e, 0xa6, 0x71, 0x71, 0x26,
	0x6c, 0x02, 0xc3, 0xad, 0x32, 0xe5, 0xeb, 0x66, 0x87, 0x52, 0xec, 0x7c, 0xd6, 0xa5, 0x61, 0x4a,
	0xda, 0x8a, 0x24, 0xf6, 0x1f, 0x06, 0xbc, 0xaa, 0x2c, 0x3a, 0x97, 0x7d, 0x0e, 0x82, 0xeb, 0xea,
	0x57, 0xd1, 0x2a, 0xec, 0x1f, 0xf4, 0x0f, 0xf5, 0x76, 0x8d, 0xef, 0xd9, 0x57, 0x3b, 0x6b, 0x84,
	0x45, 0x1f, 0xe2, 0x17, 0x6b, 0xf6, 0xf9, 0x1a, 0x92, 0x47, 0x29, 0x34, 0xf7, 0xb5, 0x45, 0x36,
	0x82, 0xae, 0x93, 0xa2, 0x69, 0x17, 0x20, 0x63, 0x10, 0x73, 0x25, 0x0c, 0xd5, 0x4a, 0x0a, 0xc2,
	0x2c, 0x83, 0xc1, 0x11, 0xad, 0x93, 0x46, 0x37, 0x85, 0x5a, 0x9a, 0xdf, 0xc1, 0x30, 0x98, 0x61,
	0xb5, 0xa8, 0x75, 0xa5, 0x90, 0xfd, 0x86, 0x8e, 0x3f, 0x35, 0x76, 0x1d, 0x7f, 0x62, 0x93, 0xb3,
	0x7f, 0x30, 0x4b, 0x6f, 0xfe, 0xcc, 0x68, 0x69, 0x3f, 0xe9, 0x14, 0x98, 0x3f, 0x40, 0x7a, 0x6f,
	0xad, 0xb1, 0x4b, 0xf4, 0x5c, 0xaa, 0x90, 0x5f, 0x9a, 0x0a, 0xc9, 0xa3, 0x57, 0x10, 0x0e, 0xf9,
	0x7b, 0x74, 0x8e, 0x0b, 0x6c, 0x6a, 0xb5, 0x34, 0x6c, 0x51, 0x6a, 0x8d, 0x96, 0x7a, 0x25, 0xc5,
	0x99, 0x2c, 0xae, 0x9e, 0x2f, 0xc3, 0x3f, 0x8a, 0xa3, 0x13, 0x95, 0x46, 0xcd, 0xa4, 0x99, 0x0b,
	0x73, 0xfd, 0xc6, 0x95, 0x42, 0x3f, 0x27, 0x75, 0xdb, 0xa7, 0xe7, 0xf6, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x74, 0x67, 0x43, 0xd1, 0xd3, 0x01, 0x00, 0x00,
}
