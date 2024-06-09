// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0--rc3
// source: skynx/protobuf/network/v1/sxsp/session.proto

package sxsp

import (
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

type SessionMsgType int32

const (
	SessionMsgType_UNDEFINED_SESSION_MSG SessionMsgType = 0
	SessionMsgType_SESSION_KEEPALIVE     SessionMsgType = 11
)

// Enum value maps for SessionMsgType.
var (
	SessionMsgType_name = map[int32]string{
		0:  "UNDEFINED_SESSION_MSG",
		11: "SESSION_KEEPALIVE",
	}
	SessionMsgType_value = map[string]int32{
		"UNDEFINED_SESSION_MSG": 0,
		"SESSION_KEEPALIVE":     11,
	}
)

func (x SessionMsgType) Enum() *SessionMsgType {
	p := new(SessionMsgType)
	*p = x
	return p
}

func (x SessionMsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SessionMsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_skynx_protobuf_network_v1_sxsp_session_proto_enumTypes[0].Descriptor()
}

func (SessionMsgType) Type() protoreflect.EnumType {
	return &file_skynx_protobuf_network_v1_sxsp_session_proto_enumTypes[0]
}

func (x SessionMsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SessionMsgType.Descriptor instead.
func (SessionMsgType) EnumDescriptor() ([]byte, []int) {
	return file_skynx_protobuf_network_v1_sxsp_session_proto_rawDescGZIP(), []int{0}
}

type SessionPDU struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      SessionMsgType `protobuf:"varint,11,opt,name=type,proto3,enum=sxsp.SessionMsgType" json:"type,omitempty"`
	SessionID string         `protobuf:"bytes,21,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
}

func (x *SessionPDU) Reset() {
	*x = SessionPDU{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_network_v1_sxsp_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionPDU) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionPDU) ProtoMessage() {}

func (x *SessionPDU) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_network_v1_sxsp_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionPDU.ProtoReflect.Descriptor instead.
func (*SessionPDU) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_network_v1_sxsp_session_proto_rawDescGZIP(), []int{0}
}

func (x *SessionPDU) GetType() SessionMsgType {
	if x != nil {
		return x.Type
	}
	return SessionMsgType_UNDEFINED_SESSION_MSG
}

func (x *SessionPDU) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

var File_skynx_protobuf_network_v1_sxsp_session_proto protoreflect.FileDescriptor

var file_skynx_protobuf_network_v1_sxsp_session_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x78, 0x73, 0x70,
	0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x73, 0x78, 0x73, 0x70, 0x22, 0x54, 0x0a, 0x0a, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50,
	0x44, 0x55, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x73, 0x78, 0x73, 0x70, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d,
	0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x2a, 0x42, 0x0a, 0x0e, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15,
	0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x4d, 0x53, 0x47, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x45, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x4b, 0x45, 0x45, 0x50, 0x41, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x0b, 0x42, 0x25,
	0x5a, 0x23, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x2d, 0x61, 0x70, 0x69,
	0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x73, 0x78, 0x73, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_skynx_protobuf_network_v1_sxsp_session_proto_rawDescOnce sync.Once
	file_skynx_protobuf_network_v1_sxsp_session_proto_rawDescData = file_skynx_protobuf_network_v1_sxsp_session_proto_rawDesc
)

func file_skynx_protobuf_network_v1_sxsp_session_proto_rawDescGZIP() []byte {
	file_skynx_protobuf_network_v1_sxsp_session_proto_rawDescOnce.Do(func() {
		file_skynx_protobuf_network_v1_sxsp_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_skynx_protobuf_network_v1_sxsp_session_proto_rawDescData)
	})
	return file_skynx_protobuf_network_v1_sxsp_session_proto_rawDescData
}

var file_skynx_protobuf_network_v1_sxsp_session_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_skynx_protobuf_network_v1_sxsp_session_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_skynx_protobuf_network_v1_sxsp_session_proto_goTypes = []interface{}{
	(SessionMsgType)(0), // 0: sxsp.SessionMsgType
	(*SessionPDU)(nil),  // 1: sxsp.SessionPDU
}
var file_skynx_protobuf_network_v1_sxsp_session_proto_depIdxs = []int32{
	0, // 0: sxsp.SessionPDU.type:type_name -> sxsp.SessionMsgType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_skynx_protobuf_network_v1_sxsp_session_proto_init() }
func file_skynx_protobuf_network_v1_sxsp_session_proto_init() {
	if File_skynx_protobuf_network_v1_sxsp_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_skynx_protobuf_network_v1_sxsp_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionPDU); i {
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
			RawDescriptor: file_skynx_protobuf_network_v1_sxsp_session_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_skynx_protobuf_network_v1_sxsp_session_proto_goTypes,
		DependencyIndexes: file_skynx_protobuf_network_v1_sxsp_session_proto_depIdxs,
		EnumInfos:         file_skynx_protobuf_network_v1_sxsp_session_proto_enumTypes,
		MessageInfos:      file_skynx_protobuf_network_v1_sxsp_session_proto_msgTypes,
	}.Build()
	File_skynx_protobuf_network_v1_sxsp_session_proto = out.File
	file_skynx_protobuf_network_v1_sxsp_session_proto_rawDesc = nil
	file_skynx_protobuf_network_v1_sxsp_session_proto_goTypes = nil
	file_skynx_protobuf_network_v1_sxsp_session_proto_depIdxs = nil
}
