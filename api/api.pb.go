// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package clusteror is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	ClusterDefinition
	ClusterConfigs
	ClusterStatusMsg
*/
package clusteror

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClusterDefinition struct {
	ClusterProvider          string          `protobuf:"bytes,1,opt,name=clusterProvider" json:"clusterProvider,omitempty"`
	ClusterConfigs           *ClusterConfigs `protobuf:"bytes,2,opt,name=clusterConfigs" json:"clusterConfigs,omitempty"`
	AutoFetchClusterProvider bool            `protobuf:"varint,3,opt,name=autoFetchClusterProvider" json:"autoFetchClusterProvider,omitempty"`
	ProviderStorePath        string          `protobuf:"bytes,4,opt,name=providerStorePath" json:"providerStorePath,omitempty"`
	CloudID                  string          `protobuf:"bytes,5,opt,name=CloudID" json:"CloudID,omitempty"`
}

func (m *ClusterDefinition) Reset()                    { *m = ClusterDefinition{} }
func (m *ClusterDefinition) String() string            { return proto.CompactTextString(m) }
func (*ClusterDefinition) ProtoMessage()               {}
func (*ClusterDefinition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClusterDefinition) GetClusterProvider() string {
	if m != nil {
		return m.ClusterProvider
	}
	return ""
}

func (m *ClusterDefinition) GetClusterConfigs() *ClusterConfigs {
	if m != nil {
		return m.ClusterConfigs
	}
	return nil
}

func (m *ClusterDefinition) GetAutoFetchClusterProvider() bool {
	if m != nil {
		return m.AutoFetchClusterProvider
	}
	return false
}

func (m *ClusterDefinition) GetProviderStorePath() string {
	if m != nil {
		return m.ProviderStorePath
	}
	return ""
}

func (m *ClusterDefinition) GetCloudID() string {
	if m != nil {
		return m.CloudID
	}
	return ""
}

type ClusterConfigs struct {
	Name              string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	CloudProviderName string `protobuf:"bytes,2,opt,name=cloudProviderName" json:"cloudProviderName,omitempty"`
}

func (m *ClusterConfigs) Reset()                    { *m = ClusterConfigs{} }
func (m *ClusterConfigs) String() string            { return proto.CompactTextString(m) }
func (*ClusterConfigs) ProtoMessage()               {}
func (*ClusterConfigs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClusterConfigs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterConfigs) GetCloudProviderName() string {
	if m != nil {
		return m.CloudProviderName
	}
	return ""
}

type ClusterStatusMsg struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Code   int64  `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
}

func (m *ClusterStatusMsg) Reset()                    { *m = ClusterStatusMsg{} }
func (m *ClusterStatusMsg) String() string            { return proto.CompactTextString(m) }
func (*ClusterStatusMsg) ProtoMessage()               {}
func (*ClusterStatusMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClusterStatusMsg) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ClusterStatusMsg) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*ClusterDefinition)(nil), "clusteror.ClusterDefinition")
	proto.RegisterType((*ClusterConfigs)(nil), "clusteror.ClusterConfigs")
	proto.RegisterType((*ClusterStatusMsg)(nil), "clusteror.ClusterStatusMsg")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ClusterCreator service

type ClusterCreatorClient interface {
	Create(ctx context.Context, in *ClusterDefinition, opts ...grpc.CallOption) (*ClusterStatusMsg, error)
	Apply(ctx context.Context, in *ClusterDefinition, opts ...grpc.CallOption) (*ClusterStatusMsg, error)
	Delete(ctx context.Context, in *ClusterDefinition, opts ...grpc.CallOption) (*ClusterStatusMsg, error)
}

type clusterCreatorClient struct {
	cc *grpc.ClientConn
}

func NewClusterCreatorClient(cc *grpc.ClientConn) ClusterCreatorClient {
	return &clusterCreatorClient{cc}
}

func (c *clusterCreatorClient) Create(ctx context.Context, in *ClusterDefinition, opts ...grpc.CallOption) (*ClusterStatusMsg, error) {
	out := new(ClusterStatusMsg)
	err := grpc.Invoke(ctx, "/clusteror.ClusterCreator/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterCreatorClient) Apply(ctx context.Context, in *ClusterDefinition, opts ...grpc.CallOption) (*ClusterStatusMsg, error) {
	out := new(ClusterStatusMsg)
	err := grpc.Invoke(ctx, "/clusteror.ClusterCreator/Apply", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterCreatorClient) Delete(ctx context.Context, in *ClusterDefinition, opts ...grpc.CallOption) (*ClusterStatusMsg, error) {
	out := new(ClusterStatusMsg)
	err := grpc.Invoke(ctx, "/clusteror.ClusterCreator/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClusterCreator service

type ClusterCreatorServer interface {
	Create(context.Context, *ClusterDefinition) (*ClusterStatusMsg, error)
	Apply(context.Context, *ClusterDefinition) (*ClusterStatusMsg, error)
	Delete(context.Context, *ClusterDefinition) (*ClusterStatusMsg, error)
}

func RegisterClusterCreatorServer(s *grpc.Server, srv ClusterCreatorServer) {
	s.RegisterService(&_ClusterCreator_serviceDesc, srv)
}

func _ClusterCreator_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterDefinition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterCreatorServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clusteror.ClusterCreator/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterCreatorServer).Create(ctx, req.(*ClusterDefinition))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterCreator_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterDefinition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterCreatorServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clusteror.ClusterCreator/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterCreatorServer).Apply(ctx, req.(*ClusterDefinition))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterCreator_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterDefinition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterCreatorServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clusteror.ClusterCreator/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterCreatorServer).Delete(ctx, req.(*ClusterDefinition))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClusterCreator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clusteror.ClusterCreator",
	HandlerType: (*ClusterCreatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ClusterCreator_Create_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _ClusterCreator_Apply_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ClusterCreator_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x41, 0x6e, 0xda, 0x40,
	0x14, 0x86, 0x65, 0x03, 0x6e, 0xfd, 0x2a, 0x41, 0x19, 0xb5, 0x95, 0x4b, 0x59, 0x20, 0xaf, 0x2c,
	0x04, 0x58, 0xa5, 0x3b, 0x16, 0x95, 0x90, 0x51, 0xa5, 0x2e, 0x12, 0x21, 0xb3, 0x89, 0xd8, 0x4d,
	0xec, 0xc1, 0x58, 0x72, 0xfc, 0xac, 0xf1, 0x98, 0x28, 0xdb, 0x28, 0x37, 0xc8, 0x0d, 0x72, 0xa5,
	0x5c, 0x21, 0x07, 0x89, 0x66, 0x6c, 0x40, 0x60, 0x65, 0x95, 0xac, 0x78, 0x6f, 0xfe, 0x7f, 0xbe,
	0xff, 0xc7, 0x36, 0x98, 0x34, 0x8b, 0x27, 0x19, 0x47, 0x81, 0xc4, 0x0c, 0x92, 0x22, 0x17, 0x8c,
	0x23, 0xef, 0xf5, 0x23, 0xc4, 0x28, 0x61, 0x2e, 0xcd, 0x62, 0x97, 0xa6, 0x29, 0x0a, 0x2a, 0x62,
	0x4c, 0xf3, 0xd2, 0xd8, 0x1b, 0xa9, 0x9f, 0x60, 0x1c, 0xb1, 0x74, 0x9c, 0xdf, 0xd2, 0x28, 0x62,
	0xdc, 0xc5, 0x4c, 0x39, 0xea, 0x6e, 0xfb, 0x41, 0x87, 0xae, 0x57, 0x92, 0x17, 0x6c, 0x13, 0xa7,
	0xb1, 0x14, 0x89, 0x03, 0x9d, 0x2a, 0x6e, 0xc9, 0x71, 0x17, 0x87, 0x8c, 0x5b, 0xda, 0x40, 0x73,
	0x4c, 0xff, 0xfc, 0x98, 0xcc, 0xa1, 0x5d, 0x1d, 0x79, 0x98, 0x6e, 0xe2, 0x28, 0xb7, 0xf4, 0x81,
	0xe6, 0x7c, 0x99, 0xfe, 0x9c, 0x1c, 0xfa, 0x4e, 0xbc, 0x13, 0x83, 0x7f, 0x76, 0x81, 0xcc, 0xc0,
	0xa2, 0x85, 0xc0, 0x7f, 0x4c, 0x04, 0x5b, 0xef, 0x2c, 0xb5, 0x31, 0xd0, 0x9c, 0xcf, 0xfe, 0x9b,
	0x3a, 0x19, 0x41, 0x37, 0xab, 0xe6, 0x95, 0x40, 0xce, 0x96, 0x54, 0x6c, 0xad, 0xa6, 0xaa, 0x5a,
	0x17, 0x88, 0x05, 0x9f, 0xbc, 0x04, 0x8b, 0xf0, 0xff, 0xc2, 0x6a, 0x29, 0xcf, 0x7e, 0xb5, 0x7d,
	0x68, 0x9f, 0xb6, 0x24, 0x04, 0x9a, 0x29, 0xbd, 0x61, 0xd5, 0xff, 0x56, 0xb3, 0x4c, 0x0b, 0xe4,
	0x85, 0x7d, 0xfc, 0xa5, 0x34, 0xe8, 0x65, 0x5a, 0x4d, 0xb0, 0xff, 0xc2, 0xd7, 0x8a, 0xb9, 0x12,
	0x54, 0x14, 0xf9, 0x45, 0x1e, 0x91, 0x1f, 0x60, 0xe4, 0x6a, 0xa9, 0xb8, 0xd5, 0x26, 0xd3, 0x02,
	0x0c, 0x4b, 0x58, 0xc3, 0x57, 0xf3, 0xf4, 0x49, 0x3f, 0x96, 0xe2, 0x8c, 0x0a, 0xe4, 0x64, 0x0d,
	0x86, 0x1a, 0x19, 0xe9, 0xd7, 0x9f, 0xef, 0xf1, 0xfd, 0xf5, 0x7e, 0xd5, 0xd5, 0x43, 0x07, 0xfb,
	0xfb, 0xfd, 0xf3, 0xcb, 0xa3, 0xde, 0xb1, 0xc1, 0xdd, 0xfd, 0x76, 0x03, 0x85, 0x9b, 0x69, 0x43,
	0x72, 0x05, 0xad, 0x79, 0x96, 0x25, 0x77, 0xef, 0x41, 0x7f, 0x53, 0xe8, 0xb6, 0x6d, 0x4a, 0x34,
	0x95, 0x34, 0x49, 0x5e, 0x83, 0xb1, 0x60, 0x09, 0xfb, 0xc0, 0xd6, 0xa1, 0xc2, 0xcd, 0xb4, 0xe1,
	0xb5, 0xa1, 0x3e, 0xe3, 0x3f, 0xaf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x87, 0x70, 0xa5, 0x2a,
	0x03, 0x00, 0x00,
}
