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
	Session int32 `protobuf:"varint,1,opt,name=session,proto3" json:"session,omitempty"`
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
	//	*SceneResponse_Exception
	//	*SceneResponse_SpaceMoveRet
	//	*SceneResponse_SpaceEnterRet
	//	*SceneResponse_SpaceLeaveRet
	//	*SceneResponse_GetStateRet
	Response  isSceneResponse_Response `protobuf_oneof:"response"`
	ScenePush string                   `protobuf:"bytes,500,opt,name=scenePush,proto3" json:"scenePush,omitempty"`
	SpacePush *SpacePush               `protobuf:"bytes,501,opt,name=spacePush" json:"spacePush,omitempty"`
}

func (m *SceneResponse) Reset()                    { *m = SceneResponse{} }
func (m *SceneResponse) String() string            { return proto.CompactTextString(m) }
func (*SceneResponse) ProtoMessage()               {}
func (*SceneResponse) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{1} }

type isSceneResponse_Response interface {
	isSceneResponse_Response()
}

type SceneResponse_Exception struct {
	Exception uint32 `protobuf:"varint,2,opt,name=exception,proto3,oneof"`
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

func (*SceneResponse_Exception) isSceneResponse_Response()     {}
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

func (m *SceneResponse) GetException() uint32 {
	if x, ok := m.GetResponse().(*SceneResponse_Exception); ok {
		return x.Exception
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

func (m *SceneResponse) GetScenePush() string {
	if m != nil {
		return m.ScenePush
	}
	return ""
}

func (m *SceneResponse) GetSpacePush() *SpacePush {
	if m != nil {
		return m.SpacePush
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SceneResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SceneResponse_OneofMarshaler, _SceneResponse_OneofUnmarshaler, _SceneResponse_OneofSizer, []interface{}{
		(*SceneResponse_Exception)(nil),
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
	case *SceneResponse_Exception:
		_ = b.EncodeVarint(2<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Exception))
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
	case 2: // response.exception
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Response = &SceneResponse_Exception{uint32(x)}
		return true, err
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
	case *SceneResponse_Exception:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Exception))
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

func init() {
	proto.RegisterType((*SceneRequest)(nil), "scene.SceneRequest")
	proto.RegisterType((*SceneResponse)(nil), "scene.SceneResponse")
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptorProtocol) }

var fileDescriptorProtocol = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x4b, 0xfb, 0x40,
	0x14, 0xc4, 0x93, 0x96, 0x36, 0xcd, 0x4b, 0xdb, 0xff, 0x9f, 0x3d, 0x2d, 0x82, 0x12, 0x7a, 0x0a,
	0x0a, 0x39, 0x54, 0x04, 0xcf, 0x05, 0x21, 0x07, 0x05, 0xd9, 0x7e, 0x80, 0x52, 0xc3, 0xa3, 0x15,
	0x6a, 0x36, 0x66, 0xb7, 0xc5, 0xcf, 0xec, 0x59, 0xcf, 0xca, 0xbe, 0x6c, 0x37, 0xd9, 0x8b, 0xa7,
	0x64, 0x67, 0x7e, 0x73, 0x98, 0xe1, 0xc1, 0xbc, 0x6e, 0xa4, 0x96, 0xa5, 0x3c, 0xe4, 0xf4, 0xc3,
	0x46, 0xaa, 0xc4, 0x0a, 0x2f, 0x12, 0xfa, 0xb4, 0xda, 0xe2, 0x33, 0x84, 0xe9, 0xda, 0xbc, 0x05,
	0xbe, 0x1f, 0x51, 0x69, 0xc6, 0x21, 0x52, 0xa8, 0xd4, 0xab, 0xac, 0x78, 0x98, 0x86, 0xd9, 0x48,
	0x9c, 0x9f, 0xec, 0x06, 0x62, 0x55, 0x6f, 0x4b, 0x7c, 0x92, 0x27, 0xe4, 0xa3, 0x34, 0xcc, 0x92,
	0x65, 0x92, 0x93, 0xb2, 0x79, 0x93, 0x27, 0x2c, 0x02, 0xd1, 0xf9, 0x2c, 0x07, 0xa0, 0xc7, 0x43,
	0xa5, 0xb1, 0xe1, 0x63, 0xa2, 0xa7, 0x96, 0x46, 0xa3, 0x15, 0x81, 0xe8, 0x11, 0x8e, 0x7f, 0xc4,
	0xed, 0x09, 0x79, 0xe4, 0xf1, 0x07, 0xa3, 0x39, 0x9e, 0x08, 0x96, 0xc1, 0x64, 0x87, 0x7a, 0xad,
	0xb7, 0x1a, 0xf9, 0x84, 0x68, 0xc8, 0x77, 0xa8, 0x37, 0xca, 0x28, 0x45, 0x20, 0x9c, 0xbb, 0x8a,
	0x21, 0x6a, 0xda, 0x6e, 0x8b, 0x9f, 0x01, 0xcc, 0x6c, 0x59, 0x55, 0xcb, 0x4a, 0xe1, 0x1f, 0x6d,
	0xaf, 0x20, 0xc6, 0x8f, 0x12, 0x6b, 0x6d, 0xbc, 0x41, 0x1a, 0x66, 0x33, 0x53, 0xd0, 0x49, 0xec,
	0x0e, 0xa6, 0xae, 0xad, 0x40, 0x6d, 0x07, 0xf9, 0xd7, 0x1b, 0x64, 0xd3, 0xa0, 0x2e, 0x02, 0xe1,
	0x61, 0xec, 0x1e, 0x66, 0x5d, 0x6b, 0x93, 0x6b, 0xa7, 0xf9, 0xdf, 0x9f, 0xc6, 0x06, 0x7d, 0xd0,
	0x25, 0xa9, 0xbf, 0x49, 0x46, 0x5e, 0x92, 0x46, 0xf2, 0x92, 0x67, 0x90, 0x2d, 0x21, 0x39, 0xaf,
	0x61, 0x72, 0xed, 0x5c, 0xf3, 0x6e, 0x2e, 0x9b, 0xea, 0x43, 0xec, 0x12, 0x62, 0x3a, 0x93, 0xe7,
	0xa3, 0xda, 0xf3, 0xaf, 0x61, 0x1a, 0x66, 0xb1, 0xe8, 0x14, 0x76, 0x6d, 0x6f, 0x81, 0xec, 0xef,
	0xa1, 0x77, 0x0c, 0xf5, 0x51, 0xed, 0x45, 0x67, 0xaf, 0x00, 0x26, 0x8d, 0xdd, 0xfb, 0x65, 0x4c,
	0x57, 0x77, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x6f, 0x61, 0x53, 0x9b, 0x02, 0x00, 0x00,
}
