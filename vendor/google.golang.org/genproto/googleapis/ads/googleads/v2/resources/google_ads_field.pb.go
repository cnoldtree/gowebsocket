// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/google_ads_field.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// A field or resource (artifact) used by GoogleAdsService.
type GoogleAdsField struct {
	// The resource name of the artifact.
	// Artifact resource names have the form:
	//
	// `googleAdsFields/{name}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The name of the artifact.
	Name *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The category of the artifact.
	Category enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory `protobuf:"varint,3,opt,name=category,proto3,enum=google.ads.googleads.v2.enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory" json:"category,omitempty"`
	// Whether the artifact can be used in a SELECT clause in search
	// queries.
	Selectable *wrappers.BoolValue `protobuf:"bytes,4,opt,name=selectable,proto3" json:"selectable,omitempty"`
	// Whether the artifact can be used in a WHERE clause in search
	// queries.
	Filterable *wrappers.BoolValue `protobuf:"bytes,5,opt,name=filterable,proto3" json:"filterable,omitempty"`
	// Whether the artifact can be used in a ORDER BY clause in search
	// queries.
	Sortable *wrappers.BoolValue `protobuf:"bytes,6,opt,name=sortable,proto3" json:"sortable,omitempty"`
	// The names of all resources, segments, and metrics that are selectable with
	// the described artifact.
	SelectableWith []*wrappers.StringValue `protobuf:"bytes,7,rep,name=selectable_with,json=selectableWith,proto3" json:"selectable_with,omitempty"`
	// The names of all resources that are selectable with the described
	// artifact. Fields from these resources do not segment metrics when included
	// in search queries.
	//
	// This field is only set for artifacts whose category is RESOURCE.
	AttributeResources []*wrappers.StringValue `protobuf:"bytes,8,rep,name=attribute_resources,json=attributeResources,proto3" json:"attribute_resources,omitempty"`
	// At and beyond version V1 this field lists the names of all metrics that are
	// selectable with the described artifact when it is used in the FROM clause.
	// It is only set for artifacts whose category is RESOURCE.
	//
	// Before version V1 this field lists the names of all metrics that are
	// selectable with the described artifact. It is only set for artifacts whose
	// category is either RESOURCE or SEGMENT
	Metrics []*wrappers.StringValue `protobuf:"bytes,9,rep,name=metrics,proto3" json:"metrics,omitempty"`
	// At and beyond version V1 this field lists the names of all artifacts,
	// whether a segment or another resource, that segment metrics when included
	// in search queries and when the described artifact is used in the FROM
	// clause. It is only set for artifacts whose category is RESOURCE.
	//
	// Before version V1 this field lists the names of all artifacts, whether a
	// segment or another resource, that segment metrics when included in search
	// queries. It is only set for artifacts of category RESOURCE, SEGMENT or
	// METRIC.
	Segments []*wrappers.StringValue `protobuf:"bytes,10,rep,name=segments,proto3" json:"segments,omitempty"`
	// Values the artifact can assume if it is a field of type ENUM.
	//
	// This field is only set for artifacts of category SEGMENT or ATTRIBUTE.
	EnumValues []*wrappers.StringValue `protobuf:"bytes,11,rep,name=enum_values,json=enumValues,proto3" json:"enum_values,omitempty"`
	// This field determines the operators that can be used with the artifact
	// in WHERE clauses.
	DataType enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType `protobuf:"varint,12,opt,name=data_type,json=dataType,proto3,enum=google.ads.googleads.v2.enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType" json:"data_type,omitempty"`
	// The URL of proto describing the artifact's data type.
	TypeUrl *wrappers.StringValue `protobuf:"bytes,13,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// Whether the field artifact is repeated.
	IsRepeated           *wrappers.BoolValue `protobuf:"bytes,14,opt,name=is_repeated,json=isRepeated,proto3" json:"is_repeated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GoogleAdsField) Reset()         { *m = GoogleAdsField{} }
func (m *GoogleAdsField) String() string { return proto.CompactTextString(m) }
func (*GoogleAdsField) ProtoMessage()    {}
func (*GoogleAdsField) Descriptor() ([]byte, []int) {
	return fileDescriptor_e51b907878502641, []int{0}
}

func (m *GoogleAdsField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleAdsField.Unmarshal(m, b)
}
func (m *GoogleAdsField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleAdsField.Marshal(b, m, deterministic)
}
func (m *GoogleAdsField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleAdsField.Merge(m, src)
}
func (m *GoogleAdsField) XXX_Size() int {
	return xxx_messageInfo_GoogleAdsField.Size(m)
}
func (m *GoogleAdsField) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleAdsField.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleAdsField proto.InternalMessageInfo

func (m *GoogleAdsField) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *GoogleAdsField) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *GoogleAdsField) GetCategory() enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory {
	if m != nil {
		return m.Category
	}
	return enums.GoogleAdsFieldCategoryEnum_UNSPECIFIED
}

func (m *GoogleAdsField) GetSelectable() *wrappers.BoolValue {
	if m != nil {
		return m.Selectable
	}
	return nil
}

func (m *GoogleAdsField) GetFilterable() *wrappers.BoolValue {
	if m != nil {
		return m.Filterable
	}
	return nil
}

func (m *GoogleAdsField) GetSortable() *wrappers.BoolValue {
	if m != nil {
		return m.Sortable
	}
	return nil
}

func (m *GoogleAdsField) GetSelectableWith() []*wrappers.StringValue {
	if m != nil {
		return m.SelectableWith
	}
	return nil
}

func (m *GoogleAdsField) GetAttributeResources() []*wrappers.StringValue {
	if m != nil {
		return m.AttributeResources
	}
	return nil
}

func (m *GoogleAdsField) GetMetrics() []*wrappers.StringValue {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *GoogleAdsField) GetSegments() []*wrappers.StringValue {
	if m != nil {
		return m.Segments
	}
	return nil
}

func (m *GoogleAdsField) GetEnumValues() []*wrappers.StringValue {
	if m != nil {
		return m.EnumValues
	}
	return nil
}

func (m *GoogleAdsField) GetDataType() enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType {
	if m != nil {
		return m.DataType
	}
	return enums.GoogleAdsFieldDataTypeEnum_UNSPECIFIED
}

func (m *GoogleAdsField) GetTypeUrl() *wrappers.StringValue {
	if m != nil {
		return m.TypeUrl
	}
	return nil
}

func (m *GoogleAdsField) GetIsRepeated() *wrappers.BoolValue {
	if m != nil {
		return m.IsRepeated
	}
	return nil
}

func init() {
	proto.RegisterType((*GoogleAdsField)(nil), "google.ads.googleads.v2.resources.GoogleAdsField")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/google_ads_field.proto", fileDescriptor_e51b907878502641)
}

var fileDescriptor_e51b907878502641 = []byte{
	// 582 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdf, 0x6a, 0x13, 0x4f,
	0x14, 0xc7, 0x49, 0xd2, 0x5f, 0xb3, 0x99, 0xb4, 0xf9, 0xc1, 0xf4, 0x66, 0x09, 0x45, 0x52, 0xa5,
	0x90, 0xab, 0x59, 0x59, 0xa1, 0x96, 0x2d, 0x15, 0x36, 0x5a, 0x0b, 0x82, 0x12, 0x56, 0x8d, 0x20,
	0x81, 0x65, 0x92, 0x3d, 0xd9, 0x0e, 0xec, 0xee, 0x2c, 0x33, 0xb3, 0x29, 0xb9, 0xf3, 0x59, 0xbc,
	0xf4, 0xce, 0xd7, 0xf0, 0x51, 0x7c, 0x0a, 0xd9, 0x7f, 0x93, 0x96, 0x1a, 0x13, 0xbd, 0x3b, 0x39,
	0xe7, 0xfb, 0x39, 0xff, 0x72, 0x76, 0xd0, 0x79, 0xc8, 0x79, 0x18, 0x81, 0x45, 0x03, 0x69, 0x95,
	0x66, 0x6e, 0x2d, 0x6d, 0x4b, 0x80, 0xe4, 0x99, 0x98, 0x43, 0xed, 0xf6, 0x69, 0x20, 0xfd, 0x05,
	0x83, 0x28, 0x20, 0xa9, 0xe0, 0x8a, 0xe3, 0x93, 0xd2, 0x4f, 0x68, 0x20, 0x89, 0x26, 0xc9, 0xd2,
	0x26, 0x9a, 0xec, 0x5f, 0x6e, 0x4a, 0x0e, 0x49, 0x16, 0x3f, 0x4c, 0xec, 0xcf, 0xa9, 0x82, 0x90,
	0x8b, 0x55, 0x59, 0xa1, 0xff, 0xe2, 0x2f, 0xf1, 0x80, 0x2a, 0xea, 0xab, 0x55, 0x0a, 0x15, 0xff,
	0xa8, 0xe2, 0x8b, 0x5f, 0xb3, 0x6c, 0x61, 0xdd, 0x0a, 0x9a, 0xa6, 0x20, 0x64, 0x15, 0x3f, 0xae,
	0xf3, 0xa7, 0xcc, 0xa2, 0x49, 0xc2, 0x15, 0x55, 0x8c, 0x27, 0x55, 0xf4, 0xf1, 0xf7, 0x36, 0xea,
	0x5d, 0x17, 0x02, 0x37, 0x90, 0xaf, 0xf3, 0x02, 0xf8, 0x09, 0x3a, 0xac, 0x87, 0xf3, 0x13, 0x1a,
	0x83, 0xd9, 0x18, 0x34, 0x86, 0x1d, 0xef, 0xa0, 0x76, 0xbe, 0xa3, 0x31, 0xe0, 0xa7, 0x68, 0xaf,
	0x88, 0x35, 0x07, 0x8d, 0x61, 0xd7, 0x3e, 0xae, 0x76, 0x43, 0xea, 0x26, 0xc8, 0x7b, 0x25, 0x58,
	0x12, 0x4e, 0x68, 0x94, 0x81, 0x57, 0x28, 0x71, 0x84, 0x8c, 0x7a, 0x72, 0xb3, 0x35, 0x68, 0x0c,
	0x7b, 0xf6, 0x98, 0x6c, 0x5a, 0x6e, 0x31, 0x3a, 0xb9, 0xdf, 0xd7, 0xcb, 0x0a, 0xbe, 0x4a, 0xb2,
	0x78, 0x43, 0xc8, 0xd3, 0x15, 0xb0, 0x83, 0x90, 0x84, 0x08, 0xe6, 0x8a, 0xce, 0x22, 0x30, 0xf7,
	0x8a, 0x2e, 0xfb, 0x0f, 0xba, 0x1c, 0x71, 0x1e, 0x95, 0x3d, 0xde, 0x51, 0xe7, 0xec, 0x82, 0x45,
	0x0a, 0x44, 0xc1, 0xfe, 0xb7, 0x9d, 0x5d, 0xab, 0xf1, 0x19, 0x32, 0x24, 0x17, 0x65, 0xd5, 0xfd,
	0xad, 0xa4, 0xd6, 0xe2, 0x2b, 0xf4, 0xff, 0xba, 0x03, 0xff, 0x96, 0xa9, 0x1b, 0xb3, 0x3d, 0x68,
	0x6d, 0x5d, 0x6d, 0x6f, 0x0d, 0x7d, 0x62, 0xea, 0x06, 0xbf, 0x45, 0x47, 0x54, 0x29, 0xc1, 0x66,
	0x99, 0x02, 0x5f, 0x9f, 0xa8, 0x69, 0xec, 0x90, 0x0a, 0x6b, 0xd0, 0xab, 0x39, 0x7c, 0x86, 0xda,
	0x31, 0x28, 0xc1, 0xe6, 0xd2, 0xec, 0xec, 0x90, 0xa2, 0x16, 0xe3, 0x73, 0x64, 0x48, 0x08, 0x63,
	0x48, 0x94, 0x34, 0xd1, 0x0e, 0xa0, 0x56, 0xe3, 0x4b, 0xd4, 0xcd, 0xff, 0x7c, 0x7f, 0x99, 0xfb,
	0xa5, 0xd9, 0xdd, 0x01, 0x46, 0x39, 0x50, 0x98, 0x12, 0xc7, 0xa8, 0xa3, 0xbf, 0x0f, 0xf3, 0xe0,
	0x1f, 0xae, 0xec, 0x15, 0x55, 0xf4, 0xc3, 0x2a, 0x85, 0xdf, 0x5c, 0x59, 0x1d, 0xf2, 0x8c, 0xa0,
	0xb2, 0xf0, 0x73, 0x64, 0xe4, 0x95, 0xfc, 0x4c, 0x44, 0xe6, 0xe1, 0x0e, 0x5f, 0x42, 0x3b, 0x57,
	0x7f, 0x14, 0x11, 0xbe, 0x40, 0x5d, 0x26, 0x7d, 0x01, 0x29, 0x50, 0x05, 0x81, 0xd9, 0xdb, 0x7e,
	0x63, 0x4c, 0x7a, 0x95, 0x7a, 0xf4, 0xa5, 0x89, 0x4e, 0xe7, 0x3c, 0x26, 0x5b, 0x9f, 0xa6, 0xd1,
	0xd1, 0xfd, 0x09, 0xc6, 0x79, 0xde, 0x71, 0xe3, 0xf3, 0x9b, 0x8a, 0x0c, 0x79, 0x44, 0x93, 0x90,
	0x70, 0x11, 0x5a, 0x21, 0x24, 0x45, 0xd5, 0xfa, 0x09, 0x4a, 0x99, 0xfc, 0xc3, 0x6b, 0x79, 0xa1,
	0xad, 0xaf, 0xcd, 0xd6, 0xb5, 0xeb, 0x7e, 0x6b, 0x9e, 0x94, 0x95, 0x88, 0x1b, 0xdc, 0xd9, 0x28,
	0x99, 0xd8, 0x44, 0x1f, 0xd3, 0x8f, 0x5a, 0x33, 0x75, 0x03, 0x39, 0xd5, 0x9a, 0xe9, 0xc4, 0x9e,
	0x6a, 0xcd, 0xcf, 0xe6, 0x69, 0x19, 0x70, 0x1c, 0x37, 0x90, 0x8e, 0xa3, 0x55, 0x8e, 0x33, 0xb1,
	0x1d, 0x47, 0xeb, 0x66, 0xfb, 0x45, 0xb3, 0xcf, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x88,
	0xa4, 0x05, 0xd9, 0x05, 0x00, 0x00,
}
