// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/resultstore/v2/invocation.proto

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

// An Invocation typically represents the result of running a tool. Each has a
// unique ID, typically generated by the server. Target resources under each
// Invocation contain the bulk of the data.
type Invocation struct {
	// The resource name.  Its format must be:
	// invocations/${INVOCATION_ID}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resource ID components that identify the Invocation. They must match
	// the resource name after proper encoding.
	Id *Invocation_Id `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The aggregate status of the invocation.
	StatusAttributes *StatusAttributes `protobuf:"bytes,3,opt,name=status_attributes,json=statusAttributes,proto3" json:"status_attributes,omitempty"`
	// When this invocation started and its duration.
	Timing *Timing `protobuf:"bytes,4,opt,name=timing,proto3" json:"timing,omitempty"`
	// Attributes of this invocation.
	InvocationAttributes *InvocationAttributes `protobuf:"bytes,5,opt,name=invocation_attributes,json=invocationAttributes,proto3" json:"invocation_attributes,omitempty"`
	// The workspace the tool was run in.
	WorkspaceInfo *WorkspaceInfo `protobuf:"bytes,6,opt,name=workspace_info,json=workspaceInfo,proto3" json:"workspace_info,omitempty"`
	// Arbitrary name-value pairs.
	// This is implemented as a multi-map. Multiple properties are allowed with
	// the same key. Properties will be returned in lexicographical order by key.
	Properties []*Property `protobuf:"bytes,7,rep,name=properties,proto3" json:"properties,omitempty"`
	// A list of file references for invocation level files.
	// The file IDs must be unique within this list. Duplicate file IDs will
	// result in an error. Files will be returned in lexicographical order by ID.
	// Use this field to specify build logs, and other invocation level logs.
	Files []*File `protobuf:"bytes,8,rep,name=files,proto3" json:"files,omitempty"`
	// Summary of aggregate coverage across all Actions in this Invocation.
	// the server populates this for you in the post-processing phase.
	CoverageSummaries    []*LanguageCoverageSummary `protobuf:"bytes,9,rep,name=coverage_summaries,json=coverageSummaries,proto3" json:"coverage_summaries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Invocation) Reset()         { *m = Invocation{} }
func (m *Invocation) String() string { return proto.CompactTextString(m) }
func (*Invocation) ProtoMessage()    {}
func (*Invocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_invocation_04b66137ff87776e, []int{0}
}
func (m *Invocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invocation.Unmarshal(m, b)
}
func (m *Invocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invocation.Marshal(b, m, deterministic)
}
func (dst *Invocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invocation.Merge(dst, src)
}
func (m *Invocation) XXX_Size() int {
	return xxx_messageInfo_Invocation.Size(m)
}
func (m *Invocation) XXX_DiscardUnknown() {
	xxx_messageInfo_Invocation.DiscardUnknown(m)
}

var xxx_messageInfo_Invocation proto.InternalMessageInfo

func (m *Invocation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Invocation) GetId() *Invocation_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Invocation) GetStatusAttributes() *StatusAttributes {
	if m != nil {
		return m.StatusAttributes
	}
	return nil
}

func (m *Invocation) GetTiming() *Timing {
	if m != nil {
		return m.Timing
	}
	return nil
}

func (m *Invocation) GetInvocationAttributes() *InvocationAttributes {
	if m != nil {
		return m.InvocationAttributes
	}
	return nil
}

func (m *Invocation) GetWorkspaceInfo() *WorkspaceInfo {
	if m != nil {
		return m.WorkspaceInfo
	}
	return nil
}

func (m *Invocation) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Invocation) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *Invocation) GetCoverageSummaries() []*LanguageCoverageSummary {
	if m != nil {
		return m.CoverageSummaries
	}
	return nil
}

// The resource ID components that identify the Invocation.
type Invocation_Id struct {
	// The Invocation ID.
	InvocationId         string   `protobuf:"bytes,1,opt,name=invocation_id,json=invocationId,proto3" json:"invocation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Invocation_Id) Reset()         { *m = Invocation_Id{} }
func (m *Invocation_Id) String() string { return proto.CompactTextString(m) }
func (*Invocation_Id) ProtoMessage()    {}
func (*Invocation_Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_invocation_04b66137ff87776e, []int{0, 0}
}
func (m *Invocation_Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invocation_Id.Unmarshal(m, b)
}
func (m *Invocation_Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invocation_Id.Marshal(b, m, deterministic)
}
func (dst *Invocation_Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invocation_Id.Merge(dst, src)
}
func (m *Invocation_Id) XXX_Size() int {
	return xxx_messageInfo_Invocation_Id.Size(m)
}
func (m *Invocation_Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Invocation_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Invocation_Id proto.InternalMessageInfo

func (m *Invocation_Id) GetInvocationId() string {
	if m != nil {
		return m.InvocationId
	}
	return ""
}

// If known, represents the state of the user/build-system workspace.
type WorkspaceContext struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkspaceContext) Reset()         { *m = WorkspaceContext{} }
func (m *WorkspaceContext) String() string { return proto.CompactTextString(m) }
func (*WorkspaceContext) ProtoMessage()    {}
func (*WorkspaceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_invocation_04b66137ff87776e, []int{1}
}
func (m *WorkspaceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkspaceContext.Unmarshal(m, b)
}
func (m *WorkspaceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkspaceContext.Marshal(b, m, deterministic)
}
func (dst *WorkspaceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkspaceContext.Merge(dst, src)
}
func (m *WorkspaceContext) XXX_Size() int {
	return xxx_messageInfo_WorkspaceContext.Size(m)
}
func (m *WorkspaceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkspaceContext.DiscardUnknown(m)
}

var xxx_messageInfo_WorkspaceContext proto.InternalMessageInfo

// Describes the workspace under which the tool was invoked, this includes
// information that was fed into the command, the source code referenced, and
// the tool itself.
type WorkspaceInfo struct {
	// Data about the workspace that might be useful for debugging.
	WorkspaceContext *WorkspaceContext `protobuf:"bytes,1,opt,name=workspace_context,json=workspaceContext,proto3" json:"workspace_context,omitempty"`
	// Where the tool was invoked
	Hostname string `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// The client's working directory where the build/test was run from.
	WorkingDirectory string `protobuf:"bytes,4,opt,name=working_directory,json=workingDirectory,proto3" json:"working_directory,omitempty"`
	// Tools should set tool_tag to the name of the tool or use case.
	ToolTag string `protobuf:"bytes,5,opt,name=tool_tag,json=toolTag,proto3" json:"tool_tag,omitempty"`
	// The command lines invoked. The first command line is the one typed by the
	// user, then each one after that should be an expansion of the previous
	// command line.
	CommandLines         []*CommandLine `protobuf:"bytes,7,rep,name=command_lines,json=commandLines,proto3" json:"command_lines,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *WorkspaceInfo) Reset()         { *m = WorkspaceInfo{} }
func (m *WorkspaceInfo) String() string { return proto.CompactTextString(m) }
func (*WorkspaceInfo) ProtoMessage()    {}
func (*WorkspaceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_invocation_04b66137ff87776e, []int{2}
}
func (m *WorkspaceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkspaceInfo.Unmarshal(m, b)
}
func (m *WorkspaceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkspaceInfo.Marshal(b, m, deterministic)
}
func (dst *WorkspaceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkspaceInfo.Merge(dst, src)
}
func (m *WorkspaceInfo) XXX_Size() int {
	return xxx_messageInfo_WorkspaceInfo.Size(m)
}
func (m *WorkspaceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkspaceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_WorkspaceInfo proto.InternalMessageInfo

func (m *WorkspaceInfo) GetWorkspaceContext() *WorkspaceContext {
	if m != nil {
		return m.WorkspaceContext
	}
	return nil
}

func (m *WorkspaceInfo) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *WorkspaceInfo) GetWorkingDirectory() string {
	if m != nil {
		return m.WorkingDirectory
	}
	return ""
}

func (m *WorkspaceInfo) GetToolTag() string {
	if m != nil {
		return m.ToolTag
	}
	return ""
}

func (m *WorkspaceInfo) GetCommandLines() []*CommandLine {
	if m != nil {
		return m.CommandLines
	}
	return nil
}

// The command and arguments that produced this Invocation.
type CommandLine struct {
	// A label describing this command line.
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// The command-line tool that is run: argv[0].
	Tool string `protobuf:"bytes,2,opt,name=tool,proto3" json:"tool,omitempty"`
	// The arguments to the above tool: argv[1]...argv[N].
	Args []string `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	// The actual command that was run with the tool.  (e.g. "build", or "test")
	// Omit if the tool doesn't accept a command.
	// This is a duplicate of one of the fields in args.
	Command              string   `protobuf:"bytes,4,opt,name=command,proto3" json:"command,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandLine) Reset()         { *m = CommandLine{} }
func (m *CommandLine) String() string { return proto.CompactTextString(m) }
func (*CommandLine) ProtoMessage()    {}
func (*CommandLine) Descriptor() ([]byte, []int) {
	return fileDescriptor_invocation_04b66137ff87776e, []int{3}
}
func (m *CommandLine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandLine.Unmarshal(m, b)
}
func (m *CommandLine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandLine.Marshal(b, m, deterministic)
}
func (dst *CommandLine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandLine.Merge(dst, src)
}
func (m *CommandLine) XXX_Size() int {
	return xxx_messageInfo_CommandLine.Size(m)
}
func (m *CommandLine) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandLine.DiscardUnknown(m)
}

var xxx_messageInfo_CommandLine proto.InternalMessageInfo

func (m *CommandLine) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *CommandLine) GetTool() string {
	if m != nil {
		return m.Tool
	}
	return ""
}

func (m *CommandLine) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *CommandLine) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

// Attributes that apply to all invocations.
type InvocationAttributes struct {
	// The project ID this invocation is associated with. This must be
	// set in the CreateInvocation call, and can't be changed.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The list of users in the command chain.  The first user in this sequence
	// is the one who instigated the first command in the chain.
	Users []string `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	// Labels to categorize this invocation.
	// This is implemented as a set. All labels will be unique. Any duplicate
	// labels added will be ignored. Labels will be returned in lexicographical
	// order. Labels should be short, easy to read, and you
	// shouldn't have more than a handful.
	// Labels should match regex \w([- \w]*\w)?
	// Labels should not be used for unique properties such as unique IDs.
	// Use properties in cases that don't meet these conditions.
	Labels               []string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvocationAttributes) Reset()         { *m = InvocationAttributes{} }
func (m *InvocationAttributes) String() string { return proto.CompactTextString(m) }
func (*InvocationAttributes) ProtoMessage()    {}
func (*InvocationAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_invocation_04b66137ff87776e, []int{4}
}
func (m *InvocationAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvocationAttributes.Unmarshal(m, b)
}
func (m *InvocationAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvocationAttributes.Marshal(b, m, deterministic)
}
func (dst *InvocationAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvocationAttributes.Merge(dst, src)
}
func (m *InvocationAttributes) XXX_Size() int {
	return xxx_messageInfo_InvocationAttributes.Size(m)
}
func (m *InvocationAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_InvocationAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_InvocationAttributes proto.InternalMessageInfo

func (m *InvocationAttributes) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *InvocationAttributes) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *InvocationAttributes) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*Invocation)(nil), "google.devtools.resultstore.v2.Invocation")
	proto.RegisterType((*Invocation_Id)(nil), "google.devtools.resultstore.v2.Invocation.Id")
	proto.RegisterType((*WorkspaceContext)(nil), "google.devtools.resultstore.v2.WorkspaceContext")
	proto.RegisterType((*WorkspaceInfo)(nil), "google.devtools.resultstore.v2.WorkspaceInfo")
	proto.RegisterType((*CommandLine)(nil), "google.devtools.resultstore.v2.CommandLine")
	proto.RegisterType((*InvocationAttributes)(nil), "google.devtools.resultstore.v2.InvocationAttributes")
}

func init() {
	proto.RegisterFile("google/devtools/resultstore/v2/invocation.proto", fileDescriptor_invocation_04b66137ff87776e)
}

var fileDescriptor_invocation_04b66137ff87776e = []byte{
	// 631 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0x71, 0x6b, 0x13, 0x3f,
	0x18, 0xc7, 0x69, 0xbb, 0x76, 0xeb, 0xb3, 0xf5, 0xc7, 0x16, 0xf6, 0x93, 0xb3, 0xa0, 0x8c, 0x2a,
	0xd2, 0x31, 0xbc, 0x93, 0xaa, 0x08, 0x8a, 0x82, 0x4e, 0xc4, 0x83, 0xfd, 0x31, 0xb2, 0x81, 0x20,
	0xc8, 0x91, 0xdd, 0xa5, 0x31, 0x7a, 0x97, 0xd4, 0x24, 0xd7, 0xb9, 0xf7, 0xe3, 0x8b, 0xf2, 0xe5,
	0x48, 0x72, 0x77, 0xbd, 0x5b, 0x99, 0xde, 0xfe, 0xbb, 0xe7, 0x69, 0xbe, 0x9f, 0x3c, 0x79, 0xf2,
	0xcd, 0x53, 0x08, 0x98, 0x94, 0x2c, 0xa5, 0x41, 0x42, 0x97, 0x46, 0xca, 0x54, 0x07, 0x8a, 0xea,
	0x3c, 0x35, 0xda, 0x48, 0x45, 0x83, 0xe5, 0x2c, 0xe0, 0x62, 0x29, 0x63, 0x62, 0xb8, 0x14, 0xfe,
	0x42, 0x49, 0x23, 0xd1, 0xfd, 0x42, 0xe0, 0x57, 0x02, 0xbf, 0x21, 0xf0, 0x97, 0xb3, 0xf1, 0x51,
	0x0b, 0x30, 0x96, 0x59, 0x56, 0xc1, 0xc6, 0xcf, 0x5b, 0x17, 0x2f, 0xa9, 0x22, 0x8c, 0x46, 0x3a,
	0xcf, 0x32, 0xa2, 0xae, 0x4a, 0xd9, 0x61, 0x8b, 0x6c, 0xce, 0x53, 0x5a, 0x2c, 0x9d, 0xfc, 0xee,
	0x03, 0x84, 0xab, 0x33, 0x20, 0x04, 0x1b, 0x82, 0x64, 0xd4, 0xeb, 0x1c, 0x74, 0xa6, 0x43, 0xec,
	0xbe, 0xd1, 0x6b, 0xe8, 0xf2, 0xc4, 0xeb, 0x1e, 0x74, 0xa6, 0xdb, 0xb3, 0xc7, 0xfe, 0xbf, 0x8f,
	0xe7, 0xd7, 0x2c, 0x3f, 0x4c, 0x70, 0x97, 0x27, 0xe8, 0x0b, 0xec, 0x69, 0x43, 0x4c, 0xae, 0x23,
	0x62, 0x8c, 0xe2, 0x17, 0xb9, 0xa1, 0xda, 0xeb, 0x39, 0xda, 0x93, 0x36, 0xda, 0x99, 0x13, 0xbe,
	0x5d, 0xe9, 0xf0, 0xae, 0x5e, 0xcb, 0xa0, 0x37, 0x30, 0x30, 0x3c, 0xe3, 0x82, 0x79, 0x1b, 0x8e,
	0xf9, 0xa8, 0x8d, 0x79, 0xee, 0x56, 0xe3, 0x52, 0x85, 0x38, 0xfc, 0x5f, 0xdf, 0x61, 0xb3, 0xc4,
	0xbe, 0xc3, 0x3d, 0xbb, 0xfd, 0x81, 0x1b, 0x65, 0xee, 0xf3, 0x1b, 0xb2, 0xe8, 0x1c, 0xfe, 0xbb,
	0x94, 0xea, 0xbb, 0x5e, 0x90, 0x98, 0x46, 0x5c, 0xcc, 0xa5, 0x37, 0xb8, 0x5d, 0x53, 0x3f, 0x55,
	0xaa, 0x50, 0xcc, 0x25, 0x1e, 0x5d, 0x36, 0x43, 0xf4, 0x11, 0x60, 0xa1, 0xe4, 0x82, 0x2a, 0xc3,
	0xa9, 0xf6, 0x36, 0x0f, 0x7a, 0xd3, 0xed, 0xd9, 0xb4, 0x8d, 0x78, 0x5a, 0x28, 0xae, 0x70, 0x43,
	0x8b, 0x5e, 0x42, 0xdf, 0x3a, 0x43, 0x7b, 0x5b, 0x0e, 0xf2, 0xb0, 0x0d, 0xf2, 0x81, 0xa7, 0x14,
	0x17, 0x12, 0x34, 0x07, 0xb4, 0x66, 0x46, 0x5b, 0xcd, 0xd0, 0x81, 0x5e, 0xb4, 0x81, 0x4e, 0x88,
	0x60, 0x39, 0x61, 0xf4, 0xb8, 0x24, 0x9c, 0x15, 0x6e, 0xc6, 0x7b, 0xf1, 0xb5, 0x04, 0xa7, 0x7a,
	0x7c, 0x08, 0xdd, 0x30, 0x41, 0x0f, 0x60, 0xd4, 0xb8, 0x34, 0x9e, 0x94, 0x7e, 0xdd, 0xa9, 0x93,
	0x61, 0x32, 0x41, 0xb0, 0xbb, 0x6a, 0xdc, 0xb1, 0x14, 0x86, 0xfe, 0x34, 0x93, 0x5f, 0x5d, 0x18,
	0x5d, 0xeb, 0xa6, 0xb5, 0x67, 0x7d, 0x29, 0x71, 0xb1, 0xcc, 0xe1, 0x6e, 0x61, 0xcf, 0x75, 0x3c,
	0xde, 0xbd, 0x5c, 0xcb, 0xa0, 0x31, 0x6c, 0x7d, 0x95, 0xda, 0xb8, 0x47, 0xd5, 0x73, 0x45, 0xae,
	0x62, 0x74, 0x54, 0x6c, 0xcd, 0x05, 0x8b, 0x12, 0xae, 0x68, 0x6c, 0xa4, 0xba, 0x72, 0x2e, 0x1e,
	0x16, 0x20, 0x2e, 0xd8, 0xfb, 0x2a, 0x8f, 0xee, 0xc2, 0x96, 0xad, 0x21, 0x32, 0x84, 0x39, 0x6b,
	0x0e, 0xf1, 0xa6, 0x8d, 0xcf, 0x09, 0x43, 0xa7, 0x30, 0xb2, 0x53, 0x83, 0x88, 0x24, 0x4a, 0xb9,
	0x58, 0x99, 0xe0, 0xa8, 0xad, 0xfc, 0xe3, 0x42, 0x74, 0xc2, 0x05, 0xc5, 0x3b, 0x71, 0x1d, 0xe8,
	0x09, 0x85, 0xed, 0xc6, 0x8f, 0x68, 0x1f, 0xfa, 0x29, 0xb9, 0xa0, 0x69, 0xd9, 0xe6, 0x22, 0xb0,
	0xb3, 0xc2, 0x62, 0xdd, 0x64, 0x18, 0x62, 0xf7, 0x6d, 0x73, 0x44, 0x31, 0xfb, 0xbe, 0x7b, 0x36,
	0x67, 0xbf, 0x91, 0x07, 0x9b, 0x25, 0xbc, 0x3c, 0x5c, 0x15, 0x4e, 0x62, 0xd8, 0xbf, 0xe9, 0xf9,
	0xa0, 0x7b, 0xce, 0xd2, 0xdf, 0x68, 0x6c, 0xea, 0xbb, 0x1d, 0x96, 0x99, 0x30, 0xb1, 0xe5, 0xe4,
	0x9a, 0x2a, 0xed, 0x75, 0xdd, 0x2e, 0x45, 0x80, 0xee, 0xc0, 0xc0, 0xd5, 0x55, 0x6d, 0x5e, 0x46,
	0xef, 0x7e, 0xc0, 0x24, 0x96, 0x59, 0x4b, 0x2f, 0x4e, 0x3b, 0x9f, 0xc3, 0x72, 0x05, 0x93, 0x29,
	0x11, 0xcc, 0x97, 0x8a, 0x05, 0x8c, 0x0a, 0x37, 0x25, 0xcb, 0x3f, 0x01, 0xb2, 0xe0, 0xfa, 0x6f,
	0x33, 0xf5, 0x55, 0x23, 0xbc, 0x18, 0x38, 0xd5, 0xd3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfa,
	0x5b, 0x4b, 0x30, 0x3d, 0x06, 0x00, 0x00,
}
