// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/config/pullorder.proto

package config

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PullOrder int32

const (
	PullOrderRandom        PullOrder = 0
	PullOrderAlphabetic    PullOrder = 1
	PullOrderSmallestFirst PullOrder = 2
	PullOrderLargestFirst  PullOrder = 3
	PullOrderOldestFirst   PullOrder = 4
	PullOrderNewestFirst   PullOrder = 5
)

var PullOrder_name = map[int32]string{
	0: "PULL_ORDER_RANDOM",
	1: "PULL_ORDER_ALPHABETIC",
	2: "PULL_ORDER_SMALLEST_FIRST",
	3: "PULL_ORDER_LARGEST_FIRST",
	4: "PULL_ORDER_OLDEST_FIRST",
	5: "PULL_ORDER_NEWEST_FIRST",
}

var PullOrder_value = map[string]int32{
	"PULL_ORDER_RANDOM":         0,
	"PULL_ORDER_ALPHABETIC":     1,
	"PULL_ORDER_SMALLEST_FIRST": 2,
	"PULL_ORDER_LARGEST_FIRST":  3,
	"PULL_ORDER_OLDEST_FIRST":   4,
	"PULL_ORDER_NEWEST_FIRST":   5,
}

func (PullOrder) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2fa3f5222a7755bf, []int{0}
}

func init() {
	proto.RegisterEnum("config.PullOrder", PullOrder_name, PullOrder_value)
}

func init() { proto.RegisterFile("lib/config/pullorder.proto", fileDescriptor_2fa3f5222a7755bf) }

var fileDescriptor_2fa3f5222a7755bf = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd1, 0xbf, 0x4a, 0xc3, 0x40,
	0x1c, 0xc0, 0xf1, 0xa4, 0xd6, 0x82, 0xb7, 0x58, 0x53, 0x6b, 0xdb, 0x1b, 0x8e, 0x80, 0x93, 0x1d,
	0x1a, 0x50, 0x44, 0x1c, 0x53, 0x9b, 0x6a, 0xf1, 0xda, 0x94, 0xa4, 0x22, 0xb8, 0x94, 0x24, 0x4d,
	0xd3, 0xc0, 0x35, 0x17, 0xf2, 0x07, 0xf1, 0x15, 0x32, 0xf9, 0x02, 0x01, 0x07, 0x07, 0x1f, 0xa5,
	0x63, 0xc1, 0xc5, 0xb5, 0xcd, 0x8b, 0x88, 0x29, 0x26, 0x41, 0x70, 0xfb, 0xdd, 0x8f, 0xef, 0xe7,
	0x96, 0x1f, 0x80, 0xc4, 0xd6, 0x05, 0x83, 0x3a, 0x73, 0xdb, 0x12, 0xdc, 0x90, 0x10, 0xea, 0xcd,
	0x4c, 0xaf, 0xe3, 0x7a, 0x34, 0xa0, 0x5c, 0x65, 0xb7, 0x87, 0xa7, 0x9e, 0xe9, 0x52, 0x5f, 0x48,
	0x97, 0x7a, 0x38, 0x17, 0x2c, 0x6a, 0xd1, 0xf4, 0x91, 0x4e, 0xbb, 0xb8, 0xfd, 0x59, 0x02, 0x07,
	0xe3, 0x90, 0x10, 0xf9, 0xe7, 0x03, 0xae, 0x0d, 0x8e, 0xc6, 0x0f, 0x18, 0x4f, 0x65, 0xa5, 0x27,
	0x29, 0x53, 0x45, 0x1c, 0xf5, 0xe4, 0x61, 0x95, 0x81, 0xb5, 0x28, 0xe6, 0x0f, 0xb3, 0x4a, 0xd1,
	0x9c, 0x19, 0x5d, 0x72, 0xe7, 0xa0, 0x5e, 0x68, 0x45, 0x3c, 0xbe, 0x13, 0xbb, 0xd2, 0x64, 0x70,
	0x53, 0x65, 0x61, 0x23, 0x8a, 0xf9, 0x5a, 0xd6, 0x8b, 0xc4, 0x5d, 0x68, 0xba, 0x19, 0xd8, 0x06,
	0x77, 0x0d, 0x5a, 0x05, 0xa3, 0x0e, 0x45, 0x8c, 0x25, 0x75, 0x32, 0xed, 0x0f, 0x14, 0x75, 0x52,
	0x2d, 0x41, 0x18, 0xc5, 0xfc, 0x49, 0xe6, 0xd4, 0xa5, 0x46, 0x88, 0xe9, 0x07, 0x7d, 0xdb, 0xf3,
	0x03, 0xee, 0x0a, 0x34, 0x0b, 0x14, 0x8b, 0xca, 0x6d, 0x2e, 0xf7, 0x60, 0x2b, 0x8a, 0xf9, 0x7a,
	0x26, 0xb1, 0xe6, 0x59, 0x19, 0xbc, 0x04, 0x8d, 0x02, 0x94, 0x71, 0x2f, 0x77, 0x65, 0xd8, 0x8c,
	0x62, 0xfe, 0x38, 0x73, 0x32, 0x99, 0xfd, 0xc3, 0x46, 0xd2, 0x63, 0xce, 0xf6, 0xff, 0xb0, 0x91,
	0xf9, 0xfc, 0xcb, 0x60, 0xf9, 0xe3, 0x1d, 0x31, 0xdd, 0xfb, 0xd5, 0x06, 0x31, 0xeb, 0x0d, 0x62,
	0x56, 0x5b, 0xc4, 0xae, 0xb7, 0x88, 0x7d, 0x4d, 0x10, 0xf3, 0x96, 0x20, 0x76, 0x9d, 0x20, 0xe6,
	0x2b, 0x41, 0xcc, 0xd3, 0x99, 0x65, 0x07, 0x8b, 0x50, 0xef, 0x18, 0x74, 0x29, 0xf8, 0x2f, 0x8e,
	0x11, 0x2c, 0x6c, 0xc7, 0x2a, 0x4c, 0xf9, 0x7d, 0xf5, 0x4a, 0x7a, 0xa9, 0x8b, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x09, 0x68, 0xa0, 0x7d, 0xf4, 0x01, 0x00, 0x00,
}