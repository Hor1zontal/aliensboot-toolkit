// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protocol.proto

/*
	Package scene is a generated protocol buffer package.

	It is generated from these files:
		protocol.proto
		scene.proto

	It has these top-level messages:
		SceneRequest
		SceneResponse
		ScenePush
		PushResult
		Vector
		Entity
		SpaceEnter
		SpaceEnterRet
		SpaceMove
		SpaceMoveRet
		SpaceLeave
		SpaceLeaveRet
		GetState
		GetStateRet
*/
package scene

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SceneRequest struct {
	Sequence         *int32      `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	Session          *int32      `protobuf:"varint,2,opt,name=session" json:"session,omitempty"`
	ClientID         *int32      `protobuf:"varint,3,opt,name=clientID" json:"clientID,omitempty"`
	SpaceMove        *SpaceMove  `protobuf:"bytes,5,opt,name=spaceMove" json:"spaceMove,omitempty"`
	SpaceEnter       *SpaceEnter `protobuf:"bytes,6,opt,name=spaceEnter" json:"spaceEnter,omitempty"`
	SpaceLeave       *SpaceLeave `protobuf:"bytes,7,opt,name=spaceLeave" json:"spaceLeave,omitempty"`
	GetState         *GetState   `protobuf:"bytes,8,opt,name=getState" json:"getState,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SceneRequest) Reset()                    { *m = SceneRequest{} }
func (m *SceneRequest) String() string            { return proto.CompactTextString(m) }
func (*SceneRequest) ProtoMessage()               {}
func (*SceneRequest) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{0} }

func (m *SceneRequest) GetSequence() int32 {
	if m != nil && m.Sequence != nil {
		return *m.Sequence
	}
	return 0
}

func (m *SceneRequest) GetSession() int32 {
	if m != nil && m.Session != nil {
		return *m.Session
	}
	return 0
}

func (m *SceneRequest) GetClientID() int32 {
	if m != nil && m.ClientID != nil {
		return *m.ClientID
	}
	return 0
}

func (m *SceneRequest) GetSpaceMove() *SpaceMove {
	if m != nil {
		return m.SpaceMove
	}
	return nil
}

func (m *SceneRequest) GetSpaceEnter() *SpaceEnter {
	if m != nil {
		return m.SpaceEnter
	}
	return nil
}

func (m *SceneRequest) GetSpaceLeave() *SpaceLeave {
	if m != nil {
		return m.SpaceLeave
	}
	return nil
}

func (m *SceneRequest) GetGetState() *GetState {
	if m != nil {
		return m.GetState
	}
	return nil
}

type SceneResponse struct {
	Sequence         *int32         `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	Session          *int32         `protobuf:"varint,2,opt,name=session" json:"session,omitempty"`
	SpaceMoveRet     *SpaceMoveRet  `protobuf:"bytes,5,opt,name=spaceMoveRet" json:"spaceMoveRet,omitempty"`
	SpaceEnterRet    *SpaceEnterRet `protobuf:"bytes,6,opt,name=spaceEnterRet" json:"spaceEnterRet,omitempty"`
	SpaceLeaveRet    *SpaceLeaveRet `protobuf:"bytes,7,opt,name=spaceLeaveRet" json:"spaceLeaveRet,omitempty"`
	GetStateRet      *GetStateRet   `protobuf:"bytes,8,opt,name=getStateRet" json:"getStateRet,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *SceneResponse) Reset()                    { *m = SceneResponse{} }
func (m *SceneResponse) String() string            { return proto.CompactTextString(m) }
func (*SceneResponse) ProtoMessage()               {}
func (*SceneResponse) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{1} }

func (m *SceneResponse) GetSequence() int32 {
	if m != nil && m.Sequence != nil {
		return *m.Sequence
	}
	return 0
}

func (m *SceneResponse) GetSession() int32 {
	if m != nil && m.Session != nil {
		return *m.Session
	}
	return 0
}

func (m *SceneResponse) GetSpaceMoveRet() *SpaceMoveRet {
	if m != nil {
		return m.SpaceMoveRet
	}
	return nil
}

func (m *SceneResponse) GetSpaceEnterRet() *SpaceEnterRet {
	if m != nil {
		return m.SpaceEnterRet
	}
	return nil
}

func (m *SceneResponse) GetSpaceLeaveRet() *SpaceLeaveRet {
	if m != nil {
		return m.SpaceLeaveRet
	}
	return nil
}

func (m *SceneResponse) GetGetStateRet() *GetStateRet {
	if m != nil {
		return m.GetStateRet
	}
	return nil
}

type ScenePush struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ScenePush) Reset()                    { *m = ScenePush{} }
func (m *ScenePush) String() string            { return proto.CompactTextString(m) }
func (*ScenePush) ProtoMessage()               {}
func (*ScenePush) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{2} }

type PushResult struct {
	ClientID         *int32 `protobuf:"varint,1,opt,name=clientID" json:"clientID,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PushResult) Reset()                    { *m = PushResult{} }
func (m *PushResult) String() string            { return proto.CompactTextString(m) }
func (*PushResult) ProtoMessage()               {}
func (*PushResult) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{3} }

func (m *PushResult) GetClientID() int32 {
	if m != nil && m.ClientID != nil {
		return *m.ClientID
	}
	return 0
}

func init() {
	proto.RegisterType((*SceneRequest)(nil), "scene.SceneRequest")
	proto.RegisterType((*SceneResponse)(nil), "scene.SceneResponse")
	proto.RegisterType((*ScenePush)(nil), "scene.ScenePush")
	proto.RegisterType((*PushResult)(nil), "scene.PushResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RPCService service

type RPCServiceClient interface {
	Request(ctx context.Context, in *SceneRequest, opts ...grpc.CallOption) (*SceneResponse, error)
}

type rPCServiceClient struct {
	cc *grpc.ClientConn
}

func NewRPCServiceClient(cc *grpc.ClientConn) RPCServiceClient {
	return &rPCServiceClient{cc}
}

func (c *rPCServiceClient) Request(ctx context.Context, in *SceneRequest, opts ...grpc.CallOption) (*SceneResponse, error) {
	out := new(SceneResponse)
	err := grpc.Invoke(ctx, "/scene.RPCService/request", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RPCService service

type RPCServiceServer interface {
	Request(context.Context, *SceneRequest) (*SceneResponse, error)
}

func RegisterRPCServiceServer(s *grpc.Server, srv RPCServiceServer) {
	s.RegisterService(&_RPCService_serviceDesc, srv)
}

func _RPCService_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SceneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scene.RPCService/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).Request(ctx, req.(*SceneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scene.RPCService",
	HandlerType: (*RPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "request",
			Handler:    _RPCService_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol.proto",
}

// Client API for RPCPush service

type RPCPushClient interface {
	Push(ctx context.Context, opts ...grpc.CallOption) (RPCPush_PushClient, error)
}

type rPCPushClient struct {
	cc *grpc.ClientConn
}

func NewRPCPushClient(cc *grpc.ClientConn) RPCPushClient {
	return &rPCPushClient{cc}
}

func (c *rPCPushClient) Push(ctx context.Context, opts ...grpc.CallOption) (RPCPush_PushClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RPCPush_serviceDesc.Streams[0], c.cc, "/scene.RPCPush/push", opts...)
	if err != nil {
		return nil, err
	}
	x := &rPCPushPushClient{stream}
	return x, nil
}

type RPCPush_PushClient interface {
	Send(*ScenePush) error
	CloseAndRecv() (*PushResult, error)
	grpc.ClientStream
}

type rPCPushPushClient struct {
	grpc.ClientStream
}

func (x *rPCPushPushClient) Send(m *ScenePush) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rPCPushPushClient) CloseAndRecv() (*PushResult, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PushResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RPCPush service

type RPCPushServer interface {
	Push(RPCPush_PushServer) error
}

func RegisterRPCPushServer(s *grpc.Server, srv RPCPushServer) {
	s.RegisterService(&_RPCPush_serviceDesc, srv)
}

func _RPCPush_Push_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RPCPushServer).Push(&rPCPushPushServer{stream})
}

type RPCPush_PushServer interface {
	SendAndClose(*PushResult) error
	Recv() (*ScenePush, error)
	grpc.ServerStream
}

type rPCPushPushServer struct {
	grpc.ServerStream
}

func (x *rPCPushPushServer) SendAndClose(m *PushResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rPCPushPushServer) Recv() (*ScenePush, error) {
	m := new(ScenePush)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RPCPush_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scene.RPCPush",
	HandlerType: (*RPCPushServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "push",
			Handler:       _RPCPush_Push_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "protocol.proto",
}

func (m *SceneRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SceneRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Sequence != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(*m.Sequence))
	}
	if m.Session != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(*m.Session))
	}
	if m.ClientID != nil {
		dAtA[i] = 0x18
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(*m.ClientID))
	}
	if m.SpaceMove != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.SpaceMove.Size()))
		n1, err := m.SpaceMove.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.SpaceEnter != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.SpaceEnter.Size()))
		n2, err := m.SpaceEnter.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.SpaceLeave != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.SpaceLeave.Size()))
		n3, err := m.SpaceLeave.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.GetState != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.GetState.Size()))
		n4, err := m.GetState.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SceneResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SceneResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Sequence != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(*m.Sequence))
	}
	if m.Session != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(*m.Session))
	}
	if m.SpaceMoveRet != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.SpaceMoveRet.Size()))
		n5, err := m.SpaceMoveRet.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.SpaceEnterRet != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.SpaceEnterRet.Size()))
		n6, err := m.SpaceEnterRet.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.SpaceLeaveRet != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.SpaceLeaveRet.Size()))
		n7, err := m.SpaceLeaveRet.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.GetStateRet != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.GetStateRet.Size()))
		n8, err := m.GetStateRet.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ScenePush) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScenePush) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PushResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ClientID != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(*m.ClientID))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Protocol(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Protocol(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintProtocol(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SceneRequest) Size() (n int) {
	var l int
	_ = l
	if m.Sequence != nil {
		n += 1 + sovProtocol(uint64(*m.Sequence))
	}
	if m.Session != nil {
		n += 1 + sovProtocol(uint64(*m.Session))
	}
	if m.ClientID != nil {
		n += 1 + sovProtocol(uint64(*m.ClientID))
	}
	if m.SpaceMove != nil {
		l = m.SpaceMove.Size()
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.SpaceEnter != nil {
		l = m.SpaceEnter.Size()
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.SpaceLeave != nil {
		l = m.SpaceLeave.Size()
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.GetState != nil {
		l = m.GetState.Size()
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SceneResponse) Size() (n int) {
	var l int
	_ = l
	if m.Sequence != nil {
		n += 1 + sovProtocol(uint64(*m.Sequence))
	}
	if m.Session != nil {
		n += 1 + sovProtocol(uint64(*m.Session))
	}
	if m.SpaceMoveRet != nil {
		l = m.SpaceMoveRet.Size()
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.SpaceEnterRet != nil {
		l = m.SpaceEnterRet.Size()
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.SpaceLeaveRet != nil {
		l = m.SpaceLeaveRet.Size()
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.GetStateRet != nil {
		l = m.GetStateRet.Size()
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ScenePush) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PushResult) Size() (n int) {
	var l int
	_ = l
	if m.ClientID != nil {
		n += 1 + sovProtocol(uint64(*m.ClientID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProtocol(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProtocol(x uint64) (n int) {
	return sovProtocol(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SceneRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SceneRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SceneRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Sequence = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Session", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Session = &v
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ClientID = &v
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceMove", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SpaceMove == nil {
				m.SpaceMove = &SpaceMove{}
			}
			if err := m.SpaceMove.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceEnter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SpaceEnter == nil {
				m.SpaceEnter = &SpaceEnter{}
			}
			if err := m.SpaceEnter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceLeave", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SpaceLeave == nil {
				m.SpaceLeave = &SpaceLeave{}
			}
			if err := m.SpaceLeave.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GetState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GetState == nil {
				m.GetState = &GetState{}
			}
			if err := m.GetState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SceneResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SceneResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SceneResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Sequence = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Session", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Session = &v
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceMoveRet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SpaceMoveRet == nil {
				m.SpaceMoveRet = &SpaceMoveRet{}
			}
			if err := m.SpaceMoveRet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceEnterRet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SpaceEnterRet == nil {
				m.SpaceEnterRet = &SpaceEnterRet{}
			}
			if err := m.SpaceEnterRet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceLeaveRet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SpaceLeaveRet == nil {
				m.SpaceLeaveRet = &SpaceLeaveRet{}
			}
			if err := m.SpaceLeaveRet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GetStateRet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GetStateRet == nil {
				m.GetStateRet = &GetStateRet{}
			}
			if err := m.GetStateRet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ScenePush) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ScenePush: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScenePush: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PushResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PushResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ClientID = &v
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProtocol(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtocol
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
					return 0, ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProtocol
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthProtocol
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProtocol
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipProtocol(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthProtocol = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtocol   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("protocol.proto", fileDescriptorProtocol) }

var fileDescriptorProtocol = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x52, 0xcd, 0x4e, 0xea, 0x40,
	0x14, 0x66, 0xb8, 0x97, 0x5b, 0x38, 0x05, 0x2e, 0x77, 0xae, 0x8b, 0x49, 0x17, 0x0d, 0xe9, 0xc2,
	0xd4, 0xc4, 0x54, 0x83, 0x09, 0x0b, 0x97, 0x82, 0x0b, 0x13, 0x4d, 0xc8, 0xf0, 0x00, 0x84, 0x34,
	0x27, 0x48, 0x52, 0xdb, 0xda, 0x19, 0xfa, 0x1a, 0x6e, 0x5d, 0xf8, 0x40, 0x2e, 0x7d, 0x04, 0x83,
	0x2f, 0x62, 0x66, 0xfa, 0x37, 0xdd, 0xba, 0x3a, 0x73, 0xbe, 0x9f, 0xc0, 0xf7, 0xf5, 0xc0, 0x38,
	0xcd, 0x12, 0x99, 0x84, 0x49, 0x14, 0xe8, 0x07, 0xed, 0x89, 0x10, 0x63, 0x74, 0x6c, 0x3d, 0x0a,
	0xcc, 0x7b, 0xe9, 0xc2, 0x70, 0xad, 0x76, 0x8e, 0xcf, 0x07, 0x14, 0x92, 0x3a, 0xd0, 0x17, 0xea,
	0x19, 0x87, 0xc8, 0xc8, 0x94, 0xf8, 0x3d, 0x5e, 0xef, 0x94, 0x81, 0x25, 0x50, 0x88, 0x7d, 0x12,
	0xb3, 0xae, 0xa6, 0xaa, 0x55, 0xb9, 0xc2, 0x68, 0x8f, 0xb1, 0xbc, 0x5b, 0xb2, 0x5f, 0x85, 0xab,
	0xda, 0xe9, 0x19, 0x0c, 0x44, 0xba, 0x0d, 0xf1, 0x21, 0xc9, 0x91, 0xf5, 0xa6, 0xc4, 0xb7, 0x67,
	0x76, 0xa0, 0x91, 0xcd, 0x53, 0x92, 0x23, 0x6f, 0x58, 0x7a, 0x0e, 0xa0, 0x97, 0xdb, 0x58, 0x62,
	0xc6, 0xfe, 0x68, 0xed, 0xb0, 0xd4, 0xa2, 0xc2, 0xb8, 0xc1, 0xd7, 0xea, 0x7b, 0xdc, 0xe6, 0xc8,
	0xac, 0x96, 0x3a, 0x52, 0x18, 0x37, 0x78, 0x7a, 0x0a, 0xfd, 0x1d, 0xca, 0xb5, 0xdc, 0x4a, 0x64,
	0x7d, 0xad, 0x85, 0x60, 0x87, 0x72, 0x23, 0x14, 0xc2, 0x6b, 0xce, 0x7b, 0xeb, 0xc2, 0xa8, 0x6c,
	0x44, 0xa4, 0x49, 0x2c, 0xf0, 0x87, 0x95, 0x5c, 0xc1, 0xb0, 0x0e, 0xc6, 0x51, 0x96, 0xc9, 0xff,
	0x1a, 0xc9, 0x37, 0x19, 0x4a, 0xde, 0x12, 0xd1, 0x39, 0x8c, 0x9a, 0x80, 0xca, 0x55, 0x74, 0x30,
	0x31, 0x3b, 0xd0, 0xb6, 0xb6, 0xac, 0xf6, 0xe9, 0xa8, 0xca, 0x67, 0xb5, 0x7c, 0xba, 0x0d, 0xc3,
	0x57, 0xc9, 0xe8, 0x25, 0xd8, 0x55, 0x70, 0xe5, 0x2a, 0x7a, 0x19, 0x37, 0xbd, 0x68, 0x8f, 0x29,
	0xf1, 0x6c, 0x18, 0xe8, 0x76, 0x56, 0x07, 0xf1, 0xe8, 0xf9, 0x00, 0x6a, 0x72, 0x14, 0x87, 0x48,
	0xb6, 0x8e, 0x80, 0xb4, 0x8f, 0x60, 0xb6, 0x04, 0xe0, 0xab, 0xc5, 0x1a, 0xb3, 0x7c, 0x1f, 0x22,
	0x9d, 0x83, 0x95, 0x95, 0xf7, 0xf6, 0x3f, 0x28, 0xce, 0xd1, 0x3c, 0x42, 0xe7, 0xa4, 0x0d, 0x16,
	0xdf, 0xc1, 0xeb, 0xcc, 0xae, 0xc1, 0xe2, 0xab, 0x85, 0xfa, 0x49, 0x7a, 0x01, 0xbf, 0x53, 0x35,
	0x27, 0xa6, 0x54, 0x31, 0xce, 0xbf, 0x12, 0x69, 0xfe, 0x99, 0xd7, 0xf1, 0xc9, 0xcd, 0xe4, 0xfd,
	0xe8, 0x92, 0x8f, 0xa3, 0x4b, 0x3e, 0x8f, 0x2e, 0x79, 0xfd, 0x72, 0x3b, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xdc, 0x57, 0x5e, 0x90, 0x20, 0x03, 0x00, 0x00,
}
