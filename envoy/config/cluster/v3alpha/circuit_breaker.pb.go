// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/cluster/v3alpha/circuit_breaker.proto

package envoy_config_cluster_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/config/core/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type CircuitBreakers struct {
	Thresholds           []*CircuitBreakers_Thresholds `protobuf:"bytes,1,rep,name=thresholds,proto3" json:"thresholds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CircuitBreakers) Reset()         { *m = CircuitBreakers{} }
func (m *CircuitBreakers) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers) ProtoMessage()    {}
func (*CircuitBreakers) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cdfac73f383a10, []int{0}
}

func (m *CircuitBreakers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakers.Unmarshal(m, b)
}
func (m *CircuitBreakers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakers.Marshal(b, m, deterministic)
}
func (m *CircuitBreakers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers.Merge(m, src)
}
func (m *CircuitBreakers) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakers.Size(m)
}
func (m *CircuitBreakers) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers proto.InternalMessageInfo

func (m *CircuitBreakers) GetThresholds() []*CircuitBreakers_Thresholds {
	if m != nil {
		return m.Thresholds
	}
	return nil
}

type CircuitBreakers_Thresholds struct {
	Priority             v3alpha.RoutingPriority `protobuf:"varint,1,opt,name=priority,proto3,enum=envoy.config.core.v3alpha.RoutingPriority" json:"priority,omitempty"`
	MaxConnections       *wrappers.UInt32Value   `protobuf:"bytes,2,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	MaxPendingRequests   *wrappers.UInt32Value   `protobuf:"bytes,3,opt,name=max_pending_requests,json=maxPendingRequests,proto3" json:"max_pending_requests,omitempty"`
	MaxRequests          *wrappers.UInt32Value   `protobuf:"bytes,4,opt,name=max_requests,json=maxRequests,proto3" json:"max_requests,omitempty"`
	MaxRetries           *wrappers.UInt32Value   `protobuf:"bytes,5,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	TrackRemaining       bool                    `protobuf:"varint,6,opt,name=track_remaining,json=trackRemaining,proto3" json:"track_remaining,omitempty"`
	MaxConnectionPools   *wrappers.UInt32Value   `protobuf:"bytes,7,opt,name=max_connection_pools,json=maxConnectionPools,proto3" json:"max_connection_pools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CircuitBreakers_Thresholds) Reset()         { *m = CircuitBreakers_Thresholds{} }
func (m *CircuitBreakers_Thresholds) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers_Thresholds) ProtoMessage()    {}
func (*CircuitBreakers_Thresholds) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cdfac73f383a10, []int{0, 0}
}

func (m *CircuitBreakers_Thresholds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Unmarshal(m, b)
}
func (m *CircuitBreakers_Thresholds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Marshal(b, m, deterministic)
}
func (m *CircuitBreakers_Thresholds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers_Thresholds.Merge(m, src)
}
func (m *CircuitBreakers_Thresholds) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Size(m)
}
func (m *CircuitBreakers_Thresholds) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers_Thresholds.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers_Thresholds proto.InternalMessageInfo

func (m *CircuitBreakers_Thresholds) GetPriority() v3alpha.RoutingPriority {
	if m != nil {
		return m.Priority
	}
	return v3alpha.RoutingPriority_DEFAULT
}

func (m *CircuitBreakers_Thresholds) GetMaxConnections() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxPendingRequests() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxPendingRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRequests() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRetries() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxRetries
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetTrackRemaining() bool {
	if m != nil {
		return m.TrackRemaining
	}
	return false
}

func (m *CircuitBreakers_Thresholds) GetMaxConnectionPools() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConnectionPools
	}
	return nil
}

func init() {
	proto.RegisterType((*CircuitBreakers)(nil), "envoy.config.cluster.v3alpha.CircuitBreakers")
	proto.RegisterType((*CircuitBreakers_Thresholds)(nil), "envoy.config.cluster.v3alpha.CircuitBreakers.Thresholds")
}

func init() {
	proto.RegisterFile("envoy/config/cluster/v3alpha/circuit_breaker.proto", fileDescriptor_c7cdfac73f383a10)
}

var fileDescriptor_c7cdfac73f383a10 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x8a, 0xd4, 0x30,
	0x18, 0xc7, 0xe9, 0xee, 0xce, 0x3a, 0x64, 0x64, 0x46, 0xaa, 0x60, 0x19, 0x96, 0x65, 0x94, 0x45,
	0x87, 0x15, 0x12, 0xec, 0x80, 0xc8, 0x82, 0x08, 0x1d, 0x3c, 0x78, 0x91, 0x52, 0x54, 0xbc, 0x95,
	0x4c, 0x27, 0xdb, 0x09, 0xdb, 0xe6, 0x8b, 0x49, 0x5a, 0x3b, 0x57, 0x4f, 0x3e, 0x83, 0x0f, 0xe0,
	0x7b, 0x78, 0xf0, 0xa5, 0x3c, 0x49, 0x9b, 0x4e, 0xc7, 0x5d, 0xd8, 0xa5, 0xb7, 0x36, 0x5f, 0x7e,
	0xbf, 0xe4, 0xfb, 0xf2, 0x47, 0x3e, 0x13, 0x25, 0x6c, 0x49, 0x02, 0xe2, 0x92, 0xa7, 0x24, 0xc9,
	0x0a, 0x6d, 0x98, 0x22, 0xe5, 0x82, 0x66, 0x72, 0x43, 0x49, 0xc2, 0x55, 0x52, 0x70, 0x13, 0xaf,
	0x14, 0xa3, 0x57, 0x4c, 0x61, 0xa9, 0xc0, 0x80, 0x7b, 0xd2, 0x30, 0xd8, 0x32, 0xb8, 0x65, 0x70,
	0xcb, 0x4c, 0xcf, 0xae, 0x1b, 0x41, 0xb1, 0x4e, 0xb7, 0xa2, 0x9a, 0x59, 0xc7, 0xf4, 0x34, 0x05,
	0x48, 0x33, 0x46, 0x9a, 0xbf, 0x55, 0x71, 0x49, 0xbe, 0x29, 0x2a, 0x25, 0x53, 0xba, 0xad, 0x3f,
	0x29, 0xd6, 0x92, 0x12, 0x2a, 0x04, 0x18, 0x6a, 0x38, 0x08, 0x4d, 0x4a, 0xa6, 0x34, 0x07, 0xc1,
	0x45, 0xda, 0x6e, 0x79, 0x5c, 0xd2, 0x8c, 0xaf, 0xa9, 0x61, 0x64, 0xf7, 0x61, 0x0b, 0x4f, 0x7f,
	0x0f, 0xd0, 0x64, 0x69, 0x6f, 0x1e, 0xd8, 0x8b, 0x6b, 0xf7, 0x0b, 0x42, 0x66, 0xa3, 0x98, 0xde,
	0x40, 0xb6, 0xd6, 0x9e, 0x33, 0x3b, 0x9c, 0x8f, 0xfc, 0xd7, 0xf8, 0xae, 0x46, 0xf0, 0x0d, 0x05,
	0xfe, 0xd8, 0xf1, 0xd1, 0x7f, 0xae, 0xe9, 0xaf, 0x23, 0x84, 0xf6, 0x25, 0x37, 0x44, 0x43, 0xa9,
	0x38, 0x28, 0x6e, 0xb6, 0x9e, 0x33, 0x73, 0xe6, 0x63, 0xff, 0xfc, 0xc6, 0x31, 0xa0, 0x58, 0x77,
	0x46, 0x04, 0x85, 0xe1, 0x22, 0x0d, 0x5b, 0x22, 0x18, 0xfe, 0x0d, 0x06, 0xdf, 0x9d, 0x83, 0x07,
	0x4e, 0xd4, 0x59, 0xdc, 0x77, 0x68, 0x92, 0xd3, 0x2a, 0x4e, 0x40, 0x08, 0x96, 0x34, 0xb3, 0xf0,
	0x0e, 0x66, 0xce, 0x7c, 0xe4, 0x9f, 0x60, 0x3b, 0x44, 0xbc, 0x1b, 0x22, 0xfe, 0xf4, 0x5e, 0x98,
	0x85, 0xff, 0x99, 0x66, 0x05, 0x8b, 0xc6, 0x39, 0xad, 0x96, 0x7b, 0xc6, 0xfd, 0x80, 0x1e, 0xd5,
	0x1a, 0xc9, 0xc4, 0x9a, 0x8b, 0x34, 0x56, 0xec, 0x6b, 0xc1, 0xb4, 0xd1, 0xde, 0x61, 0x0f, 0x97,
	0x9b, 0xd3, 0x2a, 0xb4, 0x60, 0xd4, 0x72, 0xee, 0x5b, 0x74, 0xbf, 0xf6, 0x75, 0x9e, 0xa3, 0x1e,
	0x9e, 0x51, 0x4e, 0xab, 0x4e, 0xf0, 0x06, 0x8d, 0xac, 0xc0, 0x28, 0xce, 0xb4, 0x37, 0xe8, 0xc1,
	0xa3, 0x86, 0x6f, 0xf6, 0xbb, 0xcf, 0xd1, 0xc4, 0x28, 0x9a, 0x5c, 0xc5, 0x8a, 0xe5, 0x94, 0xd7,
	0xb9, 0xf0, 0x8e, 0x67, 0xce, 0x7c, 0x18, 0x8d, 0x9b, 0xe5, 0x68, 0xb7, 0xba, 0x6b, 0x7c, 0x3f,
	0xbf, 0x58, 0x02, 0x64, 0xda, 0xbb, 0xd7, 0xb3, 0xf1, 0xfd, 0x10, 0xc3, 0x9a, 0xbb, 0x78, 0xf5,
	0xf3, 0xcf, 0x8f, 0xd3, 0x97, 0x88, 0xd8, 0x57, 0xa5, 0x92, 0xe3, 0xd2, 0xef, 0xc2, 0x73, 0x7b,
	0x68, 0x2e, 0x5e, 0xd4, 0xdc, 0x33, 0x74, 0xd6, 0x87, 0x0b, 0x96, 0xe8, 0x9c, 0x83, 0x0d, 0x8e,
	0x54, 0x50, 0x6d, 0xef, 0x8c, 0x6a, 0xf0, 0xf0, 0x3a, 0x1e, 0xd6, 0xad, 0x84, 0xce, 0xea, 0xb8,
	0xe9, 0x69, 0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x65, 0x8b, 0x8f, 0x52, 0xe5, 0x03, 0x00, 0x00,
}
