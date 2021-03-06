// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/ad_serving_optimization_status.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible serving statuses.
type AdServingOptimizationStatusEnum_AdServingOptimizationStatus int32

const (
	// No value has been specified.
	AdServingOptimizationStatusEnum_UNSPECIFIED AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdServingOptimizationStatusEnum_UNKNOWN AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 1
	// Ad serving is optimized based on CTR for the campaign.
	AdServingOptimizationStatusEnum_OPTIMIZE AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 2
	// Ad serving is optimized based on CTR * Conversion for the campaign. If
	// the campaign is not in the conversion optimizer bidding strategy, it will
	// default to OPTIMIZED.
	AdServingOptimizationStatusEnum_CONVERSION_OPTIMIZE AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 3
	// Ads are rotated evenly for 90 days, then optimized for clicks.
	AdServingOptimizationStatusEnum_ROTATE AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 4
	// Show lower performing ads more evenly with higher performing ads, and do
	// not optimize.
	AdServingOptimizationStatusEnum_ROTATE_INDEFINITELY AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 5
	// Ad serving optimization status is not available.
	AdServingOptimizationStatusEnum_UNAVAILABLE AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 6
)

var AdServingOptimizationStatusEnum_AdServingOptimizationStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "OPTIMIZE",
	3: "CONVERSION_OPTIMIZE",
	4: "ROTATE",
	5: "ROTATE_INDEFINITELY",
	6: "UNAVAILABLE",
}

var AdServingOptimizationStatusEnum_AdServingOptimizationStatus_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"OPTIMIZE":            2,
	"CONVERSION_OPTIMIZE": 3,
	"ROTATE":              4,
	"ROTATE_INDEFINITELY": 5,
	"UNAVAILABLE":         6,
}

func (x AdServingOptimizationStatusEnum_AdServingOptimizationStatus) String() string {
	return proto.EnumName(AdServingOptimizationStatusEnum_AdServingOptimizationStatus_name, int32(x))
}

func (AdServingOptimizationStatusEnum_AdServingOptimizationStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_68f0c6c0e1ebc7e3, []int{0, 0}
}

// Possible ad serving statuses of a campaign.
type AdServingOptimizationStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdServingOptimizationStatusEnum) Reset()         { *m = AdServingOptimizationStatusEnum{} }
func (m *AdServingOptimizationStatusEnum) String() string { return proto.CompactTextString(m) }
func (*AdServingOptimizationStatusEnum) ProtoMessage()    {}
func (*AdServingOptimizationStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_68f0c6c0e1ebc7e3, []int{0}
}

func (m *AdServingOptimizationStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdServingOptimizationStatusEnum.Unmarshal(m, b)
}
func (m *AdServingOptimizationStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdServingOptimizationStatusEnum.Marshal(b, m, deterministic)
}
func (m *AdServingOptimizationStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdServingOptimizationStatusEnum.Merge(m, src)
}
func (m *AdServingOptimizationStatusEnum) XXX_Size() int {
	return xxx_messageInfo_AdServingOptimizationStatusEnum.Size(m)
}
func (m *AdServingOptimizationStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdServingOptimizationStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdServingOptimizationStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.AdServingOptimizationStatusEnum_AdServingOptimizationStatus", AdServingOptimizationStatusEnum_AdServingOptimizationStatus_name, AdServingOptimizationStatusEnum_AdServingOptimizationStatus_value)
	proto.RegisterType((*AdServingOptimizationStatusEnum)(nil), "google.ads.googleads.v3.enums.AdServingOptimizationStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/ad_serving_optimization_status.proto", fileDescriptor_68f0c6c0e1ebc7e3)
}

var fileDescriptor_68f0c6c0e1ebc7e3 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x8a, 0xdb, 0x30,
	0x14, 0xac, 0x9d, 0x36, 0x2d, 0x4a, 0xa1, 0xc6, 0x3d, 0x14, 0xda, 0x86, 0x36, 0xf9, 0x00, 0xf9,
	0xe0, 0x9b, 0x7a, 0x92, 0x13, 0x25, 0x88, 0xa6, 0xb2, 0x89, 0x1d, 0x97, 0x06, 0x83, 0x71, 0xd7,
	0xc6, 0x18, 0x62, 0xc9, 0x44, 0x4e, 0x0e, 0xfb, 0x25, 0x7b, 0xde, 0xe3, 0x7e, 0xc3, 0x7e, 0xc1,
	0x7e, 0xca, 0xde, 0xf7, 0xbe, 0x58, 0x4a, 0xcc, 0x5e, 0x36, 0x17, 0x31, 0x68, 0xe6, 0xcd, 0x3c,
	0xe6, 0x01, 0xaf, 0x14, 0xa2, 0xdc, 0x15, 0x4e, 0x96, 0x4b, 0x47, 0xc3, 0x0e, 0x1d, 0x5d, 0xa7,
	0xe0, 0x87, 0x5a, 0x3a, 0x59, 0x9e, 0xca, 0x62, 0x7f, 0xac, 0x78, 0x99, 0x8a, 0xa6, 0xad, 0xea,
	0xea, 0x3a, 0x6b, 0x2b, 0xc1, 0x53, 0xd9, 0x66, 0xed, 0x41, 0xc2, 0x66, 0x2f, 0x5a, 0x61, 0x8f,
	0xf5, 0x20, 0xcc, 0x72, 0x09, 0x7b, 0x0f, 0x78, 0x74, 0xa1, 0xf2, 0xf8, 0xfa, 0xfd, 0x1c, 0xd1,
	0x54, 0x4e, 0xc6, 0xb9, 0x68, 0x95, 0xc5, 0x69, 0x78, 0x7a, 0x6f, 0x80, 0x1f, 0x38, 0x0f, 0x75,
	0x88, 0xff, 0x22, 0x23, 0x54, 0x11, 0x84, 0x1f, 0xea, 0xe9, 0x8d, 0x01, 0xbe, 0x5d, 0xd0, 0xd8,
	0x9f, 0xc0, 0x68, 0xc3, 0xc2, 0x80, 0xcc, 0xe8, 0x82, 0x92, 0xb9, 0xf5, 0xc6, 0x1e, 0x81, 0xf7,
	0x1b, 0xf6, 0x9b, 0xf9, 0x7f, 0x99, 0x65, 0xd8, 0x1f, 0xc1, 0x07, 0x3f, 0x88, 0xe8, 0x1f, 0xba,
	0x25, 0x96, 0x69, 0x7f, 0x01, 0x9f, 0x67, 0x3e, 0x8b, 0xc9, 0x3a, 0xa4, 0x3e, 0x4b, 0x7b, 0x62,
	0x60, 0x03, 0x30, 0x5c, 0xfb, 0x11, 0x8e, 0x88, 0xf5, 0xb6, 0x13, 0x69, 0x9c, 0x52, 0x36, 0x27,
	0x0b, 0xca, 0x68, 0x44, 0x56, 0xff, 0xac, 0x77, 0x3a, 0x09, 0xc7, 0x98, 0xae, 0xb0, 0xb7, 0x22,
	0xd6, 0xd0, 0x7b, 0x32, 0xc0, 0xe4, 0x4a, 0xd4, 0xf0, 0x62, 0x05, 0xde, 0xcf, 0x0b, 0xdb, 0x07,
	0x5d, 0x0d, 0x81, 0xb1, 0x3d, 0x5d, 0x02, 0x96, 0x62, 0x97, 0xf1, 0x12, 0x8a, 0x7d, 0xe9, 0x94,
	0x05, 0x57, 0x25, 0x9d, 0x2f, 0xd3, 0x54, 0xf2, 0x95, 0x43, 0xfd, 0x52, 0xef, 0xad, 0x39, 0x58,
	0x62, 0x7c, 0x67, 0x8e, 0x97, 0xda, 0x0a, 0xe7, 0x12, 0x6a, 0xd8, 0xa1, 0xd8, 0x85, 0x5d, 0x9b,
	0xf2, 0xe1, 0xcc, 0x27, 0x38, 0x97, 0x49, 0xcf, 0x27, 0xb1, 0x9b, 0x28, 0xfe, 0xd1, 0x9c, 0xe8,
	0x4f, 0x84, 0x70, 0x2e, 0x11, 0xea, 0x15, 0x08, 0xc5, 0x2e, 0x42, 0x4a, 0xf3, 0x7f, 0xa8, 0x16,
	0x73, 0x9f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x79, 0x51, 0xc9, 0x40, 0x02, 0x00, 0x00,
}
