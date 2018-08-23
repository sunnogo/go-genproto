// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/streetview/publish/v1/resources.proto

package publish // import "github.com/sunnogo/go-genproto/googleapis/streetview/publish/v1"

import proto "github.com/sunnogo/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/sunnogo/protobuf/ptypes/timestamp"
import _ "github.com/sunnogo/go-genproto/googleapis/api/annotations"
import latlng "github.com/sunnogo/go-genproto/googleapis/type/latlng"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Upload reference for media files.
type UploadRef struct {
	// Required. An upload reference should be unique for each user. It follows
	// the form:
	// "https://streetviewpublish.googleapis.com/media/user/<account_id>/photo/<upload_reference>"
	UploadUrl            string   `protobuf:"bytes,1,opt,name=upload_url,json=uploadUrl,proto3" json:"upload_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadRef) Reset()         { *m = UploadRef{} }
func (m *UploadRef) String() string { return proto.CompactTextString(m) }
func (*UploadRef) ProtoMessage()    {}
func (*UploadRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_8db553b896211977, []int{0}
}
func (m *UploadRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRef.Unmarshal(m, b)
}
func (m *UploadRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRef.Marshal(b, m, deterministic)
}
func (dst *UploadRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRef.Merge(dst, src)
}
func (m *UploadRef) XXX_Size() int {
	return xxx_messageInfo_UploadRef.Size(m)
}
func (m *UploadRef) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRef.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRef proto.InternalMessageInfo

func (m *UploadRef) GetUploadUrl() string {
	if m != nil {
		return m.UploadUrl
	}
	return ""
}

// Identifier for a photo.
type PhotoId struct {
	// Required. A base64 encoded identifier.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhotoId) Reset()         { *m = PhotoId{} }
func (m *PhotoId) String() string { return proto.CompactTextString(m) }
func (*PhotoId) ProtoMessage()    {}
func (*PhotoId) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_8db553b896211977, []int{1}
}
func (m *PhotoId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhotoId.Unmarshal(m, b)
}
func (m *PhotoId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhotoId.Marshal(b, m, deterministic)
}
func (dst *PhotoId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhotoId.Merge(dst, src)
}
func (m *PhotoId) XXX_Size() int {
	return xxx_messageInfo_PhotoId.Size(m)
}
func (m *PhotoId) XXX_DiscardUnknown() {
	xxx_messageInfo_PhotoId.DiscardUnknown(m)
}

var xxx_messageInfo_PhotoId proto.InternalMessageInfo

func (m *PhotoId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Level information containing level number and its corresponding name.
type Level struct {
	// Floor number, used for ordering. 0 indicates the ground level, 1 indicates
	// the first level above ground level, -1 indicates the first level under
	// ground level. Non-integer values are OK.
	Number float64 `protobuf:"fixed64,1,opt,name=number,proto3" json:"number,omitempty"`
	// Required. A name assigned to this Level, restricted to 3 characters.
	// Consider how the elevator buttons would be labeled for this level if there
	// was an elevator.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Level) Reset()         { *m = Level{} }
func (m *Level) String() string { return proto.CompactTextString(m) }
func (*Level) ProtoMessage()    {}
func (*Level) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_8db553b896211977, []int{2}
}
func (m *Level) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Level.Unmarshal(m, b)
}
func (m *Level) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Level.Marshal(b, m, deterministic)
}
func (dst *Level) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Level.Merge(dst, src)
}
func (m *Level) XXX_Size() int {
	return xxx_messageInfo_Level.Size(m)
}
func (m *Level) XXX_DiscardUnknown() {
	xxx_messageInfo_Level.DiscardUnknown(m)
}

var xxx_messageInfo_Level proto.InternalMessageInfo

func (m *Level) GetNumber() float64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Level) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Raw pose measurement for an entity.
type Pose struct {
	// Latitude and longitude pair of the pose, as explained here:
	// https://cloud.google.com/datastore/docs/reference/rest/Shared.Types/LatLng
	// When creating a photo, if the latitude and longitude pair are not provided
	// here, the geolocation from the exif header will be used.
	// If the latitude and longitude pair is not provided and cannot be found in
	// the exif header, the create photo process will fail.
	LatLngPair *latlng.LatLng `protobuf:"bytes,1,opt,name=lat_lng_pair,json=latLngPair,proto3" json:"lat_lng_pair,omitempty"`
	// Altitude of the pose in meters above ground level (as defined by WGS84).
	// NaN indicates an unmeasured quantity.
	Altitude float64 `protobuf:"fixed64,2,opt,name=altitude,proto3" json:"altitude,omitempty"`
	// Compass heading, measured at the center of the photo in degrees clockwise
	// from North. Value must be >=0 and <360.
	// NaN indicates an unmeasured quantity.
	Heading float64 `protobuf:"fixed64,3,opt,name=heading,proto3" json:"heading,omitempty"`
	// Pitch, measured at the center of the photo in degrees. Value must be >=-90
	// and <= 90. A value of -90 means looking directly down, and a value of 90
	// means looking directly up.
	// NaN indicates an unmeasured quantity.
	Pitch float64 `protobuf:"fixed64,4,opt,name=pitch,proto3" json:"pitch,omitempty"`
	// Roll, measured in degrees. Value must be >= 0 and <360. A value of 0
	// means level with the horizon.
	// NaN indicates an unmeasured quantity.
	Roll float64 `protobuf:"fixed64,5,opt,name=roll,proto3" json:"roll,omitempty"`
	// Level (the floor in a building) used to configure vertical navigation.
	Level                *Level   `protobuf:"bytes,7,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pose) Reset()         { *m = Pose{} }
func (m *Pose) String() string { return proto.CompactTextString(m) }
func (*Pose) ProtoMessage()    {}
func (*Pose) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_8db553b896211977, []int{3}
}
func (m *Pose) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pose.Unmarshal(m, b)
}
func (m *Pose) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pose.Marshal(b, m, deterministic)
}
func (dst *Pose) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pose.Merge(dst, src)
}
func (m *Pose) XXX_Size() int {
	return xxx_messageInfo_Pose.Size(m)
}
func (m *Pose) XXX_DiscardUnknown() {
	xxx_messageInfo_Pose.DiscardUnknown(m)
}

var xxx_messageInfo_Pose proto.InternalMessageInfo

func (m *Pose) GetLatLngPair() *latlng.LatLng {
	if m != nil {
		return m.LatLngPair
	}
	return nil
}

func (m *Pose) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Pose) GetHeading() float64 {
	if m != nil {
		return m.Heading
	}
	return 0
}

func (m *Pose) GetPitch() float64 {
	if m != nil {
		return m.Pitch
	}
	return 0
}

func (m *Pose) GetRoll() float64 {
	if m != nil {
		return m.Roll
	}
	return 0
}

func (m *Pose) GetLevel() *Level {
	if m != nil {
		return m.Level
	}
	return nil
}

// Place metadata for an entity.
type Place struct {
	// Required. Place identifier, as described in
	// https://developers.google.com/places/place-id.
	PlaceId              string   `protobuf:"bytes,1,opt,name=place_id,json=placeId,proto3" json:"place_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Place) Reset()         { *m = Place{} }
func (m *Place) String() string { return proto.CompactTextString(m) }
func (*Place) ProtoMessage()    {}
func (*Place) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_8db553b896211977, []int{4}
}
func (m *Place) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Place.Unmarshal(m, b)
}
func (m *Place) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Place.Marshal(b, m, deterministic)
}
func (dst *Place) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Place.Merge(dst, src)
}
func (m *Place) XXX_Size() int {
	return xxx_messageInfo_Place.Size(m)
}
func (m *Place) XXX_DiscardUnknown() {
	xxx_messageInfo_Place.DiscardUnknown(m)
}

var xxx_messageInfo_Place proto.InternalMessageInfo

func (m *Place) GetPlaceId() string {
	if m != nil {
		return m.PlaceId
	}
	return ""
}

// A connection is the link from a source photo to a destination photo.
type Connection struct {
	// Required. The destination of the connection from the containing photo to
	// another photo.
	Target               *PhotoId `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connection) Reset()         { *m = Connection{} }
func (m *Connection) String() string { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()    {}
func (*Connection) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_8db553b896211977, []int{5}
}
func (m *Connection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connection.Unmarshal(m, b)
}
func (m *Connection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connection.Marshal(b, m, deterministic)
}
func (dst *Connection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connection.Merge(dst, src)
}
func (m *Connection) XXX_Size() int {
	return xxx_messageInfo_Connection.Size(m)
}
func (m *Connection) XXX_DiscardUnknown() {
	xxx_messageInfo_Connection.DiscardUnknown(m)
}

var xxx_messageInfo_Connection proto.InternalMessageInfo

func (m *Connection) GetTarget() *PhotoId {
	if m != nil {
		return m.Target
	}
	return nil
}

// Photo is used to store 360 photos along with photo metadata.
type Photo struct {
	// Output only. Identifier for the photo, which is unique among all photos in
	// Google.
	PhotoId *PhotoId `protobuf:"bytes,1,opt,name=photo_id,json=photoId,proto3" json:"photo_id,omitempty"`
	// Required (when creating photo). Input only. The resource URL where the
	// photo bytes are uploaded to.
	UploadReference *UploadRef `protobuf:"bytes,2,opt,name=upload_reference,json=uploadReference,proto3" json:"upload_reference,omitempty"`
	// Output only. The download URL for the photo bytes. This field is set only
	// when the `view` parameter in a `GetPhotoRequest` is set to
	// `INCLUDE_DOWNLOAD_URL`.
	DownloadUrl string `protobuf:"bytes,3,opt,name=download_url,json=downloadUrl,proto3" json:"download_url,omitempty"`
	// Output only. The thumbnail URL for showing a preview of the given photo.
	ThumbnailUrl string `protobuf:"bytes,9,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"`
	// Output only. The share link for the photo.
	ShareLink string `protobuf:"bytes,11,opt,name=share_link,json=shareLink,proto3" json:"share_link,omitempty"`
	// Pose of the photo.
	Pose *Pose `protobuf:"bytes,4,opt,name=pose,proto3" json:"pose,omitempty"`
	// Connections to other photos. A connection represents the link from this
	// photo to another photo.
	Connections []*Connection `protobuf:"bytes,5,rep,name=connections,proto3" json:"connections,omitempty"`
	// Absolute time when the photo was captured.
	// When the photo has no exif timestamp, this is used to set a timestamp in
	// the photo metadata.
	CaptureTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=capture_time,json=captureTime,proto3" json:"capture_time,omitempty"`
	// Places where this photo belongs.
	Places []*Place `protobuf:"bytes,7,rep,name=places,proto3" json:"places,omitempty"`
	// Output only. View count of the photo.
	ViewCount            int64    `protobuf:"varint,10,opt,name=view_count,json=viewCount,proto3" json:"view_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Photo) Reset()         { *m = Photo{} }
func (m *Photo) String() string { return proto.CompactTextString(m) }
func (*Photo) ProtoMessage()    {}
func (*Photo) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_8db553b896211977, []int{6}
}
func (m *Photo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Photo.Unmarshal(m, b)
}
func (m *Photo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Photo.Marshal(b, m, deterministic)
}
func (dst *Photo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Photo.Merge(dst, src)
}
func (m *Photo) XXX_Size() int {
	return xxx_messageInfo_Photo.Size(m)
}
func (m *Photo) XXX_DiscardUnknown() {
	xxx_messageInfo_Photo.DiscardUnknown(m)
}

var xxx_messageInfo_Photo proto.InternalMessageInfo

func (m *Photo) GetPhotoId() *PhotoId {
	if m != nil {
		return m.PhotoId
	}
	return nil
}

func (m *Photo) GetUploadReference() *UploadRef {
	if m != nil {
		return m.UploadReference
	}
	return nil
}

func (m *Photo) GetDownloadUrl() string {
	if m != nil {
		return m.DownloadUrl
	}
	return ""
}

func (m *Photo) GetThumbnailUrl() string {
	if m != nil {
		return m.ThumbnailUrl
	}
	return ""
}

func (m *Photo) GetShareLink() string {
	if m != nil {
		return m.ShareLink
	}
	return ""
}

func (m *Photo) GetPose() *Pose {
	if m != nil {
		return m.Pose
	}
	return nil
}

func (m *Photo) GetConnections() []*Connection {
	if m != nil {
		return m.Connections
	}
	return nil
}

func (m *Photo) GetCaptureTime() *timestamp.Timestamp {
	if m != nil {
		return m.CaptureTime
	}
	return nil
}

func (m *Photo) GetPlaces() []*Place {
	if m != nil {
		return m.Places
	}
	return nil
}

func (m *Photo) GetViewCount() int64 {
	if m != nil {
		return m.ViewCount
	}
	return 0
}

func init() {
	proto.RegisterType((*UploadRef)(nil), "google.streetview.publish.v1.UploadRef")
	proto.RegisterType((*PhotoId)(nil), "google.streetview.publish.v1.PhotoId")
	proto.RegisterType((*Level)(nil), "google.streetview.publish.v1.Level")
	proto.RegisterType((*Pose)(nil), "google.streetview.publish.v1.Pose")
	proto.RegisterType((*Place)(nil), "google.streetview.publish.v1.Place")
	proto.RegisterType((*Connection)(nil), "google.streetview.publish.v1.Connection")
	proto.RegisterType((*Photo)(nil), "google.streetview.publish.v1.Photo")
}

func init() {
	proto.RegisterFile("google/streetview/publish/v1/resources.proto", fileDescriptor_resources_8db553b896211977)
}

var fileDescriptor_resources_8db553b896211977 = []byte{
	// 651 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdb, 0x6a, 0xdb, 0x4c,
	0x10, 0xc6, 0xf1, 0x29, 0x1e, 0xf9, 0x3f, 0xb0, 0xff, 0x4f, 0x51, 0x4c, 0x43, 0x53, 0x85, 0x52,
	0x53, 0x8a, 0x44, 0x1c, 0x5a, 0x28, 0x21, 0x50, 0x92, 0xab, 0xb4, 0xbe, 0x30, 0xdb, 0xa6, 0x17,
	0xbd, 0x11, 0x6b, 0x69, 0x22, 0x2f, 0x59, 0xef, 0x2e, 0xab, 0x95, 0x43, 0x9f, 0xa1, 0x8f, 0xd1,
	0x97, 0xea, 0xe3, 0x14, 0xad, 0x56, 0x4e, 0x2f, 0x82, 0xd3, 0x2b, 0xcf, 0x7c, 0xf3, 0x7d, 0xe3,
	0x39, 0xad, 0xe0, 0x75, 0xa1, 0x54, 0x21, 0x30, 0x29, 0xad, 0x41, 0xb4, 0x1b, 0x8e, 0x77, 0x89,
	0xae, 0x96, 0x82, 0x97, 0xab, 0x64, 0x73, 0x92, 0x18, 0x2c, 0x55, 0x65, 0x32, 0x2c, 0x63, 0x6d,
	0x94, 0x55, 0xe4, 0x69, 0xc3, 0x8e, 0xef, 0xd9, 0xb1, 0x67, 0xc7, 0x9b, 0x93, 0x89, 0x8f, 0x26,
	0x4c, 0xf3, 0x84, 0x49, 0xa9, 0x2c, 0xb3, 0x5c, 0x49, 0xaf, 0x9d, 0x3c, 0xf3, 0x51, 0xe7, 0x2d,
	0xab, 0x9b, 0xc4, 0xf2, 0x35, 0x96, 0x96, 0xad, 0xb5, 0x27, 0x84, 0x9e, 0x60, 0xbf, 0x69, 0x4c,
	0x04, 0xb3, 0x42, 0x16, 0x4d, 0x24, 0x7a, 0x05, 0xa3, 0x6b, 0x2d, 0x14, 0xcb, 0x29, 0xde, 0x90,
	0x43, 0x80, 0xca, 0x39, 0x69, 0x65, 0x44, 0xd8, 0x39, 0xea, 0x4c, 0x47, 0x74, 0xd4, 0x20, 0xd7,
	0x46, 0x44, 0x07, 0x30, 0x5c, 0xac, 0x94, 0x55, 0x57, 0x39, 0xf9, 0x1b, 0xf6, 0x78, 0xee, 0x19,
	0x7b, 0x3c, 0x8f, 0x4e, 0xa1, 0x3f, 0xc7, 0x0d, 0x0a, 0xf2, 0x04, 0x06, 0xb2, 0x5a, 0x2f, 0xd1,
	0xb8, 0x60, 0x87, 0x7a, 0x8f, 0x10, 0xe8, 0x49, 0xb6, 0xc6, 0x70, 0xcf, 0x49, 0x9c, 0x1d, 0xfd,
	0xec, 0x40, 0x6f, 0xa1, 0x4a, 0x24, 0x6f, 0x60, 0x2c, 0x98, 0x4d, 0x85, 0x2c, 0x52, 0xcd, 0x78,
	0x23, 0x0d, 0x66, 0xff, 0xc5, 0x7e, 0x24, 0x75, 0xd5, 0xf1, 0x9c, 0xd9, 0xb9, 0x2c, 0x28, 0x08,
	0xf7, 0xbb, 0x60, 0xdc, 0x90, 0x09, 0xec, 0x33, 0x61, 0xb9, 0xad, 0xf2, 0x26, 0x6f, 0x87, 0x6e,
	0x7d, 0x12, 0xc2, 0x70, 0x85, 0x2c, 0xe7, 0xb2, 0x08, 0xbb, 0x2e, 0xd4, 0xba, 0xe4, 0x7f, 0xe8,
	0x6b, 0x6e, 0xb3, 0x55, 0xd8, 0x73, 0x78, 0xe3, 0xd4, 0xf5, 0x19, 0x25, 0x44, 0xd8, 0x77, 0xa0,
	0xb3, 0xc9, 0x3b, 0xe8, 0x8b, 0xba, 0xa9, 0x70, 0xe8, 0xea, 0x39, 0x8e, 0x77, 0xad, 0x28, 0x76,
	0xfd, 0xd3, 0x46, 0x11, 0x45, 0xd0, 0x5f, 0x08, 0x96, 0x21, 0x39, 0x80, 0x7d, 0x5d, 0x1b, 0xe9,
	0x76, 0x5c, 0x43, 0xe7, 0x5f, 0xe5, 0xd1, 0x47, 0x80, 0x4b, 0x25, 0x25, 0x66, 0xf5, 0x2a, 0xc9,
	0x39, 0x0c, 0x2c, 0x33, 0x05, 0x5a, 0xdf, 0xfd, 0x8b, 0xdd, 0xff, 0xe6, 0x17, 0x41, 0xbd, 0x28,
	0xfa, 0xd1, 0x83, 0xbe, 0xc3, 0xc8, 0x7b, 0xd8, 0xd7, 0xb5, 0xd1, 0xfe, 0xe3, 0x1f, 0xa7, 0x1a,
	0x6a, 0xbf, 0x5c, 0x0a, 0xff, 0xfa, 0x33, 0x30, 0x78, 0x83, 0x06, 0x65, 0xd6, 0xcc, 0x37, 0x98,
	0xbd, 0xdc, 0x9d, 0x69, 0x7b, 0x49, 0xf4, 0x9f, 0xaa, 0x35, 0x1b, 0x3d, 0x79, 0x0e, 0xe3, 0x5c,
	0xdd, 0xc9, 0xed, 0x71, 0x75, 0xdd, 0x2c, 0x82, 0x16, 0xbb, 0x36, 0x82, 0x1c, 0xc3, 0x5f, 0x76,
	0x55, 0xad, 0x97, 0x92, 0x71, 0xe1, 0x38, 0x23, 0xc7, 0x19, 0x6f, 0xc1, 0x9a, 0x74, 0x08, 0x50,
	0xae, 0x98, 0xc1, 0x54, 0x70, 0x79, 0x1b, 0x06, 0xcd, 0x89, 0x3a, 0x64, 0xce, 0xe5, 0x2d, 0x79,
	0x0b, 0x3d, 0xad, 0x4a, 0x74, 0xbb, 0x0d, 0x66, 0xd1, 0x23, 0x8d, 0xab, 0x12, 0xa9, 0xe3, 0x93,
	0x0f, 0x10, 0x64, 0xdb, 0x5d, 0x94, 0x61, 0xff, 0xa8, 0x3b, 0x0d, 0x66, 0xd3, 0xdd, 0xf2, 0xfb,
	0xe5, 0xd1, 0xdf, 0xc5, 0xe4, 0x1c, 0xc6, 0x19, 0xd3, 0xb6, 0x32, 0x98, 0xd6, 0xef, 0x30, 0x1c,
	0xb8, 0x5a, 0x26, 0x6d, 0xb2, 0xf6, 0x91, 0xc6, 0x9f, 0xdb, 0x47, 0x4a, 0x03, 0xcf, 0xaf, 0x11,
	0x72, 0x06, 0x03, 0x77, 0x21, 0x65, 0x38, 0x74, 0x55, 0x3c, 0x72, 0x76, 0xee, 0xcc, 0xa8, 0x97,
	0xd4, 0xe3, 0xa9, 0x09, 0x69, 0xa6, 0x2a, 0x69, 0x43, 0x38, 0xea, 0x4c, 0xbb, 0x74, 0x54, 0x23,
	0x97, 0x35, 0x70, 0xf1, 0xbd, 0x03, 0xd3, 0x4c, 0xad, 0xdb, 0x8c, 0x05, 0xaa, 0xb8, 0x2a, 0xb2,
	0x87, 0x33, 0x5f, 0x4c, 0x3e, 0x39, 0xf8, 0x0b, 0xc7, 0xbb, 0x45, 0x83, 0xd2, 0xf6, 0x9b, 0xf5,
	0xf5, 0xb2, 0xcd, 0xa0, 0x04, 0x93, 0x45, 0xac, 0x4c, 0x91, 0x14, 0x28, 0x5d, 0x6b, 0x49, 0x13,
	0x62, 0x9a, 0x97, 0x0f, 0x7f, 0xfa, 0xce, 0xbc, 0xb9, 0x1c, 0x38, 0xfe, 0xe9, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xad, 0x4e, 0x7a, 0x51, 0x29, 0x05, 0x00, 0x00,
}
