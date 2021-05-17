// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/brocwoodworthIBLX/hack-db/pkg/pb/service.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// TODO: Structure your own protobuf messages. Each protocol buffer message is a
// small logical record of information, containing a series of name-value pairs.
type VersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73f7111cf4d69d9, []int{0}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*VersionResponse)(nil), "service.VersionResponse")
}

func init() {
	proto.RegisterFile("github.com/brocwoodworthIBLX/hack-db/pkg/pb/service.proto", fileDescriptor_f73f7111cf4d69d9)
}

var fileDescriptor_f73f7111cf4d69d9 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x50, 0xbb, 0x4e, 0x03, 0x31,
	0x10, 0x54, 0x28, 0x12, 0x70, 0x03, 0xba, 0x02, 0x45, 0x07, 0x05, 0x4a, 0x85, 0x04, 0xe7, 0x45,
	0x50, 0x01, 0x5d, 0x04, 0x02, 0x24, 0x0a, 0x94, 0x02, 0xa1, 0x74, 0xb6, 0xb3, 0xf1, 0x59, 0x97,
	0xf3, 0x5a, 0xb6, 0x93, 0x83, 0x96, 0x5f, 0xe0, 0xd3, 0xf8, 0x05, 0x3e, 0x04, 0xe5, 0x1e, 0x28,
	0x42, 0x34, 0x54, 0xd6, 0xec, 0xce, 0xce, 0x78, 0x86, 0x5d, 0x6a, 0x13, 0xf3, 0xa5, 0xe4, 0x8a,
	0x4a, 0x90, 0x9e, 0x54, 0x45, 0x34, 0xab, 0xc8, 0xc7, 0xfc, 0x61, 0xfc, 0xf8, 0x02, 0xb9, 0x50,
	0x45, 0x36, 0x93, 0xe0, 0x0a, 0x0d, 0x4e, 0x42, 0x40, 0xbf, 0x32, 0x0a, 0xb9, 0xf3, 0x14, 0x29,
	0x19, 0xb4, 0x30, 0x3d, 0xd0, 0x44, 0x7a, 0x81, 0x50, 0x8f, 0xe5, 0x72, 0x0e, 0x58, 0xba, 0xf8,
	0xd6, 0xb0, 0xd2, 0xc3, 0x76, 0x29, 0x9c, 0x01, 0x61, 0x2d, 0x45, 0x11, 0x0d, 0xd9, 0xd0, 0x6e,
	0x4f, 0xeb, 0x47, 0x65, 0x1a, 0x6d, 0x16, 0x2a, 0xa1, 0x35, 0x7a, 0x20, 0x57, 0x33, 0xfe, 0x60,
	0x5f, 0x6d, 0x7c, 0xd6, 0xd8, 0x39, 0xc9, 0x05, 0xbd, 0x92, 0x43, 0x0b, 0x1b, 0x2a, 0x9a, 0x7c,
	0xf9, 0x23, 0xb1, 0x06, 0xcd, 0xed, 0xe8, 0x84, 0xed, 0x3e, 0xa3, 0x0f, 0x86, 0xec, 0x04, 0x83,
	0x23, 0x1b, 0x30, 0x19, 0xb2, 0xc1, 0xaa, 0x19, 0x0d, 0x7b, 0x47, 0xbd, 0xe3, 0x9d, 0x49, 0x07,
	0xcf, 0xa7, 0xac, 0x7f, 0x2f, 0x54, 0x71, 0x23, 0x93, 0x27, 0xc6, 0xee, 0x30, 0xb6, 0x97, 0xc9,
	0x3e, 0x6f, 0xd2, 0xf0, 0x2e, 0x2a, 0xbf, 0x5d, 0x47, 0x4d, 0x87, 0xbc, 0xab, 0xe6, 0x97, 0xc7,
	0x68, 0xef, 0xfd, 0xf3, 0xeb, 0x63, 0x8b, 0x25, 0xdb, 0xd0, 0x6a, 0x8f, 0xcf, 0xa6, 0xfc, 0x1f,
	0x9d, 0x5f, 0x3b, 0x29, 0xfb, 0xb5, 0xdb, 0xc5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x0b,
	0x74, 0xfd, 0xac, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HackDbClient is the client API for HackDb service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HackDbClient interface {
	GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type hackDbClient struct {
	cc *grpc.ClientConn
}

func NewHackDbClient(cc *grpc.ClientConn) HackDbClient {
	return &hackDbClient{cc}
}

func (c *hackDbClient) GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/service.HackDb/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HackDbServer is the server API for HackDb service.
type HackDbServer interface {
	GetVersion(context.Context, *empty.Empty) (*VersionResponse, error)
}

func RegisterHackDbServer(s *grpc.Server, srv HackDbServer) {
	s.RegisterService(&_HackDb_serviceDesc, srv)
}

func _HackDb_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HackDbServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.HackDb/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HackDbServer).GetVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _HackDb_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.HackDb",
	HandlerType: (*HackDbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _HackDb_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/brocwoodworthIBLX/hack-db/pkg/pb/service.proto",
}
