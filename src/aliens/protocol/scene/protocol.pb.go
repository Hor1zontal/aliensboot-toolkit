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
	SpacePush
*/
package scene

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// request
type SceneRequest struct {
	Session  int32 `protobuf:"varint,1,opt,name=session,proto3" json:"session,omitempty"`
	ClientID int32 `protobuf:"varint,3,opt,name=clientID,proto3" json:"clientID,omitempty"`
	// Types that are valid to be assigned to Request:
	//	*SceneRequest_SpaceMove
	//	*SceneRequest_SpaceEnter
	//	*SceneRequest_SpaceLeave
	//	*SceneRequest_GetState
	Request isSceneRequest_Request `protobuf_oneof:"request"`
}

func (m *SceneRequest) Reset()                    { *m = SceneRequest{} }
func (m *SceneRequest) String() string            { return proto.CompactTextString(m) }
func (*SceneRequest) ProtoMessage()               {}
func (*SceneRequest) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{0} }

type isSceneRequest_Request interface {
	isSceneRequest_Request()
}

type SceneRequest_SpaceMove struct {
	SpaceMove *SpaceMove `protobuf:"bytes,5,opt,name=spaceMove,oneof"`
}
type SceneRequest_SpaceEnter struct {
	SpaceEnter *SpaceEnter `protobuf:"bytes,6,opt,name=spaceEnter,oneof"`
}
type SceneRequest_SpaceLeave struct {
	SpaceLeave *SpaceLeave `protobuf:"bytes,7,opt,name=spaceLeave,oneof"`
}
type SceneRequest_GetState struct {
	GetState *GetState `protobuf:"bytes,8,opt,name=getState,oneof"`
}

func (*SceneRequest_SpaceMove) isSceneRequest_Request()  {}
func (*SceneRequest_SpaceEnter) isSceneRequest_Request() {}
func (*SceneRequest_SpaceLeave) isSceneRequest_Request() {}
func (*SceneRequest_GetState) isSceneRequest_Request()   {}

func (m *SceneRequest) GetRequest() isSceneRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *SceneRequest) GetSession() int32 {
	if m != nil {
		return m.Session
	}
	return 0
}

func (m *SceneRequest) GetClientID() int32 {
	if m != nil {
		return m.ClientID
	}
	return 0
}

func (m *SceneRequest) GetSpaceMove() *SpaceMove {
	if x, ok := m.GetRequest().(*SceneRequest_SpaceMove); ok {
		return x.SpaceMove
	}
	return nil
}

func (m *SceneRequest) GetSpaceEnter() *SpaceEnter {
	if x, ok := m.GetRequest().(*SceneRequest_SpaceEnter); ok {
		return x.SpaceEnter
	}
	return nil
}

func (m *SceneRequest) GetSpaceLeave() *SpaceLeave {
	if x, ok := m.GetRequest().(*SceneRequest_SpaceLeave); ok {
		return x.SpaceLeave
	}
	return nil
}

func (m *SceneRequest) GetGetState() *GetState {
	if x, ok := m.GetRequest().(*SceneRequest_GetState); ok {
		return x.GetState
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SceneRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SceneRequest_OneofMarshaler, _SceneRequest_OneofUnmarshaler, _SceneRequest_OneofSizer, []interface{}{
		(*SceneRequest_SpaceMove)(nil),
		(*SceneRequest_SpaceEnter)(nil),
		(*SceneRequest_SpaceLeave)(nil),
		(*SceneRequest_GetState)(nil),
	}
}

func _SceneRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SceneRequest)
	// request
	switch x := m.Request.(type) {
	case *SceneRequest_SpaceMove:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SpaceMove); err != nil {
			return err
		}
	case *SceneRequest_SpaceEnter:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SpaceEnter); err != nil {
			return err
		}
	case *SceneRequest_SpaceLeave:
		_ = b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SpaceLeave); err != nil {
			return err
		}
	case *SceneRequest_GetState:
		_ = b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GetState); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SceneRequest.Request has unexpected type %T", x)
	}
	return nil
}

func _SceneRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SceneRequest)
	switch tag {
	case 5: // request.spaceMove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SpaceMove)
		err := b.DecodeMessage(msg)
		m.Request = &SceneRequest_SpaceMove{msg}
		return true, err
	case 6: // request.spaceEnter
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SpaceEnter)
		err := b.DecodeMessage(msg)
		m.Request = &SceneRequest_SpaceEnter{msg}
		return true, err
	case 7: // request.spaceLeave
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SpaceLeave)
		err := b.DecodeMessage(msg)
		m.Request = &SceneRequest_SpaceLeave{msg}
		return true, err
	case 8: // request.getState
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GetState)
		err := b.DecodeMessage(msg)
		m.Request = &SceneRequest_GetState{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SceneRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SceneRequest)
	// request
	switch x := m.Request.(type) {
	case *SceneRequest_SpaceMove:
		s := proto.Size(x.SpaceMove)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SceneRequest_SpaceEnter:
		s := proto.Size(x.SpaceEnter)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SceneRequest_SpaceLeave:
		s := proto.Size(x.SpaceLeave)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SceneRequest_GetState:
		s := proto.Size(x.GetState)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// response
type SceneResponse struct {
	Session int32 `protobuf:"varint,1,opt,name=session,proto3" json:"session,omitempty"`
	// Types that are valid to be assigned to Response:
	//	*SceneResponse_SpaceMoveRet
	//	*SceneResponse_SpaceEnterRet
	//	*SceneResponse_SpaceLeaveRet
	//	*SceneResponse_GetStateRet
	Response isSceneResponse_Response `protobuf_oneof:"response"`
}

func (m *SceneResponse) Reset()                    { *m = SceneResponse{} }
func (m *SceneResponse) String() string            { return proto.CompactTextString(m) }
func (*SceneResponse) ProtoMessage()               {}
func (*SceneResponse) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{1} }

type isSceneResponse_Response interface {
	isSceneResponse_Response()
}

type SceneResponse_SpaceMoveRet struct {
	SpaceMoveRet *SpaceMoveRet `protobuf:"bytes,5,opt,name=spaceMoveRet,oneof"`
}
type SceneResponse_SpaceEnterRet struct {
	SpaceEnterRet *SpaceEnterRet `protobuf:"bytes,6,opt,name=spaceEnterRet,oneof"`
}
type SceneResponse_SpaceLeaveRet struct {
	SpaceLeaveRet *SpaceLeaveRet `protobuf:"bytes,7,opt,name=spaceLeaveRet,oneof"`
}
type SceneResponse_GetStateRet struct {
	GetStateRet *GetStateRet `protobuf:"bytes,8,opt,name=getStateRet,oneof"`
}

func (*SceneResponse_SpaceMoveRet) isSceneResponse_Response()  {}
func (*SceneResponse_SpaceEnterRet) isSceneResponse_Response() {}
func (*SceneResponse_SpaceLeaveRet) isSceneResponse_Response() {}
func (*SceneResponse_GetStateRet) isSceneResponse_Response()   {}

func (m *SceneResponse) GetResponse() isSceneResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *SceneResponse) GetSession() int32 {
	if m != nil {
		return m.Session
	}
	return 0
}

func (m *SceneResponse) GetSpaceMoveRet() *SpaceMoveRet {
	if x, ok := m.GetResponse().(*SceneResponse_SpaceMoveRet); ok {
		return x.SpaceMoveRet
	}
	return nil
}

func (m *SceneResponse) GetSpaceEnterRet() *SpaceEnterRet {
	if x, ok := m.GetResponse().(*SceneResponse_SpaceEnterRet); ok {
		return x.SpaceEnterRet
	}
	return nil
}

func (m *SceneResponse) GetSpaceLeaveRet() *SpaceLeaveRet {
	if x, ok := m.GetResponse().(*SceneResponse_SpaceLeaveRet); ok {
		return x.SpaceLeaveRet
	}
	return nil
}

func (m *SceneResponse) GetGetStateRet() *GetStateRet {
	if x, ok := m.GetResponse().(*SceneResponse_GetStateRet); ok {
		return x.GetStateRet
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SceneResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SceneResponse_OneofMarshaler, _SceneResponse_OneofUnmarshaler, _SceneResponse_OneofSizer, []interface{}{
		(*SceneResponse_SpaceMoveRet)(nil),
		(*SceneResponse_SpaceEnterRet)(nil),
		(*SceneResponse_SpaceLeaveRet)(nil),
		(*SceneResponse_GetStateRet)(nil),
	}
}

func _SceneResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SceneResponse)
	// response
	switch x := m.Response.(type) {
	case *SceneResponse_SpaceMoveRet:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SpaceMoveRet); err != nil {
			return err
		}
	case *SceneResponse_SpaceEnterRet:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SpaceEnterRet); err != nil {
			return err
		}
	case *SceneResponse_SpaceLeaveRet:
		_ = b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SpaceLeaveRet); err != nil {
			return err
		}
	case *SceneResponse_GetStateRet:
		_ = b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GetStateRet); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SceneResponse.Response has unexpected type %T", x)
	}
	return nil
}

func _SceneResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SceneResponse)
	switch tag {
	case 5: // response.spaceMoveRet
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SpaceMoveRet)
		err := b.DecodeMessage(msg)
		m.Response = &SceneResponse_SpaceMoveRet{msg}
		return true, err
	case 6: // response.spaceEnterRet
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SpaceEnterRet)
		err := b.DecodeMessage(msg)
		m.Response = &SceneResponse_SpaceEnterRet{msg}
		return true, err
	case 7: // response.spaceLeaveRet
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SpaceLeaveRet)
		err := b.DecodeMessage(msg)
		m.Response = &SceneResponse_SpaceLeaveRet{msg}
		return true, err
	case 8: // response.getStateRet
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GetStateRet)
		err := b.DecodeMessage(msg)
		m.Response = &SceneResponse_GetStateRet{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SceneResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SceneResponse)
	// response
	switch x := m.Response.(type) {
	case *SceneResponse_SpaceMoveRet:
		s := proto.Size(x.SpaceMoveRet)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SceneResponse_SpaceEnterRet:
		s := proto.Size(x.SpaceEnterRet)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SceneResponse_SpaceLeaveRet:
		s := proto.Size(x.SpaceLeaveRet)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SceneResponse_GetStateRet:
		s := proto.Size(x.GetStateRet)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// push
type ScenePush struct {
	SpacePush *SpacePush `protobuf:"bytes,1,opt,name=spacePush" json:"spacePush,omitempty"`
}

func (m *ScenePush) Reset()                    { *m = ScenePush{} }
func (m *ScenePush) String() string            { return proto.CompactTextString(m) }
func (*ScenePush) ProtoMessage()               {}
func (*ScenePush) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{2} }

func (m *ScenePush) GetSpacePush() *SpacePush {
	if m != nil {
		return m.SpacePush
	}
	return nil
}

func init() {
	proto.RegisterType((*SceneRequest)(nil), "scene.SceneRequest")
	proto.RegisterType((*SceneResponse)(nil), "scene.SceneResponse")
	proto.RegisterType((*ScenePush)(nil), "scene.ScenePush")
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptorProtocol) }

var fileDescriptorProtocol = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x18, 0x84, 0xd3, 0x4a, 0x9b, 0xf4, 0x4f, 0x5b, 0x65, 0x4f, 0x4b, 0x4f, 0xa5, 0xa7, 0x8a, 0x90,
	0x43, 0x45, 0xf1, 0x5c, 0x14, 0x22, 0x28, 0xc8, 0xf6, 0x01, 0x4a, 0x0c, 0x3f, 0x6d, 0x21, 0x66,
	0x63, 0x76, 0x93, 0x47, 0xf0, 0x09, 0x7c, 0x60, 0xd9, 0x3f, 0x9b, 0x4d, 0x72, 0xf1, 0xd4, 0xce,
	0xfc, 0xdf, 0x1c, 0x66, 0x36, 0xb0, 0x2c, 0x4a, 0xa9, 0x65, 0x2a, 0xb3, 0x88, 0xfe, 0xb0, 0x89,
	0x4a, 0x31, 0xc7, 0x55, 0x48, 0x3f, 0x8d, 0xb7, 0xf9, 0x19, 0xc3, 0xfc, 0x60, 0xb4, 0xc0, 0xef,
	0x0a, 0x95, 0x66, 0x1c, 0x7c, 0x85, 0x4a, 0x5d, 0x64, 0xce, 0x47, 0xeb, 0xd1, 0x76, 0x22, 0x5a,
	0xc9, 0x56, 0x10, 0xa4, 0xd9, 0x05, 0x73, 0xfd, 0xfa, 0xcc, 0xaf, 0xe8, 0xe4, 0x34, 0xbb, 0x83,
	0x99, 0x2a, 0x92, 0x14, 0xdf, 0x65, 0x8d, 0x7c, 0xb2, 0x1e, 0x6d, 0xc3, 0x5d, 0x18, 0x91, 0x73,
	0xfc, 0x92, 0x35, 0xc6, 0x9e, 0xe8, 0xee, 0x2c, 0x02, 0x20, 0xf1, 0x92, 0x6b, 0x2c, 0xf9, 0x94,
	0xe8, 0xb9, 0xa5, 0xd1, 0x78, 0xb1, 0x27, 0x7a, 0x84, 0xe3, 0xdf, 0x30, 0xa9, 0x91, 0xfb, 0x03,
	0x3e, 0x33, 0x9e, 0xe3, 0x89, 0x60, 0x5b, 0x08, 0x4e, 0xa8, 0x0f, 0x3a, 0xd1, 0xc8, 0x03, 0xa2,
	0x21, 0x3a, 0xa1, 0x3e, 0x2a, 0xe3, 0xc4, 0x9e, 0x70, 0xd7, 0xfd, 0x0c, 0xfc, 0xb2, 0xe9, 0xbd,
	0xf9, 0x1d, 0xc3, 0xc2, 0x0e, 0xa1, 0x0a, 0x99, 0x2b, 0xfc, 0x67, 0x89, 0x07, 0x98, 0xbb, 0x36,
	0x02, 0xb5, 0x2d, 0x7c, 0xdd, 0x2b, 0x7c, 0x2c, 0x51, 0xc7, 0x9e, 0x18, 0x60, 0xec, 0x09, 0x16,
	0x5d, 0x2b, 0x93, 0x6b, 0xaa, 0xdf, 0xf4, 0xab, 0xdb, 0xe0, 0x10, 0x74, 0x49, 0xea, 0x67, 0x92,
	0xfe, 0x20, 0x49, 0x23, 0x0c, 0x92, 0x2d, 0xc8, 0x76, 0x10, 0xb6, 0x6d, 0x4d, 0xae, 0x99, 0x63,
	0xd9, 0xcd, 0x61, 0x53, 0x7d, 0x68, 0x0f, 0x10, 0x94, 0x76, 0x84, 0xcd, 0x23, 0xcc, 0x68, 0x95,
	0x8f, 0x4a, 0x9d, 0xd9, 0xad, 0x7d, 0x65, 0x23, 0x68, 0x93, 0xee, 0x95, 0x8b, 0x4a, 0x9d, 0x45,
	0x77, 0xfd, 0x9c, 0xd2, 0xe7, 0x75, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xd6, 0x86, 0x7a,
	0x84, 0x02, 0x00, 0x00,
}
