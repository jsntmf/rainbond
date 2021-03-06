// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/squash/v3/squash.proto

package envoy_extensions_filters_http_squash_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Squash struct {
	Cluster              string             `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	AttachmentTemplate   *_struct.Struct    `protobuf:"bytes,2,opt,name=attachment_template,json=attachmentTemplate,proto3" json:"attachment_template,omitempty"`
	RequestTimeout       *duration.Duration `protobuf:"bytes,3,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
	AttachmentTimeout    *duration.Duration `protobuf:"bytes,4,opt,name=attachment_timeout,json=attachmentTimeout,proto3" json:"attachment_timeout,omitempty"`
	AttachmentPollPeriod *duration.Duration `protobuf:"bytes,5,opt,name=attachment_poll_period,json=attachmentPollPeriod,proto3" json:"attachment_poll_period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Squash) Reset()         { *m = Squash{} }
func (m *Squash) String() string { return proto.CompactTextString(m) }
func (*Squash) ProtoMessage()    {}
func (*Squash) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dba825384a5c607, []int{0}
}

func (m *Squash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Squash.Unmarshal(m, b)
}
func (m *Squash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Squash.Marshal(b, m, deterministic)
}
func (m *Squash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Squash.Merge(m, src)
}
func (m *Squash) XXX_Size() int {
	return xxx_messageInfo_Squash.Size(m)
}
func (m *Squash) XXX_DiscardUnknown() {
	xxx_messageInfo_Squash.DiscardUnknown(m)
}

var xxx_messageInfo_Squash proto.InternalMessageInfo

func (m *Squash) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Squash) GetAttachmentTemplate() *_struct.Struct {
	if m != nil {
		return m.AttachmentTemplate
	}
	return nil
}

func (m *Squash) GetRequestTimeout() *duration.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentTimeout() *duration.Duration {
	if m != nil {
		return m.AttachmentTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentPollPeriod() *duration.Duration {
	if m != nil {
		return m.AttachmentPollPeriod
	}
	return nil
}

func init() {
	proto.RegisterType((*Squash)(nil), "envoy.extensions.filters.http.squash.v3.Squash")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/squash/v3/squash.proto", fileDescriptor_2dba825384a5c607)
}

var fileDescriptor_2dba825384a5c607 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x8b, 0x1a, 0x31,
	0x14, 0xc7, 0x19, 0xb5, 0x4a, 0x23, 0xb4, 0x74, 0x5a, 0xea, 0x54, 0x8a, 0x68, 0x2f, 0xb5, 0x14,
	0x92, 0xa2, 0xed, 0xa5, 0xc7, 0x61, 0x59, 0xbc, 0xed, 0xa0, 0xde, 0x25, 0xce, 0xc4, 0x31, 0x10,
	0x93, 0x31, 0x79, 0x19, 0xf4, 0x0b, 0x2c, 0xfb, 0x19, 0xf6, 0xfb, 0xec, 0x97, 0xda, 0xd3, 0x62,
	0x26, 0xa2, 0xbb, 0x7b, 0x70, 0x6f, 0x61, 0xde, 0xfb, 0xfd, 0xe6, 0xcf, 0x7b, 0x0f, 0xfd, 0x65,
	0xb2, 0x54, 0x7b, 0xc2, 0x76, 0xc0, 0xa4, 0xe1, 0x4a, 0x1a, 0xb2, 0xe2, 0x02, 0x98, 0x36, 0x64,
	0x0d, 0x50, 0x10, 0xb3, 0xb5, 0xd4, 0xac, 0x49, 0x39, 0xf6, 0x2f, 0x5c, 0x68, 0x05, 0x2a, 0xfc,
	0xe9, 0x28, 0x7c, 0xa2, 0xb0, 0xa7, 0xf0, 0x81, 0xc2, 0xbe, 0xb7, 0x1c, 0x77, 0x7b, 0xb9, 0x52,
	0xb9, 0x60, 0xc4, 0x61, 0x4b, 0xbb, 0x22, 0x99, 0xd5, 0x14, 0xb8, 0x92, 0x95, 0xa8, 0xfb, 0xfd,
	0x65, 0xdd, 0x80, 0xb6, 0x29, 0xf8, 0xea, 0xc0, 0x66, 0x05, 0x25, 0x54, 0x4a, 0x05, 0x0e, 0x32,
	0xa4, 0x64, 0xfa, 0xf0, 0x3f, 0x2e, 0x73, 0xdf, 0xd2, 0x29, 0xa9, 0xe0, 0x19, 0x05, 0x46, 0x8e,
	0x8f, 0xaa, 0xf0, 0xe3, 0xb6, 0x8e, 0x9a, 0x33, 0x97, 0x23, 0x1c, 0xa0, 0x56, 0x2a, 0xac, 0x01,
	0xa6, 0xa3, 0xa0, 0x1f, 0x0c, 0xdf, 0xc7, 0xad, 0xc7, 0xb8, 0xa1, 0x6b, 0xfd, 0x60, 0x7a, 0xfc,
	0x1e, 0x4e, 0xd0, 0x67, 0x0a, 0x40, 0xd3, 0xf5, 0x86, 0x49, 0x58, 0x00, 0xdb, 0x14, 0x82, 0x02,
	0x8b, 0x6a, 0xfd, 0x60, 0xd8, 0x1e, 0x75, 0x70, 0x95, 0x12, 0x1f, 0x53, 0xe2, 0x99, 0x4b, 0x39,
	0x0d, 0x4f, 0xcc, 0xdc, 0x23, 0x61, 0x8c, 0x3e, 0x6a, 0xb6, 0xb5, 0xcc, 0xc0, 0x02, 0xf8, 0x86,
	0x29, 0x0b, 0x51, 0xdd, 0x59, 0xbe, 0xbd, 0xb2, 0x5c, 0xf9, 0x59, 0x4c, 0x3f, 0x78, 0x62, 0x5e,
	0x01, 0xe1, 0x04, 0x85, 0xe7, 0x69, 0xbc, 0xa6, 0x71, 0x49, 0xf3, 0xe9, 0x2c, 0x8e, 0x37, 0xdd,
	0xa0, 0xaf, 0x67, 0xa6, 0x42, 0x09, 0xb1, 0x28, 0x98, 0xe6, 0x2a, 0x8b, 0xde, 0x5d, 0xb2, 0x7d,
	0x39, 0x81, 0x89, 0x12, 0x22, 0x71, 0xd8, 0xff, 0x3f, 0xf7, 0x0f, 0x77, 0xbd, 0xdf, 0xe8, 0x57,
	0x75, 0x00, 0xa9, 0x92, 0x2b, 0x9e, 0xfb, 0xe5, 0x3f, 0xdf, 0xfd, 0x08, 0x57, 0xd3, 0x8f, 0xaf,
	0xd1, 0x3f, 0xae, 0xb0, 0xeb, 0x2f, 0xb4, 0xda, 0xed, 0xf1, 0x1b, 0x6f, 0x27, 0x6e, 0x57, 0x82,
	0xe4, 0x90, 0x2c, 0x09, 0x96, 0x4d, 0x17, 0x71, 0xfc, 0x14, 0x00, 0x00, 0xff, 0xff, 0x29, 0xf3,
	0xcd, 0x3b, 0xb2, 0x02, 0x00, 0x00,
}
