// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/store/item.proto

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

type PriceDuration int32

const (
	PriceDuration_DAY   PriceDuration = 0
	PriceDuration_WEEK  PriceDuration = 1
	PriceDuration_MONTH PriceDuration = 2
	PriceDuration_YEAR  PriceDuration = 3
)

var PriceDuration_name = map[int32]string{
	0: "DAY",
	1: "WEEK",
	2: "MONTH",
	3: "YEAR",
}

var PriceDuration_value = map[string]int32{
	"DAY":   0,
	"WEEK":  1,
	"MONTH": 2,
	"YEAR":  3,
}

func (x PriceDuration) String() string {
	return proto.EnumName(PriceDuration_name, int32(x))
}

func (PriceDuration) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f845c3cdd9c1555f, []int{0}
}

type UnitofMeasurement int32

const (
	UnitofMeasurement_MM  UnitofMeasurement = 0
	UnitofMeasurement_CM  UnitofMeasurement = 1
	UnitofMeasurement_DM  UnitofMeasurement = 2
	UnitofMeasurement_M   UnitofMeasurement = 3
	UnitofMeasurement_DAM UnitofMeasurement = 4
	UnitofMeasurement_HM  UnitofMeasurement = 5
	UnitofMeasurement_KM  UnitofMeasurement = 6
)

var UnitofMeasurement_name = map[int32]string{
	0: "MM",
	1: "CM",
	2: "DM",
	3: "M",
	4: "DAM",
	5: "HM",
	6: "KM",
}

var UnitofMeasurement_value = map[string]int32{
	"MM":  0,
	"CM":  1,
	"DM":  2,
	"M":   3,
	"DAM": 4,
	"HM":  5,
	"KM":  6,
}

func (x UnitofMeasurement) String() string {
	return proto.EnumName(UnitofMeasurement_name, int32(x))
}

func (UnitofMeasurement) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f845c3cdd9c1555f, []int{1}
}

type ItemDimension struct {
	Length               int32             `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Width                int32             `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32             `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	UnitOfMeasurement    UnitofMeasurement `protobuf:"varint,4,opt,name=unit_of_measurement,json=unitOfMeasurement,proto3,enum=aggregates.store.UnitofMeasurement" json:"unit_of_measurement,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ItemDimension) Reset()         { *m = ItemDimension{} }
func (m *ItemDimension) String() string { return proto.CompactTextString(m) }
func (*ItemDimension) ProtoMessage()    {}
func (*ItemDimension) Descriptor() ([]byte, []int) {
	return fileDescriptor_f845c3cdd9c1555f, []int{0}
}

func (m *ItemDimension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemDimension.Unmarshal(m, b)
}
func (m *ItemDimension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemDimension.Marshal(b, m, deterministic)
}
func (m *ItemDimension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemDimension.Merge(m, src)
}
func (m *ItemDimension) XXX_Size() int {
	return xxx_messageInfo_ItemDimension.Size(m)
}
func (m *ItemDimension) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemDimension.DiscardUnknown(m)
}

var xxx_messageInfo_ItemDimension proto.InternalMessageInfo

func (m *ItemDimension) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *ItemDimension) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ItemDimension) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ItemDimension) GetUnitOfMeasurement() UnitofMeasurement {
	if m != nil {
		return m.UnitOfMeasurement
	}
	return UnitofMeasurement_MM
}

type Item struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Storefront           *Storefront          `protobuf:"bytes,3,opt,name=storefront,proto3" json:"storefront,omitempty"`
	Name                 string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Amount               int32                `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Unit                 string               `protobuf:"bytes,6,opt,name=unit,proto3" json:"unit,omitempty"`
	Price                int32                `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty"`
	PriceDuration        PriceDuration        `protobuf:"varint,8,opt,name=price_duration,json=priceDuration,proto3,enum=aggregates.store.PriceDuration" json:"price_duration,omitempty"`
	Description          string               `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	Photo                string               `protobuf:"bytes,10,opt,name=photo,proto3" json:"photo,omitempty"`
	Rating               float64              `protobuf:"fixed64,11,opt,name=rating,proto3" json:"rating,omitempty"`
	ItemDimension        *ItemDimension       `protobuf:"bytes,12,opt,name=item_dimension,json=itemDimension,proto3" json:"item_dimension,omitempty"`
	Color                string               `protobuf:"bytes,13,opt,name=color,proto3" json:"color,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_f845c3cdd9c1555f, []int{1}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Item) GetStorefront() *Storefront {
	if m != nil {
		return m.Storefront
	}
	return nil
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Item) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *Item) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Item) GetPriceDuration() PriceDuration {
	if m != nil {
		return m.PriceDuration
	}
	return PriceDuration_DAY
}

func (m *Item) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Item) GetPhoto() string {
	if m != nil {
		return m.Photo
	}
	return ""
}

func (m *Item) GetRating() float64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Item) GetItemDimension() *ItemDimension {
	if m != nil {
		return m.ItemDimension
	}
	return nil
}

func (m *Item) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *Item) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("aggregates.store.PriceDuration", PriceDuration_name, PriceDuration_value)
	proto.RegisterEnum("aggregates.store.UnitofMeasurement", UnitofMeasurement_name, UnitofMeasurement_value)
	proto.RegisterType((*ItemDimension)(nil), "aggregates.store.ItemDimension")
	proto.RegisterType((*Item)(nil), "aggregates.store.Item")
}

func init() { proto.RegisterFile("aggregates/store/item.proto", fileDescriptor_f845c3cdd9c1555f) }

var fileDescriptor_f845c3cdd9c1555f = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x5e, 0x92, 0xa6, 0x5b, 0x5d, 0x5a, 0x79, 0x66, 0x42, 0x56, 0x41, 0x5a, 0x19, 0x2f, 0xd5,
	0xa4, 0x25, 0xd2, 0x78, 0x40, 0x48, 0xbc, 0x14, 0x5a, 0x54, 0xa8, 0xc2, 0x50, 0x36, 0x84, 0xc6,
	0x4b, 0x95, 0x26, 0x6e, 0x62, 0xad, 0xb6, 0xab, 0xe4, 0x02, 0xff, 0x89, 0x3f, 0xc7, 0x5f, 0x40,
	0xb6, 0x5b, 0xda, 0xae, 0xbc, 0xc4, 0xfe, 0xee, 0x3e, 0xdf, 0x7d, 0x77, 0xfa, 0x82, 0x9e, 0x27,
	0x79, 0x5e, 0xb2, 0x3c, 0x01, 0x56, 0x85, 0x15, 0xa8, 0x92, 0x85, 0x1c, 0x98, 0x08, 0x56, 0xa5,
	0x02, 0x45, 0xf0, 0x36, 0x19, 0x98, 0x64, 0xef, 0x3c, 0x57, 0x2a, 0x5f, 0xb2, 0xd0, 0xe4, 0xe7,
	0xf5, 0x22, 0x04, 0x2e, 0x58, 0x05, 0x89, 0x58, 0xd9, 0x27, 0xbd, 0x97, 0x07, 0xf5, 0xcc, 0x77,
	0x51, 0x2a, 0x09, 0x96, 0x72, 0xf1, 0xdb, 0x41, 0x9d, 0x4f, 0xc0, 0xc4, 0x88, 0x0b, 0x26, 0x2b,
	0xae, 0x24, 0x79, 0x86, 0x9a, 0x4b, 0x26, 0x73, 0x28, 0xa8, 0xd3, 0x77, 0x06, 0x7e, 0xbc, 0x46,
	0xe4, 0x0c, 0xf9, 0xbf, 0x78, 0x06, 0x05, 0x75, 0x4d, 0xd8, 0x02, 0xcd, 0x2e, 0x18, 0xcf, 0x0b,
	0xa0, 0x9e, 0x65, 0x5b, 0x44, 0x6e, 0xd1, 0xd3, 0x5a, 0x72, 0x98, 0xa9, 0xc5, 0x4c, 0xb0, 0xa4,
	0xaa, 0x4b, 0x26, 0x98, 0x04, 0xda, 0xe8, 0x3b, 0x83, 0xee, 0xf5, 0xab, 0xe0, 0xf1, 0x2c, 0xc1,
	0x37, 0xc9, 0x41, 0x2d, 0xa2, 0x2d, 0x35, 0x3e, 0xd5, 0xef, 0x6f, 0x76, 0x43, 0x17, 0x7f, 0x3c,
	0xd4, 0xd0, 0x62, 0x49, 0x17, 0xb9, 0x3c, 0x33, 0xfa, 0xbc, 0xd8, 0xe5, 0x19, 0x21, 0xa8, 0x51,
	0xd7, 0x3c, 0x33, 0xd2, 0x5a, 0xb1, 0xb9, 0x93, 0x77, 0x08, 0x6d, 0xa7, 0x35, 0xea, 0xda, 0xd7,
	0x2f, 0x0e, 0x1b, 0xdf, 0xfe, 0xe3, 0xc4, 0x3b, 0x7c, 0x5d, 0x51, 0x26, 0x82, 0x19, 0xc1, 0xad,
	0xd8, 0xdc, 0xf5, 0xac, 0x89, 0x50, 0xb5, 0x04, 0xea, 0xdb, 0x59, 0x2d, 0x32, 0xdd, 0x25, 0x07,
	0xda, 0x5c, 0x77, 0x97, 0x1c, 0xf4, 0xb6, 0x56, 0x25, 0x4f, 0x19, 0x3d, 0xb6, 0xdb, 0x32, 0x80,
	0x7c, 0x44, 0x5d, 0x73, 0x99, 0x65, 0x75, 0x99, 0x00, 0x57, 0x92, 0x9e, 0x98, 0x85, 0x9c, 0x1f,
	0xea, 0xfa, 0xaa, 0x79, 0xa3, 0x35, 0x2d, 0xee, 0xac, 0x76, 0x21, 0xe9, 0xa3, 0x76, 0xc6, 0xaa,
	0xb4, 0xe4, 0x2b, 0x53, 0xa4, 0x65, 0x1a, 0xef, 0x86, 0x4c, 0xff, 0x42, 0x81, 0xa2, 0xc8, 0xe4,
	0x2c, 0xd0, 0x13, 0xe8, 0x0a, 0x32, 0xa7, 0xed, 0xbe, 0x33, 0x70, 0xe2, 0x35, 0xd2, 0xba, 0xb4,
	0xd3, 0x66, 0xd9, 0xc6, 0x05, 0xf4, 0x89, 0xd9, 0xd7, 0x7f, 0x74, 0xed, 0x99, 0x25, 0xee, 0xf0,
	0x3d, 0xef, 0x9c, 0x21, 0x3f, 0x55, 0x4b, 0x55, 0xd2, 0x8e, 0xed, 0x6a, 0x00, 0x79, 0x8b, 0x50,
	0x5a, 0xb2, 0x04, 0x58, 0x36, 0x4b, 0x80, 0x76, 0x4d, 0xe5, 0x5e, 0x60, 0xcd, 0x1b, 0x6c, 0xcc,
	0x1b, 0xdc, 0x6d, 0xcc, 0x1b, 0xb7, 0xd6, 0xec, 0x21, 0x5c, 0xbe, 0x41, 0x9d, 0xbd, 0x45, 0x90,
	0x63, 0xe4, 0x8d, 0x86, 0xf7, 0xf8, 0x88, 0x9c, 0xa0, 0xc6, 0xf7, 0xf1, 0x78, 0x8a, 0x1d, 0xd2,
	0x42, 0x7e, 0x74, 0xf3, 0xe5, 0x6e, 0x82, 0x5d, 0x1d, 0xbc, 0x1f, 0x0f, 0x63, 0xec, 0x5d, 0x4e,
	0xd1, 0xe9, 0x81, 0xa5, 0x48, 0x13, 0xb9, 0x51, 0x84, 0x8f, 0xf4, 0xf9, 0x21, 0xc2, 0x8e, 0x3e,
	0x47, 0x11, 0x76, 0x89, 0x8f, 0x9c, 0x08, 0x7b, 0xb6, 0x76, 0x84, 0x1b, 0x3a, 0x3e, 0x89, 0xb0,
	0xaf, 0xcf, 0x69, 0x84, 0x9b, 0xef, 0x3f, 0xff, 0x98, 0xe4, 0x1c, 0x8a, 0x7a, 0x1e, 0xa4, 0x4a,
	0x84, 0x7c, 0x59, 0x24, 0x42, 0x14, 0x59, 0x16, 0x3e, 0xd4, 0x59, 0xf2, 0xc0, 0xaf, 0xb6, 0xb6,
	0xb9, 0xaa, 0x58, 0xf9, 0x93, 0xa7, 0x2c, 0x64, 0x12, 0x38, 0x70, 0x56, 0x85, 0x8f, 0xff, 0xc0,
	0x79, 0xd3, 0x0c, 0xfc, 0xfa, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xc0, 0x06, 0x2e, 0xec,
	0x03, 0x00, 0x00,
}
