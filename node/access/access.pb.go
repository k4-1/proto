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
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: node/access/access.proto

package access

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

type Level int32

const (
	Level_NONE  Level = 0
	Level_READ  Level = 1
	Level_MGMT  Level = 2
	Level_ADMIN Level = 3
	Level_ROOT  Level = 4
)

// Enum value maps for Level.
var (
	Level_name = map[int32]string{
		0: "NONE",
		1: "READ",
		2: "MGMT",
		3: "ADMIN",
		4: "ROOT",
	}
	Level_value = map[string]int32{
		"NONE":  0,
		"READ":  1,
		"MGMT":  2,
		"ADMIN": 3,
		"ROOT":  4,
	}
)

func (x Level) Enum() *Level {
	p := new(Level)
	*p = x
	return p
}

func (x Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Level) Descriptor() protoreflect.EnumDescriptor {
	return file_node_access_access_proto_enumTypes[0].Descriptor()
}

func (Level) Type() protoreflect.EnumType {
	return &file_node_access_access_proto_enumTypes[0]
}

func (x Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Level.Descriptor instead.
func (Level) EnumDescriptor() ([]byte, []int) {
	return file_node_access_access_proto_rawDescGZIP(), []int{0}
}

type Role int32

const (
	Role_UNSET Role = 0
	Role_OWNER Role = 1
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "UNSET",
		1: "OWNER",
	}
	Role_value = map[string]int32{
		"UNSET": 0,
		"OWNER": 1,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_node_access_access_proto_enumTypes[1].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_node_access_access_proto_enumTypes[1]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_node_access_access_proto_rawDescGZIP(), []int{1}
}

type Access struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level     Level   `protobuf:"varint,1,opt,name=level,proto3,enum=infinimesh.node.access.Level" json:"level,omitempty"`
	Role      Role    `protobuf:"varint,2,opt,name=role,proto3,enum=infinimesh.node.access.Role" json:"role,omitempty"`
	Namespace *string `protobuf:"bytes,3,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
}

func (x *Access) Reset() {
	*x = Access{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_access_access_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Access) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Access) ProtoMessage() {}

func (x *Access) ProtoReflect() protoreflect.Message {
	mi := &file_node_access_access_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Access.ProtoReflect.Descriptor instead.
func (*Access) Descriptor() ([]byte, []int) {
	return file_node_access_access_proto_rawDescGZIP(), []int{0}
}

func (x *Access) GetLevel() Level {
	if x != nil {
		return x.Level
	}
	return Level_NONE
}

func (x *Access) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_UNSET
}

func (x *Access) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node   string `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Parent string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Edge   string `protobuf:"bytes,3,opt,name=edge,proto3" json:"edge,omitempty"` // @private
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_access_access_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_node_access_access_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_node_access_access_proto_rawDescGZIP(), []int{1}
}

func (x *Node) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *Node) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *Node) GetEdge() string {
	if x != nil {
		return x.Edge
	}
	return ""
}

type Nodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *Nodes) Reset() {
	*x = Nodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_access_access_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nodes) ProtoMessage() {}

func (x *Nodes) ProtoReflect() protoreflect.Message {
	mi := &file_node_access_access_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nodes.ProtoReflect.Descriptor instead.
func (*Nodes) Descriptor() ([]byte, []int) {
	return file_node_access_access_proto_rawDescGZIP(), []int{2}
}

func (x *Nodes) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

var File_node_access_access_proto protoreflect.FileDescriptor

var file_node_access_access_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x69, 0x6e, 0x66, 0x69,
	0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0xa0, 0x01, 0x0a, 0x06, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x69,
	0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x30, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1c, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x46, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x64, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x64, 0x67, 0x65, 0x22, 0x3b, 0x0a,
	0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2a, 0x3a, 0x0a, 0x05, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x47, 0x4d, 0x54, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04,
	0x52, 0x4f, 0x4f, 0x54, 0x10, 0x04, 0x2a, 0x1c, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x09,
	0x0a, 0x05, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x57, 0x4e,
	0x45, 0x52, 0x10, 0x01, 0x42, 0xcc, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x66,
	0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x42, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0xa2, 0x02, 0x03, 0x49, 0x4e,
	0x41, 0xaa, 0x02, 0x16, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0xca, 0x02, 0x16, 0x49, 0x6e, 0x66,
	0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x4e, 0x6f, 0x64, 0x65, 0x5c, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0xe2, 0x02, 0x22, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68,
	0x5c, 0x4e, 0x6f, 0x64, 0x65, 0x5c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x49, 0x6e, 0x66, 0x69, 0x6e,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x4e, 0x6f, 0x64, 0x65, 0x3a, 0x3a, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_access_access_proto_rawDescOnce sync.Once
	file_node_access_access_proto_rawDescData = file_node_access_access_proto_rawDesc
)

func file_node_access_access_proto_rawDescGZIP() []byte {
	file_node_access_access_proto_rawDescOnce.Do(func() {
		file_node_access_access_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_access_access_proto_rawDescData)
	})
	return file_node_access_access_proto_rawDescData
}

var file_node_access_access_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_node_access_access_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_node_access_access_proto_goTypes = []interface{}{
	(Level)(0),     // 0: infinimesh.node.access.Level
	(Role)(0),      // 1: infinimesh.node.access.Role
	(*Access)(nil), // 2: infinimesh.node.access.Access
	(*Node)(nil),   // 3: infinimesh.node.access.Node
	(*Nodes)(nil),  // 4: infinimesh.node.access.Nodes
}
var file_node_access_access_proto_depIdxs = []int32{
	0, // 0: infinimesh.node.access.Access.level:type_name -> infinimesh.node.access.Level
	1, // 1: infinimesh.node.access.Access.role:type_name -> infinimesh.node.access.Role
	3, // 2: infinimesh.node.access.Nodes.nodes:type_name -> infinimesh.node.access.Node
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_node_access_access_proto_init() }
func file_node_access_access_proto_init() {
	if File_node_access_access_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_access_access_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Access); i {
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
		file_node_access_access_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_node_access_access_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nodes); i {
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
	file_node_access_access_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_node_access_access_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_node_access_access_proto_goTypes,
		DependencyIndexes: file_node_access_access_proto_depIdxs,
		EnumInfos:         file_node_access_access_proto_enumTypes,
		MessageInfos:      file_node_access_access_proto_msgTypes,
	}.Build()
	File_node_access_access_proto = out.File
	file_node_access_access_proto_rawDesc = nil
	file_node_access_access_proto_goTypes = nil
	file_node_access_access_proto_depIdxs = nil
}
