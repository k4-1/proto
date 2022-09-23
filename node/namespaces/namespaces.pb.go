//
//Copyright © 2022 Infinite Devices GmbH
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: node/namespaces/namespaces.proto

package namespaces

import (
	access "github.com/infinimesh/proto/node/access"
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

type Namespace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   string         `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title  string         `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Access *access.Access `protobuf:"bytes,3,opt,name=access,proto3,oneof" json:"access,omitempty"`
	Plugin *string        `protobuf:"bytes,4,opt,name=plugin,proto3,oneof" json:"plugin,omitempty"`
}

func (x *Namespace) Reset() {
	*x = Namespace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_namespaces_namespaces_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Namespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Namespace) ProtoMessage() {}

func (x *Namespace) ProtoReflect() protoreflect.Message {
	mi := &file_node_namespaces_namespaces_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Namespace.ProtoReflect.Descriptor instead.
func (*Namespace) Descriptor() ([]byte, []int) {
	return file_node_namespaces_namespaces_proto_rawDescGZIP(), []int{0}
}

func (x *Namespace) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Namespace) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Namespace) GetAccess() *access.Access {
	if x != nil {
		return x.Access
	}
	return nil
}

func (x *Namespace) GetPlugin() string {
	if x != nil && x.Plugin != nil {
		return *x.Plugin
	}
	return ""
}

type Namespaces struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespaces []*Namespace `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
}

func (x *Namespaces) Reset() {
	*x = Namespaces{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_namespaces_namespaces_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Namespaces) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Namespaces) ProtoMessage() {}

func (x *Namespaces) ProtoReflect() protoreflect.Message {
	mi := &file_node_namespaces_namespaces_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Namespaces.ProtoReflect.Descriptor instead.
func (*Namespaces) Descriptor() ([]byte, []int) {
	return file_node_namespaces_namespaces_proto_rawDescGZIP(), []int{1}
}

func (x *Namespaces) GetNamespaces() []*Namespace {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

var File_node_namespaces_namespaces_proto protoreflect.FileDescriptor

var file_node_namespaces_namespaces_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x1a, 0x18,
	0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x09, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x3b, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x48, 0x00, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a,
	0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x22, 0x53, 0x0a, 0x0a, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x45,
	0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x42, 0xe8, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e,
	0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x42, 0x0f, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x49, 0x4e, 0x4e, 0xaa, 0x02,
	0x1a, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0xca, 0x02, 0x1a, 0x49, 0x6e,
	0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x4e, 0x6f, 0x64, 0x65, 0x5c, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0xe2, 0x02, 0x26, 0x49, 0x6e, 0x66, 0x69, 0x6e,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x4e, 0x6f, 0x64, 0x65, 0x5c, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x1c, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x3a, 0x3a,
	0x4e, 0x6f, 0x64, 0x65, 0x3a, 0x3a, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_namespaces_namespaces_proto_rawDescOnce sync.Once
	file_node_namespaces_namespaces_proto_rawDescData = file_node_namespaces_namespaces_proto_rawDesc
)

func file_node_namespaces_namespaces_proto_rawDescGZIP() []byte {
	file_node_namespaces_namespaces_proto_rawDescOnce.Do(func() {
		file_node_namespaces_namespaces_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_namespaces_namespaces_proto_rawDescData)
	})
	return file_node_namespaces_namespaces_proto_rawDescData
}

var file_node_namespaces_namespaces_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_node_namespaces_namespaces_proto_goTypes = []interface{}{
	(*Namespace)(nil),     // 0: infinimesh.node.namespaces.Namespace
	(*Namespaces)(nil),    // 1: infinimesh.node.namespaces.Namespaces
	(*access.Access)(nil), // 2: infinimesh.node.access.Access
}
var file_node_namespaces_namespaces_proto_depIdxs = []int32{
	2, // 0: infinimesh.node.namespaces.Namespace.access:type_name -> infinimesh.node.access.Access
	0, // 1: infinimesh.node.namespaces.Namespaces.namespaces:type_name -> infinimesh.node.namespaces.Namespace
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_node_namespaces_namespaces_proto_init() }
func file_node_namespaces_namespaces_proto_init() {
	if File_node_namespaces_namespaces_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_namespaces_namespaces_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Namespace); i {
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
		file_node_namespaces_namespaces_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Namespaces); i {
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
	file_node_namespaces_namespaces_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_node_namespaces_namespaces_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_node_namespaces_namespaces_proto_goTypes,
		DependencyIndexes: file_node_namespaces_namespaces_proto_depIdxs,
		MessageInfos:      file_node_namespaces_namespaces_proto_msgTypes,
	}.Build()
	File_node_namespaces_namespaces_proto = out.File
	file_node_namespaces_namespaces_proto_rawDesc = nil
	file_node_namespaces_namespaces_proto_goTypes = nil
	file_node_namespaces_namespaces_proto_depIdxs = nil
}
