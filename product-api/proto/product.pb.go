// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Product struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price                float32  `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{0}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Product) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type ProductList struct {
	Products             []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProductList) Reset()         { *m = ProductList{} }
func (m *ProductList) String() string { return proto.CompactTextString(m) }
func (*ProductList) ProtoMessage()    {}
func (*ProductList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{1}
}

func (m *ProductList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductList.Unmarshal(m, b)
}
func (m *ProductList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductList.Marshal(b, m, deterministic)
}
func (m *ProductList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductList.Merge(m, src)
}
func (m *ProductList) XXX_Size() int {
	return xxx_messageInfo_ProductList.Size(m)
}
func (m *ProductList) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductList.DiscardUnknown(m)
}

var xxx_messageInfo_ProductList proto.InternalMessageInfo

func (m *ProductList) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

type ProductQuery struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductQuery) Reset()         { *m = ProductQuery{} }
func (m *ProductQuery) String() string { return proto.CompactTextString(m) }
func (*ProductQuery) ProtoMessage()    {}
func (*ProductQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{2}
}

func (m *ProductQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductQuery.Unmarshal(m, b)
}
func (m *ProductQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductQuery.Marshal(b, m, deterministic)
}
func (m *ProductQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductQuery.Merge(m, src)
}
func (m *ProductQuery) XXX_Size() int {
	return xxx_messageInfo_ProductQuery.Size(m)
}
func (m *ProductQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ProductQuery proto.InternalMessageInfo

func (m *ProductQuery) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Product)(nil), "proto.Product")
	proto.RegisterType((*ProductList)(nil), "proto.ProductList")
	proto.RegisterType((*ProductQuery)(nil), "proto.ProductQuery")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5) }

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xb3, 0x69, 0xe3, 0x9f, 0x89, 0xcd, 0x61, 0x14, 0x09, 0x11, 0x24, 0xec, 0x29, 0x78,
	0x48, 0xa1, 0x05, 0xa1, 0x78, 0xf2, 0x20, 0x5e, 0x3c, 0x68, 0xc4, 0x0f, 0xd0, 0x66, 0xc7, 0xb2,
	0x60, 0xbb, 0xcb, 0x66, 0x23, 0xf4, 0x83, 0x7b, 0x97, 0x6e, 0xd6, 0x6a, 0x0a, 0x1e, 0x7a, 0xda,
	0x9d, 0x37, 0xef, 0x31, 0x3f, 0x1e, 0x8c, 0xb4, 0x51, 0xa2, 0xad, 0x6d, 0xa9, 0x8d, 0xb2, 0x0a,
	0x23, 0xf7, 0x64, 0x57, 0x4b, 0xa5, 0x96, 0x1f, 0x34, 0x76, 0xd3, 0xa2, 0x7d, 0x1f, 0xd3, 0x4a,
	0xdb, 0x4d, 0xe7, 0xe1, 0x04, 0xc7, 0xcf, 0x5d, 0x08, 0x13, 0x08, 0xa5, 0x48, 0x59, 0xce, 0x8a,
	0xa8, 0x0a, 0xa5, 0x40, 0x84, 0xe1, 0x7a, 0xbe, 0xa2, 0x34, 0xcc, 0x59, 0x71, 0x5a, 0xb9, 0x3f,
	0xe6, 0x10, 0x0b, 0x6a, 0x6a, 0x23, 0xb5, 0x95, 0x6a, 0x9d, 0x0e, 0xdc, 0xea, 0xaf, 0x84, 0x17,
	0x10, 0x69, 0x23, 0x6b, 0x4a, 0x87, 0x39, 0x2b, 0xc2, 0xaa, 0x1b, 0xf8, 0x0c, 0x62, 0x7f, 0xe6,
	0x49, 0x36, 0x16, 0x6f, 0xe0, 0xc4, 0xa3, 0x36, 0x29, 0xcb, 0x07, 0x45, 0x3c, 0x49, 0x3a, 0x9e,
	0xd2, 0xbb, 0xaa, 0xdd, 0x9e, 0x5f, 0xc3, 0x99, 0x17, 0x5f, 0x5a, 0x32, 0x9b, 0x7d, 0xcc, 0xc9,
	0x17, 0x83, 0xc4, 0x1b, 0x5e, 0xc9, 0x7c, 0xca, 0x9a, 0xf0, 0x0e, 0xe2, 0x47, 0xb2, 0x5e, 0x6c,
	0xf0, 0xb2, 0xec, 0x1a, 0x28, 0x7f, 0x1a, 0x28, 0x1f, 0xb6, 0x0d, 0x64, 0xd8, 0xbf, 0xb9, 0x25,
	0xe3, 0x01, 0x4e, 0x01, 0x7e, 0xc3, 0x78, 0xde, 0xf7, 0x38, 0x84, 0x6c, 0x0f, 0x96, 0x07, 0x78,
	0x0b, 0x70, 0x2f, 0xc4, 0xae, 0xc9, 0xfe, 0x3e, 0xfb, 0x07, 0x80, 0x07, 0x38, 0x83, 0xd1, 0x9b,
	0x16, 0x73, 0x4b, 0x07, 0x47, 0x17, 0x47, 0x4e, 0x99, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x61,
	0xb4, 0x60, 0x22, 0xf5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	GetProducts(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProductList, error)
	GetProduct(ctx context.Context, in *ProductQuery, opts ...grpc.CallOption) (*Product, error)
	AddProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*empty.Empty, error)
}

type productServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductServiceClient(cc *grpc.ClientConn) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetProducts(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProductList, error) {
	out := new(ProductList)
	err := c.cc.Invoke(ctx, "/proto.ProductService/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProduct(ctx context.Context, in *ProductQuery, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/proto.ProductService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) AddProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.ProductService/AddProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.ProductService/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	GetProducts(context.Context, *empty.Empty) (*ProductList, error)
	GetProduct(context.Context, *ProductQuery) (*Product, error)
	AddProduct(context.Context, *Product) (*empty.Empty, error)
	UpdateProduct(context.Context, *Product) (*empty.Empty, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) GetProducts(ctx context.Context, req *empty.Empty) (*ProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (*UnimplementedProductServiceServer) GetProduct(ctx context.Context, req *ProductQuery) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (*UnimplementedProductServiceServer) AddProduct(ctx context.Context, req *Product) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (*UnimplementedProductServiceServer) UpdateProduct(ctx context.Context, req *Product) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProducts(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProduct(ctx, req.(*ProductQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/AddProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProducts",
			Handler:    _ProductService_GetProducts_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductService_GetProduct_Handler,
		},
		{
			MethodName: "AddProduct",
			Handler:    _ProductService_AddProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProductService_UpdateProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
