// Code generated by protoc-gen-go.
// source: trillian.proto
// DO NOT EDIT!

package trillian

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import sigpb "github.com/google/trillian/crypto/sigpb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Defines the way empty / node / leaf hashes are constructed incorporating
// preimage protection, which can be application specific.
type HashStrategy int32

const (
	// Hash strategy cannot be determined. Included to enable detection of
	// mismatched proto versions being used. Represents an invalid value.
	HashStrategy_UNKNOWN_HASH_STRATEGY HashStrategy = 0
	// Certificate Transparency strategy: leaf hash prefix = 0x00, node prefix =
	// 0x01, empty hash is digest([]byte{}), as defined in the specification.
	HashStrategy_RFC_6962 HashStrategy = 1
)

var HashStrategy_name = map[int32]string{
	0: "UNKNOWN_HASH_STRATEGY",
	1: "RFC_6962",
}
var HashStrategy_value = map[string]int32{
	"UNKNOWN_HASH_STRATEGY": 0,
	"RFC_6962":              1,
}

func (x HashStrategy) String() string {
	return proto.EnumName(HashStrategy_name, int32(x))
}
func (HashStrategy) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// State of the tree.
type TreeState int32

const (
	// Tree state cannot be determined. Included to enable detection of
	// mismatched proto versions being used. Represents an invalid value.
	TreeState_UNKNOWN_TREE_STATE TreeState = 0
	// Active trees are able to respond to both read and write requests.
	TreeState_ACTIVE TreeState = 1
	// Frozen trees are only able to respond to read requests, writing to a frozen
	// tree is forbidden.
	TreeState_FROZEN TreeState = 2
	// Tree was been deleted, therefore is invisible and acts similarly to a
	// non-existing tree for all requests.
	// A soft deleted tree may be undeleted while the soft-deletion period has not
	// passed.
	TreeState_SOFT_DELETED TreeState = 3
	// A hard deleted tree was been definitely deleted and cannot be recovered.
	// Acts an a non-existing tree for all read and write requests, but blocks the
	// tree ID from ever being reused.
	TreeState_HARD_DELETED TreeState = 4
)

var TreeState_name = map[int32]string{
	0: "UNKNOWN_TREE_STATE",
	1: "ACTIVE",
	2: "FROZEN",
	3: "SOFT_DELETED",
	4: "HARD_DELETED",
}
var TreeState_value = map[string]int32{
	"UNKNOWN_TREE_STATE": 0,
	"ACTIVE":             1,
	"FROZEN":             2,
	"SOFT_DELETED":       3,
	"HARD_DELETED":       4,
}

func (x TreeState) String() string {
	return proto.EnumName(TreeState_name, int32(x))
}
func (TreeState) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// Type of the tree.
type TreeType int32

const (
	// Tree type cannot be determined. Included to enable detection of
	// mismatched proto versions being used. Represents an invalid value.
	TreeType_UNKNOWN_TREE_TYPE TreeType = 0
	// Tree represents a verifiable log.
	TreeType_LOG TreeType = 1
	// Tree represents a verifiable map.
	TreeType_MAP TreeType = 2
)

var TreeType_name = map[int32]string{
	0: "UNKNOWN_TREE_TYPE",
	1: "LOG",
	2: "MAP",
}
var TreeType_value = map[string]int32{
	"UNKNOWN_TREE_TYPE": 0,
	"LOG":               1,
	"MAP":               2,
}

func (x TreeType) String() string {
	return proto.EnumName(TreeType_name, int32(x))
}
func (TreeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

// Duplicate policy of a tree.
type DuplicatePolicy int32

const (
	// Policy cannot be determined. Included to enable detection of
	// mismatched proto versions being used. Represents an invalid value.
	DuplicatePolicy_UNKNOWN_DUPLICATE_POLICY DuplicatePolicy = 0
	// Duplicates are not allowed in the tree and will cause errors on insertion.
	DuplicatePolicy_DUPLICATES_NOT_ALLOWED DuplicatePolicy = 1
	// Duplicates are allowed in the tree.
	DuplicatePolicy_DUPLICATES_ALLOWED DuplicatePolicy = 2
)

var DuplicatePolicy_name = map[int32]string{
	0: "UNKNOWN_DUPLICATE_POLICY",
	1: "DUPLICATES_NOT_ALLOWED",
	2: "DUPLICATES_ALLOWED",
}
var DuplicatePolicy_value = map[string]int32{
	"UNKNOWN_DUPLICATE_POLICY": 0,
	"DUPLICATES_NOT_ALLOWED":   1,
	"DUPLICATES_ALLOWED":       2,
}

func (x DuplicatePolicy) String() string {
	return proto.EnumName(DuplicatePolicy_name, int32(x))
}
func (DuplicatePolicy) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

// Represents a tree, which may be either a verifiable log or map.
// Readonly attributes are assigned at tree creation, after which they may not
// be modified.
type Tree struct {
	// ID of the tree.
	// Readonly.
	TreeId int64 `protobuf:"varint,1,opt,name=tree_id,json=treeId" json:"tree_id,omitempty"`
	// State of the tree.
	// Trees are active after creation. At any point the tree may transition
	// between ACTIVE and FROZEN.
	// Deleted trees are set as SOFT_DELETED for a certain time period, after
	// which they'll automatically transition to HARD_DELETED.
	TreeState TreeState `protobuf:"varint,2,opt,name=tree_state,json=treeState,enum=trillian.TreeState" json:"tree_state,omitempty"`
	// Type of the tree.
	// Readonly.
	TreeType TreeType `protobuf:"varint,3,opt,name=tree_type,json=treeType,enum=trillian.TreeType" json:"tree_type,omitempty"`
	// Hash strategy to be used by the tree.
	// Readonly.
	HashStrategy HashStrategy `protobuf:"varint,4,opt,name=hash_strategy,json=hashStrategy,enum=trillian.HashStrategy" json:"hash_strategy,omitempty"`
	// Hash algorithm to be used by the tree.
	// Readonly.
	HashAlgorithm sigpb.DigitallySigned_HashAlgorithm `protobuf:"varint,5,opt,name=hash_algorithm,json=hashAlgorithm,enum=sigpb.DigitallySigned_HashAlgorithm" json:"hash_algorithm,omitempty"`
	// Signature algorithm to be used by the tree.
	// Readonly.
	SignatureAlgorithm sigpb.DigitallySigned_SignatureAlgorithm `protobuf:"varint,6,opt,name=signature_algorithm,json=signatureAlgorithm,enum=sigpb.DigitallySigned_SignatureAlgorithm" json:"signature_algorithm,omitempty"`
	// Duplicate policy to be used by the tree.
	// Readonly.
	DuplicatePolicy DuplicatePolicy `protobuf:"varint,7,opt,name=duplicate_policy,json=duplicatePolicy,enum=trillian.DuplicatePolicy" json:"duplicate_policy,omitempty"`
	// Display name of the tree.
	// Optional.
	DisplayName string `protobuf:"bytes,8,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	// Description of the tree,
	// Optional.
	Description string `protobuf:"bytes,9,opt,name=description" json:"description,omitempty"`
	// Timestamp of tree creation.
	// Readonly.
	CreateTimeMillisSinceEpoch int64 `protobuf:"varint,10,opt,name=create_time_millis_since_epoch,json=createTimeMillisSinceEpoch" json:"create_time_millis_since_epoch,omitempty"`
	// Timestamp of last tree update.
	// Readonly (automatically assigned on updates).
	UpdateTimeMillisSinceEpoch int64 `protobuf:"varint,11,opt,name=update_time_millis_since_epoch,json=updateTimeMillisSinceEpoch" json:"update_time_millis_since_epoch,omitempty"`
}

func (m *Tree) Reset()                    { *m = Tree{} }
func (m *Tree) String() string            { return proto.CompactTextString(m) }
func (*Tree) ProtoMessage()               {}
func (*Tree) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Tree) GetTreeId() int64 {
	if m != nil {
		return m.TreeId
	}
	return 0
}

func (m *Tree) GetTreeState() TreeState {
	if m != nil {
		return m.TreeState
	}
	return TreeState_UNKNOWN_TREE_STATE
}

func (m *Tree) GetTreeType() TreeType {
	if m != nil {
		return m.TreeType
	}
	return TreeType_UNKNOWN_TREE_TYPE
}

func (m *Tree) GetHashStrategy() HashStrategy {
	if m != nil {
		return m.HashStrategy
	}
	return HashStrategy_UNKNOWN_HASH_STRATEGY
}

func (m *Tree) GetHashAlgorithm() sigpb.DigitallySigned_HashAlgorithm {
	if m != nil {
		return m.HashAlgorithm
	}
	return sigpb.DigitallySigned_NONE
}

func (m *Tree) GetSignatureAlgorithm() sigpb.DigitallySigned_SignatureAlgorithm {
	if m != nil {
		return m.SignatureAlgorithm
	}
	return sigpb.DigitallySigned_ANONYMOUS
}

func (m *Tree) GetDuplicatePolicy() DuplicatePolicy {
	if m != nil {
		return m.DuplicatePolicy
	}
	return DuplicatePolicy_UNKNOWN_DUPLICATE_POLICY
}

func (m *Tree) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Tree) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Tree) GetCreateTimeMillisSinceEpoch() int64 {
	if m != nil {
		return m.CreateTimeMillisSinceEpoch
	}
	return 0
}

func (m *Tree) GetUpdateTimeMillisSinceEpoch() int64 {
	if m != nil {
		return m.UpdateTimeMillisSinceEpoch
	}
	return 0
}

type SignedEntryTimestamp struct {
	TimestampNanos int64                  `protobuf:"varint,1,opt,name=timestamp_nanos,json=timestampNanos" json:"timestamp_nanos,omitempty"`
	LogId          int64                  `protobuf:"varint,2,opt,name=log_id,json=logId" json:"log_id,omitempty"`
	Signature      *sigpb.DigitallySigned `protobuf:"bytes,3,opt,name=signature" json:"signature,omitempty"`
}

func (m *SignedEntryTimestamp) Reset()                    { *m = SignedEntryTimestamp{} }
func (m *SignedEntryTimestamp) String() string            { return proto.CompactTextString(m) }
func (*SignedEntryTimestamp) ProtoMessage()               {}
func (*SignedEntryTimestamp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *SignedEntryTimestamp) GetTimestampNanos() int64 {
	if m != nil {
		return m.TimestampNanos
	}
	return 0
}

func (m *SignedEntryTimestamp) GetLogId() int64 {
	if m != nil {
		return m.LogId
	}
	return 0
}

func (m *SignedEntryTimestamp) GetSignature() *sigpb.DigitallySigned {
	if m != nil {
		return m.Signature
	}
	return nil
}

// SignedLogRoot represents a commitment by a Log to a particular tree.
type SignedLogRoot struct {
	// epoch nanoseconds, good until 2500ish
	TimestampNanos int64  `protobuf:"varint,1,opt,name=timestamp_nanos,json=timestampNanos" json:"timestamp_nanos,omitempty"`
	RootHash       []byte `protobuf:"bytes,2,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`
	// TreeSize is the number of entries in the tree.
	TreeSize int64 `protobuf:"varint,3,opt,name=tree_size,json=treeSize" json:"tree_size,omitempty"`
	// TODO(al): define serialized format for the signature scheme.
	Signature    *sigpb.DigitallySigned `protobuf:"bytes,4,opt,name=signature" json:"signature,omitempty"`
	LogId        int64                  `protobuf:"varint,5,opt,name=log_id,json=logId" json:"log_id,omitempty"`
	TreeRevision int64                  `protobuf:"varint,6,opt,name=tree_revision,json=treeRevision" json:"tree_revision,omitempty"`
}

func (m *SignedLogRoot) Reset()                    { *m = SignedLogRoot{} }
func (m *SignedLogRoot) String() string            { return proto.CompactTextString(m) }
func (*SignedLogRoot) ProtoMessage()               {}
func (*SignedLogRoot) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *SignedLogRoot) GetTimestampNanos() int64 {
	if m != nil {
		return m.TimestampNanos
	}
	return 0
}

func (m *SignedLogRoot) GetRootHash() []byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

func (m *SignedLogRoot) GetTreeSize() int64 {
	if m != nil {
		return m.TreeSize
	}
	return 0
}

func (m *SignedLogRoot) GetSignature() *sigpb.DigitallySigned {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignedLogRoot) GetLogId() int64 {
	if m != nil {
		return m.LogId
	}
	return 0
}

func (m *SignedLogRoot) GetTreeRevision() int64 {
	if m != nil {
		return m.TreeRevision
	}
	return 0
}

type MapperMetadata struct {
	SourceLogId                  []byte `protobuf:"bytes,1,opt,name=source_log_id,json=sourceLogId,proto3" json:"source_log_id,omitempty"`
	HighestFullyCompletedSeq     int64  `protobuf:"varint,2,opt,name=highest_fully_completed_seq,json=highestFullyCompletedSeq" json:"highest_fully_completed_seq,omitempty"`
	HighestPartiallyCompletedSeq int64  `protobuf:"varint,3,opt,name=highest_partially_completed_seq,json=highestPartiallyCompletedSeq" json:"highest_partially_completed_seq,omitempty"`
}

func (m *MapperMetadata) Reset()                    { *m = MapperMetadata{} }
func (m *MapperMetadata) String() string            { return proto.CompactTextString(m) }
func (*MapperMetadata) ProtoMessage()               {}
func (*MapperMetadata) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *MapperMetadata) GetSourceLogId() []byte {
	if m != nil {
		return m.SourceLogId
	}
	return nil
}

func (m *MapperMetadata) GetHighestFullyCompletedSeq() int64 {
	if m != nil {
		return m.HighestFullyCompletedSeq
	}
	return 0
}

func (m *MapperMetadata) GetHighestPartiallyCompletedSeq() int64 {
	if m != nil {
		return m.HighestPartiallyCompletedSeq
	}
	return 0
}

// SignedMapRoot represents a commitment by a Map to a particular tree.
type SignedMapRoot struct {
	TimestampNanos int64           `protobuf:"varint,1,opt,name=timestamp_nanos,json=timestampNanos" json:"timestamp_nanos,omitempty"`
	RootHash       []byte          `protobuf:"bytes,2,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`
	Metadata       *MapperMetadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
	// TODO(al): define serialized format for the signature scheme.
	Signature   *sigpb.DigitallySigned `protobuf:"bytes,4,opt,name=signature" json:"signature,omitempty"`
	MapId       int64                  `protobuf:"varint,5,opt,name=map_id,json=mapId" json:"map_id,omitempty"`
	MapRevision int64                  `protobuf:"varint,6,opt,name=map_revision,json=mapRevision" json:"map_revision,omitempty"`
}

func (m *SignedMapRoot) Reset()                    { *m = SignedMapRoot{} }
func (m *SignedMapRoot) String() string            { return proto.CompactTextString(m) }
func (*SignedMapRoot) ProtoMessage()               {}
func (*SignedMapRoot) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *SignedMapRoot) GetTimestampNanos() int64 {
	if m != nil {
		return m.TimestampNanos
	}
	return 0
}

func (m *SignedMapRoot) GetRootHash() []byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

func (m *SignedMapRoot) GetMetadata() *MapperMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *SignedMapRoot) GetSignature() *sigpb.DigitallySigned {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignedMapRoot) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *SignedMapRoot) GetMapRevision() int64 {
	if m != nil {
		return m.MapRevision
	}
	return 0
}

func init() {
	proto.RegisterType((*Tree)(nil), "trillian.Tree")
	proto.RegisterType((*SignedEntryTimestamp)(nil), "trillian.SignedEntryTimestamp")
	proto.RegisterType((*SignedLogRoot)(nil), "trillian.SignedLogRoot")
	proto.RegisterType((*MapperMetadata)(nil), "trillian.MapperMetadata")
	proto.RegisterType((*SignedMapRoot)(nil), "trillian.SignedMapRoot")
	proto.RegisterEnum("trillian.HashStrategy", HashStrategy_name, HashStrategy_value)
	proto.RegisterEnum("trillian.TreeState", TreeState_name, TreeState_value)
	proto.RegisterEnum("trillian.TreeType", TreeType_name, TreeType_value)
	proto.RegisterEnum("trillian.DuplicatePolicy", DuplicatePolicy_name, DuplicatePolicy_value)
}

func init() { proto.RegisterFile("trillian.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 881 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xae, 0x6c, 0xd7, 0xb1, 0x8f, 0x1d, 0x47, 0x63, 0x97, 0x4c, 0x4d, 0x8b, 0x2d, 0xf3, 0x06,
	0x2c, 0xcb, 0x45, 0x0c, 0x24, 0x45, 0x87, 0x61, 0xd8, 0x85, 0x67, 0x2b, 0x8d, 0x51, 0xff, 0x41,
	0x52, 0x57, 0xb4, 0x37, 0x04, 0x23, 0x71, 0x32, 0x01, 0xc9, 0x64, 0x25, 0x7a, 0x80, 0xfa, 0x0c,
	0x7b, 0xa2, 0x3d, 0xcf, 0xde, 0x62, 0x18, 0x30, 0x90, 0xfa, 0xb1, 0xdd, 0x76, 0x43, 0x31, 0xec,
	0xc6, 0x20, 0xbf, 0xf3, 0x7d, 0x9f, 0xce, 0xe1, 0x39, 0xa4, 0xa1, 0x27, 0x13, 0x16, 0x45, 0x8c,
	0xac, 0x2f, 0x45, 0xc2, 0x25, 0x47, 0xad, 0x72, 0x7f, 0x7a, 0x1d, 0x32, 0xb9, 0xda, 0xdc, 0x5d,
	0xfa, 0x3c, 0x1e, 0x84, 0x9c, 0x87, 0x11, 0x1d, 0x94, 0xb1, 0x81, 0x9f, 0x64, 0x42, 0xf2, 0x41,
	0xca, 0x42, 0x71, 0x97, 0xff, 0xe6, 0xf2, 0xfe, 0x5f, 0x0d, 0x68, 0x78, 0x09, 0xa5, 0xe8, 0x33,
	0x38, 0x90, 0x09, 0xa5, 0x98, 0x05, 0x96, 0x71, 0x66, 0x9c, 0xd7, 0x9d, 0xa6, 0xda, 0x4e, 0x02,
	0x74, 0x05, 0xa0, 0x03, 0xa9, 0x24, 0x92, 0x5a, 0xb5, 0x33, 0xe3, 0xbc, 0x77, 0xf5, 0xe0, 0xb2,
	0xca, 0x42, 0x89, 0x5d, 0x15, 0x72, 0xda, 0xb2, 0x5c, 0xa2, 0x01, 0xe8, 0x0d, 0x96, 0x99, 0xa0,
	0x56, 0x5d, 0x4b, 0xd0, 0xbe, 0xc4, 0xcb, 0x04, 0x75, 0x5a, 0xb2, 0x58, 0xa1, 0x1f, 0xe0, 0x70,
	0x45, 0xd2, 0x15, 0x4e, 0x65, 0x42, 0x24, 0x0d, 0x33, 0xab, 0xa1, 0x45, 0x27, 0x5b, 0xd1, 0x2d,
	0x49, 0x57, 0x6e, 0x11, 0x75, 0xba, 0xab, 0x9d, 0x1d, 0x7a, 0x0e, 0x3d, 0x2d, 0x26, 0x51, 0xc8,
	0x13, 0x26, 0x57, 0xb1, 0x75, 0x5f, 0xab, 0xbf, 0xbe, 0xcc, 0x2b, 0x1d, 0xb3, 0x90, 0x49, 0x12,
	0x45, 0x99, 0xcb, 0xc2, 0x35, 0x0d, 0xb4, 0xd5, 0xb0, 0xe4, 0x3a, 0xfa, 0xc3, 0xd5, 0x16, 0xbd,
	0x86, 0x07, 0x29, 0x0b, 0xd7, 0x44, 0x6e, 0x12, 0xba, 0xe3, 0xd8, 0xd4, 0x8e, 0xdf, 0xfe, 0x83,
	0xa3, 0x5b, 0x2a, 0xb6, 0xb6, 0x28, 0x7d, 0x0f, 0x43, 0x63, 0x30, 0x83, 0x8d, 0x88, 0x98, 0x4f,
	0x24, 0xc5, 0x82, 0x47, 0xcc, 0xcf, 0xac, 0x03, 0x6d, 0xfc, 0x70, 0x5b, 0xe8, 0xb8, 0x64, 0x2c,
	0x35, 0xc1, 0x39, 0x0a, 0xf6, 0x01, 0xf4, 0x25, 0x74, 0x03, 0x96, 0x8a, 0x88, 0x64, 0x78, 0x4d,
	0x62, 0x6a, 0xb5, 0xce, 0x8c, 0xf3, 0xb6, 0xd3, 0x29, 0xb0, 0x39, 0x89, 0x29, 0x3a, 0x83, 0x4e,
	0x40, 0x53, 0x3f, 0x61, 0x42, 0x32, 0xbe, 0xb6, 0xda, 0x05, 0x63, 0x0b, 0xa1, 0x9f, 0xe0, 0x73,
	0x3f, 0xa1, 0x2a, 0x0f, 0xc9, 0x62, 0x8a, 0x63, 0xf5, 0xf1, 0x14, 0xa7, 0x6c, 0xed, 0x53, 0x4c,
	0x05, 0xf7, 0x57, 0x16, 0xe8, 0x29, 0x38, 0xcd, 0x59, 0x1e, 0x8b, 0xe9, 0x4c, 0x73, 0x5c, 0x45,
	0xb1, 0x15, 0x43, 0x79, 0x6c, 0x44, 0xf0, 0x6f, 0x1e, 0x9d, 0xdc, 0x23, 0x67, 0x7d, 0xc8, 0xa3,
	0xff, 0x9b, 0x01, 0x9f, 0xe6, 0x87, 0x68, 0xaf, 0x65, 0x92, 0x29, 0x4e, 0x2a, 0x49, 0x2c, 0xd0,
	0x37, 0x70, 0x24, 0xcb, 0x0d, 0x5e, 0x93, 0x35, 0x4f, 0x8b, 0xb9, 0xec, 0x55, 0xf0, 0x5c, 0xa1,
	0xe8, 0x18, 0x9a, 0x11, 0x0f, 0xd5, 0xdc, 0xd6, 0x74, 0xfc, 0x7e, 0xc4, 0xc3, 0x49, 0x80, 0x9e,
	0x40, 0xbb, 0xea, 0x80, 0x1e, 0xc1, 0xce, 0xd5, 0xc9, 0x87, 0xbb, 0xe7, 0x6c, 0x89, 0xfd, 0x3f,
	0x0c, 0x38, 0xcc, 0xd1, 0x29, 0x0f, 0x1d, 0xce, 0xe5, 0xc7, 0xe7, 0xf1, 0x08, 0xda, 0x09, 0xe7,
	0x12, 0xab, 0x71, 0xd2, 0xa9, 0x74, 0x9d, 0x96, 0x02, 0xd4, 0xb4, 0xa9, 0x60, 0x7e, 0x89, 0xd8,
	0xdb, 0x3c, 0x9b, 0x7a, 0x3e, 0xfc, 0x2e, 0x7b, 0x4b, 0xf7, 0x53, 0x6d, 0x7c, 0x64, 0xaa, 0x3b,
	0x75, 0xdf, 0xdf, 0xad, 0xfb, 0x2b, 0x38, 0xd4, 0x5f, 0x4a, 0xe8, 0xaf, 0x2c, 0x55, 0xcd, 0x6f,
	0xea, 0x68, 0x57, 0x81, 0x4e, 0x81, 0xf5, 0x7f, 0x37, 0xa0, 0x37, 0x23, 0x42, 0xd0, 0x64, 0x46,
	0x25, 0x09, 0x88, 0x24, 0xa8, 0x0f, 0x87, 0x29, 0xdf, 0x24, 0x3e, 0xc5, 0x85, 0xab, 0xa1, 0x4b,
	0xe8, 0xe4, 0xe0, 0x54, 0x7b, 0xff, 0x08, 0x8f, 0x56, 0x2c, 0x5c, 0xd1, 0x54, 0xe2, 0x5f, 0x36,
	0x51, 0x94, 0x61, 0x9f, 0xc7, 0x22, 0xa2, 0x92, 0x06, 0x38, 0xa5, 0x6f, 0x8a, 0xf3, 0xb7, 0x0a,
	0xca, 0x8d, 0x62, 0x8c, 0x4a, 0x82, 0x4b, 0xdf, 0x20, 0x1b, 0xbe, 0x28, 0xe5, 0x82, 0x24, 0x92,
	0x91, 0xf7, 0x2d, 0xf2, 0xa3, 0x79, 0x5c, 0xd0, 0x96, 0x25, 0x6b, 0xd7, 0xa6, 0xff, 0x67, 0xd5,
	0xa3, 0x19, 0x11, 0xff, 0x63, 0x8f, 0x9e, 0x40, 0x2b, 0x2e, 0x4e, 0xa3, 0x18, 0x18, 0x6b, 0x7b,
	0x2b, 0xf7, 0x4f, 0xcb, 0xa9, 0x98, 0xff, 0xbd, 0x79, 0x31, 0x11, 0x3b, 0xcd, 0x8b, 0x89, 0x98,
	0x04, 0xea, 0x6a, 0x2b, 0xf8, 0x9d, 0xde, 0x75, 0x62, 0x22, 0xca, 0xd6, 0x5d, 0x7c, 0x07, 0xdd,
	0xdd, 0xa7, 0x10, 0x3d, 0x84, 0xe3, 0x17, 0xf3, 0xe7, 0xf3, 0xc5, 0xcb, 0x39, 0xbe, 0x1d, 0xba,
	0xb7, 0xd8, 0xf5, 0x9c, 0xa1, 0x67, 0x3f, 0x7b, 0x65, 0xde, 0x43, 0x5d, 0x68, 0x39, 0x37, 0x23,
	0xfc, 0xf4, 0xfb, 0xa7, 0x57, 0xa6, 0x71, 0x81, 0xa1, 0x5d, 0xbd, 0xd5, 0xe8, 0x04, 0x50, 0xa9,
	0xf2, 0x1c, 0xdb, 0xc6, 0xae, 0x37, 0xf4, 0x6c, 0xf3, 0x1e, 0x02, 0x68, 0x0e, 0x47, 0xde, 0xe4,
	0x67, 0xdb, 0x34, 0xd4, 0xfa, 0xc6, 0x59, 0xbc, 0xb6, 0xe7, 0x66, 0x0d, 0x99, 0xd0, 0x75, 0x17,
	0x37, 0x1e, 0x1e, 0xdb, 0x53, 0xdb, 0xb3, 0xc7, 0x66, 0x5d, 0x21, 0xb7, 0x43, 0x67, 0x5c, 0x21,
	0x8d, 0x8b, 0x6b, 0x68, 0x95, 0x2f, 0x3b, 0x3a, 0x86, 0x4f, 0xf6, 0xfc, 0xbd, 0x57, 0x4b, 0x65,
	0x7f, 0x00, 0xf5, 0xe9, 0xe2, 0x99, 0x69, 0xa8, 0xc5, 0x6c, 0xb8, 0x34, 0x6b, 0x17, 0x3e, 0x1c,
	0xbd, 0xf3, 0xe0, 0xa1, 0xc7, 0x60, 0x95, 0xda, 0xf1, 0x8b, 0xe5, 0x74, 0x32, 0x1a, 0x7a, 0x36,
	0x5e, 0x2e, 0xa6, 0x93, 0x91, 0x2a, 0xea, 0x14, 0x4e, 0x2a, 0xd4, 0xc5, 0xf3, 0x85, 0x87, 0x87,
	0xd3, 0xe9, 0xe2, 0xa5, 0x3d, 0x36, 0x0d, 0x55, 0xd5, 0x4e, 0xac, 0xc4, 0x6b, 0x77, 0x4d, 0xfd,
	0x5f, 0x77, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x3f, 0x34, 0x38, 0x3c, 0x07, 0x00,
	0x00,
}
