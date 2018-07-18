// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fim.proto

package fim

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type FimdConfig struct {
	HostUid              string               `protobuf:"bytes,1,opt,name=hostUid,proto3" json:"hostUid,omitempty"`
	ContainerId          []string             `protobuf:"bytes,2,rep,name=containerId,proto3" json:"containerId,omitempty"`
	Subject              []*FimWatcherSubject `protobuf:"bytes,3,rep,name=subject,proto3" json:"subject,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FimdConfig) Reset()         { *m = FimdConfig{} }
func (m *FimdConfig) String() string { return proto.CompactTextString(m) }
func (*FimdConfig) ProtoMessage()    {}
func (*FimdConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_fim_d7f44d3e17185854, []int{0}
}
func (m *FimdConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FimdConfig.Unmarshal(m, b)
}
func (m *FimdConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FimdConfig.Marshal(b, m, deterministic)
}
func (dst *FimdConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FimdConfig.Merge(dst, src)
}
func (m *FimdConfig) XXX_Size() int {
	return xxx_messageInfo_FimdConfig.Size(m)
}
func (m *FimdConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FimdConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FimdConfig proto.InternalMessageInfo

func (m *FimdConfig) GetHostUid() string {
	if m != nil {
		return m.HostUid
	}
	return ""
}

func (m *FimdConfig) GetContainerId() []string {
	if m != nil {
		return m.ContainerId
	}
	return nil
}

func (m *FimdConfig) GetSubject() []*FimWatcherSubject {
	if m != nil {
		return m.Subject
	}
	return nil
}

type FimWatcherSubject struct {
	Path                 []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
	Event                []string `protobuf:"bytes,2,rep,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FimWatcherSubject) Reset()         { *m = FimWatcherSubject{} }
func (m *FimWatcherSubject) String() string { return proto.CompactTextString(m) }
func (*FimWatcherSubject) ProtoMessage()    {}
func (*FimWatcherSubject) Descriptor() ([]byte, []int) {
	return fileDescriptor_fim_d7f44d3e17185854, []int{1}
}
func (m *FimWatcherSubject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FimWatcherSubject.Unmarshal(m, b)
}
func (m *FimWatcherSubject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FimWatcherSubject.Marshal(b, m, deterministic)
}
func (dst *FimWatcherSubject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FimWatcherSubject.Merge(dst, src)
}
func (m *FimWatcherSubject) XXX_Size() int {
	return xxx_messageInfo_FimWatcherSubject.Size(m)
}
func (m *FimWatcherSubject) XXX_DiscardUnknown() {
	xxx_messageInfo_FimWatcherSubject.DiscardUnknown(m)
}

var xxx_messageInfo_FimWatcherSubject proto.InternalMessageInfo

func (m *FimWatcherSubject) GetPath() []string {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *FimWatcherSubject) GetEvent() []string {
	if m != nil {
		return m.Event
	}
	return nil
}

type FimdHandle struct {
	HostUid              string   `protobuf:"bytes,1,opt,name=hostUid,proto3" json:"hostUid,omitempty"`
	Pid                  []int32  `protobuf:"varint,2,rep,packed,name=pid,proto3" json:"pid,omitempty"`
	ProcessEventfd       []int32  `protobuf:"varint,3,rep,packed,name=processEventfd,proto3" json:"processEventfd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FimdHandle) Reset()         { *m = FimdHandle{} }
func (m *FimdHandle) String() string { return proto.CompactTextString(m) }
func (*FimdHandle) ProtoMessage()    {}
func (*FimdHandle) Descriptor() ([]byte, []int) {
	return fileDescriptor_fim_d7f44d3e17185854, []int{2}
}
func (m *FimdHandle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FimdHandle.Unmarshal(m, b)
}
func (m *FimdHandle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FimdHandle.Marshal(b, m, deterministic)
}
func (dst *FimdHandle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FimdHandle.Merge(dst, src)
}
func (m *FimdHandle) XXX_Size() int {
	return xxx_messageInfo_FimdHandle.Size(m)
}
func (m *FimdHandle) XXX_DiscardUnknown() {
	xxx_messageInfo_FimdHandle.DiscardUnknown(m)
}

var xxx_messageInfo_FimdHandle proto.InternalMessageInfo

func (m *FimdHandle) GetHostUid() string {
	if m != nil {
		return m.HostUid
	}
	return ""
}

func (m *FimdHandle) GetPid() []int32 {
	if m != nil {
		return m.Pid
	}
	return nil
}

func (m *FimdHandle) GetProcessEventfd() []int32 {
	if m != nil {
		return m.ProcessEventfd
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_fim_d7f44d3e17185854, []int{3}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FimdConfig)(nil), "fim.FimdConfig")
	proto.RegisterType((*FimWatcherSubject)(nil), "fim.FimWatcherSubject")
	proto.RegisterType((*FimdHandle)(nil), "fim.FimdHandle")
	proto.RegisterType((*Empty)(nil), "fim.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FimdClient is the client API for Fimd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FimdClient interface {
	CreateWatch(ctx context.Context, in *FimdConfig, opts ...grpc.CallOption) (*FimdHandle, error)
	DestroyWatch(ctx context.Context, in *FimdConfig, opts ...grpc.CallOption) (*Empty, error)
}

type fimdClient struct {
	cc *grpc.ClientConn
}

func NewFimdClient(cc *grpc.ClientConn) FimdClient {
	return &fimdClient{cc}
}

func (c *fimdClient) CreateWatch(ctx context.Context, in *FimdConfig, opts ...grpc.CallOption) (*FimdHandle, error) {
	out := new(FimdHandle)
	err := c.cc.Invoke(ctx, "/fim.Fimd/CreateWatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fimdClient) DestroyWatch(ctx context.Context, in *FimdConfig, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/fim.Fimd/DestroyWatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FimdServer is the server API for Fimd service.
type FimdServer interface {
	CreateWatch(context.Context, *FimdConfig) (*FimdHandle, error)
	DestroyWatch(context.Context, *FimdConfig) (*Empty, error)
}

func RegisterFimdServer(s *grpc.Server, srv FimdServer) {
	s.RegisterService(&_Fimd_serviceDesc, srv)
}

func _Fimd_CreateWatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FimdConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FimdServer).CreateWatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fim.Fimd/CreateWatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FimdServer).CreateWatch(ctx, req.(*FimdConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fimd_DestroyWatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FimdConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FimdServer).DestroyWatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fim.Fimd/DestroyWatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FimdServer).DestroyWatch(ctx, req.(*FimdConfig))
	}
	return interceptor(ctx, in, info, handler)
}

var _Fimd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fim.Fimd",
	HandlerType: (*FimdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWatch",
			Handler:    _Fimd_CreateWatch_Handler,
		},
		{
			MethodName: "DestroyWatch",
			Handler:    _Fimd_DestroyWatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fim.proto",
}

func init() { proto.RegisterFile("fim.proto", fileDescriptor_fim_d7f44d3e17185854) }

var fileDescriptor_fim_d7f44d3e17185854 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4f, 0x84, 0x30,
	0x10, 0x47, 0x17, 0x59, 0x24, 0x0c, 0xc6, 0x3f, 0x13, 0x63, 0x88, 0x27, 0xd2, 0x83, 0xe1, 0xe2,
	0x46, 0xd7, 0xb3, 0xa7, 0x75, 0x37, 0x7a, 0xc5, 0x18, 0xaf, 0xb2, 0x30, 0x48, 0x8d, 0x50, 0xd2,
	0x8e, 0x26, 0xeb, 0xa7, 0x37, 0xb4, 0x4b, 0x34, 0x1a, 0xbd, 0xcd, 0xfc, 0xa6, 0xed, 0x7b, 0x6d,
	0x21, 0xaa, 0x65, 0x3b, 0xeb, 0xb5, 0x62, 0x85, 0x7e, 0x2d, 0x5b, 0xf1, 0x01, 0xb0, 0x92, 0x6d,
	0xb5, 0x50, 0x5d, 0x2d, 0x9f, 0x31, 0x81, 0xb0, 0x51, 0x86, 0x1f, 0x64, 0x95, 0x78, 0xa9, 0x97,
	0x45, 0xf9, 0xd8, 0x62, 0x0a, 0x71, 0xa9, 0x3a, 0x2e, 0x64, 0x47, 0xfa, 0xae, 0x4a, 0x76, 0x52,
	0x3f, 0x8b, 0xf2, 0xef, 0x11, 0x5e, 0x40, 0x68, 0xde, 0xd6, 0x2f, 0x54, 0x72, 0xe2, 0xa7, 0x7e,
	0x16, 0xcf, 0x4f, 0x66, 0x03, 0x6b, 0x25, 0xdb, 0xc7, 0x82, 0xcb, 0x86, 0xf4, 0xbd, 0x9b, 0xe6,
	0xe3, 0x32, 0x71, 0x0d, 0x47, 0xbf, 0xa6, 0x88, 0x30, 0xed, 0x0b, 0x6e, 0x12, 0xcf, 0x12, 0x6c,
	0x8d, 0xc7, 0x10, 0xd0, 0x3b, 0x75, 0xbc, 0xc5, 0xba, 0x46, 0x3c, 0x39, 0xf5, 0xdb, 0xa2, 0xab,
	0x5e, 0xe9, 0x1f, 0xf5, 0x43, 0xf0, 0x7b, 0xe9, 0x94, 0x83, 0x7c, 0x28, 0xf1, 0x0c, 0xf6, 0x7b,
	0xad, 0x4a, 0x32, 0x66, 0x39, 0x9c, 0x54, 0x57, 0xd6, 0x38, 0xc8, 0x7f, 0xa4, 0x22, 0x84, 0x60,
	0xd9, 0xf6, 0xbc, 0x99, 0x37, 0x30, 0x1d, 0x50, 0x78, 0x09, 0xf1, 0x42, 0x53, 0xc1, 0x64, 0xa5,
	0xf1, 0x60, 0xbc, 0xe1, 0xf6, 0xfd, 0x4e, 0xbf, 0x02, 0x67, 0x25, 0x26, 0x78, 0x0e, 0x7b, 0x37,
	0x64, 0x58, 0xab, 0xcd, 0x1f, 0x7b, 0xc0, 0x06, 0x96, 0x23, 0x26, 0xeb, 0x5d, 0xfb, 0x37, 0x57,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x5f, 0xfe, 0x08, 0xa8, 0x01, 0x00, 0x00,
}