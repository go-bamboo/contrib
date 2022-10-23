// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: pkg/client/confluent/conf.proto

package confluent

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SASL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mechanisms string `protobuf:"bytes,1,opt,name=mechanisms,proto3" json:"mechanisms,omitempty"`
	User       string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Password   string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SASL) Reset() {
	*x = SASL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_confluent_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SASL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SASL) ProtoMessage() {}

func (x *SASL) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_confluent_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SASL.ProtoReflect.Descriptor instead.
func (*SASL) Descriptor() ([]byte, []int) {
	return file_pkg_client_confluent_conf_proto_rawDescGZIP(), []int{0}
}

func (x *SASL) GetMechanisms() string {
	if x != nil {
		return x.Mechanisms
	}
	return ""
}

func (x *SASL) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *SASL) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SSL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaLocation string `protobuf:"bytes,1,opt,name=caLocation,proto3" json:"caLocation,omitempty"`
	CaPem      string `protobuf:"bytes,2,opt,name=caPem,proto3" json:"caPem,omitempty"`
}

func (x *SSL) Reset() {
	*x = SSL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_confluent_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSL) ProtoMessage() {}

func (x *SSL) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_confluent_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSL.ProtoReflect.Descriptor instead.
func (*SSL) Descriptor() ([]byte, []int) {
	return file_pkg_client_confluent_conf_proto_rawDescGZIP(), []int{1}
}

func (x *SSL) GetCaLocation() string {
	if x != nil {
		return x.CaLocation
	}
	return ""
}

func (x *SSL) GetCaPem() string {
	if x != nil {
		return x.CaPem
	}
	return ""
}

type Conf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BootstrapServers string               `protobuf:"bytes,1,opt,name=bootstrapServers,proto3" json:"bootstrapServers,omitempty"`
	SecurityProtocol string               `protobuf:"bytes,2,opt,name=securityProtocol,proto3" json:"securityProtocol,omitempty"` //plaintext, ssl, sasl_plaintext, sasl_ssl
	Sasl             *SASL                `protobuf:"bytes,3,opt,name=sasl,proto3" json:"sasl,omitempty"`
	Ssl              *SSL                 `protobuf:"bytes,4,opt,name=ssl,proto3" json:"ssl,omitempty"`
	Topic            string               `protobuf:"bytes,5,opt,name=topic,proto3" json:"topic,omitempty"`
	Group            string               `protobuf:"bytes,6,opt,name=group,proto3" json:"group,omitempty"`
	ReadTimeout      *durationpb.Duration `protobuf:"bytes,7,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout     *durationpb.Duration `protobuf:"bytes,8,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
	Acks             int32                `protobuf:"varint,9,opt,name=acks,proto3" json:"acks,omitempty"`
}

func (x *Conf) Reset() {
	*x = Conf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_confluent_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conf) ProtoMessage() {}

func (x *Conf) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_confluent_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conf.ProtoReflect.Descriptor instead.
func (*Conf) Descriptor() ([]byte, []int) {
	return file_pkg_client_confluent_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Conf) GetBootstrapServers() string {
	if x != nil {
		return x.BootstrapServers
	}
	return ""
}

func (x *Conf) GetSecurityProtocol() string {
	if x != nil {
		return x.SecurityProtocol
	}
	return ""
}

func (x *Conf) GetSasl() *SASL {
	if x != nil {
		return x.Sasl
	}
	return nil
}

func (x *Conf) GetSsl() *SSL {
	if x != nil {
		return x.Ssl
	}
	return nil
}

func (x *Conf) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Conf) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Conf) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Conf) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

func (x *Conf) GetAcks() int32 {
	if x != nil {
		return x.Acks
	}
	return 0
}

var File_pkg_client_confluent_conf_proto protoreflect.FileDescriptor

var file_pkg_client_confluent_conf_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x6c, 0x75, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x70, 0x6b, 0x67, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x74, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x04, 0x53, 0x41, 0x53, 0x4c, 0x12,
	0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x3b, 0x0a, 0x03, 0x53, 0x53, 0x4c, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x50, 0x65, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x50, 0x65, 0x6d, 0x22, 0xf9, 0x02, 0x0a,
	0x04, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x2a, 0x0a, 0x10, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72,
	0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x2e, 0x0a,
	0x04, 0x73, 0x61, 0x73, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65,
	0x6e, 0x74, 0x2e, 0x53, 0x41, 0x53, 0x4c, 0x52, 0x04, 0x73, 0x61, 0x73, 0x6c, 0x12, 0x2b, 0x0a,
	0x03, 0x73, 0x73, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6b, 0x67,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e,
	0x74, 0x2e, 0x53, 0x53, 0x4c, 0x52, 0x03, 0x73, 0x73, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x61, 0x63, 0x6b, 0x73, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x67, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65,
	0x6e, 0x74, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_client_confluent_conf_proto_rawDescOnce sync.Once
	file_pkg_client_confluent_conf_proto_rawDescData = file_pkg_client_confluent_conf_proto_rawDesc
)

func file_pkg_client_confluent_conf_proto_rawDescGZIP() []byte {
	file_pkg_client_confluent_conf_proto_rawDescOnce.Do(func() {
		file_pkg_client_confluent_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_client_confluent_conf_proto_rawDescData)
	})
	return file_pkg_client_confluent_conf_proto_rawDescData
}

var file_pkg_client_confluent_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_client_confluent_conf_proto_goTypes = []interface{}{
	(*SASL)(nil),                // 0: pkg.client.confluent.SASL
	(*SSL)(nil),                 // 1: pkg.client.confluent.SSL
	(*Conf)(nil),                // 2: pkg.client.confluent.Conf
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
}
var file_pkg_client_confluent_conf_proto_depIdxs = []int32{
	0, // 0: pkg.client.confluent.Conf.sasl:type_name -> pkg.client.confluent.SASL
	1, // 1: pkg.client.confluent.Conf.ssl:type_name -> pkg.client.confluent.SSL
	3, // 2: pkg.client.confluent.Conf.read_timeout:type_name -> google.protobuf.Duration
	3, // 3: pkg.client.confluent.Conf.write_timeout:type_name -> google.protobuf.Duration
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_client_confluent_conf_proto_init() }
func file_pkg_client_confluent_conf_proto_init() {
	if File_pkg_client_confluent_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_client_confluent_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SASL); i {
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
		file_pkg_client_confluent_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSL); i {
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
		file_pkg_client_confluent_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conf); i {
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
			RawDescriptor: file_pkg_client_confluent_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_client_confluent_conf_proto_goTypes,
		DependencyIndexes: file_pkg_client_confluent_conf_proto_depIdxs,
		MessageInfos:      file_pkg_client_confluent_conf_proto_msgTypes,
	}.Build()
	File_pkg_client_confluent_conf_proto = out.File
	file_pkg_client_confluent_conf_proto_rawDesc = nil
	file_pkg_client_confluent_conf_proto_goTypes = nil
	file_pkg_client_confluent_conf_proto_depIdxs = nil
}