// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: discovery.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for HelloWorld
type DiscoveryServiceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ServiceID     string                 `protobuf:"bytes,1,opt,name=ServiceID,proto3" json:"ServiceID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DiscoveryServiceRequest) Reset() {
	*x = DiscoveryServiceRequest{}
	mi := &file_discovery_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscoveryServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveryServiceRequest) ProtoMessage() {}

func (x *DiscoveryServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoveryServiceRequest.ProtoReflect.Descriptor instead.
func (*DiscoveryServiceRequest) Descriptor() ([]byte, []int) {
	return file_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *DiscoveryServiceRequest) GetServiceID() string {
	if x != nil {
		return x.ServiceID
	}
	return ""
}

// Response message for HelloWorld
type DiscoveryServiceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Host          string                 `protobuf:"bytes,1,opt,name=Host,proto3" json:"Host,omitempty"`
	Port          uint64                 `protobuf:"varint,2,opt,name=Port,proto3" json:"Port,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DiscoveryServiceResponse) Reset() {
	*x = DiscoveryServiceResponse{}
	mi := &file_discovery_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscoveryServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveryServiceResponse) ProtoMessage() {}

func (x *DiscoveryServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoveryServiceResponse.ProtoReflect.Descriptor instead.
func (*DiscoveryServiceResponse) Descriptor() ([]byte, []int) {
	return file_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *DiscoveryServiceResponse) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *DiscoveryServiceResponse) GetPort() uint64 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_discovery_proto protoreflect.FileDescriptor

const file_discovery_proto_rawDesc = "" +
	"\n" +
	"\x0fdiscovery.proto\x12\tdiscovery\"7\n" +
	"\x17DiscoveryServiceRequest\x12\x1c\n" +
	"\tServiceID\x18\x01 \x01(\tR\tServiceID\"B\n" +
	"\x18DiscoveryServiceResponse\x12\x12\n" +
	"\x04Host\x18\x01 \x01(\tR\x04Host\x12\x12\n" +
	"\x04Port\x18\x02 \x01(\x04R\x04Port2o\n" +
	"\x15DiscoveryProtoService\x12V\n" +
	"\tDiscovery\x12\".discovery.DiscoveryServiceRequest\x1a#.discovery.DiscoveryServiceResponse\"\x00B\x04Z\x02./b\x06proto3"

var (
	file_discovery_proto_rawDescOnce sync.Once
	file_discovery_proto_rawDescData []byte
)

func file_discovery_proto_rawDescGZIP() []byte {
	file_discovery_proto_rawDescOnce.Do(func() {
		file_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_discovery_proto_rawDesc), len(file_discovery_proto_rawDesc)))
	})
	return file_discovery_proto_rawDescData
}

var file_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_discovery_proto_goTypes = []any{
	(*DiscoveryServiceRequest)(nil),  // 0: discovery.DiscoveryServiceRequest
	(*DiscoveryServiceResponse)(nil), // 1: discovery.DiscoveryServiceResponse
}
var file_discovery_proto_depIdxs = []int32{
	0, // 0: discovery.DiscoveryProtoService.Discovery:input_type -> discovery.DiscoveryServiceRequest
	1, // 1: discovery.DiscoveryProtoService.Discovery:output_type -> discovery.DiscoveryServiceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_discovery_proto_init() }
func file_discovery_proto_init() {
	if File_discovery_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_discovery_proto_rawDesc), len(file_discovery_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_discovery_proto_goTypes,
		DependencyIndexes: file_discovery_proto_depIdxs,
		MessageInfos:      file_discovery_proto_msgTypes,
	}.Build()
	File_discovery_proto = out.File
	file_discovery_proto_goTypes = nil
	file_discovery_proto_depIdxs = nil
}
