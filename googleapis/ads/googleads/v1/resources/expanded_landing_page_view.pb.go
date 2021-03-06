// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/expanded_landing_page_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A landing page view with metrics aggregated at the expanded final URL
// level.
type ExpandedLandingPageView struct {
	// Output only. The resource name of the expanded landing page view.
	// Expanded landing page view resource names have the form:
	//
	// `customers/{customer_id}/expandedLandingPageViews/{expanded_final_url_fingerprint}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The final URL that clicks are directed to.
	ExpandedFinalUrl     *wrappers.StringValue `protobuf:"bytes,2,opt,name=expanded_final_url,json=expandedFinalUrl,proto3" json:"expanded_final_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ExpandedLandingPageView) Reset()         { *m = ExpandedLandingPageView{} }
func (m *ExpandedLandingPageView) String() string { return proto.CompactTextString(m) }
func (*ExpandedLandingPageView) ProtoMessage()    {}
func (*ExpandedLandingPageView) Descriptor() ([]byte, []int) {
	return fileDescriptor_b526ad7687a265f6, []int{0}
}

func (m *ExpandedLandingPageView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpandedLandingPageView.Unmarshal(m, b)
}
func (m *ExpandedLandingPageView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpandedLandingPageView.Marshal(b, m, deterministic)
}
func (m *ExpandedLandingPageView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpandedLandingPageView.Merge(m, src)
}
func (m *ExpandedLandingPageView) XXX_Size() int {
	return xxx_messageInfo_ExpandedLandingPageView.Size(m)
}
func (m *ExpandedLandingPageView) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpandedLandingPageView.DiscardUnknown(m)
}

var xxx_messageInfo_ExpandedLandingPageView proto.InternalMessageInfo

func (m *ExpandedLandingPageView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ExpandedLandingPageView) GetExpandedFinalUrl() *wrappers.StringValue {
	if m != nil {
		return m.ExpandedFinalUrl
	}
	return nil
}

func init() {
	proto.RegisterType((*ExpandedLandingPageView)(nil), "google.ads.googleads.v1.resources.ExpandedLandingPageView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/expanded_landing_page_view.proto", fileDescriptor_b526ad7687a265f6)
}

var fileDescriptor_b526ad7687a265f6 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcf, 0x8a, 0xd4, 0x30,
	0x18, 0xa7, 0x1d, 0x10, 0xac, 0x0a, 0xd2, 0x8b, 0xeb, 0x32, 0xe8, 0xae, 0xb0, 0xb0, 0xa7, 0xc4,
	0xae, 0x17, 0x89, 0xa7, 0x14, 0x74, 0x61, 0x11, 0x1d, 0x46, 0xec, 0x41, 0x46, 0x4a, 0x66, 0xf2,
	0x4d, 0x0c, 0xb4, 0x49, 0x49, 0xda, 0x19, 0x61, 0x59, 0xc1, 0xa3, 0xaf, 0xe1, 0xd1, 0x47, 0x11,
	0x7c, 0x87, 0x3d, 0xef, 0x23, 0x78, 0x92, 0x4e, 0x93, 0xec, 0x20, 0x8c, 0xb2, 0xb7, 0x5f, 0xf9,
	0xfd, 0xc9, 0xf7, 0xfb, 0xbe, 0x26, 0xb9, 0xd0, 0x5a, 0x54, 0x80, 0x19, 0xb7, 0x78, 0x80, 0x3d,
	0x5a, 0x65, 0xd8, 0x80, 0xd5, 0x9d, 0x59, 0x80, 0xc5, 0xf0, 0xb9, 0x61, 0x8a, 0x03, 0x2f, 0x2b,
	0xa6, 0xb8, 0x54, 0xa2, 0x6c, 0x98, 0x80, 0x72, 0x25, 0x61, 0x8d, 0x1a, 0xa3, 0x5b, 0x9d, 0x1e,
	0x0e, 0x46, 0xc4, 0xb8, 0x45, 0x21, 0x03, 0xad, 0x32, 0x14, 0x32, 0xf6, 0x1f, 0xfb, 0x67, 0x1a,
	0x89, 0x97, 0x12, 0x2a, 0x5e, 0xce, 0xe1, 0x13, 0x5b, 0x49, 0x6d, 0x86, 0x8c, 0xfd, 0x87, 0x5b,
	0x02, 0x6f, 0x73, 0xd4, 0x23, 0x47, 0x6d, 0xbe, 0xe6, 0xdd, 0x12, 0xaf, 0x0d, 0x6b, 0x1a, 0x30,
	0xd6, 0xf1, 0xe3, 0x2d, 0x2b, 0x53, 0x4a, 0xb7, 0xac, 0x95, 0x5a, 0x39, 0xf6, 0xc9, 0xaf, 0x38,
	0x79, 0xf0, 0xd2, 0x35, 0x78, 0x3d, 0x14, 0x98, 0x30, 0x01, 0x85, 0x84, 0x75, 0xfa, 0x31, 0xb9,
	0xe7, 0xdf, 0x2a, 0x15, 0xab, 0x61, 0x2f, 0x3a, 0x88, 0x8e, 0x6f, 0xe7, 0xcf, 0x2f, 0xe9, 0xe8,
	0x37, 0x3d, 0x49, 0x9e, 0x5e, 0x97, 0x71, 0xa8, 0x91, 0x16, 0x2d, 0x74, 0x8d, 0x77, 0x04, 0x4e,
	0xef, 0xfa, 0xb8, 0x37, 0xac, 0x86, 0xf4, 0x6d, 0x92, 0x86, 0xdd, 0x2d, 0xa5, 0x62, 0x55, 0xd9,
	0x99, 0x6a, 0x2f, 0x3e, 0x88, 0x8e, 0xef, 0x9c, 0x8c, 0x5d, 0x24, 0xf2, 0xad, 0xd0, 0xbb, 0xd6,
	0x48, 0x25, 0x0a, 0x56, 0x75, 0x90, 0x8f, 0x2e, 0xe9, 0x68, 0x7a, 0xdf, 0x9b, 0x5f, 0xf5, 0xde,
	0xf7, 0xa6, 0x22, 0x5f, 0xa3, 0x2b, 0xfa, 0xe5, 0xe6, 0x63, 0xa5, 0x67, 0x8b, 0xce, 0xb6, 0xba,
	0x06, 0x63, 0xf1, 0xb9, 0x87, 0x17, 0xe1, 0xae, 0x7f, 0xa9, 0x2d, 0x3e, 0xdf, 0x7d, 0xf1, 0x8b,
	0xfc, 0x5b, 0x9c, 0x1c, 0x2d, 0x74, 0x8d, 0xfe, 0x7b, 0xf3, 0x7c, 0xbc, 0x63, 0x9c, 0x49, 0xdf,
	0x78, 0x12, 0x7d, 0x38, 0x73, 0x11, 0x42, 0x57, 0x4c, 0x09, 0xa4, 0x8d, 0xc0, 0x02, 0xd4, 0x66,
	0x1f, 0xf8, 0xba, 0xd8, 0x3f, 0xfe, 0xcc, 0x17, 0x01, 0x7d, 0x8f, 0x47, 0xa7, 0x94, 0xfe, 0x88,
	0x0f, 0x4f, 0x87, 0x48, 0xca, 0x2d, 0x1a, 0x60, 0x8f, 0x8a, 0x0c, 0x4d, 0xbd, 0xf2, 0xa7, 0xd7,
	0xcc, 0x28, 0xb7, 0xb3, 0xa0, 0x99, 0x15, 0xd9, 0x2c, 0x68, 0xae, 0xe2, 0xa3, 0x81, 0x20, 0x84,
	0x72, 0x4b, 0x48, 0x50, 0x11, 0x52, 0x64, 0x84, 0x04, 0xdd, 0xfc, 0xd6, 0x66, 0xd8, 0x67, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x2e, 0x53, 0x8e, 0x45, 0x03, 0x00, 0x00,
}
