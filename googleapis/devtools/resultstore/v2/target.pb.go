// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/resultstore/v2/target.proto

package resultstore // import "github.com/sunnogo/go-genproto/googleapis/devtools/resultstore/v2"

import proto "github.com/sunnogo/protobuf/proto"
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

// These correspond to the suffix of the rule name. Eg cc_test has type TEST.
type TargetType int32

const (
	// Unspecified by the build system.
	TargetType_TARGET_TYPE_UNSPECIFIED TargetType = 0
	// An application e.g. ios_application.
	TargetType_APPLICATION TargetType = 1
	// A binary target e.g. cc_binary.
	TargetType_BINARY TargetType = 2
	// A library target e.g. java_library
	TargetType_LIBRARY TargetType = 3
	// A package
	TargetType_PACKAGE TargetType = 4
	// Any test target, in bazel that means a rule with a '_test' suffix.
	TargetType_TEST TargetType = 5
)

var TargetType_name = map[int32]string{
	0: "TARGET_TYPE_UNSPECIFIED",
	1: "APPLICATION",
	2: "BINARY",
	3: "LIBRARY",
	4: "PACKAGE",
	5: "TEST",
}
var TargetType_value = map[string]int32{
	"TARGET_TYPE_UNSPECIFIED": 0,
	"APPLICATION":             1,
	"BINARY":                  2,
	"LIBRARY":                 3,
	"PACKAGE":                 4,
	"TEST":                    5,
}

func (x TargetType) String() string {
	return proto.EnumName(TargetType_name, int32(x))
}
func (TargetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_target_25848505c9398e1a, []int{0}
}

// Indicates how big the user indicated the test action was.
type TestSize int32

const (
	// Unspecified by the user.
	TestSize_TEST_SIZE_UNSPECIFIED TestSize = 0
	// Unit test taking less than 1 minute.
	TestSize_SMALL TestSize = 1
	// Integration tests taking less than 5 minutes.
	TestSize_MEDIUM TestSize = 2
	// End-to-end tests taking less than 15 minutes.
	TestSize_LARGE TestSize = 3
	// Even bigger than LARGE.
	TestSize_ENORMOUS TestSize = 4
	// Something that doesn't fit into the above categories.
	TestSize_OTHER_SIZE TestSize = 5
)

var TestSize_name = map[int32]string{
	0: "TEST_SIZE_UNSPECIFIED",
	1: "SMALL",
	2: "MEDIUM",
	3: "LARGE",
	4: "ENORMOUS",
	5: "OTHER_SIZE",
}
var TestSize_value = map[string]int32{
	"TEST_SIZE_UNSPECIFIED": 0,
	"SMALL":                 1,
	"MEDIUM":                2,
	"LARGE":                 3,
	"ENORMOUS":              4,
	"OTHER_SIZE":            5,
}

func (x TestSize) String() string {
	return proto.EnumName(TestSize_name, int32(x))
}
func (TestSize) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_target_25848505c9398e1a, []int{1}
}

// Each Target represents data for a given target in a given Invocation.
// ConfiguredTarget and Action resources under each Target contain the bulk of
// the data.
type Target struct {
	// The resource name.  Its format must be:
	// invocations/${INVOCATION_ID}/targets/${TARGET_ID}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resource ID components that identify the Target. They must match the
	// resource name after proper encoding.
	Id *Target_Id `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// This is the aggregate status of the target.
	StatusAttributes *StatusAttributes `protobuf:"bytes,3,opt,name=status_attributes,json=statusAttributes,proto3" json:"status_attributes,omitempty"`
	// When this target started and its duration.
	Timing *Timing `protobuf:"bytes,4,opt,name=timing,proto3" json:"timing,omitempty"`
	// Attributes that apply to all targets.
	TargetAttributes *TargetAttributes `protobuf:"bytes,5,opt,name=target_attributes,json=targetAttributes,proto3" json:"target_attributes,omitempty"`
	// Attributes that apply to all test actions under this target.
	TestAttributes *TestAttributes `protobuf:"bytes,6,opt,name=test_attributes,json=testAttributes,proto3" json:"test_attributes,omitempty"`
	// Arbitrary name-value pairs.
	// This is implemented as a multi-map. Multiple properties are allowed with
	// the same key. Properties will be returned in lexicographical order by key.
	Properties []*Property `protobuf:"bytes,7,rep,name=properties,proto3" json:"properties,omitempty"`
	// A list of file references for target level files.
	// The file IDs must be unique within this list. Duplicate file IDs will
	// result in an error. Files will be returned in lexicographical order by ID.
	// Use this field to specify outputs not related to a configuration.
	Files []*File `protobuf:"bytes,8,rep,name=files,proto3" json:"files,omitempty"`
	// Provides a hint to clients as to whether to display the Target to users.
	// If true then clients likely want to display the Target by default.
	// Once set to true, this may not be set back to false.
	IsVisible            bool     `protobuf:"varint,9,opt,name=is_visible,json=isVisible,proto3" json:"is_visible,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Target) Reset()         { *m = Target{} }
func (m *Target) String() string { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()    {}
func (*Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_target_25848505c9398e1a, []int{0}
}
func (m *Target) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Target.Unmarshal(m, b)
}
func (m *Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Target.Marshal(b, m, deterministic)
}
func (dst *Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Target.Merge(dst, src)
}
func (m *Target) XXX_Size() int {
	return xxx_messageInfo_Target.Size(m)
}
func (m *Target) XXX_DiscardUnknown() {
	xxx_messageInfo_Target.DiscardUnknown(m)
}

var xxx_messageInfo_Target proto.InternalMessageInfo

func (m *Target) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Target) GetId() *Target_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Target) GetStatusAttributes() *StatusAttributes {
	if m != nil {
		return m.StatusAttributes
	}
	return nil
}

func (m *Target) GetTiming() *Timing {
	if m != nil {
		return m.Timing
	}
	return nil
}

func (m *Target) GetTargetAttributes() *TargetAttributes {
	if m != nil {
		return m.TargetAttributes
	}
	return nil
}

func (m *Target) GetTestAttributes() *TestAttributes {
	if m != nil {
		return m.TestAttributes
	}
	return nil
}

func (m *Target) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Target) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *Target) GetIsVisible() bool {
	if m != nil {
		return m.IsVisible
	}
	return false
}

// The resource ID components that identify the Target.
type Target_Id struct {
	// The Invocation ID.
	InvocationId string `protobuf:"bytes,1,opt,name=invocation_id,json=invocationId,proto3" json:"invocation_id,omitempty"`
	// The Target ID.
	TargetId             string   `protobuf:"bytes,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Target_Id) Reset()         { *m = Target_Id{} }
func (m *Target_Id) String() string { return proto.CompactTextString(m) }
func (*Target_Id) ProtoMessage()    {}
func (*Target_Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_target_25848505c9398e1a, []int{0, 0}
}
func (m *Target_Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Target_Id.Unmarshal(m, b)
}
func (m *Target_Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Target_Id.Marshal(b, m, deterministic)
}
func (dst *Target_Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Target_Id.Merge(dst, src)
}
func (m *Target_Id) XXX_Size() int {
	return xxx_messageInfo_Target_Id.Size(m)
}
func (m *Target_Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Target_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Target_Id proto.InternalMessageInfo

func (m *Target_Id) GetInvocationId() string {
	if m != nil {
		return m.InvocationId
	}
	return ""
}

func (m *Target_Id) GetTargetId() string {
	if m != nil {
		return m.TargetId
	}
	return ""
}

// Attributes that apply to all targets.
type TargetAttributes struct {
	// If known, indicates the type of this target.  In bazel this corresponds
	// to the rule-suffix.
	Type TargetType `protobuf:"varint,1,opt,name=type,proto3,enum=google.devtools.resultstore.v2.TargetType" json:"type,omitempty"`
	// If known, the main language of this target, e.g. java, cc, python, etc.
	Language Language `protobuf:"varint,2,opt,name=language,proto3,enum=google.devtools.resultstore.v2.Language" json:"language,omitempty"`
	// The tags attribute of the build rule. These should be short, descriptive
	// words, and there should only be a few of them.
	// This is implemented as a set. All tags will be unique. Any duplicate tags
	// will be ignored. Tags will be returned in lexicographical order.
	Tags                 []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TargetAttributes) Reset()         { *m = TargetAttributes{} }
func (m *TargetAttributes) String() string { return proto.CompactTextString(m) }
func (*TargetAttributes) ProtoMessage()    {}
func (*TargetAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_target_25848505c9398e1a, []int{1}
}
func (m *TargetAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetAttributes.Unmarshal(m, b)
}
func (m *TargetAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetAttributes.Marshal(b, m, deterministic)
}
func (dst *TargetAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetAttributes.Merge(dst, src)
}
func (m *TargetAttributes) XXX_Size() int {
	return xxx_messageInfo_TargetAttributes.Size(m)
}
func (m *TargetAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_TargetAttributes proto.InternalMessageInfo

func (m *TargetAttributes) GetType() TargetType {
	if m != nil {
		return m.Type
	}
	return TargetType_TARGET_TYPE_UNSPECIFIED
}

func (m *TargetAttributes) GetLanguage() Language {
	if m != nil {
		return m.Language
	}
	return Language_LANGUAGE_UNSPECIFIED
}

func (m *TargetAttributes) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// Attributes that apply only to test actions under this target.
type TestAttributes struct {
	// Indicates how big the user indicated the test action was.
	Size                 TestSize `protobuf:"varint,1,opt,name=size,proto3,enum=google.devtools.resultstore.v2.TestSize" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestAttributes) Reset()         { *m = TestAttributes{} }
func (m *TestAttributes) String() string { return proto.CompactTextString(m) }
func (*TestAttributes) ProtoMessage()    {}
func (*TestAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_target_25848505c9398e1a, []int{2}
}
func (m *TestAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestAttributes.Unmarshal(m, b)
}
func (m *TestAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestAttributes.Marshal(b, m, deterministic)
}
func (dst *TestAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestAttributes.Merge(dst, src)
}
func (m *TestAttributes) XXX_Size() int {
	return xxx_messageInfo_TestAttributes.Size(m)
}
func (m *TestAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_TestAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_TestAttributes proto.InternalMessageInfo

func (m *TestAttributes) GetSize() TestSize {
	if m != nil {
		return m.Size
	}
	return TestSize_TEST_SIZE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*Target)(nil), "google.devtools.resultstore.v2.Target")
	proto.RegisterType((*Target_Id)(nil), "google.devtools.resultstore.v2.Target.Id")
	proto.RegisterType((*TargetAttributes)(nil), "google.devtools.resultstore.v2.TargetAttributes")
	proto.RegisterType((*TestAttributes)(nil), "google.devtools.resultstore.v2.TestAttributes")
	proto.RegisterEnum("google.devtools.resultstore.v2.TargetType", TargetType_name, TargetType_value)
	proto.RegisterEnum("google.devtools.resultstore.v2.TestSize", TestSize_name, TestSize_value)
}

func init() {
	proto.RegisterFile("google/devtools/resultstore/v2/target.proto", fileDescriptor_target_25848505c9398e1a)
}

var fileDescriptor_target_25848505c9398e1a = []byte{
	// 634 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xed, 0x6a, 0xdb, 0x4a,
	0x10, 0xbd, 0xb2, 0x65, 0x47, 0x1e, 0xe7, 0x3a, 0xba, 0x0b, 0x97, 0xaa, 0x09, 0x2d, 0xc6, 0x2d,
	0xc5, 0x49, 0x41, 0x2e, 0xee, 0xaf, 0x7e, 0x10, 0x50, 0x12, 0x25, 0x11, 0xf5, 0x87, 0x58, 0x29,
	0x2d, 0x09, 0x14, 0xa3, 0x44, 0x5b, 0xb1, 0x45, 0xd6, 0xba, 0xda, 0xb5, 0x21, 0x79, 0x95, 0xbe,
	0x42, 0x1f, 0xb2, 0x68, 0xd7, 0x49, 0x1c, 0x43, 0xab, 0xfc, 0xdb, 0x19, 0x9d, 0x73, 0xe6, 0x68,
	0x66, 0x76, 0xe1, 0x75, 0xc2, 0x58, 0x92, 0x92, 0x5e, 0x4c, 0x16, 0x82, 0xb1, 0x94, 0xf7, 0x72,
	0xc2, 0xe7, 0xa9, 0xe0, 0x82, 0xe5, 0xa4, 0xb7, 0xe8, 0xf7, 0x44, 0x94, 0x27, 0x44, 0xd8, 0xb3,
	0x9c, 0x09, 0x86, 0x9e, 0x2b, 0xb0, 0x7d, 0x0b, 0xb6, 0x57, 0xc0, 0xf6, 0xa2, 0xbf, 0x5d, 0x26,
	0x76, 0xc5, 0xa6, 0x53, 0x96, 0x29, 0xb1, 0xed, 0xdd, 0x12, 0xf0, 0x37, 0x9a, 0x12, 0x05, 0xed,
	0xfc, 0xac, 0x41, 0x3d, 0x94, 0x46, 0x10, 0x02, 0x3d, 0x8b, 0xa6, 0xc4, 0xd2, 0xda, 0x5a, 0xb7,
	0x81, 0xe5, 0x19, 0xbd, 0x83, 0x0a, 0x8d, 0xad, 0x4a, 0x5b, 0xeb, 0x36, 0xfb, 0xbb, 0xf6, 0xdf,
	0x3d, 0xda, 0x4a, 0xc7, 0xf6, 0x62, 0x5c, 0xa1, 0x31, 0xfa, 0x0a, 0xff, 0x71, 0x11, 0x89, 0x39,
	0x9f, 0x44, 0x42, 0xe4, 0xf4, 0x72, 0x2e, 0x08, 0xb7, 0xaa, 0x52, 0xe9, 0x4d, 0x99, 0x52, 0x20,
	0x89, 0xce, 0x1d, 0x0f, 0x9b, 0x7c, 0x2d, 0x83, 0xf6, 0xa1, 0x2e, 0xe8, 0x94, 0x66, 0x89, 0xa5,
	0x4b, 0xcd, 0x57, 0xa5, 0xee, 0x24, 0x1a, 0x2f, 0x59, 0x85, 0x3d, 0x35, 0x80, 0x55, 0x7b, 0xb5,
	0xc7, 0xd9, 0x53, 0x3f, 0xba, 0x6a, 0x4f, 0xac, 0x65, 0xd0, 0x17, 0xd8, 0x12, 0x84, 0x3f, 0x10,
	0xaf, 0x4b, 0x71, 0xbb, 0x54, 0x9c, 0xf0, 0x55, 0xe9, 0x96, 0x78, 0x10, 0xa3, 0x53, 0x80, 0x59,
	0xce, 0x66, 0x24, 0x17, 0x94, 0x70, 0x6b, 0xa3, 0x5d, 0xed, 0x36, 0xfb, 0xdd, 0x32, 0x4d, 0x5f,
	0x31, 0xae, 0xf1, 0x0a, 0x17, 0xbd, 0x87, 0x5a, 0xb1, 0x08, 0xdc, 0x32, 0xa4, 0xc8, 0xcb, 0x32,
	0x91, 0x63, 0x9a, 0x12, 0xac, 0x28, 0xe8, 0x19, 0x00, 0xe5, 0x93, 0x05, 0xe5, 0xf4, 0x32, 0x25,
	0x56, 0xa3, 0xad, 0x75, 0x0d, 0xdc, 0xa0, 0xfc, 0xb3, 0x4a, 0x6c, 0x1f, 0x43, 0xc5, 0x8b, 0xd1,
	0x0b, 0xf8, 0x97, 0x66, 0x0b, 0x76, 0x15, 0x09, 0xca, 0xb2, 0x09, 0x8d, 0x97, 0x9b, 0xb5, 0x79,
	0x9f, 0xf4, 0x62, 0xb4, 0x03, 0x8d, 0xe5, 0x1c, 0x96, 0x8b, 0xd6, 0xc0, 0x86, 0x4a, 0x78, 0x71,
	0xe7, 0x97, 0x06, 0xe6, 0x7a, 0xb3, 0xd1, 0x3e, 0xe8, 0xe2, 0x7a, 0xa6, 0xf6, 0xb4, 0xd5, 0xdf,
	0x7b, 0xdc, 0xb0, 0xc2, 0xeb, 0x19, 0xc1, 0x92, 0x87, 0x8e, 0xc0, 0x48, 0xa3, 0x2c, 0x99, 0x47,
	0x09, 0x91, 0x05, 0x5b, 0xe5, 0xfd, 0x1b, 0x2c, 0xf1, 0xf8, 0x8e, 0x59, 0xdc, 0x16, 0x11, 0x25,
	0xc5, 0x46, 0x57, 0x8b, 0xdb, 0x52, 0x9c, 0x3b, 0x23, 0x68, 0x3d, 0x9c, 0x1e, 0xfa, 0x08, 0x3a,
	0xa7, 0x37, 0xb7, 0x5e, 0xbb, 0x8f, 0x99, 0x7d, 0x40, 0x6f, 0x08, 0x96, 0xac, 0xbd, 0xef, 0x00,
	0xf7, 0xee, 0xd1, 0x0e, 0x3c, 0x09, 0x1d, 0x7c, 0xe2, 0x86, 0x93, 0xf0, 0xdc, 0x77, 0x27, 0x67,
	0xa3, 0xc0, 0x77, 0x0f, 0xbd, 0x63, 0xcf, 0x3d, 0x32, 0xff, 0x41, 0x5b, 0xd0, 0x74, 0x7c, 0x7f,
	0xe0, 0x1d, 0x3a, 0xa1, 0x37, 0x1e, 0x99, 0x1a, 0x02, 0xa8, 0x1f, 0x78, 0x23, 0x07, 0x9f, 0x9b,
	0x15, 0xd4, 0x84, 0x8d, 0x81, 0x77, 0x80, 0x8b, 0xa0, 0x5a, 0x04, 0xbe, 0x73, 0xf8, 0xc9, 0x39,
	0x71, 0x4d, 0x1d, 0x19, 0xa0, 0x87, 0x6e, 0x10, 0x9a, 0xb5, 0x3d, 0x02, 0xc6, 0x6d, 0x75, 0xf4,
	0x14, 0xfe, 0x2f, 0xb2, 0x93, 0xc0, 0xbb, 0x58, 0xaf, 0xd3, 0x80, 0x5a, 0x30, 0x74, 0x06, 0x03,
	0x55, 0x61, 0xe8, 0x1e, 0x79, 0x67, 0x43, 0xb3, 0x52, 0xa4, 0x07, 0x85, 0x37, 0xb3, 0x8a, 0x36,
	0xc1, 0x70, 0x47, 0x63, 0x3c, 0x1c, 0x9f, 0x05, 0xa6, 0x8e, 0x5a, 0x00, 0xe3, 0xf0, 0xd4, 0xc5,
	0x52, 0xcb, 0xac, 0x1d, 0xfc, 0x80, 0xce, 0x15, 0x9b, 0x96, 0xf4, 0xc1, 0xd7, 0x2e, 0xbc, 0x25,
	0x22, 0x61, 0x45, 0xc7, 0x6d, 0x96, 0x27, 0xbd, 0x84, 0x64, 0xf2, 0xcd, 0xea, 0xa9, 0x4f, 0xd1,
	0x8c, 0xf2, 0x3f, 0xbd, 0x70, 0x1f, 0x56, 0xc2, 0xcb, 0xba, 0x64, 0xbd, 0xfd, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0xc9, 0x5f, 0x96, 0x8a, 0x90, 0x05, 0x00, 0x00,
}
