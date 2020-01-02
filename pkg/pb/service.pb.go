// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/seizadi/test/pkg/pb/service.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "github.com/lyft/protoc-gen-validate/validate"
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
	return fileDescriptor_123714ce407bb634, []int{0}
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
	proto.RegisterFile("github.com/seizadi/test/pkg/pb/service.proto", fileDescriptor_123714ce407bb634)
}

var fileDescriptor_123714ce407bb634 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xe9, 0x9f, 0x3f, 0x56, 0xf7, 0xa2, 0xe4, 0x20, 0x25, 0x7a, 0xd0, 0x7a, 0x11, 0x6c,
	0x77, 0x40, 0x6f, 0xf6, 0x26, 0x88, 0x57, 0x29, 0x22, 0xe2, 0x6d, 0x37, 0x9d, 0xac, 0x8b, 0xc9,
	0xce, 0x92, 0x9d, 0x46, 0xeb, 0xd1, 0x57, 0xf0, 0xd1, 0x7c, 0x05, 0x1f, 0x44, 0x9a, 0x6c, 0x24,
	0xa8, 0x78, 0xdb, 0x6f, 0x66, 0xbe, 0xef, 0x07, 0xdf, 0x8a, 0x89, 0xb1, 0xfc, 0xb0, 0xd4, 0x32,
	0xa3, 0x12, 0x02, 0xda, 0x17, 0xb5, 0xb0, 0xc0, 0x18, 0x18, 0xfc, 0xa3, 0x01, 0xaf, 0x21, 0x60,
	0x55, 0xdb, 0x0c, 0xa5, 0xaf, 0x88, 0x29, 0x19, 0x46, 0x99, 0xee, 0x19, 0x22, 0x53, 0x20, 0x34,
	0x63, 0xbd, 0xcc, 0x01, 0x4b, 0xcf, 0xab, 0xf6, 0x2a, 0xdd, 0x8f, 0x4b, 0xe5, 0x2d, 0x28, 0xe7,
	0x88, 0x15, 0x5b, 0x72, 0x21, 0x6e, 0x67, 0x3d, 0x62, 0xb1, 0xca, 0xb9, 0xcd, 0xc8, 0xa6, 0x06,
	0xdd, 0xb4, 0x56, 0x85, 0x5d, 0x28, 0x46, 0xf8, 0xf1, 0x88, 0xe6, 0x49, 0xef, 0x38, 0x3c, 0x29,
	0x63, 0xb0, 0x02, 0xf2, 0x4d, 0xfc, 0x2f, 0xa8, 0xf3, 0x1e, 0xca, 0xba, 0x9c, 0x74, 0x41, 0xcf,
	0xe4, 0xd1, 0xf5, 0x91, 0x86, 0xaa, 0xf2, 0x2b, 0x62, 0x2d, 0x5a, 0xef, 0xf8, 0x44, 0x6c, 0xdf,
	0x62, 0x15, 0x2c, 0xb9, 0x39, 0x06, 0x4f, 0x2e, 0x60, 0x32, 0x12, 0xc3, 0xba, 0x1d, 0x8d, 0x06,
	0x07, 0x83, 0xe3, 0xad, 0x79, 0x27, 0x4f, 0xef, 0xc4, 0xff, 0x1b, 0x0c, 0x9c, 0x5c, 0x0b, 0x71,
	0x85, 0x1c, 0x7d, 0xc9, 0xae, 0x6c, 0x8b, 0x90, 0x5d, 0x4b, 0xf2, 0x72, 0xdd, 0x52, 0x3a, 0x92,
	0x5d, 0xab, 0xdf, 0x08, 0xe3, 0x9d, 0xd7, 0xf7, 0x8f, 0xb7, 0x7f, 0x22, 0xd9, 0x84, 0x98, 0x7c,
	0x71, 0x74, 0x7f, 0xf8, 0xf7, 0x0f, 0xcd, 0xbc, 0xd6, 0x1b, 0x0d, 0xe0, 0xec, 0x33, 0x00, 0x00,
	0xff, 0xff, 0x1e, 0x77, 0x4a, 0x06, 0xcd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestClient interface {
	GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type testClient struct {
	cc *grpc.ClientConn
}

func NewTestClient(cc *grpc.ClientConn) TestClient {
	return &testClient{cc}
}

func (c *testClient) GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/service.Test/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServer is the server API for Test service.
type TestServer interface {
	GetVersion(context.Context, *empty.Empty) (*VersionResponse, error)
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Test/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).GetVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Test_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/seizadi/test/pkg/pb/service.proto",
}