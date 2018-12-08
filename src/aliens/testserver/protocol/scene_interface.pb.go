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

// entity
type MigrateOut struct {
	EntityID  string `protobuf:"bytes,1,opt,name=entityID,proto3" json:"entityID,omitempty"`
	ToSpaceID string `protobuf:"bytes,2,opt,name=toSpaceID,proto3" json:"toSpaceID,omitempty"`
}

func (m *MigrateOut) Reset()                    { *m = MigrateOut{} }
func (m *MigrateOut) String() string            { return proto.CompactTextString(m) }
func (*MigrateOut) ProtoMessage()               {}
func (*MigrateOut) Descriptor() ([]byte, []int) { return fileDescriptorSceneInterface, []int{2} }

func (m *MigrateOut) GetEntityID() string {
	if m != nil {
		return m.EntityID
	}
	return ""
}

func (m *MigrateOut) GetToSpaceID() string {
	if m != nil {
		return m.ToSpaceID
	}
	return ""
}

// entity
type MigrateIn struct {
	SpaceID  string `protobuf:"bytes,1,opt,name=spaceID,proto3" json:"spaceID,omitempty"`
	EntityID string `protobuf:"bytes,2,opt,name=entityID,proto3" json:"entityID,omitempty"`
	Data     []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *MigrateIn) Reset()                    { *m = MigrateIn{} }
func (m *MigrateIn) String() string            { return proto.CompactTextString(m) }
func (*MigrateIn) ProtoMessage()               {}
func (*MigrateIn) Descriptor() ([]byte, []int) { return fileDescriptorSceneInterface, []int{3} }

func (m *MigrateIn) GetSpaceID() string {
	if m != nil {
		return m.SpaceID
	}
	return ""
}

func (m *MigrateIn) GetEntityID() string {
	if m != nil {
		return m.EntityID
	}
	return ""
}

func (m *MigrateIn) GetData() []byte {
	if m != nil {
		return m.Data
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
func (*EntityCall) Descriptor() ([]byte, []int) { return fileDescriptorSceneInterface, []int{4} }

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
func (*EntityPush) Descriptor() ([]byte, []int) { return fileDescriptorSceneInterface, []int{5} }

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
	proto.RegisterType((*MigrateOut)(nil), "protocol.MigrateOut")
	proto.RegisterType((*MigrateIn)(nil), "protocol.MigrateIn")
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

func (m *MigrateOut) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MigrateOut) MarshalTo(dAtA []byte) (int, error) {
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
	if len(m.ToSpaceID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSceneInterface(dAtA, i, uint64(len(m.ToSpaceID)))
		i += copy(dAtA[i:], m.ToSpaceID)
	}
	return i, nil
}

func (m *MigrateIn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MigrateIn) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.SpaceID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSceneInterface(dAtA, i, uint64(len(m.SpaceID)))
		i += copy(dAtA[i:], m.SpaceID)
	}
	if len(m.EntityID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSceneInterface(dAtA, i, uint64(len(m.EntityID)))
		i += copy(dAtA[i:], m.EntityID)
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSceneInterface(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
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

func (m *MigrateOut) Size() (n int) {
	var l int
	_ = l
	l = len(m.EntityID)
	if l > 0 {
		n += 1 + l + sovSceneInterface(uint64(l))
	}
	l = len(m.ToSpaceID)
	if l > 0 {
		n += 1 + l + sovSceneInterface(uint64(l))
	}
	return n
}

func (m *MigrateIn) Size() (n int) {
	var l int
	_ = l
	l = len(m.SpaceID)
	if l > 0 {
		n += 1 + l + sovSceneInterface(uint64(l))
	}
	l = len(m.EntityID)
	if l > 0 {
		n += 1 + l + sovSceneInterface(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
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
func (m *MigrateOut) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MigrateOut: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MigrateOut: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field ToSpaceID", wireType)
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
			m.ToSpaceID = string(dAtA[iNdEx:postIndex])
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
func (m *MigrateIn) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MigrateIn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MigrateIn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
		case 2:
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
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
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
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
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xd9, 0x46, 0x62, 0x33, 0x56, 0xd4, 0x05, 0x4b, 0x28, 0x52, 0x42, 0x4e, 0x39, 0xe5,
	0xa0, 0x27, 0xaf, 0x5a, 0x85, 0x82, 0xa2, 0x6c, 0xd5, 0xab, 0x6c, 0x93, 0x6d, 0xb2, 0x90, 0x66,
	0x4b, 0x76, 0x7a, 0xe8, 0x3f, 0xf4, 0xe8, 0x4f, 0x90, 0xfe, 0x12, 0xc9, 0x66, 0x63, 0x6c, 0x11,
	0x4f, 0xd9, 0xf7, 0xde, 0xf0, 0xe6, 0xcb, 0xc0, 0xb9, 0x4e, 0x44, 0x29, 0xde, 0x65, 0x89, 0xa2,
	0x5a, 0xf0, 0x44, 0xc4, 0xab, 0x4a, 0xa1, 0xa2, 0x7d, 0xf3, 0x49, 0x54, 0x31, 0x3a, 0x6b, 0x06,
	0x96, 0x2a, 0x15, 0x45, 0x13, 0x86, 0x6f, 0x00, 0x0f, 0x2a, 0x93, 0xe5, 0xac, 0x4e, 0xe8, 0x10,
	0x5c, 0xbe, 0xc6, 0x7c, 0x3a, 0xf1, 0x49, 0x40, 0x22, 0x87, 0x59, 0x55, 0xfb, 0x19, 0x47, 0x31,
	0x9d, 0xf8, 0xbd, 0x80, 0x44, 0x1e, 0xb3, 0x8a, 0xfa, 0x70, 0xa8, 0x57, 0x3c, 0xa9, 0x03, 0xc7,
	0x04, 0xad, 0x0c, 0xaf, 0xe1, 0xb8, 0xeb, 0x65, 0x02, 0x69, 0x04, 0xae, 0x28, 0x51, 0xe2, 0xc6,
	0x54, 0x1f, 0x5d, 0x9e, 0xc6, 0x2d, 0x56, 0x7c, 0x67, 0x7c, 0x66, 0xf3, 0xf0, 0x1e, 0xe0, 0x51,
	0x66, 0x15, 0x47, 0xf1, 0xb4, 0x46, 0x3a, 0x82, 0x7e, 0xe3, 0x5b, 0x28, 0x8f, 0xfd, 0x68, 0x7a,
	0x01, 0x1e, 0xaa, 0x99, 0x05, 0x68, 0xc8, 0x3a, 0x23, 0x7c, 0x05, 0xcf, 0xf6, 0x4c, 0xcb, 0xdf,
	0xa4, 0x64, 0x87, 0x74, 0x67, 0x41, 0x6f, 0x6f, 0x01, 0x85, 0x83, 0x94, 0x23, 0x37, 0x3f, 0x37,
	0x60, 0xe6, 0x1d, 0xbe, 0x00, 0x34, 0xf9, 0x2d, 0x2f, 0x8a, 0x7f, 0xf1, 0x86, 0xe0, 0x2e, 0x05,
	0xe6, 0x2a, 0x6d, 0xaf, 0xd6, 0xa8, 0xba, 0x95, 0x57, 0x99, 0xf6, 0x9d, 0xc0, 0xa9, 0x5b, 0xeb,
	0x77, 0xb8, 0x68, 0x5b, 0x9f, 0xd7, 0x3a, 0xa7, 0x31, 0x78, 0xa5, 0x90, 0x59, 0x3e, 0x57, 0x95,
	0xf6, 0x49, 0xe0, 0xfc, 0x79, 0xaf, 0x6e, 0x84, 0x46, 0x70, 0x92, 0x4a, 0x8d, 0xaa, 0xda, 0x98,
	0x4c, 0x0a, 0xed, 0xf7, 0x02, 0x27, 0xf2, 0xd8, 0xbe, 0x7d, 0x33, 0xf8, 0xd8, 0x8e, 0xc9, 0xe7,
	0x76, 0x4c, 0xbe, 0xb6, 0x63, 0x32, 0x77, 0x4d, 0xe7, 0xd5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x51, 0x91, 0x10, 0xab, 0x3a, 0x02, 0x00, 0x00,
}
