// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/talent/v4beta1/company.proto

package talent // import "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A Company resource represents a company in the service. A company is the
// entity that owns job postings, that is, the hiring entity responsible for
// employing applicants for the job position.
type Company struct {
	// Required during company update.
	//
	// The resource name for a company. This is generated by the service when a
	// company is created.
	//
	// The format is "projects/{project_id}/companies/{company_id}", for example,
	// "projects/api-test-project/companies/foo".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required.
	//
	// The display name of the company, for example, "Google, LLC".
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required.
	//
	// Client side company identifier, used to uniquely identify the
	// company.
	//
	// The maximum number of allowed characters is 255.
	ExternalId string `protobuf:"bytes,3,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// Optional.
	//
	// The employer's company size.
	Size CompanySize `protobuf:"varint,4,opt,name=size,proto3,enum=google.cloud.talent.v4beta1.CompanySize" json:"size,omitempty"`
	// Optional.
	//
	// The street address of the company's main headquarters, which may be
	// different from the job location. The service attempts
	// to geolocate the provided address, and populates a more specific
	// location wherever possible in
	// [DerivedInfo.headquarters_location][google.cloud.talent.v4beta1.Company.DerivedInfo.headquarters_location].
	HeadquartersAddress string `protobuf:"bytes,5,opt,name=headquarters_address,json=headquartersAddress,proto3" json:"headquarters_address,omitempty"`
	// Optional.
	//
	// Set to true if it is the hiring agency that post jobs for other
	// employers.
	//
	// Defaults to false if not provided.
	HiringAgency bool `protobuf:"varint,6,opt,name=hiring_agency,json=hiringAgency,proto3" json:"hiring_agency,omitempty"`
	// Optional.
	//
	// Equal Employment Opportunity legal disclaimer text to be
	// associated with all jobs, and typically to be displayed in all
	// roles.
	//
	// The maximum number of allowed characters is 500.
	EeoText string `protobuf:"bytes,7,opt,name=eeo_text,json=eeoText,proto3" json:"eeo_text,omitempty"`
	// Optional.
	//
	// The URI representing the company's primary web site or home page,
	// for example, "https://www.google.com".
	//
	// The maximum number of allowed characters is 255.
	WebsiteUri string `protobuf:"bytes,8,opt,name=website_uri,json=websiteUri,proto3" json:"website_uri,omitempty"`
	// Optional.
	//
	// The URI to employer's career site or careers page on the employer's web
	// site, for example, "https://careers.google.com".
	CareerSiteUri string `protobuf:"bytes,9,opt,name=career_site_uri,json=careerSiteUri,proto3" json:"career_site_uri,omitempty"`
	// Optional.
	//
	// A URI that hosts the employer's company logo.
	ImageUri string `protobuf:"bytes,10,opt,name=image_uri,json=imageUri,proto3" json:"image_uri,omitempty"`
	// Optional.
	//
	// A list of keys of filterable
	// [Job.custom_attributes][google.cloud.talent.v4beta1.Job.custom_attributes],
	// whose corresponding `string_values` are used in keyword search. Jobs with
	// `string_values` under these specified field keys are returned if any
	// of the values matches the search keyword. Custom field values with
	// parenthesis, brackets and special symbols won't be properly searchable,
	// and those keyword queries need to be surrounded by quotes.
	KeywordSearchableJobCustomAttributes []string `protobuf:"bytes,11,rep,name=keyword_searchable_job_custom_attributes,json=keywordSearchableJobCustomAttributes,proto3" json:"keyword_searchable_job_custom_attributes,omitempty"`
	// Output only. Derived details about the company.
	DerivedInfo *Company_DerivedInfo `protobuf:"bytes,12,opt,name=derived_info,json=derivedInfo,proto3" json:"derived_info,omitempty"`
	// Output only. Indicates whether a company is flagged to be suspended from
	// public availability by the service when job content appears suspicious,
	// abusive, or spammy.
	Suspended            bool     `protobuf:"varint,13,opt,name=suspended,proto3" json:"suspended,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Company) Reset()         { *m = Company{} }
func (m *Company) String() string { return proto.CompactTextString(m) }
func (*Company) ProtoMessage()    {}
func (*Company) Descriptor() ([]byte, []int) {
	return fileDescriptor_company_fc503f1af4088a2a, []int{0}
}
func (m *Company) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Company.Unmarshal(m, b)
}
func (m *Company) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Company.Marshal(b, m, deterministic)
}
func (dst *Company) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Company.Merge(dst, src)
}
func (m *Company) XXX_Size() int {
	return xxx_messageInfo_Company.Size(m)
}
func (m *Company) XXX_DiscardUnknown() {
	xxx_messageInfo_Company.DiscardUnknown(m)
}

var xxx_messageInfo_Company proto.InternalMessageInfo

func (m *Company) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Company) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Company) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

func (m *Company) GetSize() CompanySize {
	if m != nil {
		return m.Size
	}
	return CompanySize_COMPANY_SIZE_UNSPECIFIED
}

func (m *Company) GetHeadquartersAddress() string {
	if m != nil {
		return m.HeadquartersAddress
	}
	return ""
}

func (m *Company) GetHiringAgency() bool {
	if m != nil {
		return m.HiringAgency
	}
	return false
}

func (m *Company) GetEeoText() string {
	if m != nil {
		return m.EeoText
	}
	return ""
}

func (m *Company) GetWebsiteUri() string {
	if m != nil {
		return m.WebsiteUri
	}
	return ""
}

func (m *Company) GetCareerSiteUri() string {
	if m != nil {
		return m.CareerSiteUri
	}
	return ""
}

func (m *Company) GetImageUri() string {
	if m != nil {
		return m.ImageUri
	}
	return ""
}

func (m *Company) GetKeywordSearchableJobCustomAttributes() []string {
	if m != nil {
		return m.KeywordSearchableJobCustomAttributes
	}
	return nil
}

func (m *Company) GetDerivedInfo() *Company_DerivedInfo {
	if m != nil {
		return m.DerivedInfo
	}
	return nil
}

func (m *Company) GetSuspended() bool {
	if m != nil {
		return m.Suspended
	}
	return false
}

// Derived details about the company.
type Company_DerivedInfo struct {
	// A structured headquarters location of the company, resolved from
	// [Company.hq_location][] if provided.
	HeadquartersLocation *Location `protobuf:"bytes,1,opt,name=headquarters_location,json=headquartersLocation,proto3" json:"headquarters_location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Company_DerivedInfo) Reset()         { *m = Company_DerivedInfo{} }
func (m *Company_DerivedInfo) String() string { return proto.CompactTextString(m) }
func (*Company_DerivedInfo) ProtoMessage()    {}
func (*Company_DerivedInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_company_fc503f1af4088a2a, []int{0, 0}
}
func (m *Company_DerivedInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Company_DerivedInfo.Unmarshal(m, b)
}
func (m *Company_DerivedInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Company_DerivedInfo.Marshal(b, m, deterministic)
}
func (dst *Company_DerivedInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Company_DerivedInfo.Merge(dst, src)
}
func (m *Company_DerivedInfo) XXX_Size() int {
	return xxx_messageInfo_Company_DerivedInfo.Size(m)
}
func (m *Company_DerivedInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Company_DerivedInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Company_DerivedInfo proto.InternalMessageInfo

func (m *Company_DerivedInfo) GetHeadquartersLocation() *Location {
	if m != nil {
		return m.HeadquartersLocation
	}
	return nil
}

func init() {
	proto.RegisterType((*Company)(nil), "google.cloud.talent.v4beta1.Company")
	proto.RegisterType((*Company_DerivedInfo)(nil), "google.cloud.talent.v4beta1.Company.DerivedInfo")
}

func init() {
	proto.RegisterFile("google/cloud/talent/v4beta1/company.proto", fileDescriptor_company_fc503f1af4088a2a)
}

var fileDescriptor_company_fc503f1af4088a2a = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x95, 0xb5, 0x5b, 0x5b, 0xa7, 0x05, 0xc9, 0x0c, 0x29, 0x74, 0x93, 0x56, 0xfe, 0x2a,
	0x5c, 0x12, 0x56, 0xb8, 0xc1, 0xa5, 0x2b, 0x97, 0x21, 0x84, 0xa6, 0x74, 0x70, 0xd8, 0xc5, 0x72,
	0xe2, 0x77, 0xa9, 0x21, 0xb1, 0x83, 0xed, 0x6c, 0xed, 0x8e, 0x7c, 0x14, 0x3e, 0x00, 0x9f, 0x11,
	0xd5, 0xce, 0xba, 0x4d, 0x42, 0x65, 0xb7, 0xe4, 0x79, 0x7e, 0x8f, 0x5f, 0xbf, 0xf6, 0x6b, 0xf4,
	0x3a, 0x97, 0x32, 0x2f, 0x20, 0xce, 0x0a, 0x59, 0xb3, 0xd8, 0xd0, 0x02, 0x84, 0x89, 0x2f, 0xde,
	0xa5, 0x60, 0xe8, 0x61, 0x9c, 0xc9, 0xb2, 0xa2, 0x62, 0x19, 0x55, 0x4a, 0x1a, 0x89, 0xf7, 0x1c,
	0x1a, 0x59, 0x34, 0x72, 0x68, 0xd4, 0xa0, 0xc3, 0xfd, 0x66, 0x1d, 0x5a, 0xf1, 0x98, 0x0a, 0x21,
	0x0d, 0x35, 0x5c, 0x0a, 0xed, 0xa2, 0xc3, 0xf0, 0x3f, 0x55, 0x4a, 0x29, 0x1c, 0xf9, 0xec, 0xcf,
	0x36, 0xea, 0x4c, 0x5d, 0x59, 0x8c, 0x51, 0x5b, 0xd0, 0x12, 0x02, 0x6f, 0xe4, 0x85, 0xbd, 0xc4,
	0x7e, 0xe3, 0xa7, 0xa8, 0xcf, 0xb8, 0xae, 0x0a, 0xba, 0x24, 0xd6, 0xdb, 0xb2, 0x9e, 0xdf, 0x68,
	0x5f, 0x56, 0xc8, 0x01, 0xf2, 0x61, 0x61, 0x40, 0x09, 0x5a, 0x10, 0xce, 0x82, 0x96, 0x25, 0xd0,
	0xb5, 0x74, 0xcc, 0xf0, 0x07, 0xd4, 0xd6, 0xfc, 0x0a, 0x82, 0xf6, 0xc8, 0x0b, 0x1f, 0x8c, 0xc3,
	0x68, 0x43, 0x5f, 0x51, 0xb3, 0x97, 0x19, 0xbf, 0x82, 0xc4, 0xa6, 0xf0, 0x21, 0xda, 0x9d, 0x03,
	0x65, 0x3f, 0x6b, 0xaa, 0x0c, 0x28, 0x4d, 0x28, 0x63, 0x0a, 0xb4, 0x0e, 0xb6, 0x6d, 0x9d, 0x47,
	0xb7, 0xbd, 0x89, 0xb3, 0xf0, 0x73, 0x34, 0x98, 0x73, 0xc5, 0x45, 0x4e, 0x68, 0x0e, 0x22, 0x5b,
	0x06, 0x3b, 0x23, 0x2f, 0xec, 0x26, 0x7d, 0x27, 0x4e, 0xac, 0x86, 0x9f, 0xa0, 0x2e, 0x80, 0x24,
	0x06, 0x16, 0x26, 0xe8, 0xd8, 0xb5, 0x3a, 0x00, 0xf2, 0x14, 0x16, 0x66, 0xd5, 0xd1, 0x25, 0xa4,
	0x9a, 0x1b, 0x20, 0xb5, 0xe2, 0x41, 0xd7, 0x75, 0xd4, 0x48, 0x5f, 0x15, 0xc7, 0xaf, 0xd0, 0xc3,
	0x8c, 0x2a, 0x00, 0x45, 0xd6, 0x50, 0xcf, 0x42, 0x03, 0x27, 0xcf, 0x1a, 0x6e, 0x0f, 0xf5, 0x78,
	0x49, 0x73, 0x47, 0x20, 0x4b, 0x74, 0xad, 0xb0, 0x32, 0xbf, 0xa1, 0xf0, 0x07, 0x2c, 0x2f, 0xa5,
	0x62, 0x44, 0x03, 0x55, 0xd9, 0x9c, 0xa6, 0x05, 0x90, 0xef, 0x32, 0x25, 0x59, 0xad, 0x8d, 0x2c,
	0x09, 0x35, 0x46, 0xf1, 0xb4, 0x36, 0xa0, 0x03, 0x7f, 0xd4, 0x0a, 0x7b, 0xc9, 0x8b, 0x86, 0x9f,
	0xad, 0xf1, 0x4f, 0x32, 0x9d, 0x5a, 0x78, 0xb2, 0x66, 0xf1, 0x0c, 0xf5, 0x19, 0x28, 0x7e, 0x01,
	0x8c, 0x70, 0x71, 0x2e, 0x83, 0xfe, 0xc8, 0x0b, 0xfd, 0xf1, 0x9b, 0xfb, 0x1c, 0x7b, 0xf4, 0xd1,
	0x05, 0x8f, 0xc5, 0xb9, 0x4c, 0x7c, 0x76, 0xf3, 0x83, 0xf7, 0x51, 0x4f, 0xd7, 0xba, 0x02, 0xc1,
	0x80, 0x05, 0x03, 0x7b, 0x9c, 0x37, 0xc2, 0x90, 0x23, 0xff, 0x56, 0x12, 0x9f, 0xa1, 0xc7, 0x77,
	0xae, 0xac, 0x90, 0x99, 0x1d, 0x4f, 0x3b, 0x59, 0xfe, 0xf8, 0xe5, 0xc6, 0xad, 0x7c, 0x6e, 0xe0,
	0xe4, 0xce, 0xb5, 0x5f, 0xab, 0x47, 0xbf, 0x3c, 0x74, 0x90, 0xc9, 0x72, 0xd3, 0x12, 0x47, 0xbb,
	0x4d, 0x3b, 0x09, 0x68, 0x59, 0xab, 0x0c, 0x4e, 0x56, 0xa3, 0x7e, 0xe2, 0x9d, 0x4d, 0x9a, 0x50,
	0x2e, 0x0b, 0x2a, 0xf2, 0x48, 0xaa, 0x3c, 0xce, 0x41, 0xd8, 0x87, 0x10, 0x3b, 0x8b, 0x56, 0x5c,
	0xff, 0xf3, 0xd5, 0xbc, 0x77, 0xbf, 0xbf, 0xb7, 0x5a, 0xd3, 0xd3, 0x59, 0xba, 0x63, 0x33, 0x6f,
	0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x72, 0xd8, 0xd3, 0x40, 0xce, 0x03, 0x00, 0x00,
}
