// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/user/reset_password.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ResetPassword struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	User                 *User                `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Token                string               `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ResetPassword) Reset()         { *m = ResetPassword{} }
func (m *ResetPassword) String() string { return proto.CompactTextString(m) }
func (*ResetPassword) ProtoMessage()    {}
func (*ResetPassword) Descriptor() ([]byte, []int) {
	return fileDescriptor_174109f3c45dea44, []int{0}
}

func (m *ResetPassword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPassword.Unmarshal(m, b)
}
func (m *ResetPassword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPassword.Marshal(b, m, deterministic)
}
func (m *ResetPassword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPassword.Merge(m, src)
}
func (m *ResetPassword) XXX_Size() int {
	return xxx_messageInfo_ResetPassword.Size(m)
}
func (m *ResetPassword) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPassword.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPassword proto.InternalMessageInfo

func (m *ResetPassword) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ResetPassword) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ResetPassword) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ResetPassword) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*ResetPassword)(nil), "aggregates.user.ResetPassword")
}

func init() {
	proto.RegisterFile("aggregates/user/reset_password.proto", fileDescriptor_174109f3c45dea44)
}

var fileDescriptor_174109f3c45dea44 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xe9, 0xee, 0x2a, 0x6c, 0x44, 0x85, 0xa2, 0x50, 0x7a, 0xb1, 0x88, 0x87, 0x7a, 0xd8,
	0x04, 0xf4, 0xe4, 0x51, 0x2f, 0xe2, 0x4d, 0x8a, 0x5e, 0xbc, 0x2c, 0x69, 0x33, 0x9b, 0x86, 0x6e,
	0x9a, 0x32, 0x33, 0xd5, 0x1f, 0xe3, 0x9f, 0x95, 0xb6, 0xbb, 0x08, 0xbd, 0x04, 0x1e, 0x7c, 0x79,
	0xdf, 0x63, 0xc4, 0x9d, 0xb6, 0x16, 0xc1, 0x6a, 0x06, 0x52, 0x3d, 0x01, 0x2a, 0x04, 0x02, 0xde,
	0x76, 0x9a, 0xe8, 0x27, 0xa0, 0x91, 0x1d, 0x06, 0x0e, 0xf1, 0xe5, 0x3f, 0x25, 0x07, 0x2a, 0x4d,
	0xe7, 0xdf, 0x86, 0x67, 0x82, 0xd3, 0x1b, 0x1b, 0x82, 0xdd, 0x83, 0x1a, 0x53, 0xd9, 0xef, 0x14,
	0x3b, 0x0f, 0xc4, 0xda, 0x77, 0x13, 0x70, 0xfb, 0x1b, 0x89, 0xf3, 0x62, 0xd0, 0xbc, 0x1f, 0x2c,
	0xf1, 0x85, 0x58, 0x38, 0x93, 0x44, 0x59, 0x94, 0x2f, 0x8b, 0x85, 0x33, 0xf1, 0xbd, 0x58, 0x0d,
	0x85, 0xc9, 0x22, 0x8b, 0xf2, 0xb3, 0x87, 0x6b, 0x39, 0xd3, 0xcb, 0x4f, 0x02, 0x2c, 0x46, 0x24,
	0xbe, 0x12, 0x27, 0x1c, 0x1a, 0x68, 0x93, 0x65, 0x16, 0xe5, 0xeb, 0x62, 0x0a, 0xf1, 0x93, 0x10,
	0x15, 0x82, 0x66, 0x30, 0x5b, 0xcd, 0xc9, 0x6a, 0xac, 0x49, 0xe5, 0x34, 0x4c, 0x1e, 0x87, 0xc9,
	0x8f, 0xe3, 0xb0, 0x62, 0x7d, 0xa0, 0x9f, 0xf9, 0xe5, 0xed, 0xeb, 0xd5, 0x3a, 0xae, 0xfb, 0x52,
	0x56, 0xc1, 0x2b, 0xb7, 0xaf, 0xb5, 0xf7, 0xb5, 0x31, 0xaa, 0xe9, 0x8d, 0x6e, 0xdc, 0x86, 0x38,
	0x20, 0xec, 0x30, 0xb4, 0xbc, 0x21, 0xc0, 0x6f, 0x57, 0x81, 0x82, 0x96, 0x1d, 0x3b, 0x20, 0x35,
	0x3b, 0x4a, 0x79, 0x3a, 0x9a, 0x1e, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x09, 0x04, 0xf9, 0x21,
	0x65, 0x01, 0x00, 0x00,
}
