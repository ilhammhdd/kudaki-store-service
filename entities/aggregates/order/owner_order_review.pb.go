// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/order/owner_order_review.proto

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

type OwnerOrderReview struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	TenantUuid           string               `protobuf:"bytes,3,opt,name=tenant_uuid,json=tenantUuid,proto3" json:"tenant_uuid,omitempty"`
	OwnerOrder           *OwnerOrder          `protobuf:"bytes,4,opt,name=owner_order,json=ownerOrder,proto3" json:"owner_order,omitempty"`
	Rating               float64              `protobuf:"fixed64,5,opt,name=rating,proto3" json:"rating,omitempty"`
	Review               string               `protobuf:"bytes,6,opt,name=review,proto3" json:"review,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OwnerOrderReview) Reset()         { *m = OwnerOrderReview{} }
func (m *OwnerOrderReview) String() string { return proto.CompactTextString(m) }
func (*OwnerOrderReview) ProtoMessage()    {}
func (*OwnerOrderReview) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e0316944777b9fa, []int{0}
}

func (m *OwnerOrderReview) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OwnerOrderReview.Unmarshal(m, b)
}
func (m *OwnerOrderReview) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OwnerOrderReview.Marshal(b, m, deterministic)
}
func (m *OwnerOrderReview) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnerOrderReview.Merge(m, src)
}
func (m *OwnerOrderReview) XXX_Size() int {
	return xxx_messageInfo_OwnerOrderReview.Size(m)
}
func (m *OwnerOrderReview) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnerOrderReview.DiscardUnknown(m)
}

var xxx_messageInfo_OwnerOrderReview proto.InternalMessageInfo

func (m *OwnerOrderReview) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OwnerOrderReview) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *OwnerOrderReview) GetTenantUuid() string {
	if m != nil {
		return m.TenantUuid
	}
	return ""
}

func (m *OwnerOrderReview) GetOwnerOrder() *OwnerOrder {
	if m != nil {
		return m.OwnerOrder
	}
	return nil
}

func (m *OwnerOrderReview) GetRating() float64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *OwnerOrderReview) GetReview() string {
	if m != nil {
		return m.Review
	}
	return ""
}

func (m *OwnerOrderReview) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*OwnerOrderReview)(nil), "aggregates.order.OwnerOrderReview")
}

func init() {
	proto.RegisterFile("aggregates/order/owner_order_review.proto", fileDescriptor_3e0316944777b9fa)
}

var fileDescriptor_3e0316944777b9fa = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x6f, 0xc2, 0x30,
	0x10, 0x85, 0x95, 0x40, 0xa9, 0x30, 0x52, 0x85, 0x3c, 0x54, 0x11, 0xaa, 0x44, 0xc4, 0x94, 0x0e,
	0xd8, 0x52, 0x3b, 0x75, 0xe8, 0xd0, 0x4e, 0x55, 0x17, 0xa4, 0xa8, 0x5d, 0xba, 0x44, 0x06, 0x1f,
	0xe6, 0x04, 0x89, 0x91, 0x73, 0x81, 0xb9, 0xff, 0xbc, 0xe2, 0x02, 0x2d, 0x62, 0xe8, 0x76, 0xf7,
	0xde, 0xb3, 0x3f, 0xfb, 0x89, 0x7b, 0xe3, 0x5c, 0x00, 0x67, 0x08, 0x6a, 0xed, 0x83, 0x85, 0xa0,
	0xfd, 0xbe, 0x82, 0x50, 0xf0, 0x5c, 0x04, 0xd8, 0x21, 0xec, 0xd5, 0x36, 0x78, 0xf2, 0x72, 0xf8,
	0x17, 0x55, 0x6c, 0x8f, 0xc6, 0xce, 0x7b, 0xb7, 0x01, 0xcd, 0xfe, 0xbc, 0x59, 0x6a, 0xc2, 0x12,
	0x6a, 0x32, 0xe5, 0xb6, 0x3d, 0x32, 0x9a, 0xfc, 0x77, 0x7b, 0x9b, 0x99, 0x7c, 0xc7, 0x62, 0x38,
	0x3b, 0xa8, 0xb3, 0x83, 0x98, 0x33, 0x51, 0xde, 0x88, 0x18, 0x6d, 0x12, 0xa5, 0x51, 0xd6, 0xc9,
	0x63, 0xb4, 0x52, 0x8a, 0x6e, 0xd3, 0xa0, 0x4d, 0xe2, 0x34, 0xca, 0xfa, 0x39, 0xcf, 0x72, 0x2c,
	0x06, 0x04, 0x95, 0xa9, 0xa8, 0x60, 0xab, 0xc3, 0x96, 0x68, 0xa5, 0xcf, 0x43, 0xe0, 0x59, 0x0c,
	0xce, 0x70, 0x49, 0x37, 0x8d, 0xb2, 0xc1, 0xc3, 0x9d, 0xba, 0xfc, 0x86, 0x3a, 0xa3, 0x0b, 0xff,
	0x3b, 0xcb, 0x5b, 0xd1, 0x0b, 0x86, 0xb0, 0x72, 0xc9, 0x55, 0x1a, 0x65, 0x51, 0x7e, 0xdc, 0x58,
	0xe7, 0x57, 0x26, 0x3d, 0x46, 0x1e, 0x37, 0xf9, 0x24, 0xc4, 0x22, 0x80, 0x21, 0xb0, 0x85, 0xa1,
	0xe4, 0x9a, 0x69, 0x23, 0xd5, 0x56, 0xa4, 0x4e, 0x15, 0xa9, 0x8f, 0x53, 0x45, 0x79, 0xff, 0x98,
	0x7e, 0xa1, 0xd7, 0xf7, 0xaf, 0x37, 0x87, 0xb4, 0x6a, 0xe6, 0x6a, 0xe1, 0x4b, 0x8d, 0x9b, 0x95,
	0x29, 0xcb, 0x95, 0xb5, 0x7a, 0xdd, 0x58, 0xb3, 0xc6, 0x69, 0x4d, 0x3e, 0xc0, 0x32, 0xf8, 0x8a,
	0xa6, 0x35, 0x84, 0x1d, 0x2e, 0x40, 0x43, 0x45, 0x48, 0x08, 0xb5, 0xbe, 0x6c, 0x78, 0xde, 0x63,
	0xd4, 0xe3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xdc, 0x18, 0xbd, 0xda, 0x01, 0x00, 0x00,
}
