// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: contrib/envoy/extensions/filters/http/sxg/v3alpha/sxg.proto

package v3alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	duration "github.com/golang/protobuf/ptypes/duration"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// [#next-free-field: 10]
type SXG struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The SDS configuration for the public key data for the SSL certificate that will be used to sign the
	// SXG response.
	Certificate *v3.SdsSecretConfig `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	// The SDS configuration for the private key data for the SSL certificate that will be used to sign the
	// SXG response.
	PrivateKey *v3.SdsSecretConfig `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// The duration for which the generated SXG package will be valid. Default is 604800s (7 days in seconds).
	// Note that in order to account for clock skew, the timestamp will be backdated by a day. So, if duration
	// is set to 7 days, that will be 7 days from 24 hours ago (6 days from now). Also note that while 6/7 days
	// is appropriate for most content, if the downstream service is serving Javascript, or HTML with inline
	// Javascript, 1 day (so, with backdated expiry, 2 days, or 172800 seconds) is more appropriate.
	Duration *duration.Duration `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	// The SXG response payload is Merkle Integrity Content Encoding (MICE) encoded (specification is [here](https://datatracker.ietf.org/doc/html/draft-thomson-http-mice-03))
	// This value indicates the record size in the encoded payload. The default value is 4096.
	MiRecordSize uint64 `protobuf:"varint,4,opt,name=mi_record_size,json=miRecordSize,proto3" json:"mi_record_size,omitempty"`
	// The URI of certificate CBOR file published. Since it is required that the certificate CBOR file
	// be served from the same domain as the SXG document, this should be a relative URI.
	CborUrl string `protobuf:"bytes,5,opt,name=cbor_url,json=cborUrl,proto3" json:"cbor_url,omitempty"`
	// URL to retrieve validity data for signature, a CBOR map. See specification [here](https://tools.ietf.org/html/draft-yasskin-httpbis-origin-signed-exchanges-impl-00#section-3.6)
	ValidityUrl string `protobuf:"bytes,6,opt,name=validity_url,json=validityUrl,proto3" json:"validity_url,omitempty"`
	// Header that will be set if it is determined that the client can accept SXG (typically “accept: application/signed-exchange;v=b3“)
	// If not set, filter will default to: “x-client-can-accept-sxg“
	ClientCanAcceptSxgHeader string `protobuf:"bytes,7,opt,name=client_can_accept_sxg_header,json=clientCanAcceptSxgHeader,proto3" json:"client_can_accept_sxg_header,omitempty"`
	// Header set by downstream service to signal that the response should be transformed to SXG If not set,
	// filter will default to: “x-should-encode-sxg“
	ShouldEncodeSxgHeader string `protobuf:"bytes,8,opt,name=should_encode_sxg_header,json=shouldEncodeSxgHeader,proto3" json:"should_encode_sxg_header,omitempty"`
	// Headers that will be stripped from the SXG document, by listing a prefix (i.e. “x-custom-“ will cause
	// all headers prefixed by “x-custom-“ to be omitted from the SXG document)
	HeaderPrefixFilters []string `protobuf:"bytes,9,rep,name=header_prefix_filters,json=headerPrefixFilters,proto3" json:"header_prefix_filters,omitempty"`
}

func (x *SXG) Reset() {
	*x = SXG{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SXG) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SXG) ProtoMessage() {}

func (x *SXG) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SXG.ProtoReflect.Descriptor instead.
func (*SXG) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_rawDescGZIP(), []int{0}
}

func (x *SXG) GetCertificate() *v3.SdsSecretConfig {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *SXG) GetPrivateKey() *v3.SdsSecretConfig {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *SXG) GetDuration() *duration.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *SXG) GetMiRecordSize() uint64 {
	if x != nil {
		return x.MiRecordSize
	}
	return 0
}

func (x *SXG) GetCborUrl() string {
	if x != nil {
		return x.CborUrl
	}
	return ""
}

func (x *SXG) GetValidityUrl() string {
	if x != nil {
		return x.ValidityUrl
	}
	return ""
}

func (x *SXG) GetClientCanAcceptSxgHeader() string {
	if x != nil {
		return x.ClientCanAcceptSxgHeader
	}
	return ""
}

func (x *SXG) GetShouldEncodeSxgHeader() string {
	if x != nil {
		return x.ShouldEncodeSxgHeader
	}
	return ""
}

func (x *SXG) GetHeaderPrefixFilters() []string {
	if x != nil {
		return x.HeaderPrefixFilters
	}
	return nil
}

var File_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto protoreflect.FileDescriptor

var file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x73, 0x78, 0x67, 0x2f, 0x76, 0x33, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x73, 0x78, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x78, 0x67,
	0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x36, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x6c, 0x73,
	0x2f, 0x76, 0x33, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x04, 0x0a, 0x03, 0x53, 0x58, 0x47,
	0x12, 0x5c, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x76,
	0x33, 0x2e, 0x53, 0x64, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x5b,
	0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x76, 0x33, 0x2e,
	0x53, 0x64, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x69, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6d, 0x69, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x62, 0x6f, 0x72,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72,
	0x05, 0x10, 0x01, 0x3a, 0x01, 0x2f, 0x52, 0x07, 0x63, 0x62, 0x6f, 0x72, 0x55, 0x72, 0x6c, 0x12,
	0x2d, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x3a, 0x01,
	0x2f, 0x52, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x4e,
	0x0a, 0x1c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x6e, 0x5f, 0x61, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x5f, 0x73, 0x78, 0x67, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x72, 0x09, 0xc8, 0x01, 0x00, 0xd0, 0x01,
	0x01, 0xc0, 0x01, 0x01, 0x52, 0x18, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6e, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x53, 0x78, 0x67, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x47,
	0x0a, 0x18, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x73, 0x78, 0x67, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x72, 0x09, 0xc8, 0x01, 0x00, 0xd0, 0x01, 0x01, 0xc0, 0x01, 0x01,
	0x52, 0x15, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x78,
	0x67, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x15, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x42, 0x10, 0xfa, 0x42, 0x0d, 0x92, 0x01, 0x0a, 0x22, 0x08,
	0x72, 0x06, 0xc8, 0x01, 0x00, 0xc0, 0x01, 0x01, 0x52, 0x13, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x42, 0xa9, 0x01,
	0xba, 0x80, 0xc8, 0xd1, 0x06, 0x04, 0x08, 0x01, 0x10, 0x02, 0x0a, 0x37, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x78, 0x67, 0x2e, 0x76, 0x33, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x42, 0x08, 0x53, 0x78, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x73, 0x78,
	0x67, 0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_rawDescOnce sync.Once
	file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_rawDescData = file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_rawDesc
)

func file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_rawDescGZIP() []byte {
	file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_rawDescOnce.Do(func() {
		file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_rawDescData = protoimpl.X.CompressGZIP(file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_rawDescData)
	})
	return file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_rawDescData
}

var file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_goTypes = []interface{}{
	(*SXG)(nil),                // 0: envoy.extensions.filters.http.sxg.v3alpha.SXG
	(*v3.SdsSecretConfig)(nil), // 1: envoy.extensions.transport_sockets.tls.v3.SdsSecretConfig
	(*duration.Duration)(nil),  // 2: google.protobuf.Duration
}
var file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.http.sxg.v3alpha.SXG.certificate:type_name -> envoy.extensions.transport_sockets.tls.v3.SdsSecretConfig
	1, // 1: envoy.extensions.filters.http.sxg.v3alpha.SXG.private_key:type_name -> envoy.extensions.transport_sockets.tls.v3.SdsSecretConfig
	2, // 2: envoy.extensions.filters.http.sxg.v3alpha.SXG.duration:type_name -> google.protobuf.Duration
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_init() }
func file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_init() {
	if File_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SXG); i {
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
			RawDescriptor: file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_goTypes,
		DependencyIndexes: file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_depIdxs,
		MessageInfos:      file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_msgTypes,
	}.Build()
	File_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto = out.File
	file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_rawDesc = nil
	file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_goTypes = nil
	file_contrib_envoy_extensions_filters_http_sxg_v3alpha_sxg_proto_depIdxs = nil
}
