// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/order/item_feedback.proto

package order

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

type ItemFeedback struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Order                *Order               `protobuf:"bytes,3,opt,name=order,proto3" json:"order,omitempty"`
	Rating               float64              `protobuf:"fixed64,4,opt,name=rating,proto3" json:"rating,omitempty"`
	Description          string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ItemFeedback) Reset()         { *m = ItemFeedback{} }
func (m *ItemFeedback) String() string { return proto.CompactTextString(m) }
func (*ItemFeedback) ProtoMessage()    {}
func (*ItemFeedback) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d156fff11ae0264, []int{0}
}

func (m *ItemFeedback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemFeedback.Unmarshal(m, b)
}
func (m *ItemFeedback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemFeedback.Marshal(b, m, deterministic)
}
func (m *ItemFeedback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemFeedback.Merge(m, src)
}
func (m *ItemFeedback) XXX_Size() int {
	return xxx_messageInfo_ItemFeedback.Size(m)
}
func (m *ItemFeedback) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemFeedback.DiscardUnknown(m)
}

var xxx_messageInfo_ItemFeedback proto.InternalMessageInfo

func (m *ItemFeedback) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ItemFeedback) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ItemFeedback) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *ItemFeedback) GetRating() float64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *ItemFeedback) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ItemFeedback) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*ItemFeedback)(nil), "aggregates.order.ItemFeedback")
}

func init() {
	proto.RegisterFile("aggregates/order/item_feedback.proto", fileDescriptor_6d156fff11ae0264)
}

var fileDescriptor_6d156fff11ae0264 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x87, 0x49, 0xff, 0x41, 0x53, 0x11, 0xc9, 0x41, 0x43, 0x11, 0x5c, 0xc4, 0xc3, 0x5e, 0x9a,
	0x80, 0x9e, 0x3c, 0xea, 0x41, 0xd4, 0x8b, 0xb0, 0x78, 0xf2, 0x52, 0xb2, 0x9b, 0x69, 0x3a, 0xb4,
	0xd9, 0x94, 0x64, 0xd6, 0x47, 0xf5, 0x79, 0xc4, 0x6c, 0x8b, 0x52, 0x2f, 0x43, 0x92, 0xf9, 0x32,
	0x5f, 0x7e, 0xe1, 0x37, 0xc6, 0xb9, 0x08, 0xce, 0x10, 0x24, 0x1d, 0xa2, 0x85, 0xa8, 0x91, 0xc0,
	0x2f, 0x57, 0x00, 0xb6, 0x36, 0xcd, 0x46, 0xed, 0x62, 0xa0, 0x20, 0xce, 0x7e, 0x29, 0x95, 0xa9,
	0xf9, 0x95, 0x0b, 0xc1, 0x6d, 0x41, 0xe7, 0x7e, 0xdd, 0xad, 0x34, 0xa1, 0x87, 0x44, 0xc6, 0xef,
	0xfa, 0x2b, 0xf3, 0xcb, 0x7f, 0x83, 0x73, 0xed, 0xbb, 0xd7, 0x5f, 0x8c, 0x9f, 0xbc, 0x10, 0xf8,
	0xa7, 0xbd, 0x47, 0x9c, 0xf2, 0x01, 0x5a, 0xc9, 0x0a, 0x56, 0x0e, 0xab, 0x01, 0x5a, 0x21, 0xf8,
	0xa8, 0xeb, 0xd0, 0xca, 0x41, 0xc1, 0xca, 0x69, 0x95, 0xd7, 0x62, 0xc1, 0xc7, 0x79, 0x86, 0x1c,
	0x16, 0xac, 0x9c, 0xdd, 0x5e, 0xa8, 0xe3, 0x57, 0xa9, 0xb7, 0x9f, 0x5a, 0xf5, 0x94, 0x38, 0xe7,
	0x93, 0x68, 0x08, 0x5b, 0x27, 0x47, 0x05, 0x2b, 0x59, 0xb5, 0xdf, 0x89, 0x82, 0xcf, 0x2c, 0xa4,
	0x26, 0xe2, 0x8e, 0x30, 0xb4, 0x72, 0x9c, 0x0d, 0x7f, 0x8f, 0xc4, 0x3d, 0xe7, 0x4d, 0x04, 0x43,
	0x60, 0x97, 0x86, 0xe4, 0x24, 0xdb, 0xe6, 0xaa, 0x4f, 0xac, 0x0e, 0x89, 0xd5, 0xfb, 0x21, 0x71,
	0x35, 0xdd, 0xd3, 0x0f, 0xf4, 0xf8, 0xfa, 0xf1, 0xec, 0x90, 0xd6, 0x5d, 0xad, 0x9a, 0xe0, 0x35,
	0x6e, 0xd7, 0xc6, 0xfb, 0xb5, 0xb5, 0x7a, 0xd3, 0x59, 0xb3, 0xc1, 0x45, 0xa2, 0x10, 0x61, 0x15,
	0x43, 0x4b, 0x8b, 0x04, 0xf1, 0x13, 0x1b, 0xd0, 0xd0, 0x12, 0x12, 0x42, 0xd2, 0xc7, 0x1f, 0x56,
	0x4f, 0xb2, 0xea, 0xee, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x69, 0x48, 0xea, 0xa4, 0x01, 0x00,
	0x00,
}