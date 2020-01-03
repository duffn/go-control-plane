// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/ratelimit/v3alpha/rate_limit.proto

package envoy_extensions_filters_network_ratelimit_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha1 "github.com/envoyproxy/go-control-plane/envoy/config/ratelimit/v3alpha"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/extensions/common/ratelimit/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type RateLimit struct {
	StatPrefix           string                           `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	Domain               string                           `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Descriptors          []*v3alpha.RateLimitDescriptor   `protobuf:"bytes,3,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	Timeout              *duration.Duration               `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	FailureModeDeny      bool                             `protobuf:"varint,5,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	RateLimitService     *v3alpha1.RateLimitServiceConfig `protobuf:"bytes,6,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6135d4202787954, []int{0}
}

func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimit.Unmarshal(m, b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(m, src)
}
func (m *RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimit.Size(m)
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetDescriptors() []*v3alpha.RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimit) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func (m *RateLimit) GetRateLimitService() *v3alpha1.RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.extensions.filters.network.ratelimit.v3alpha.RateLimit")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/ratelimit/v3alpha/rate_limit.proto", fileDescriptor_d6135d4202787954)
}

var fileDescriptor_d6135d4202787954 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0xe5, 0x76, 0xb7, 0xdb, 0x75, 0x25, 0x58, 0x72, 0x21, 0xf4, 0xb0, 0x04, 0x4e, 0x11,
	0x07, 0x5b, 0x6a, 0x05, 0x12, 0x7b, 0x40, 0x28, 0xdb, 0x23, 0x48, 0x55, 0xe0, 0x1e, 0x79, 0x9b,
	0x49, 0x31, 0x24, 0x9e, 0xc8, 0x71, 0x42, 0xfb, 0x06, 0x88, 0x23, 0x47, 0xde, 0x87, 0x33, 0xef,
	0xb3, 0x27, 0x54, 0xdb, 0x6d, 0xb7, 0x14, 0x90, 0xb8, 0x25, 0x33, 0xe3, 0xef, 0x9f, 0xf9, 0xf5,
	0xd3, 0x6b, 0x50, 0x1d, 0xae, 0x39, 0xac, 0x0c, 0xa8, 0x46, 0xa2, 0x6a, 0x78, 0x21, 0x4b, 0x03,
	0xba, 0xe1, 0x0a, 0xcc, 0x67, 0xd4, 0x9f, 0xb8, 0x16, 0x06, 0x4a, 0x59, 0x49, 0xc3, 0xbb, 0xa9,
	0x28, 0xeb, 0x0f, 0xc2, 0x56, 0x32, 0x5b, 0x62, 0xb5, 0x46, 0x83, 0xc1, 0xc4, 0x42, 0xd8, 0x1e,
	0xc2, 0x3c, 0x84, 0x79, 0x08, 0xdb, 0x41, 0x98, 0x87, 0x8c, 0x63, 0x27, 0xbc, 0x40, 0x55, 0xc8,
	0xe5, 0x9f, 0x44, 0xca, 0xc6, 0xd1, 0xc7, 0x2f, 0x8f, 0x56, 0x5c, 0x60, 0x55, 0xa1, 0xfa, 0xcb,
	0x66, 0x77, 0x16, 0x1b, 0x5f, 0x2e, 0x11, 0x97, 0x25, 0x70, 0xfb, 0x77, 0xd3, 0x16, 0x3c, 0x6f,
	0xb5, 0x30, 0x12, 0x95, 0xef, 0x3f, 0x69, 0xf3, 0x5a, 0x70, 0xa1, 0x14, 0x1a, 0x5b, 0x6e, 0x78,
	0x07, 0x7a, 0xa3, 0x21, 0xd5, 0xd2, 0x8f, 0x3c, 0xec, 0x44, 0x29, 0x73, 0x61, 0x80, 0x6f, 0x3f,
	0x5c, 0xe3, 0xe9, 0xcf, 0x3e, 0x3d, 0x4f, 0x85, 0x81, 0x37, 0x1b, 0xbd, 0x20, 0xa6, 0xa3, 0xc6,
	0x08, 0x93, 0xd5, 0x1a, 0x0a, 0xb9, 0x0a, 0x49, 0x44, 0xe2, 0xf3, 0xe4, 0xec, 0x36, 0x39, 0xd1,
	0xbd, 0x88, 0xa4, 0x74, 0xd3, 0x9b, 0xdb, 0x56, 0xf0, 0x98, 0x0e, 0x72, 0xac, 0x84, 0x54, 0x61,
	0xef, 0x70, 0xc8, 0x97, 0x83, 0x8f, 0x74, 0x94, 0x43, 0xb3, 0xd0, 0xb2, 0x36, 0xa8, 0x9b, 0xb0,
	0x1f, 0xf5, 0xe3, 0xd1, 0xe4, 0x15, 0x3b, 0xf2, 0xd8, 0xb9, 0x70, 0x6c, 0x2d, 0xdb, 0x6d, 0x35,
	0xdb, 0x61, 0x92, 0xe1, 0x6d, 0x72, 0xfa, 0x8d, 0xf4, 0x86, 0x24, 0xbd, 0x0b, 0x0f, 0xa6, 0xf4,
	0xcc, 0xc8, 0x0a, 0xb0, 0x35, 0xe1, 0x49, 0x44, 0xe2, 0xd1, 0xe4, 0x11, 0x73, 0x96, 0xb1, 0xad,
	0x65, 0x6c, 0xe6, 0x2d, 0x4b, 0xb7, 0x93, 0xc1, 0x33, 0xfa, 0xa0, 0x10, 0xb2, 0x6c, 0x35, 0x64,
	0x15, 0xe6, 0x90, 0xe5, 0xa0, 0xd6, 0xe1, 0x69, 0x44, 0xe2, 0x61, 0x7a, 0xdf, 0x37, 0xde, 0x62,
	0x0e, 0x33, 0x50, 0xeb, 0x40, 0xd1, 0x60, 0x1f, 0x97, 0xac, 0x01, 0xdd, 0xc9, 0x05, 0x84, 0x03,
	0xab, 0xf5, 0xc2, 0xdf, 0xe4, 0x32, 0xf0, 0xaf, 0x43, 0xde, 0xb9, 0x77, 0xd7, 0x76, 0xd0, 0xde,
	0xf2, 0x95, 0xf4, 0x2e, 0x48, 0x7a, 0xa1, 0x7f, 0x9b, 0xb8, 0xba, 0xfa, 0xfe, 0xe3, 0xcb, 0xe5,
	0x73, 0x3a, 0x3d, 0x20, 0xbb, 0x34, 0x1e, 0x84, 0xd1, 0xe7, 0xb7, 0x9b, 0xec, 0x45, 0x92, 0xf7,
	0xf4, 0xb5, 0x44, 0xb7, 0x53, 0xad, 0x71, 0xb5, 0x66, 0xff, 0x1f, 0xeb, 0xe4, 0xde, 0x0e, 0x37,
	0xdf, 0x18, 0x38, 0x27, 0x37, 0x03, 0xeb, 0xe4, 0xf4, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87,
	0xbf, 0xcf, 0xbe, 0x6a, 0x03, 0x00, 0x00,
}
