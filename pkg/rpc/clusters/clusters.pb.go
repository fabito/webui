// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        (unknown)
// source: pkg/rpc/clusters/clusters.proto

package clusters

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Context struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Context) Reset() {
	*x = Context{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_clusters_clusters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Context) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Context) ProtoMessage() {}

func (x *Context) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_clusters_clusters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Context.ProtoReflect.Descriptor instead.
func (*Context) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_clusters_clusters_proto_rawDescGZIP(), []int{0}
}

func (x *Context) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListContextsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListContextsReq) Reset() {
	*x = ListContextsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_clusters_clusters_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListContextsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListContextsReq) ProtoMessage() {}

func (x *ListContextsReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_clusters_clusters_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListContextsReq.ProtoReflect.Descriptor instead.
func (*ListContextsReq) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_clusters_clusters_proto_rawDescGZIP(), []int{1}
}

type ListContextsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contexts []*Context `protobuf:"bytes,1,rep,name=contexts,proto3" json:"contexts,omitempty"`
}

func (x *ListContextsRes) Reset() {
	*x = ListContextsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_clusters_clusters_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListContextsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListContextsRes) ProtoMessage() {}

func (x *ListContextsRes) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_clusters_clusters_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListContextsRes.ProtoReflect.Descriptor instead.
func (*ListContextsRes) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_clusters_clusters_proto_rawDescGZIP(), []int{2}
}

func (x *ListContextsRes) GetContexts() []*Context {
	if x != nil {
		return x.Contexts
	}
	return nil
}

type Kustomization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Kustomization) Reset() {
	*x = Kustomization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_clusters_clusters_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kustomization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kustomization) ProtoMessage() {}

func (x *Kustomization) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_clusters_clusters_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kustomization.ProtoReflect.Descriptor instead.
func (*Kustomization) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_clusters_clusters_proto_rawDescGZIP(), []int{3}
}

func (x *Kustomization) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListKustomizationsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContextName string `protobuf:"bytes,1,opt,name=contextName,proto3" json:"contextName,omitempty"`
}

func (x *ListKustomizationsReq) Reset() {
	*x = ListKustomizationsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_clusters_clusters_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKustomizationsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKustomizationsReq) ProtoMessage() {}

func (x *ListKustomizationsReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_clusters_clusters_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKustomizationsReq.ProtoReflect.Descriptor instead.
func (*ListKustomizationsReq) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_clusters_clusters_proto_rawDescGZIP(), []int{4}
}

func (x *ListKustomizationsReq) GetContextName() string {
	if x != nil {
		return x.ContextName
	}
	return ""
}

type ListKustomizationsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kustomizations []*Kustomization `protobuf:"bytes,1,rep,name=kustomizations,proto3" json:"kustomizations,omitempty"`
}

func (x *ListKustomizationsRes) Reset() {
	*x = ListKustomizationsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_clusters_clusters_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKustomizationsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKustomizationsRes) ProtoMessage() {}

func (x *ListKustomizationsRes) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_clusters_clusters_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKustomizationsRes.ProtoReflect.Descriptor instead.
func (*ListKustomizationsRes) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_clusters_clusters_proto_rawDescGZIP(), []int{5}
}

func (x *ListKustomizationsRes) GetKustomizations() []*Kustomization {
	if x != nil {
		return x.Kustomizations
	}
	return nil
}

var File_pkg_rpc_clusters_clusters_proto protoreflect.FileDescriptor

var file_pkg_rpc_clusters_clusters_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x1d, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x52, 0x65, 0x71, 0x22, 0x40, 0x0a,
	0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x12, 0x2d, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x22,
	0x23, 0x0a, 0x0d, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x58, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0e, 0x6b, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x4b, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x6b, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xa8, 0x01, 0x0a, 0x08, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x44, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x1a, 0x19, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x52, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1f, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_rpc_clusters_clusters_proto_rawDescOnce sync.Once
	file_pkg_rpc_clusters_clusters_proto_rawDescData = file_pkg_rpc_clusters_clusters_proto_rawDesc
)

func file_pkg_rpc_clusters_clusters_proto_rawDescGZIP() []byte {
	file_pkg_rpc_clusters_clusters_proto_rawDescOnce.Do(func() {
		file_pkg_rpc_clusters_clusters_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_rpc_clusters_clusters_proto_rawDescData)
	})
	return file_pkg_rpc_clusters_clusters_proto_rawDescData
}

var file_pkg_rpc_clusters_clusters_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_rpc_clusters_clusters_proto_goTypes = []interface{}{
	(*Context)(nil),               // 0: clusters.Context
	(*ListContextsReq)(nil),       // 1: clusters.ListContextsReq
	(*ListContextsRes)(nil),       // 2: clusters.ListContextsRes
	(*Kustomization)(nil),         // 3: clusters.Kustomization
	(*ListKustomizationsReq)(nil), // 4: clusters.ListKustomizationsReq
	(*ListKustomizationsRes)(nil), // 5: clusters.ListKustomizationsRes
}
var file_pkg_rpc_clusters_clusters_proto_depIdxs = []int32{
	0, // 0: clusters.ListContextsRes.contexts:type_name -> clusters.Context
	3, // 1: clusters.ListKustomizationsRes.kustomizations:type_name -> clusters.Kustomization
	1, // 2: clusters.Clusters.ListContexts:input_type -> clusters.ListContextsReq
	4, // 3: clusters.Clusters.ListKustomizations:input_type -> clusters.ListKustomizationsReq
	2, // 4: clusters.Clusters.ListContexts:output_type -> clusters.ListContextsRes
	5, // 5: clusters.Clusters.ListKustomizations:output_type -> clusters.ListKustomizationsRes
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_rpc_clusters_clusters_proto_init() }
func file_pkg_rpc_clusters_clusters_proto_init() {
	if File_pkg_rpc_clusters_clusters_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_rpc_clusters_clusters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Context); i {
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
		file_pkg_rpc_clusters_clusters_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListContextsReq); i {
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
		file_pkg_rpc_clusters_clusters_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListContextsRes); i {
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
		file_pkg_rpc_clusters_clusters_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kustomization); i {
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
		file_pkg_rpc_clusters_clusters_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKustomizationsReq); i {
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
		file_pkg_rpc_clusters_clusters_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKustomizationsRes); i {
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
			RawDescriptor: file_pkg_rpc_clusters_clusters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_rpc_clusters_clusters_proto_goTypes,
		DependencyIndexes: file_pkg_rpc_clusters_clusters_proto_depIdxs,
		MessageInfos:      file_pkg_rpc_clusters_clusters_proto_msgTypes,
	}.Build()
	File_pkg_rpc_clusters_clusters_proto = out.File
	file_pkg_rpc_clusters_clusters_proto_rawDesc = nil
	file_pkg_rpc_clusters_clusters_proto_goTypes = nil
	file_pkg_rpc_clusters_clusters_proto_depIdxs = nil
}
