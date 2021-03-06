// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/store/item_review.proto

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

type ItemReview struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid             string               `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Item                 *Item                `protobuf:"bytes,4,opt,name=item,proto3" json:"item,omitempty"`
	Review               string               `protobuf:"bytes,5,opt,name=review,proto3" json:"review,omitempty"`
	Rating               float64              `protobuf:"fixed64,6,opt,name=rating,proto3" json:"rating,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ItemReview) Reset()         { *m = ItemReview{} }
func (m *ItemReview) String() string { return proto.CompactTextString(m) }
func (*ItemReview) ProtoMessage()    {}
func (*ItemReview) Descriptor() ([]byte, []int) {
	return fileDescriptor_3806212aa403f188, []int{0}
}

func (m *ItemReview) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemReview.Unmarshal(m, b)
}
func (m *ItemReview) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemReview.Marshal(b, m, deterministic)
}
func (m *ItemReview) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemReview.Merge(m, src)
}
func (m *ItemReview) XXX_Size() int {
	return xxx_messageInfo_ItemReview.Size(m)
}
func (m *ItemReview) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemReview.DiscardUnknown(m)
}

var xxx_messageInfo_ItemReview proto.InternalMessageInfo

func (m *ItemReview) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ItemReview) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ItemReview) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *ItemReview) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *ItemReview) GetReview() string {
	if m != nil {
		return m.Review
	}
	return ""
}

func (m *ItemReview) GetRating() float64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *ItemReview) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*ItemReview)(nil), "aggregates.store.ItemReview")
}

func init() { proto.RegisterFile("aggregates/store/item_review.proto", fileDescriptor_3806212aa403f188) }

var fileDescriptor_3806212aa403f188 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4f, 0x02, 0x31,
	0x10, 0x85, 0x53, 0x40, 0x94, 0x9a, 0x18, 0xd3, 0x03, 0xd9, 0xc0, 0xc1, 0x0d, 0xa7, 0x8d, 0x09,
	0x6d, 0xa2, 0x27, 0x8f, 0x7a, 0x52, 0x8f, 0x1b, 0xbd, 0x78, 0x21, 0x85, 0x0e, 0x65, 0x02, 0xa5,
	0xa4, 0x3b, 0xc5, 0x3f, 0xec, 0x0f, 0x31, 0xcc, 0x42, 0x4c, 0x88, 0xb7, 0x9d, 0xf9, 0xde, 0xbe,
	0x37, 0xaf, 0x72, 0x62, 0xbd, 0x4f, 0xe0, 0x2d, 0x41, 0x63, 0x1a, 0x8a, 0x09, 0x0c, 0x12, 0x84,
	0x59, 0x82, 0x3d, 0xc2, 0xb7, 0xde, 0xa5, 0x48, 0x51, 0xdd, 0xfe, 0x69, 0x34, 0x6b, 0x46, 0x77,
	0x3e, 0x46, 0xbf, 0x01, 0xc3, 0x7c, 0x9e, 0x97, 0x86, 0x30, 0x40, 0x43, 0x36, 0xec, 0xda, 0x5f,
	0x46, 0xe3, 0x7f, 0x6d, 0x5b, 0x38, 0xf9, 0x11, 0x52, 0xbe, 0x11, 0x84, 0x9a, 0x43, 0xd4, 0x8d,
	0xec, 0xa0, 0x2b, 0x44, 0x29, 0xaa, 0x6e, 0xdd, 0x41, 0xa7, 0x94, 0xec, 0xe5, 0x8c, 0xae, 0xe8,
	0x94, 0xa2, 0x1a, 0xd4, 0xfc, 0xad, 0xc6, 0x72, 0x90, 0x1b, 0x48, 0x33, 0x06, 0x5d, 0x06, 0x57,
	0x87, 0xc5, 0xe7, 0x01, 0xde, 0xcb, 0xde, 0xc1, 0xbd, 0xe8, 0x95, 0xa2, 0xba, 0x7e, 0x18, 0xea,
	0xf3, 0x73, 0x35, 0x87, 0xb1, 0x46, 0x0d, 0x65, 0xbf, 0xed, 0x56, 0x5c, 0xb0, 0xcb, 0x71, 0xe2,
	0xbd, 0x25, 0xdc, 0xfa, 0xa2, 0x5f, 0x8a, 0x4a, 0xd4, 0xc7, 0x49, 0x3d, 0x49, 0xb9, 0x48, 0x60,
	0x09, 0xdc, 0xcc, 0x52, 0x71, 0xc9, 0x09, 0x23, 0xdd, 0xd6, 0xd7, 0xa7, 0xfa, 0xfa, 0xe3, 0x54,
	0xbf, 0x1e, 0x1c, 0xd5, 0xcf, 0xf4, 0xf2, 0xfe, 0xf5, 0xea, 0x91, 0x56, 0x79, 0xae, 0x17, 0x31,
	0x18, 0xdc, 0xac, 0x6c, 0x08, 0x2b, 0xe7, 0xcc, 0x3a, 0x3b, 0xbb, 0xc6, 0x29, 0x9f, 0xb6, 0x4c,
	0x71, 0x4b, 0xd3, 0x06, 0xd2, 0x1e, 0x17, 0x60, 0x60, 0x4b, 0x48, 0x08, 0x8d, 0x39, 0x7f, 0xbd,
	0x79, 0x9f, 0xa3, 0x1e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x2d, 0x53, 0x93, 0xaf, 0x01,
	0x00, 0x00,
}
