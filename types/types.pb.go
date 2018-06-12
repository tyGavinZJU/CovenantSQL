// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package types

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

type TxType int32

const (
	TxType_QUERY   TxType = 0
	TxType_STORAGE TxType = 1
)

var TxType_name = map[int32]string{
	0: "QUERY",
	1: "STORAGE",
}
var TxType_value = map[string]int32{
	"QUERY":   0,
	"STORAGE": 1,
}

func (x TxType) String() string {
	return proto.EnumName(TxType_name, int32(x))
}
func (TxType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_types_15b9862bddc8f363, []int{0}
}

type Signature struct {
	R                    string   `protobuf:"bytes,1,opt,name=R,proto3" json:"R,omitempty"`
	S                    string   `protobuf:"bytes,2,opt,name=S,proto3" json:"S,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_15b9862bddc8f363, []int{0}
}
func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (dst *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(dst, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetR() string {
	if m != nil {
		return m.R
	}
	return ""
}

func (m *Signature) GetS() string {
	if m != nil {
		return m.S
	}
	return ""
}

type PublicKey struct {
	PublicKey            []byte   `protobuf:"bytes,1,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicKey) Reset()         { *m = PublicKey{} }
func (m *PublicKey) String() string { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()    {}
func (*PublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_15b9862bddc8f363, []int{1}
}
func (m *PublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicKey.Unmarshal(m, b)
}
func (m *PublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicKey.Marshal(b, m, deterministic)
}
func (dst *PublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKey.Merge(dst, src)
}
func (m *PublicKey) XXX_Size() int {
	return xxx_messageInfo_PublicKey.Size(m)
}
func (m *PublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKey proto.InternalMessageInfo

func (m *PublicKey) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type Hash struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hash) Reset()         { *m = Hash{} }
func (m *Hash) String() string { return proto.CompactTextString(m) }
func (*Hash) ProtoMessage()    {}
func (*Hash) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_15b9862bddc8f363, []int{2}
}
func (m *Hash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hash.Unmarshal(m, b)
}
func (m *Hash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hash.Marshal(b, m, deterministic)
}
func (dst *Hash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hash.Merge(dst, src)
}
func (m *Hash) XXX_Size() int {
	return xxx_messageInfo_Hash.Size(m)
}
func (m *Hash) XXX_DiscardUnknown() {
	xxx_messageInfo_Hash.DiscardUnknown(m)
}

var xxx_messageInfo_Hash proto.InternalMessageInfo

func (m *Hash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type UtxoEntry struct {
	IsCoinbase           bool             `protobuf:"varint,1,opt,name=IsCoinbase,proto3" json:"IsCoinbase,omitempty"`
	FromMainChain        bool             `protobuf:"varint,2,opt,name=FromMainChain,proto3" json:"FromMainChain,omitempty"`
	BlockHeight          uint32           `protobuf:"varint,3,opt,name=BlockHeight,proto3" json:"BlockHeight,omitempty"`
	SparseOutputs        map[uint32]*Utxo `protobuf:"bytes,4,rep,name=SparseOutputs,proto3" json:"SparseOutputs,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UtxoEntry) Reset()         { *m = UtxoEntry{} }
func (m *UtxoEntry) String() string { return proto.CompactTextString(m) }
func (*UtxoEntry) ProtoMessage()    {}
func (*UtxoEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_15b9862bddc8f363, []int{3}
}
func (m *UtxoEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UtxoEntry.Unmarshal(m, b)
}
func (m *UtxoEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UtxoEntry.Marshal(b, m, deterministic)
}
func (dst *UtxoEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UtxoEntry.Merge(dst, src)
}
func (m *UtxoEntry) XXX_Size() int {
	return xxx_messageInfo_UtxoEntry.Size(m)
}
func (m *UtxoEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_UtxoEntry.DiscardUnknown(m)
}

var xxx_messageInfo_UtxoEntry proto.InternalMessageInfo

func (m *UtxoEntry) GetIsCoinbase() bool {
	if m != nil {
		return m.IsCoinbase
	}
	return false
}

func (m *UtxoEntry) GetFromMainChain() bool {
	if m != nil {
		return m.FromMainChain
	}
	return false
}

func (m *UtxoEntry) GetBlockHeight() uint32 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *UtxoEntry) GetSparseOutputs() map[uint32]*Utxo {
	if m != nil {
		return m.SparseOutputs
	}
	return nil
}

type Utxo struct {
	UtxoHeader           *UtxoHeader `protobuf:"bytes,1,opt,name=UtxoHeader,proto3" json:"UtxoHeader,omitempty"`
	Spent                bool        `protobuf:"varint,2,opt,name=Spent,proto3" json:"Spent,omitempty"`
	Amount               uint64      `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Utxo) Reset()         { *m = Utxo{} }
func (m *Utxo) String() string { return proto.CompactTextString(m) }
func (*Utxo) ProtoMessage()    {}
func (*Utxo) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_15b9862bddc8f363, []int{4}
}
func (m *Utxo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Utxo.Unmarshal(m, b)
}
func (m *Utxo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Utxo.Marshal(b, m, deterministic)
}
func (dst *Utxo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Utxo.Merge(dst, src)
}
func (m *Utxo) XXX_Size() int {
	return xxx_messageInfo_Utxo.Size(m)
}
func (m *Utxo) XXX_DiscardUnknown() {
	xxx_messageInfo_Utxo.DiscardUnknown(m)
}

var xxx_messageInfo_Utxo proto.InternalMessageInfo

func (m *Utxo) GetUtxoHeader() *UtxoHeader {
	if m != nil {
		return m.UtxoHeader
	}
	return nil
}

func (m *Utxo) GetSpent() bool {
	if m != nil {
		return m.Spent
	}
	return false
}

func (m *Utxo) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type UtxoHeader struct {
	Version              int32      `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	PrevTxHash           *Hash      `protobuf:"bytes,2,opt,name=PrevTxHash,proto3" json:"PrevTxHash,omitempty"`
	Signee               *PublicKey `protobuf:"bytes,3,opt,name=Signee,proto3" json:"Signee,omitempty"`
	Signature            *Signature `protobuf:"bytes,4,opt,name=Signature,proto3" json:"Signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UtxoHeader) Reset()         { *m = UtxoHeader{} }
func (m *UtxoHeader) String() string { return proto.CompactTextString(m) }
func (*UtxoHeader) ProtoMessage()    {}
func (*UtxoHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_15b9862bddc8f363, []int{5}
}
func (m *UtxoHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UtxoHeader.Unmarshal(m, b)
}
func (m *UtxoHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UtxoHeader.Marshal(b, m, deterministic)
}
func (dst *UtxoHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UtxoHeader.Merge(dst, src)
}
func (m *UtxoHeader) XXX_Size() int {
	return xxx_messageInfo_UtxoHeader.Size(m)
}
func (m *UtxoHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_UtxoHeader.DiscardUnknown(m)
}

var xxx_messageInfo_UtxoHeader proto.InternalMessageInfo

func (m *UtxoHeader) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *UtxoHeader) GetPrevTxHash() *Hash {
	if m != nil {
		return m.PrevTxHash
	}
	return nil
}

func (m *UtxoHeader) GetSignee() *PublicKey {
	if m != nil {
		return m.Signee
	}
	return nil
}

func (m *UtxoHeader) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Tx struct {
	UtxoIn               []*Utxo  `protobuf:"bytes,1,rep,name=UtxoIn,proto3" json:"UtxoIn,omitempty"`
	UtxoOut              []*Utxo  `protobuf:"bytes,2,rep,name=UtxoOut,proto3" json:"UtxoOut,omitempty"`
	Type                 TxType   `protobuf:"varint,3,opt,name=type,proto3,enum=types.TxType" json:"type,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tx) Reset()         { *m = Tx{} }
func (m *Tx) String() string { return proto.CompactTextString(m) }
func (*Tx) ProtoMessage()    {}
func (*Tx) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_15b9862bddc8f363, []int{6}
}
func (m *Tx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tx.Unmarshal(m, b)
}
func (m *Tx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tx.Marshal(b, m, deterministic)
}
func (dst *Tx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tx.Merge(dst, src)
}
func (m *Tx) XXX_Size() int {
	return xxx_messageInfo_Tx.Size(m)
}
func (m *Tx) XXX_DiscardUnknown() {
	xxx_messageInfo_Tx.DiscardUnknown(m)
}

var xxx_messageInfo_Tx proto.InternalMessageInfo

func (m *Tx) GetUtxoIn() []*Utxo {
	if m != nil {
		return m.UtxoIn
	}
	return nil
}

func (m *Tx) GetUtxoOut() []*Utxo {
	if m != nil {
		return m.UtxoOut
	}
	return nil
}

func (m *Tx) GetType() TxType {
	if m != nil {
		return m.Type
	}
	return TxType_QUERY
}

func (m *Tx) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type NodeID struct {
	NodeID               string   `protobuf:"bytes,1,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeID) Reset()         { *m = NodeID{} }
func (m *NodeID) String() string { return proto.CompactTextString(m) }
func (*NodeID) ProtoMessage()    {}
func (*NodeID) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_15b9862bddc8f363, []int{7}
}
func (m *NodeID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeID.Unmarshal(m, b)
}
func (m *NodeID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeID.Marshal(b, m, deterministic)
}
func (dst *NodeID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeID.Merge(dst, src)
}
func (m *NodeID) XXX_Size() int {
	return xxx_messageInfo_NodeID.Size(m)
}
func (m *NodeID) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeID.DiscardUnknown(m)
}

var xxx_messageInfo_NodeID proto.InternalMessageInfo

func (m *NodeID) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

type Header struct {
	Version              int32    `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Producer             *NodeID  `protobuf:"bytes,2,opt,name=Producer,proto3" json:"Producer,omitempty"`
	Root                 *Hash    `protobuf:"bytes,3,opt,name=Root,proto3" json:"Root,omitempty"`
	Parent               *Hash    `protobuf:"bytes,4,opt,name=Parent,proto3" json:"Parent,omitempty"`
	MerkleRoot           *Hash    `protobuf:"bytes,5,opt,name=MerkleRoot,proto3" json:"MerkleRoot,omitempty"`
	Timestamp            int64    `protobuf:"varint,6,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_15b9862bddc8f363, []int{8}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (dst *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(dst, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Header) GetProducer() *NodeID {
	if m != nil {
		return m.Producer
	}
	return nil
}

func (m *Header) GetRoot() *Hash {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *Header) GetParent() *Hash {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *Header) GetMerkleRoot() *Hash {
	if m != nil {
		return m.MerkleRoot
	}
	return nil
}

func (m *Header) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type SignedHeader struct {
	Header               *Header    `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	BlockHash            *Hash      `protobuf:"bytes,2,opt,name=BlockHash,proto3" json:"BlockHash,omitempty"`
	Signee               *PublicKey `protobuf:"bytes,3,opt,name=Signee,proto3" json:"Signee,omitempty"`
	Signature            *Signature `protobuf:"bytes,4,opt,name=Signature,proto3" json:"Signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SignedHeader) Reset()         { *m = SignedHeader{} }
func (m *SignedHeader) String() string { return proto.CompactTextString(m) }
func (*SignedHeader) ProtoMessage()    {}
func (*SignedHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_15b9862bddc8f363, []int{9}
}
func (m *SignedHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedHeader.Unmarshal(m, b)
}
func (m *SignedHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedHeader.Marshal(b, m, deterministic)
}
func (dst *SignedHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedHeader.Merge(dst, src)
}
func (m *SignedHeader) XXX_Size() int {
	return xxx_messageInfo_SignedHeader.Size(m)
}
func (m *SignedHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedHeader.DiscardUnknown(m)
}

var xxx_messageInfo_SignedHeader proto.InternalMessageInfo

func (m *SignedHeader) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SignedHeader) GetBlockHash() *Hash {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *SignedHeader) GetSignee() *PublicKey {
	if m != nil {
		return m.Signee
	}
	return nil
}

func (m *SignedHeader) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type State struct {
	Head                 *Hash    `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	Height               int32    `protobuf:"varint,2,opt,name=Height,proto3" json:"Height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_15b9862bddc8f363, []int{10}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (dst *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(dst, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetHead() *Hash {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *State) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

type BPTx struct {
	TxHash               *Hash     `protobuf:"bytes,1,opt,name=TxHash,proto3" json:"TxHash,omitempty"`
	Txdata               *BPTxData `protobuf:"bytes,2,opt,name=txdata,proto3" json:"txdata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BPTx) Reset()         { *m = BPTx{} }
func (m *BPTx) String() string { return proto.CompactTextString(m) }
func (*BPTx) ProtoMessage()    {}
func (*BPTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_15b9862bddc8f363, []int{11}
}
func (m *BPTx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BPTx.Unmarshal(m, b)
}
func (m *BPTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BPTx.Marshal(b, m, deterministic)
}
func (dst *BPTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BPTx.Merge(dst, src)
}
func (m *BPTx) XXX_Size() int {
	return xxx_messageInfo_BPTx.Size(m)
}
func (m *BPTx) XXX_DiscardUnknown() {
	xxx_messageInfo_BPTx.DiscardUnknown(m)
}

var xxx_messageInfo_BPTx proto.InternalMessageInfo

func (m *BPTx) GetTxHash() *Hash {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *BPTx) GetTxdata() *BPTxData {
	if m != nil {
		return m.Txdata
	}
	return nil
}

type BPTxData struct {
	AccountNonce         uint64     `protobuf:"varint,1,opt,name=AccountNonce,proto3" json:"AccountNonce,omitempty"`
	Recipient            *NodeID    `protobuf:"bytes,2,opt,name=Recipient,proto3" json:"Recipient,omitempty"`
	Amount               []byte     `protobuf:"bytes,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Payload              []byte     `protobuf:"bytes,4,opt,name=Payload,proto3" json:"Payload,omitempty"`
	Sig                  *Signature `protobuf:"bytes,5,opt,name=Sig,proto3" json:"Sig,omitempty"`
	PublicKey            *PublicKey `protobuf:"bytes,6,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BPTxData) Reset()         { *m = BPTxData{} }
func (m *BPTxData) String() string { return proto.CompactTextString(m) }
func (*BPTxData) ProtoMessage()    {}
func (*BPTxData) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_15b9862bddc8f363, []int{12}
}
func (m *BPTxData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BPTxData.Unmarshal(m, b)
}
func (m *BPTxData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BPTxData.Marshal(b, m, deterministic)
}
func (dst *BPTxData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BPTxData.Merge(dst, src)
}
func (m *BPTxData) XXX_Size() int {
	return xxx_messageInfo_BPTxData.Size(m)
}
func (m *BPTxData) XXX_DiscardUnknown() {
	xxx_messageInfo_BPTxData.DiscardUnknown(m)
}

var xxx_messageInfo_BPTxData proto.InternalMessageInfo

func (m *BPTxData) GetAccountNonce() uint64 {
	if m != nil {
		return m.AccountNonce
	}
	return 0
}

func (m *BPTxData) GetRecipient() *NodeID {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *BPTxData) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *BPTxData) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *BPTxData) GetSig() *Signature {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *BPTxData) GetPublicKey() *PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type BPHeader struct {
	Version              int32    `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Producer             *NodeID  `protobuf:"bytes,2,opt,name=Producer,proto3" json:"Producer,omitempty"`
	Root                 *Hash    `protobuf:"bytes,3,opt,name=Root,proto3" json:"Root,omitempty"`
	Parent               *Hash    `protobuf:"bytes,4,opt,name=Parent,proto3" json:"Parent,omitempty"`
	MerkleRoot           *Hash    `protobuf:"bytes,5,opt,name=MerkleRoot,proto3" json:"MerkleRoot,omitempty"`
	Timestamp            int64    `protobuf:"varint,6,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BPHeader) Reset()         { *m = BPHeader{} }
func (m *BPHeader) String() string { return proto.CompactTextString(m) }
func (*BPHeader) ProtoMessage()    {}
func (*BPHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_15b9862bddc8f363, []int{13}
}
func (m *BPHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BPHeader.Unmarshal(m, b)
}
func (m *BPHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BPHeader.Marshal(b, m, deterministic)
}
func (dst *BPHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BPHeader.Merge(dst, src)
}
func (m *BPHeader) XXX_Size() int {
	return xxx_messageInfo_BPHeader.Size(m)
}
func (m *BPHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BPHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BPHeader proto.InternalMessageInfo

func (m *BPHeader) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BPHeader) GetProducer() *NodeID {
	if m != nil {
		return m.Producer
	}
	return nil
}

func (m *BPHeader) GetRoot() *Hash {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *BPHeader) GetParent() *Hash {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *BPHeader) GetMerkleRoot() *Hash {
	if m != nil {
		return m.MerkleRoot
	}
	return nil
}

func (m *BPHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type BPSignedHeader struct {
	Header               *Header    `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	BlockHash            *Hash      `protobuf:"bytes,2,opt,name=BlockHash,proto3" json:"BlockHash,omitempty"`
	Signee               *PublicKey `protobuf:"bytes,3,opt,name=Signee,proto3" json:"Signee,omitempty"`
	Signature            *Signature `protobuf:"bytes,4,opt,name=Signature,proto3" json:"Signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BPSignedHeader) Reset()         { *m = BPSignedHeader{} }
func (m *BPSignedHeader) String() string { return proto.CompactTextString(m) }
func (*BPSignedHeader) ProtoMessage()    {}
func (*BPSignedHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_15b9862bddc8f363, []int{14}
}
func (m *BPSignedHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BPSignedHeader.Unmarshal(m, b)
}
func (m *BPSignedHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BPSignedHeader.Marshal(b, m, deterministic)
}
func (dst *BPSignedHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BPSignedHeader.Merge(dst, src)
}
func (m *BPSignedHeader) XXX_Size() int {
	return xxx_messageInfo_BPSignedHeader.Size(m)
}
func (m *BPSignedHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BPSignedHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BPSignedHeader proto.InternalMessageInfo

func (m *BPSignedHeader) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BPSignedHeader) GetBlockHash() *Hash {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *BPSignedHeader) GetSignee() *PublicKey {
	if m != nil {
		return m.Signee
	}
	return nil
}

func (m *BPSignedHeader) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type BPBlock struct {
	Header               *BPSignedHeader `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	Tx                   []*BPTx         `protobuf:"bytes,2,rep,name=Tx,proto3" json:"Tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BPBlock) Reset()         { *m = BPBlock{} }
func (m *BPBlock) String() string { return proto.CompactTextString(m) }
func (*BPBlock) ProtoMessage()    {}
func (*BPBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_15b9862bddc8f363, []int{15}
}
func (m *BPBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BPBlock.Unmarshal(m, b)
}
func (m *BPBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BPBlock.Marshal(b, m, deterministic)
}
func (dst *BPBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BPBlock.Merge(dst, src)
}
func (m *BPBlock) XXX_Size() int {
	return xxx_messageInfo_BPBlock.Size(m)
}
func (m *BPBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_BPBlock.DiscardUnknown(m)
}

var xxx_messageInfo_BPBlock proto.InternalMessageInfo

func (m *BPBlock) GetHeader() *BPSignedHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BPBlock) GetTx() []*BPTx {
	if m != nil {
		return m.Tx
	}
	return nil
}

func init() {
	proto.RegisterType((*Signature)(nil), "types.Signature")
	proto.RegisterType((*PublicKey)(nil), "types.PublicKey")
	proto.RegisterType((*Hash)(nil), "types.Hash")
	proto.RegisterType((*UtxoEntry)(nil), "types.UtxoEntry")
	proto.RegisterMapType((map[uint32]*Utxo)(nil), "types.UtxoEntry.SparseOutputsEntry")
	proto.RegisterType((*Utxo)(nil), "types.Utxo")
	proto.RegisterType((*UtxoHeader)(nil), "types.UtxoHeader")
	proto.RegisterType((*Tx)(nil), "types.Tx")
	proto.RegisterType((*NodeID)(nil), "types.NodeID")
	proto.RegisterType((*Header)(nil), "types.Header")
	proto.RegisterType((*SignedHeader)(nil), "types.SignedHeader")
	proto.RegisterType((*State)(nil), "types.State")
	proto.RegisterType((*BPTx)(nil), "types.BPTx")
	proto.RegisterType((*BPTxData)(nil), "types.BPTxData")
	proto.RegisterType((*BPHeader)(nil), "types.BPHeader")
	proto.RegisterType((*BPSignedHeader)(nil), "types.BPSignedHeader")
	proto.RegisterType((*BPBlock)(nil), "types.BPBlock")
	proto.RegisterEnum("types.TxType", TxType_name, TxType_value)
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_types_15b9862bddc8f363) }

var fileDescriptor_types_15b9862bddc8f363 = []byte{
	// 795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0x51, 0x6f, 0xe2, 0x46,
	0x10, 0xae, 0xc1, 0x18, 0x18, 0x43, 0x4a, 0x57, 0x6d, 0x85, 0xd2, 0xaa, 0x25, 0x4e, 0xa3, 0x90,
	0x46, 0x45, 0x2a, 0x7d, 0xa9, 0xfa, 0xd4, 0x90, 0xa4, 0x0d, 0xaa, 0x92, 0xb8, 0x6b, 0xe7, 0xa4,
	0x7b, 0xdc, 0xc0, 0x8a, 0x58, 0x01, 0xdb, 0xb2, 0xd7, 0x11, 0xfc, 0x88, 0xfc, 0x94, 0xfb, 0x03,
	0xa7, 0xfb, 0x2d, 0xa7, 0xbb, 0x7f, 0x72, 0x9a, 0xf5, 0x1a, 0xdb, 0x90, 0xe8, 0x9e, 0xee, 0xe1,
	0xee, 0x89, 0x9d, 0xf9, 0xc6, 0xbb, 0x33, 0xdf, 0x7c, 0x33, 0x02, 0x4c, 0xb1, 0x0a, 0x79, 0x3c,
	0x08, 0xa3, 0x40, 0x04, 0xa4, 0x26, 0x0d, 0xeb, 0x10, 0x9a, 0x8e, 0x37, 0xf3, 0x99, 0x48, 0x22,
	0x4e, 0x5a, 0xa0, 0xd1, 0xae, 0xd6, 0xd3, 0xfa, 0x4d, 0xaa, 0x51, 0xb4, 0x9c, 0x6e, 0x25, 0xb5,
	0x1c, 0xeb, 0x08, 0x9a, 0x76, 0x72, 0x3b, 0xf7, 0x26, 0xff, 0xf1, 0x15, 0xf9, 0xb1, 0x60, 0xc8,
	0x0f, 0x5a, 0x34, 0x77, 0x58, 0xbb, 0xa0, 0x5f, 0xb0, 0xf8, 0x8e, 0x90, 0xf4, 0x57, 0x05, 0xc8,
	0xb3, 0xf5, 0x58, 0x81, 0xe6, 0x8d, 0x58, 0x06, 0xe7, 0xbe, 0x88, 0x56, 0xe4, 0x27, 0x80, 0x71,
	0x7c, 0x1a, 0x78, 0xfe, 0x2d, 0x8b, 0xb9, 0x8c, 0x6b, 0xd0, 0x82, 0x87, 0xfc, 0x02, 0xed, 0x7f,
	0xa2, 0x60, 0x71, 0xc9, 0x3c, 0xff, 0xf4, 0x8e, 0x79, 0xbe, 0x4c, 0xa7, 0x41, 0xcb, 0x4e, 0xd2,
	0x03, 0x73, 0x34, 0x0f, 0x26, 0xf7, 0x17, 0xdc, 0x9b, 0xdd, 0x89, 0x6e, 0xb5, 0xa7, 0xf5, 0xdb,
	0xb4, 0xe8, 0x22, 0x63, 0x68, 0x3b, 0x21, 0x8b, 0x62, 0x7e, 0x9d, 0x88, 0x30, 0x11, 0x71, 0x57,
	0xef, 0x55, 0xfb, 0xe6, 0x70, 0x7f, 0x90, 0x32, 0xb2, 0x4e, 0x68, 0x50, 0x8a, 0x92, 0x2e, 0x5a,
	0xfe, 0x72, 0xf7, 0x12, 0xc8, 0x76, 0x10, 0xe9, 0x40, 0xf5, 0x5e, 0x51, 0xd1, 0xa6, 0x78, 0x24,
	0x7b, 0x50, 0x7b, 0x60, 0xf3, 0x84, 0xcb, 0x94, 0xcd, 0xa1, 0x59, 0x78, 0x8a, 0xa6, 0xc8, 0x5f,
	0x95, 0x3f, 0x35, 0x6b, 0x06, 0x3a, 0xba, 0xc8, 0xef, 0x00, 0xf8, 0x7b, 0xc1, 0xd9, 0x94, 0x47,
	0xf2, 0x1e, 0x73, 0xf8, 0x4d, 0xe1, 0x9b, 0x14, 0xa0, 0x85, 0x20, 0xf2, 0x2d, 0xd4, 0x9c, 0x90,
	0xfb, 0x42, 0x91, 0x92, 0x1a, 0xe4, 0x7b, 0x30, 0xd8, 0x22, 0x48, 0xfc, 0x94, 0x07, 0x9d, 0x2a,
	0xcb, 0x7a, 0xa5, 0x15, 0x5f, 0x20, 0x5d, 0xa8, 0xbf, 0xe0, 0x51, 0xec, 0x05, 0xbe, 0x7c, 0xac,
	0x46, 0x33, 0x93, 0x1c, 0x03, 0xd8, 0x11, 0x7f, 0x70, 0x97, 0xb2, 0x77, 0xe5, 0xec, 0xd1, 0x45,
	0x0b, 0x30, 0xe9, 0x83, 0x81, 0xf2, 0xe1, 0x5c, 0xbe, 0x66, 0x0e, 0x3b, 0x2a, 0x70, 0x2d, 0x06,
	0xaa, 0x70, 0x32, 0x28, 0x08, 0xad, 0xab, 0x97, 0x82, 0xd7, 0x7e, 0x9a, 0x87, 0x58, 0x8f, 0x1a,
	0x54, 0xdc, 0x25, 0xd9, 0x07, 0x03, 0xb3, 0x1e, 0x63, 0x9a, 0xd5, 0x4d, 0x1e, 0x15, 0x44, 0x0e,
	0xa0, 0x8e, 0xa7, 0xeb, 0x04, 0xb9, 0xd8, 0x8a, 0xca, 0x30, 0xb2, 0x07, 0x3a, 0xba, 0x65, 0xaa,
	0x3b, 0xc3, 0xb6, 0x8a, 0x71, 0x97, 0xee, 0x2a, 0xe4, 0x54, 0x42, 0x48, 0xcb, 0x69, 0xe0, 0x0b,
	0x64, 0x55, 0x97, 0xca, 0xcf, 0x4c, 0xab, 0x07, 0xc6, 0x55, 0x30, 0xe5, 0xe3, 0x33, 0x64, 0x38,
	0x3d, 0xa9, 0x51, 0x51, 0x96, 0xf5, 0x56, 0x03, 0xe3, 0xa3, 0xec, 0x1e, 0x41, 0xc3, 0x8e, 0x82,
	0x69, 0x32, 0xe1, 0x91, 0xe2, 0x36, 0xcb, 0x23, 0xbd, 0x85, 0xae, 0x61, 0xf2, 0x33, 0xe8, 0x34,
	0x08, 0x84, 0x62, 0xb6, 0xd4, 0x02, 0x09, 0x20, 0x37, 0x36, 0x8b, 0xb2, 0x5c, 0x37, 0x42, 0x14,
	0x84, 0xed, 0xbc, 0xe4, 0xd1, 0xfd, 0x9c, 0xcb, 0xbb, 0x6a, 0x4f, 0xb4, 0x33, 0x87, 0x71, 0xae,
	0x5d, 0x6f, 0xc1, 0x63, 0xc1, 0x16, 0x61, 0xd7, 0xe8, 0x69, 0xfd, 0x2a, 0xcd, 0x1d, 0xd6, 0x6b,
	0x0d, 0x5a, 0xb2, 0x9b, 0x53, 0x55, 0xe6, 0x41, 0x56, 0xb0, 0x12, 0x6c, 0x56, 0x8a, 0x12, 0x6b,
	0xc6, 0xc6, 0x11, 0x34, 0xd3, 0x61, 0x7c, 0x46, 0x50, 0x39, 0xfa, 0x09, 0xf5, 0xf4, 0x37, 0xd4,
	0x1c, 0xc1, 0x04, 0x47, 0x5a, 0x31, 0x2f, 0x95, 0x72, 0x99, 0x56, 0x04, 0xb0, 0xbf, 0x6a, 0x93,
	0x54, 0x64, 0xef, 0x94, 0x65, 0xb9, 0xa0, 0x8f, 0xec, 0x54, 0x92, 0x6a, 0x38, 0x9e, 0xb8, 0x42,
	0x41, 0xe4, 0x10, 0x0c, 0xb1, 0x9c, 0x32, 0xc1, 0x54, 0xc1, 0x5f, 0xab, 0x20, 0xbc, 0xe1, 0x8c,
	0x09, 0x46, 0x15, 0x6c, 0xbd, 0xd7, 0xa0, 0x91, 0x39, 0x89, 0x05, 0xad, 0x93, 0xc9, 0x04, 0xe7,
	0xf5, 0x2a, 0xf0, 0x27, 0xe9, 0x46, 0xd4, 0x69, 0xc9, 0x47, 0x8e, 0xa1, 0x49, 0xf9, 0xc4, 0x0b,
	0xbd, 0x6c, 0xf4, 0xb7, 0x24, 0x94, 0xe3, 0x58, 0xcb, 0x49, 0xbe, 0x0d, 0x5a, 0x54, 0x59, 0x28,
	0x50, 0x9b, 0xad, 0xe6, 0x01, 0x9b, 0x4a, 0xee, 0x5a, 0x34, 0x33, 0x89, 0x05, 0x55, 0xc7, 0x9b,
	0x29, 0xa1, 0x6c, 0x33, 0x8a, 0x20, 0x72, 0x1f, 0xae, 0xd7, 0xbf, 0xf1, 0x4c, 0xa3, 0xf2, 0x10,
	0xeb, 0x9d, 0xac, 0xf1, 0x8b, 0x9e, 0x8d, 0x37, 0x1a, 0xec, 0x8c, 0xec, 0xcf, 0x76, 0x3a, 0x6e,
	0xa0, 0x3e, 0xb2, 0xe5, 0x43, 0xe4, 0xb7, 0x8d, 0xb4, 0xbf, 0x5b, 0x2b, 0xb7, 0x58, 0xdd, 0x3a,
	0xfd, 0x1f, 0x70, 0x4d, 0x6f, 0xac, 0x5d, 0xd4, 0x33, 0xad, 0xb8, 0xcb, 0x5f, 0x7b, 0x38, 0x2a,
	0xb8, 0x5e, 0x49, 0x13, 0x6a, 0xff, 0xdf, 0x9c, 0xd3, 0x97, 0x9d, 0xaf, 0x88, 0x09, 0x75, 0xc7,
	0xbd, 0xa6, 0x27, 0xff, 0x9e, 0x77, 0xb4, 0x5b, 0x43, 0xfe, 0x1b, 0xf9, 0xe3, 0x43, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x97, 0xed, 0x86, 0xe5, 0x9c, 0x08, 0x00, 0x00,
}
