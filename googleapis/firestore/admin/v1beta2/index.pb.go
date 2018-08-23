// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/firestore/admin/v1beta2/index.proto

package admin // import "github.com/sunnogo/go-genproto/googleapis/firestore/admin/v1beta2"

import proto "github.com/sunnogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/sunnogo/go-genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Query Scope defines the scope at which a query is run. This is specified on
// a StructuredQuery's `from` field.
type Index_QueryScope int32

const (
	// The query scope is unspecified. Not a valid option.
	Index_QUERY_SCOPE_UNSPECIFIED Index_QueryScope = 0
	// Indexes with a collection query scope specified allow queries
	// against a collection that is the child of a specific document, specified
	// at query time, and that has the collection id specified by the index.
	Index_COLLECTION Index_QueryScope = 1
)

var Index_QueryScope_name = map[int32]string{
	0: "QUERY_SCOPE_UNSPECIFIED",
	1: "COLLECTION",
}
var Index_QueryScope_value = map[string]int32{
	"QUERY_SCOPE_UNSPECIFIED": 0,
	"COLLECTION":              1,
}

func (x Index_QueryScope) String() string {
	return proto.EnumName(Index_QueryScope_name, int32(x))
}
func (Index_QueryScope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_index_47cf3a5e57bc9ffa, []int{0, 0}
}

// The state of an index. During index creation, an index will be in the
// `CREATING` state. If the index is created successfully, it will transition
// to the `READY` state. If the index creation encounters a problem, the index
// will transition to the `NEEDS_REPAIR` state.
type Index_State int32

const (
	// The state is unspecified.
	Index_STATE_UNSPECIFIED Index_State = 0
	// The index is being created.
	// There is an active long-running operation for the index.
	// The index is updated when writing a document.
	// Some index data may exist.
	Index_CREATING Index_State = 1
	// The index is ready to be used.
	// The index is updated when writing a document.
	// The index is fully populated from all stored documents it applies to.
	Index_READY Index_State = 2
	// The index was being created, but something went wrong.
	// There is no active long-running operation for the index,
	// and the most recently finished long-running operation failed.
	// The index is not updated when writing a document.
	// Some index data may exist.
	// Use the google.longrunning.Operations API to determine why the operation
	// that last attempted to create this index failed, then re-create the
	// index.
	Index_NEEDS_REPAIR Index_State = 3
)

var Index_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "CREATING",
	2: "READY",
	3: "NEEDS_REPAIR",
}
var Index_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"CREATING":          1,
	"READY":             2,
	"NEEDS_REPAIR":      3,
}

func (x Index_State) String() string {
	return proto.EnumName(Index_State_name, int32(x))
}
func (Index_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_index_47cf3a5e57bc9ffa, []int{0, 1}
}

// The supported orderings.
type Index_IndexField_Order int32

const (
	// The ordering is unspecified. Not a valid option.
	Index_IndexField_ORDER_UNSPECIFIED Index_IndexField_Order = 0
	// The field is ordered by ascending field value.
	Index_IndexField_ASCENDING Index_IndexField_Order = 1
	// The field is ordered by descending field value.
	Index_IndexField_DESCENDING Index_IndexField_Order = 2
)

var Index_IndexField_Order_name = map[int32]string{
	0: "ORDER_UNSPECIFIED",
	1: "ASCENDING",
	2: "DESCENDING",
}
var Index_IndexField_Order_value = map[string]int32{
	"ORDER_UNSPECIFIED": 0,
	"ASCENDING":         1,
	"DESCENDING":        2,
}

func (x Index_IndexField_Order) String() string {
	return proto.EnumName(Index_IndexField_Order_name, int32(x))
}
func (Index_IndexField_Order) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_index_47cf3a5e57bc9ffa, []int{0, 0, 0}
}

// The supported array value configurations.
type Index_IndexField_ArrayConfig int32

const (
	// The index does not support additional array queries.
	Index_IndexField_ARRAY_CONFIG_UNSPECIFIED Index_IndexField_ArrayConfig = 0
	// The index supports array containment queries.
	Index_IndexField_CONTAINS Index_IndexField_ArrayConfig = 1
)

var Index_IndexField_ArrayConfig_name = map[int32]string{
	0: "ARRAY_CONFIG_UNSPECIFIED",
	1: "CONTAINS",
}
var Index_IndexField_ArrayConfig_value = map[string]int32{
	"ARRAY_CONFIG_UNSPECIFIED": 0,
	"CONTAINS":                 1,
}

func (x Index_IndexField_ArrayConfig) String() string {
	return proto.EnumName(Index_IndexField_ArrayConfig_name, int32(x))
}
func (Index_IndexField_ArrayConfig) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_index_47cf3a5e57bc9ffa, []int{0, 0, 1}
}

// Cloud Firestore indexes enable simple and complex queries against
// documents in a database.
type Index struct {
	// Output only.
	// A server defined name for this index.
	// The form of this name for composite indexes will be:
	// `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{composite_index_id}`
	// For single field indexes, this field will be empty.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Indexes with a collection query scope specified allow queries
	// against a collection that is the child of a specific document, specified at
	// query time, and that has the same collection id.
	//
	// Indexes with a collection group query scope specified allow queries against
	// all collections descended from a specific document, specified at query
	// time, and that have the same collection id as this index.
	QueryScope Index_QueryScope `protobuf:"varint,2,opt,name=query_scope,json=queryScope,proto3,enum=google.firestore.admin.v1beta2.Index_QueryScope" json:"query_scope,omitempty"`
	// The fields supported by this index.
	//
	// For composite indexes, this is always 2 or more fields.
	// The last field entry is always for the field path `__name__`. If, on
	// creation, `__name__` was not specified as the last field, it will be added
	// automatically with the same direction as that of the last field defined. If
	// the final field in a composite index is not directional, the `__name__`
	// will be ordered ASCENDING (unless explicitly specified).
	//
	// For single field indexes, this will always be exactly one entry with a
	// field path equal to the field path of the associated field.
	Fields []*Index_IndexField `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
	// Output only.
	// The serving state of the index.
	State                Index_State `protobuf:"varint,4,opt,name=state,proto3,enum=google.firestore.admin.v1beta2.Index_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Index) Reset()         { *m = Index{} }
func (m *Index) String() string { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()    {}
func (*Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_index_47cf3a5e57bc9ffa, []int{0}
}
func (m *Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index.Unmarshal(m, b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index.Marshal(b, m, deterministic)
}
func (dst *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(dst, src)
}
func (m *Index) XXX_Size() int {
	return xxx_messageInfo_Index.Size(m)
}
func (m *Index) XXX_DiscardUnknown() {
	xxx_messageInfo_Index.DiscardUnknown(m)
}

var xxx_messageInfo_Index proto.InternalMessageInfo

func (m *Index) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Index) GetQueryScope() Index_QueryScope {
	if m != nil {
		return m.QueryScope
	}
	return Index_QUERY_SCOPE_UNSPECIFIED
}

func (m *Index) GetFields() []*Index_IndexField {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Index) GetState() Index_State {
	if m != nil {
		return m.State
	}
	return Index_STATE_UNSPECIFIED
}

// A field in an index.
// The field_path describes which field is indexed, the value_mode describes
// how the field value is indexed.
type Index_IndexField struct {
	// Can be __name__.
	// For single field indexes, this must match the name of the field or may
	// be omitted.
	FieldPath string `protobuf:"bytes,1,opt,name=field_path,json=fieldPath,proto3" json:"field_path,omitempty"`
	// How the field value is indexed.
	//
	// Types that are valid to be assigned to ValueMode:
	//	*Index_IndexField_Order_
	//	*Index_IndexField_ArrayConfig_
	ValueMode            isIndex_IndexField_ValueMode `protobuf_oneof:"value_mode"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Index_IndexField) Reset()         { *m = Index_IndexField{} }
func (m *Index_IndexField) String() string { return proto.CompactTextString(m) }
func (*Index_IndexField) ProtoMessage()    {}
func (*Index_IndexField) Descriptor() ([]byte, []int) {
	return fileDescriptor_index_47cf3a5e57bc9ffa, []int{0, 0}
}
func (m *Index_IndexField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index_IndexField.Unmarshal(m, b)
}
func (m *Index_IndexField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index_IndexField.Marshal(b, m, deterministic)
}
func (dst *Index_IndexField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index_IndexField.Merge(dst, src)
}
func (m *Index_IndexField) XXX_Size() int {
	return xxx_messageInfo_Index_IndexField.Size(m)
}
func (m *Index_IndexField) XXX_DiscardUnknown() {
	xxx_messageInfo_Index_IndexField.DiscardUnknown(m)
}

var xxx_messageInfo_Index_IndexField proto.InternalMessageInfo

func (m *Index_IndexField) GetFieldPath() string {
	if m != nil {
		return m.FieldPath
	}
	return ""
}

type isIndex_IndexField_ValueMode interface {
	isIndex_IndexField_ValueMode()
}

type Index_IndexField_Order_ struct {
	Order Index_IndexField_Order `protobuf:"varint,2,opt,name=order,proto3,enum=google.firestore.admin.v1beta2.Index_IndexField_Order,oneof"`
}

type Index_IndexField_ArrayConfig_ struct {
	ArrayConfig Index_IndexField_ArrayConfig `protobuf:"varint,3,opt,name=array_config,json=arrayConfig,proto3,enum=google.firestore.admin.v1beta2.Index_IndexField_ArrayConfig,oneof"`
}

func (*Index_IndexField_Order_) isIndex_IndexField_ValueMode() {}

func (*Index_IndexField_ArrayConfig_) isIndex_IndexField_ValueMode() {}

func (m *Index_IndexField) GetValueMode() isIndex_IndexField_ValueMode {
	if m != nil {
		return m.ValueMode
	}
	return nil
}

func (m *Index_IndexField) GetOrder() Index_IndexField_Order {
	if x, ok := m.GetValueMode().(*Index_IndexField_Order_); ok {
		return x.Order
	}
	return Index_IndexField_ORDER_UNSPECIFIED
}

func (m *Index_IndexField) GetArrayConfig() Index_IndexField_ArrayConfig {
	if x, ok := m.GetValueMode().(*Index_IndexField_ArrayConfig_); ok {
		return x.ArrayConfig
	}
	return Index_IndexField_ARRAY_CONFIG_UNSPECIFIED
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Index_IndexField) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Index_IndexField_OneofMarshaler, _Index_IndexField_OneofUnmarshaler, _Index_IndexField_OneofSizer, []interface{}{
		(*Index_IndexField_Order_)(nil),
		(*Index_IndexField_ArrayConfig_)(nil),
	}
}

func _Index_IndexField_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Index_IndexField)
	// value_mode
	switch x := m.ValueMode.(type) {
	case *Index_IndexField_Order_:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Order))
	case *Index_IndexField_ArrayConfig_:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.ArrayConfig))
	case nil:
	default:
		return fmt.Errorf("Index_IndexField.ValueMode has unexpected type %T", x)
	}
	return nil
}

func _Index_IndexField_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Index_IndexField)
	switch tag {
	case 2: // value_mode.order
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueMode = &Index_IndexField_Order_{Index_IndexField_Order(x)}
		return true, err
	case 3: // value_mode.array_config
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueMode = &Index_IndexField_ArrayConfig_{Index_IndexField_ArrayConfig(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Index_IndexField_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Index_IndexField)
	// value_mode
	switch x := m.ValueMode.(type) {
	case *Index_IndexField_Order_:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Order))
	case *Index_IndexField_ArrayConfig_:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.ArrayConfig))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Index)(nil), "google.firestore.admin.v1beta2.Index")
	proto.RegisterType((*Index_IndexField)(nil), "google.firestore.admin.v1beta2.Index.IndexField")
	proto.RegisterEnum("google.firestore.admin.v1beta2.Index_QueryScope", Index_QueryScope_name, Index_QueryScope_value)
	proto.RegisterEnum("google.firestore.admin.v1beta2.Index_State", Index_State_name, Index_State_value)
	proto.RegisterEnum("google.firestore.admin.v1beta2.Index_IndexField_Order", Index_IndexField_Order_name, Index_IndexField_Order_value)
	proto.RegisterEnum("google.firestore.admin.v1beta2.Index_IndexField_ArrayConfig", Index_IndexField_ArrayConfig_name, Index_IndexField_ArrayConfig_value)
}

func init() {
	proto.RegisterFile("google/firestore/admin/v1beta2/index.proto", fileDescriptor_index_47cf3a5e57bc9ffa)
}

var fileDescriptor_index_47cf3a5e57bc9ffa = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6b, 0xdb, 0x4c,
	0x10, 0xc6, 0x23, 0x39, 0x0a, 0xaf, 0xc7, 0x7e, 0x83, 0xba, 0x50, 0x2a, 0xd2, 0xb4, 0x18, 0xd1,
	0x83, 0x69, 0x41, 0x6a, 0x52, 0x28, 0x84, 0xb6, 0x07, 0x59, 0x5a, 0xdb, 0x82, 0x20, 0xc9, 0x2b,
	0xa7, 0xe0, 0x5e, 0xc4, 0xc6, 0x5a, 0x2b, 0x02, 0x5b, 0xeb, 0x48, 0x72, 0x68, 0xbe, 0x4e, 0xa1,
	0x97, 0x7e, 0xb0, 0x1e, 0xfa, 0x29, 0x8a, 0x56, 0xc2, 0x2e, 0xa1, 0x7f, 0x92, 0x8b, 0xd8, 0x91,
	0x9e, 0xe7, 0x37, 0xb3, 0xa3, 0x19, 0x78, 0x99, 0x70, 0x9e, 0x2c, 0x99, 0xb9, 0x48, 0x73, 0x56,
	0x94, 0x3c, 0x67, 0x26, 0x8d, 0x57, 0x69, 0x66, 0xde, 0x9c, 0x5c, 0xb2, 0x92, 0x9e, 0x9a, 0x69,
	0x16, 0xb3, 0xcf, 0xc6, 0x3a, 0xe7, 0x25, 0x47, 0xcf, 0x6b, 0xad, 0xb1, 0xd5, 0x1a, 0x42, 0x6b,
	0x34, 0xda, 0xa3, 0xe3, 0x86, 0x45, 0xd7, 0xa9, 0x49, 0xb3, 0x8c, 0x97, 0xb4, 0x4c, 0x79, 0x56,
	0xd4, 0x6e, 0xfd, 0x87, 0x02, 0x8a, 0x5b, 0xd1, 0x10, 0x82, 0xfd, 0x8c, 0xae, 0x98, 0x26, 0xf5,
	0xa4, 0x7e, 0x9b, 0x88, 0x33, 0x9a, 0x40, 0xe7, 0x7a, 0xc3, 0xf2, 0xdb, 0xa8, 0x98, 0xf3, 0x35,
	0xd3, 0xe4, 0x9e, 0xd4, 0x3f, 0x3c, 0x7d, 0x6d, 0xfc, 0x3d, 0xa3, 0x21, 0x78, 0xc6, 0xa4, 0x32,
	0x86, 0x95, 0x8f, 0xc0, 0xf5, 0xf6, 0x8c, 0xc6, 0x70, 0xb0, 0x48, 0xd9, 0x32, 0x2e, 0xb4, 0x56,
	0xaf, 0xd5, 0xef, 0xdc, 0x97, 0x26, 0x9e, 0xc3, 0xca, 0x48, 0x1a, 0x3f, 0xb2, 0x40, 0x29, 0x4a,
	0x5a, 0x32, 0x6d, 0x5f, 0x94, 0xf5, 0xea, 0x7e, 0xa0, 0xb0, 0xb2, 0x90, 0xda, 0x79, 0xf4, 0x5d,
	0x06, 0xd8, 0x91, 0xd1, 0x33, 0x00, 0xc1, 0x8e, 0xd6, 0xb4, 0xbc, 0x6a, 0x1a, 0xd1, 0x16, 0x6f,
	0x02, 0x5a, 0x5e, 0x21, 0x0f, 0x14, 0x9e, 0xc7, 0x2c, 0x6f, 0xfa, 0xf0, 0xf6, 0xa1, 0x95, 0x1b,
	0x7e, 0xe5, 0x1e, 0xef, 0x91, 0x1a, 0x83, 0x28, 0x74, 0x69, 0x9e, 0xd3, 0xdb, 0x68, 0xce, 0xb3,
	0x45, 0x9a, 0x68, 0x2d, 0x81, 0x7d, 0xff, 0x60, 0xac, 0x55, 0x41, 0x6c, 0xc1, 0x18, 0xef, 0x91,
	0x0e, 0xdd, 0x85, 0xfa, 0x07, 0x50, 0x44, 0x52, 0xf4, 0x18, 0x1e, 0xf9, 0xc4, 0xc1, 0x24, 0xba,
	0xf0, 0xc2, 0x00, 0xdb, 0xee, 0xd0, 0xc5, 0x8e, 0xba, 0x87, 0xfe, 0x87, 0xb6, 0x15, 0xda, 0xd8,
	0x73, 0x5c, 0x6f, 0xa4, 0x4a, 0xe8, 0x10, 0xc0, 0xc1, 0xdb, 0x58, 0xd6, 0xcf, 0xa0, 0xf3, 0x0b,
	0x1c, 0x1d, 0x83, 0x66, 0x11, 0x62, 0xcd, 0x22, 0xdb, 0xf7, 0x86, 0xee, 0xe8, 0x0e, 0xab, 0x0b,
	0xff, 0xd9, 0xbe, 0x37, 0xb5, 0x5c, 0x2f, 0x54, 0xa5, 0x41, 0x17, 0xe0, 0x86, 0x2e, 0x37, 0x2c,
	0x5a, 0xf1, 0x98, 0xe9, 0x67, 0x00, 0xbb, 0x79, 0x40, 0x4f, 0xe1, 0xc9, 0xe4, 0x02, 0x93, 0x59,
	0x14, 0xda, 0x7e, 0x80, 0xef, 0x60, 0x0e, 0x01, 0x6c, 0xff, 0xfc, 0x1c, 0xdb, 0x53, 0xd7, 0xf7,
	0x54, 0x49, 0x77, 0x41, 0x11, 0xff, 0xac, 0xba, 0x42, 0x38, 0xb5, 0xa6, 0xf8, 0x37, 0x69, 0x09,
	0xb6, 0xa6, 0xf5, 0x0d, 0xda, 0xa0, 0x10, 0x6c, 0x39, 0x33, 0x55, 0x46, 0x2a, 0x74, 0x3d, 0x8c,
	0x9d, 0x30, 0x22, 0x38, 0xb0, 0x5c, 0xa2, 0xb6, 0x06, 0x5f, 0x25, 0xd0, 0xe7, 0x7c, 0xf5, 0x8f,
	0x06, 0x0f, 0xea, 0x91, 0x08, 0xaa, 0xfd, 0x08, 0xa4, 0x4f, 0x76, 0xa3, 0x4e, 0xf8, 0x92, 0x66,
	0x89, 0xc1, 0xf3, 0xc4, 0x4c, 0x58, 0x26, 0xb6, 0xc7, 0xac, 0x3f, 0xd1, 0x75, 0x5a, 0xfc, 0x69,
	0x55, 0xdf, 0x89, 0xe8, 0x8b, 0xbc, 0x3f, 0xb2, 0x87, 0xe1, 0x37, 0xf9, 0xc5, 0xa8, 0x86, 0xd9,
	0x4b, 0xbe, 0x89, 0x8d, 0xe1, 0xb6, 0x00, 0x4b, 0x14, 0xf0, 0xf1, 0x64, 0x50, 0x79, 0x2e, 0x0f,
	0x04, 0xfd, 0xcd, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xcd, 0x75, 0x16, 0x07, 0x04, 0x00,
	0x00,
}
