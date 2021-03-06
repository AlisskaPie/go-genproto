// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/customer_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// Request message for [CustomerService.GetCustomer][google.ads.googleads.v3.services.CustomerService.GetCustomer].
type GetCustomerRequest struct {
	// Required. The resource name of the customer to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerRequest) Reset()         { *m = GetCustomerRequest{} }
func (m *GetCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerRequest) ProtoMessage()    {}
func (*GetCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62d137bd13ed426f, []int{0}
}

func (m *GetCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerRequest.Unmarshal(m, b)
}
func (m *GetCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerRequest.Marshal(b, m, deterministic)
}
func (m *GetCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerRequest.Merge(m, src)
}
func (m *GetCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerRequest.Size(m)
}
func (m *GetCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerRequest proto.InternalMessageInfo

func (m *GetCustomerRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CustomerService.MutateCustomer][google.ads.googleads.v3.services.CustomerService.MutateCustomer].
type MutateCustomerRequest struct {
	// Required. The ID of the customer being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The operation to perform on the customer
	Operation *CustomerOperation `protobuf:"bytes,4,opt,name=operation,proto3" json:"operation,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,5,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCustomerRequest) Reset()         { *m = MutateCustomerRequest{} }
func (m *MutateCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerRequest) ProtoMessage()    {}
func (*MutateCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62d137bd13ed426f, []int{1}
}

func (m *MutateCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerRequest.Unmarshal(m, b)
}
func (m *MutateCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerRequest.Marshal(b, m, deterministic)
}
func (m *MutateCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerRequest.Merge(m, src)
}
func (m *MutateCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerRequest.Size(m)
}
func (m *MutateCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerRequest proto.InternalMessageInfo

func (m *MutateCustomerRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCustomerRequest) GetOperation() *CustomerOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *MutateCustomerRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// Request message for [CustomerService.CreateCustomerClient][google.ads.googleads.v3.services.CustomerService.CreateCustomerClient].
type CreateCustomerClientRequest struct {
	// Required. The ID of the Manager under whom client customer is being created.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The new client customer to create. The resource name on this customer
	// will be ignored.
	CustomerClient *resources.Customer `protobuf:"bytes,2,opt,name=customer_client,json=customerClient,proto3" json:"customer_client,omitempty"`
	// Email address of the user who should be invited on the created client
	// customer. Accessible to whitelisted customers only.
	EmailAddress *wrappers.StringValue `protobuf:"bytes,3,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	// The proposed role of user on the created client customer.
	// Accessible to whitelisted customers only.
	AccessRole           enums.AccessRoleEnum_AccessRole `protobuf:"varint,4,opt,name=access_role,json=accessRole,proto3,enum=google.ads.googleads.v3.enums.AccessRoleEnum_AccessRole" json:"access_role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *CreateCustomerClientRequest) Reset()         { *m = CreateCustomerClientRequest{} }
func (m *CreateCustomerClientRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCustomerClientRequest) ProtoMessage()    {}
func (*CreateCustomerClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62d137bd13ed426f, []int{2}
}

func (m *CreateCustomerClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCustomerClientRequest.Unmarshal(m, b)
}
func (m *CreateCustomerClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCustomerClientRequest.Marshal(b, m, deterministic)
}
func (m *CreateCustomerClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCustomerClientRequest.Merge(m, src)
}
func (m *CreateCustomerClientRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCustomerClientRequest.Size(m)
}
func (m *CreateCustomerClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCustomerClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCustomerClientRequest proto.InternalMessageInfo

func (m *CreateCustomerClientRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *CreateCustomerClientRequest) GetCustomerClient() *resources.Customer {
	if m != nil {
		return m.CustomerClient
	}
	return nil
}

func (m *CreateCustomerClientRequest) GetEmailAddress() *wrappers.StringValue {
	if m != nil {
		return m.EmailAddress
	}
	return nil
}

func (m *CreateCustomerClientRequest) GetAccessRole() enums.AccessRoleEnum_AccessRole {
	if m != nil {
		return m.AccessRole
	}
	return enums.AccessRoleEnum_UNSPECIFIED
}

// A single update on a customer.
type CustomerOperation struct {
	// Mutate operation. Only updates are supported for customer.
	Update *resources.Customer `protobuf:"bytes,1,opt,name=update,proto3" json:"update,omitempty"`
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CustomerOperation) Reset()         { *m = CustomerOperation{} }
func (m *CustomerOperation) String() string { return proto.CompactTextString(m) }
func (*CustomerOperation) ProtoMessage()    {}
func (*CustomerOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_62d137bd13ed426f, []int{3}
}

func (m *CustomerOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerOperation.Unmarshal(m, b)
}
func (m *CustomerOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerOperation.Marshal(b, m, deterministic)
}
func (m *CustomerOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerOperation.Merge(m, src)
}
func (m *CustomerOperation) XXX_Size() int {
	return xxx_messageInfo_CustomerOperation.Size(m)
}
func (m *CustomerOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerOperation proto.InternalMessageInfo

func (m *CustomerOperation) GetUpdate() *resources.Customer {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *CustomerOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// Response message for CreateCustomerClient mutate.
type CreateCustomerClientResponse struct {
	// The resource name of the newly created customer client.
	ResourceName         string   `protobuf:"bytes,2,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCustomerClientResponse) Reset()         { *m = CreateCustomerClientResponse{} }
func (m *CreateCustomerClientResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCustomerClientResponse) ProtoMessage()    {}
func (*CreateCustomerClientResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62d137bd13ed426f, []int{4}
}

func (m *CreateCustomerClientResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCustomerClientResponse.Unmarshal(m, b)
}
func (m *CreateCustomerClientResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCustomerClientResponse.Marshal(b, m, deterministic)
}
func (m *CreateCustomerClientResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCustomerClientResponse.Merge(m, src)
}
func (m *CreateCustomerClientResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCustomerClientResponse.Size(m)
}
func (m *CreateCustomerClientResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCustomerClientResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCustomerClientResponse proto.InternalMessageInfo

func (m *CreateCustomerClientResponse) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Response message for customer mutate.
type MutateCustomerResponse struct {
	// Result for the mutate.
	Result               *MutateCustomerResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MutateCustomerResponse) Reset()         { *m = MutateCustomerResponse{} }
func (m *MutateCustomerResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerResponse) ProtoMessage()    {}
func (*MutateCustomerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62d137bd13ed426f, []int{5}
}

func (m *MutateCustomerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerResponse.Unmarshal(m, b)
}
func (m *MutateCustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerResponse.Marshal(b, m, deterministic)
}
func (m *MutateCustomerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerResponse.Merge(m, src)
}
func (m *MutateCustomerResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerResponse.Size(m)
}
func (m *MutateCustomerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerResponse proto.InternalMessageInfo

func (m *MutateCustomerResponse) GetResult() *MutateCustomerResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// The result for the customer mutate.
type MutateCustomerResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCustomerResult) Reset()         { *m = MutateCustomerResult{} }
func (m *MutateCustomerResult) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerResult) ProtoMessage()    {}
func (*MutateCustomerResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_62d137bd13ed426f, []int{6}
}

func (m *MutateCustomerResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerResult.Unmarshal(m, b)
}
func (m *MutateCustomerResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerResult.Marshal(b, m, deterministic)
}
func (m *MutateCustomerResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerResult.Merge(m, src)
}
func (m *MutateCustomerResult) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerResult.Size(m)
}
func (m *MutateCustomerResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerResult proto.InternalMessageInfo

func (m *MutateCustomerResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CustomerService.ListAccessibleCustomers][google.ads.googleads.v3.services.CustomerService.ListAccessibleCustomers].
type ListAccessibleCustomersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAccessibleCustomersRequest) Reset()         { *m = ListAccessibleCustomersRequest{} }
func (m *ListAccessibleCustomersRequest) String() string { return proto.CompactTextString(m) }
func (*ListAccessibleCustomersRequest) ProtoMessage()    {}
func (*ListAccessibleCustomersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62d137bd13ed426f, []int{7}
}

func (m *ListAccessibleCustomersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAccessibleCustomersRequest.Unmarshal(m, b)
}
func (m *ListAccessibleCustomersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAccessibleCustomersRequest.Marshal(b, m, deterministic)
}
func (m *ListAccessibleCustomersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAccessibleCustomersRequest.Merge(m, src)
}
func (m *ListAccessibleCustomersRequest) XXX_Size() int {
	return xxx_messageInfo_ListAccessibleCustomersRequest.Size(m)
}
func (m *ListAccessibleCustomersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAccessibleCustomersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAccessibleCustomersRequest proto.InternalMessageInfo

// Response message for [CustomerService.ListAccessibleCustomers][google.ads.googleads.v3.services.CustomerService.ListAccessibleCustomers].
type ListAccessibleCustomersResponse struct {
	// Resource name of customers directly accessible by the
	// user authenticating the call.
	ResourceNames        []string `protobuf:"bytes,1,rep,name=resource_names,json=resourceNames,proto3" json:"resource_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAccessibleCustomersResponse) Reset()         { *m = ListAccessibleCustomersResponse{} }
func (m *ListAccessibleCustomersResponse) String() string { return proto.CompactTextString(m) }
func (*ListAccessibleCustomersResponse) ProtoMessage()    {}
func (*ListAccessibleCustomersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62d137bd13ed426f, []int{8}
}

func (m *ListAccessibleCustomersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAccessibleCustomersResponse.Unmarshal(m, b)
}
func (m *ListAccessibleCustomersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAccessibleCustomersResponse.Marshal(b, m, deterministic)
}
func (m *ListAccessibleCustomersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAccessibleCustomersResponse.Merge(m, src)
}
func (m *ListAccessibleCustomersResponse) XXX_Size() int {
	return xxx_messageInfo_ListAccessibleCustomersResponse.Size(m)
}
func (m *ListAccessibleCustomersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAccessibleCustomersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAccessibleCustomersResponse proto.InternalMessageInfo

func (m *ListAccessibleCustomersResponse) GetResourceNames() []string {
	if m != nil {
		return m.ResourceNames
	}
	return nil
}

func init() {
	proto.RegisterType((*GetCustomerRequest)(nil), "google.ads.googleads.v3.services.GetCustomerRequest")
	proto.RegisterType((*MutateCustomerRequest)(nil), "google.ads.googleads.v3.services.MutateCustomerRequest")
	proto.RegisterType((*CreateCustomerClientRequest)(nil), "google.ads.googleads.v3.services.CreateCustomerClientRequest")
	proto.RegisterType((*CustomerOperation)(nil), "google.ads.googleads.v3.services.CustomerOperation")
	proto.RegisterType((*CreateCustomerClientResponse)(nil), "google.ads.googleads.v3.services.CreateCustomerClientResponse")
	proto.RegisterType((*MutateCustomerResponse)(nil), "google.ads.googleads.v3.services.MutateCustomerResponse")
	proto.RegisterType((*MutateCustomerResult)(nil), "google.ads.googleads.v3.services.MutateCustomerResult")
	proto.RegisterType((*ListAccessibleCustomersRequest)(nil), "google.ads.googleads.v3.services.ListAccessibleCustomersRequest")
	proto.RegisterType((*ListAccessibleCustomersResponse)(nil), "google.ads.googleads.v3.services.ListAccessibleCustomersResponse")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/customer_service.proto", fileDescriptor_62d137bd13ed426f)
}

var fileDescriptor_62d137bd13ed426f = []byte{
	// 922 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x4f, 0x6f, 0xdc, 0x44,
	0x14, 0xd7, 0x38, 0x10, 0x91, 0xd9, 0x24, 0x15, 0xa3, 0x94, 0x2e, 0x9b, 0xa8, 0xd9, 0xba, 0xa9,
	0x48, 0x03, 0xd8, 0x28, 0x8b, 0x48, 0xe5, 0x28, 0x88, 0xc9, 0x0a, 0xb6, 0x48, 0x34, 0xad, 0x5c,
	0x69, 0xa5, 0xc2, 0x4a, 0xab, 0x59, 0x7b, 0xba, 0xb5, 0x6a, 0x7b, 0x8c, 0xc7, 0x5e, 0x54, 0x55,
	0xbd, 0xf0, 0x15, 0x10, 0x5f, 0x80, 0x1b, 0x7c, 0x04, 0xc4, 0x27, 0xa8, 0xe0, 0xc4, 0x2d, 0xa7,
	0x1e, 0x38, 0x71, 0xe0, 0xc0, 0x11, 0x2e, 0xc8, 0xe3, 0x19, 0xaf, 0xf7, 0x8f, 0xbb, 0x6c, 0x7a,
	0x7b, 0x3b, 0xf3, 0x7e, 0xbf, 0xf7, 0x7e, 0xef, 0xbd, 0x79, 0x5e, 0x78, 0x34, 0x64, 0x6c, 0xe8,
	0x53, 0x93, 0xb8, 0xdc, 0xcc, 0xcd, 0xcc, 0x1a, 0xb5, 0x4c, 0x4e, 0xe3, 0x91, 0xe7, 0x50, 0x6e,
	0x3a, 0x29, 0x4f, 0x58, 0x40, 0xe3, 0xbe, 0x3c, 0x31, 0xa2, 0x98, 0x25, 0x0c, 0x35, 0x73, 0x6f,
	0x83, 0xb8, 0xdc, 0x28, 0x80, 0xc6, 0xa8, 0x65, 0x28, 0x60, 0xc3, 0xac, 0xa2, 0xa6, 0x61, 0x1a,
	0x70, 0x93, 0x38, 0x0e, 0xe5, 0xbc, 0x1f, 0x33, 0x5f, 0x52, 0x36, 0x3e, 0xa8, 0x02, 0xc4, 0x94,
	0xb3, 0x34, 0x2e, 0x27, 0x23, 0x11, 0x3b, 0x0a, 0x11, 0x79, 0x26, 0x09, 0x43, 0x96, 0x90, 0xc4,
	0x63, 0x21, 0x97, 0xb7, 0x57, 0x4a, 0xb7, 0x8e, 0xef, 0xd1, 0x30, 0x91, 0x17, 0xbb, 0xa5, 0x8b,
	0x87, 0x1e, 0xf5, 0xdd, 0xfe, 0x80, 0x3e, 0x22, 0x23, 0x8f, 0x29, 0xde, 0xb7, 0x4b, 0x0e, 0x2a,
	0xb8, 0xbc, 0x92, 0xba, 0x4d, 0xf1, 0x6b, 0x90, 0x3e, 0x94, 0x04, 0x01, 0xe1, 0x8f, 0xa5, 0xc7,
	0xd5, 0x69, 0x8f, 0x6f, 0x62, 0x12, 0x45, 0x34, 0x96, 0x69, 0xe9, 0x2e, 0x44, 0x1d, 0x9a, 0xb4,
	0xa5, 0x12, 0x9b, 0x7e, 0x9d, 0x52, 0x9e, 0xa0, 0x33, 0xb8, 0xa1, 0x22, 0xf5, 0x43, 0x12, 0xd0,
	0x3a, 0x68, 0x82, 0xfd, 0xb5, 0xd3, 0x9b, 0x2f, 0xb0, 0xf6, 0x0f, 0xbe, 0x0e, 0xaf, 0x8d, 0x6b,
	0x2c, 0xad, 0xc8, 0xe3, 0x86, 0xc3, 0x02, 0xb3, 0x20, 0x5a, 0x57, 0xf8, 0x33, 0x12, 0x50, 0xfd,
	0x67, 0x00, 0x2f, 0xdf, 0x49, 0x13, 0x92, 0xd0, 0xe9, 0x48, 0x7b, 0xb0, 0x56, 0xf4, 0xd4, 0x73,
	0x65, 0x9c, 0x95, 0x17, 0x58, 0xb3, 0xa1, 0x3a, 0xff, 0xdc, 0x45, 0x5d, 0xb8, 0xc6, 0x22, 0x1a,
	0x8b, 0x82, 0xd6, 0x5f, 0x6b, 0x82, 0xfd, 0xda, 0x61, 0xcb, 0x58, 0xd4, 0x73, 0x43, 0xc5, 0xba,
	0xab, 0xa0, 0x39, 0xf1, 0x98, 0x0a, 0x5d, 0x87, 0x1b, 0x23, 0xe2, 0x7b, 0x2e, 0x49, 0x68, 0x9f,
	0x85, 0xfe, 0x93, 0xfa, 0xeb, 0x4d, 0xb0, 0xff, 0x86, 0xbd, 0xae, 0x0e, 0xef, 0x86, 0xfe, 0x13,
	0xfd, 0x17, 0x0d, 0x6e, 0xb7, 0x63, 0x5a, 0x4a, 0xbe, 0x2d, 0xfa, 0xb7, 0x9c, 0x84, 0x07, 0xf0,
	0x52, 0xe1, 0x95, 0xf7, 0xbf, 0xae, 0x09, 0x21, 0xef, 0x56, 0x0a, 0x29, 0x26, 0xad, 0x50, 0x92,
	0xd3, 0x6e, 0x3a, 0x13, 0x79, 0x20, 0x0c, 0x37, 0x68, 0x40, 0x3c, 0xbf, 0x4f, 0x5c, 0x37, 0xa6,
	0x9c, 0xd7, 0x57, 0x04, 0xf1, 0x8e, 0x22, 0x56, 0xbd, 0x37, 0xee, 0x27, 0xb1, 0x17, 0x0e, 0xbb,
	0xc4, 0x4f, 0xa9, 0xbd, 0x2e, 0x20, 0x38, 0x47, 0xa0, 0x07, 0xb0, 0x56, 0x7a, 0x02, 0xa2, 0xc4,
	0x9b, 0x87, 0xb7, 0x2a, 0x33, 0x13, 0x8f, 0xc6, 0xc0, 0x02, 0x61, 0x33, 0x9f, 0x7e, 0x1a, 0xa6,
	0x41, 0xe9, 0xa7, 0x0d, 0x49, 0x61, 0xeb, 0xdf, 0x03, 0xf8, 0xe6, 0x4c, 0x27, 0x50, 0x1b, 0xae,
	0xa6, 0x51, 0x56, 0x62, 0x51, 0xaf, 0xe5, 0xaa, 0x60, 0x4b, 0x28, 0x3a, 0x86, 0xb5, 0xdc, 0x12,
	0x13, 0x2f, 0xeb, 0xd9, 0x98, 0x91, 0xfd, 0x59, 0xf6, 0x28, 0xee, 0x10, 0xfe, 0xd8, 0x86, 0xb9,
	0x7b, 0x66, 0xeb, 0x6d, 0xb8, 0x33, 0xbf, 0xab, 0x3c, 0x62, 0x21, 0xa7, 0xd9, 0x6c, 0x4c, 0xbe,
	0x81, 0x8c, 0x7e, 0x6d, 0x6a, 0xb0, 0x1f, 0xc1, 0xb7, 0xa6, 0xe7, 0x5a, 0xc2, 0xcf, 0xe0, 0x6a,
	0x4c, 0x79, 0xea, 0xab, 0x36, 0x7f, 0xb4, 0x78, 0x5e, 0x67, 0x98, 0x52, 0x3f, 0xb1, 0x25, 0x8b,
	0x7e, 0x0c, 0xb7, 0xe6, 0xdd, 0xcf, 0xa6, 0x09, 0xe6, 0xa4, 0xd9, 0x84, 0x57, 0xbf, 0xf0, 0x78,
	0x92, 0x77, 0xc8, 0x1b, 0xf8, 0x05, 0x09, 0x97, 0x43, 0xac, 0xdf, 0x86, 0xbb, 0x95, 0x1e, 0x52,
	0xd1, 0x0d, 0xb8, 0x39, 0x11, 0x89, 0xd7, 0x41, 0x73, 0x65, 0x7f, 0xcd, 0xde, 0x28, 0x87, 0xe2,
	0x87, 0x7f, 0xad, 0xc2, 0x4b, 0x0a, 0x7c, 0x3f, 0x97, 0x86, 0x7e, 0x04, 0xb0, 0x56, 0x5a, 0x33,
	0xe8, 0xc3, 0xc5, 0xc5, 0x98, 0xdd, 0x4a, 0x8d, 0x65, 0x66, 0x44, 0x3f, 0x3a, 0xc7, 0x93, 0x85,
	0xf9, 0xf6, 0xf7, 0x3f, 0xbe, 0xd3, 0xae, 0xa1, 0xdd, 0x6c, 0x87, 0x3f, 0x9d, 0xb8, 0x39, 0x51,
	0xef, 0x89, 0x9b, 0x07, 0xcf, 0xd0, 0xaf, 0x00, 0x6e, 0x4e, 0x56, 0x1a, 0x1d, 0x2d, 0xdf, 0xbb,
	0x3c, 0xe3, 0x5b, 0x17, 0x68, 0xba, 0x28, 0xb6, 0xde, 0x39, 0xc7, 0x97, 0x4b, 0x5b, 0xe5, 0xbd,
	0x62, 0x67, 0x09, 0x19, 0x37, 0xf5, 0xbd, 0x4c, 0xc6, 0x38, 0xef, 0xa7, 0x25, 0xdf, 0x93, 0x83,
	0x67, 0x56, 0x20, 0x58, 0x2d, 0x70, 0x80, 0x7e, 0x03, 0xf0, 0x4a, 0x45, 0x67, 0xd1, 0x27, 0x8b,
	0xd3, 0x7b, 0xf9, 0xd8, 0x34, 0xf0, 0x2b, 0x30, 0x48, 0xa5, 0xef, 0x0b, 0x45, 0xef, 0xa0, 0x1b,
	0x13, 0x8a, 0x2c, 0xbf, 0x22, 0xe7, 0xbf, 0x01, 0xdc, 0x9a, 0xf7, 0x6e, 0xd1, 0xc9, 0xff, 0xf8,
	0x20, 0x54, 0x6f, 0xf1, 0xc6, 0xc7, 0x17, 0x85, 0x4b, 0x19, 0x5f, 0x9d, 0xe3, 0xed, 0x72, 0xc3,
	0xa6, 0x96, 0xbd, 0x10, 0x79, 0xa4, 0x1f, 0xbe, 0xbc, 0x6d, 0xce, 0x9c, 0x08, 0x16, 0x38, 0x68,
	0x6c, 0x3f, 0xc7, 0xf5, 0xaa, 0x6f, 0xee, 0xe9, 0xbf, 0x00, 0xee, 0x39, 0x2c, 0x58, 0x98, 0xff,
	0xe9, 0xd6, 0xd4, 0xb3, 0xbc, 0x97, 0x2d, 0xc8, 0x7b, 0xe0, 0xcb, 0xdb, 0x12, 0x39, 0x64, 0x3e,
	0x09, 0x87, 0x06, 0x8b, 0x87, 0xe6, 0x90, 0x86, 0x62, 0x7d, 0x9a, 0xe3, 0x58, 0xd5, 0xff, 0xca,
	0x8e, 0x95, 0xf1, 0x83, 0xb6, 0xd2, 0xc1, 0xf8, 0x27, 0xad, 0xd9, 0xc9, 0x09, 0xb1, 0xcb, 0x8d,
	0xdc, 0xcc, 0xac, 0x6e, 0xcb, 0x90, 0x81, 0xf9, 0x73, 0xe5, 0xd2, 0xc3, 0x2e, 0xef, 0x15, 0x2e,
	0xbd, 0x6e, 0xab, 0xa7, 0x5c, 0xfe, 0xd4, 0xf6, 0xf2, 0x73, 0xcb, 0xc2, 0x2e, 0xb7, 0xac, 0xc2,
	0xc9, 0xb2, 0xba, 0x2d, 0xcb, 0x52, 0x6e, 0x83, 0x55, 0x91, 0x67, 0xeb, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x8d, 0x76, 0xc5, 0x4e, 0x3c, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerServiceClient interface {
	// Returns the requested customer in full detail.
	GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*resources.Customer, error)
	// Updates a customer. Operation statuses are returned.
	MutateCustomer(ctx context.Context, in *MutateCustomerRequest, opts ...grpc.CallOption) (*MutateCustomerResponse, error)
	// Returns resource names of customers directly accessible by the
	// user authenticating the call.
	ListAccessibleCustomers(ctx context.Context, in *ListAccessibleCustomersRequest, opts ...grpc.CallOption) (*ListAccessibleCustomersResponse, error)
	// Creates a new client under manager. The new client customer is returned.
	CreateCustomerClient(ctx context.Context, in *CreateCustomerClientRequest, opts ...grpc.CallOption) (*CreateCustomerClientResponse, error)
}

type customerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceClient(cc grpc.ClientConnInterface) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*resources.Customer, error) {
	out := new(resources.Customer)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CustomerService/GetCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) MutateCustomer(ctx context.Context, in *MutateCustomerRequest, opts ...grpc.CallOption) (*MutateCustomerResponse, error) {
	out := new(MutateCustomerResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CustomerService/MutateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) ListAccessibleCustomers(ctx context.Context, in *ListAccessibleCustomersRequest, opts ...grpc.CallOption) (*ListAccessibleCustomersResponse, error) {
	out := new(ListAccessibleCustomersResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CustomerService/ListAccessibleCustomers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) CreateCustomerClient(ctx context.Context, in *CreateCustomerClientRequest, opts ...grpc.CallOption) (*CreateCustomerClientResponse, error) {
	out := new(CreateCustomerClientResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CustomerService/CreateCustomerClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
type CustomerServiceServer interface {
	// Returns the requested customer in full detail.
	GetCustomer(context.Context, *GetCustomerRequest) (*resources.Customer, error)
	// Updates a customer. Operation statuses are returned.
	MutateCustomer(context.Context, *MutateCustomerRequest) (*MutateCustomerResponse, error)
	// Returns resource names of customers directly accessible by the
	// user authenticating the call.
	ListAccessibleCustomers(context.Context, *ListAccessibleCustomersRequest) (*ListAccessibleCustomersResponse, error)
	// Creates a new client under manager. The new client customer is returned.
	CreateCustomerClient(context.Context, *CreateCustomerClientRequest) (*CreateCustomerClientResponse, error)
}

// UnimplementedCustomerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (*UnimplementedCustomerServiceServer) GetCustomer(ctx context.Context, req *GetCustomerRequest) (*resources.Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomer not implemented")
}
func (*UnimplementedCustomerServiceServer) MutateCustomer(ctx context.Context, req *MutateCustomerRequest) (*MutateCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomer not implemented")
}
func (*UnimplementedCustomerServiceServer) ListAccessibleCustomers(ctx context.Context, req *ListAccessibleCustomersRequest) (*ListAccessibleCustomersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessibleCustomers not implemented")
}
func (*UnimplementedCustomerServiceServer) CreateCustomerClient(ctx context.Context, req *CreateCustomerClientRequest) (*CreateCustomerClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomerClient not implemented")
}

func RegisterCustomerServiceServer(s *grpc.Server, srv CustomerServiceServer) {
	s.RegisterService(&_CustomerService_serviceDesc, srv)
}

func _CustomerService_GetCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CustomerService/GetCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetCustomer(ctx, req.(*GetCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_MutateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).MutateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CustomerService/MutateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).MutateCustomer(ctx, req.(*MutateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_ListAccessibleCustomers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccessibleCustomersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).ListAccessibleCustomers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CustomerService/ListAccessibleCustomers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).ListAccessibleCustomers(ctx, req.(*ListAccessibleCustomersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_CreateCustomerClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomerClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).CreateCustomerClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CustomerService/CreateCustomerClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).CreateCustomerClient(ctx, req.(*CreateCustomerClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomer",
			Handler:    _CustomerService_GetCustomer_Handler,
		},
		{
			MethodName: "MutateCustomer",
			Handler:    _CustomerService_MutateCustomer_Handler,
		},
		{
			MethodName: "ListAccessibleCustomers",
			Handler:    _CustomerService_ListAccessibleCustomers_Handler,
		},
		{
			MethodName: "CreateCustomerClient",
			Handler:    _CustomerService_CreateCustomerClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/customer_service.proto",
}
