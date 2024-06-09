// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0--rc3
// source: skynx/protobuf/resources/v1/topology/proxy.proto

package topology

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

type Proxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationID string      `protobuf:"bytes,1,opt,name=locationID,proto3" json:"locationID,omitempty"`
	ProxyID    string      `protobuf:"bytes,5,opt,name=proxyID,proto3" json:"proxyID,omitempty"`
	Cfg        *ProxyCfg   `protobuf:"bytes,41,opt,name=cfg,proto3" json:"cfg,omitempty"`
	Agent      *ProxyAgent `protobuf:"bytes,51,opt,name=agent,proto3" json:"agent,omitempty"`
	IPv6       string      `protobuf:"bytes,81,opt,name=IPv6,proto3" json:"IPv6,omitempty"`
	LastSeen   int64       `protobuf:"varint,201,opt,name=lastSeen,proto3" json:"lastSeen,omitempty"`
}

func (x *Proxy) Reset() {
	*x = Proxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_topology_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proxy) ProtoMessage() {}

func (x *Proxy) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_topology_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proxy.ProtoReflect.Descriptor instead.
func (*Proxy) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_topology_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *Proxy) GetLocationID() string {
	if x != nil {
		return x.LocationID
	}
	return ""
}

func (x *Proxy) GetProxyID() string {
	if x != nil {
		return x.ProxyID
	}
	return ""
}

func (x *Proxy) GetCfg() *ProxyCfg {
	if x != nil {
		return x.Cfg
	}
	return nil
}

func (x *Proxy) GetAgent() *ProxyAgent {
	if x != nil {
		return x.Agent
	}
	return nil
}

func (x *Proxy) GetIPv6() string {
	if x != nil {
		return x.IPv6
	}
	return ""
}

func (x *Proxy) GetLastSeen() int64 {
	if x != nil {
		return x.LastSeen
	}
	return 0
}

type ProxyCfg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProxyName   string `protobuf:"bytes,1,opt,name=proxyName,proto3" json:"proxyName,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Priority    int32  `protobuf:"varint,21,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *ProxyCfg) Reset() {
	*x = ProxyCfg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_topology_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyCfg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyCfg) ProtoMessage() {}

func (x *ProxyCfg) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_topology_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyCfg.ProtoReflect.Descriptor instead.
func (*ProxyCfg) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_topology_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *ProxyCfg) GetProxyName() string {
	if x != nil {
		return x.ProxyName
	}
	return ""
}

func (x *ProxyCfg) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProxyCfg) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type ProxyAgent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P2PHostID    string   `protobuf:"bytes,1,opt,name=P2PHostID,proto3" json:"P2PHostID,omitempty"`
	Hostname     string   `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Port         int32    `protobuf:"varint,11,opt,name=port,proto3" json:"port,omitempty"` // string transport = 12;
	ExternalIPv4 string   `protobuf:"bytes,21,opt,name=externalIPv4,proto3" json:"externalIPv4,omitempty"`
	MAddrs       []string `protobuf:"bytes,31,rep,name=MAddrs,proto3" json:"MAddrs,omitempty"`
	Version      string   `protobuf:"bytes,1000,opt,name=version,proto3" json:"version,omitempty"`
	DevMode      bool     `protobuf:"varint,1001,opt,name=devMode,proto3" json:"devMode,omitempty"`
}

func (x *ProxyAgent) Reset() {
	*x = ProxyAgent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_topology_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyAgent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyAgent) ProtoMessage() {}

func (x *ProxyAgent) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_topology_proxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyAgent.ProtoReflect.Descriptor instead.
func (*ProxyAgent) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_topology_proxy_proto_rawDescGZIP(), []int{2}
}

func (x *ProxyAgent) GetP2PHostID() string {
	if x != nil {
		return x.P2PHostID
	}
	return ""
}

func (x *ProxyAgent) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *ProxyAgent) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ProxyAgent) GetExternalIPv4() string {
	if x != nil {
		return x.ExternalIPv4
	}
	return ""
}

func (x *ProxyAgent) GetMAddrs() []string {
	if x != nil {
		return x.MAddrs
	}
	return nil
}

func (x *ProxyAgent) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ProxyAgent) GetDevMode() bool {
	if x != nil {
		return x.DevMode
	}
	return false
}

var File_skynx_protobuf_resources_v1_topology_proxy_proto protoreflect.FileDescriptor

var file_skynx_protobuf_resources_v1_topology_proxy_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f,
	0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x22, 0xc4, 0x01, 0x0a,
	0x05, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x49,
	0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x49, 0x44,
	0x12, 0x24, 0x0a, 0x03, 0x63, 0x66, 0x67, 0x18, 0x29, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x66,
	0x67, 0x52, 0x03, 0x63, 0x66, 0x67, 0x12, 0x2a, 0x0a, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18,
	0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79,
	0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x50, 0x76, 0x36, 0x18, 0x51, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x49, 0x50, 0x76, 0x36, 0x12, 0x1b, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65,
	0x65, 0x6e, 0x18, 0xc9, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x53,
	0x65, 0x65, 0x6e, 0x22, 0x66, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x66, 0x67, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0xcc, 0x01, 0x0a, 0x0a,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x32,
	0x50, 0x48, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50,
	0x32, 0x50, 0x48, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x49, 0x50, 0x76, 0x34, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x50, 0x76, 0x34, 0x12, 0x16, 0x0a, 0x06,
	0x4d, 0x41, 0x64, 0x64, 0x72, 0x73, 0x18, 0x1f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x41,
	0x64, 0x64, 0x72, 0x73, 0x12, 0x19, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0xe8, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x19, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x64, 0x65, 0x76, 0x4d, 0x6f, 0x64, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x73, 0x6b,
	0x79, 0x6e, 0x78, 0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x74,
	0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_skynx_protobuf_resources_v1_topology_proxy_proto_rawDescOnce sync.Once
	file_skynx_protobuf_resources_v1_topology_proxy_proto_rawDescData = file_skynx_protobuf_resources_v1_topology_proxy_proto_rawDesc
)

func file_skynx_protobuf_resources_v1_topology_proxy_proto_rawDescGZIP() []byte {
	file_skynx_protobuf_resources_v1_topology_proxy_proto_rawDescOnce.Do(func() {
		file_skynx_protobuf_resources_v1_topology_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_skynx_protobuf_resources_v1_topology_proxy_proto_rawDescData)
	})
	return file_skynx_protobuf_resources_v1_topology_proxy_proto_rawDescData
}

var file_skynx_protobuf_resources_v1_topology_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_skynx_protobuf_resources_v1_topology_proxy_proto_goTypes = []interface{}{
	(*Proxy)(nil),      // 0: topology.Proxy
	(*ProxyCfg)(nil),   // 1: topology.ProxyCfg
	(*ProxyAgent)(nil), // 2: topology.ProxyAgent
}
var file_skynx_protobuf_resources_v1_topology_proxy_proto_depIdxs = []int32{
	1, // 0: topology.Proxy.cfg:type_name -> topology.ProxyCfg
	2, // 1: topology.Proxy.agent:type_name -> topology.ProxyAgent
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_skynx_protobuf_resources_v1_topology_proxy_proto_init() }
func file_skynx_protobuf_resources_v1_topology_proxy_proto_init() {
	if File_skynx_protobuf_resources_v1_topology_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_skynx_protobuf_resources_v1_topology_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proxy); i {
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
		file_skynx_protobuf_resources_v1_topology_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyCfg); i {
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
		file_skynx_protobuf_resources_v1_topology_proxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyAgent); i {
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
			RawDescriptor: file_skynx_protobuf_resources_v1_topology_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_skynx_protobuf_resources_v1_topology_proxy_proto_goTypes,
		DependencyIndexes: file_skynx_protobuf_resources_v1_topology_proxy_proto_depIdxs,
		MessageInfos:      file_skynx_protobuf_resources_v1_topology_proxy_proto_msgTypes,
	}.Build()
	File_skynx_protobuf_resources_v1_topology_proxy_proto = out.File
	file_skynx_protobuf_resources_v1_topology_proxy_proto_rawDesc = nil
	file_skynx_protobuf_resources_v1_topology_proxy_proto_goTypes = nil
	file_skynx_protobuf_resources_v1_topology_proxy_proto_depIdxs = nil
}
