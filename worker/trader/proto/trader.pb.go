// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.3
// source: trader.proto

package trader

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

type CreateMarketOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpType string `protobuf:"bytes,1,opt,name=opType,proto3" json:"opType,omitempty"`
	Ticker string `protobuf:"bytes,2,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Qty    int32  `protobuf:"varint,3,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *CreateMarketOrderRequest) Reset() {
	*x = CreateMarketOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMarketOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMarketOrderRequest) ProtoMessage() {}

func (x *CreateMarketOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMarketOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateMarketOrderRequest) Descriptor() ([]byte, []int) {
	return file_trader_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMarketOrderRequest) GetOpType() string {
	if x != nil {
		return x.OpType
	}
	return ""
}

func (x *CreateMarketOrderRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *CreateMarketOrderRequest) GetQty() int32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

type CreateMarketOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateMarketOrderResponse) Reset() {
	*x = CreateMarketOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMarketOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMarketOrderResponse) ProtoMessage() {}

func (x *CreateMarketOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMarketOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateMarketOrderResponse) Descriptor() ([]byte, []int) {
	return file_trader_proto_rawDescGZIP(), []int{1}
}

type CreateLimitOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpType string  `protobuf:"bytes,1,opt,name=opType,proto3" json:"opType,omitempty"`
	Ticker string  `protobuf:"bytes,2,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Price  float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Qty    int32   `protobuf:"varint,4,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *CreateLimitOrderRequest) Reset() {
	*x = CreateLimitOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trader_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLimitOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLimitOrderRequest) ProtoMessage() {}

func (x *CreateLimitOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trader_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLimitOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateLimitOrderRequest) Descriptor() ([]byte, []int) {
	return file_trader_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLimitOrderRequest) GetOpType() string {
	if x != nil {
		return x.OpType
	}
	return ""
}

func (x *CreateLimitOrderRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *CreateLimitOrderRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateLimitOrderRequest) GetQty() int32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

type CreateLimitOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateLimitOrderResponse) Reset() {
	*x = CreateLimitOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trader_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLimitOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLimitOrderResponse) ProtoMessage() {}

func (x *CreateLimitOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trader_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLimitOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateLimitOrderResponse) Descriptor() ([]byte, []int) {
	return file_trader_proto_rawDescGZIP(), []int{3}
}

var File_trader_proto protoreflect.FileDescriptor

var file_trader_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x22, 0x5c, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x71, 0x74, 0x79, 0x22, 0x1b, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x71, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x71, 0x74, 0x79, 0x22, 0x1a, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xbd, 0x01, 0x0a, 0x06, 0x54, 0x72, 0x61, 0x64, 0x65, 0x72, 0x12, 0x5a, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x20, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trader_proto_rawDescOnce sync.Once
	file_trader_proto_rawDescData = file_trader_proto_rawDesc
)

func file_trader_proto_rawDescGZIP() []byte {
	file_trader_proto_rawDescOnce.Do(func() {
		file_trader_proto_rawDescData = protoimpl.X.CompressGZIP(file_trader_proto_rawDescData)
	})
	return file_trader_proto_rawDescData
}

var file_trader_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_trader_proto_goTypes = []interface{}{
	(*CreateMarketOrderRequest)(nil),  // 0: trader.CreateMarketOrderRequest
	(*CreateMarketOrderResponse)(nil), // 1: trader.CreateMarketOrderResponse
	(*CreateLimitOrderRequest)(nil),   // 2: trader.CreateLimitOrderRequest
	(*CreateLimitOrderResponse)(nil),  // 3: trader.CreateLimitOrderResponse
}
var file_trader_proto_depIdxs = []int32{
	0, // 0: trader.Trader.CreateMarketOrder:input_type -> trader.CreateMarketOrderRequest
	2, // 1: trader.Trader.CreateLimitOrder:input_type -> trader.CreateLimitOrderRequest
	1, // 2: trader.Trader.CreateMarketOrder:output_type -> trader.CreateMarketOrderResponse
	3, // 3: trader.Trader.CreateLimitOrder:output_type -> trader.CreateLimitOrderResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_trader_proto_init() }
func file_trader_proto_init() {
	if File_trader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMarketOrderRequest); i {
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
		file_trader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMarketOrderResponse); i {
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
		file_trader_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLimitOrderRequest); i {
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
		file_trader_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLimitOrderResponse); i {
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
			RawDescriptor: file_trader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trader_proto_goTypes,
		DependencyIndexes: file_trader_proto_depIdxs,
		MessageInfos:      file_trader_proto_msgTypes,
	}.Build()
	File_trader_proto = out.File
	file_trader_proto_rawDesc = nil
	file_trader_proto_goTypes = nil
	file_trader_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TraderClient is the client API for Trader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraderClient interface {
	CreateMarketOrder(ctx context.Context, in *CreateMarketOrderRequest, opts ...grpc.CallOption) (*CreateMarketOrderResponse, error)
	CreateLimitOrder(ctx context.Context, in *CreateLimitOrderRequest, opts ...grpc.CallOption) (*CreateLimitOrderResponse, error)
}

type traderClient struct {
	cc grpc.ClientConnInterface
}

func NewTraderClient(cc grpc.ClientConnInterface) TraderClient {
	return &traderClient{cc}
}

func (c *traderClient) CreateMarketOrder(ctx context.Context, in *CreateMarketOrderRequest, opts ...grpc.CallOption) (*CreateMarketOrderResponse, error) {
	out := new(CreateMarketOrderResponse)
	err := c.cc.Invoke(ctx, "/trader.Trader/CreateMarketOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traderClient) CreateLimitOrder(ctx context.Context, in *CreateLimitOrderRequest, opts ...grpc.CallOption) (*CreateLimitOrderResponse, error) {
	out := new(CreateLimitOrderResponse)
	err := c.cc.Invoke(ctx, "/trader.Trader/CreateLimitOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TraderServer is the server API for Trader service.
type TraderServer interface {
	CreateMarketOrder(context.Context, *CreateMarketOrderRequest) (*CreateMarketOrderResponse, error)
	CreateLimitOrder(context.Context, *CreateLimitOrderRequest) (*CreateLimitOrderResponse, error)
}

// UnimplementedTraderServer can be embedded to have forward compatible implementations.
type UnimplementedTraderServer struct {
}

func (*UnimplementedTraderServer) CreateMarketOrder(context.Context, *CreateMarketOrderRequest) (*CreateMarketOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMarketOrder not implemented")
}
func (*UnimplementedTraderServer) CreateLimitOrder(context.Context, *CreateLimitOrderRequest) (*CreateLimitOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLimitOrder not implemented")
}

func RegisterTraderServer(s *grpc.Server, srv TraderServer) {
	s.RegisterService(&_Trader_serviceDesc, srv)
}

func _Trader_CreateMarketOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMarketOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraderServer).CreateMarketOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trader.Trader/CreateMarketOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraderServer).CreateMarketOrder(ctx, req.(*CreateMarketOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trader_CreateLimitOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLimitOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraderServer).CreateLimitOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trader.Trader/CreateLimitOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraderServer).CreateLimitOrder(ctx, req.(*CreateLimitOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Trader_serviceDesc = grpc.ServiceDesc{
	ServiceName: "trader.Trader",
	HandlerType: (*TraderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMarketOrder",
			Handler:    _Trader_CreateMarketOrder_Handler,
		},
		{
			MethodName: "CreateLimitOrder",
			Handler:    _Trader_CreateLimitOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trader.proto",
}
