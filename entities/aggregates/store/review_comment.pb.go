// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/store/review_comment.proto

package store

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

type ReviewComment struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	ItemReview           *ItemReview          `protobuf:"bytes,3,opt,name=item_review,json=itemReview,proto3" json:"item_review,omitempty"`
	UserUuid             string               `protobuf:"bytes,4,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Comment              string               `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ReviewComment) Reset()         { *m = ReviewComment{} }
func (m *ReviewComment) String() string { return proto.CompactTextString(m) }
func (*ReviewComment) ProtoMessage()    {}
func (*ReviewComment) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f9167cebc150470, []int{0}
}

func (m *ReviewComment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReviewComment.Unmarshal(m, b)
}
func (m *ReviewComment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReviewComment.Marshal(b, m, deterministic)
}
func (m *ReviewComment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReviewComment.Merge(m, src)
}
func (m *ReviewComment) XXX_Size() int {
	return xxx_messageInfo_ReviewComment.Size(m)
}
func (m *ReviewComment) XXX_DiscardUnknown() {
	xxx_messageInfo_ReviewComment.DiscardUnknown(m)
}

var xxx_messageInfo_ReviewComment proto.InternalMessageInfo

func (m *ReviewComment) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReviewComment) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ReviewComment) GetItemReview() *ItemReview {
	if m != nil {
		return m.ItemReview
	}
	return nil
}

func (m *ReviewComment) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *ReviewComment) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *ReviewComment) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*ReviewComment)(nil), "aggregates.store.ReviewComment")
}

func init() {
	proto.RegisterFile("aggregates/store/review_comment.proto", fileDescriptor_6f9167cebc150470)
}

var fileDescriptor_6f9167cebc150470 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4f, 0x02, 0x31,
	0x10, 0x85, 0xb3, 0x80, 0x28, 0x25, 0x1a, 0xd3, 0xd3, 0x06, 0x4d, 0x24, 0x24, 0x26, 0x5c, 0x68,
	0x13, 0x3d, 0x79, 0xf0, 0xa0, 0x5e, 0xd4, 0xe3, 0x46, 0x2f, 0x5e, 0x36, 0x65, 0x3b, 0x94, 0x09,
	0x94, 0x92, 0x76, 0x8a, 0x7f, 0xd8, 0x1f, 0x62, 0x68, 0x41, 0x0d, 0xde, 0xa6, 0x33, 0xaf, 0xdf,
	0x9b, 0x79, 0xec, 0x5a, 0x19, 0xe3, 0xc1, 0x28, 0x82, 0x20, 0x03, 0x39, 0x0f, 0xd2, 0xc3, 0x06,
	0xe1, 0xb3, 0x6e, 0x9c, 0xb5, 0xb0, 0x22, 0xb1, 0xf6, 0x8e, 0x1c, 0x3f, 0xff, 0x95, 0x89, 0x24,
	0x1b, 0x5c, 0x19, 0xe7, 0xcc, 0x12, 0x64, 0x9a, 0x4f, 0xe3, 0x4c, 0x12, 0x5a, 0x08, 0xa4, 0xec,
	0x3a, 0x7f, 0x19, 0x8c, 0xfe, 0x91, 0x91, 0xc0, 0xd6, 0x19, 0x9f, 0x35, 0xa3, 0xaf, 0x82, 0x9d,
	0x56, 0xa9, 0xf1, 0x94, 0xed, 0xf8, 0x19, 0x6b, 0xa1, 0x2e, 0x8b, 0x61, 0x31, 0x6e, 0x57, 0x2d,
	0xd4, 0x9c, 0xb3, 0x4e, 0x8c, 0xa8, 0xcb, 0xd6, 0xb0, 0x18, 0xf7, 0xaa, 0x54, 0xf3, 0x7b, 0xd6,
	0xff, 0x83, 0x2a, 0xdb, 0xc3, 0x62, 0xdc, 0xbf, 0xb9, 0x14, 0x87, 0x2b, 0x8a, 0x17, 0x02, 0x9b,
	0xe9, 0x15, 0xc3, 0x9f, 0x9a, 0x5f, 0xb0, 0x5e, 0x0c, 0xe0, 0xeb, 0xc4, 0xed, 0x24, 0xee, 0xc9,
	0xb6, 0xf1, 0xbe, 0x65, 0x97, 0xec, 0x78, 0x77, 0x79, 0x79, 0x94, 0x46, 0xfb, 0x27, 0xbf, 0x63,
	0xac, 0xf1, 0xa0, 0x08, 0x74, 0xad, 0xa8, 0xec, 0x26, 0xd3, 0x81, 0xc8, 0x29, 0x88, 0x7d, 0x0a,
	0xe2, 0x6d, 0x9f, 0x42, 0xd5, 0xdb, 0xa9, 0x1f, 0xe8, 0xf1, 0xf5, 0xe3, 0xd9, 0x20, 0xcd, 0xe3,
	0x54, 0x34, 0xce, 0x4a, 0x5c, 0xce, 0x95, 0xb5, 0x73, 0xad, 0xe5, 0x22, 0x6a, 0xb5, 0xc0, 0x49,
	0xda, 0x76, 0xe6, 0xdd, 0x8a, 0x26, 0x01, 0xfc, 0x06, 0x1b, 0x90, 0xb0, 0x22, 0x24, 0x84, 0x20,
	0x0f, 0x43, 0x9c, 0x76, 0x93, 0xd5, 0xed, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x20, 0x84,
	0xfa, 0xb9, 0x01, 0x00, 0x00,
}
