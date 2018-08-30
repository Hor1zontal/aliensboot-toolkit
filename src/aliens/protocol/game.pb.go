// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: game.proto

/*
	Package protocol is a generated protocol buffer package.

	It is generated from these files:
		game.proto
		game_model.proto
		gate.proto
		passport.proto
		passport_model.proto
		protocol.proto
		resultcode.proto
		scene.proto
		scene_model.proto

	It has these top-level messages:
		GetUserInfo
		GetUserInfoRet
		LoginRole
		LoginRoleRet
		CreateRole
		CreateRoleRet
		RemoveRole
		RemoveRoleRet
		GameUser
		Role
		RoleInfo
		Equip
		Skill
		KickOut
		BindService
		LoginRegister
		LoginRegisterRet
		LoginLogin
		LoginLoginRet
		ChannelLogin
		ChannelLoginRet
		TokenLogin
		TokenLoginRet
		User
		Request
		Response
		Push
		SpaceEnter
		SpaceEnterRet
		SpaceMove
		SpaceMoveRet
		SpaceLeave
		SpaceLeaveRet
		GetState
		GetStateRet
		SpacePush
		Vector
		Entity
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// 登录游戏服务器
type GetUserInfo struct {
}

func (m *GetUserInfo) Reset()                    { *m = GetUserInfo{} }
func (m *GetUserInfo) String() string            { return proto.CompactTextString(m) }
func (*GetUserInfo) ProtoMessage()               {}
func (*GetUserInfo) Descriptor() ([]byte, []int) { return fileDescriptorGame, []int{0} }

// 登录游戏服务器返回
type GetUserInfoRet struct {
	User *GameUser `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *GetUserInfoRet) Reset()                    { *m = GetUserInfoRet{} }
func (m *GetUserInfoRet) String() string            { return proto.CompactTextString(m) }
func (*GetUserInfoRet) ProtoMessage()               {}
func (*GetUserInfoRet) Descriptor() ([]byte, []int) { return fileDescriptorGame, []int{1} }

func (m *GetUserInfoRet) GetUser() *GameUser {
	if m != nil {
		return m.User
	}
	return nil
}

// 角色登录
type LoginRole struct {
	RoleID int64 `protobuf:"varint,1,opt,name=roleID,proto3" json:"roleID,omitempty"`
}

func (m *LoginRole) Reset()                    { *m = LoginRole{} }
func (m *LoginRole) String() string            { return proto.CompactTextString(m) }
func (*LoginRole) ProtoMessage()               {}
func (*LoginRole) Descriptor() ([]byte, []int) { return fileDescriptorGame, []int{2} }

func (m *LoginRole) GetRoleID() int64 {
	if m != nil {
		return m.RoleID
	}
	return 0
}

type LoginRoleRet struct {
	RoleInfo   *RoleInfo `protobuf:"bytes,1,opt,name=roleInfo" json:"roleInfo,omitempty"`
	ServerTime int64     `protobuf:"varint,2,opt,name=serverTime,proto3" json:"serverTime,omitempty"`
}

func (m *LoginRoleRet) Reset()                    { *m = LoginRoleRet{} }
func (m *LoginRoleRet) String() string            { return proto.CompactTextString(m) }
func (*LoginRoleRet) ProtoMessage()               {}
func (*LoginRoleRet) Descriptor() ([]byte, []int) { return fileDescriptorGame, []int{3} }

func (m *LoginRoleRet) GetRoleInfo() *RoleInfo {
	if m != nil {
		return m.RoleInfo
	}
	return nil
}

func (m *LoginRoleRet) GetServerTime() int64 {
	if m != nil {
		return m.ServerTime
	}
	return 0
}

// 创建角色
type CreateRole struct {
	Role *Role `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
}

func (m *CreateRole) Reset()                    { *m = CreateRole{} }
func (m *CreateRole) String() string            { return proto.CompactTextString(m) }
func (*CreateRole) ProtoMessage()               {}
func (*CreateRole) Descriptor() ([]byte, []int) { return fileDescriptorGame, []int{4} }

func (m *CreateRole) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type CreateRoleRet struct {
	Role *Role `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
}

func (m *CreateRoleRet) Reset()                    { *m = CreateRoleRet{} }
func (m *CreateRoleRet) String() string            { return proto.CompactTextString(m) }
func (*CreateRoleRet) ProtoMessage()               {}
func (*CreateRoleRet) Descriptor() ([]byte, []int) { return fileDescriptorGame, []int{5} }

func (m *CreateRoleRet) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

// 删除角色
type RemoveRole struct {
	RoleID int64 `protobuf:"varint,1,opt,name=roleID,proto3" json:"roleID,omitempty"`
}

func (m *RemoveRole) Reset()                    { *m = RemoveRole{} }
func (m *RemoveRole) String() string            { return proto.CompactTextString(m) }
func (*RemoveRole) ProtoMessage()               {}
func (*RemoveRole) Descriptor() ([]byte, []int) { return fileDescriptorGame, []int{6} }

func (m *RemoveRole) GetRoleID() int64 {
	if m != nil {
		return m.RoleID
	}
	return 0
}

type RemoveRoleRet struct {
}

func (m *RemoveRoleRet) Reset()                    { *m = RemoveRoleRet{} }
func (m *RemoveRoleRet) String() string            { return proto.CompactTextString(m) }
func (*RemoveRoleRet) ProtoMessage()               {}
func (*RemoveRoleRet) Descriptor() ([]byte, []int) { return fileDescriptorGame, []int{7} }

func init() {
	proto.RegisterType((*GetUserInfo)(nil), "protocol.get_user_info")
	proto.RegisterType((*GetUserInfoRet)(nil), "protocol.get_user_info_ret")
	proto.RegisterType((*LoginRole)(nil), "protocol.login_role")
	proto.RegisterType((*LoginRoleRet)(nil), "protocol.login_role_ret")
	proto.RegisterType((*CreateRole)(nil), "protocol.create_role")
	proto.RegisterType((*CreateRoleRet)(nil), "protocol.create_role_ret")
	proto.RegisterType((*RemoveRole)(nil), "protocol.remove_role")
	proto.RegisterType((*RemoveRoleRet)(nil), "protocol.remove_role_ret")
}
func (m *GetUserInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetUserInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GetUserInfoRet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetUserInfoRet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.User != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGame(dAtA, i, uint64(m.User.Size()))
		n1, err := m.User.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *LoginRole) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginRole) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RoleID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGame(dAtA, i, uint64(m.RoleID))
	}
	return i, nil
}

func (m *LoginRoleRet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginRoleRet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RoleInfo != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGame(dAtA, i, uint64(m.RoleInfo.Size()))
		n2, err := m.RoleInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.ServerTime != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGame(dAtA, i, uint64(m.ServerTime))
	}
	return i, nil
}

func (m *CreateRole) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateRole) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Role != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGame(dAtA, i, uint64(m.Role.Size()))
		n3, err := m.Role.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *CreateRoleRet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateRoleRet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Role != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGame(dAtA, i, uint64(m.Role.Size()))
		n4, err := m.Role.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *RemoveRole) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveRole) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RoleID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGame(dAtA, i, uint64(m.RoleID))
	}
	return i, nil
}

func (m *RemoveRoleRet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveRoleRet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintGame(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GetUserInfo) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *GetUserInfoRet) Size() (n int) {
	var l int
	_ = l
	if m.User != nil {
		l = m.User.Size()
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func (m *LoginRole) Size() (n int) {
	var l int
	_ = l
	if m.RoleID != 0 {
		n += 1 + sovGame(uint64(m.RoleID))
	}
	return n
}

func (m *LoginRoleRet) Size() (n int) {
	var l int
	_ = l
	if m.RoleInfo != nil {
		l = m.RoleInfo.Size()
		n += 1 + l + sovGame(uint64(l))
	}
	if m.ServerTime != 0 {
		n += 1 + sovGame(uint64(m.ServerTime))
	}
	return n
}

func (m *CreateRole) Size() (n int) {
	var l int
	_ = l
	if m.Role != nil {
		l = m.Role.Size()
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func (m *CreateRoleRet) Size() (n int) {
	var l int
	_ = l
	if m.Role != nil {
		l = m.Role.Size()
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func (m *RemoveRole) Size() (n int) {
	var l int
	_ = l
	if m.RoleID != 0 {
		n += 1 + sovGame(uint64(m.RoleID))
	}
	return n
}

func (m *RemoveRoleRet) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovGame(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGame(x uint64) (n int) {
	return sovGame(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetUserInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: get_user_info: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: get_user_info: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
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
func (m *GetUserInfoRet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: get_user_info_ret: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: get_user_info_ret: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.User == nil {
				m.User = &GameUser{}
			}
			if err := m.User.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
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
func (m *LoginRole) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: login_role: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: login_role: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleID", wireType)
			}
			m.RoleID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RoleID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
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
func (m *LoginRoleRet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: login_role_ret: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: login_role_ret: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RoleInfo == nil {
				m.RoleInfo = &RoleInfo{}
			}
			if err := m.RoleInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerTime", wireType)
			}
			m.ServerTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServerTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
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
func (m *CreateRole) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: create_role: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: create_role: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Role == nil {
				m.Role = &Role{}
			}
			if err := m.Role.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
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
func (m *CreateRoleRet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: create_role_ret: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: create_role_ret: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Role == nil {
				m.Role = &Role{}
			}
			if err := m.Role.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
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
func (m *RemoveRole) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: remove_role: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: remove_role: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleID", wireType)
			}
			m.RoleID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RoleID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
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
func (m *RemoveRoleRet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: remove_role_ret: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: remove_role_ret: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
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
func skipGame(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGame
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
					return 0, ErrIntOverflowGame
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
					return 0, ErrIntOverflowGame
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
				return 0, ErrInvalidLengthGame
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGame
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
				next, err := skipGame(dAtA[start:])
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
	ErrInvalidLengthGame = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGame   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("game.proto", fileDescriptorGame) }

var fileDescriptorGame = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0x8d, 0x96, 0x52, 0x26, 0xd8, 0xb4, 0x7b, 0x90, 0xe2, 0x61, 0x91, 0x45, 0xc5, 0x53,
	0x40, 0xc5, 0x93, 0x37, 0x11, 0xc4, 0xeb, 0xa2, 0xe7, 0x35, 0xd6, 0x69, 0x08, 0x64, 0x33, 0x32,
	0x89, 0x7d, 0x16, 0x1f, 0xc9, 0xa3, 0x8f, 0x20, 0xf1, 0x45, 0x64, 0xa7, 0x51, 0x53, 0x04, 0xe9,
	0x69, 0xd8, 0xff, 0xc7, 0x6f, 0x66, 0x01, 0xf2, 0xcc, 0x63, 0xfa, 0xcc, 0xd4, 0x90, 0x1a, 0xc9,
	0x98, 0x53, 0xb9, 0x3f, 0x09, 0xaa, 0xf3, 0xf4, 0x84, 0xe5, 0xca, 0x33, 0x09, 0xec, 0xe6, 0xd8,
	0xb8, 0x97, 0x1a, 0xd9, 0x15, 0xd5, 0x82, 0xcc, 0x25, 0x4c, 0xd7, 0x04, 0xc7, 0xd8, 0xa8, 0x63,
	0x18, 0x04, 0x61, 0x16, 0x1d, 0x44, 0x27, 0xf1, 0x99, 0x4a, 0xbf, 0x81, 0xe9, 0x4d, 0xe6, 0xf1,
	0xbe, 0x46, 0xb6, 0xe2, 0x9b, 0x43, 0x80, 0x92, 0xf2, 0xa2, 0x72, 0x4c, 0x25, 0xaa, 0x3d, 0x18,
	0x86, 0x79, 0x7b, 0x2d, 0xbd, 0x1d, 0xdb, 0xbd, 0xcc, 0x03, 0x8c, 0x7f, 0x53, 0xc2, 0x4f, 0x61,
	0x24, 0x5e, 0xb5, 0xa0, 0xbf, 0x3b, 0x6c, 0xe7, 0xd8, 0x9f, 0x8c, 0xd2, 0x00, 0x35, 0xf2, 0x12,
	0xf9, 0xae, 0xf0, 0x38, 0xdb, 0x16, 0x7a, 0x4f, 0x31, 0xa7, 0x10, 0xcf, 0x19, 0xb3, 0x06, 0x57,
	0x87, 0x18, 0x18, 0x84, 0xd9, 0xa1, 0xc7, 0xeb, 0x68, 0x2b, 0x9e, 0xb9, 0x80, 0xa4, 0x57, 0x91,
	0xab, 0x36, 0xa9, 0x1d, 0x41, 0xcc, 0xe8, 0x69, 0x89, 0xff, 0x7f, 0x79, 0x0a, 0x49, 0x2f, 0x16,
	0xe8, 0x57, 0x93, 0xb7, 0x56, 0x47, 0xef, 0xad, 0x8e, 0x3e, 0x5a, 0x1d, 0xbd, 0x7e, 0xea, 0xad,
	0xc7, 0xa1, 0x2c, 0x38, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x3c, 0xf7, 0x27, 0xbc, 0x01,
	0x00, 0x00,
}
