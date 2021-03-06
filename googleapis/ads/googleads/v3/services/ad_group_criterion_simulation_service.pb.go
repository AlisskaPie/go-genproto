// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/ad_group_criterion_simulation_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for
// [AdGroupCriterionSimulationService.GetAdGroupCriterionSimulation][google.ads.googleads.v3.services.AdGroupCriterionSimulationService.GetAdGroupCriterionSimulation].
type GetAdGroupCriterionSimulationRequest struct {
	// Required. The resource name of the ad group criterion simulation to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupCriterionSimulationRequest) Reset()         { *m = GetAdGroupCriterionSimulationRequest{} }
func (m *GetAdGroupCriterionSimulationRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupCriterionSimulationRequest) ProtoMessage()    {}
func (*GetAdGroupCriterionSimulationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da600f48a422603f, []int{0}
}

func (m *GetAdGroupCriterionSimulationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupCriterionSimulationRequest.Unmarshal(m, b)
}
func (m *GetAdGroupCriterionSimulationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupCriterionSimulationRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupCriterionSimulationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupCriterionSimulationRequest.Merge(m, src)
}
func (m *GetAdGroupCriterionSimulationRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupCriterionSimulationRequest.Size(m)
}
func (m *GetAdGroupCriterionSimulationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupCriterionSimulationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupCriterionSimulationRequest proto.InternalMessageInfo

func (m *GetAdGroupCriterionSimulationRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupCriterionSimulationRequest)(nil), "google.ads.googleads.v3.services.GetAdGroupCriterionSimulationRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/ad_group_criterion_simulation_service.proto", fileDescriptor_da600f48a422603f)
}

var fileDescriptor_da600f48a422603f = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x31, 0x8b, 0xd4, 0x40,
	0x1c, 0xc5, 0x49, 0x04, 0xc1, 0xa0, 0x4d, 0x1a, 0xcf, 0x55, 0x71, 0x3d, 0x16, 0x91, 0x2b, 0x66,
	0xc0, 0x60, 0x33, 0xc7, 0x22, 0xb3, 0xa2, 0x6b, 0x21, 0x72, 0x78, 0xb8, 0x85, 0x04, 0xe2, 0x5c,
	0x66, 0x8c, 0x03, 0x49, 0x26, 0xce, 0x7f, 0x92, 0x46, 0x6c, 0xac, 0x6c, 0xac, 0xfc, 0x06, 0x96,
	0x7e, 0x0a, 0xeb, 0x6b, 0xed, 0xac, 0x2c, 0xac, 0xfc, 0x08, 0x82, 0x20, 0xc9, 0x64, 0x72, 0x7b,
	0x45, 0x36, 0x76, 0x8f, 0xbc, 0x97, 0xdf, 0xff, 0x9f, 0x37, 0x93, 0xe0, 0x69, 0xa6, 0x54, 0x96,
	0x0b, 0xcc, 0x38, 0x60, 0x2b, 0x5b, 0xd5, 0x44, 0x18, 0x84, 0x6e, 0x64, 0x2a, 0x00, 0x33, 0x9e,
	0x64, 0x5a, 0xd5, 0x55, 0x92, 0x6a, 0x69, 0x84, 0x96, 0xaa, 0x4c, 0x40, 0x16, 0x75, 0xce, 0x4c,
	0x27, 0x6d, 0x0c, 0x55, 0x5a, 0x19, 0x15, 0xce, 0x2d, 0x02, 0x31, 0x0e, 0x68, 0xa0, 0xa1, 0x26,
	0x42, 0x8e, 0x36, 0x7b, 0x34, 0x36, 0x4f, 0x0b, 0x50, 0xb5, 0x9e, 0x1c, 0x68, 0x07, 0xcd, 0x6e,
	0x38, 0x4c, 0x25, 0x31, 0x2b, 0x4b, 0x65, 0x3a, 0x13, 0x7a, 0xf7, 0xea, 0x96, 0x9b, 0xe6, 0x52,
	0x94, 0xa6, 0x37, 0x6e, 0x6d, 0x19, 0xaf, 0xa5, 0xc8, 0x79, 0x72, 0x22, 0xde, 0xb0, 0x46, 0x2a,
	0xdd, 0x07, 0xae, 0x6d, 0x05, 0xdc, 0x46, 0xd6, 0xda, 0xff, 0xe8, 0x05, 0x8b, 0xb5, 0x30, 0x94,
	0xaf, 0xdb, 0xe5, 0x1e, 0xba, 0xdd, 0x8e, 0x87, 0xd5, 0x9e, 0x8b, 0xb7, 0xb5, 0x00, 0x13, 0xbe,
	0x0a, 0xae, 0xb8, 0x57, 0x93, 0x92, 0x15, 0x62, 0xcf, 0x9b, 0x7b, 0x77, 0x2f, 0xad, 0x0e, 0x7f,
	0x52, 0xff, 0x0f, 0xbd, 0x1f, 0x44, 0x67, 0xc5, 0xf4, 0xaa, 0x92, 0x80, 0x52, 0x55, 0xe0, 0x1d,
	0xe8, 0xcb, 0x8e, 0xf8, 0x8c, 0x15, 0xe2, 0xde, 0x37, 0x3f, 0xb8, 0x3d, 0x1e, 0x3e, 0xb6, 0x5d,
	0x87, 0x7f, 0xbd, 0xe0, 0xe6, 0xce, 0x85, 0xc3, 0xc7, 0x68, 0xea, 0xbc, 0xd0, 0xff, 0x7c, 0xf1,
	0x6c, 0x39, 0xca, 0x19, 0x4e, 0x15, 0x8d, 0x53, 0xf6, 0x5f, 0xfc, 0xa0, 0xe7, 0x1b, 0xfb, 0xf0,
	0xfd, 0xd7, 0x67, 0xff, 0x41, 0xb8, 0x6c, 0xef, 0xc5, 0xbb, 0x73, 0xce, 0x32, 0xad, 0xc1, 0xa8,
	0x42, 0x68, 0xc0, 0x07, 0x98, 0x8d, 0x22, 0x01, 0x1f, 0xbc, 0x9f, 0x5d, 0x3f, 0xa5, 0x7b, 0x63,
	0x5d, 0xaf, 0x3e, 0xf9, 0xc1, 0x22, 0x55, 0xc5, 0x64, 0x01, 0xab, 0x3b, 0x93, 0x45, 0x1f, 0xb5,
	0xd7, 0xe3, 0xc8, 0x7b, 0xf9, 0xa4, 0x67, 0x65, 0x2a, 0x67, 0x65, 0x86, 0x94, 0xce, 0x70, 0x26,
	0xca, 0xee, 0xf2, 0xe0, 0xb3, 0xe9, 0xe3, 0x7f, 0xda, 0xa1, 0x13, 0x5f, 0xfc, 0x0b, 0x6b, 0x4a,
	0xbf, 0xfa, 0xf3, 0xb5, 0x05, 0x52, 0x0e, 0xc8, 0xca, 0x56, 0x6d, 0x22, 0xd4, 0x0f, 0x86, 0x53,
	0x17, 0x89, 0x29, 0x87, 0x78, 0x88, 0xc4, 0x9b, 0x28, 0x76, 0x91, 0xdf, 0xfe, 0xc2, 0x3e, 0x27,
	0x84, 0x72, 0x20, 0x64, 0x08, 0x11, 0xb2, 0x89, 0x08, 0x71, 0xb1, 0x93, 0x8b, 0xdd, 0x9e, 0xd1,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0xc1, 0x61, 0x2b, 0x10, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdGroupCriterionSimulationServiceClient is the client API for AdGroupCriterionSimulationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupCriterionSimulationServiceClient interface {
	// Returns the requested ad group criterion simulation in full detail.
	GetAdGroupCriterionSimulation(ctx context.Context, in *GetAdGroupCriterionSimulationRequest, opts ...grpc.CallOption) (*resources.AdGroupCriterionSimulation, error)
}

type adGroupCriterionSimulationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupCriterionSimulationServiceClient(cc grpc.ClientConnInterface) AdGroupCriterionSimulationServiceClient {
	return &adGroupCriterionSimulationServiceClient{cc}
}

func (c *adGroupCriterionSimulationServiceClient) GetAdGroupCriterionSimulation(ctx context.Context, in *GetAdGroupCriterionSimulationRequest, opts ...grpc.CallOption) (*resources.AdGroupCriterionSimulation, error) {
	out := new(resources.AdGroupCriterionSimulation)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.AdGroupCriterionSimulationService/GetAdGroupCriterionSimulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupCriterionSimulationServiceServer is the server API for AdGroupCriterionSimulationService service.
type AdGroupCriterionSimulationServiceServer interface {
	// Returns the requested ad group criterion simulation in full detail.
	GetAdGroupCriterionSimulation(context.Context, *GetAdGroupCriterionSimulationRequest) (*resources.AdGroupCriterionSimulation, error)
}

// UnimplementedAdGroupCriterionSimulationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupCriterionSimulationServiceServer struct {
}

func (*UnimplementedAdGroupCriterionSimulationServiceServer) GetAdGroupCriterionSimulation(ctx context.Context, req *GetAdGroupCriterionSimulationRequest) (*resources.AdGroupCriterionSimulation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdGroupCriterionSimulation not implemented")
}

func RegisterAdGroupCriterionSimulationServiceServer(s *grpc.Server, srv AdGroupCriterionSimulationServiceServer) {
	s.RegisterService(&_AdGroupCriterionSimulationService_serviceDesc, srv)
}

func _AdGroupCriterionSimulationService_GetAdGroupCriterionSimulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupCriterionSimulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupCriterionSimulationServiceServer).GetAdGroupCriterionSimulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.AdGroupCriterionSimulationService/GetAdGroupCriterionSimulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupCriterionSimulationServiceServer).GetAdGroupCriterionSimulation(ctx, req.(*GetAdGroupCriterionSimulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupCriterionSimulationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.AdGroupCriterionSimulationService",
	HandlerType: (*AdGroupCriterionSimulationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupCriterionSimulation",
			Handler:    _AdGroupCriterionSimulationService_GetAdGroupCriterionSimulation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/ad_group_criterion_simulation_service.proto",
}
