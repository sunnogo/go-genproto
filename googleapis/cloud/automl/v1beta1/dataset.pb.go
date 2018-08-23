// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/dataset.proto

package automl // import "github.com/sunnogo/go-genproto/googleapis/cloud/automl/v1beta1"

import proto "github.com/sunnogo/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/sunnogo/protobuf/ptypes/timestamp"
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

// A workspace for solving a single, particular machine learning (ML) problem.
// A workspace contains examples that may be annotated.
type Dataset struct {
	// Required.
	// The dataset metadata that is specific to the problem type.
	//
	// Types that are valid to be assigned to DatasetMetadata:
	//	*Dataset_TranslationDatasetMetadata
	//	*Dataset_ImageClassificationDatasetMetadata
	//	*Dataset_TextClassificationDatasetMetadata
	DatasetMetadata isDataset_DatasetMetadata `protobuf_oneof:"dataset_metadata"`
	// Output only. The resource name of the dataset.
	// Form: `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The name of the dataset to show in the interface. The name can be
	// up to 32 characters
	// long and can consist only of ASCII Latin letters A-Z and a-z, underscores
	// (_), and ASCII digits 0-9.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. The number of examples in the dataset.
	ExampleCount int32 `protobuf:"varint,21,opt,name=example_count,json=exampleCount,proto3" json:"example_count,omitempty"`
	// Output only. Timestamp when this dataset was created.
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,14,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Dataset) Reset()         { *m = Dataset{} }
func (m *Dataset) String() string { return proto.CompactTextString(m) }
func (*Dataset) ProtoMessage()    {}
func (*Dataset) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataset_a8678ffeb33098d6, []int{0}
}
func (m *Dataset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dataset.Unmarshal(m, b)
}
func (m *Dataset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dataset.Marshal(b, m, deterministic)
}
func (dst *Dataset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dataset.Merge(dst, src)
}
func (m *Dataset) XXX_Size() int {
	return xxx_messageInfo_Dataset.Size(m)
}
func (m *Dataset) XXX_DiscardUnknown() {
	xxx_messageInfo_Dataset.DiscardUnknown(m)
}

var xxx_messageInfo_Dataset proto.InternalMessageInfo

type isDataset_DatasetMetadata interface {
	isDataset_DatasetMetadata()
}

type Dataset_TranslationDatasetMetadata struct {
	TranslationDatasetMetadata *TranslationDatasetMetadata `protobuf:"bytes,23,opt,name=translation_dataset_metadata,json=translationDatasetMetadata,proto3,oneof"`
}

type Dataset_ImageClassificationDatasetMetadata struct {
	ImageClassificationDatasetMetadata *ImageClassificationDatasetMetadata `protobuf:"bytes,24,opt,name=image_classification_dataset_metadata,json=imageClassificationDatasetMetadata,proto3,oneof"`
}

type Dataset_TextClassificationDatasetMetadata struct {
	TextClassificationDatasetMetadata *TextClassificationDatasetMetadata `protobuf:"bytes,25,opt,name=text_classification_dataset_metadata,json=textClassificationDatasetMetadata,proto3,oneof"`
}

func (*Dataset_TranslationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_ImageClassificationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TextClassificationDatasetMetadata) isDataset_DatasetMetadata() {}

func (m *Dataset) GetDatasetMetadata() isDataset_DatasetMetadata {
	if m != nil {
		return m.DatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTranslationDatasetMetadata() *TranslationDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TranslationDatasetMetadata); ok {
		return x.TranslationDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetImageClassificationDatasetMetadata() *ImageClassificationDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_ImageClassificationDatasetMetadata); ok {
		return x.ImageClassificationDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTextClassificationDatasetMetadata() *TextClassificationDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TextClassificationDatasetMetadata); ok {
		return x.TextClassificationDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Dataset) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Dataset) GetExampleCount() int32 {
	if m != nil {
		return m.ExampleCount
	}
	return 0
}

func (m *Dataset) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Dataset) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Dataset_OneofMarshaler, _Dataset_OneofUnmarshaler, _Dataset_OneofSizer, []interface{}{
		(*Dataset_TranslationDatasetMetadata)(nil),
		(*Dataset_ImageClassificationDatasetMetadata)(nil),
		(*Dataset_TextClassificationDatasetMetadata)(nil),
	}
}

func _Dataset_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Dataset)
	// dataset_metadata
	switch x := m.DatasetMetadata.(type) {
	case *Dataset_TranslationDatasetMetadata:
		b.EncodeVarint(23<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TranslationDatasetMetadata); err != nil {
			return err
		}
	case *Dataset_ImageClassificationDatasetMetadata:
		b.EncodeVarint(24<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ImageClassificationDatasetMetadata); err != nil {
			return err
		}
	case *Dataset_TextClassificationDatasetMetadata:
		b.EncodeVarint(25<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TextClassificationDatasetMetadata); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Dataset.DatasetMetadata has unexpected type %T", x)
	}
	return nil
}

func _Dataset_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Dataset)
	switch tag {
	case 23: // dataset_metadata.translation_dataset_metadata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TranslationDatasetMetadata)
		err := b.DecodeMessage(msg)
		m.DatasetMetadata = &Dataset_TranslationDatasetMetadata{msg}
		return true, err
	case 24: // dataset_metadata.image_classification_dataset_metadata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ImageClassificationDatasetMetadata)
		err := b.DecodeMessage(msg)
		m.DatasetMetadata = &Dataset_ImageClassificationDatasetMetadata{msg}
		return true, err
	case 25: // dataset_metadata.text_classification_dataset_metadata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TextClassificationDatasetMetadata)
		err := b.DecodeMessage(msg)
		m.DatasetMetadata = &Dataset_TextClassificationDatasetMetadata{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Dataset_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Dataset)
	// dataset_metadata
	switch x := m.DatasetMetadata.(type) {
	case *Dataset_TranslationDatasetMetadata:
		s := proto.Size(x.TranslationDatasetMetadata)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Dataset_ImageClassificationDatasetMetadata:
		s := proto.Size(x.ImageClassificationDatasetMetadata)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Dataset_TextClassificationDatasetMetadata:
		s := proto.Size(x.TextClassificationDatasetMetadata)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Dataset)(nil), "google.cloud.automl.v1beta1.Dataset")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/dataset.proto", fileDescriptor_dataset_a8678ffeb33098d6)
}

var fileDescriptor_dataset_a8678ffeb33098d6 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x5d, 0x6b, 0x14, 0x31,
	0x14, 0x86, 0x1d, 0xa9, 0x8a, 0xd9, 0x2a, 0x12, 0x10, 0xc7, 0xb1, 0xd0, 0x6d, 0xfd, 0x5a, 0x41,
	0x13, 0xaa, 0x82, 0x17, 0x05, 0xc5, 0xd6, 0x0b, 0xbd, 0x50, 0x64, 0xe9, 0x95, 0x37, 0xe1, 0xec,
	0x6c, 0x76, 0x08, 0xe4, 0x63, 0xd8, 0x9c, 0x91, 0x2d, 0xfe, 0x0a, 0xc1, 0xdf, 0xe7, 0x6f, 0x91,
	0x49, 0x32, 0xd4, 0xaf, 0xcd, 0xdc, 0x25, 0xe7, 0x3c, 0x79, 0xf3, 0x9e, 0x97, 0x43, 0x9e, 0x34,
	0xce, 0x35, 0x5a, 0xf2, 0x5a, 0xbb, 0x6e, 0xc9, 0xa1, 0x43, 0x67, 0x34, 0xff, 0x7a, 0xb4, 0x90,
	0x08, 0x47, 0x7c, 0x09, 0x08, 0x5e, 0x22, 0x6b, 0xd7, 0x0e, 0x1d, 0xbd, 0x17, 0x51, 0x16, 0x50,
	0x16, 0x51, 0x96, 0xd0, 0x6a, 0x2f, 0xe9, 0x40, 0xab, 0x38, 0x58, 0xeb, 0x10, 0x50, 0x39, 0xeb,
	0xe3, 0xd3, 0xea, 0x65, 0xee, 0x97, 0x0b, 0x5c, 0xb4, 0x70, 0xae, 0x1d, 0x2c, 0xd3, 0xab, 0xa7,
	0x63, 0xde, 0x84, 0x42, 0x69, 0x86, 0x3f, 0x1e, 0xe7, 0x68, 0x65, 0xa0, 0x91, 0x09, 0x7c, 0x94,
	0x03, 0x51, 0x6e, 0xd2, 0xbc, 0xd5, 0xb3, 0x2c, 0xb7, 0x06, 0xeb, 0x75, 0x70, 0x9d, 0xf0, 0xfd,
	0x84, 0x87, 0xdb, 0xa2, 0x5b, 0x71, 0x54, 0x46, 0x7a, 0x04, 0xd3, 0x46, 0xe0, 0xf0, 0xe7, 0x0e,
	0xb9, 0xf6, 0x2e, 0x26, 0x4a, 0xbf, 0x91, 0xbd, 0xdf, 0x14, 0x44, 0x0a, 0x5a, 0x18, 0x89, 0xd0,
	0x9f, 0xcb, 0x3b, 0xd3, 0x62, 0x36, 0x79, 0xfe, 0x8a, 0x65, 0x22, 0x67, 0x67, 0x17, 0x02, 0x49,
	0xf6, 0x63, 0x7a, 0xfe, 0xfe, 0xd2, 0xbc, 0xc2, 0xad, 0x5d, 0xfa, 0xa3, 0x20, 0x0f, 0x43, 0x20,
	0xa2, 0xd6, 0xe0, 0xbd, 0x5a, 0xa9, 0x7a, 0x8b, 0x8d, 0x32, 0xd8, 0x78, 0x93, 0xb5, 0xf1, 0xa1,
	0x57, 0x3a, 0xfd, 0x43, 0xe8, 0x5f, 0x3b, 0x87, 0x6a, 0x94, 0xa2, 0xdf, 0x0b, 0xf2, 0xa0, 0x8f,
	0x7f, 0xd4, 0xd5, 0xdd, 0xe0, 0xea, 0x75, 0x3e, 0x1c, 0xb9, 0xc1, 0x31, 0x53, 0x07, 0x38, 0x06,
	0x51, 0x4a, 0x76, 0x2c, 0x18, 0x59, 0x16, 0xd3, 0x62, 0x76, 0x7d, 0x1e, 0xce, 0xf4, 0x80, 0xec,
	0x2e, 0x95, 0x6f, 0x35, 0x9c, 0x8b, 0xd0, 0xbb, 0x1c, 0x7a, 0x93, 0x54, 0xfb, 0xd4, 0x23, 0xf7,
	0xc9, 0x0d, 0xb9, 0x01, 0xd3, 0x6a, 0x29, 0x6a, 0xd7, 0x59, 0x2c, 0x6f, 0x4f, 0x8b, 0xd9, 0x95,
	0xf9, 0x6e, 0x2a, 0x9e, 0xf6, 0x35, 0x7a, 0x4c, 0x26, 0xf5, 0x5a, 0x02, 0x4a, 0xd1, 0x6f, 0x4a,
	0x79, 0x33, 0x4c, 0x55, 0x0d, 0x53, 0x0d, 0x6b, 0xc4, 0xce, 0x86, 0x35, 0x9a, 0x93, 0x88, 0xf7,
	0x85, 0x13, 0x4a, 0x6e, 0xfd, 0x9d, 0xcb, 0xc9, 0x8a, 0xec, 0xd7, 0xce, 0xe4, 0x62, 0xf9, 0x5c,
	0x7c, 0x79, 0x9b, 0xda, 0x8d, 0xd3, 0x60, 0x1b, 0xe6, 0xd6, 0x0d, 0x6f, 0xa4, 0x0d, 0xbf, 0xf1,
	0xd8, 0x82, 0x56, 0xf9, 0xff, 0x2e, 0xfd, 0x71, 0xbc, 0x2e, 0xae, 0x06, 0xfa, 0xc5, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x84, 0x83, 0x13, 0xa3, 0x3c, 0x04, 0x00, 0x00,
}
