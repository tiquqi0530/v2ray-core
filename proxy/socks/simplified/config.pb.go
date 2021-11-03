// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proxy/socks/simplified/config.proto

package simplified

import (
	net "github.com/v2fly/v2ray-core/v4/common/net"
	_ "github.com/v2fly/v2ray-core/v4/common/protoext"
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

type ServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    *net.IPOrDomain `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	UdpEnabled bool            `protobuf:"varint,4,opt,name=udp_enabled,json=udpEnabled,proto3" json:"udp_enabled,omitempty"`
}

func (x *ServerConfig) Reset() {
	*x = ServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_socks_simplified_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig) ProtoMessage() {}

func (x *ServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_socks_simplified_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig.ProtoReflect.Descriptor instead.
func (*ServerConfig) Descriptor() ([]byte, []int) {
	return file_proxy_socks_simplified_config_proto_rawDescGZIP(), []int{0}
}

func (x *ServerConfig) GetAddress() *net.IPOrDomain {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *ServerConfig) GetUdpEnabled() bool {
	if x != nil {
		return x.UdpEnabled
	}
	return false
}

type ClientConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *net.IPOrDomain `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port    uint32          `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ClientConfig) Reset() {
	*x = ClientConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_socks_simplified_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConfig) ProtoMessage() {}

func (x *ClientConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_socks_simplified_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConfig.ProtoReflect.Descriptor instead.
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return file_proxy_socks_simplified_config_proto_rawDescGZIP(), []int{1}
}

func (x *ClientConfig) GetAddress() *net.IPOrDomain {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *ClientConfig) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_proxy_socks_simplified_config_proto protoreflect.FileDescriptor

var file_proxy_socks_simplified_config_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x2f, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x2e, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x49,
	0x50, 0x4f, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x64, 0x70, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x75, 0x64, 0x70, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x3a, 0x18, 0x82, 0xb5, 0x18, 0x09, 0x0a, 0x07, 0x69, 0x6e, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x82, 0xb5, 0x18, 0x07, 0x12, 0x05, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x7a, 0x0a,
	0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x49, 0x50, 0x4f, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x19,
	0x82, 0xb5, 0x18, 0x0a, 0x0a, 0x08, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x82, 0xb5,
	0x18, 0x07, 0x12, 0x05, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x42, 0x84, 0x01, 0x0a, 0x25, 0x63, 0x6f,
	0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x34, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x73, 0x6f, 0x63, 0x6b,
	0x73, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0xaa, 0x02, 0x21, 0x56,
	0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x53, 0x6f, 0x63, 0x6b, 0x73, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proxy_socks_simplified_config_proto_rawDescOnce sync.Once
	file_proxy_socks_simplified_config_proto_rawDescData = file_proxy_socks_simplified_config_proto_rawDesc
)

func file_proxy_socks_simplified_config_proto_rawDescGZIP() []byte {
	file_proxy_socks_simplified_config_proto_rawDescOnce.Do(func() {
		file_proxy_socks_simplified_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_socks_simplified_config_proto_rawDescData)
	})
	return file_proxy_socks_simplified_config_proto_rawDescData
}

var file_proxy_socks_simplified_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proxy_socks_simplified_config_proto_goTypes = []interface{}{
	(*ServerConfig)(nil),   // 0: v2ray.core.proxy.socks.simplified.ServerConfig
	(*ClientConfig)(nil),   // 1: v2ray.core.proxy.socks.simplified.ClientConfig
	(*net.IPOrDomain)(nil), // 2: v2ray.core.common.net.IPOrDomain
}
var file_proxy_socks_simplified_config_proto_depIdxs = []int32{
	2, // 0: v2ray.core.proxy.socks.simplified.ServerConfig.address:type_name -> v2ray.core.common.net.IPOrDomain
	2, // 1: v2ray.core.proxy.socks.simplified.ClientConfig.address:type_name -> v2ray.core.common.net.IPOrDomain
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proxy_socks_simplified_config_proto_init() }
func file_proxy_socks_simplified_config_proto_init() {
	if File_proxy_socks_simplified_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_socks_simplified_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfig); i {
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
		file_proxy_socks_simplified_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConfig); i {
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
			RawDescriptor: file_proxy_socks_simplified_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proxy_socks_simplified_config_proto_goTypes,
		DependencyIndexes: file_proxy_socks_simplified_config_proto_depIdxs,
		MessageInfos:      file_proxy_socks_simplified_config_proto_msgTypes,
	}.Build()
	File_proxy_socks_simplified_config_proto = out.File
	file_proxy_socks_simplified_config_proto_rawDesc = nil
	file_proxy_socks_simplified_config_proto_goTypes = nil
	file_proxy_socks_simplified_config_proto_depIdxs = nil
}
