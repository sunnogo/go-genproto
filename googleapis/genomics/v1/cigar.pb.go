// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/genomics/v1/cigar.proto
// DO NOT EDIT!

package google_genomics_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Describes the different types of CIGAR alignment operations that exist.
// Used wherever CIGAR alignments are used.
type CigarUnit_Operation int32

const (
	CigarUnit_OPERATION_UNSPECIFIED CigarUnit_Operation = 0
	// An alignment match indicates that a sequence can be aligned to the
	// reference without evidence of an INDEL. Unlike the
	// `SEQUENCE_MATCH` and `SEQUENCE_MISMATCH` operators,
	// the `ALIGNMENT_MATCH` operator does not indicate whether the
	// reference and read sequences are an exact match. This operator is
	// equivalent to SAM's `M`.
	CigarUnit_ALIGNMENT_MATCH CigarUnit_Operation = 1
	// The insert operator indicates that the read contains evidence of bases
	// being inserted into the reference. This operator is equivalent to SAM's
	// `I`.
	CigarUnit_INSERT CigarUnit_Operation = 2
	// The delete operator indicates that the read contains evidence of bases
	// being deleted from the reference. This operator is equivalent to SAM's
	// `D`.
	CigarUnit_DELETE CigarUnit_Operation = 3
	// The skip operator indicates that this read skips a long segment of the
	// reference, but the bases have not been deleted. This operator is commonly
	// used when working with RNA-seq data, where reads may skip long segments
	// of the reference between exons. This operator is equivalent to SAM's
	// `N`.
	CigarUnit_SKIP CigarUnit_Operation = 4
	// The soft clip operator indicates that bases at the start/end of a read
	// have not been considered during alignment. This may occur if the majority
	// of a read maps, except for low quality bases at the start/end of a read.
	// This operator is equivalent to SAM's `S`. Bases that are soft
	// clipped will still be stored in the read.
	CigarUnit_CLIP_SOFT CigarUnit_Operation = 5
	// The hard clip operator indicates that bases at the start/end of a read
	// have been omitted from this alignment. This may occur if this linear
	// alignment is part of a chimeric alignment, or if the read has been
	// trimmed (for example, during error correction or to trim poly-A tails for
	// RNA-seq). This operator is equivalent to SAM's `H`.
	CigarUnit_CLIP_HARD CigarUnit_Operation = 6
	// The pad operator indicates that there is padding in an alignment. This
	// operator is equivalent to SAM's `P`.
	CigarUnit_PAD CigarUnit_Operation = 7
	// This operator indicates that this portion of the aligned sequence exactly
	// matches the reference. This operator is equivalent to SAM's `=`.
	CigarUnit_SEQUENCE_MATCH CigarUnit_Operation = 8
	// This operator indicates that this portion of the aligned sequence is an
	// alignment match to the reference, but a sequence mismatch. This can
	// indicate a SNP or a read error. This operator is equivalent to SAM's
	// `X`.
	CigarUnit_SEQUENCE_MISMATCH CigarUnit_Operation = 9
)

var CigarUnit_Operation_name = map[int32]string{
	0: "OPERATION_UNSPECIFIED",
	1: "ALIGNMENT_MATCH",
	2: "INSERT",
	3: "DELETE",
	4: "SKIP",
	5: "CLIP_SOFT",
	6: "CLIP_HARD",
	7: "PAD",
	8: "SEQUENCE_MATCH",
	9: "SEQUENCE_MISMATCH",
}
var CigarUnit_Operation_value = map[string]int32{
	"OPERATION_UNSPECIFIED": 0,
	"ALIGNMENT_MATCH":       1,
	"INSERT":                2,
	"DELETE":                3,
	"SKIP":                  4,
	"CLIP_SOFT":             5,
	"CLIP_HARD":             6,
	"PAD":                   7,
	"SEQUENCE_MATCH":        8,
	"SEQUENCE_MISMATCH":     9,
}

func (x CigarUnit_Operation) String() string {
	return proto.EnumName(CigarUnit_Operation_name, int32(x))
}
func (CigarUnit_Operation) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

// A single CIGAR operation.
type CigarUnit struct {
	Operation CigarUnit_Operation `protobuf:"varint,1,opt,name=operation,enum=google.genomics.v1.CigarUnit_Operation" json:"operation,omitempty"`
	// The number of genomic bases that the operation runs for. Required.
	OperationLength int64 `protobuf:"varint,2,opt,name=operation_length,json=operationLength" json:"operation_length,omitempty"`
	// `referenceSequence` is only used at mismatches
	// (`SEQUENCE_MISMATCH`) and deletions (`DELETE`).
	// Filling this field replaces SAM's MD tag. If the relevant information is
	// not available, this field is unset.
	ReferenceSequence string `protobuf:"bytes,3,opt,name=reference_sequence,json=referenceSequence" json:"reference_sequence,omitempty"`
}

func (m *CigarUnit) Reset()                    { *m = CigarUnit{} }
func (m *CigarUnit) String() string            { return proto.CompactTextString(m) }
func (*CigarUnit) ProtoMessage()               {}
func (*CigarUnit) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*CigarUnit)(nil), "google.genomics.v1.CigarUnit")
	proto.RegisterEnum("google.genomics.v1.CigarUnit_Operation", CigarUnit_Operation_name, CigarUnit_Operation_value)
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/genomics/v1/cigar.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x51, 0xc1, 0x92, 0x9a, 0x40,
	0x10, 0x0d, 0x62, 0x54, 0xba, 0x2a, 0x3a, 0x4e, 0xca, 0x94, 0xc9, 0x29, 0xe5, 0x25, 0xf1, 0x10,
	0x28, 0x93, 0x73, 0x0e, 0x08, 0x63, 0x9c, 0x0a, 0x02, 0x01, 0x3c, 0x53, 0x84, 0x1a, 0x27, 0x54,
	0xe9, 0x8c, 0x01, 0xd6, 0xdf, 0xda, 0x3f, 0xda, 0xef, 0xd8, 0xe3, 0x02, 0x2a, 0x1e, 0x76, 0x0f,
	0x7b, 0x99, 0xea, 0x7e, 0xef, 0xf5, 0xeb, 0xa9, 0xd7, 0xf0, 0x93, 0x4b, 0xc9, 0xf7, 0x4c, 0xe7,
	0x72, 0x9f, 0x08, 0xae, 0xcb, 0x9c, 0x1b, 0x9c, 0x89, 0x63, 0x2e, 0x4b, 0x69, 0x9c, 0xa9, 0xe4,
	0x98, 0x15, 0x35, 0x26, 0x0f, 0x59, 0x5a, 0x18, 0xa7, 0x85, 0x91, 0x66, 0x3c, 0xc9, 0xf5, 0x46,
	0x82, 0xf1, 0x75, 0xfc, 0xc2, 0xeb, 0xa7, 0xc5, 0x27, 0xfa, 0x3a, 0xcb, 0xea, 0x31, 0x0a, 0x96,
	0x9f, 0xb2, 0x94, 0xa5, 0x52, 0xec, 0x32, 0x6e, 0x24, 0x42, 0xc8, 0x32, 0x29, 0x33, 0x29, 0x8a,
	0xb3, 0xfd, 0xec, 0xa1, 0x03, 0x9a, 0x55, 0xaf, 0xdb, 0x8a, 0xac, 0xc4, 0x04, 0x34, 0x79, 0x64,
	0x79, 0xa3, 0x98, 0x2a, 0x9f, 0x95, 0xaf, 0xc3, 0xef, 0x5f, 0xf4, 0xe7, 0x1f, 0xd0, 0xdb, 0x09,
	0xdd, 0xbb, 0xca, 0x83, 0xdb, 0x24, 0x9e, 0x03, 0x6a, 0x9b, 0x78, 0xcf, 0x04, 0x2f, 0xff, 0x4d,
	0x3b, 0x95, 0x9b, 0x1a, 0x8c, 0x5a, 0xdc, 0x69, 0x60, 0xfc, 0x0d, 0x70, 0xce, 0x76, 0x2c, 0x67,
	0x22, 0x65, 0x71, 0xc1, 0xfe, 0xdf, 0xd5, 0xc5, 0x54, 0xad, 0xc4, 0x5a, 0x30, 0x6e, 0x99, 0xf0,
	0x42, 0xcc, 0xee, 0x15, 0xd0, 0xda, 0x95, 0xf8, 0x23, 0x4c, 0x3c, 0x9f, 0x04, 0x66, 0x44, 0x3d,
	0x37, 0xde, 0xba, 0xa1, 0x4f, 0x2c, 0xba, 0xa2, 0xc4, 0x46, 0x6f, 0xf0, 0x7b, 0x18, 0x99, 0x0e,
	0xfd, 0xe5, 0x6e, 0x88, 0x1b, 0xc5, 0x1b, 0x33, 0xb2, 0xd6, 0x48, 0xc1, 0x00, 0x3d, 0xea, 0x86,
	0x24, 0x88, 0x50, 0xa7, 0xae, 0x6d, 0xe2, 0x90, 0x88, 0x20, 0x15, 0x0f, 0xa0, 0x1b, 0xfe, 0xa6,
	0x3e, 0xea, 0xe2, 0x77, 0x55, 0x1a, 0x0e, 0xf5, 0xe3, 0xd0, 0x5b, 0x45, 0xe8, 0x6d, 0xdb, 0xae,
	0xcd, 0xc0, 0x46, 0x3d, 0xdc, 0x07, 0xd5, 0x37, 0x6d, 0xd4, 0xc7, 0x18, 0x86, 0x21, 0xf9, 0xb3,
	0x25, 0xae, 0x45, 0x2e, 0xe6, 0x03, 0x3c, 0x81, 0xf1, 0x0d, 0xa3, 0xe1, 0x19, 0xd6, 0x96, 0x73,
	0xf8, 0x90, 0xca, 0xc3, 0x0b, 0x21, 0x2e, 0xa1, 0x49, 0xd1, 0xaf, 0xcf, 0xe0, 0x2b, 0x8f, 0x8a,
	0xf2, 0xb7, 0xd7, 0x9c, 0xe4, 0xc7, 0x53, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x86, 0x00, 0x39,
	0x32, 0x02, 0x00, 0x00,
}