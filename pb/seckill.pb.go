// Code generated by protoc-gen-go. DO NOT EDIT.
// source: seckill.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type SecRequest struct {
	ProductId            int64    `protobuf:"varint,1,opt,name=ProductId,json=productId,proto3" json:"ProductId,omitempty"`
	Source               string   `protobuf:"bytes,2,opt,name=Source,json=source,proto3" json:"Source,omitempty"`
	AuthCode             string   `protobuf:"bytes,3,opt,name=AuthCode,json=authCode,proto3" json:"AuthCode,omitempty"`
	SecTime              string   `protobuf:"bytes,4,opt,name=SecTime,json=secTime,proto3" json:"SecTime,omitempty"`
	Nance                string   `protobuf:"bytes,5,opt,name=Nance,json=nance,proto3" json:"Nance,omitempty"`
	UserId               int64    `protobuf:"varint,6,opt,name=UserId,json=userId,proto3" json:"UserId,omitempty"`
	UserAuthSign         string   `protobuf:"bytes,7,opt,name=UserAuthSign,json=userAuthSign,proto3" json:"UserAuthSign,omitempty"`
	AccessTime           int64    `protobuf:"varint,8,opt,name=AccessTime,json=accessTime,proto3" json:"AccessTime,omitempty"`
	ClientAddr           string   `protobuf:"bytes,9,opt,name=ClientAddr,json=clientAddr,proto3" json:"ClientAddr,omitempty"`
	ClientRefence        string   `protobuf:"bytes,10,opt,name=ClientRefence,json=clientRefence,proto3" json:"ClientRefence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecRequest) Reset()         { *m = SecRequest{} }
func (m *SecRequest) String() string { return proto.CompactTextString(m) }
func (*SecRequest) ProtoMessage()    {}
func (*SecRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1202afb06d2a3a7a, []int{0}
}

func (m *SecRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecRequest.Unmarshal(m, b)
}
func (m *SecRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecRequest.Marshal(b, m, deterministic)
}
func (m *SecRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecRequest.Merge(m, src)
}
func (m *SecRequest) XXX_Size() int {
	return xxx_messageInfo_SecRequest.Size(m)
}
func (m *SecRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SecRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SecRequest proto.InternalMessageInfo

func (m *SecRequest) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *SecRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *SecRequest) GetAuthCode() string {
	if m != nil {
		return m.AuthCode
	}
	return ""
}

func (m *SecRequest) GetSecTime() string {
	if m != nil {
		return m.SecTime
	}
	return ""
}

func (m *SecRequest) GetNance() string {
	if m != nil {
		return m.Nance
	}
	return ""
}

func (m *SecRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SecRequest) GetUserAuthSign() string {
	if m != nil {
		return m.UserAuthSign
	}
	return ""
}

func (m *SecRequest) GetAccessTime() int64 {
	if m != nil {
		return m.AccessTime
	}
	return 0
}

func (m *SecRequest) GetClientAddr() string {
	if m != nil {
		return m.ClientAddr
	}
	return ""
}

func (m *SecRequest) GetClientRefence() string {
	if m != nil {
		return m.ClientRefence
	}
	return ""
}

type SecResponse struct {
	ProductId            int64    `protobuf:"varint,1,opt,name=ProductId,json=productId,proto3" json:"ProductId,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=UserId,json=userId,proto3" json:"UserId,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=Token,json=token,proto3" json:"Token,omitempty"`
	TokenTime            int64    `protobuf:"varint,4,opt,name=TokenTime,json=tokenTime,proto3" json:"TokenTime,omitempty"`
	Code                 int64    `protobuf:"varint,5,opt,name=Code,json=code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecResponse) Reset()         { *m = SecResponse{} }
func (m *SecResponse) String() string { return proto.CompactTextString(m) }
func (*SecResponse) ProtoMessage()    {}
func (*SecResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1202afb06d2a3a7a, []int{1}
}

func (m *SecResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecResponse.Unmarshal(m, b)
}
func (m *SecResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecResponse.Marshal(b, m, deterministic)
}
func (m *SecResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecResponse.Merge(m, src)
}
func (m *SecResponse) XXX_Size() int {
	return xxx_messageInfo_SecResponse.Size(m)
}
func (m *SecResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SecResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SecResponse proto.InternalMessageInfo

func (m *SecResponse) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *SecResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SecResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SecResponse) GetTokenTime() int64 {
	if m != nil {
		return m.TokenTime
	}
	return 0
}

func (m *SecResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*SecRequest)(nil), "pb.SecRequest")
	proto.RegisterType((*SecResponse)(nil), "pb.SecResponse")
}

func init() { proto.RegisterFile("seckill.proto", fileDescriptor_1202afb06d2a3a7a) }

var fileDescriptor_1202afb06d2a3a7a = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4e, 0xfa, 0x40,
	0x10, 0xc6, 0xff, 0x2d, 0xb4, 0xd0, 0xf9, 0x03, 0x26, 0x13, 0x62, 0x36, 0xc4, 0x18, 0xd2, 0x78,
	0xe0, 0x60, 0x38, 0xe8, 0xdd, 0x84, 0x70, 0x22, 0x26, 0xc6, 0xb4, 0xf8, 0x00, 0x30, 0x1d, 0xb5,
	0xa1, 0x76, 0x6b, 0x77, 0xeb, 0x4b, 0xf8, 0x32, 0x3e, 0xa2, 0xd9, 0x01, 0x2a, 0x9e, 0x3c, 0x7e,
	0xbf, 0xaf, 0xbb, 0xdf, 0xcc, 0xd7, 0x85, 0xa1, 0x61, 0xda, 0xe5, 0x45, 0x31, 0xaf, 0x6a, 0x6d,
	0x35, 0xfa, 0xd5, 0x36, 0xfe, 0xf2, 0x01, 0x52, 0xa6, 0x84, 0xdf, 0x1b, 0x36, 0x16, 0x2f, 0x20,
	0x7a, 0xac, 0x75, 0xd6, 0x90, 0x5d, 0x65, 0xca, 0x9b, 0x7a, 0xb3, 0x4e, 0x12, 0x55, 0x47, 0x80,
	0xe7, 0x10, 0xa6, 0xba, 0xa9, 0x89, 0x95, 0x3f, 0xf5, 0x66, 0x51, 0x12, 0x1a, 0x51, 0x38, 0x81,
	0xfe, 0xa2, 0xb1, 0xaf, 0x4b, 0x9d, 0xb1, 0xea, 0x88, 0xd3, 0xdf, 0x1c, 0x34, 0x2a, 0xe8, 0xa5,
	0x4c, 0xeb, 0xfc, 0x8d, 0x55, 0x57, 0xac, 0x9e, 0xd9, 0x4b, 0x1c, 0x43, 0xf0, 0xb0, 0x29, 0x89,
	0x55, 0x20, 0x3c, 0x28, 0x9d, 0x70, 0x19, 0x4f, 0x86, 0xeb, 0x55, 0xa6, 0x42, 0x89, 0x0f, 0x1b,
	0x51, 0x18, 0xc3, 0xc0, 0x71, 0x97, 0x93, 0xe6, 0x2f, 0xa5, 0xea, 0xc9, 0xa1, 0x41, 0x73, 0xc2,
	0xf0, 0x12, 0x60, 0x41, 0xc4, 0xc6, 0x48, 0x5c, 0x5f, 0xce, 0xc3, 0xa6, 0x25, 0xce, 0x5f, 0x16,
	0x39, 0x97, 0x76, 0x91, 0x65, 0xb5, 0x8a, 0xe4, 0x06, 0xa0, 0x96, 0xe0, 0x15, 0x0c, 0xf7, 0x7e,
	0xc2, 0xcf, 0xec, 0x26, 0x03, 0xf9, 0x64, 0x48, 0xa7, 0x30, 0xfe, 0xf4, 0xe0, 0xbf, 0x54, 0x66,
	0x2a, 0x5d, 0x1a, 0xfe, 0xbb, 0xb3, 0xc3, 0x3e, 0xfe, 0xaf, 0x7d, 0xc6, 0x10, 0xac, 0xf5, 0x8e,
	0xcb, 0x43, 0x61, 0x81, 0x75, 0xc2, 0xdd, 0x25, 0xb4, 0xed, 0xab, 0x93, 0x44, 0xf6, 0x08, 0x10,
	0xa1, 0x2b, 0x1d, 0x07, 0x62, 0x74, 0x49, 0x67, 0x7c, 0x73, 0x07, 0xa3, 0x94, 0xe9, 0x3e, 0x2f,
	0x8a, 0x94, 0xeb, 0x8f, 0x9c, 0x18, 0xaf, 0xc1, 0x55, 0xec, 0x08, 0x8e, 0xe6, 0xd5, 0x76, 0xfe,
	0xf3, 0x7b, 0x27, 0x67, 0xad, 0xde, 0xcf, 0x1e, 0xff, 0xdb, 0x86, 0xf2, 0x16, 0x6e, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xe3, 0x7f, 0x85, 0x32, 0x1c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecKillServiceClient is the client API for SecKillService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecKillServiceClient interface {
	SecKill(ctx context.Context, in *SecRequest, opts ...grpc.CallOption) (*SecResponse, error)
}

type secKillServiceClient struct {
	cc *grpc.ClientConn
}

func NewSecKillServiceClient(cc *grpc.ClientConn) SecKillServiceClient {
	return &secKillServiceClient{cc}
}

func (c *secKillServiceClient) SecKill(ctx context.Context, in *SecRequest, opts ...grpc.CallOption) (*SecResponse, error) {
	out := new(SecResponse)
	err := c.cc.Invoke(ctx, "/pb.SecKillService/secKill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecKillServiceServer is the server API for SecKillService service.
type SecKillServiceServer interface {
	SecKill(context.Context, *SecRequest) (*SecResponse, error)
}

// UnimplementedSecKillServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSecKillServiceServer struct {
}

func (*UnimplementedSecKillServiceServer) SecKill(ctx context.Context, req *SecRequest) (*SecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SecKill not implemented")
}

func RegisterSecKillServiceServer(s *grpc.Server, srv SecKillServiceServer) {
	s.RegisterService(&_SecKillService_serviceDesc, srv)
}

func _SecKillService_SecKill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecKillServiceServer).SecKill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SecKillService/SecKill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecKillServiceServer).SecKill(ctx, req.(*SecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecKillService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SecKillService",
	HandlerType: (*SecKillServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "secKill",
			Handler:    _SecKillService_SecKill_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "seckill.proto",
}
