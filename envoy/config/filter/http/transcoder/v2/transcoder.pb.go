// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/transcoder/v2/transcoder.proto

package envoy_config_filter_http_transcoder_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type GrpcJsonTranscoder struct {
	// Types that are valid to be assigned to DescriptorSet:
	//	*GrpcJsonTranscoder_ProtoDescriptor
	//	*GrpcJsonTranscoder_ProtoDescriptorBin
	DescriptorSet                isGrpcJsonTranscoder_DescriptorSet `protobuf_oneof:"descriptor_set"`
	Services                     []string                           `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
	PrintOptions                 *GrpcJsonTranscoder_PrintOptions   `protobuf:"bytes,3,opt,name=print_options,json=printOptions,proto3" json:"print_options,omitempty"`
	MatchIncomingRequestRoute    bool                               `protobuf:"varint,5,opt,name=match_incoming_request_route,json=matchIncomingRequestRoute,proto3" json:"match_incoming_request_route,omitempty"`
	IgnoredQueryParameters       []string                           `protobuf:"bytes,6,rep,name=ignored_query_parameters,json=ignoredQueryParameters,proto3" json:"ignored_query_parameters,omitempty"`
	AutoMapping                  bool                               `protobuf:"varint,7,opt,name=auto_mapping,json=autoMapping,proto3" json:"auto_mapping,omitempty"`
	IgnoreUnknownQueryParameters bool                               `protobuf:"varint,8,opt,name=ignore_unknown_query_parameters,json=ignoreUnknownQueryParameters,proto3" json:"ignore_unknown_query_parameters,omitempty"`
	ConvertGrpcStatus            bool                               `protobuf:"varint,9,opt,name=convert_grpc_status,json=convertGrpcStatus,proto3" json:"convert_grpc_status,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                           `json:"-"`
	XXX_unrecognized             []byte                             `json:"-"`
	XXX_sizecache                int32                              `json:"-"`
}

func (m *GrpcJsonTranscoder) Reset()         { *m = GrpcJsonTranscoder{} }
func (m *GrpcJsonTranscoder) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder) ProtoMessage()    {}
func (*GrpcJsonTranscoder) Descriptor() ([]byte, []int) {
	return fileDescriptor_540f2f6de8b0585c, []int{0}
}

func (m *GrpcJsonTranscoder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcJsonTranscoder.Unmarshal(m, b)
}
func (m *GrpcJsonTranscoder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcJsonTranscoder.Marshal(b, m, deterministic)
}
func (m *GrpcJsonTranscoder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder.Merge(m, src)
}
func (m *GrpcJsonTranscoder) XXX_Size() int {
	return xxx_messageInfo_GrpcJsonTranscoder.Size(m)
}
func (m *GrpcJsonTranscoder) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder proto.InternalMessageInfo

type isGrpcJsonTranscoder_DescriptorSet interface {
	isGrpcJsonTranscoder_DescriptorSet()
}

type GrpcJsonTranscoder_ProtoDescriptor struct {
	ProtoDescriptor string `protobuf:"bytes,1,opt,name=proto_descriptor,json=protoDescriptor,proto3,oneof"`
}

type GrpcJsonTranscoder_ProtoDescriptorBin struct {
	ProtoDescriptorBin []byte `protobuf:"bytes,4,opt,name=proto_descriptor_bin,json=protoDescriptorBin,proto3,oneof"`
}

func (*GrpcJsonTranscoder_ProtoDescriptor) isGrpcJsonTranscoder_DescriptorSet() {}

func (*GrpcJsonTranscoder_ProtoDescriptorBin) isGrpcJsonTranscoder_DescriptorSet() {}

func (m *GrpcJsonTranscoder) GetDescriptorSet() isGrpcJsonTranscoder_DescriptorSet {
	if m != nil {
		return m.DescriptorSet
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetProtoDescriptor() string {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptor); ok {
		return x.ProtoDescriptor
	}
	return ""
}

func (m *GrpcJsonTranscoder) GetProtoDescriptorBin() []byte {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptorBin); ok {
		return x.ProtoDescriptorBin
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetPrintOptions() *GrpcJsonTranscoder_PrintOptions {
	if m != nil {
		return m.PrintOptions
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetMatchIncomingRequestRoute() bool {
	if m != nil {
		return m.MatchIncomingRequestRoute
	}
	return false
}

func (m *GrpcJsonTranscoder) GetIgnoredQueryParameters() []string {
	if m != nil {
		return m.IgnoredQueryParameters
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetAutoMapping() bool {
	if m != nil {
		return m.AutoMapping
	}
	return false
}

func (m *GrpcJsonTranscoder) GetIgnoreUnknownQueryParameters() bool {
	if m != nil {
		return m.IgnoreUnknownQueryParameters
	}
	return false
}

func (m *GrpcJsonTranscoder) GetConvertGrpcStatus() bool {
	if m != nil {
		return m.ConvertGrpcStatus
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GrpcJsonTranscoder) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GrpcJsonTranscoder_ProtoDescriptor)(nil),
		(*GrpcJsonTranscoder_ProtoDescriptorBin)(nil),
	}
}

type GrpcJsonTranscoder_PrintOptions struct {
	AddWhitespace              bool     `protobuf:"varint,1,opt,name=add_whitespace,json=addWhitespace,proto3" json:"add_whitespace,omitempty"`
	AlwaysPrintPrimitiveFields bool     `protobuf:"varint,2,opt,name=always_print_primitive_fields,json=alwaysPrintPrimitiveFields,proto3" json:"always_print_primitive_fields,omitempty"`
	AlwaysPrintEnumsAsInts     bool     `protobuf:"varint,3,opt,name=always_print_enums_as_ints,json=alwaysPrintEnumsAsInts,proto3" json:"always_print_enums_as_ints,omitempty"`
	PreserveProtoFieldNames    bool     `protobuf:"varint,4,opt,name=preserve_proto_field_names,json=preserveProtoFieldNames,proto3" json:"preserve_proto_field_names,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *GrpcJsonTranscoder_PrintOptions) Reset()         { *m = GrpcJsonTranscoder_PrintOptions{} }
func (m *GrpcJsonTranscoder_PrintOptions) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder_PrintOptions) ProtoMessage()    {}
func (*GrpcJsonTranscoder_PrintOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_540f2f6de8b0585c, []int{0, 0}
}

func (m *GrpcJsonTranscoder_PrintOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Unmarshal(m, b)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Marshal(b, m, deterministic)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Merge(m, src)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Size() int {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Size(m)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder_PrintOptions proto.InternalMessageInfo

func (m *GrpcJsonTranscoder_PrintOptions) GetAddWhitespace() bool {
	if m != nil {
		return m.AddWhitespace
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintPrimitiveFields() bool {
	if m != nil {
		return m.AlwaysPrintPrimitiveFields
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintEnumsAsInts() bool {
	if m != nil {
		return m.AlwaysPrintEnumsAsInts
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetPreserveProtoFieldNames() bool {
	if m != nil {
		return m.PreserveProtoFieldNames
	}
	return false
}

func init() {
	proto.RegisterType((*GrpcJsonTranscoder)(nil), "envoy.config.filter.http.transcoder.v2.GrpcJsonTranscoder")
	proto.RegisterType((*GrpcJsonTranscoder_PrintOptions)(nil), "envoy.config.filter.http.transcoder.v2.GrpcJsonTranscoder.PrintOptions")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/transcoder/v2/transcoder.proto", fileDescriptor_540f2f6de8b0585c)
}

var fileDescriptor_540f2f6de8b0585c = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x51, 0x4f, 0x14, 0x3b,
	0x14, 0xc7, 0x19, 0xb8, 0x70, 0x87, 0xb2, 0xc0, 0xbd, 0xbd, 0x57, 0x18, 0x37, 0xa8, 0xab, 0x51,
	0xb3, 0x89, 0xc9, 0x4c, 0xb2, 0x98, 0x68, 0xd0, 0xc4, 0x30, 0x11, 0x01, 0x13, 0x75, 0x1d, 0x35,
	0x3e, 0x36, 0x65, 0xe6, 0xb0, 0x5b, 0xdd, 0x69, 0x4b, 0xdb, 0x19, 0xd8, 0xcf, 0xe0, 0x83, 0x89,
	0x4f, 0x7e, 0x05, 0x3f, 0x9e, 0x3e, 0x19, 0x1f, 0x8c, 0x69, 0xbb, 0xac, 0x23, 0xbc, 0xf0, 0x36,
	0xed, 0xff, 0xff, 0x3f, 0x67, 0xce, 0xe9, 0x0f, 0xdd, 0x03, 0x5e, 0x8b, 0x71, 0x92, 0x0b, 0x7e,
	0xc8, 0x06, 0xc9, 0x21, 0x1b, 0x19, 0x50, 0xc9, 0xd0, 0x18, 0x99, 0x18, 0x45, 0xb9, 0xce, 0x45,
	0x01, 0x2a, 0xa9, 0x7b, 0x8d, 0x53, 0x2c, 0x95, 0x30, 0x02, 0xdf, 0x76, 0xc1, 0xd8, 0x07, 0x63,
	0x1f, 0x8c, 0x6d, 0x30, 0x6e, 0x58, 0xeb, 0x5e, 0xfb, 0x6a, 0x55, 0x48, 0x9a, 0x50, 0xce, 0x85,
	0xa1, 0x86, 0x09, 0xae, 0x93, 0x92, 0x0d, 0x14, 0x35, 0xe0, 0xeb, 0xb4, 0xd7, 0x6b, 0x3a, 0x62,
	0x05, 0x35, 0x90, 0x9c, 0x7e, 0x78, 0xe1, 0xc6, 0x97, 0x05, 0x84, 0x77, 0x95, 0xcc, 0x9f, 0x6a,
	0xc1, 0x5f, 0x4f, 0x4b, 0xe2, 0x3b, 0xe8, 0x1f, 0xa7, 0x93, 0x02, 0x74, 0xae, 0x98, 0x34, 0x42,
	0x45, 0x41, 0x27, 0xe8, 0x2e, 0xee, 0xcd, 0x64, 0xab, 0x4e, 0x79, 0x3c, 0x15, 0x70, 0x0f, 0xfd,
	0x7f, 0xd6, 0x4c, 0x0e, 0x18, 0x8f, 0xfe, 0xea, 0x04, 0xdd, 0xd6, 0xde, 0x4c, 0x86, 0xcf, 0x04,
	0x52, 0xc6, 0xf1, 0x4d, 0x14, 0x6a, 0x50, 0x35, 0xcb, 0x41, 0x47, 0xb3, 0x9d, 0xb9, 0xee, 0x62,
	0x1a, 0xfe, 0x48, 0xe7, 0x3f, 0x05, 0xb3, 0x61, 0x90, 0x4d, 0x15, 0x3c, 0x42, 0xcb, 0x52, 0x31,
	0x6e, 0x88, 0x90, 0x6e, 0xaa, 0x68, 0xae, 0x13, 0x74, 0x97, 0x7a, 0xbb, 0xf1, 0xc5, 0xd6, 0x12,
	0x9f, 0x9f, 0x2c, 0xee, 0xdb, 0x7a, 0x2f, 0x7c, 0xb9, 0xac, 0x25, 0x1b, 0x27, 0xfc, 0x08, 0x6d,
	0x94, 0xd4, 0xe4, 0x43, 0xc2, 0x78, 0x2e, 0x4a, 0xc6, 0x07, 0x44, 0xc1, 0x51, 0x05, 0xda, 0x10,
	0x25, 0x2a, 0x03, 0xd1, 0x7c, 0x27, 0xe8, 0x86, 0xd9, 0x65, 0xe7, 0xd9, 0x9f, 0x58, 0x32, 0xef,
	0xc8, 0xac, 0x01, 0xdf, 0x47, 0x11, 0x1b, 0x70, 0xa1, 0xa0, 0x20, 0x47, 0x15, 0xa8, 0x31, 0x91,
	0x54, 0xd1, 0x12, 0x0c, 0x28, 0x1d, 0x2d, 0xd8, 0x21, 0xb3, 0xb5, 0x89, 0xfe, 0xd2, 0xca, 0xfd,
	0xa9, 0x8a, 0xaf, 0xa3, 0x16, 0xad, 0x8c, 0x20, 0x25, 0x95, 0x92, 0xf1, 0x41, 0xf4, 0xb7, 0x6b,
	0xb5, 0x64, 0xef, 0x9e, 0xf9, 0x2b, 0xbc, 0x83, 0xae, 0xf9, 0x30, 0xa9, 0xf8, 0x7b, 0x2e, 0x8e,
	0xf9, 0xf9, 0x1e, 0xa1, 0x4b, 0x6d, 0x78, 0xdb, 0x1b, 0xef, 0x3a, 0xdb, 0x29, 0x46, 0xff, 0xe5,
	0x82, 0xd7, 0xa0, 0x0c, 0x19, 0x28, 0x99, 0x13, 0x6d, 0xa8, 0xa9, 0x74, 0xb4, 0xe8, 0xa2, 0xff,
	0x4e, 0x24, 0xbb, 0xb7, 0x57, 0x4e, 0x68, 0x7f, 0x0d, 0x50, 0xab, 0xb9, 0x33, 0x7c, 0x0b, 0xad,
	0xd0, 0xa2, 0x20, 0xc7, 0x43, 0x66, 0x40, 0x4b, 0x9a, 0x83, 0x03, 0x23, 0xcc, 0x96, 0x69, 0x51,
	0xbc, 0x9d, 0x5e, 0xe2, 0x6d, 0x74, 0x85, 0x8e, 0x8e, 0xe9, 0x58, 0x13, 0xff, 0x82, 0x52, 0xb1,
	0x92, 0x19, 0x56, 0x03, 0x39, 0x64, 0x30, 0x2a, 0xec, 0xab, 0xdb, 0x54, 0xdb, 0x9b, 0x5c, 0x87,
	0xfe, 0xa9, 0xe5, 0x89, 0x73, 0xe0, 0x2d, 0xd4, 0xfe, 0xa3, 0x04, 0xf0, 0xaa, 0xd4, 0x84, 0x6a,
	0xc2, 0xb8, 0xf1, 0x28, 0x84, 0xd9, 0x5a, 0x23, 0xbf, 0x63, 0xf5, 0x6d, 0xbd, 0xcf, 0x8d, 0xc6,
	0x0f, 0x50, 0x5b, 0x2a, 0xb0, 0x20, 0x01, 0xf1, 0x70, 0xba, 0xb6, 0x84, 0xd3, 0x12, 0xb4, 0x23,
	0x33, 0xcc, 0xd6, 0x4f, 0x1d, 0x7d, 0x6b, 0x70, 0x4d, 0x9f, 0x5b, 0x39, 0xbd, 0x84, 0x56, 0x1a,
	0x28, 0x6b, 0x30, 0x78, 0xee, 0x7b, 0x1a, 0xa4, 0x1f, 0x82, 0x6f, 0x9f, 0x7f, 0x7e, 0x9c, 0x7f,
	0x88, 0xb7, 0x3c, 0x7e, 0x70, 0x62, 0x80, 0x6b, 0xbb, 0x94, 0x09, 0x82, 0xda, 0x33, 0xe8, 0xf6,
	0xfa, 0x4e, 0x0b, 0x4e, 0x9a, 0x34, 0x6e, 0xd2, 0x91, 0x1c, 0x52, 0x74, 0x97, 0x09, 0x4f, 0xaf,
	0x54, 0xe2, 0x64, 0x7c, 0x41, 0x90, 0xd3, 0xd5, 0xdf, 0x04, 0xbb, 0x9f, 0xed, 0x07, 0x07, 0x0b,
	0x6e, 0xac, 0xcd, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x95, 0xdb, 0x3c, 0x5c, 0x04, 0x00,
	0x00,
}
