// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.28.3
// source: envoy/extensions/filters/http/grpc_json_reverse_transcoder/v3/transcoder.proto

package grpc_json_reverse_transcoderv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// [#next-free-field: 6]
// “GrpcJsonReverseTranscoder“ is the filter configuration for the gRPC JSON
// reverse transcoder. The reverse transcoder acts as a bridge between a gRPC
// client and an HTTP/JSON server, converting the gRPC request into HTTP/JSON
// for the HTTP backend and the HTTP/JSON response back to gRPC for the gRPC
// client. This effectively reverses the behavior of the
// :ref:`grpc_json_transcoder filter <config_http_filters_grpc_json_transcoder>`,
// allowing a gRPC client to communicate with an HTTP/JSON server.
type GrpcJsonReverseTranscoder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Supplies the filename of
	// :ref:`the proto descriptor set
	// <config_grpc_json_reverse_transcoder_generate_proto_descriptor_set>` for the gRPC services.
	// If set, takes precedence over the “descriptor_binary“ field.
	DescriptorPath string `protobuf:"bytes,1,opt,name=descriptor_path,json=descriptorPath,proto3" json:"descriptor_path,omitempty"`
	// Supplies the binary content of
	// :ref:`the proto descriptor set
	// <config_grpc_json_reverse_transcoder_generate_proto_descriptor_set>` for the gRPC services.
	// If “descriptor_path“ is set, this field is not used.
	DescriptorBinary []byte `protobuf:"bytes,2,opt,name=descriptor_binary,json=descriptorBinary,proto3" json:"descriptor_binary,omitempty"`
	// The maximum size of a request body to be transcoded, in bytes. A body exceeding this size will
	// provoke a “gRPC status: ResourceExhausted“ response.
	//
	// Large values may cause envoy to use a lot of memory if there are many
	// concurrent requests.
	//
	// If unset, the current stream buffer size is used.
	MaxRequestBodySize *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=max_request_body_size,json=maxRequestBodySize,proto3" json:"max_request_body_size,omitempty"`
	// The maximum size of a response body to be transcoded, in bytes. A body exceeding this size will
	// provoke a “gRPC status: Internal“ response.
	//
	// Large values may cause envoy to use a lot of memory if there are many
	// concurrent requests.
	//
	// If unset, the current stream buffer size is used.
	MaxResponseBodySize *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=max_response_body_size,json=maxResponseBodySize,proto3" json:"max_response_body_size,omitempty"`
	// The name of the header field that has the API version of the request.
	ApiVersionHeader string `protobuf:"bytes,5,opt,name=api_version_header,json=apiVersionHeader,proto3" json:"api_version_header,omitempty"`
}

func (x *GrpcJsonReverseTranscoder) Reset() {
	*x = GrpcJsonReverseTranscoder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcJsonReverseTranscoder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcJsonReverseTranscoder) ProtoMessage() {}

func (x *GrpcJsonReverseTranscoder) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcJsonReverseTranscoder.ProtoReflect.Descriptor instead.
func (*GrpcJsonReverseTranscoder) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_rawDescGZIP(), []int{0}
}

func (x *GrpcJsonReverseTranscoder) GetDescriptorPath() string {
	if x != nil {
		return x.DescriptorPath
	}
	return ""
}

func (x *GrpcJsonReverseTranscoder) GetDescriptorBinary() []byte {
	if x != nil {
		return x.DescriptorBinary
	}
	return nil
}

func (x *GrpcJsonReverseTranscoder) GetMaxRequestBodySize() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxRequestBodySize
	}
	return nil
}

func (x *GrpcJsonReverseTranscoder) GetMaxResponseBodySize() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxResponseBodySize
	}
	return nil
}

func (x *GrpcJsonReverseTranscoder) GetApiVersionHeader() string {
	if x != nil {
		return x.ApiVersionHeader
	}
	return ""
}

var File_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_rawDesc = []byte{
	0x0a, 0x4e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x3d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x02, 0x0a, 0x19, 0x47, 0x72, 0x70, 0x63,
	0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x63, 0x6f, 0x64, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2b,
	0x0a, 0x11, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x62, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x58, 0x0a, 0x15, 0x6d,
	0x61, 0x78, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20,
	0x00, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64,
	0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x5a, 0x0a, 0x16, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x13, 0x6d, 0x61,
	0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61,
	0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42,
	0xee, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x4b, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6a, 0x73, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63,
	0x6f, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f,
	0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x83, 0x01, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2f, 0x76,
	0x33, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x76, 0x33,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_rawDescData = file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_rawDesc
)

func file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_rawDescData
}

var file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_goTypes = []interface{}{
	(*GrpcJsonReverseTranscoder)(nil), // 0: envoy.extensions.filters.http.grpc_json_reverse_transcoder.v3.GrpcJsonReverseTranscoder
	(*wrapperspb.UInt32Value)(nil),    // 1: google.protobuf.UInt32Value
}
var file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.http.grpc_json_reverse_transcoder.v3.GrpcJsonReverseTranscoder.max_request_body_size:type_name -> google.protobuf.UInt32Value
	1, // 1: envoy.extensions.filters.http.grpc_json_reverse_transcoder.v3.GrpcJsonReverseTranscoder.max_response_body_size:type_name -> google.protobuf.UInt32Value
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_init()
}
func file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_init() {
	if File_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcJsonReverseTranscoder); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto = out.File
	file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_rawDesc = nil
	file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_goTypes = nil
	file_envoy_extensions_filters_http_grpc_json_reverse_transcoder_v3_transcoder_proto_depIdxs = nil
}
