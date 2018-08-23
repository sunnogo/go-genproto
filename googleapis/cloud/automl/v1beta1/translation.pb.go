// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/translation.proto

package automl // import "github.com/sunnogo/go-genproto/googleapis/cloud/automl/v1beta1"

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

// Dataset metadata that is specific to translation.
type TranslationDatasetMetadata struct {
	// Required. The BCP-47 language code of the source language.
	SourceLanguageCode string `protobuf:"bytes,1,opt,name=source_language_code,json=sourceLanguageCode,proto3" json:"source_language_code,omitempty"`
	// Required. The BCP-47 language code of the target language.
	TargetLanguageCode   string   `protobuf:"bytes,2,opt,name=target_language_code,json=targetLanguageCode,proto3" json:"target_language_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranslationDatasetMetadata) Reset()         { *m = TranslationDatasetMetadata{} }
func (m *TranslationDatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*TranslationDatasetMetadata) ProtoMessage()    {}
func (*TranslationDatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_translation_23667d09236b8210, []int{0}
}
func (m *TranslationDatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslationDatasetMetadata.Unmarshal(m, b)
}
func (m *TranslationDatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslationDatasetMetadata.Marshal(b, m, deterministic)
}
func (dst *TranslationDatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslationDatasetMetadata.Merge(dst, src)
}
func (m *TranslationDatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_TranslationDatasetMetadata.Size(m)
}
func (m *TranslationDatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslationDatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TranslationDatasetMetadata proto.InternalMessageInfo

func (m *TranslationDatasetMetadata) GetSourceLanguageCode() string {
	if m != nil {
		return m.SourceLanguageCode
	}
	return ""
}

func (m *TranslationDatasetMetadata) GetTargetLanguageCode() string {
	if m != nil {
		return m.TargetLanguageCode
	}
	return ""
}

// Evaluation metrics for the dataset.
type TranslationEvaluationMetrics struct {
	// Output only. BLEU score.
	BleuScore float64 `protobuf:"fixed64,1,opt,name=bleu_score,json=bleuScore,proto3" json:"bleu_score,omitempty"`
	// Output only. BLEU score for base model.
	BaseBleuScore        float64  `protobuf:"fixed64,2,opt,name=base_bleu_score,json=baseBleuScore,proto3" json:"base_bleu_score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranslationEvaluationMetrics) Reset()         { *m = TranslationEvaluationMetrics{} }
func (m *TranslationEvaluationMetrics) String() string { return proto.CompactTextString(m) }
func (*TranslationEvaluationMetrics) ProtoMessage()    {}
func (*TranslationEvaluationMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_translation_23667d09236b8210, []int{1}
}
func (m *TranslationEvaluationMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslationEvaluationMetrics.Unmarshal(m, b)
}
func (m *TranslationEvaluationMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslationEvaluationMetrics.Marshal(b, m, deterministic)
}
func (dst *TranslationEvaluationMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslationEvaluationMetrics.Merge(dst, src)
}
func (m *TranslationEvaluationMetrics) XXX_Size() int {
	return xxx_messageInfo_TranslationEvaluationMetrics.Size(m)
}
func (m *TranslationEvaluationMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslationEvaluationMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_TranslationEvaluationMetrics proto.InternalMessageInfo

func (m *TranslationEvaluationMetrics) GetBleuScore() float64 {
	if m != nil {
		return m.BleuScore
	}
	return 0
}

func (m *TranslationEvaluationMetrics) GetBaseBleuScore() float64 {
	if m != nil {
		return m.BaseBleuScore
	}
	return 0
}

// Model metadata that is specific to translation.
type TranslationModelMetadata struct {
	// The resource name of the model to use as a baseline to train the custom
	// model. If unset, we use the default base model provided by Google
	// Translate. Format:
	// `projects/{project_id}/locations/{location_id}/models/{model_id}`
	BaseModel string `protobuf:"bytes,1,opt,name=base_model,json=baseModel,proto3" json:"base_model,omitempty"`
	// Output only. Inferred from the dataset.
	// The source languge (The BCP-47 language code) that is used for training.
	SourceLanguageCode string `protobuf:"bytes,2,opt,name=source_language_code,json=sourceLanguageCode,proto3" json:"source_language_code,omitempty"`
	// Output only. The target languge (The BCP-47 language code) that is used for
	// training.
	TargetLanguageCode   string   `protobuf:"bytes,3,opt,name=target_language_code,json=targetLanguageCode,proto3" json:"target_language_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranslationModelMetadata) Reset()         { *m = TranslationModelMetadata{} }
func (m *TranslationModelMetadata) String() string { return proto.CompactTextString(m) }
func (*TranslationModelMetadata) ProtoMessage()    {}
func (*TranslationModelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_translation_23667d09236b8210, []int{2}
}
func (m *TranslationModelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslationModelMetadata.Unmarshal(m, b)
}
func (m *TranslationModelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslationModelMetadata.Marshal(b, m, deterministic)
}
func (dst *TranslationModelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslationModelMetadata.Merge(dst, src)
}
func (m *TranslationModelMetadata) XXX_Size() int {
	return xxx_messageInfo_TranslationModelMetadata.Size(m)
}
func (m *TranslationModelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslationModelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TranslationModelMetadata proto.InternalMessageInfo

func (m *TranslationModelMetadata) GetBaseModel() string {
	if m != nil {
		return m.BaseModel
	}
	return ""
}

func (m *TranslationModelMetadata) GetSourceLanguageCode() string {
	if m != nil {
		return m.SourceLanguageCode
	}
	return ""
}

func (m *TranslationModelMetadata) GetTargetLanguageCode() string {
	if m != nil {
		return m.TargetLanguageCode
	}
	return ""
}

// Annotation details specific to translation.
type TranslationAnnotation struct {
	// Output only . The translated content.
	TranslatedContent    *TextSnippet `protobuf:"bytes,1,opt,name=translated_content,json=translatedContent,proto3" json:"translated_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TranslationAnnotation) Reset()         { *m = TranslationAnnotation{} }
func (m *TranslationAnnotation) String() string { return proto.CompactTextString(m) }
func (*TranslationAnnotation) ProtoMessage()    {}
func (*TranslationAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_translation_23667d09236b8210, []int{3}
}
func (m *TranslationAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslationAnnotation.Unmarshal(m, b)
}
func (m *TranslationAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslationAnnotation.Marshal(b, m, deterministic)
}
func (dst *TranslationAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslationAnnotation.Merge(dst, src)
}
func (m *TranslationAnnotation) XXX_Size() int {
	return xxx_messageInfo_TranslationAnnotation.Size(m)
}
func (m *TranslationAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslationAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_TranslationAnnotation proto.InternalMessageInfo

func (m *TranslationAnnotation) GetTranslatedContent() *TextSnippet {
	if m != nil {
		return m.TranslatedContent
	}
	return nil
}

func init() {
	proto.RegisterType((*TranslationDatasetMetadata)(nil), "google.cloud.automl.v1beta1.TranslationDatasetMetadata")
	proto.RegisterType((*TranslationEvaluationMetrics)(nil), "google.cloud.automl.v1beta1.TranslationEvaluationMetrics")
	proto.RegisterType((*TranslationModelMetadata)(nil), "google.cloud.automl.v1beta1.TranslationModelMetadata")
	proto.RegisterType((*TranslationAnnotation)(nil), "google.cloud.automl.v1beta1.TranslationAnnotation")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/translation.proto", fileDescriptor_translation_23667d09236b8210)
}

var fileDescriptor_translation_23667d09236b8210 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0x91, 0x0b, 0x05, 0x6f, 0x29, 0x6d, 0x45, 0x0b, 0xc6, 0x75, 0x69, 0xf1, 0xa1, 0xf8,
	0xd0, 0x4a, 0x75, 0x73, 0xcc, 0xc9, 0x76, 0x72, 0x8b, 0xc1, 0xd8, 0x86, 0x40, 0x2e, 0x62, 0x24,
	0x0d, 0x8b, 0x60, 0xb5, 0x23, 0xb4, 0x23, 0xe3, 0x63, 0x5e, 0x24, 0xef, 0x1a, 0x76, 0x57, 0xb6,
	0x45, 0x88, 0x0d, 0xb9, 0x49, 0x33, 0xbf, 0x6f, 0xfe, 0xed, 0x27, 0xfe, 0x4a, 0x22, 0xa9, 0x30,
	0xce, 0x14, 0x35, 0x79, 0x0c, 0x0d, 0x53, 0xa9, 0xe2, 0xdd, 0x34, 0x45, 0x86, 0x69, 0xcc, 0x35,
	0x68, 0xa3, 0x80, 0x0b, 0xd2, 0x51, 0x55, 0x13, 0x53, 0xf8, 0xdd, 0xe3, 0x91, 0xc3, 0x23, 0x8f,
	0x47, 0x2d, 0x3e, 0x1c, 0xb5, 0xb5, 0xa0, 0x2a, 0x62, 0xd0, 0x9a, 0xd8, 0x29, 0x8d, 0x97, 0x0e,
	0xff, 0x5c, 0xea, 0x94, 0x03, 0x43, 0x52, 0x30, 0x96, 0x2d, 0x3d, 0x7e, 0x0c, 0xc4, 0x70, 0x7b,
	0x6a, 0x7f, 0x03, 0x0c, 0x06, 0x79, 0x89, 0x0c, 0x16, 0x0d, 0xff, 0x89, 0xaf, 0x86, 0x9a, 0x3a,
	0xc3, 0x44, 0x81, 0x96, 0x0d, 0x48, 0x4c, 0x32, 0xca, 0x71, 0x10, 0xfc, 0x0a, 0x26, 0xfd, 0x75,
	0xe8, 0x73, 0x77, 0x6d, 0x6a, 0x41, 0x39, 0x5a, 0x05, 0x43, 0x2d, 0x91, 0x5f, 0x28, 0x7a, 0x5e,
	0xe1, 0x73, 0x5d, 0xc5, 0x18, 0xc5, 0xa8, 0x33, 0xc1, 0xed, 0x0e, 0x54, 0xe3, 0xbe, 0x96, 0xc8,
	0x75, 0x91, 0x99, 0xf0, 0x87, 0x10, 0xa9, 0xc2, 0x26, 0x31, 0x19, 0xd5, 0xbe, 0x73, 0xb0, 0xee,
	0xdb, 0xc8, 0xc6, 0x06, 0xc2, 0xdf, 0xe2, 0x53, 0x0a, 0x06, 0x93, 0x0e, 0xd3, 0x73, 0xcc, 0x47,
	0x1b, 0x9e, 0x1f, 0xb8, 0xf1, 0x53, 0x20, 0x06, 0x9d, 0x3e, 0x4b, 0xca, 0x51, 0x1d, 0xf7, 0xb4,
	0x3d, 0x6c, 0x91, 0xd2, 0x46, 0xdb, 0xed, 0xfa, 0x36, 0xe2, 0xb0, 0xb3, 0x67, 0xe8, 0xbd, 0xf9,
	0x0c, 0xef, 0xce, 0x9e, 0xa1, 0x12, 0xdf, 0x3a, 0xe3, 0xcd, 0x8e, 0xef, 0x1a, 0xde, 0x8b, 0xf0,
	0x60, 0x10, 0xcc, 0x93, 0x8c, 0x34, 0xa3, 0x66, 0x37, 0xe3, 0x87, 0xff, 0x93, 0xe8, 0x82, 0x51,
	0xa2, 0x2d, 0xee, 0x79, 0xa3, 0x8b, 0xaa, 0x42, 0x5e, 0x7f, 0x39, 0xd5, 0x58, 0xf8, 0x12, 0xf3,
	0xbd, 0xf8, 0x99, 0x51, 0x79, 0xa9, 0xc2, 0xfc, 0x73, 0x67, 0xa4, 0x95, 0x35, 0xcc, 0x2a, 0x78,
	0x98, 0xb5, 0x02, 0x49, 0x76, 0xb5, 0x88, 0x6a, 0x19, 0x4b, 0xd4, 0xce, 0x4e, 0xb1, 0x4f, 0x41,
	0x55, 0x98, 0x57, 0xfd, 0x77, 0xed, 0x7f, 0xd3, 0xf7, 0x8e, 0xbe, 0x7a, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0xfb, 0xc7, 0xb3, 0x37, 0x16, 0x03, 0x00, 0x00,
}
