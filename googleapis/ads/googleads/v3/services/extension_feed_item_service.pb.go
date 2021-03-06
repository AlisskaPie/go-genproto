// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/extension_feed_item_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for [ExtensionFeedItemService.GetExtensionFeedItem][google.ads.googleads.v3.services.ExtensionFeedItemService.GetExtensionFeedItem].
type GetExtensionFeedItemRequest struct {
	// Required. The resource name of the extension feed item to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetExtensionFeedItemRequest) Reset()         { *m = GetExtensionFeedItemRequest{} }
func (m *GetExtensionFeedItemRequest) String() string { return proto.CompactTextString(m) }
func (*GetExtensionFeedItemRequest) ProtoMessage()    {}
func (*GetExtensionFeedItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67cac715d7a110a8, []int{0}
}

func (m *GetExtensionFeedItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetExtensionFeedItemRequest.Unmarshal(m, b)
}
func (m *GetExtensionFeedItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetExtensionFeedItemRequest.Marshal(b, m, deterministic)
}
func (m *GetExtensionFeedItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetExtensionFeedItemRequest.Merge(m, src)
}
func (m *GetExtensionFeedItemRequest) XXX_Size() int {
	return xxx_messageInfo_GetExtensionFeedItemRequest.Size(m)
}
func (m *GetExtensionFeedItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetExtensionFeedItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetExtensionFeedItemRequest proto.InternalMessageInfo

func (m *GetExtensionFeedItemRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [ExtensionFeedItemService.MutateExtensionFeedItems][google.ads.googleads.v3.services.ExtensionFeedItemService.MutateExtensionFeedItems].
type MutateExtensionFeedItemsRequest struct {
	// Required. The ID of the customer whose extension feed items are being
	// modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual extension feed items.
	Operations []*ExtensionFeedItemOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateExtensionFeedItemsRequest) Reset()         { *m = MutateExtensionFeedItemsRequest{} }
func (m *MutateExtensionFeedItemsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateExtensionFeedItemsRequest) ProtoMessage()    {}
func (*MutateExtensionFeedItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67cac715d7a110a8, []int{1}
}

func (m *MutateExtensionFeedItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateExtensionFeedItemsRequest.Unmarshal(m, b)
}
func (m *MutateExtensionFeedItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateExtensionFeedItemsRequest.Marshal(b, m, deterministic)
}
func (m *MutateExtensionFeedItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateExtensionFeedItemsRequest.Merge(m, src)
}
func (m *MutateExtensionFeedItemsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateExtensionFeedItemsRequest.Size(m)
}
func (m *MutateExtensionFeedItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateExtensionFeedItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateExtensionFeedItemsRequest proto.InternalMessageInfo

func (m *MutateExtensionFeedItemsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateExtensionFeedItemsRequest) GetOperations() []*ExtensionFeedItemOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateExtensionFeedItemsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateExtensionFeedItemsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on an extension feed item.
type ExtensionFeedItemOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*ExtensionFeedItemOperation_Create
	//	*ExtensionFeedItemOperation_Update
	//	*ExtensionFeedItemOperation_Remove
	Operation            isExtensionFeedItemOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *ExtensionFeedItemOperation) Reset()         { *m = ExtensionFeedItemOperation{} }
func (m *ExtensionFeedItemOperation) String() string { return proto.CompactTextString(m) }
func (*ExtensionFeedItemOperation) ProtoMessage()    {}
func (*ExtensionFeedItemOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_67cac715d7a110a8, []int{2}
}

func (m *ExtensionFeedItemOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionFeedItemOperation.Unmarshal(m, b)
}
func (m *ExtensionFeedItemOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionFeedItemOperation.Marshal(b, m, deterministic)
}
func (m *ExtensionFeedItemOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionFeedItemOperation.Merge(m, src)
}
func (m *ExtensionFeedItemOperation) XXX_Size() int {
	return xxx_messageInfo_ExtensionFeedItemOperation.Size(m)
}
func (m *ExtensionFeedItemOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionFeedItemOperation.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionFeedItemOperation proto.InternalMessageInfo

func (m *ExtensionFeedItemOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isExtensionFeedItemOperation_Operation interface {
	isExtensionFeedItemOperation_Operation()
}

type ExtensionFeedItemOperation_Create struct {
	Create *resources.ExtensionFeedItem `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type ExtensionFeedItemOperation_Update struct {
	Update *resources.ExtensionFeedItem `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type ExtensionFeedItemOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*ExtensionFeedItemOperation_Create) isExtensionFeedItemOperation_Operation() {}

func (*ExtensionFeedItemOperation_Update) isExtensionFeedItemOperation_Operation() {}

func (*ExtensionFeedItemOperation_Remove) isExtensionFeedItemOperation_Operation() {}

func (m *ExtensionFeedItemOperation) GetOperation() isExtensionFeedItemOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *ExtensionFeedItemOperation) GetCreate() *resources.ExtensionFeedItem {
	if x, ok := m.GetOperation().(*ExtensionFeedItemOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *ExtensionFeedItemOperation) GetUpdate() *resources.ExtensionFeedItem {
	if x, ok := m.GetOperation().(*ExtensionFeedItemOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *ExtensionFeedItemOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*ExtensionFeedItemOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExtensionFeedItemOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExtensionFeedItemOperation_Create)(nil),
		(*ExtensionFeedItemOperation_Update)(nil),
		(*ExtensionFeedItemOperation_Remove)(nil),
	}
}

// Response message for an extension feed item mutate.
type MutateExtensionFeedItemsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateExtensionFeedItemResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *MutateExtensionFeedItemsResponse) Reset()         { *m = MutateExtensionFeedItemsResponse{} }
func (m *MutateExtensionFeedItemsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateExtensionFeedItemsResponse) ProtoMessage()    {}
func (*MutateExtensionFeedItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67cac715d7a110a8, []int{3}
}

func (m *MutateExtensionFeedItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateExtensionFeedItemsResponse.Unmarshal(m, b)
}
func (m *MutateExtensionFeedItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateExtensionFeedItemsResponse.Marshal(b, m, deterministic)
}
func (m *MutateExtensionFeedItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateExtensionFeedItemsResponse.Merge(m, src)
}
func (m *MutateExtensionFeedItemsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateExtensionFeedItemsResponse.Size(m)
}
func (m *MutateExtensionFeedItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateExtensionFeedItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateExtensionFeedItemsResponse proto.InternalMessageInfo

func (m *MutateExtensionFeedItemsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateExtensionFeedItemsResponse) GetResults() []*MutateExtensionFeedItemResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the extension feed item mutate.
type MutateExtensionFeedItemResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateExtensionFeedItemResult) Reset()         { *m = MutateExtensionFeedItemResult{} }
func (m *MutateExtensionFeedItemResult) String() string { return proto.CompactTextString(m) }
func (*MutateExtensionFeedItemResult) ProtoMessage()    {}
func (*MutateExtensionFeedItemResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_67cac715d7a110a8, []int{4}
}

func (m *MutateExtensionFeedItemResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateExtensionFeedItemResult.Unmarshal(m, b)
}
func (m *MutateExtensionFeedItemResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateExtensionFeedItemResult.Marshal(b, m, deterministic)
}
func (m *MutateExtensionFeedItemResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateExtensionFeedItemResult.Merge(m, src)
}
func (m *MutateExtensionFeedItemResult) XXX_Size() int {
	return xxx_messageInfo_MutateExtensionFeedItemResult.Size(m)
}
func (m *MutateExtensionFeedItemResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateExtensionFeedItemResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateExtensionFeedItemResult proto.InternalMessageInfo

func (m *MutateExtensionFeedItemResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetExtensionFeedItemRequest)(nil), "google.ads.googleads.v3.services.GetExtensionFeedItemRequest")
	proto.RegisterType((*MutateExtensionFeedItemsRequest)(nil), "google.ads.googleads.v3.services.MutateExtensionFeedItemsRequest")
	proto.RegisterType((*ExtensionFeedItemOperation)(nil), "google.ads.googleads.v3.services.ExtensionFeedItemOperation")
	proto.RegisterType((*MutateExtensionFeedItemsResponse)(nil), "google.ads.googleads.v3.services.MutateExtensionFeedItemsResponse")
	proto.RegisterType((*MutateExtensionFeedItemResult)(nil), "google.ads.googleads.v3.services.MutateExtensionFeedItemResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/extension_feed_item_service.proto", fileDescriptor_67cac715d7a110a8)
}

var fileDescriptor_67cac715d7a110a8 = []byte{
	// 793 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcb, 0x6a, 0xeb, 0x46,
	0x18, 0xae, 0xe4, 0x90, 0x36, 0xe3, 0xa4, 0x85, 0xe9, 0x4d, 0x75, 0x1a, 0x62, 0xd4, 0x40, 0x83,
	0x09, 0x12, 0xd8, 0xa1, 0x05, 0x25, 0xa1, 0xc8, 0x34, 0x4e, 0x02, 0xcd, 0x05, 0x05, 0x52, 0xda,
	0x1a, 0xc4, 0x44, 0x1a, 0xbb, 0x22, 0x92, 0x46, 0x9d, 0x19, 0x99, 0x86, 0x90, 0x4d, 0xe9, 0x1b,
	0xb4, 0x0f, 0x50, 0xba, 0xec, 0x7b, 0x74, 0x93, 0x6d, 0x57, 0xcd, 0x2a, 0x8b, 0xb3, 0x3a, 0x4f,
	0x70, 0x38, 0x8b, 0xc3, 0x41, 0x1a, 0x8d, 0x2f, 0x71, 0x14, 0x43, 0xb2, 0xfb, 0x3d, 0xff, 0xa7,
	0xef, 0xfb, 0xef, 0x06, 0xed, 0x3e, 0x21, 0xfd, 0x10, 0x9b, 0xc8, 0x67, 0xa6, 0x30, 0x33, 0x6b,
	0xd0, 0x32, 0x19, 0xa6, 0x83, 0xc0, 0xc3, 0xcc, 0xc4, 0xbf, 0x72, 0x1c, 0xb3, 0x80, 0xc4, 0x6e,
	0x0f, 0x63, 0xdf, 0x0d, 0x38, 0x8e, 0xdc, 0xc2, 0x69, 0x24, 0x94, 0x70, 0x02, 0xeb, 0xe2, 0x43,
	0x03, 0xf9, 0xcc, 0x18, 0x72, 0x18, 0x83, 0x96, 0x21, 0x39, 0x6a, 0x5b, 0x65, 0x2a, 0x14, 0x33,
	0x92, 0xd2, 0x12, 0x19, 0x41, 0x5f, 0xfb, 0x5c, 0x7e, 0x9c, 0x04, 0x26, 0x8a, 0x63, 0xc2, 0x11,
	0x0f, 0x48, 0xcc, 0x0a, 0xef, 0xa7, 0x63, 0x5e, 0x2f, 0x0c, 0x70, 0xcc, 0x0b, 0xc7, 0xea, 0x98,
	0xa3, 0x17, 0xe0, 0xd0, 0x77, 0xcf, 0xf1, 0xcf, 0x68, 0x10, 0x10, 0x5a, 0x00, 0x3e, 0x1b, 0x03,
	0xc8, 0x38, 0x0a, 0x57, 0x91, 0x91, 0x99, 0xff, 0x3a, 0x4f, 0x7b, 0x05, 0x41, 0x84, 0xd8, 0xc5,
	0x3d, 0x59, 0x9a, 0x78, 0x26, 0xe3, 0x88, 0xa7, 0x45, 0x3c, 0xfa, 0x00, 0x2c, 0xef, 0x61, 0xbe,
	0x2b, 0xb3, 0xe9, 0x60, 0xec, 0x1f, 0x70, 0x1c, 0x39, 0xf8, 0x97, 0x14, 0x33, 0x0e, 0xbf, 0x07,
	0x4b, 0x52, 0xcb, 0x8d, 0x51, 0x84, 0x35, 0xa5, 0xae, 0xac, 0x2f, 0xb4, 0x9b, 0x77, 0xb6, 0xfa,
	0xda, 0xde, 0x00, 0x8d, 0x51, 0xfd, 0x0a, 0x2b, 0x09, 0x98, 0xe1, 0x91, 0xc8, 0x9c, 0x66, 0x5c,
	0x94, 0x44, 0x47, 0x28, 0xc2, 0xfa, 0x2b, 0x05, 0xac, 0x1e, 0xa6, 0x1c, 0x71, 0x3c, 0x85, 0x64,
	0x52, 0x7c, 0x0d, 0x54, 0xbd, 0x94, 0x71, 0x12, 0x61, 0xea, 0x06, 0x7e, 0x21, 0x5d, 0xb9, 0xb3,
	0x55, 0x07, 0xc8, 0xf7, 0x03, 0x1f, 0x22, 0x00, 0x48, 0x82, 0xa9, 0xa8, 0xb2, 0xa6, 0xd6, 0x2b,
	0xeb, 0xd5, 0xe6, 0xb6, 0x31, 0xab, 0xc7, 0xc6, 0x94, 0xec, 0xb1, 0x24, 0x29, 0x24, 0x46, 0xa4,
	0xf0, 0x4b, 0xf0, 0x41, 0x82, 0x28, 0x0f, 0x50, 0xe8, 0xf6, 0x50, 0x10, 0xa6, 0x14, 0x6b, 0x95,
	0xba, 0xb2, 0xfe, 0x9e, 0xf3, 0x7e, 0xf1, 0xdc, 0x11, 0xaf, 0xf0, 0x0b, 0xb0, 0x34, 0x40, 0x61,
	0xe0, 0x23, 0x8e, 0x5d, 0x12, 0x87, 0x97, 0xda, 0x5c, 0x0e, 0x5b, 0x94, 0x8f, 0xc7, 0x71, 0x78,
	0xa9, 0xff, 0xa5, 0x82, 0x5a, 0xb9, 0x3a, 0xdc, 0x02, 0xd5, 0x34, 0xc9, 0x19, 0xb2, 0xfe, 0xe5,
	0x0c, 0xd5, 0x66, 0x4d, 0x26, 0x24, 0x5b, 0x6c, 0x74, 0xb2, 0x16, 0x1f, 0x22, 0x76, 0xe1, 0x00,
	0x01, 0xcf, 0x6c, 0x78, 0x04, 0xe6, 0x3d, 0x8a, 0x11, 0x17, 0x8d, 0xaa, 0x36, 0x37, 0x4b, 0x0b,
	0x31, 0x1c, 0xe5, 0xe9, 0x4a, 0xec, 0xbf, 0xe3, 0x14, 0x2c, 0x19, 0x9f, 0x60, 0xd7, 0xd4, 0xe7,
	0xf1, 0x09, 0x16, 0xa8, 0x81, 0x79, 0x8a, 0x23, 0x32, 0x10, 0x05, 0x5c, 0xc8, 0x3c, 0xe2, 0x77,
	0xbb, 0x0a, 0x16, 0x86, 0x15, 0xd7, 0xff, 0x55, 0x40, 0xbd, 0x7c, 0x3a, 0x58, 0x42, 0x62, 0x86,
	0x61, 0x07, 0x7c, 0x7c, 0xaf, 0x2b, 0x2e, 0xa6, 0x94, 0xd0, 0x9c, 0xba, 0xda, 0x84, 0x32, 0x54,
	0x9a, 0x78, 0xc6, 0x69, 0x3e, 0xf3, 0xce, 0x87, 0x93, 0xfd, 0xda, 0xcd, 0xe0, 0xf0, 0x07, 0xf0,
	0x2e, 0xc5, 0x2c, 0x0d, 0xb9, 0x9c, 0x9e, 0x6f, 0x66, 0x4f, 0x4f, 0x49, 0x70, 0x4e, 0xce, 0xe3,
	0x48, 0x3e, 0xfd, 0x5b, 0xb0, 0xf2, 0x28, 0x32, 0x1b, 0x98, 0x07, 0xf6, 0x6b, 0x72, 0x57, 0x9a,
	0x7f, 0xce, 0x01, 0x6d, 0x8a, 0xe0, 0x54, 0x84, 0x02, 0xff, 0x57, 0xc0, 0x47, 0x0f, 0x6d, 0x30,
	0xdc, 0x99, 0x9d, 0xc5, 0x23, 0x9b, 0x5f, 0x7b, 0x52, 0xa7, 0xf5, 0xef, 0x6e, 0xed, 0xc9, 0x84,
	0x7e, 0xfb, 0xef, 0xc5, 0x1f, 0xea, 0x57, 0x70, 0x33, 0xbb, 0x9e, 0x57, 0x13, 0x9e, 0x1d, 0xb9,
	0xc3, 0xcc, 0x6c, 0x8c, 0xce, 0xe9, 0xb0, 0xcd, 0x66, 0xe3, 0x1a, 0xbe, 0x51, 0x80, 0x56, 0x36,
	0x06, 0xd0, 0x7e, 0x72, 0x97, 0xe4, 0x81, 0xa9, 0xb5, 0x9f, 0x43, 0x21, 0xa6, 0x50, 0xff, 0xe9,
	0xd6, 0xfe, 0x64, 0xec, 0x4a, 0x6d, 0x8c, 0xce, 0x46, 0x9e, 0xfa, 0xb6, 0xfe, 0x75, 0x96, 0xfa,
	0x28, 0xd7, 0xab, 0x31, 0xf0, 0x4e, 0xe3, 0xfa, 0x81, 0xcc, 0xad, 0x28, 0xd7, 0xb4, 0x94, 0x46,
	0x6d, 0xf9, 0xc6, 0xd6, 0xca, 0x4e, 0x6c, 0xfb, 0x77, 0x15, 0xac, 0x79, 0x24, 0x9a, 0x99, 0x43,
	0x7b, 0xa5, 0x6c, 0x78, 0x4e, 0xb2, 0x63, 0x72, 0xa2, 0xfc, 0xb8, 0x5f, 0x50, 0xf4, 0x49, 0x88,
	0xe2, 0xbe, 0x41, 0x68, 0xdf, 0xec, 0xe3, 0x38, 0x3f, 0x35, 0xe6, 0x48, 0xb4, 0xfc, 0x4f, 0x77,
	0x4b, 0x1a, 0x7f, 0xab, 0x95, 0x3d, 0xdb, 0xfe, 0x47, 0xad, 0xef, 0x09, 0x42, 0xdb, 0x67, 0x86,
	0x30, 0x33, 0xeb, 0xac, 0x65, 0x14, 0xc2, 0xec, 0x46, 0x42, 0xba, 0xb6, 0xcf, 0xba, 0x43, 0x48,
	0xf7, 0xac, 0xd5, 0x95, 0x90, 0x97, 0xea, 0x9a, 0x78, 0xb7, 0x2c, 0xdb, 0x67, 0x96, 0x35, 0x04,
	0x59, 0xd6, 0x59, 0xcb, 0xb2, 0x24, 0xec, 0x7c, 0x3e, 0x8f, 0xb3, 0xf5, 0x36, 0x00, 0x00, 0xff,
	0xff, 0xcd, 0x62, 0xb9, 0x3e, 0x1b, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExtensionFeedItemServiceClient is the client API for ExtensionFeedItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExtensionFeedItemServiceClient interface {
	// Returns the requested extension feed item in full detail.
	GetExtensionFeedItem(ctx context.Context, in *GetExtensionFeedItemRequest, opts ...grpc.CallOption) (*resources.ExtensionFeedItem, error)
	// Creates, updates, or removes extension feed items. Operation
	// statuses are returned.
	MutateExtensionFeedItems(ctx context.Context, in *MutateExtensionFeedItemsRequest, opts ...grpc.CallOption) (*MutateExtensionFeedItemsResponse, error)
}

type extensionFeedItemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExtensionFeedItemServiceClient(cc grpc.ClientConnInterface) ExtensionFeedItemServiceClient {
	return &extensionFeedItemServiceClient{cc}
}

func (c *extensionFeedItemServiceClient) GetExtensionFeedItem(ctx context.Context, in *GetExtensionFeedItemRequest, opts ...grpc.CallOption) (*resources.ExtensionFeedItem, error) {
	out := new(resources.ExtensionFeedItem)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.ExtensionFeedItemService/GetExtensionFeedItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extensionFeedItemServiceClient) MutateExtensionFeedItems(ctx context.Context, in *MutateExtensionFeedItemsRequest, opts ...grpc.CallOption) (*MutateExtensionFeedItemsResponse, error) {
	out := new(MutateExtensionFeedItemsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.ExtensionFeedItemService/MutateExtensionFeedItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionFeedItemServiceServer is the server API for ExtensionFeedItemService service.
type ExtensionFeedItemServiceServer interface {
	// Returns the requested extension feed item in full detail.
	GetExtensionFeedItem(context.Context, *GetExtensionFeedItemRequest) (*resources.ExtensionFeedItem, error)
	// Creates, updates, or removes extension feed items. Operation
	// statuses are returned.
	MutateExtensionFeedItems(context.Context, *MutateExtensionFeedItemsRequest) (*MutateExtensionFeedItemsResponse, error)
}

// UnimplementedExtensionFeedItemServiceServer can be embedded to have forward compatible implementations.
type UnimplementedExtensionFeedItemServiceServer struct {
}

func (*UnimplementedExtensionFeedItemServiceServer) GetExtensionFeedItem(ctx context.Context, req *GetExtensionFeedItemRequest) (*resources.ExtensionFeedItem, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetExtensionFeedItem not implemented")
}
func (*UnimplementedExtensionFeedItemServiceServer) MutateExtensionFeedItems(ctx context.Context, req *MutateExtensionFeedItemsRequest) (*MutateExtensionFeedItemsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateExtensionFeedItems not implemented")
}

func RegisterExtensionFeedItemServiceServer(s *grpc.Server, srv ExtensionFeedItemServiceServer) {
	s.RegisterService(&_ExtensionFeedItemService_serviceDesc, srv)
}

func _ExtensionFeedItemService_GetExtensionFeedItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExtensionFeedItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionFeedItemServiceServer).GetExtensionFeedItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.ExtensionFeedItemService/GetExtensionFeedItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionFeedItemServiceServer).GetExtensionFeedItem(ctx, req.(*GetExtensionFeedItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExtensionFeedItemService_MutateExtensionFeedItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateExtensionFeedItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionFeedItemServiceServer).MutateExtensionFeedItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.ExtensionFeedItemService/MutateExtensionFeedItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionFeedItemServiceServer).MutateExtensionFeedItems(ctx, req.(*MutateExtensionFeedItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExtensionFeedItemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.ExtensionFeedItemService",
	HandlerType: (*ExtensionFeedItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExtensionFeedItem",
			Handler:    _ExtensionFeedItemService_GetExtensionFeedItem_Handler,
		},
		{
			MethodName: "MutateExtensionFeedItems",
			Handler:    _ExtensionFeedItemService_MutateExtensionFeedItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/extension_feed_item_service.proto",
}
