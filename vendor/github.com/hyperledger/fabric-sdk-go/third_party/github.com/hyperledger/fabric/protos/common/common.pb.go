/*
Notice: This file has been modified for Hyperledger Fabric SDK Go usage.
Please review third_party pinning scripts and patches for more details.
*/
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common // import "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// These status codes are intended to resemble selected HTTP status codes
type Status int32

const (
	Status_UNKNOWN                  Status = 0
	Status_SUCCESS                  Status = 200
	Status_BAD_REQUEST              Status = 400
	Status_FORBIDDEN                Status = 403
	Status_NOT_FOUND                Status = 404
	Status_REQUEST_ENTITY_TOO_LARGE Status = 413
	Status_INTERNAL_SERVER_ERROR    Status = 500
	Status_NOT_IMPLEMENTED          Status = 501
	Status_SERVICE_UNAVAILABLE      Status = 503
)

var Status_name = map[int32]string{
	0:   "UNKNOWN",
	200: "SUCCESS",
	400: "BAD_REQUEST",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	413: "REQUEST_ENTITY_TOO_LARGE",
	500: "INTERNAL_SERVER_ERROR",
	501: "NOT_IMPLEMENTED",
	503: "SERVICE_UNAVAILABLE",
}
var Status_value = map[string]int32{
	"UNKNOWN":                  0,
	"SUCCESS":                  200,
	"BAD_REQUEST":              400,
	"FORBIDDEN":                403,
	"NOT_FOUND":                404,
	"REQUEST_ENTITY_TOO_LARGE": 413,
	"INTERNAL_SERVER_ERROR":    500,
	"NOT_IMPLEMENTED":          501,
	"SERVICE_UNAVAILABLE":      503,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_b374fafc5e1c956e, []int{0}
}

type HeaderType int32

const (
	HeaderType_MESSAGE              HeaderType = 0
	HeaderType_CONFIG               HeaderType = 1
	HeaderType_CONFIG_UPDATE        HeaderType = 2
	HeaderType_ENDORSER_TRANSACTION HeaderType = 3
	HeaderType_ORDERER_TRANSACTION  HeaderType = 4
	HeaderType_DELIVER_SEEK_INFO    HeaderType = 5
	HeaderType_CHAINCODE_PACKAGE    HeaderType = 6
	HeaderType_PEER_ADMIN_OPERATION HeaderType = 8
	HeaderType_TOKEN_TRANSACTION    HeaderType = 9
)

var HeaderType_name = map[int32]string{
	0: "MESSAGE",
	1: "CONFIG",
	2: "CONFIG_UPDATE",
	3: "ENDORSER_TRANSACTION",
	4: "ORDERER_TRANSACTION",
	5: "DELIVER_SEEK_INFO",
	6: "CHAINCODE_PACKAGE",
	8: "PEER_ADMIN_OPERATION",
	9: "TOKEN_TRANSACTION",
}
var HeaderType_value = map[string]int32{
	"MESSAGE":              0,
	"CONFIG":               1,
	"CONFIG_UPDATE":        2,
	"ENDORSER_TRANSACTION": 3,
	"ORDERER_TRANSACTION":  4,
	"DELIVER_SEEK_INFO":    5,
	"CHAINCODE_PACKAGE":    6,
	"PEER_ADMIN_OPERATION": 8,
	"TOKEN_TRANSACTION":    9,
}

func (x HeaderType) String() string {
	return proto.EnumName(HeaderType_name, int32(x))
}
func (HeaderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_b374fafc5e1c956e, []int{1}
}

// This enum enlists indexes of the block metadata array
type BlockMetadataIndex int32

const (
	BlockMetadataIndex_SIGNATURES          BlockMetadataIndex = 0
	BlockMetadataIndex_LAST_CONFIG         BlockMetadataIndex = 1
	BlockMetadataIndex_TRANSACTIONS_FILTER BlockMetadataIndex = 2
	BlockMetadataIndex_ORDERER             BlockMetadataIndex = 3
)

var BlockMetadataIndex_name = map[int32]string{
	0: "SIGNATURES",
	1: "LAST_CONFIG",
	2: "TRANSACTIONS_FILTER",
	3: "ORDERER",
}
var BlockMetadataIndex_value = map[string]int32{
	"SIGNATURES":          0,
	"LAST_CONFIG":         1,
	"TRANSACTIONS_FILTER": 2,
	"ORDERER":             3,
}

func (x BlockMetadataIndex) String() string {
	return proto.EnumName(BlockMetadataIndex_name, int32(x))
}
func (BlockMetadataIndex) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_b374fafc5e1c956e, []int{2}
}

// LastConfig is the encoded value for the Metadata message which is encoded in the LAST_CONFIGURATION block metadata index
type LastConfig struct {
	Index                uint64   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LastConfig) Reset()         { *m = LastConfig{} }
func (m *LastConfig) String() string { return proto.CompactTextString(m) }
func (*LastConfig) ProtoMessage()    {}
func (*LastConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b374fafc5e1c956e, []int{0}
}
func (m *LastConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LastConfig.Unmarshal(m, b)
}
func (m *LastConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LastConfig.Marshal(b, m, deterministic)
}
func (dst *LastConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastConfig.Merge(dst, src)
}
func (m *LastConfig) XXX_Size() int {
	return xxx_messageInfo_LastConfig.Size(m)
}
func (m *LastConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LastConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LastConfig proto.InternalMessageInfo

func (m *LastConfig) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

// Metadata is a common structure to be used to encode block metadata
type Metadata struct {
	Value                []byte               `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Signatures           []*MetadataSignature `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b374fafc5e1c956e, []int{1}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (dst *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(dst, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Metadata) GetSignatures() []*MetadataSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type MetadataSignature struct {
	SignatureHeader      []byte   `protobuf:"bytes,1,opt,name=signature_header,json=signatureHeader,proto3" json:"signature_header,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetadataSignature) Reset()         { *m = MetadataSignature{} }
func (m *MetadataSignature) String() string { return proto.CompactTextString(m) }
func (*MetadataSignature) ProtoMessage()    {}
func (*MetadataSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b374fafc5e1c956e, []int{2}
}
func (m *MetadataSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataSignature.Unmarshal(m, b)
}
func (m *MetadataSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataSignature.Marshal(b, m, deterministic)
}
func (dst *MetadataSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataSignature.Merge(dst, src)
}
func (m *MetadataSignature) XXX_Size() int {
	return xxx_messageInfo_MetadataSignature.Size(m)
}
func (m *MetadataSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataSignature.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataSignature proto.InternalMessageInfo

func (m *MetadataSignature) GetSignatureHeader() []byte {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

func (m *MetadataSignature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Header struct {
	ChannelHeader        []byte   `protobuf:"bytes,1,opt,name=channel_header,json=channelHeader,proto3" json:"channel_header,omitempty"`
	SignatureHeader      []byte   `protobuf:"bytes,2,opt,name=signature_header,json=signatureHeader,proto3" json:"signature_header,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b374fafc5e1c956e, []int{3}
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

func (m *Header) GetChannelHeader() []byte {
	if m != nil {
		return m.ChannelHeader
	}
	return nil
}

func (m *Header) GetSignatureHeader() []byte {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

// Header is a generic replay prevention and identity message to include in a signed payload
type ChannelHeader struct {
	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	// Timestamp is the local time when the message was created
	// by the sender
	Timestamp *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Identifier of the channel this message is bound for
	ChannelId string `protobuf:"bytes,4,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// An unique identifier that is used end-to-end.
	//  -  set by higher layers such as end user or SDK
	//  -  passed to the endorser (which will check for uniqueness)
	//  -  as the header is passed along unchanged, it will be
	//     be retrieved by the committer (uniqueness check here as well)
	//  -  to be stored in the ledger
	TxId string `protobuf:"bytes,5,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	// The epoch in which this header was generated, where epoch is defined based on block height
	// Epoch in which the response has been generated. This field identifies a
	// logical window of time. A proposal response is accepted by a peer only if
	// two conditions hold:
	// 1. the epoch specified in the message is the current epoch
	// 2. this message has been only seen once during this epoch (i.e. it hasn't
	//    been replayed)
	Epoch uint64 `protobuf:"varint,6,opt,name=epoch,proto3" json:"epoch,omitempty"`
	// Extension that may be attached based on the header type
	Extension []byte `protobuf:"bytes,7,opt,name=extension,proto3" json:"extension,omitempty"`
	// If mutual TLS is employed, this represents
	// the hash of the client's TLS certificate
	TlsCertHash          []byte   `protobuf:"bytes,8,opt,name=tls_cert_hash,json=tlsCertHash,proto3" json:"tls_cert_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelHeader) Reset()         { *m = ChannelHeader{} }
func (m *ChannelHeader) String() string { return proto.CompactTextString(m) }
func (*ChannelHeader) ProtoMessage()    {}
func (*ChannelHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b374fafc5e1c956e, []int{4}
}
func (m *ChannelHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelHeader.Unmarshal(m, b)
}
func (m *ChannelHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelHeader.Marshal(b, m, deterministic)
}
func (dst *ChannelHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelHeader.Merge(dst, src)
}
func (m *ChannelHeader) XXX_Size() int {
	return xxx_messageInfo_ChannelHeader.Size(m)
}
func (m *ChannelHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelHeader proto.InternalMessageInfo

func (m *ChannelHeader) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ChannelHeader) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ChannelHeader) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ChannelHeader) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *ChannelHeader) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *ChannelHeader) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *ChannelHeader) GetExtension() []byte {
	if m != nil {
		return m.Extension
	}
	return nil
}

func (m *ChannelHeader) GetTlsCertHash() []byte {
	if m != nil {
		return m.TlsCertHash
	}
	return nil
}

type SignatureHeader struct {
	// Creator of the message, a marshaled msp.SerializedIdentity
	Creator []byte `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// Arbitrary number that may only be used once. Can be used to detect replay attacks.
	Nonce                []byte   `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignatureHeader) Reset()         { *m = SignatureHeader{} }
func (m *SignatureHeader) String() string { return proto.CompactTextString(m) }
func (*SignatureHeader) ProtoMessage()    {}
func (*SignatureHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b374fafc5e1c956e, []int{5}
}
func (m *SignatureHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureHeader.Unmarshal(m, b)
}
func (m *SignatureHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureHeader.Marshal(b, m, deterministic)
}
func (dst *SignatureHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureHeader.Merge(dst, src)
}
func (m *SignatureHeader) XXX_Size() int {
	return xxx_messageInfo_SignatureHeader.Size(m)
}
func (m *SignatureHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureHeader.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureHeader proto.InternalMessageInfo

func (m *SignatureHeader) GetCreator() []byte {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *SignatureHeader) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

// Payload is the message contents (and header to allow for signing)
type Payload struct {
	// Header is included to provide identity and prevent replay
	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Data, the encoding of which is defined by the type in the header
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b374fafc5e1c956e, []int{6}
}
func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (dst *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(dst, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Payload) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Envelope wraps a Payload with a signature so that the message may be authenticated
type Envelope struct {
	// A marshaled Payload
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// A signature by the creator specified in the Payload header
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Envelope) Reset()         { *m = Envelope{} }
func (m *Envelope) String() string { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()    {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b374fafc5e1c956e, []int{7}
}
func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Envelope.Unmarshal(m, b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
}
func (dst *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(dst, src)
}
func (m *Envelope) XXX_Size() int {
	return xxx_messageInfo_Envelope.Size(m)
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func (m *Envelope) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Envelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// This is finalized block structure to be shared among the orderer and peer
// Note that the BlockHeader chains to the previous BlockHeader, and the BlockData hash is embedded
// in the BlockHeader.  This makes it natural and obvious that the Data is included in the hash, but
// the Metadata is not.
type Block struct {
	Header               *BlockHeader   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data                 *BlockData     `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Metadata             *BlockMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b374fafc5e1c956e, []int{8}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetData() *BlockData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Block) GetMetadata() *BlockMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// BlockHeader is the element of the block which forms the block chain
// The block header is hashed using the configured chain hashing algorithm
// over the ASN.1 encoding of the BlockHeader
type BlockHeader struct {
	Number               uint64   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	PreviousHash         []byte   `protobuf:"bytes,2,opt,name=previous_hash,json=previousHash,proto3" json:"previous_hash,omitempty"`
	DataHash             []byte   `protobuf:"bytes,3,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b374fafc5e1c956e, []int{9}
}
func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (dst *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(dst, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *BlockHeader) GetPreviousHash() []byte {
	if m != nil {
		return m.PreviousHash
	}
	return nil
}

func (m *BlockHeader) GetDataHash() []byte {
	if m != nil {
		return m.DataHash
	}
	return nil
}

type BlockData struct {
	Data                 [][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockData) Reset()         { *m = BlockData{} }
func (m *BlockData) String() string { return proto.CompactTextString(m) }
func (*BlockData) ProtoMessage()    {}
func (*BlockData) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b374fafc5e1c956e, []int{10}
}
func (m *BlockData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockData.Unmarshal(m, b)
}
func (m *BlockData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockData.Marshal(b, m, deterministic)
}
func (dst *BlockData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockData.Merge(dst, src)
}
func (m *BlockData) XXX_Size() int {
	return xxx_messageInfo_BlockData.Size(m)
}
func (m *BlockData) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockData.DiscardUnknown(m)
}

var xxx_messageInfo_BlockData proto.InternalMessageInfo

func (m *BlockData) GetData() [][]byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type BlockMetadata struct {
	Metadata             [][]byte `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockMetadata) Reset()         { *m = BlockMetadata{} }
func (m *BlockMetadata) String() string { return proto.CompactTextString(m) }
func (*BlockMetadata) ProtoMessage()    {}
func (*BlockMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b374fafc5e1c956e, []int{11}
}
func (m *BlockMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockMetadata.Unmarshal(m, b)
}
func (m *BlockMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockMetadata.Marshal(b, m, deterministic)
}
func (dst *BlockMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockMetadata.Merge(dst, src)
}
func (m *BlockMetadata) XXX_Size() int {
	return xxx_messageInfo_BlockMetadata.Size(m)
}
func (m *BlockMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_BlockMetadata proto.InternalMessageInfo

func (m *BlockMetadata) GetMetadata() [][]byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*LastConfig)(nil), "sdk.common.LastConfig")
	proto.RegisterType((*Metadata)(nil), "sdk.common.Metadata")
	proto.RegisterType((*MetadataSignature)(nil), "sdk.common.MetadataSignature")
	proto.RegisterType((*Header)(nil), "sdk.common.Header")
	proto.RegisterType((*ChannelHeader)(nil), "sdk.common.ChannelHeader")
	proto.RegisterType((*SignatureHeader)(nil), "sdk.common.SignatureHeader")
	proto.RegisterType((*Payload)(nil), "sdk.common.Payload")
	proto.RegisterType((*Envelope)(nil), "sdk.common.Envelope")
	proto.RegisterType((*Block)(nil), "sdk.common.Block")
	proto.RegisterType((*BlockHeader)(nil), "sdk.common.BlockHeader")
	proto.RegisterType((*BlockData)(nil), "sdk.common.BlockData")
	proto.RegisterType((*BlockMetadata)(nil), "sdk.common.BlockMetadata")
	proto.RegisterEnum("sdk.common.Status", Status_name, Status_value)
	proto.RegisterEnum("sdk.common.HeaderType", HeaderType_name, HeaderType_value)
	proto.RegisterEnum("sdk.common.BlockMetadataIndex", BlockMetadataIndex_name, BlockMetadataIndex_value)
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_common_b374fafc5e1c956e) }

var fileDescriptor_common_b374fafc5e1c956e = []byte{
	// 959 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xcf, 0x6f, 0xe3, 0x44,
	0x18, 0xdd, 0xc4, 0xf9, 0xf9, 0xa5, 0x69, 0xdd, 0x49, 0xcb, 0x9a, 0xc2, 0x6a, 0x2b, 0xc3, 0xa2,
	0xd2, 0x4a, 0xa9, 0x28, 0x17, 0x38, 0x3a, 0xf6, 0xb4, 0xb5, 0x9a, 0x8e, 0xc3, 0xd8, 0x59, 0xc4,
	0x2e, 0x92, 0xe5, 0x24, 0xd3, 0x24, 0x22, 0xb1, 0x23, 0x7b, 0x52, 0xb5, 0x67, 0xee, 0x08, 0x09,
	0xae, 0xfc, 0x2f, 0x1c, 0x11, 0x7f, 0x0f, 0x88, 0x2b, 0x1a, 0x8f, 0xed, 0x4d, 0xca, 0x4a, 0x9c,
	0xe2, 0xf7, 0xe6, 0xcd, 0xf7, 0xbd, 0xf9, 0xde, 0xc4, 0x86, 0xce, 0x38, 0x5a, 0x2e, 0xa3, 0xf0,
	0x5c, 0xfe, 0x74, 0x57, 0x71, 0xc4, 0x23, 0x54, 0x93, 0xe8, 0xe8, 0xe5, 0x34, 0x8a, 0xa6, 0x0b,
	0x76, 0x9e, 0xb2, 0xa3, 0xf5, 0xdd, 0x39, 0x9f, 0x2f, 0x59, 0xc2, 0x83, 0xe5, 0x4a, 0x0a, 0x75,
	0x1d, 0xa0, 0x1f, 0x24, 0xdc, 0x8c, 0xc2, 0xbb, 0xf9, 0x14, 0x1d, 0x40, 0x75, 0x1e, 0x4e, 0xd8,
	0x83, 0x56, 0x3a, 0x2e, 0x9d, 0x54, 0xa8, 0x04, 0xfa, 0x5b, 0x68, 0xdc, 0x32, 0x1e, 0x4c, 0x02,
	0x1e, 0x08, 0xc5, 0x7d, 0xb0, 0x58, 0xb3, 0x54, 0xb1, 0x43, 0x25, 0x40, 0x5f, 0x03, 0x24, 0xf3,
	0x69, 0x18, 0xf0, 0x75, 0xcc, 0x12, 0xad, 0x7c, 0xac, 0x9c, 0xb4, 0x2e, 0x3e, 0xec, 0x66, 0x8e,
	0xf2, 0xbd, 0x6e, 0xae, 0xa0, 0x1b, 0x62, 0xfd, 0x7b, 0xd8, 0xff, 0x8f, 0x00, 0x7d, 0x0e, 0x6a,
	0x21, 0xf1, 0x67, 0x2c, 0x98, 0xb0, 0x38, 0x6b, 0xb8, 0x57, 0xf0, 0xd7, 0x29, 0x8d, 0x3e, 0x86,
	0x66, 0x41, 0x69, 0xe5, 0x54, 0xf3, 0x8e, 0xd0, 0xdf, 0x40, 0x2d, 0xd3, 0xbd, 0x82, 0xdd, 0xf1,
	0x2c, 0x08, 0x43, 0xb6, 0xd8, 0x2e, 0xd8, 0xce, 0xd8, 0x4c, 0xf6, 0xbe, 0xce, 0xe5, 0xf7, 0x76,
	0xd6, 0x7f, 0x2c, 0x43, 0xdb, 0xdc, 0xda, 0x8c, 0xa0, 0xc2, 0x1f, 0x57, 0x72, 0x36, 0x55, 0x9a,
	0x3e, 0x23, 0x0d, 0xea, 0xf7, 0x2c, 0x4e, 0xe6, 0x51, 0x98, 0xd6, 0xa9, 0xd2, 0x1c, 0xa2, 0xaf,
	0xa0, 0x59, 0xa4, 0xa1, 0x29, 0xc7, 0xa5, 0x93, 0xd6, 0xc5, 0x51, 0x57, 0xe6, 0xd5, 0xcd, 0xf3,
	0xea, 0x7a, 0xb9, 0x82, 0xbe, 0x13, 0xa3, 0x17, 0x00, 0xf9, 0x59, 0xe6, 0x13, 0xad, 0x72, 0x5c,
	0x3a, 0x69, 0xd2, 0x66, 0xc6, 0xd8, 0x13, 0xd4, 0x81, 0x2a, 0x7f, 0x10, 0x2b, 0xd5, 0x74, 0xa5,
	0xc2, 0x1f, 0xec, 0x89, 0x08, 0x8e, 0xad, 0xa2, 0xf1, 0x4c, 0xab, 0xc9, 0x68, 0x53, 0x20, 0xa6,
	0xc7, 0x1e, 0x38, 0x0b, 0x53, 0x7f, 0x75, 0x39, 0xbd, 0x82, 0x40, 0x3a, 0xb4, 0xf9, 0x22, 0xf1,
	0xc7, 0x2c, 0xe6, 0xfe, 0x2c, 0x48, 0x66, 0x5a, 0x23, 0x55, 0xb4, 0xf8, 0x22, 0x31, 0x59, 0xcc,
	0xaf, 0x83, 0x64, 0xa6, 0x1b, 0xb0, 0xe7, 0x3e, 0x89, 0x44, 0x83, 0xfa, 0x38, 0x66, 0x01, 0x8f,
	0xf2, 0x19, 0xe7, 0x50, 0x98, 0x08, 0xa3, 0x70, 0x9c, 0x07, 0x25, 0x81, 0x8e, 0xa1, 0x3e, 0x08,
	0x1e, 0x17, 0x51, 0x30, 0x41, 0x9f, 0x41, 0x6d, 0x23, 0x9d, 0xd6, 0xc5, 0x6e, 0x7e, 0x89, 0x64,
	0x69, 0x9a, 0xad, 0x8a, 0x49, 0x8b, 0x1b, 0x93, 0xd5, 0x49, 0x9f, 0xf5, 0x1e, 0x34, 0x70, 0x78,
	0xcf, 0x16, 0x91, 0x9c, 0xfa, 0x4a, 0x96, 0xcc, 0x2d, 0x64, 0xf0, 0x7f, 0xee, 0xcb, 0x4f, 0x25,
	0xa8, 0xf6, 0x16, 0xd1, 0xf8, 0x07, 0x74, 0xf6, 0xc4, 0x49, 0x27, 0x77, 0x92, 0x2e, 0x3f, 0xb1,
	0xf3, 0x6a, 0xc3, 0x4e, 0xeb, 0x62, 0x7f, 0x4b, 0x6a, 0x05, 0x3c, 0x90, 0x0e, 0xd1, 0x17, 0xd0,
	0x58, 0x66, 0x77, 0x3d, 0x0b, 0xfc, 0x70, 0x4b, 0x9a, 0xff, 0x11, 0x68, 0x21, 0xd3, 0xa7, 0xd0,
	0xda, 0x68, 0x88, 0x3e, 0x80, 0x5a, 0xb8, 0x5e, 0x8e, 0x32, 0x57, 0x15, 0x9a, 0x21, 0xf4, 0x09,
	0xb4, 0x57, 0x31, 0xbb, 0x9f, 0x47, 0xeb, 0x44, 0x26, 0x25, 0x4f, 0xb6, 0x93, 0x93, 0x22, 0x2a,
	0xf4, 0x11, 0x34, 0x45, 0x4d, 0x29, 0x50, 0x52, 0x41, 0x43, 0x10, 0x69, 0x8e, 0x2f, 0xa1, 0x59,
	0xd8, 0x2d, 0xc6, 0x5b, 0x3a, 0x56, 0x8a, 0xf1, 0x9e, 0x41, 0x7b, 0xcb, 0x24, 0x3a, 0xda, 0x38,
	0x8d, 0x14, 0x16, 0xf8, 0xf4, 0xf7, 0x12, 0xd4, 0x5c, 0x1e, 0xf0, 0x75, 0x82, 0x5a, 0x50, 0x1f,
	0x92, 0x1b, 0xe2, 0x7c, 0x4b, 0xd4, 0x67, 0x68, 0x07, 0xea, 0xee, 0xd0, 0x34, 0xb1, 0xeb, 0xaa,
	0x7f, 0x94, 0x90, 0x0a, 0xad, 0x9e, 0x61, 0xf9, 0x14, 0x7f, 0x33, 0xc4, 0xae, 0xa7, 0xfe, 0xac,
	0xa0, 0x5d, 0x68, 0x5e, 0x3a, 0xb4, 0x67, 0x5b, 0x16, 0x26, 0xea, 0x2f, 0x29, 0x26, 0x8e, 0xe7,
	0x5f, 0x3a, 0x43, 0x62, 0xa9, 0xbf, 0x2a, 0xe8, 0x05, 0x68, 0x99, 0xda, 0xc7, 0xc4, 0xb3, 0xbd,
	0xef, 0x7c, 0xcf, 0x71, 0xfc, 0xbe, 0x41, 0xaf, 0xb0, 0xfa, 0x9b, 0x82, 0x8e, 0xe0, 0xd0, 0x26,
	0x1e, 0xa6, 0xc4, 0xe8, 0xfb, 0x2e, 0xa6, 0xaf, 0x31, 0xf5, 0x31, 0xa5, 0x0e, 0x55, 0xff, 0x52,
	0xd0, 0x01, 0xec, 0x89, 0x52, 0xf6, 0xed, 0xa0, 0x8f, 0x6f, 0x31, 0xf1, 0xb0, 0xa5, 0xfe, 0xad,
	0x20, 0x0d, 0x3a, 0x42, 0x68, 0x9b, 0xd8, 0x1f, 0x12, 0xe3, 0xb5, 0x61, 0xf7, 0x8d, 0x5e, 0x1f,
	0xab, 0xff, 0x28, 0xa7, 0x7f, 0x96, 0x00, 0xe4, 0xd4, 0x3d, 0xf1, 0x3f, 0x6e, 0x41, 0xfd, 0x16,
	0xbb, 0xae, 0x71, 0x85, 0xd5, 0x67, 0x08, 0xa0, 0x66, 0x3a, 0xe4, 0xd2, 0xbe, 0x52, 0x4b, 0x68,
	0x1f, 0xda, 0xf2, 0xd9, 0x1f, 0x0e, 0x2c, 0xc3, 0xc3, 0x6a, 0x19, 0x69, 0x70, 0x80, 0x89, 0xe5,
	0x50, 0x17, 0x53, 0xdf, 0xa3, 0x06, 0x71, 0x0d, 0xd3, 0xb3, 0x1d, 0xa2, 0x2a, 0xe8, 0x39, 0x74,
	0x1c, 0x6a, 0x61, 0xfa, 0x64, 0xa1, 0x82, 0x0e, 0x61, 0xdf, 0xc2, 0x7d, 0x5b, 0x38, 0x76, 0x31,
	0xbe, 0xf1, 0x6d, 0x72, 0xe9, 0xa8, 0x55, 0x41, 0x9b, 0xd7, 0x86, 0x4d, 0x4c, 0xc7, 0xc2, 0xfe,
	0xc0, 0x30, 0x6f, 0x44, 0xff, 0x9a, 0x68, 0x30, 0xc0, 0x98, 0xfa, 0x86, 0x75, 0x6b, 0x13, 0xdf,
	0x19, 0x60, 0x6a, 0xa4, 0x75, 0x1a, 0x62, 0x83, 0xe7, 0xdc, 0x60, 0xb2, 0x55, 0xbe, 0x79, 0xfa,
	0x16, 0xd0, 0x56, 0x78, 0xb6, 0x78, 0xb1, 0xa3, 0x5d, 0x00, 0xd7, 0xbe, 0x22, 0x86, 0x37, 0xa4,
	0xd8, 0x55, 0x9f, 0xa1, 0x3d, 0x68, 0xf5, 0x0d, 0xd7, 0xf3, 0x8b, 0xb3, 0x3d, 0x87, 0xce, 0x46,
	0x1d, 0xd7, 0xbf, 0xb4, 0xfb, 0x1e, 0xa6, 0x6a, 0x59, 0x4c, 0x23, 0x3b, 0x87, 0xaa, 0xf4, 0x5c,
	0xf8, 0x34, 0x8a, 0xa7, 0xdd, 0xd9, 0xe3, 0x8a, 0xc5, 0x0b, 0x36, 0x99, 0xb2, 0xb8, 0x7b, 0x17,
	0x8c, 0xe2, 0xf9, 0x58, 0xbe, 0xc6, 0x92, 0xec, 0x8e, 0xbf, 0x39, 0x9b, 0xce, 0xf9, 0x6c, 0x3d,
	0x12, 0xf0, 0x7c, 0x43, 0x7c, 0x2e, 0xc5, 0xf2, 0x1b, 0x95, 0x64, 0xdf, 0xb1, 0x51, 0x2d, 0x85,
	0x5f, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x63, 0x9b, 0x9f, 0x8b, 0xdf, 0x06, 0x00, 0x00,
}
