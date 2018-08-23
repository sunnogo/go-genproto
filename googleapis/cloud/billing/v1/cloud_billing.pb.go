// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/billing/v1/cloud_billing.proto

package billing // import "github.com/sunnogo/go-genproto/googleapis/cloud/billing/v1"

import proto "github.com/sunnogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/sunnogo/go-genproto/googleapis/api/annotations"

import (
	context "github.com/sunnogo/net/context"
	grpc "github.com/sunnogo/grpc-go"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A billing account in [Google Cloud
// Console](https://console.cloud.google.com/). You can assign a billing account
// to one or more projects.
type BillingAccount struct {
	// The resource name of the billing account. The resource name has the form
	// `billingAccounts/{billing_account_id}`. For example,
	// `billingAccounts/012345-567890-ABCDEF` would be the resource name for
	// billing account `012345-567890-ABCDEF`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// True if the billing account is open, and will therefore be charged for any
	// usage on associated projects. False if the billing account is closed, and
	// therefore projects associated with it will be unable to use paid services.
	Open bool `protobuf:"varint,2,opt,name=open,proto3" json:"open,omitempty"`
	// The display name given to the billing account, such as `My Billing
	// Account`. This name is displayed in the Google Cloud Console.
	DisplayName          string   `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BillingAccount) Reset()         { *m = BillingAccount{} }
func (m *BillingAccount) String() string { return proto.CompactTextString(m) }
func (*BillingAccount) ProtoMessage()    {}
func (*BillingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_cloud_billing_868724f2532f6f7f, []int{0}
}
func (m *BillingAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillingAccount.Unmarshal(m, b)
}
func (m *BillingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillingAccount.Marshal(b, m, deterministic)
}
func (dst *BillingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillingAccount.Merge(dst, src)
}
func (m *BillingAccount) XXX_Size() int {
	return xxx_messageInfo_BillingAccount.Size(m)
}
func (m *BillingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_BillingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_BillingAccount proto.InternalMessageInfo

func (m *BillingAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BillingAccount) GetOpen() bool {
	if m != nil {
		return m.Open
	}
	return false
}

func (m *BillingAccount) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

// Encapsulation of billing information for a Cloud Console project. A project
// has at most one associated billing account at a time (but a billing account
// can be assigned to multiple projects).
type ProjectBillingInfo struct {
	// The resource name for the `ProjectBillingInfo`; has the form
	// `projects/{project_id}/billingInfo`. For example, the resource name for the
	// billing information for project `tokyo-rain-123` would be
	// `projects/tokyo-rain-123/billingInfo`. This field is read-only.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The ID of the project that this `ProjectBillingInfo` represents, such as
	// `tokyo-rain-123`. This is a convenience field so that you don't need to
	// parse the `name` field to obtain a project ID. This field is read-only.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The resource name of the billing account associated with the project, if
	// any. For example, `billingAccounts/012345-567890-ABCDEF`.
	BillingAccountName string `protobuf:"bytes,3,opt,name=billing_account_name,json=billingAccountName,proto3" json:"billing_account_name,omitempty"`
	// True if the project is associated with an open billing account, to which
	// usage on the project is charged. False if the project is associated with a
	// closed billing account, or no billing account at all, and therefore cannot
	// use paid services. This field is read-only.
	BillingEnabled       bool     `protobuf:"varint,4,opt,name=billing_enabled,json=billingEnabled,proto3" json:"billing_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectBillingInfo) Reset()         { *m = ProjectBillingInfo{} }
func (m *ProjectBillingInfo) String() string { return proto.CompactTextString(m) }
func (*ProjectBillingInfo) ProtoMessage()    {}
func (*ProjectBillingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cloud_billing_868724f2532f6f7f, []int{1}
}
func (m *ProjectBillingInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectBillingInfo.Unmarshal(m, b)
}
func (m *ProjectBillingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectBillingInfo.Marshal(b, m, deterministic)
}
func (dst *ProjectBillingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectBillingInfo.Merge(dst, src)
}
func (m *ProjectBillingInfo) XXX_Size() int {
	return xxx_messageInfo_ProjectBillingInfo.Size(m)
}
func (m *ProjectBillingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectBillingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectBillingInfo proto.InternalMessageInfo

func (m *ProjectBillingInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProjectBillingInfo) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *ProjectBillingInfo) GetBillingAccountName() string {
	if m != nil {
		return m.BillingAccountName
	}
	return ""
}

func (m *ProjectBillingInfo) GetBillingEnabled() bool {
	if m != nil {
		return m.BillingEnabled
	}
	return false
}

// Request message for `GetBillingAccount`.
type GetBillingAccountRequest struct {
	// The resource name of the billing account to retrieve. For example,
	// `billingAccounts/012345-567890-ABCDEF`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBillingAccountRequest) Reset()         { *m = GetBillingAccountRequest{} }
func (m *GetBillingAccountRequest) String() string { return proto.CompactTextString(m) }
func (*GetBillingAccountRequest) ProtoMessage()    {}
func (*GetBillingAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cloud_billing_868724f2532f6f7f, []int{2}
}
func (m *GetBillingAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBillingAccountRequest.Unmarshal(m, b)
}
func (m *GetBillingAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBillingAccountRequest.Marshal(b, m, deterministic)
}
func (dst *GetBillingAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBillingAccountRequest.Merge(dst, src)
}
func (m *GetBillingAccountRequest) XXX_Size() int {
	return xxx_messageInfo_GetBillingAccountRequest.Size(m)
}
func (m *GetBillingAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBillingAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBillingAccountRequest proto.InternalMessageInfo

func (m *GetBillingAccountRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for `ListBillingAccounts`.
type ListBillingAccountsRequest struct {
	// Requested page size. The maximum page size is 100; this is also the
	// default.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A token identifying a page of results to return. This should be a
	// `next_page_token` value returned from a previous `ListBillingAccounts`
	// call. If unspecified, the first page of results is returned.
	PageToken            string   `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBillingAccountsRequest) Reset()         { *m = ListBillingAccountsRequest{} }
func (m *ListBillingAccountsRequest) String() string { return proto.CompactTextString(m) }
func (*ListBillingAccountsRequest) ProtoMessage()    {}
func (*ListBillingAccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cloud_billing_868724f2532f6f7f, []int{3}
}
func (m *ListBillingAccountsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBillingAccountsRequest.Unmarshal(m, b)
}
func (m *ListBillingAccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBillingAccountsRequest.Marshal(b, m, deterministic)
}
func (dst *ListBillingAccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBillingAccountsRequest.Merge(dst, src)
}
func (m *ListBillingAccountsRequest) XXX_Size() int {
	return xxx_messageInfo_ListBillingAccountsRequest.Size(m)
}
func (m *ListBillingAccountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBillingAccountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBillingAccountsRequest proto.InternalMessageInfo

func (m *ListBillingAccountsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListBillingAccountsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Response message for `ListBillingAccounts`.
type ListBillingAccountsResponse struct {
	// A list of billing accounts.
	BillingAccounts []*BillingAccount `protobuf:"bytes,1,rep,name=billing_accounts,json=billingAccounts,proto3" json:"billing_accounts,omitempty"`
	// A token to retrieve the next page of results. To retrieve the next page,
	// call `ListBillingAccounts` again with the `page_token` field set to this
	// value. This field is empty if there are no more results to retrieve.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBillingAccountsResponse) Reset()         { *m = ListBillingAccountsResponse{} }
func (m *ListBillingAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*ListBillingAccountsResponse) ProtoMessage()    {}
func (*ListBillingAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cloud_billing_868724f2532f6f7f, []int{4}
}
func (m *ListBillingAccountsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBillingAccountsResponse.Unmarshal(m, b)
}
func (m *ListBillingAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBillingAccountsResponse.Marshal(b, m, deterministic)
}
func (dst *ListBillingAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBillingAccountsResponse.Merge(dst, src)
}
func (m *ListBillingAccountsResponse) XXX_Size() int {
	return xxx_messageInfo_ListBillingAccountsResponse.Size(m)
}
func (m *ListBillingAccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBillingAccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBillingAccountsResponse proto.InternalMessageInfo

func (m *ListBillingAccountsResponse) GetBillingAccounts() []*BillingAccount {
	if m != nil {
		return m.BillingAccounts
	}
	return nil
}

func (m *ListBillingAccountsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request message for `ListProjectBillingInfo`.
type ListProjectBillingInfoRequest struct {
	// The resource name of the billing account associated with the projects that
	// you want to list. For example, `billingAccounts/012345-567890-ABCDEF`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Requested page size. The maximum page size is 100; this is also the
	// default.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A token identifying a page of results to be returned. This should be a
	// `next_page_token` value returned from a previous `ListProjectBillingInfo`
	// call. If unspecified, the first page of results is returned.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListProjectBillingInfoRequest) Reset()         { *m = ListProjectBillingInfoRequest{} }
func (m *ListProjectBillingInfoRequest) String() string { return proto.CompactTextString(m) }
func (*ListProjectBillingInfoRequest) ProtoMessage()    {}
func (*ListProjectBillingInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cloud_billing_868724f2532f6f7f, []int{5}
}
func (m *ListProjectBillingInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProjectBillingInfoRequest.Unmarshal(m, b)
}
func (m *ListProjectBillingInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProjectBillingInfoRequest.Marshal(b, m, deterministic)
}
func (dst *ListProjectBillingInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProjectBillingInfoRequest.Merge(dst, src)
}
func (m *ListProjectBillingInfoRequest) XXX_Size() int {
	return xxx_messageInfo_ListProjectBillingInfoRequest.Size(m)
}
func (m *ListProjectBillingInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProjectBillingInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListProjectBillingInfoRequest proto.InternalMessageInfo

func (m *ListProjectBillingInfoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListProjectBillingInfoRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListProjectBillingInfoRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Request message for `ListProjectBillingInfoResponse`.
type ListProjectBillingInfoResponse struct {
	// A list of `ProjectBillingInfo` resources representing the projects
	// associated with the billing account.
	ProjectBillingInfo []*ProjectBillingInfo `protobuf:"bytes,1,rep,name=project_billing_info,json=projectBillingInfo,proto3" json:"project_billing_info,omitempty"`
	// A token to retrieve the next page of results. To retrieve the next page,
	// call `ListProjectBillingInfo` again with the `page_token` field set to this
	// value. This field is empty if there are no more results to retrieve.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListProjectBillingInfoResponse) Reset()         { *m = ListProjectBillingInfoResponse{} }
func (m *ListProjectBillingInfoResponse) String() string { return proto.CompactTextString(m) }
func (*ListProjectBillingInfoResponse) ProtoMessage()    {}
func (*ListProjectBillingInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cloud_billing_868724f2532f6f7f, []int{6}
}
func (m *ListProjectBillingInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProjectBillingInfoResponse.Unmarshal(m, b)
}
func (m *ListProjectBillingInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProjectBillingInfoResponse.Marshal(b, m, deterministic)
}
func (dst *ListProjectBillingInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProjectBillingInfoResponse.Merge(dst, src)
}
func (m *ListProjectBillingInfoResponse) XXX_Size() int {
	return xxx_messageInfo_ListProjectBillingInfoResponse.Size(m)
}
func (m *ListProjectBillingInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProjectBillingInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListProjectBillingInfoResponse proto.InternalMessageInfo

func (m *ListProjectBillingInfoResponse) GetProjectBillingInfo() []*ProjectBillingInfo {
	if m != nil {
		return m.ProjectBillingInfo
	}
	return nil
}

func (m *ListProjectBillingInfoResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request message for `GetProjectBillingInfo`.
type GetProjectBillingInfoRequest struct {
	// The resource name of the project for which billing information is
	// retrieved. For example, `projects/tokyo-rain-123`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProjectBillingInfoRequest) Reset()         { *m = GetProjectBillingInfoRequest{} }
func (m *GetProjectBillingInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetProjectBillingInfoRequest) ProtoMessage()    {}
func (*GetProjectBillingInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cloud_billing_868724f2532f6f7f, []int{7}
}
func (m *GetProjectBillingInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProjectBillingInfoRequest.Unmarshal(m, b)
}
func (m *GetProjectBillingInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProjectBillingInfoRequest.Marshal(b, m, deterministic)
}
func (dst *GetProjectBillingInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProjectBillingInfoRequest.Merge(dst, src)
}
func (m *GetProjectBillingInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetProjectBillingInfoRequest.Size(m)
}
func (m *GetProjectBillingInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProjectBillingInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProjectBillingInfoRequest proto.InternalMessageInfo

func (m *GetProjectBillingInfoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for `UpdateProjectBillingInfo`.
type UpdateProjectBillingInfoRequest struct {
	// The resource name of the project associated with the billing information
	// that you want to update. For example, `projects/tokyo-rain-123`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The new billing information for the project. Read-only fields are ignored;
	// thus, you may leave empty all fields except `billing_account_name`.
	ProjectBillingInfo   *ProjectBillingInfo `protobuf:"bytes,2,opt,name=project_billing_info,json=projectBillingInfo,proto3" json:"project_billing_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UpdateProjectBillingInfoRequest) Reset()         { *m = UpdateProjectBillingInfoRequest{} }
func (m *UpdateProjectBillingInfoRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateProjectBillingInfoRequest) ProtoMessage()    {}
func (*UpdateProjectBillingInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cloud_billing_868724f2532f6f7f, []int{8}
}
func (m *UpdateProjectBillingInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProjectBillingInfoRequest.Unmarshal(m, b)
}
func (m *UpdateProjectBillingInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProjectBillingInfoRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateProjectBillingInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProjectBillingInfoRequest.Merge(dst, src)
}
func (m *UpdateProjectBillingInfoRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateProjectBillingInfoRequest.Size(m)
}
func (m *UpdateProjectBillingInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProjectBillingInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProjectBillingInfoRequest proto.InternalMessageInfo

func (m *UpdateProjectBillingInfoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateProjectBillingInfoRequest) GetProjectBillingInfo() *ProjectBillingInfo {
	if m != nil {
		return m.ProjectBillingInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*BillingAccount)(nil), "google.cloud.billing.v1.BillingAccount")
	proto.RegisterType((*ProjectBillingInfo)(nil), "google.cloud.billing.v1.ProjectBillingInfo")
	proto.RegisterType((*GetBillingAccountRequest)(nil), "google.cloud.billing.v1.GetBillingAccountRequest")
	proto.RegisterType((*ListBillingAccountsRequest)(nil), "google.cloud.billing.v1.ListBillingAccountsRequest")
	proto.RegisterType((*ListBillingAccountsResponse)(nil), "google.cloud.billing.v1.ListBillingAccountsResponse")
	proto.RegisterType((*ListProjectBillingInfoRequest)(nil), "google.cloud.billing.v1.ListProjectBillingInfoRequest")
	proto.RegisterType((*ListProjectBillingInfoResponse)(nil), "google.cloud.billing.v1.ListProjectBillingInfoResponse")
	proto.RegisterType((*GetProjectBillingInfoRequest)(nil), "google.cloud.billing.v1.GetProjectBillingInfoRequest")
	proto.RegisterType((*UpdateProjectBillingInfoRequest)(nil), "google.cloud.billing.v1.UpdateProjectBillingInfoRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CloudBillingClient is the client API for CloudBilling service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/github.com/sunnogo/grpc-go#ClientConn.NewStream.
type CloudBillingClient interface {
	// Gets information about a billing account. The current authenticated user
	// must be an [owner of the billing
	// account](https://support.google.com/cloud/answer/4430947).
	GetBillingAccount(ctx context.Context, in *GetBillingAccountRequest, opts ...grpc.CallOption) (*BillingAccount, error)
	// Lists the billing accounts that the current authenticated user
	// [owns](https://support.google.com/cloud/answer/4430947).
	ListBillingAccounts(ctx context.Context, in *ListBillingAccountsRequest, opts ...grpc.CallOption) (*ListBillingAccountsResponse, error)
	// Lists the projects associated with a billing account. The current
	// authenticated user must be an [owner of the billing
	// account](https://support.google.com/cloud/answer/4430947).
	ListProjectBillingInfo(ctx context.Context, in *ListProjectBillingInfoRequest, opts ...grpc.CallOption) (*ListProjectBillingInfoResponse, error)
	// Gets the billing information for a project. The current authenticated user
	// must have [permission to view the
	// project](https://cloud.google.com/docs/permissions-overview#h.bgs0oxofvnoo
	// ).
	GetProjectBillingInfo(ctx context.Context, in *GetProjectBillingInfoRequest, opts ...grpc.CallOption) (*ProjectBillingInfo, error)
	// Sets or updates the billing account associated with a project. You specify
	// the new billing account by setting the `billing_account_name` in the
	// `ProjectBillingInfo` resource to the resource name of a billing account.
	// Associating a project with an open billing account enables billing on the
	// project and allows charges for resource usage. If the project already had a
	// billing account, this method changes the billing account used for resource
	// usage charges.
	//
	// *Note:* Incurred charges that have not yet been reported in the transaction
	// history of the Google Cloud Console may be billed to the new billing
	// account, even if the charge occurred before the new billing account was
	// assigned to the project.
	//
	// The current authenticated user must have ownership privileges for both the
	// [project](https://cloud.google.com/docs/permissions-overview#h.bgs0oxofvnoo
	// ) and the [billing
	// account](https://support.google.com/cloud/answer/4430947).
	//
	// You can disable billing on the project by setting the
	// `billing_account_name` field to empty. This action disassociates the
	// current billing account from the project. Any billable activity of your
	// in-use services will stop, and your application could stop functioning as
	// expected. Any unbilled charges to date will be billed to the previously
	// associated account. The current authenticated user must be either an owner
	// of the project or an owner of the billing account for the project.
	//
	// Note that associating a project with a *closed* billing account will have
	// much the same effect as disabling billing on the project: any paid
	// resources used by the project will be shut down. Thus, unless you wish to
	// disable billing, you should always call this method with the name of an
	// *open* billing account.
	UpdateProjectBillingInfo(ctx context.Context, in *UpdateProjectBillingInfoRequest, opts ...grpc.CallOption) (*ProjectBillingInfo, error)
}

type cloudBillingClient struct {
	cc *grpc.ClientConn
}

func NewCloudBillingClient(cc *grpc.ClientConn) CloudBillingClient {
	return &cloudBillingClient{cc}
}

func (c *cloudBillingClient) GetBillingAccount(ctx context.Context, in *GetBillingAccountRequest, opts ...grpc.CallOption) (*BillingAccount, error) {
	out := new(BillingAccount)
	err := c.cc.Invoke(ctx, "/google.cloud.billing.v1.CloudBilling/GetBillingAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBillingClient) ListBillingAccounts(ctx context.Context, in *ListBillingAccountsRequest, opts ...grpc.CallOption) (*ListBillingAccountsResponse, error) {
	out := new(ListBillingAccountsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.billing.v1.CloudBilling/ListBillingAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBillingClient) ListProjectBillingInfo(ctx context.Context, in *ListProjectBillingInfoRequest, opts ...grpc.CallOption) (*ListProjectBillingInfoResponse, error) {
	out := new(ListProjectBillingInfoResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.billing.v1.CloudBilling/ListProjectBillingInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBillingClient) GetProjectBillingInfo(ctx context.Context, in *GetProjectBillingInfoRequest, opts ...grpc.CallOption) (*ProjectBillingInfo, error) {
	out := new(ProjectBillingInfo)
	err := c.cc.Invoke(ctx, "/google.cloud.billing.v1.CloudBilling/GetProjectBillingInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBillingClient) UpdateProjectBillingInfo(ctx context.Context, in *UpdateProjectBillingInfoRequest, opts ...grpc.CallOption) (*ProjectBillingInfo, error) {
	out := new(ProjectBillingInfo)
	err := c.cc.Invoke(ctx, "/google.cloud.billing.v1.CloudBilling/UpdateProjectBillingInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudBillingServer is the server API for CloudBilling service.
type CloudBillingServer interface {
	// Gets information about a billing account. The current authenticated user
	// must be an [owner of the billing
	// account](https://support.google.com/cloud/answer/4430947).
	GetBillingAccount(context.Context, *GetBillingAccountRequest) (*BillingAccount, error)
	// Lists the billing accounts that the current authenticated user
	// [owns](https://support.google.com/cloud/answer/4430947).
	ListBillingAccounts(context.Context, *ListBillingAccountsRequest) (*ListBillingAccountsResponse, error)
	// Lists the projects associated with a billing account. The current
	// authenticated user must be an [owner of the billing
	// account](https://support.google.com/cloud/answer/4430947).
	ListProjectBillingInfo(context.Context, *ListProjectBillingInfoRequest) (*ListProjectBillingInfoResponse, error)
	// Gets the billing information for a project. The current authenticated user
	// must have [permission to view the
	// project](https://cloud.google.com/docs/permissions-overview#h.bgs0oxofvnoo
	// ).
	GetProjectBillingInfo(context.Context, *GetProjectBillingInfoRequest) (*ProjectBillingInfo, error)
	// Sets or updates the billing account associated with a project. You specify
	// the new billing account by setting the `billing_account_name` in the
	// `ProjectBillingInfo` resource to the resource name of a billing account.
	// Associating a project with an open billing account enables billing on the
	// project and allows charges for resource usage. If the project already had a
	// billing account, this method changes the billing account used for resource
	// usage charges.
	//
	// *Note:* Incurred charges that have not yet been reported in the transaction
	// history of the Google Cloud Console may be billed to the new billing
	// account, even if the charge occurred before the new billing account was
	// assigned to the project.
	//
	// The current authenticated user must have ownership privileges for both the
	// [project](https://cloud.google.com/docs/permissions-overview#h.bgs0oxofvnoo
	// ) and the [billing
	// account](https://support.google.com/cloud/answer/4430947).
	//
	// You can disable billing on the project by setting the
	// `billing_account_name` field to empty. This action disassociates the
	// current billing account from the project. Any billable activity of your
	// in-use services will stop, and your application could stop functioning as
	// expected. Any unbilled charges to date will be billed to the previously
	// associated account. The current authenticated user must be either an owner
	// of the project or an owner of the billing account for the project.
	//
	// Note that associating a project with a *closed* billing account will have
	// much the same effect as disabling billing on the project: any paid
	// resources used by the project will be shut down. Thus, unless you wish to
	// disable billing, you should always call this method with the name of an
	// *open* billing account.
	UpdateProjectBillingInfo(context.Context, *UpdateProjectBillingInfoRequest) (*ProjectBillingInfo, error)
}

func RegisterCloudBillingServer(s *grpc.Server, srv CloudBillingServer) {
	s.RegisterService(&_CloudBilling_serviceDesc, srv)
}

func _CloudBilling_GetBillingAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBillingAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBillingServer).GetBillingAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.billing.v1.CloudBilling/GetBillingAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBillingServer).GetBillingAccount(ctx, req.(*GetBillingAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBilling_ListBillingAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBillingAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBillingServer).ListBillingAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.billing.v1.CloudBilling/ListBillingAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBillingServer).ListBillingAccounts(ctx, req.(*ListBillingAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBilling_ListProjectBillingInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectBillingInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBillingServer).ListProjectBillingInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.billing.v1.CloudBilling/ListProjectBillingInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBillingServer).ListProjectBillingInfo(ctx, req.(*ListProjectBillingInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBilling_GetProjectBillingInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectBillingInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBillingServer).GetProjectBillingInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.billing.v1.CloudBilling/GetProjectBillingInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBillingServer).GetProjectBillingInfo(ctx, req.(*GetProjectBillingInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBilling_UpdateProjectBillingInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectBillingInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBillingServer).UpdateProjectBillingInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.billing.v1.CloudBilling/UpdateProjectBillingInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBillingServer).UpdateProjectBillingInfo(ctx, req.(*UpdateProjectBillingInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CloudBilling_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.billing.v1.CloudBilling",
	HandlerType: (*CloudBillingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBillingAccount",
			Handler:    _CloudBilling_GetBillingAccount_Handler,
		},
		{
			MethodName: "ListBillingAccounts",
			Handler:    _CloudBilling_ListBillingAccounts_Handler,
		},
		{
			MethodName: "ListProjectBillingInfo",
			Handler:    _CloudBilling_ListProjectBillingInfo_Handler,
		},
		{
			MethodName: "GetProjectBillingInfo",
			Handler:    _CloudBilling_GetProjectBillingInfo_Handler,
		},
		{
			MethodName: "UpdateProjectBillingInfo",
			Handler:    _CloudBilling_UpdateProjectBillingInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/billing/v1/cloud_billing.proto",
}

func init() {
	proto.RegisterFile("google/cloud/billing/v1/cloud_billing.proto", fileDescriptor_cloud_billing_868724f2532f6f7f)
}

var fileDescriptor_cloud_billing_868724f2532f6f7f = []byte{
	// 667 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x4e, 0xd4, 0x40,
	0x14, 0xce, 0x00, 0x12, 0xf6, 0x80, 0x20, 0x03, 0xe8, 0x66, 0x17, 0x10, 0xea, 0x0f, 0x28, 0xb1,
	0x15, 0xf0, 0xdf, 0xa8, 0x11, 0x63, 0x08, 0x89, 0x31, 0x9b, 0xaa, 0x89, 0xd1, 0x98, 0x66, 0x76,
	0x3b, 0x34, 0xd5, 0x32, 0x53, 0x77, 0x0a, 0x51, 0x8c, 0x37, 0xbe, 0x82, 0x7a, 0xe1, 0x85, 0x37,
	0x5e, 0xe8, 0x2b, 0x78, 0xed, 0x2b, 0xf8, 0x0a, 0xde, 0xfb, 0x0a, 0x66, 0xa6, 0x53, 0x65, 0xbb,
	0x9d, 0x85, 0xc6, 0xbb, 0xe6, 0x9b, 0x73, 0xfa, 0x7d, 0xe7, 0x9b, 0xef, 0x6c, 0x17, 0x96, 0x02,
	0xce, 0x83, 0x88, 0x3a, 0xad, 0x88, 0x6f, 0xfb, 0x4e, 0x33, 0x8c, 0xa2, 0x90, 0x05, 0xce, 0xce,
	0x72, 0x0a, 0x78, 0x1a, 0xb0, 0xe3, 0x36, 0x4f, 0x38, 0x3e, 0x96, 0x16, 0xdb, 0xea, 0xcc, 0xce,
	0xce, 0x76, 0x96, 0x6b, 0xd3, 0xfa, 0x2d, 0x24, 0x0e, 0x1d, 0xc2, 0x18, 0x4f, 0x48, 0x12, 0x72,
	0x26, 0xd2, 0x36, 0xeb, 0x29, 0x8c, 0xae, 0xa5, 0xb5, 0xb7, 0x5b, 0x2d, 0xbe, 0xcd, 0x12, 0x8c,
	0x61, 0x80, 0x91, 0x2d, 0x5a, 0x45, 0x73, 0x68, 0xb1, 0xe2, 0xaa, 0x67, 0x89, 0xf1, 0x98, 0xb2,
	0x6a, 0xdf, 0x1c, 0x5a, 0x1c, 0x72, 0xd5, 0x33, 0x9e, 0x87, 0x11, 0x3f, 0x14, 0x71, 0x44, 0x5e,
	0x7b, 0xaa, 0xbe, 0x5f, 0xd5, 0x0f, 0x6b, 0xec, 0x3e, 0xd9, 0xa2, 0xd6, 0x17, 0x04, 0xb8, 0xd1,
	0xe6, 0xcf, 0x69, 0x2b, 0xd1, 0x24, 0x1b, 0x6c, 0x93, 0x17, 0x32, 0xcc, 0x00, 0xc4, 0x69, 0xa5,
	0x17, 0xfa, 0x8a, 0xa7, 0xe2, 0x56, 0x34, 0xb2, 0xe1, 0xe3, 0xf3, 0x30, 0xa9, 0x47, 0xf2, 0x48,
	0xaa, 0x73, 0x2f, 0x29, 0x6e, 0x76, 0x8c, 0x20, 0xb9, 0xf1, 0x02, 0x8c, 0x65, 0x1d, 0x94, 0x91,
	0x66, 0x44, 0xfd, 0xea, 0x80, 0x52, 0x3f, 0xaa, 0xe1, 0xbb, 0x29, 0x6a, 0xd9, 0x50, 0x5d, 0xa7,
	0x49, 0xa7, 0x09, 0x2e, 0x7d, 0xb9, 0x4d, 0x45, 0xa1, 0x17, 0xd6, 0x63, 0xa8, 0xdd, 0x0b, 0x45,
	0xae, 0x41, 0x64, 0x1d, 0x75, 0xa8, 0xc4, 0x24, 0xa0, 0x9e, 0x08, 0x77, 0xd3, 0xb6, 0x43, 0xee,
	0x90, 0x04, 0x1e, 0x84, 0xbb, 0xe9, 0x90, 0xf2, 0x30, 0xe1, 0x2f, 0xb4, 0x99, 0x72, 0x48, 0x12,
	0xd0, 0x87, 0x12, 0xb0, 0x3e, 0x21, 0xa8, 0x17, 0xbe, 0x5a, 0xc4, 0x9c, 0x09, 0x8a, 0x5d, 0x38,
	0x92, 0x33, 0x41, 0x54, 0xd1, 0x5c, 0xff, 0xe2, 0xf0, 0xca, 0x82, 0x6d, 0xb8, 0x7d, 0x3b, 0x37,
	0xd7, 0x58, 0xa7, 0x53, 0x02, 0x9f, 0x86, 0x31, 0x46, 0x5f, 0x25, 0x5e, 0x97, 0xae, 0xc3, 0x12,
	0x6e, 0xfc, 0xd5, 0xc6, 0x61, 0x46, 0x4a, 0xeb, 0xbe, 0xcd, 0x1e, 0x56, 0x75, 0x9a, 0xd1, 0xd7,
	0xd3, 0x8c, 0xfe, 0xbc, 0x19, 0xdf, 0x10, 0xcc, 0x9a, 0x18, 0xb5, 0x1f, 0xcf, 0x60, 0x32, 0xcb,
	0x4c, 0xe6, 0x4b, 0xc8, 0x36, 0xb9, 0xf6, 0x64, 0xc9, 0xe8, 0x49, 0xc1, 0x2b, 0x71, 0xdc, 0x1d,
	0xd3, 0x83, 0x5a, 0xb3, 0x02, 0xd3, 0xeb, 0xb4, 0x9c, 0x33, 0xd6, 0x07, 0x04, 0xc7, 0x1f, 0xc5,
	0x3e, 0x49, 0x68, 0x39, 0x47, 0x4d, 0x23, 0x4b, 0x61, 0xff, 0x3f, 0xf2, 0xca, 0xef, 0x41, 0x18,
	0xb9, 0x23, 0x7b, 0x35, 0x88, 0x3f, 0x22, 0x18, 0xef, 0xda, 0x0e, 0xbc, 0x6c, 0xe4, 0x31, 0x6d,
	0x52, 0xed, 0xa0, 0x09, 0xb5, 0x4e, 0xbe, 0xfb, 0xf9, 0xeb, 0x7d, 0xdf, 0x2c, 0x9e, 0x96, 0x3f,
	0x74, 0x6f, 0xe4, 0xd0, 0x37, 0x72, 0x99, 0x75, 0xce, 0xbe, 0xc5, 0x9f, 0x11, 0x4c, 0x14, 0xac,
	0x0a, 0x5e, 0x35, 0xd2, 0x98, 0x77, 0xb6, 0x76, 0xa1, 0x5c, 0x53, 0x9a, 0x3e, 0xab, 0xae, 0x84,
	0x4e, 0xe1, 0x09, 0x29, 0x34, 0xbf, 0x56, 0xdf, 0x11, 0x1c, 0x2d, 0x4e, 0x2f, 0xbe, 0xd4, 0x93,
	0xcd, 0x18, 0x87, 0xda, 0xe5, 0xd2, 0x7d, 0x5a, 0xe8, 0x39, 0x25, 0x74, 0x01, 0x9f, 0xea, 0xe5,
	0xa8, 0xa3, 0xd3, 0x20, 0xf0, 0x57, 0x04, 0x53, 0x85, 0x79, 0xc6, 0x17, 0x7b, 0x5d, 0xbb, 0x59,
	0x78, 0x99, 0x54, 0x5a, 0x67, 0x94, 0xd8, 0x13, 0x78, 0xfe, 0x9f, 0xd8, 0x4c, 0x99, 0x54, 0xd9,
	0xdc, 0x23, 0xe7, 0x07, 0x82, 0xaa, 0x69, 0x87, 0xf0, 0x15, 0x23, 0xe9, 0x3e, 0x6b, 0x57, 0x4e,
	0xee, 0x2d, 0x25, 0xf7, 0x6a, 0x6d, 0x7f, 0xb9, 0xd7, 0x0a, 0x17, 0x77, 0xad, 0x0d, 0xf5, 0x16,
	0xdf, 0x32, 0x51, 0xae, 0x8d, 0xef, 0xdd, 0xc6, 0x86, 0xfc, 0x62, 0x37, 0xd0, 0x93, 0x9b, 0xba,
	0x3a, 0xe0, 0x11, 0x61, 0x81, 0xcd, 0xdb, 0x81, 0x13, 0x50, 0xa6, 0xbe, 0xe7, 0x4e, 0x7a, 0x44,
	0xe2, 0x50, 0x74, 0xfd, 0x6d, 0xb8, 0xae, 0x1f, 0x9b, 0x83, 0xaa, 0x74, 0xf5, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc0, 0x01, 0x24, 0x32, 0x60, 0x08, 0x00, 0x00,
}
