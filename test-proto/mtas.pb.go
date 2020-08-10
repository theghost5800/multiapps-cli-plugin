// Copyright 2015 The gRPC Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: mtas.proto

package test_proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MtasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token            string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	OrganizationName string `protobuf:"bytes,2,opt,name=organizationName,proto3" json:"organizationName,omitempty"`
	SpaceName        string `protobuf:"bytes,3,opt,name=spaceName,proto3" json:"spaceName,omitempty"`
	ApiUrl           string `protobuf:"bytes,4,opt,name=apiUrl,proto3" json:"apiUrl,omitempty"`
}

func (x *MtasRequest) Reset() {
	*x = MtasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MtasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MtasRequest) ProtoMessage() {}

func (x *MtasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mtas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MtasRequest.ProtoReflect.Descriptor instead.
func (*MtasRequest) Descriptor() ([]byte, []int) {
	return file_mtas_proto_rawDescGZIP(), []int{0}
}

func (x *MtasRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *MtasRequest) GetOrganizationName() string {
	if x != nil {
		return x.OrganizationName
	}
	return ""
}

func (x *MtasRequest) GetSpaceName() string {
	if x != nil {
		return x.SpaceName
	}
	return ""
}

func (x *MtasRequest) GetApiUrl() string {
	if x != nil {
		return x.ApiUrl
	}
	return ""
}

type MtasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mtas []*MtaMetadata `protobuf:"bytes,1,rep,name=mtas,proto3" json:"mtas,omitempty"`
}

func (x *MtasResponse) Reset() {
	*x = MtasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtas_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MtasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MtasResponse) ProtoMessage() {}

func (x *MtasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mtas_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MtasResponse.ProtoReflect.Descriptor instead.
func (*MtasResponse) Descriptor() ([]byte, []int) {
	return file_mtas_proto_rawDescGZIP(), []int{1}
}

func (x *MtasResponse) GetMtas() []*MtaMetadata {
	if x != nil {
		return x.Mtas
	}
	return nil
}

type MtaMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *MtaMetadata) Reset() {
	*x = MtaMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtas_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MtaMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MtaMetadata) ProtoMessage() {}

func (x *MtaMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_mtas_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MtaMetadata.ProtoReflect.Descriptor instead.
func (*MtaMetadata) Descriptor() ([]byte, []int) {
	return file_mtas_proto_rawDescGZIP(), []int{2}
}

func (x *MtaMetadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MtaMetadata) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtas_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_mtas_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_mtas_proto_rawDescGZIP(), []int{3}
}

var File_mtas_proto protoreflect.FileDescriptor

var file_mtas_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x74, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x63, 0x6f,
	0x6d, 0x2e, 0x73, 0x61, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6d, 0x2e, 0x73,
	0x6c, 0x2e, 0x78, 0x73, 0x32, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x22, 0x85, 0x01, 0x0a, 0x0b,
	0x4d, 0x74, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x70, 0x69, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69,
	0x55, 0x72, 0x6c, 0x22, 0x4e, 0x0a, 0x0c, 0x4d, 0x74, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x04, 0x6d, 0x74, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6c, 0x6d, 0x2e, 0x73, 0x6c, 0x2e, 0x78, 0x73, 0x32, 0x2e, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x2e, 0x4d, 0x74, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d,
	0x74, 0x61, 0x73, 0x22, 0x37, 0x0a, 0x0b, 0x4d, 0x74, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xc6, 0x01, 0x0a, 0x04, 0x4d, 0x74, 0x61, 0x73, 0x12, 0x64,
	0x0a, 0x07, 0x67, 0x65, 0x74, 0x4d, 0x74, 0x61, 0x73, 0x12, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x73, 0x61, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6d, 0x2e, 0x73, 0x6c, 0x2e,
	0x78, 0x73, 0x32, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x4d, 0x74, 0x61, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x70, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6d, 0x2e, 0x73, 0x6c, 0x2e, 0x78, 0x73, 0x32, 0x2e,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x4d, 0x74, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x08, 0x73, 0x68, 0x75, 0x74, 0x44, 0x6f, 0x77, 0x6e,
	0x12, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6c, 0x6d, 0x2e, 0x73, 0x6c, 0x2e, 0x78, 0x73, 0x32, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x70,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6d, 0x2e, 0x73, 0x6c, 0x2e, 0x78, 0x73, 0x32,
	0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x75,
	0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6c, 0x6d, 0x2e, 0x73, 0x6c, 0x2e, 0x78, 0x73, 0x32, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x42,
	0x09, 0x4d, 0x74, 0x61, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x72, 0x79, 0x2d, 0x69, 0x6e, 0x63, 0x75, 0x62, 0x61, 0x74, 0x6f, 0x72, 0x2f,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x61, 0x70, 0x70, 0x73, 0x2d, 0x63, 0x6c, 0x69, 0x2d, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2,
	0x02, 0x03, 0x48, 0x4c, 0x57, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mtas_proto_rawDescOnce sync.Once
	file_mtas_proto_rawDescData = file_mtas_proto_rawDesc
)

func file_mtas_proto_rawDescGZIP() []byte {
	file_mtas_proto_rawDescOnce.Do(func() {
		file_mtas_proto_rawDescData = protoimpl.X.CompressGZIP(file_mtas_proto_rawDescData)
	})
	return file_mtas_proto_rawDescData
}

var file_mtas_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mtas_proto_goTypes = []interface{}{
	(*MtasRequest)(nil),  // 0: com.sap.cloud.lm.sl.xs2.local.MtasRequest
	(*MtasResponse)(nil), // 1: com.sap.cloud.lm.sl.xs2.local.MtasResponse
	(*MtaMetadata)(nil),  // 2: com.sap.cloud.lm.sl.xs2.local.MtaMetadata
	(*Empty)(nil),        // 3: com.sap.cloud.lm.sl.xs2.local.Empty
}
var file_mtas_proto_depIdxs = []int32{
	2, // 0: com.sap.cloud.lm.sl.xs2.local.MtasResponse.mtas:type_name -> com.sap.cloud.lm.sl.xs2.local.MtaMetadata
	0, // 1: com.sap.cloud.lm.sl.xs2.local.Mtas.getMtas:input_type -> com.sap.cloud.lm.sl.xs2.local.MtasRequest
	3, // 2: com.sap.cloud.lm.sl.xs2.local.Mtas.shutDown:input_type -> com.sap.cloud.lm.sl.xs2.local.Empty
	1, // 3: com.sap.cloud.lm.sl.xs2.local.Mtas.getMtas:output_type -> com.sap.cloud.lm.sl.xs2.local.MtasResponse
	3, // 4: com.sap.cloud.lm.sl.xs2.local.Mtas.shutDown:output_type -> com.sap.cloud.lm.sl.xs2.local.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mtas_proto_init() }
func file_mtas_proto_init() {
	if File_mtas_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mtas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MtasRequest); i {
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
		file_mtas_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MtasResponse); i {
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
		file_mtas_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MtaMetadata); i {
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
		file_mtas_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_mtas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mtas_proto_goTypes,
		DependencyIndexes: file_mtas_proto_depIdxs,
		MessageInfos:      file_mtas_proto_msgTypes,
	}.Build()
	File_mtas_proto = out.File
	file_mtas_proto_rawDesc = nil
	file_mtas_proto_goTypes = nil
	file_mtas_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MtasClient is the client API for Mtas service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MtasClient interface {
	GetMtas(ctx context.Context, in *MtasRequest, opts ...grpc.CallOption) (*MtasResponse, error)
	ShutDown(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type mtasClient struct {
	cc grpc.ClientConnInterface
}

func NewMtasClient(cc grpc.ClientConnInterface) MtasClient {
	return &mtasClient{cc}
}

func (c *mtasClient) GetMtas(ctx context.Context, in *MtasRequest, opts ...grpc.CallOption) (*MtasResponse, error) {
	out := new(MtasResponse)
	err := c.cc.Invoke(ctx, "/com.sap.cloud.lm.sl.xs2.local.Mtas/getMtas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mtasClient) ShutDown(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/com.sap.cloud.lm.sl.xs2.local.Mtas/shutDown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MtasServer is the server API for Mtas service.
type MtasServer interface {
	GetMtas(context.Context, *MtasRequest) (*MtasResponse, error)
	ShutDown(context.Context, *Empty) (*Empty, error)
}

// UnimplementedMtasServer can be embedded to have forward compatible implementations.
type UnimplementedMtasServer struct {
}

func (*UnimplementedMtasServer) GetMtas(context.Context, *MtasRequest) (*MtasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMtas not implemented")
}
func (*UnimplementedMtasServer) ShutDown(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShutDown not implemented")
}

func RegisterMtasServer(s *grpc.Server, srv MtasServer) {
	s.RegisterService(&_Mtas_serviceDesc, srv)
}

func _Mtas_GetMtas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MtasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MtasServer).GetMtas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.sap.cloud.lm.sl.xs2.local.Mtas/GetMtas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MtasServer).GetMtas(ctx, req.(*MtasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mtas_ShutDown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MtasServer).ShutDown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.sap.cloud.lm.sl.xs2.local.Mtas/ShutDown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MtasServer).ShutDown(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mtas_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.sap.cloud.lm.sl.xs2.local.Mtas",
	HandlerType: (*MtasServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getMtas",
			Handler:    _Mtas_GetMtas_Handler,
		},
		{
			MethodName: "shutDown",
			Handler:    _Mtas_ShutDown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mtas.proto",
}