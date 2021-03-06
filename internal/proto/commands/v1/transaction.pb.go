// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/commands/v1/transaction.proto

package v1

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

type InputData struct {
	// A random number used to provided uniqueness and prevents against replay
	// attack.
	Nonce uint64 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// The block height associated to the transaction.
	// This should always be current height of the node at the time of sending the
	// Tx. BlockHeight is used as a mechanism for replay protection.
	BlockHeight uint64 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// Types that are valid to be assigned to Command:
	//	*InputData_OrderSubmission
	//	*InputData_OrderCancellation
	//	*InputData_OrderAmendment
	//	*InputData_WithdrawSubmission
	//	*InputData_ProposalSubmission
	//	*InputData_VoteSubmission
	//	*InputData_LiquidityProvisionSubmission
	//	*InputData_NodeRegistration
	//	*InputData_NodeVote
	//	*InputData_NodeSignature
	//	*InputData_ChainEvent
	//	*InputData_OracleDataSubmission
	Command              isInputData_Command `protobuf_oneof:"command"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *InputData) Reset()         { *m = InputData{} }
func (m *InputData) String() string { return proto.CompactTextString(m) }
func (*InputData) ProtoMessage()    {}
func (*InputData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1ec3aa815c39ed0, []int{0}
}

func (m *InputData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputData.Unmarshal(m, b)
}
func (m *InputData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputData.Marshal(b, m, deterministic)
}
func (m *InputData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputData.Merge(m, src)
}
func (m *InputData) XXX_Size() int {
	return xxx_messageInfo_InputData.Size(m)
}
func (m *InputData) XXX_DiscardUnknown() {
	xxx_messageInfo_InputData.DiscardUnknown(m)
}

var xxx_messageInfo_InputData proto.InternalMessageInfo

func (m *InputData) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *InputData) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

type isInputData_Command interface {
	isInputData_Command()
}

type InputData_OrderSubmission struct {
	OrderSubmission *OrderSubmission `protobuf:"bytes,1001,opt,name=order_submission,json=orderSubmission,proto3,oneof"`
}

type InputData_OrderCancellation struct {
	OrderCancellation *OrderCancellation `protobuf:"bytes,1002,opt,name=order_cancellation,json=orderCancellation,proto3,oneof"`
}

type InputData_OrderAmendment struct {
	OrderAmendment *OrderAmendment `protobuf:"bytes,1003,opt,name=order_amendment,json=orderAmendment,proto3,oneof"`
}

type InputData_WithdrawSubmission struct {
	WithdrawSubmission *WithdrawSubmission `protobuf:"bytes,1004,opt,name=withdraw_submission,json=withdrawSubmission,proto3,oneof"`
}

type InputData_ProposalSubmission struct {
	ProposalSubmission *ProposalSubmission `protobuf:"bytes,1005,opt,name=proposal_submission,json=proposalSubmission,proto3,oneof"`
}

type InputData_VoteSubmission struct {
	VoteSubmission *VoteSubmission `protobuf:"bytes,1006,opt,name=vote_submission,json=voteSubmission,proto3,oneof"`
}

type InputData_LiquidityProvisionSubmission struct {
	LiquidityProvisionSubmission *LiquidityProvisionSubmission `protobuf:"bytes,1007,opt,name=liquidity_provision_submission,json=liquidityProvisionSubmission,proto3,oneof"`
}

type InputData_NodeRegistration struct {
	NodeRegistration *NodeRegistration `protobuf:"bytes,2001,opt,name=node_registration,json=nodeRegistration,proto3,oneof"`
}

type InputData_NodeVote struct {
	NodeVote *NodeVote `protobuf:"bytes,2002,opt,name=node_vote,json=nodeVote,proto3,oneof"`
}

type InputData_NodeSignature struct {
	NodeSignature *NodeSignature `protobuf:"bytes,2003,opt,name=node_signature,json=nodeSignature,proto3,oneof"`
}

type InputData_ChainEvent struct {
	ChainEvent *ChainEvent `protobuf:"bytes,2004,opt,name=chain_event,json=chainEvent,proto3,oneof"`
}

type InputData_OracleDataSubmission struct {
	OracleDataSubmission *OracleDataSubmission `protobuf:"bytes,3001,opt,name=oracle_data_submission,json=oracleDataSubmission,proto3,oneof"`
}

func (*InputData_OrderSubmission) isInputData_Command() {}

func (*InputData_OrderCancellation) isInputData_Command() {}

func (*InputData_OrderAmendment) isInputData_Command() {}

func (*InputData_WithdrawSubmission) isInputData_Command() {}

func (*InputData_ProposalSubmission) isInputData_Command() {}

func (*InputData_VoteSubmission) isInputData_Command() {}

func (*InputData_LiquidityProvisionSubmission) isInputData_Command() {}

func (*InputData_NodeRegistration) isInputData_Command() {}

func (*InputData_NodeVote) isInputData_Command() {}

func (*InputData_NodeSignature) isInputData_Command() {}

func (*InputData_ChainEvent) isInputData_Command() {}

func (*InputData_OracleDataSubmission) isInputData_Command() {}

func (m *InputData) GetCommand() isInputData_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *InputData) GetOrderSubmission() *OrderSubmission {
	if x, ok := m.GetCommand().(*InputData_OrderSubmission); ok {
		return x.OrderSubmission
	}
	return nil
}

func (m *InputData) GetOrderCancellation() *OrderCancellation {
	if x, ok := m.GetCommand().(*InputData_OrderCancellation); ok {
		return x.OrderCancellation
	}
	return nil
}

func (m *InputData) GetOrderAmendment() *OrderAmendment {
	if x, ok := m.GetCommand().(*InputData_OrderAmendment); ok {
		return x.OrderAmendment
	}
	return nil
}

func (m *InputData) GetWithdrawSubmission() *WithdrawSubmission {
	if x, ok := m.GetCommand().(*InputData_WithdrawSubmission); ok {
		return x.WithdrawSubmission
	}
	return nil
}

func (m *InputData) GetProposalSubmission() *ProposalSubmission {
	if x, ok := m.GetCommand().(*InputData_ProposalSubmission); ok {
		return x.ProposalSubmission
	}
	return nil
}

func (m *InputData) GetVoteSubmission() *VoteSubmission {
	if x, ok := m.GetCommand().(*InputData_VoteSubmission); ok {
		return x.VoteSubmission
	}
	return nil
}

func (m *InputData) GetLiquidityProvisionSubmission() *LiquidityProvisionSubmission {
	if x, ok := m.GetCommand().(*InputData_LiquidityProvisionSubmission); ok {
		return x.LiquidityProvisionSubmission
	}
	return nil
}

func (m *InputData) GetNodeRegistration() *NodeRegistration {
	if x, ok := m.GetCommand().(*InputData_NodeRegistration); ok {
		return x.NodeRegistration
	}
	return nil
}

func (m *InputData) GetNodeVote() *NodeVote {
	if x, ok := m.GetCommand().(*InputData_NodeVote); ok {
		return x.NodeVote
	}
	return nil
}

func (m *InputData) GetNodeSignature() *NodeSignature {
	if x, ok := m.GetCommand().(*InputData_NodeSignature); ok {
		return x.NodeSignature
	}
	return nil
}

func (m *InputData) GetChainEvent() *ChainEvent {
	if x, ok := m.GetCommand().(*InputData_ChainEvent); ok {
		return x.ChainEvent
	}
	return nil
}

func (m *InputData) GetOracleDataSubmission() *OracleDataSubmission {
	if x, ok := m.GetCommand().(*InputData_OracleDataSubmission); ok {
		return x.OracleDataSubmission
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*InputData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*InputData_OrderSubmission)(nil),
		(*InputData_OrderCancellation)(nil),
		(*InputData_OrderAmendment)(nil),
		(*InputData_WithdrawSubmission)(nil),
		(*InputData_ProposalSubmission)(nil),
		(*InputData_VoteSubmission)(nil),
		(*InputData_LiquidityProvisionSubmission)(nil),
		(*InputData_NodeRegistration)(nil),
		(*InputData_NodeVote)(nil),
		(*InputData_NodeSignature)(nil),
		(*InputData_ChainEvent)(nil),
		(*InputData_OracleDataSubmission)(nil),
	}
}

// Represents a transaction to be sent to Vega.
type Transaction struct {
	// One of the set of Vega commands (proto marshalled).
	InputData []byte `protobuf:"bytes,1,opt,name=input_data,json=inputData,proto3" json:"input_data,omitempty"`
	// The signature of the inputData.
	Signature *Signature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	// The sender of the transaction.
	// Any of the following would be valid:
	//
	// Types that are valid to be assigned to From:
	//	*Transaction_Address
	//	*Transaction_PubKey
	From isTransaction_From `protobuf_oneof:"from"`
	// A version of the transaction, to be used in the future in case we want to
	// implement changes to the Transaction format.
	Version              uint32   `protobuf:"varint,2000,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1ec3aa815c39ed0, []int{1}
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

func (m *Transaction) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type isTransaction_From interface {
	isTransaction_From()
}

type Transaction_Address struct {
	Address string `protobuf:"bytes,1001,opt,name=address,proto3,oneof"`
}

type Transaction_PubKey struct {
	PubKey string `protobuf:"bytes,1002,opt,name=pub_key,json=pubKey,proto3,oneof"`
}

func (*Transaction_Address) isTransaction_From() {}

func (*Transaction_PubKey) isTransaction_From() {}

func (m *Transaction) GetFrom() isTransaction_From {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Transaction) GetAddress() string {
	if x, ok := m.GetFrom().(*Transaction_Address); ok {
		return x.Address
	}
	return ""
}

func (m *Transaction) GetPubKey() string {
	if x, ok := m.GetFrom().(*Transaction_PubKey); ok {
		return x.PubKey
	}
	return ""
}

func (m *Transaction) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Transaction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Transaction_Address)(nil),
		(*Transaction_PubKey)(nil),
	}
}

// A signature to be authenticate a transaction and to be verified by the vega
// network.
type Signature struct {
	// The bytes of the signature (hex-encoded).
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// The algorithm used to create the signature.
	Algo string `protobuf:"bytes,2,opt,name=algo,proto3" json:"algo,omitempty"`
	// The version of the signature used to create the signature.
	Version              uint32   `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1ec3aa815c39ed0, []int{2}
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

func (m *Signature) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Signature) GetAlgo() string {
	if m != nil {
		return m.Algo
	}
	return ""
}

func (m *Signature) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*InputData)(nil), "wallet.vega.commands.v1.InputData")
	proto.RegisterType((*Transaction)(nil), "wallet.vega.commands.v1.Transaction")
	proto.RegisterType((*Signature)(nil), "wallet.vega.commands.v1.Signature")
}

func init() {
	proto.RegisterFile("proto/commands/v1/transaction.proto", fileDescriptor_c1ec3aa815c39ed0)
}

var fileDescriptor_c1ec3aa815c39ed0 = []byte{
	// 704 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdb, 0x6e, 0x13, 0x3b,
	0x14, 0x4d, 0x7a, 0xda, 0xe6, 0x8c, 0xd3, 0xab, 0x4f, 0x75, 0x3a, 0xa7, 0x87, 0x4b, 0x9b, 0x4a,
	0x50, 0x8a, 0x9a, 0xa8, 0x20, 0x1e, 0x10, 0x2f, 0xb4, 0x05, 0x31, 0x08, 0x44, 0x2b, 0x17, 0x15,
	0x81, 0x90, 0x46, 0xce, 0x8c, 0x9b, 0x58, 0x75, 0xec, 0xc1, 0xe3, 0x99, 0xa8, 0x2f, 0x7c, 0x04,
	0x7f, 0xc3, 0x2f, 0xf0, 0xc4, 0xed, 0x27, 0xb8, 0xff, 0x02, 0xb2, 0x9d, 0x49, 0x66, 0x1a, 0x92,
	0xbe, 0x8d, 0xd7, 0xde, 0x6b, 0x2d, 0x7b, 0xdb, 0xb3, 0xc0, 0x7a, 0x24, 0x85, 0x12, 0x8d, 0x40,
	0x74, 0x3a, 0x98, 0x87, 0x71, 0x23, 0xdd, 0x6e, 0x28, 0x89, 0x79, 0x8c, 0x03, 0x45, 0x05, 0xaf,
	0x9b, 0x2a, 0x5c, 0xee, 0x62, 0xc6, 0x88, 0xaa, 0xa7, 0xa4, 0x85, 0xeb, 0x59, 0x6b, 0x3d, 0xdd,
	0x5e, 0x59, 0x1d, 0x66, 0xf7, 0xcb, 0xa6, 0xb4, 0xb2, 0x39, 0xdc, 0x91, 0x62, 0x46, 0x43, 0xac,
	0x84, 0xf4, 0xcf, 0xf4, 0x5e, 0x1e, 0xee, 0x15, 0x12, 0x07, 0x8c, 0xf4, 0x1a, 0x6a, 0x6f, 0x1c,
	0xe0, 0x3c, 0xe4, 0x51, 0xa2, 0xee, 0x61, 0x85, 0xe1, 0x12, 0x98, 0xe2, 0x82, 0x07, 0xc4, 0x2d,
	0xaf, 0x96, 0x37, 0x26, 0x91, 0x5d, 0xc0, 0x35, 0x30, 0xd3, 0x64, 0x22, 0x38, 0xf1, 0xdb, 0x84,
	0xb6, 0xda, 0xca, 0x9d, 0x30, 0xc5, 0xaa, 0xc1, 0x3c, 0x03, 0xc1, 0x23, 0xb0, 0x20, 0x64, 0x48,
	0xa4, 0x1f, 0x27, 0xcd, 0x0e, 0x8d, 0x63, 0x2a, 0xb8, 0xfb, 0xa5, 0xb2, 0x5a, 0xde, 0xa8, 0xde,
	0xd8, 0xa8, 0x8f, 0x38, 0x6a, 0x7d, 0x5f, 0x33, 0x0e, 0xfb, 0x04, 0xaf, 0x84, 0xe6, 0x45, 0x11,
	0x82, 0x2f, 0x01, 0xb4, 0xba, 0x01, 0xe6, 0x01, 0x61, 0x0c, 0xeb, 0x11, 0xba, 0x5f, 0xad, 0xf2,
	0xe6, 0x78, 0xe5, 0xbd, 0x1c, 0xc5, 0x2b, 0xa1, 0x45, 0x71, 0x16, 0x84, 0x87, 0xc0, 0x1a, 0xfa,
	0xb8, 0x43, 0x78, 0xd8, 0x21, 0x5c, 0xb9, 0xdf, 0xac, 0xf4, 0xd5, 0xf1, 0xd2, 0x3b, 0x59, 0xbf,
	0x57, 0x42, 0x73, 0xa2, 0x80, 0x40, 0x1f, 0xfc, 0xd3, 0xa5, 0xaa, 0x1d, 0x4a, 0xdc, 0xcd, 0x4f,
	0xe3, 0xbb, 0x15, 0xbe, 0x3e, 0x52, 0xf8, 0x59, 0x8f, 0x54, 0x18, 0x08, 0xec, 0x0e, 0xa1, 0xda,
	0x20, 0x92, 0x22, 0x12, 0x31, 0x66, 0x79, 0x83, 0x1f, 0xe7, 0x19, 0x1c, 0xf4, 0x48, 0x45, 0x83,
	0x68, 0x08, 0xd5, 0x63, 0x49, 0x85, 0x22, 0x79, 0xf1, 0x9f, 0xe7, 0x8d, 0xe5, 0x48, 0x28, 0x52,
	0x10, 0x9e, 0x4b, 0x0b, 0x08, 0x7c, 0x0d, 0x2e, 0x31, 0xfa, 0x2a, 0xa1, 0x21, 0x55, 0xa7, 0x7e,
	0x24, 0x45, 0x4a, 0x35, 0x9c, 0xf7, 0xf8, 0x65, 0x3d, 0x6e, 0x8d, 0xf4, 0x78, 0x9c, 0xf1, 0x0f,
	0x32, 0x7a, 0xc1, 0xf1, 0x02, 0x1b, 0x53, 0x87, 0xcf, 0xc1, 0x22, 0x17, 0x21, 0xf1, 0x25, 0x69,
	0xd1, 0x58, 0x49, 0xfb, 0x90, 0x3e, 0xcc, 0x1b, 0xcb, 0x6b, 0x23, 0x2d, 0x9f, 0x88, 0x90, 0xa0,
	0x1c, 0xc3, 0x2b, 0xa1, 0x05, 0x7e, 0x06, 0x83, 0x3b, 0xc0, 0x31, 0xd2, 0xfa, 0xc4, 0xee, 0x47,
	0x2b, 0xb9, 0x36, 0x56, 0x52, 0x4f, 0xcb, 0x2b, 0xa1, 0xbf, 0x79, 0xef, 0x1b, 0x1e, 0x80, 0x39,
	0x23, 0x11, 0xd3, 0x16, 0xc7, 0x2a, 0x91, 0xc4, 0xfd, 0x64, 0x75, 0xae, 0x8c, 0xd5, 0x39, 0xcc,
	0xda, 0xbd, 0x12, 0x9a, 0xe5, 0x79, 0x00, 0x3e, 0x00, 0xd5, 0xa0, 0x8d, 0x29, 0xf7, 0x49, 0xaa,
	0xdf, 0xf5, 0x67, 0x2b, 0xb7, 0x3e, 0x52, 0x6e, 0x4f, 0x37, 0xdf, 0x4f, 0xed, 0x9b, 0x06, 0x41,
	0x7f, 0x05, 0x8f, 0xc1, 0xbf, 0x36, 0x32, 0xfc, 0x10, 0x2b, 0x9c, 0xbf, 0xb0, 0xb7, 0xcb, 0x46,
	0x73, 0x6b, 0xcc, 0xbf, 0xa2, 0x79, 0x3a, 0x59, 0x0a, 0x17, 0xb5, 0x24, 0xfe, 0x80, 0xef, 0x3a,
	0xa0, 0xd2, 0xe3, 0xd6, 0xde, 0x95, 0x41, 0xf5, 0xe9, 0x20, 0x32, 0xe1, 0x45, 0x00, 0xa8, 0xce,
	0x28, 0xb3, 0x03, 0x93, 0x4d, 0x33, 0xc8, 0xa1, 0xfd, 0xd4, 0xba, 0x0b, 0x9c, 0xc1, 0xdc, 0x26,
	0xcc, 0x9e, 0x6a, 0x23, 0xf7, 0xd4, 0x9f, 0x10, 0x1a, 0x90, 0xe0, 0xff, 0xa0, 0x82, 0xc3, 0x50,
	0x92, 0x38, 0xb6, 0xa9, 0xe5, 0x78, 0x25, 0x94, 0x21, 0x70, 0x05, 0x54, 0xa2, 0xa4, 0xe9, 0x9f,
	0x90, 0x53, 0x1b, 0x3c, 0xba, 0x38, 0x1d, 0x25, 0xcd, 0x47, 0xe4, 0x14, 0xfe, 0x07, 0x2a, 0x29,
	0x91, 0x66, 0x1a, 0xef, 0xf5, 0x84, 0x67, 0x51, 0xb6, 0xde, 0x9d, 0x06, 0x93, 0xc7, 0x52, 0x74,
	0x6a, 0xfb, 0xc0, 0x19, 0xdc, 0xca, 0x12, 0x98, 0x4a, 0x31, 0x4b, 0x6c, 0xc0, 0x3a, 0xc8, 0x2e,
	0x20, 0x04, 0x93, 0x98, 0xb5, 0x84, 0xd9, 0xbb, 0x83, 0xcc, 0x37, 0x74, 0x07, 0xca, 0x7f, 0x15,
	0x85, 0xef, 0xbc, 0xb8, 0x1d, 0x88, 0x90, 0x98, 0xa3, 0x99, 0x10, 0x0f, 0x04, 0xab, 0x53, 0xd1,
	0x68, 0x89, 0x2d, 0x7b, 0xe8, 0x06, 0xe5, 0x8a, 0x48, 0x8e, 0x59, 0x63, 0x28, 0xfd, 0x9b, 0xd3,
	0x06, 0xba, 0xf9, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xd6, 0x43, 0xd4, 0xa5, 0x06, 0x00, 0x00,
}
