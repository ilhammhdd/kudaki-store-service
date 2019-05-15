// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/topics.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type UserTopic int32

const (
	UserTopic_SIGNED_UP                     UserTopic = 0
	UserTopic_USER_VERIFICATION_EMAIL_SENT  UserTopic = 1
	UserTopic_SIGN_UP_REQUESTED             UserTopic = 2
	UserTopic_VERIFY_USER_REQUESTED         UserTopic = 3
	UserTopic_LOGGED_IN                     UserTopic = 4
	UserTopic_LOGIN_REQUESTED               UserTopic = 5
	UserTopic_USER_AUTHENTICATION_REQUESTED UserTopic = 6
	UserTopic_USER_AUTHENTICATED            UserTopic = 7
	UserTopic_RESET_PASSWORD_REQUESTED      UserTopic = 8
	UserTopic_PASSWORD_RESETED              UserTopic = 9
	UserTopic_USER_AUTHORIZATION_REQUESTED  UserTopic = 10
	UserTopic_USER_AUTHORIZED               UserTopic = 11
	UserTopic_USER_VERIFIED                 UserTopic = 12
	UserTopic_USER_RETRIEVED                UserTopic = 13
)

var UserTopic_name = map[int32]string{
	0:  "SIGNED_UP",
	1:  "USER_VERIFICATION_EMAIL_SENT",
	2:  "SIGN_UP_REQUESTED",
	3:  "VERIFY_USER_REQUESTED",
	4:  "LOGGED_IN",
	5:  "LOGIN_REQUESTED",
	6:  "USER_AUTHENTICATION_REQUESTED",
	7:  "USER_AUTHENTICATED",
	8:  "RESET_PASSWORD_REQUESTED",
	9:  "PASSWORD_RESETED",
	10: "USER_AUTHORIZATION_REQUESTED",
	11: "USER_AUTHORIZED",
	12: "USER_VERIFIED",
	13: "USER_RETRIEVED",
}

var UserTopic_value = map[string]int32{
	"SIGNED_UP":                     0,
	"USER_VERIFICATION_EMAIL_SENT":  1,
	"SIGN_UP_REQUESTED":             2,
	"VERIFY_USER_REQUESTED":         3,
	"LOGGED_IN":                     4,
	"LOGIN_REQUESTED":               5,
	"USER_AUTHENTICATION_REQUESTED": 6,
	"USER_AUTHENTICATED":            7,
	"RESET_PASSWORD_REQUESTED":      8,
	"PASSWORD_RESETED":              9,
	"USER_AUTHORIZATION_REQUESTED":  10,
	"USER_AUTHORIZED":               11,
	"USER_VERIFIED":                 12,
	"USER_RETRIEVED":                13,
}

func (x UserTopic) String() string {
	return proto.EnumName(UserTopic_name, int32(x))
}

func (UserTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c97f2476c050d00c, []int{0}
}

type MountainTopic int32

const (
	MountainTopic_CREATE_MOUNTAIN_REQUESTED    MountainTopic = 0
	MountainTopic_MOUNTAIN_CREATED             MountainTopic = 1
	MountainTopic_RETRIEVE_MOUNTAINS_REQUESTED MountainTopic = 2
	MountainTopic_MOUNTAINS_RETRIEVED          MountainTopic = 3
)

var MountainTopic_name = map[int32]string{
	0: "CREATE_MOUNTAIN_REQUESTED",
	1: "MOUNTAIN_CREATED",
	2: "RETRIEVE_MOUNTAINS_REQUESTED",
	3: "MOUNTAINS_RETRIEVED",
}

var MountainTopic_value = map[string]int32{
	"CREATE_MOUNTAIN_REQUESTED":    0,
	"MOUNTAIN_CREATED":             1,
	"RETRIEVE_MOUNTAINS_REQUESTED": 2,
	"MOUNTAINS_RETRIEVED":          3,
}

func (x MountainTopic) String() string {
	return proto.EnumName(MountainTopic_name, int32(x))
}

func (MountainTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c97f2476c050d00c, []int{1}
}

type StoreTopic int32

const (
	StoreTopic_ADD_STOREFRONT_ITEM_REQUESTED       StoreTopic = 0
	StoreTopic_STOREFRONT_ITEM_ADDED               StoreTopic = 1
	StoreTopic_DELETE_STOREFRONT_ITEM_REQUESTED    StoreTopic = 2
	StoreTopic_STOREFRONT_ITEM_DELETED             StoreTopic = 3
	StoreTopic_RETRIEVE_STOREFRONT_ITEMS_REQUESTED StoreTopic = 4
	StoreTopic_STOREFRONT_ITEMS_RETREIVED          StoreTopic = 5
)

var StoreTopic_name = map[int32]string{
	0: "ADD_STOREFRONT_ITEM_REQUESTED",
	1: "STOREFRONT_ITEM_ADDED",
	2: "DELETE_STOREFRONT_ITEM_REQUESTED",
	3: "STOREFRONT_ITEM_DELETED",
	4: "RETRIEVE_STOREFRONT_ITEMS_REQUESTED",
	5: "STOREFRONT_ITEMS_RETREIVED",
}

var StoreTopic_value = map[string]int32{
	"ADD_STOREFRONT_ITEM_REQUESTED":       0,
	"STOREFRONT_ITEM_ADDED":               1,
	"DELETE_STOREFRONT_ITEM_REQUESTED":    2,
	"STOREFRONT_ITEM_DELETED":             3,
	"RETRIEVE_STOREFRONT_ITEMS_REQUESTED": 4,
	"STOREFRONT_ITEMS_RETREIVED":          5,
}

func (x StoreTopic) String() string {
	return proto.EnumName(StoreTopic_name, int32(x))
}

func (StoreTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c97f2476c050d00c, []int{2}
}

func init() {
	proto.RegisterEnum("event.topic.UserTopic", UserTopic_name, UserTopic_value)
	proto.RegisterEnum("event.topic.MountainTopic", MountainTopic_name, MountainTopic_value)
	proto.RegisterEnum("event.topic.StoreTopic", StoreTopic_name, StoreTopic_value)
}

func init() { proto.RegisterFile("events/topics.proto", fileDescriptor_c97f2476c050d00c) }

var fileDescriptor_c97f2476c050d00c = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xdd, 0x6e, 0x12, 0x41,
	0x14, 0x2e, 0xf4, 0x47, 0x39, 0x15, 0x9d, 0x1e, 0xc4, 0x8a, 0x16, 0xd3, 0x46, 0x13, 0x13, 0x4c,
	0xe1, 0xc2, 0x27, 0x58, 0x3b, 0xa7, 0x38, 0x09, 0xec, 0xe2, 0xcc, 0x6c, 0x8d, 0xbd, 0x99, 0xd0,
	0xb2, 0x91, 0x4d, 0x85, 0x6d, 0x60, 0xf0, 0xd6, 0xb7, 0xf4, 0xd6, 0x57, 0x31, 0xb3, 0x2b, 0xb0,
	0xbb, 0x31, 0xbd, 0xdc, 0xef, 0xe7, 0x9c, 0xef, 0x3b, 0x99, 0x85, 0x46, 0xf4, 0x33, 0x9a, 0xdb,
	0x65, 0xcf, 0x26, 0xf7, 0xf1, 0xed, 0xb2, 0x7b, 0xbf, 0x48, 0x6c, 0x82, 0x87, 0x29, 0xd8, 0x4d,
	0xb1, 0xce, 0x9f, 0x2a, 0xd4, 0xc2, 0x65, 0xb4, 0xd0, 0xee, 0x0b, 0xeb, 0x50, 0x53, 0xa2, 0xef,
	0x13, 0x37, 0xe1, 0x88, 0xed, 0xe0, 0x29, 0x9c, 0x84, 0x8a, 0xa4, 0xb9, 0x22, 0x29, 0x2e, 0xc5,
	0x85, 0xa7, 0x45, 0xe0, 0x1b, 0x1a, 0x7a, 0x62, 0x60, 0x14, 0xf9, 0x9a, 0x55, 0xb0, 0x09, 0x47,
	0xce, 0x60, 0xc2, 0x91, 0x91, 0xf4, 0x25, 0x24, 0xa5, 0x89, 0xb3, 0x2a, 0xb6, 0xa0, 0x99, 0x7a,
	0xbe, 0x99, 0xd4, 0xbf, 0xa5, 0x76, 0xdd, 0x8a, 0x41, 0xd0, 0xef, 0x13, 0x37, 0xc2, 0x67, 0x7b,
	0xd8, 0x80, 0x67, 0x83, 0xa0, 0x2f, 0xfc, 0x9c, 0x66, 0x1f, 0xcf, 0xa0, 0x9d, 0xfa, 0xbc, 0x50,
	0x7f, 0x26, 0x5f, 0xaf, 0x37, 0x6f, 0x25, 0x07, 0xf8, 0x02, 0xb0, 0x2c, 0x21, 0xce, 0x1e, 0xe1,
	0x09, 0xbc, 0x94, 0xa4, 0x48, 0x9b, 0x91, 0xa7, 0xd4, 0xd7, 0x40, 0xf2, 0x9c, 0xeb, 0x31, 0x3e,
	0x07, 0x96, 0xc3, 0x15, 0x39, 0xb4, 0xb6, 0xa9, 0xe9, 0x66, 0x05, 0x52, 0x5c, 0x97, 0xb7, 0x81,
	0x4b, 0x59, 0x50, 0x10, 0x67, 0x87, 0x78, 0x04, 0xf5, 0xdc, 0x75, 0x88, 0xb3, 0x27, 0x88, 0xf0,
	0xf4, 0x5f, 0x61, 0x2d, 0x05, 0x5d, 0x11, 0x67, 0xf5, 0xce, 0x2f, 0xa8, 0x0f, 0x93, 0xd5, 0xdc,
	0x8e, 0xe3, 0x79, 0x76, 0xe4, 0x36, 0xb4, 0x2e, 0x24, 0x79, 0x9a, 0xcc, 0x30, 0x08, 0x7d, 0xed,
	0x15, 0xca, 0xef, 0xb8, 0x8c, 0x1b, 0x3c, 0xd3, 0x71, 0x56, 0x71, 0x19, 0xd7, 0x43, 0x37, 0x36,
	0x55, 0xb8, 0xf9, 0x31, 0x34, 0xf2, 0xc4, 0x3a, 0xc0, 0x6e, 0xe7, 0x77, 0x05, 0x40, 0xd9, 0x64,
	0x11, 0x65, 0xeb, 0xcf, 0xa0, 0xed, 0x71, 0x6e, 0x94, 0x0e, 0x24, 0x5d, 0xca, 0xc0, 0xd7, 0x46,
	0x68, 0x1a, 0x16, 0x22, 0xb4, 0xa0, 0x59, 0xa6, 0x3d, 0xce, 0xd3, 0x1c, 0xef, 0xe0, 0x94, 0xd3,
	0x80, 0x34, 0x3d, 0x30, 0xa0, 0x8a, 0xaf, 0xe1, 0xb8, 0x4c, 0x67, 0x2e, 0xf7, 0x02, 0xde, 0xc3,
	0xdb, 0x4d, 0x95, 0x92, 0x2a, 0xdf, 0x68, 0x0f, 0xdf, 0xc0, 0xab, 0xff, 0xf0, 0x5a, 0x92, 0x70,
	0xc5, 0xf6, 0x3f, 0x9d, 0x5f, 0x7f, 0xf8, 0x1e, 0xdb, 0xe9, 0xea, 0xa6, 0x7b, 0x9b, 0xcc, 0x7a,
	0xf1, 0x8f, 0xe9, 0x78, 0x36, 0x9b, 0x4e, 0x26, 0xbd, 0xbb, 0xd5, 0x64, 0x7c, 0x17, 0x9f, 0x47,
	0x73, 0x1b, 0xdb, 0x38, 0x5a, 0xf6, 0xb2, 0x9f, 0xe0, 0xe6, 0x20, 0x7d, 0xfe, 0x1f, 0xff, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x38, 0xc9, 0xd7, 0x78, 0x15, 0x03, 0x00, 0x00,
}
