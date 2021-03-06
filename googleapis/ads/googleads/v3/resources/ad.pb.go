// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/ad.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v3/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
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

// An ad.
type Ad struct {
	// Immutable. The resource name of the ad.
	// Ad resource names have the form:
	//
	// `customers/{customer_id}/ads/{ad_id}`
	ResourceName string `protobuf:"bytes,37,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the ad.
	Id *wrappers.Int64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The list of possible final URLs after all cross-domain redirects for the
	// ad.
	FinalUrls []*wrappers.StringValue `protobuf:"bytes,2,rep,name=final_urls,json=finalUrls,proto3" json:"final_urls,omitempty"`
	// A list of final app URLs that will be used on mobile if the user has the
	// specific app installed.
	FinalAppUrls []*common.FinalAppUrl `protobuf:"bytes,35,rep,name=final_app_urls,json=finalAppUrls,proto3" json:"final_app_urls,omitempty"`
	// The list of possible final mobile URLs after all cross-domain redirects
	// for the ad.
	FinalMobileUrls []*wrappers.StringValue `protobuf:"bytes,16,rep,name=final_mobile_urls,json=finalMobileUrls,proto3" json:"final_mobile_urls,omitempty"`
	// The URL template for constructing a tracking URL.
	TrackingUrlTemplate *wrappers.StringValue `protobuf:"bytes,12,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3" json:"tracking_url_template,omitempty"`
	// The suffix to use when constructing a final URL.
	FinalUrlSuffix *wrappers.StringValue `protobuf:"bytes,38,opt,name=final_url_suffix,json=finalUrlSuffix,proto3" json:"final_url_suffix,omitempty"`
	// The list of mappings that can be used to substitute custom parameter tags
	// in a `tracking_url_template`, `final_urls`, or `mobile_final_urls`.
	UrlCustomParameters []*common.CustomParameter `protobuf:"bytes,10,rep,name=url_custom_parameters,json=urlCustomParameters,proto3" json:"url_custom_parameters,omitempty"`
	// The URL that appears in the ad description for some ad formats.
	DisplayUrl *wrappers.StringValue `protobuf:"bytes,4,opt,name=display_url,json=displayUrl,proto3" json:"display_url,omitempty"`
	// Output only. The type of ad.
	Type enums.AdTypeEnum_AdType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v3.enums.AdTypeEnum_AdType" json:"type,omitempty"`
	// Output only. Indicates if this ad was automatically added by Google Ads and not by a
	// user. For example, this could happen when ads are automatically created as
	// suggestions for new ads based on knowledge of how existing ads are
	// performing.
	AddedByGoogleAds *wrappers.BoolValue `protobuf:"bytes,19,opt,name=added_by_google_ads,json=addedByGoogleAds,proto3" json:"added_by_google_ads,omitempty"`
	// The device preference for the ad. You can only specify a preference for
	// mobile devices. When this preference is set the ad will be preferred over
	// other ads when being displayed on a mobile device. The ad can still be
	// displayed on other device types, e.g. if no other ads are available.
	// If unspecified (no device preference), all devices are targeted.
	// This is only supported by some ad types.
	DevicePreference enums.DeviceEnum_Device `protobuf:"varint,20,opt,name=device_preference,json=devicePreference,proto3,enum=google.ads.googleads.v3.enums.DeviceEnum_Device" json:"device_preference,omitempty"`
	// Additional URLs for the ad that are tagged with a unique identifier that
	// can be referenced from other fields in the ad.
	UrlCollections []*common.UrlCollection `protobuf:"bytes,26,rep,name=url_collections,json=urlCollections,proto3" json:"url_collections,omitempty"`
	// Immutable. The name of the ad. This is only used to be able to identify the ad. It
	// does not need to be unique and does not affect the served ad.
	Name *wrappers.StringValue `protobuf:"bytes,23,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. If this ad is system managed, then this field will indicate the source.
	// This field is read-only.
	SystemManagedResourceSource enums.SystemManagedResourceSourceEnum_SystemManagedResourceSource `protobuf:"varint,27,opt,name=system_managed_resource_source,json=systemManagedResourceSource,proto3,enum=google.ads.googleads.v3.enums.SystemManagedResourceSourceEnum_SystemManagedResourceSource" json:"system_managed_resource_source,omitempty"`
	// Details pertinent to the ad type. Exactly one value must be set.
	//
	// Types that are valid to be assigned to AdData:
	//	*Ad_TextAd
	//	*Ad_ExpandedTextAd
	//	*Ad_CallOnlyAd
	//	*Ad_ExpandedDynamicSearchAd
	//	*Ad_HotelAd
	//	*Ad_ShoppingSmartAd
	//	*Ad_ShoppingProductAd
	//	*Ad_GmailAd
	//	*Ad_ImageAd
	//	*Ad_VideoAd
	//	*Ad_ResponsiveSearchAd
	//	*Ad_LegacyResponsiveDisplayAd
	//	*Ad_AppAd
	//	*Ad_LegacyAppInstallAd
	//	*Ad_ResponsiveDisplayAd
	//	*Ad_DisplayUploadAd
	//	*Ad_AppEngagementAd
	//	*Ad_ShoppingComparisonListingAd
	AdData               isAd_AdData `protobuf_oneof:"ad_data"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Ad) Reset()         { *m = Ad{} }
func (m *Ad) String() string { return proto.CompactTextString(m) }
func (*Ad) ProtoMessage()    {}
func (*Ad) Descriptor() ([]byte, []int) {
	return fileDescriptor_65b02bf52a660238, []int{0}
}

func (m *Ad) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ad.Unmarshal(m, b)
}
func (m *Ad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ad.Marshal(b, m, deterministic)
}
func (m *Ad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ad.Merge(m, src)
}
func (m *Ad) XXX_Size() int {
	return xxx_messageInfo_Ad.Size(m)
}
func (m *Ad) XXX_DiscardUnknown() {
	xxx_messageInfo_Ad.DiscardUnknown(m)
}

var xxx_messageInfo_Ad proto.InternalMessageInfo

func (m *Ad) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Ad) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Ad) GetFinalUrls() []*wrappers.StringValue {
	if m != nil {
		return m.FinalUrls
	}
	return nil
}

func (m *Ad) GetFinalAppUrls() []*common.FinalAppUrl {
	if m != nil {
		return m.FinalAppUrls
	}
	return nil
}

func (m *Ad) GetFinalMobileUrls() []*wrappers.StringValue {
	if m != nil {
		return m.FinalMobileUrls
	}
	return nil
}

func (m *Ad) GetTrackingUrlTemplate() *wrappers.StringValue {
	if m != nil {
		return m.TrackingUrlTemplate
	}
	return nil
}

func (m *Ad) GetFinalUrlSuffix() *wrappers.StringValue {
	if m != nil {
		return m.FinalUrlSuffix
	}
	return nil
}

func (m *Ad) GetUrlCustomParameters() []*common.CustomParameter {
	if m != nil {
		return m.UrlCustomParameters
	}
	return nil
}

func (m *Ad) GetDisplayUrl() *wrappers.StringValue {
	if m != nil {
		return m.DisplayUrl
	}
	return nil
}

func (m *Ad) GetType() enums.AdTypeEnum_AdType {
	if m != nil {
		return m.Type
	}
	return enums.AdTypeEnum_UNSPECIFIED
}

func (m *Ad) GetAddedByGoogleAds() *wrappers.BoolValue {
	if m != nil {
		return m.AddedByGoogleAds
	}
	return nil
}

func (m *Ad) GetDevicePreference() enums.DeviceEnum_Device {
	if m != nil {
		return m.DevicePreference
	}
	return enums.DeviceEnum_UNSPECIFIED
}

func (m *Ad) GetUrlCollections() []*common.UrlCollection {
	if m != nil {
		return m.UrlCollections
	}
	return nil
}

func (m *Ad) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Ad) GetSystemManagedResourceSource() enums.SystemManagedResourceSourceEnum_SystemManagedResourceSource {
	if m != nil {
		return m.SystemManagedResourceSource
	}
	return enums.SystemManagedResourceSourceEnum_UNSPECIFIED
}

type isAd_AdData interface {
	isAd_AdData()
}

type Ad_TextAd struct {
	TextAd *common.TextAdInfo `protobuf:"bytes,6,opt,name=text_ad,json=textAd,proto3,oneof"`
}

type Ad_ExpandedTextAd struct {
	ExpandedTextAd *common.ExpandedTextAdInfo `protobuf:"bytes,7,opt,name=expanded_text_ad,json=expandedTextAd,proto3,oneof"`
}

type Ad_CallOnlyAd struct {
	CallOnlyAd *common.CallOnlyAdInfo `protobuf:"bytes,13,opt,name=call_only_ad,json=callOnlyAd,proto3,oneof"`
}

type Ad_ExpandedDynamicSearchAd struct {
	ExpandedDynamicSearchAd *common.ExpandedDynamicSearchAdInfo `protobuf:"bytes,14,opt,name=expanded_dynamic_search_ad,json=expandedDynamicSearchAd,proto3,oneof"`
}

type Ad_HotelAd struct {
	HotelAd *common.HotelAdInfo `protobuf:"bytes,15,opt,name=hotel_ad,json=hotelAd,proto3,oneof"`
}

type Ad_ShoppingSmartAd struct {
	ShoppingSmartAd *common.ShoppingSmartAdInfo `protobuf:"bytes,17,opt,name=shopping_smart_ad,json=shoppingSmartAd,proto3,oneof"`
}

type Ad_ShoppingProductAd struct {
	ShoppingProductAd *common.ShoppingProductAdInfo `protobuf:"bytes,18,opt,name=shopping_product_ad,json=shoppingProductAd,proto3,oneof"`
}

type Ad_GmailAd struct {
	GmailAd *common.GmailAdInfo `protobuf:"bytes,21,opt,name=gmail_ad,json=gmailAd,proto3,oneof"`
}

type Ad_ImageAd struct {
	ImageAd *common.ImageAdInfo `protobuf:"bytes,22,opt,name=image_ad,json=imageAd,proto3,oneof"`
}

type Ad_VideoAd struct {
	VideoAd *common.VideoAdInfo `protobuf:"bytes,24,opt,name=video_ad,json=videoAd,proto3,oneof"`
}

type Ad_ResponsiveSearchAd struct {
	ResponsiveSearchAd *common.ResponsiveSearchAdInfo `protobuf:"bytes,25,opt,name=responsive_search_ad,json=responsiveSearchAd,proto3,oneof"`
}

type Ad_LegacyResponsiveDisplayAd struct {
	LegacyResponsiveDisplayAd *common.LegacyResponsiveDisplayAdInfo `protobuf:"bytes,28,opt,name=legacy_responsive_display_ad,json=legacyResponsiveDisplayAd,proto3,oneof"`
}

type Ad_AppAd struct {
	AppAd *common.AppAdInfo `protobuf:"bytes,29,opt,name=app_ad,json=appAd,proto3,oneof"`
}

type Ad_LegacyAppInstallAd struct {
	LegacyAppInstallAd *common.LegacyAppInstallAdInfo `protobuf:"bytes,30,opt,name=legacy_app_install_ad,json=legacyAppInstallAd,proto3,oneof"`
}

type Ad_ResponsiveDisplayAd struct {
	ResponsiveDisplayAd *common.ResponsiveDisplayAdInfo `protobuf:"bytes,31,opt,name=responsive_display_ad,json=responsiveDisplayAd,proto3,oneof"`
}

type Ad_DisplayUploadAd struct {
	DisplayUploadAd *common.DisplayUploadAdInfo `protobuf:"bytes,33,opt,name=display_upload_ad,json=displayUploadAd,proto3,oneof"`
}

type Ad_AppEngagementAd struct {
	AppEngagementAd *common.AppEngagementAdInfo `protobuf:"bytes,34,opt,name=app_engagement_ad,json=appEngagementAd,proto3,oneof"`
}

type Ad_ShoppingComparisonListingAd struct {
	ShoppingComparisonListingAd *common.ShoppingComparisonListingAdInfo `protobuf:"bytes,36,opt,name=shopping_comparison_listing_ad,json=shoppingComparisonListingAd,proto3,oneof"`
}

func (*Ad_TextAd) isAd_AdData() {}

func (*Ad_ExpandedTextAd) isAd_AdData() {}

func (*Ad_CallOnlyAd) isAd_AdData() {}

func (*Ad_ExpandedDynamicSearchAd) isAd_AdData() {}

func (*Ad_HotelAd) isAd_AdData() {}

func (*Ad_ShoppingSmartAd) isAd_AdData() {}

func (*Ad_ShoppingProductAd) isAd_AdData() {}

func (*Ad_GmailAd) isAd_AdData() {}

func (*Ad_ImageAd) isAd_AdData() {}

func (*Ad_VideoAd) isAd_AdData() {}

func (*Ad_ResponsiveSearchAd) isAd_AdData() {}

func (*Ad_LegacyResponsiveDisplayAd) isAd_AdData() {}

func (*Ad_AppAd) isAd_AdData() {}

func (*Ad_LegacyAppInstallAd) isAd_AdData() {}

func (*Ad_ResponsiveDisplayAd) isAd_AdData() {}

func (*Ad_DisplayUploadAd) isAd_AdData() {}

func (*Ad_AppEngagementAd) isAd_AdData() {}

func (*Ad_ShoppingComparisonListingAd) isAd_AdData() {}

func (m *Ad) GetAdData() isAd_AdData {
	if m != nil {
		return m.AdData
	}
	return nil
}

func (m *Ad) GetTextAd() *common.TextAdInfo {
	if x, ok := m.GetAdData().(*Ad_TextAd); ok {
		return x.TextAd
	}
	return nil
}

func (m *Ad) GetExpandedTextAd() *common.ExpandedTextAdInfo {
	if x, ok := m.GetAdData().(*Ad_ExpandedTextAd); ok {
		return x.ExpandedTextAd
	}
	return nil
}

func (m *Ad) GetCallOnlyAd() *common.CallOnlyAdInfo {
	if x, ok := m.GetAdData().(*Ad_CallOnlyAd); ok {
		return x.CallOnlyAd
	}
	return nil
}

func (m *Ad) GetExpandedDynamicSearchAd() *common.ExpandedDynamicSearchAdInfo {
	if x, ok := m.GetAdData().(*Ad_ExpandedDynamicSearchAd); ok {
		return x.ExpandedDynamicSearchAd
	}
	return nil
}

func (m *Ad) GetHotelAd() *common.HotelAdInfo {
	if x, ok := m.GetAdData().(*Ad_HotelAd); ok {
		return x.HotelAd
	}
	return nil
}

func (m *Ad) GetShoppingSmartAd() *common.ShoppingSmartAdInfo {
	if x, ok := m.GetAdData().(*Ad_ShoppingSmartAd); ok {
		return x.ShoppingSmartAd
	}
	return nil
}

func (m *Ad) GetShoppingProductAd() *common.ShoppingProductAdInfo {
	if x, ok := m.GetAdData().(*Ad_ShoppingProductAd); ok {
		return x.ShoppingProductAd
	}
	return nil
}

func (m *Ad) GetGmailAd() *common.GmailAdInfo {
	if x, ok := m.GetAdData().(*Ad_GmailAd); ok {
		return x.GmailAd
	}
	return nil
}

func (m *Ad) GetImageAd() *common.ImageAdInfo {
	if x, ok := m.GetAdData().(*Ad_ImageAd); ok {
		return x.ImageAd
	}
	return nil
}

func (m *Ad) GetVideoAd() *common.VideoAdInfo {
	if x, ok := m.GetAdData().(*Ad_VideoAd); ok {
		return x.VideoAd
	}
	return nil
}

func (m *Ad) GetResponsiveSearchAd() *common.ResponsiveSearchAdInfo {
	if x, ok := m.GetAdData().(*Ad_ResponsiveSearchAd); ok {
		return x.ResponsiveSearchAd
	}
	return nil
}

func (m *Ad) GetLegacyResponsiveDisplayAd() *common.LegacyResponsiveDisplayAdInfo {
	if x, ok := m.GetAdData().(*Ad_LegacyResponsiveDisplayAd); ok {
		return x.LegacyResponsiveDisplayAd
	}
	return nil
}

func (m *Ad) GetAppAd() *common.AppAdInfo {
	if x, ok := m.GetAdData().(*Ad_AppAd); ok {
		return x.AppAd
	}
	return nil
}

func (m *Ad) GetLegacyAppInstallAd() *common.LegacyAppInstallAdInfo {
	if x, ok := m.GetAdData().(*Ad_LegacyAppInstallAd); ok {
		return x.LegacyAppInstallAd
	}
	return nil
}

func (m *Ad) GetResponsiveDisplayAd() *common.ResponsiveDisplayAdInfo {
	if x, ok := m.GetAdData().(*Ad_ResponsiveDisplayAd); ok {
		return x.ResponsiveDisplayAd
	}
	return nil
}

func (m *Ad) GetDisplayUploadAd() *common.DisplayUploadAdInfo {
	if x, ok := m.GetAdData().(*Ad_DisplayUploadAd); ok {
		return x.DisplayUploadAd
	}
	return nil
}

func (m *Ad) GetAppEngagementAd() *common.AppEngagementAdInfo {
	if x, ok := m.GetAdData().(*Ad_AppEngagementAd); ok {
		return x.AppEngagementAd
	}
	return nil
}

func (m *Ad) GetShoppingComparisonListingAd() *common.ShoppingComparisonListingAdInfo {
	if x, ok := m.GetAdData().(*Ad_ShoppingComparisonListingAd); ok {
		return x.ShoppingComparisonListingAd
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Ad) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Ad_TextAd)(nil),
		(*Ad_ExpandedTextAd)(nil),
		(*Ad_CallOnlyAd)(nil),
		(*Ad_ExpandedDynamicSearchAd)(nil),
		(*Ad_HotelAd)(nil),
		(*Ad_ShoppingSmartAd)(nil),
		(*Ad_ShoppingProductAd)(nil),
		(*Ad_GmailAd)(nil),
		(*Ad_ImageAd)(nil),
		(*Ad_VideoAd)(nil),
		(*Ad_ResponsiveSearchAd)(nil),
		(*Ad_LegacyResponsiveDisplayAd)(nil),
		(*Ad_AppAd)(nil),
		(*Ad_LegacyAppInstallAd)(nil),
		(*Ad_ResponsiveDisplayAd)(nil),
		(*Ad_DisplayUploadAd)(nil),
		(*Ad_AppEngagementAd)(nil),
		(*Ad_ShoppingComparisonListingAd)(nil),
	}
}

func init() {
	proto.RegisterType((*Ad)(nil), "google.ads.googleads.v3.resources.Ad")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/ad.proto", fileDescriptor_65b02bf52a660238)
}

var fileDescriptor_65b02bf52a660238 = []byte{
	// 1342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x57, 0x4d, 0x6f, 0x14, 0x37,
	0x1f, 0x7f, 0x36, 0x21, 0xc9, 0x83, 0x09, 0x79, 0x71, 0xc8, 0xc3, 0x90, 0x04, 0x08, 0xf0, 0x50,
	0xa5, 0xa0, 0xce, 0xa2, 0x6c, 0xa1, 0xd2, 0x52, 0x44, 0x67, 0x21, 0x90, 0x54, 0x40, 0xd3, 0x0d,
	0xc9, 0x01, 0xa5, 0x1d, 0x39, 0x63, 0xef, 0x64, 0x5a, 0x8f, 0x6d, 0xd9, 0xb3, 0xdb, 0x6c, 0x11,
	0x52, 0x4f, 0xfd, 0x20, 0x55, 0x4f, 0xfd, 0x28, 0xfd, 0x14, 0xa8, 0x47, 0xce, 0x3d, 0xf5, 0x54,
	0xd9, 0x9e, 0xf1, 0x66, 0x13, 0x36, 0x33, 0x97, 0xc4, 0x2f, 0xbf, 0x97, 0xbf, 0xff, 0xb6, 0xff,
	0xe3, 0x05, 0x77, 0x62, 0xce, 0x63, 0x4a, 0xea, 0x08, 0xab, 0xba, 0x6d, 0xea, 0x56, 0xaf, 0x51,
	0x97, 0x44, 0xf1, 0xae, 0x8c, 0x88, 0xaa, 0x23, 0xec, 0x0b, 0xc9, 0x33, 0x0e, 0x6f, 0x58, 0x80,
	0x8f, 0xb0, 0xf2, 0x1d, 0xd6, 0xef, 0x35, 0x7c, 0x87, 0x5d, 0x5a, 0x1f, 0x25, 0x17, 0xf1, 0x34,
	0xe5, 0xac, 0x8e, 0x70, 0x98, 0xf5, 0x05, 0x09, 0x13, 0xd6, 0xe1, 0xca, 0xca, 0x2e, 0xdd, 0x2f,
	0xe1, 0x44, 0x5d, 0x95, 0xf1, 0x34, 0x14, 0x48, 0xa2, 0x94, 0x64, 0x44, 0xe6, 0xb4, 0x32, 0xab,
	0x4e, 0xc2, 0x10, 0x0d, 0x91, 0x10, 0x61, 0x57, 0xd2, 0x9c, 0xd3, 0x28, 0xe1, 0x74, 0x25, 0x0d,
	0x23, 0x4e, 0x29, 0x89, 0xb2, 0x84, 0xb3, 0x9c, 0x74, 0x77, 0x14, 0x89, 0xb0, 0x6e, 0xaa, 0x8a,
	0x25, 0xe5, 0xe0, 0x3b, 0x67, 0x83, 0x31, 0xe9, 0x25, 0x51, 0x81, 0xfd, 0xea, 0x6c, 0xac, 0xea,
	0xab, 0x8c, 0xa4, 0x61, 0x8a, 0x18, 0x8a, 0x09, 0x0e, 0x09, 0xcb, 0x92, 0xac, 0x1f, 0xda, 0x4c,
	0xe7, 0x0a, 0xd7, 0x0b, 0x05, 0x91, 0xd4, 0x3b, 0x09, 0xa1, 0x38, 0x3c, 0x20, 0x87, 0xa8, 0x97,
	0xf0, 0x22, 0x49, 0x57, 0x8e, 0x01, 0x8a, 0x5d, 0xca, 0xa7, 0xae, 0xe5, 0x53, 0xa6, 0x77, 0xd0,
	0xed, 0xd4, 0x7f, 0x92, 0x48, 0x08, 0x22, 0x8b, 0x6d, 0x59, 0x39, 0x46, 0x45, 0x8c, 0xf1, 0x0c,
	0xe9, 0x9c, 0xe4, 0xb3, 0x37, 0xff, 0xba, 0x0c, 0xc6, 0x02, 0x0c, 0x37, 0xc1, 0xc5, 0x42, 0x36,
	0x64, 0x28, 0x25, 0xde, 0xed, 0xd5, 0xda, 0xda, 0xf9, 0xd6, 0xad, 0xf7, 0xc1, 0xc4, 0x3f, 0xc1,
	0x55, 0xb0, 0x3c, 0x38, 0x26, 0x79, 0x4b, 0x24, 0xca, 0x8f, 0x78, 0x5a, 0x0f, 0x70, 0x7b, 0xba,
	0x60, 0xbe, 0x42, 0x29, 0x81, 0xf7, 0xc0, 0x58, 0x82, 0xbd, 0xda, 0x6a, 0x6d, 0xed, 0xc2, 0xfa,
	0x72, 0x8e, 0xf6, 0x8b, 0xd8, 0xfc, 0x2d, 0x96, 0x3d, 0xf8, 0x7c, 0x0f, 0xd1, 0x2e, 0x69, 0x8d,
	0xbf, 0x0f, 0xc6, 0xdb, 0x63, 0x09, 0x86, 0x0f, 0x01, 0xb0, 0x7b, 0xdc, 0x95, 0x54, 0x79, 0x63,
	0xab, 0xe3, 0x6b, 0x17, 0xd6, 0x57, 0x4e, 0x31, 0x77, 0x32, 0x99, 0xb0, 0xd8, 0x50, 0xdb, 0xe7,
	0x0d, 0x7e, 0x57, 0x52, 0x05, 0xbf, 0x05, 0x33, 0x43, 0x07, 0x44, 0x79, 0xb7, 0x8c, 0xc0, 0x5d,
	0x7f, 0xd4, 0x21, 0xb7, 0x47, 0xc4, 0x7f, 0xa6, 0x59, 0x81, 0x10, 0xbb, 0x92, 0xb6, 0xa7, 0x3b,
	0x83, 0x8e, 0x82, 0x9b, 0x60, 0xde, 0x4a, 0xa6, 0xfc, 0x20, 0xa1, 0xc4, 0xaa, 0xce, 0x55, 0x08,
	0x6b, 0xd6, 0xd0, 0x5e, 0x1a, 0x96, 0x51, 0xda, 0x06, 0x8b, 0x99, 0x44, 0xd1, 0x8f, 0x09, 0x8b,
	0xb5, 0x4a, 0x98, 0x91, 0x54, 0x50, 0x94, 0x11, 0x6f, 0xda, 0xa4, 0xe7, 0x6c, 0xb5, 0x85, 0x82,
	0xba, 0x2b, 0xe9, 0xeb, 0x9c, 0x08, 0x9f, 0x81, 0x39, 0x97, 0xab, 0x50, 0x75, 0x3b, 0x9d, 0xe4,
	0xc8, 0xfb, 0xa4, 0x82, 0xd8, 0x4c, 0x91, 0xb1, 0x1d, 0xc3, 0x81, 0x11, 0x58, 0x34, 0x77, 0xe4,
	0xc4, 0x95, 0x54, 0x1e, 0x30, 0xeb, 0xac, 0x97, 0x65, 0xef, 0x89, 0x21, 0x6e, 0x17, 0xbc, 0xf6,
	0x42, 0x57, 0xd2, 0x13, 0x63, 0x0a, 0x3e, 0x02, 0x17, 0x70, 0xa2, 0x04, 0x45, 0x7d, 0x1d, 0xae,
	0x77, 0xae, 0x42, 0x9c, 0x20, 0x27, 0xec, 0x4a, 0x0a, 0xb7, 0xc0, 0x39, 0x7d, 0x21, 0xbd, 0x89,
	0xd5, 0xda, 0xda, 0xcc, 0xfa, 0xbd, 0x91, 0x21, 0x99, 0x5b, 0xe6, 0x07, 0xf8, 0x75, 0x5f, 0x90,
	0x0d, 0xd6, 0x4d, 0xf3, 0xa6, 0x3d, 0x60, 0x46, 0x02, 0xbe, 0x02, 0x0b, 0x08, 0x63, 0x82, 0xc3,
	0x83, 0x7e, 0x68, 0xb9, 0x21, 0xc2, 0xca, 0x5b, 0x30, 0x11, 0x2d, 0x9d, 0x8a, 0xa8, 0xc5, 0x39,
	0x3d, 0x76, 0x48, 0xe7, 0x0c, 0xb7, 0xd5, 0x7f, 0x6e, 0x60, 0x01, 0x56, 0xf0, 0x3b, 0x30, 0x6f,
	0x2b, 0x40, 0x28, 0x24, 0xe9, 0x10, 0x49, 0x58, 0x44, 0xbc, 0x4b, 0x95, 0xe2, 0x7c, 0x6a, 0x78,
	0x26, 0x4e, 0xdb, 0x6c, 0xcf, 0x59, 0xa9, 0x6d, 0xa7, 0x04, 0xf7, 0xc0, 0xec, 0x70, 0x05, 0x53,
	0xde, 0x92, 0xd9, 0x97, 0xcf, 0xca, 0xf6, 0x65, 0x57, 0xd2, 0x27, 0x8e, 0xd5, 0x9e, 0xe9, 0x1e,
	0xef, 0x2a, 0x78, 0x1f, 0x9c, 0x33, 0x97, 0xfb, 0x72, 0xf9, 0x4e, 0xe8, 0x95, 0x4f, 0xb4, 0x0d,
	0x1c, 0xfe, 0x5e, 0x03, 0xd7, 0x4e, 0x14, 0x31, 0x57, 0x2c, 0xec, 0x3f, 0x6f, 0xd9, 0xac, 0xfd,
	0x4d, 0xc9, 0xda, 0x77, 0x8c, 0xc8, 0x4b, 0xab, 0xd1, 0xce, 0x25, 0x76, 0xcc, 0x5f, 0x93, 0x90,
	0x33, 0xe6, 0xed, 0x4e, 0x2c, 0xab, 0xd1, 0x08, 0xb8, 0x01, 0xa6, 0x32, 0x72, 0x94, 0x85, 0x08,
	0x7b, 0x93, 0x66, 0x81, 0x77, 0xca, 0xb2, 0xf5, 0x9a, 0x1c, 0x65, 0x01, 0xde, 0x62, 0x1d, 0xbe,
	0xf9, 0x9f, 0xf6, 0x64, 0x66, 0x7a, 0xf0, 0x7b, 0x30, 0x47, 0x8e, 0x04, 0x62, 0xfa, 0xb8, 0x14,
	0x7a, 0x53, 0x46, 0x6f, 0xbd, 0x4c, 0x6f, 0x23, 0xe7, 0x0d, 0xe9, 0xce, 0x90, 0xa1, 0x51, 0xd8,
	0x06, 0xd3, 0x11, 0xa2, 0x34, 0xe4, 0x8c, 0xf6, 0xb5, 0xf6, 0x45, 0xa3, 0xed, 0x97, 0xde, 0x38,
	0x44, 0xe9, 0x37, 0x8c, 0xf6, 0x9d, 0x2e, 0x88, 0xdc, 0x08, 0xfc, 0x19, 0x2c, 0xb9, 0x98, 0x71,
	0x9f, 0xa1, 0x34, 0x89, 0x42, 0x45, 0x90, 0x8c, 0x0e, 0xb5, 0xc3, 0x8c, 0x71, 0x78, 0x58, 0x35,
	0xfa, 0xa7, 0x56, 0x60, 0xc7, 0xf0, 0x9d, 0xdd, 0x65, 0xf2, 0xf1, 0x69, 0xb8, 0x09, 0xfe, 0x7b,
	0xc8, 0x33, 0x42, 0xb5, 0xd3, 0xac, 0x71, 0x2a, 0xad, 0xbd, 0x9b, 0x1a, 0xef, 0x94, 0xa7, 0x0e,
	0x6d, 0x17, 0x22, 0x30, 0xaf, 0x0e, 0xb9, 0x10, 0xba, 0x5c, 0xaa, 0x14, 0x49, 0x93, 0xfa, 0x79,
	0x23, 0xd9, 0x28, 0x93, 0xdc, 0xc9, 0x89, 0x3b, 0x9a, 0xe7, 0xa4, 0x67, 0xd5, 0xf0, 0x30, 0x8c,
	0xc1, 0x82, 0xb3, 0x10, 0x92, 0xe3, 0x6e, 0x64, 0x4c, 0xa0, 0x31, 0xb9, 0x5f, 0xd5, 0x64, 0xdb,
	0x32, 0x9d, 0x8d, 0x0b, 0xdb, 0x4d, 0xe8, 0xac, 0xc4, 0x29, 0x4a, 0x4c, 0x56, 0x16, 0xab, 0x65,
	0xe5, 0xb9, 0xc6, 0x0f, 0xb2, 0x12, 0xdb, 0xae, 0x56, 0x4a, 0x52, 0x14, 0xeb, 0x8a, 0xe5, 0xfd,
	0xaf, 0x9a, 0xd2, 0x96, 0xc6, 0x0f, 0x94, 0x12, 0xdb, 0xd5, 0x4a, 0xbd, 0x04, 0x13, 0xae, 0x95,
	0xbc, 0x6a, 0x4a, 0x7b, 0x1a, 0x3f, 0x50, 0xea, 0xd9, 0x2e, 0xfc, 0x01, 0x5c, 0x92, 0x44, 0x09,
	0xce, 0x54, 0xd2, 0x23, 0xc7, 0x4e, 0xda, 0x15, 0xa3, 0xfa, 0xa0, 0x4c, 0xb5, 0xed, 0xb8, 0x27,
	0x0e, 0x19, 0x94, 0xa7, 0x66, 0xe0, 0x2f, 0x35, 0xb0, 0x42, 0x49, 0x8c, 0xa2, 0x7e, 0x78, 0xcc,
	0xb3, 0xf8, 0xb0, 0x20, 0xec, 0xad, 0x18, 0xd3, 0x47, 0x65, 0xa6, 0x2f, 0x8c, 0xc6, 0xc0, 0xfa,
	0xa9, 0x55, 0x70, 0xde, 0x57, 0xe8, 0x28, 0x00, 0x6c, 0x81, 0x49, 0xfd, 0xbc, 0x40, 0xd8, 0xbb,
	0x6a, 0xbc, 0x3e, 0x2d, 0xf3, 0x0a, 0x84, 0x70, 0xba, 0x13, 0x48, 0x77, 0xa0, 0x00, 0x8b, 0xf9,
	0x2a, 0xb4, 0x54, 0xc2, 0x54, 0xa6, 0x8b, 0x00, 0xc2, 0xde, 0xb5, 0x6a, 0x39, 0xb3, 0xe1, 0x07,
	0x42, 0x6c, 0x59, 0xaa, 0xd5, 0x37, 0x65, 0x5a, 0x27, 0x8e, 0x9e, 0x9a, 0x86, 0x29, 0x58, 0xfc,
	0x78, 0xc2, 0xae, 0x1b, 0xc7, 0x2f, 0xaa, 0xef, 0xd2, 0xc9, 0x54, 0x2d, 0xc8, 0x8f, 0x24, 0x09,
	0x81, 0x79, 0xf7, 0xb5, 0x17, 0x94, 0x23, 0xac, 0xad, 0x6e, 0x54, 0xbb, 0xbd, 0xb9, 0xca, 0xae,
	0xe1, 0x0d, 0x6e, 0x2f, 0x1e, 0x1e, 0xd6, 0x16, 0x3a, 0x79, 0x84, 0xc5, 0x28, 0x26, 0x29, 0x61,
	0xe6, 0xee, 0xde, 0xac, 0x66, 0x11, 0x08, 0xb1, 0xe1, 0x78, 0x03, 0x0b, 0x34, 0x3c, 0x0c, 0x7f,
	0xd5, 0xdf, 0xba, 0xa2, 0x42, 0x44, 0x3c, 0x15, 0x48, 0x26, 0x8a, 0xb3, 0x90, 0x26, 0x2a, 0xd3,
	0x43, 0x08, 0x7b, 0xff, 0x37, 0x86, 0x8f, 0xab, 0x16, 0x8b, 0x27, 0x4e, 0xe4, 0x85, 0xd5, 0x70,
	0xe6, 0xcb, 0x6a, 0x34, 0xa4, 0xf9, 0xf8, 0x43, 0xf0, 0xe5, 0x99, 0xef, 0x6e, 0x78, 0xd5, 0xbe,
	0xde, 0x88, 0x54, 0xf5, 0xb7, 0x45, 0xf3, 0x9d, 0xf9, 0x11, 0xf2, 0x16, 0xe1, 0x77, 0xad, 0xf3,
	0x60, 0x0a, 0xe1, 0x10, 0xa3, 0x0c, 0xb5, 0xfe, 0xae, 0x81, 0xdb, 0x11, 0x4f, 0xfd, 0xd2, 0xdf,
	0x7d, 0xad, 0xa9, 0x00, 0x6f, 0xeb, 0xd7, 0xc0, 0x76, 0xed, 0xcd, 0xd7, 0x39, 0x3a, 0xe6, 0x14,
	0xb1, 0xd8, 0xe7, 0x32, 0xae, 0xc7, 0x84, 0x99, 0xb7, 0x42, 0x7d, 0x10, 0xcc, 0x19, 0x3f, 0x38,
	0x1f, 0xba, 0xd6, 0x6f, 0x63, 0xe3, 0xcf, 0x83, 0xe0, 0x8f, 0xb1, 0x1b, 0xf6, 0x05, 0xe5, 0x07,
	0x58, 0xf9, 0xee, 0x31, 0xe5, 0xef, 0x35, 0xfc, 0xe2, 0x4b, 0xae, 0xfe, 0x2c, 0x30, 0xfb, 0x01,
	0x56, 0xfb, 0x0e, 0xb3, 0xbf, 0xd7, 0xd8, 0x77, 0x98, 0x0f, 0x63, 0xb7, 0xed, 0x44, 0xb3, 0x19,
	0x60, 0xd5, 0x6c, 0x3a, 0x54, 0xb3, 0xb9, 0xd7, 0x68, 0x36, 0x1d, 0xee, 0x60, 0xd2, 0x04, 0xdb,
	0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xb7, 0xa7, 0xed, 0x1c, 0x0f, 0x00, 0x00,
}
