// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/order/owner_order.proto

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

type OwnerOrder struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Order                *Order               `protobuf:"bytes,3,opt,name=order,proto3" json:"order,omitempty"`
	OwnerUuid            string               `protobuf:"bytes,4,opt,name=owner_uuid,json=ownerUuid,proto3" json:"owner_uuid,omitempty"`
	TotalPrice           int32                `protobuf:"varint,5,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	TotalQuantity        int64                `protobuf:"varint,6,opt,name=total_quantity,json=totalQuantity,proto3" json:"total_quantity,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	OrderStatus          OrderStatus          `protobuf:"varint,8,opt,name=order_status,json=orderStatus,proto3,enum=aggregates.order.OrderStatus" json:"order_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OwnerOrder) Reset()         { *m = OwnerOrder{} }
func (m *OwnerOrder) String() string { return proto.CompactTextString(m) }
func (*OwnerOrder) ProtoMessage()    {}
func (*OwnerOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_45d297f26665e56f, []int{0}
}

func (m *OwnerOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OwnerOrder.Unmarshal(m, b)
}
func (m *OwnerOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OwnerOrder.Marshal(b, m, deterministic)
}
func (m *OwnerOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnerOrder.Merge(m, src)
}
func (m *OwnerOrder) XXX_Size() int {
	return xxx_messageInfo_OwnerOrder.Size(m)
}
func (m *OwnerOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnerOrder.DiscardUnknown(m)
}

var xxx_messageInfo_OwnerOrder proto.InternalMessageInfo

func (m *OwnerOrder) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OwnerOrder) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *OwnerOrder) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *OwnerOrder) GetOwnerUuid() string {
	if m != nil {
		return m.OwnerUuid
	}
	return ""
}

func (m *OwnerOrder) GetTotalPrice() int32 {
	if m != nil {
		return m.TotalPrice
	}
	return 0
}

func (m *OwnerOrder) GetTotalQuantity() int64 {
	if m != nil {
		return m.TotalQuantity
	}
	return 0
}

func (m *OwnerOrder) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *OwnerOrder) GetOrderStatus() OrderStatus {
	if m != nil {
		return m.OrderStatus
	}
	return OrderStatus_PENDING
}

func init() {
	proto.RegisterType((*OwnerOrder)(nil), "aggregates.order.OwnerOrder")
}

func init() { proto.RegisterFile("aggregates/order/owner_order.proto", fileDescriptor_45d297f26665e56f) }

var fileDescriptor_45d297f26665e56f = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4d, 0x4f, 0xe3, 0x30,
	0x10, 0x86, 0x95, 0xf4, 0x63, 0xb7, 0xee, 0x6e, 0xb5, 0xf2, 0x65, 0xa3, 0x8a, 0xaa, 0x51, 0x25,
	0xa4, 0x5c, 0x6a, 0x4b, 0xe5, 0xc4, 0x0d, 0x38, 0x21, 0x2e, 0x85, 0x00, 0x17, 0x2e, 0x91, 0x1b,
	0x4f, 0x53, 0xab, 0x4d, 0x5d, 0x9c, 0x31, 0x88, 0xdf, 0xc7, 0x1f, 0x43, 0x99, 0xb4, 0xaa, 0x54,
	0xc4, 0xcd, 0x79, 0xf2, 0xe6, 0xc9, 0x3b, 0x63, 0x36, 0x51, 0x45, 0xe1, 0xa0, 0x50, 0x08, 0x95,
	0xb4, 0x4e, 0x83, 0x93, 0xf6, 0x7d, 0x0b, 0x2e, 0xa3, 0xb3, 0xd8, 0x39, 0x8b, 0x96, 0xff, 0x3b,
	0x66, 0x04, 0xf1, 0xe1, 0xb8, 0xb0, 0xb6, 0xd8, 0x80, 0xa4, 0xf7, 0x0b, 0xbf, 0x94, 0x68, 0x4a,
	0xa8, 0x50, 0x95, 0xbb, 0xe6, 0x93, 0xe1, 0xd9, 0x77, 0xed, 0x51, 0x38, 0xf9, 0x0c, 0x19, 0x9b,
	0xd7, 0xbf, 0x99, 0xd7, 0x90, 0x0f, 0x58, 0x68, 0x74, 0x14, 0xc4, 0x41, 0xd2, 0x4a, 0x43, 0xa3,
	0x39, 0x67, 0x6d, 0xef, 0x8d, 0x8e, 0xc2, 0x38, 0x48, 0x7a, 0x29, 0x9d, 0xf9, 0x94, 0x75, 0xc8,
	0x10, 0xb5, 0xe2, 0x20, 0xe9, 0xcf, 0xfe, 0x8b, 0xd3, 0x4e, 0x82, 0x5c, 0x69, 0x93, 0xe2, 0x23,
	0xc6, 0x9a, 0x39, 0x48, 0xd4, 0x26, 0x51, 0x8f, 0xc8, 0x73, 0x6d, 0x1b, 0xb3, 0x3e, 0x5a, 0x54,
	0x9b, 0x6c, 0xe7, 0x4c, 0x0e, 0x51, 0x27, 0x0e, 0x92, 0x4e, 0xca, 0x08, 0xdd, 0xd7, 0x84, 0x9f,
	0xb3, 0x41, 0x13, 0x78, 0xf5, 0x6a, 0x8b, 0x06, 0x3f, 0xa2, 0x2e, 0xd5, 0xfb, 0x4b, 0xf4, 0x61,
	0x0f, 0xf9, 0x25, 0x63, 0xb9, 0x03, 0x85, 0xa0, 0x33, 0x85, 0xd1, 0x2f, 0xaa, 0x36, 0x14, 0xcd,
	0x72, 0xc4, 0x61, 0x39, 0xe2, 0xe9, 0xb0, 0x9c, 0xb4, 0xb7, 0x4f, 0x5f, 0x23, 0xbf, 0x62, 0x7f,
	0xa8, 0x6a, 0x56, 0xa1, 0x42, 0x5f, 0x45, 0xbf, 0xe3, 0x20, 0x19, 0xcc, 0x46, 0x3f, 0xcc, 0xf5,
	0x48, 0xa1, 0xb4, 0x6f, 0x8f, 0x0f, 0x37, 0x77, 0x2f, 0xb7, 0x85, 0xc1, 0x95, 0x5f, 0x88, 0xdc,
	0x96, 0xd2, 0x6c, 0x56, 0xaa, 0x2c, 0x57, 0x5a, 0xcb, 0xb5, 0xd7, 0x6a, 0x6d, 0xa6, 0x15, 0x5a,
	0x07, 0x4b, 0x67, 0xb7, 0x38, 0xad, 0xc0, 0xbd, 0x99, 0x1c, 0x24, 0xd4, 0xad, 0x0d, 0x54, 0xf2,
	0xf4, 0x76, 0x16, 0x5d, 0x2a, 0x7b, 0xf1, 0x15, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x35, 0xd6, 0xb2,
	0x0f, 0x02, 0x00, 0x00,
}
