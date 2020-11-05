// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: cases.proto

package casesp

import (
	context "context"
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

type Case struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location     string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Age          int32  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Infectedtype string `protobuf:"bytes,5,opt,name=infectedtype,proto3" json:"infectedtype,omitempty"`
	State        string `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *Case) Reset() {
	*x = Case{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cases_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Case) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Case) ProtoMessage() {}

func (x *Case) ProtoReflect() protoreflect.Message {
	mi := &file_cases_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Case.ProtoReflect.Descriptor instead.
func (*Case) Descriptor() ([]byte, []int) {
	return file_cases_proto_rawDescGZIP(), []int{0}
}

func (x *Case) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Case) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Case) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Case) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Case) GetInfectedtype() string {
	if x != nil {
		return x.Infectedtype
	}
	return ""
}

func (x *Case) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type CaseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Case *Case `protobuf:"bytes,1,opt,name=case,proto3" json:"case,omitempty"`
}

func (x *CaseReq) Reset() {
	*x = CaseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cases_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaseReq) ProtoMessage() {}

func (x *CaseReq) ProtoReflect() protoreflect.Message {
	mi := &file_cases_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaseReq.ProtoReflect.Descriptor instead.
func (*CaseReq) Descriptor() ([]byte, []int) {
	return file_cases_proto_rawDescGZIP(), []int{1}
}

func (x *CaseReq) GetCase() *Case {
	if x != nil {
		return x.Case
	}
	return nil
}

type CaseRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Case *Case `protobuf:"bytes,1,opt,name=case,proto3" json:"case,omitempty"`
}

func (x *CaseRes) Reset() {
	*x = CaseRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cases_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaseRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaseRes) ProtoMessage() {}

func (x *CaseRes) ProtoReflect() protoreflect.Message {
	mi := &file_cases_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaseRes.ProtoReflect.Descriptor instead.
func (*CaseRes) Descriptor() ([]byte, []int) {
	return file_cases_proto_rawDescGZIP(), []int{2}
}

func (x *CaseRes) GetCase() *Case {
	if x != nil {
		return x.Case
	}
	return nil
}

var File_cases_proto protoreflect.FileDescriptor

var file_cases_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63,
	0x61, 0x73, 0x65, 0x73, 0x70, 0x22, 0x92, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x2b, 0x0a, 0x07, 0x43, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x04, 0x63, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x70, 0x2e, 0x43, 0x61, 0x73,
	0x65, 0x52, 0x04, 0x63, 0x61, 0x73, 0x65, 0x22, 0x2b, 0x0a, 0x07, 0x43, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x63, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x70, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x04,
	0x63, 0x61, 0x73, 0x65, 0x32, 0x3a, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x30,
	0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x0f, 0x2e, 0x63,
	0x61, 0x73, 0x65, 0x73, 0x70, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e,
	0x63, 0x61, 0x73, 0x65, 0x73, 0x70, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cases_proto_rawDescOnce sync.Once
	file_cases_proto_rawDescData = file_cases_proto_rawDesc
)

func file_cases_proto_rawDescGZIP() []byte {
	file_cases_proto_rawDescOnce.Do(func() {
		file_cases_proto_rawDescData = protoimpl.X.CompressGZIP(file_cases_proto_rawDescData)
	})
	return file_cases_proto_rawDescData
}

var file_cases_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cases_proto_goTypes = []interface{}{
	(*Case)(nil),    // 0: casesp.Case
	(*CaseReq)(nil), // 1: casesp.CaseReq
	(*CaseRes)(nil), // 2: casesp.CaseRes
}
var file_cases_proto_depIdxs = []int32{
	0, // 0: casesp.CaseReq.case:type_name -> casesp.Case
	0, // 1: casesp.CaseRes.case:type_name -> casesp.Case
	1, // 2: casesp.Insert.InsertCase:input_type -> casesp.CaseReq
	2, // 3: casesp.Insert.InsertCase:output_type -> casesp.CaseRes
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cases_proto_init() }
func file_cases_proto_init() {
	if File_cases_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cases_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Case); i {
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
		file_cases_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaseReq); i {
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
		file_cases_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaseRes); i {
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
			RawDescriptor: file_cases_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cases_proto_goTypes,
		DependencyIndexes: file_cases_proto_depIdxs,
		MessageInfos:      file_cases_proto_msgTypes,
	}.Build()
	File_cases_proto = out.File
	file_cases_proto_rawDesc = nil
	file_cases_proto_goTypes = nil
	file_cases_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InsertClient is the client API for Insert service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InsertClient interface {
	InsertCase(ctx context.Context, in *CaseReq, opts ...grpc.CallOption) (*CaseRes, error)
}

type insertClient struct {
	cc grpc.ClientConnInterface
}

func NewInsertClient(cc grpc.ClientConnInterface) InsertClient {
	return &insertClient{cc}
}

func (c *insertClient) InsertCase(ctx context.Context, in *CaseReq, opts ...grpc.CallOption) (*CaseRes, error) {
	out := new(CaseRes)
	err := c.cc.Invoke(ctx, "/casesp.Insert/InsertCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InsertServer is the server API for Insert service.
type InsertServer interface {
	InsertCase(context.Context, *CaseReq) (*CaseRes, error)
}

// UnimplementedInsertServer can be embedded to have forward compatible implementations.
type UnimplementedInsertServer struct {
}

func (*UnimplementedInsertServer) InsertCase(context.Context, *CaseReq) (*CaseRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertCase not implemented")
}

func RegisterInsertServer(s *grpc.Server, srv InsertServer) {
	s.RegisterService(&_Insert_serviceDesc, srv)
}

func _Insert_InsertCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsertServer).InsertCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/casesp.Insert/InsertCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsertServer).InsertCase(ctx, req.(*CaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Insert_serviceDesc = grpc.ServiceDesc{
	ServiceName: "casesp.Insert",
	HandlerType: (*InsertServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertCase",
			Handler:    _Insert_InsertCase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cases.proto",
}
