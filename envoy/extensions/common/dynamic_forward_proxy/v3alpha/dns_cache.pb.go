// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/common/dynamic_forward_proxy/v3alpha/dns_cache.proto

package envoy_extensions_common_dynamic_forward_proxy_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/service/cluster/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type DnsCacheConfig struct {
	Name                 string                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DnsLookupFamily      v3alpha.Cluster_DnsLookupFamily `protobuf:"varint,2,opt,name=dns_lookup_family,json=dnsLookupFamily,proto3,enum=envoy.service.cluster.v3alpha.Cluster_DnsLookupFamily" json:"dns_lookup_family,omitempty"`
	DnsRefreshRate       *duration.Duration              `protobuf:"bytes,3,opt,name=dns_refresh_rate,json=dnsRefreshRate,proto3" json:"dns_refresh_rate,omitempty"`
	HostTtl              *duration.Duration              `protobuf:"bytes,4,opt,name=host_ttl,json=hostTtl,proto3" json:"host_ttl,omitempty"`
	MaxHosts             *wrappers.UInt32Value           `protobuf:"bytes,5,opt,name=max_hosts,json=maxHosts,proto3" json:"max_hosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *DnsCacheConfig) Reset()         { *m = DnsCacheConfig{} }
func (m *DnsCacheConfig) String() string { return proto.CompactTextString(m) }
func (*DnsCacheConfig) ProtoMessage()    {}
func (*DnsCacheConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd4250abe0db5c5, []int{0}
}

func (m *DnsCacheConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsCacheConfig.Unmarshal(m, b)
}
func (m *DnsCacheConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsCacheConfig.Marshal(b, m, deterministic)
}
func (m *DnsCacheConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsCacheConfig.Merge(m, src)
}
func (m *DnsCacheConfig) XXX_Size() int {
	return xxx_messageInfo_DnsCacheConfig.Size(m)
}
func (m *DnsCacheConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsCacheConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DnsCacheConfig proto.InternalMessageInfo

func (m *DnsCacheConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DnsCacheConfig) GetDnsLookupFamily() v3alpha.Cluster_DnsLookupFamily {
	if m != nil {
		return m.DnsLookupFamily
	}
	return v3alpha.Cluster_AUTO
}

func (m *DnsCacheConfig) GetDnsRefreshRate() *duration.Duration {
	if m != nil {
		return m.DnsRefreshRate
	}
	return nil
}

func (m *DnsCacheConfig) GetHostTtl() *duration.Duration {
	if m != nil {
		return m.HostTtl
	}
	return nil
}

func (m *DnsCacheConfig) GetMaxHosts() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxHosts
	}
	return nil
}

func init() {
	proto.RegisterType((*DnsCacheConfig)(nil), "envoy.extensions.common.dynamic_forward_proxy.v3alpha.DnsCacheConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/common/dynamic_forward_proxy/v3alpha/dns_cache.proto", fileDescriptor_3cd4250abe0db5c5)
}

var fileDescriptor_3cd4250abe0db5c5 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x6e, 0xd4, 0x30,
	0x14, 0x87, 0x9b, 0x74, 0x4a, 0xa7, 0x46, 0x0c, 0x25, 0x1b, 0x42, 0x41, 0x55, 0x60, 0xc3, 0xa8,
	0x0b, 0x5b, 0x9a, 0x11, 0x2c, 0x58, 0x20, 0x94, 0x19, 0xfe, 0x49, 0x20, 0x55, 0x11, 0x74, 0x1b,
	0xbd, 0xc6, 0xce, 0x8c, 0x85, 0x63, 0x47, 0xb6, 0x93, 0xce, 0x6c, 0x59, 0x71, 0x06, 0x8e, 0xc0,
	0x05, 0x38, 0x01, 0x97, 0xea, 0x0a, 0xc5, 0x4e, 0x40, 0x15, 0x20, 0x10, 0xbb, 0xbc, 0x3c, 0xff,
	0xbe, 0x97, 0xf7, 0xc5, 0xe8, 0x39, 0x93, 0xad, 0xda, 0x12, 0xb6, 0xb1, 0x4c, 0x1a, 0xae, 0xa4,
	0x21, 0x85, 0xaa, 0x2a, 0x25, 0x09, 0xdd, 0x4a, 0xa8, 0x78, 0x91, 0x97, 0x4a, 0x5f, 0x80, 0xa6,
	0x79, 0xad, 0xd5, 0x66, 0x4b, 0xda, 0x39, 0x88, 0x7a, 0x0d, 0x84, 0x4a, 0x93, 0x17, 0x50, 0xac,
	0x19, 0xae, 0xb5, 0xb2, 0x2a, 0x7a, 0xe4, 0x30, 0xf8, 0x27, 0x06, 0x7b, 0x0c, 0xfe, 0x2d, 0x06,
	0xf7, 0x98, 0xa3, 0x87, 0x7e, 0xba, 0x61, 0xba, 0xe5, 0x05, 0x23, 0x85, 0x68, 0x8c, 0x65, 0xfa,
	0xc7, 0x94, 0x82, 0x1a, 0xcf, 0x3f, 0x3a, 0x5e, 0x29, 0xb5, 0x12, 0x8c, 0xb8, 0xea, 0xbc, 0x29,
	0x09, 0x6d, 0x34, 0x58, 0xae, 0xe4, 0x9f, 0xfa, 0x17, 0x1a, 0xea, 0x9a, 0xe9, 0x21, 0x7f, 0xbf,
	0xa1, 0x35, 0x10, 0x90, 0x52, 0x59, 0x17, 0x33, 0xa4, 0x65, 0xba, 0xfb, 0x50, 0x2e, 0x57, 0xfd,
	0x91, 0xdb, 0x2d, 0x08, 0x4e, 0xc1, 0x32, 0x32, 0x3c, 0xf8, 0xc6, 0x83, 0xaf, 0xbb, 0x68, 0xb2,
	0x94, 0x66, 0xd1, 0xad, 0xbb, 0x50, 0xb2, 0xe4, 0xab, 0xe8, 0x2e, 0x1a, 0x49, 0xa8, 0x58, 0x1c,
	0x24, 0xc1, 0xf4, 0x20, 0xdd, 0xbf, 0x4c, 0x47, 0x3a, 0x4c, 0x82, 0xcc, 0xbd, 0x8c, 0x04, 0xba,
	0xd5, 0xe9, 0x11, 0x4a, 0x7d, 0x68, 0xea, 0xbc, 0x84, 0x8a, 0x8b, 0x6d, 0x1c, 0x26, 0xc1, 0x74,
	0x32, 0x7b, 0x8c, 0xbd, 0xa7, 0x7e, 0x61, 0xdc, 0x2f, 0x3c, 0xf8, 0xc0, 0x8b, 0xbe, 0x5e, 0x4a,
	0xf3, 0xc6, 0xc5, 0x5f, 0xb8, 0x74, 0x3a, 0xbe, 0x4c, 0xf7, 0x3e, 0x06, 0xe1, 0x61, 0x90, 0xdd,
	0xa4, 0x57, 0x5b, 0xd1, 0x5b, 0x74, 0xd8, 0x4d, 0xd3, 0xac, 0xd4, 0xcc, 0xac, 0x73, 0x0d, 0x96,
	0xc5, 0xbb, 0x49, 0x30, 0xbd, 0x3e, 0xbb, 0x83, 0xbd, 0x14, 0x3c, 0x48, 0xc1, 0xcb, 0x5e, 0x9a,
	0xe3, 0x7d, 0x09, 0xc2, 0x93, 0x9d, 0x6c, 0x42, 0xa5, 0xc9, 0x7c, 0x36, 0x03, 0xcb, 0xa2, 0xa7,
	0x68, 0xbc, 0x56, 0xc6, 0xe6, 0xd6, 0x8a, 0x78, 0xf4, 0xef, 0x98, 0xfd, 0x2e, 0xf4, 0xce, 0x8a,
	0x28, 0x45, 0x07, 0x15, 0x6c, 0xf2, 0xae, 0x34, 0xf1, 0x9e, 0x03, 0xdc, 0xfb, 0x05, 0xf0, 0xfe,
	0xb5, 0xb4, 0xf3, 0xd9, 0x19, 0x88, 0x86, 0x39, 0x79, 0x27, 0x61, 0xb2, 0x93, 0x8d, 0x2b, 0xd8,
	0xbc, 0xea, 0x62, 0x4f, 0x5e, 0x7e, 0xfe, 0xf6, 0xe9, 0x38, 0x45, 0xcf, 0xbc, 0xab, 0xc2, 0x29,
	0xff, 0xcb, 0x7d, 0x9a, 0x79, 0x7f, 0x57, 0x7f, 0x53, 0x7a, 0x86, 0x16, 0x5c, 0x79, 0xe5, 0xfe,
	0xdc, 0x7f, 0xdd, 0xd2, 0xf4, 0xc6, 0x80, 0x3d, 0xed, 0x16, 0x38, 0x0d, 0xce, 0xaf, 0xb9, 0x4d,
	0xe6, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xae, 0x22, 0xe3, 0xe8, 0x3d, 0x03, 0x00, 0x00,
}
