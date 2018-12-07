// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scene_interface.proto

package protocol

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 内部接口玩家登录场景
type LoginScene struct {
	AuthID  int64  `protobuf:"varint,1,opt,name=authID,proto3" json:"authID,omitempty"`
	GateID  string `protobuf:"bytes,2,opt,name=gateID,proto3" json:"gateID,omitempty"`
	SpaceID string `protobuf:"bytes,3,opt,name=spaceID,proto3" json:"spaceID,omitempty"`
}

func (m *LoginScene) Reset()                    { *m = LoginScene{} }
func (m *LoginScene) String() string            { return proto.CompactTextString(m) }
func (*LoginScene) ProtoMessage()               {}
func (*LoginScene) Descriptor() ([]byte, []int) { return fileDescriptorSceneInterface, []int{0} }

func (m *LoginScene) GetAuthID() int64 {
	if m != nil {
		return m.AuthID
	}
	return 0
}

func (m *LoginScene) GetGateID() string {
	if m != nil {
		return m.GateID
	}
	return ""
}

func (m *LoginScene) GetSpaceID() string {
	if m != nil {
		return m.SpaceID
	}
	return ""
}

//
type LoginSceneRet struct {
	Entity *Entity `protobuf:"bytes,1,opt,name=entity" json:"entity,omitempty"`
}

func (m *LoginSceneRet) Reset()                    { *m = LoginSceneRet{} }
func (m *LoginSceneRet) String() string            { return proto.CompactTextString(m) }
func (*LoginSceneRet) ProtoMessage()               {}
func (*LoginSceneRet) Descriptor() ([]byte, []int) { return fileDescriptorSceneInterface, []int{1} }

func (m *LoginSceneRet) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

// 调用实体方法
type EntityCall struct {
	EntityID string   `protobuf:"bytes,1,opt,name=entityID,proto3" json:"entityID,omitempty"`
	Method   string   `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Args     [][]byte `protobuf:"bytes,3,rep,name=args" json:"args,omitempty"`
}

func (m *EntityCall) Reset()                    { *m = EntityCall{} }
func (m *EntityCall) String() string            { return proto.CompactTextString(m) }
func (*EntityCall) ProtoMessage()               {}
func (*EntityCall) Descriptor() ([]byte, []int) { return fileDescriptorSceneInterface, []int{2} }

func (m *EntityCall) GetEntityID() string {
	if m != nil {
		return m.EntityID
	}
	return ""
}

func (m *EntityCall) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *EntityCall) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

// 场景数据变更推送
type EntityPush struct {
	Neighbors       []*Entity `protobuf:"bytes,1,rep,name=neighbors" json:"neighbors,omitempty"`
	DistoryEntities []string  `protobuf:"bytes,2,rep,name=distoryEntities" json:"distoryEntities,omitempty"`
}

func (m *EntityPush) Reset()                    { *m = EntityPush{} }
func (m *EntityPush) String() string            { return proto.CompactTextString(m) }
func (*EntityPush) ProtoMessage()               {}
func (*EntityPush) Descriptor() ([]byte, []int) { return fileDescriptorSceneInterface, []int{3} }

func (m *EntityPush) GetNeighbors() []*Entity {
	if m != nil {
		return m.Neighbors
	}
	return nil
}

func (m *EntityPush) GetDistoryEntities() []string {
	if m != nil {
		return m.DistoryEntities
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginScene)(nil), "protocol.LoginScene")
	proto.RegisterType((*LoginSceneRet)(nil), "protocol.LoginSceneRet")
	proto.RegisterType((*EntityCall)(nil), "protocol.entityCall")
	proto.RegisterType((*EntityPush)(nil), "protocol.entityPush")
}
func (m *LoginScene) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginScene) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AuthID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSceneInterface(dAtA, i, uint64(m.AuthID))
	}
	if len(m.GateID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSceneInterface(dAtA, i, uint64(len(m.GateID)))
		i += copy(dAtA[i:], m.GateID)
	}
	if len(m.SpaceID) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSceneInterface(dAtA, i, uint64(len(m.SpaceID)))
		i += copy(dAtA[i:], m.SpaceID)
	}
	return i, nil
}

func (m *LoginSceneRet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginSceneRet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Entity != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSceneInterface(dAtA, i, uint64(m.Entity.Size()))
		n1, err := m.Entity.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *EntityCall) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntityCall) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.EntityID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSceneInterface(dAtA, i, uint64(len(m.EntityID)))
		i += copy(dAtA[i:], m.EntityID)
	}
	if len(m.Method) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSceneInterface(dAtA, i, uint64(len(m.Method)))
		i += copy(dAtA[i:], m.Method)
	}
	if len(m.Args) > 0 {
		for _, b := range m.Args {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintSceneInterface(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	return i, nil
}

func (m *EntityPush) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntityPush) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Neighbors) > 0 {
		for _, msg := range m.Neighbors {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSceneInterface(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.DistoryEntities) > 0 {
		for _, s := range m.DistoryEntities {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintSceneInterface(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LoginScene) Size() (n int) {
	var l int
	_ = l
	if m.AuthID != 0 {
		n += 1 + sovSceneInterface(uint64(m.AuthID))
	}
	l = len(m.GateID)
	if l > 0 {
		n += 1 + l + sovSceneInterface(uint64(l))
	}
	l = len(m.SpaceID)
	if l > 0 {
		n += 1 + l + sovSceneInterface(uint64(l))
	}
	return n
}

func (m *LoginSceneRet) Size() (n int) {
	var l int
	_ = l
	if m.Entity != nil {
		l = m.Entity.Size()
		n += 1 + l + sovSceneInterface(uint64(l))
	}
	return n
}

func (m *EntityCall) Size() (n int) {
	var l int
	_ = l
	l = len(m.EntityID)
	if l > 0 {
		n += 1 + l + sovSceneInterface(uint64(l))
	}
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovSceneInterface(uint64(l))
	}
	if len(m.Args) > 0 {
		for _, b := range m.Args {
			l = len(b)
			n += 1 + l + sovSceneInterface(uint64(l))
		}
	}
	return n
}

func (m *EntityPush) Size() (n int) {
	var l int
	_ = l
	if len(m.Neighbors) > 0 {
		for _, e := range m.Neighbors {
			l = e.Size()
			n += 1 + l + sovSceneInterface(uint64(l))
		}
	}
	if len(m.DistoryEntities) > 0 {
		for _, s := range m.DistoryEntities {
			l = len(s)
			n += 1 + l + sovSceneInterface(uint64(l))
		}
	}
	return n
}

func sovSceneInterface(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSceneInterface(x uint64) (n int) {
	return sovSceneInterface(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LoginScene) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSceneInterface
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
			return fmt.Errorf("proto: LoginScene: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginScene: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthID", wireType)
			}
			m.AuthID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSceneInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuthID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GateID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSceneInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSceneInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GateID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSceneInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSceneInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpaceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSceneInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSceneInterface
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
func (m *LoginSceneRet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSceneInterface
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
			return fmt.Errorf("proto: LoginSceneRet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginSceneRet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSceneInterface
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
				return ErrInvalidLengthSceneInterface
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Entity == nil {
				m.Entity = &Entity{}
			}
			if err := m.Entity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSceneInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSceneInterface
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
func (m *EntityCall) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSceneInterface
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
			return fmt.Errorf("proto: entityCall: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: entityCall: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSceneInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSceneInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EntityID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSceneInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSceneInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSceneInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSceneInterface
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Args = append(m.Args, make([]byte, postIndex-iNdEx))
			copy(m.Args[len(m.Args)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSceneInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSceneInterface
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
func (m *EntityPush) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSceneInterface
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
			return fmt.Errorf("proto: entityPush: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: entityPush: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Neighbors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSceneInterface
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
				return ErrInvalidLengthSceneInterface
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Neighbors = append(m.Neighbors, &Entity{})
			if err := m.Neighbors[len(m.Neighbors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistoryEntities", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSceneInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSceneInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistoryEntities = append(m.DistoryEntities, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSceneInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSceneInterface
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
func skipSceneInterface(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSceneInterface
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
					return 0, ErrIntOverflowSceneInterface
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
					return 0, ErrIntOverflowSceneInterface
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
				return 0, ErrInvalidLengthSceneInterface
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSceneInterface
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
				next, err := skipSceneInterface(dAtA[start:])
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
	ErrInvalidLengthSceneInterface = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSceneInterface   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("scene_interface.proto", fileDescriptorSceneInterface) }

var fileDescriptorSceneInterface = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8e, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xc9, 0x22, 0x73, 0xfd, 0x9c, 0xa8, 0x01, 0x25, 0xec, 0x50, 0x4a, 0x4f, 0x39, 0xf5,
	0xa0, 0x27, 0xaf, 0x3a, 0x0f, 0x03, 0x0f, 0x12, 0xc5, 0xab, 0x64, 0xed, 0xb7, 0x36, 0xd0, 0x35,
	0xa3, 0xc9, 0x0e, 0xfb, 0x87, 0x1e, 0xfd, 0x09, 0xd2, 0x5f, 0x22, 0x4d, 0x5a, 0x07, 0xe2, 0x29,
	0x79, 0xde, 0xe7, 0xe3, 0xe5, 0x85, 0x6b, 0x9b, 0x63, 0x83, 0x1f, 0xba, 0x71, 0xd8, 0x6e, 0x54,
	0x8e, 0xd9, 0xae, 0x35, 0xce, 0xb0, 0x99, 0x7f, 0x72, 0x53, 0x2f, 0xae, 0xc2, 0xc1, 0xd6, 0x14,
	0x58, 0x07, 0x99, 0xbe, 0x03, 0x3c, 0x9b, 0x52, 0x37, 0xaf, 0xbd, 0x61, 0x37, 0x30, 0x55, 0x7b,
	0x57, 0xad, 0x96, 0x9c, 0x24, 0x44, 0x50, 0x39, 0x50, 0x9f, 0x97, 0xca, 0xe1, 0x6a, 0xc9, 0x27,
	0x09, 0x11, 0x91, 0x1c, 0x88, 0x71, 0x38, 0xb5, 0x3b, 0x95, 0xf7, 0x82, 0x7a, 0x31, 0x62, 0x7a,
	0x0f, 0xe7, 0xc7, 0x5e, 0x89, 0x8e, 0x09, 0x98, 0x62, 0xe3, 0xb4, 0x3b, 0xf8, 0xea, 0xb3, 0xdb,
	0xcb, 0x6c, 0x9c, 0x95, 0x3d, 0xf9, 0x5c, 0x0e, 0x3e, 0x7d, 0x03, 0x08, 0xbf, 0x47, 0x55, 0xd7,
	0x6c, 0x01, 0xb3, 0x40, 0xc3, 0xa8, 0x48, 0xfe, 0x72, 0x3f, 0x6b, 0x8b, 0xae, 0x32, 0xc5, 0x38,
	0x2b, 0x10, 0x63, 0x70, 0xa2, 0xda, 0xd2, 0x72, 0x9a, 0x50, 0x31, 0x97, 0xfe, 0x9f, 0x6e, 0xc6,
	0xd6, 0x97, 0xbd, 0xad, 0x58, 0x06, 0x51, 0x83, 0xba, 0xac, 0xd6, 0xa6, 0xb5, 0x9c, 0x24, 0xf4,
	0xdf, 0x41, 0xc7, 0x13, 0x26, 0xe0, 0xa2, 0xd0, 0xd6, 0x99, 0xf6, 0xe0, 0x9d, 0x46, 0xcb, 0x27,
	0x09, 0x15, 0x91, 0xfc, 0x1b, 0x3f, 0xcc, 0x3f, 0xbb, 0x98, 0x7c, 0x75, 0x31, 0xf9, 0xee, 0x62,
	0xb2, 0x9e, 0xfa, 0xce, 0xbb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x96, 0x4f, 0x0a, 0x33, 0x9b,
	0x01, 0x00, 0x00,
}