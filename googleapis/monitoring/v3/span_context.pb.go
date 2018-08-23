// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/v3/span_context.proto

package monitoring // import "github.com/sunnogo/go-genproto/googleapis/monitoring/v3"

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

// The context of a span, attached to google.api.Distribution.Exemplars
// in google.api.Distribution values during aggregation.
//
// It contains the name of a span with format:
//     projects/[PROJECT_ID]/traces/[TRACE_ID]/spans/[SPAN_ID]
type SpanContext struct {
	// The resource name of the span in the following format:
	//
	//     projects/[PROJECT_ID]/traces/[TRACE_ID]/spans/[SPAN_ID]
	//
	// [TRACE_ID] is a unique identifier for a trace within a project;
	// it is a 32-character hexadecimal encoding of a 16-byte array.
	//
	// [SPAN_ID] is a unique identifier for a span within a trace; it
	// is a 16-character hexadecimal encoding of an 8-byte array.
	SpanName             string   `protobuf:"bytes,1,opt,name=span_name,json=spanName,proto3" json:"span_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpanContext) Reset()         { *m = SpanContext{} }
func (m *SpanContext) String() string { return proto.CompactTextString(m) }
func (*SpanContext) ProtoMessage()    {}
func (*SpanContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_span_context_3190ff95558fbb1a, []int{0}
}
func (m *SpanContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanContext.Unmarshal(m, b)
}
func (m *SpanContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanContext.Marshal(b, m, deterministic)
}
func (dst *SpanContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanContext.Merge(dst, src)
}
func (m *SpanContext) XXX_Size() int {
	return xxx_messageInfo_SpanContext.Size(m)
}
func (m *SpanContext) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanContext.DiscardUnknown(m)
}

var xxx_messageInfo_SpanContext proto.InternalMessageInfo

func (m *SpanContext) GetSpanName() string {
	if m != nil {
		return m.SpanName
	}
	return ""
}

func init() {
	proto.RegisterType((*SpanContext)(nil), "google.monitoring.v3.SpanContext")
}

func init() {
	proto.RegisterFile("google/monitoring/v3/span_context.proto", fileDescriptor_span_context_3190ff95558fbb1a)
}

var fileDescriptor_span_context_3190ff95558fbb1a = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0xcf, 0xcd, 0xcf, 0xcb, 0x2c, 0xc9, 0x2f, 0xca, 0xcc, 0x4b, 0xd7, 0x2f, 0x33,
	0xd6, 0x2f, 0x2e, 0x48, 0xcc, 0x8b, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xad, 0x28, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0x81, 0x28, 0xd4, 0x43, 0x28, 0xd4, 0x2b, 0x33, 0x56, 0xd2, 0xe2,
	0xe2, 0x0e, 0x2e, 0x48, 0xcc, 0x73, 0x86, 0x28, 0x15, 0x92, 0xe6, 0xe2, 0x04, 0x6b, 0xcd, 0x4b,
	0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x00, 0x09, 0xf8, 0x25, 0xe6, 0xa6,
	0x3a, 0xad, 0x60, 0xe4, 0x92, 0x48, 0xce, 0xcf, 0xd5, 0xc3, 0x66, 0x90, 0x93, 0x00, 0x92, 0x31,
	0x01, 0x20, 0x0b, 0x03, 0x18, 0xa3, 0xec, 0xa0, 0x2a, 0xd3, 0xf3, 0x73, 0x12, 0xf3, 0xd2, 0xf5,
	0xf2, 0x8b, 0xd2, 0xf5, 0xd3, 0x53, 0xf3, 0xc0, 0xce, 0xd1, 0x87, 0x48, 0x25, 0x16, 0x64, 0x16,
	0xa3, 0x3a, 0xdd, 0x1a, 0xc1, 0x5b, 0xc5, 0x24, 0xe5, 0x0e, 0x31, 0xc0, 0x39, 0x27, 0xbf, 0x34,
	0x45, 0xcf, 0x17, 0x61, 0x61, 0x98, 0xf1, 0x29, 0x98, 0x64, 0x0c, 0x58, 0x32, 0x06, 0x21, 0x19,
	0x13, 0x66, 0x9c, 0xc4, 0x06, 0xb6, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x19, 0x01,
	0xcb, 0x1e, 0x01, 0x00, 0x00,
}
