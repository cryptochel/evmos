// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evmos/ibc/evm/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgIBCEthereumTx defines an Ethereum tx to
type MsgIBCEthereumTx struct {
	// the port on which the packet will be sent
	SourcePort string `protobuf:"bytes,1,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty" yaml:"source_port"`
	// the channel by which the packet will be sent
	SourceChannel string `protobuf:"bytes,2,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty" yaml:"source_channel"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight types.Height `protobuf:"bytes,3,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height" yaml:"timeout_height"`
	// Timeout timestamp (in nanoseconds) relative to the current block timestamp.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64 `protobuf:"varint,4,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
	// Ethereum tx RLP encoded bytes
	EthereumTx []byte `protobuf:"bytes,5,opt,name=ethereum_tx,json=ethereumTx,proto3" json:"ethereum_tx,omitempty"`
}

func (m *MsgIBCEthereumTx) Reset()         { *m = MsgIBCEthereumTx{} }
func (m *MsgIBCEthereumTx) String() string { return proto.CompactTextString(m) }
func (*MsgIBCEthereumTx) ProtoMessage()    {}
func (*MsgIBCEthereumTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40d98923f08b90e, []int{0}
}
func (m *MsgIBCEthereumTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIBCEthereumTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIBCEthereumTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIBCEthereumTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIBCEthereumTx.Merge(m, src)
}
func (m *MsgIBCEthereumTx) XXX_Size() int {
	return m.Size()
}
func (m *MsgIBCEthereumTx) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIBCEthereumTx.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIBCEthereumTx proto.InternalMessageInfo

// MsgIBCEthereumTxResponse defines the Msg/MsgIBCEthereumTx response type.
type MsgIBCEthereumTxResponse struct {
}

func (m *MsgIBCEthereumTxResponse) Reset()         { *m = MsgIBCEthereumTxResponse{} }
func (m *MsgIBCEthereumTxResponse) String() string { return proto.CompactTextString(m) }
func (*MsgIBCEthereumTxResponse) ProtoMessage()    {}
func (*MsgIBCEthereumTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40d98923f08b90e, []int{1}
}
func (m *MsgIBCEthereumTxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIBCEthereumTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIBCEthereumTxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIBCEthereumTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIBCEthereumTxResponse.Merge(m, src)
}
func (m *MsgIBCEthereumTxResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgIBCEthereumTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIBCEthereumTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIBCEthereumTxResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgIBCEthereumTx)(nil), "evmos.ibc.evm.v1.MsgIBCEthereumTx")
	proto.RegisterType((*MsgIBCEthereumTxResponse)(nil), "evmos.ibc.evm.v1.MsgIBCEthereumTxResponse")
}

func init() { proto.RegisterFile("evmos/ibc/evm/v1/tx.proto", fileDescriptor_a40d98923f08b90e) }

var fileDescriptor_a40d98923f08b90e = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0xce, 0x6c, 0xab, 0xe8, 0xac, 0x5b, 0xd6, 0xa0, 0x92, 0x2e, 0x35, 0x09, 0x73, 0x0a, 0x1e,
	0x32, 0xa4, 0x1e, 0x84, 0x9e, 0x24, 0x45, 0xb0, 0x87, 0x82, 0x84, 0x9e, 0xbc, 0xac, 0x49, 0x78,
	0x24, 0x03, 0x9b, 0xbc, 0x90, 0x99, 0x0d, 0xdb, 0xab, 0x27, 0xbd, 0x29, 0xfe, 0x03, 0xfd, 0x73,
	0x7a, 0x2c, 0x78, 0xf1, 0xb4, 0xc8, 0xae, 0x82, 0xe7, 0xfd, 0x0b, 0x24, 0x3f, 0x76, 0xed, 0x2e,
	0x42, 0x4f, 0x79, 0xf3, 0x7d, 0xdf, 0x7b, 0xdf, 0xcc, 0x97, 0x47, 0x0f, 0xa1, 0xca, 0x50, 0x72,
	0x11, 0xc5, 0x1c, 0xaa, 0x8c, 0x57, 0x1e, 0x57, 0x33, 0xb7, 0x28, 0x51, 0xa1, 0x3e, 0x6c, 0x28,
	0x57, 0x44, 0xb1, 0x0b, 0x55, 0xe6, 0x56, 0xde, 0xe8, 0x28, 0x41, 0x4c, 0x26, 0xc0, 0xc3, 0x42,
	0xf0, 0x30, 0xcf, 0x51, 0x85, 0x4a, 0x60, 0x2e, 0x5b, 0xfd, 0xe8, 0x49, 0x82, 0x09, 0x36, 0x25,
	0xaf, 0xab, 0x0e, 0x35, 0x63, 0x94, 0xb5, 0x43, 0x14, 0x4a, 0xe0, 0x95, 0x17, 0x81, 0x0a, 0x3d,
	0x1e, 0xa3, 0xc8, 0x3b, 0xde, 0xaa, 0xad, 0x63, 0x2c, 0x81, 0xc7, 0x13, 0x01, 0xb9, 0xaa, 0xaf,
	0xd0, 0x56, 0xad, 0x80, 0xfd, 0xee, 0xd1, 0xe1, 0xb9, 0x4c, 0xce, 0xfc, 0xd3, 0x37, 0x2a, 0x85,
	0x12, 0xa6, 0xd9, 0xc5, 0x4c, 0x7f, 0x45, 0xfb, 0x12, 0xa7, 0x65, 0x0c, 0xe3, 0x02, 0x4b, 0x65,
	0x10, 0x9b, 0x38, 0x0f, 0xfd, 0x67, 0xab, 0xb9, 0xa5, 0x5f, 0x86, 0xd9, 0xe4, 0x84, 0xdd, 0x22,
	0x59, 0x40, 0xdb, 0xd3, 0x3b, 0x2c, 0x95, 0xfe, 0x9a, 0x1e, 0x74, 0x5c, 0x9c, 0x86, 0x79, 0x0e,
	0x13, 0xa3, 0xd7, 0xf4, 0x1e, 0xae, 0xe6, 0xd6, 0xd3, 0xad, 0xde, 0x8e, 0x67, 0xc1, 0xa0, 0x05,
	0x4e, 0xdb, 0xb3, 0xfe, 0x81, 0x1e, 0x28, 0x91, 0x01, 0x4e, 0xd5, 0x38, 0x05, 0x91, 0xa4, 0xca,
	0xd8, 0xb3, 0x89, 0xd3, 0x3f, 0x1e, 0x35, 0x49, 0xd5, 0x2f, 0x71, 0xbb, 0xfb, 0x57, 0x9e, 0xfb,
	0xb6, 0x51, 0xf8, 0xcf, 0xaf, 0xe7, 0x96, 0xf6, 0xcf, 0x61, 0xbb, 0x9f, 0x05, 0x83, 0x0e, 0x68,
	0xd5, 0xfa, 0x19, 0x7d, 0xbc, 0x56, 0xd4, 0x5f, 0xa9, 0xc2, 0xac, 0x30, 0xf6, 0x6d, 0xe2, 0xec,
	0xfb, 0x47, 0xab, 0xb9, 0x65, 0x6c, 0x0f, 0xd9, 0x48, 0x58, 0x30, 0xec, 0xb0, 0x8b, 0x35, 0xa4,
	0x5b, 0xb4, 0x0f, 0x5d, 0x6a, 0x63, 0x35, 0x33, 0xee, 0xd9, 0xc4, 0x79, 0x14, 0x50, 0xd8, 0x04,
	0x79, 0xf2, 0xe0, 0xd3, 0x95, 0xa5, 0xfd, 0xb9, 0xb2, 0x34, 0x36, 0xa2, 0xc6, 0x6e, 0xcc, 0x01,
	0xc8, 0x02, 0x73, 0x09, 0xc7, 0x5f, 0x09, 0xdd, 0x3b, 0x97, 0x89, 0xfe, 0x99, 0xd0, 0xc1, 0xf6,
	0x8f, 0x60, 0xee, 0xee, 0x96, 0xb8, 0xbb, 0x53, 0x46, 0x2f, 0xee, 0xd6, 0xac, 0x9d, 0x98, 0xf3,
	0xf1, 0xfb, 0xaf, 0x6f, 0x3d, 0xa6, 0xdb, 0xfc, 0x3f, 0x8b, 0xc9, 0x6f, 0xbd, 0xc5, 0xf7, 0xaf,
	0x17, 0x26, 0xb9, 0x59, 0x98, 0xe4, 0xe7, 0xc2, 0x24, 0x5f, 0x96, 0xa6, 0x76, 0xb3, 0x34, 0xb5,
	0x1f, 0x4b, 0x53, 0x7b, 0xef, 0x24, 0x42, 0xa5, 0xd3, 0xc8, 0x8d, 0x31, 0xe3, 0x2a, 0x0d, 0x4b,
	0x29, 0x64, 0x37, 0x6d, 0xb6, 0x99, 0xa7, 0x2e, 0x0b, 0x90, 0xd1, 0xfd, 0x66, 0xc5, 0x5e, 0xfe,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xa9, 0x47, 0x54, 0x06, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// IBCEthereumTx
	IBCEthereumTx(ctx context.Context, in *MsgIBCEthereumTx, opts ...grpc.CallOption) (*MsgIBCEthereumTxResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) IBCEthereumTx(ctx context.Context, in *MsgIBCEthereumTx, opts ...grpc.CallOption) (*MsgIBCEthereumTxResponse, error) {
	out := new(MsgIBCEthereumTxResponse)
	err := c.cc.Invoke(ctx, "/evmos.ibc.evm.v1.Msg/IBCEthereumTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// IBCEthereumTx
	IBCEthereumTx(context.Context, *MsgIBCEthereumTx) (*MsgIBCEthereumTxResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) IBCEthereumTx(ctx context.Context, req *MsgIBCEthereumTx) (*MsgIBCEthereumTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IBCEthereumTx not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_IBCEthereumTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgIBCEthereumTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).IBCEthereumTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evmos.ibc.evm.v1.Msg/IBCEthereumTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).IBCEthereumTx(ctx, req.(*MsgIBCEthereumTx))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "evmos.ibc.evm.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IBCEthereumTx",
			Handler:    _Msg_IBCEthereumTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "evmos/ibc/evm/v1/tx.proto",
}

func (m *MsgIBCEthereumTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIBCEthereumTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIBCEthereumTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EthereumTx) > 0 {
		i -= len(m.EthereumTx)
		copy(dAtA[i:], m.EthereumTx)
		i = encodeVarintTx(dAtA, i, uint64(len(m.EthereumTx)))
		i--
		dAtA[i] = 0x2a
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SourcePort) > 0 {
		i -= len(m.SourcePort)
		copy(dAtA[i:], m.SourcePort)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourcePort)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgIBCEthereumTxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIBCEthereumTxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIBCEthereumTxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgIBCEthereumTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourcePort)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovTx(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTx(uint64(m.TimeoutTimestamp))
	}
	l = len(m.EthereumTx)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgIBCEthereumTxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgIBCEthereumTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgIBCEthereumTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIBCEthereumTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourcePort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthereumTx", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthereumTx = append(m.EthereumTx[:0], dAtA[iNdEx:postIndex]...)
			if m.EthereumTx == nil {
				m.EthereumTx = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgIBCEthereumTxResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgIBCEthereumTxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIBCEthereumTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
