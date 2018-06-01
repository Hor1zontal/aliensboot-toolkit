// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model.proto

/*
	Package passport is a generated protocol buffer package.

	It is generated from these files:
		model.proto
		passport.proto
		protocol.proto

	It has these top-level messages:
		User
		LoginRegister
		LoginRegisterRet
		LoginLogin
		LoginLoginRet
		ChannelLogin
		ChannelLoginRet
		NewInterface
		NewInterfaceRet
		PassportRequest
		PassportResponse
*/
package passport

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import time "time"

import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id         int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id" gorm:"AUTO_INCREMENT"`
	Username   string    `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty" bson:"username" rorm:"uname"  unique:"true"`
	Password   string    `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" bson:"password" rorm:"pwd"`
	Salt       string    `protobuf:"bytes,4,opt,name=salt,proto3" json:"salt,omitempty" bson:"salt" rorm:"salt"`
	Channeluid string    `protobuf:"bytes,5,opt,name=channeluid,proto3" json:"channeluid,omitempty" bson:"cuid" rorm:"cuid"`
	Channel    string    `protobuf:"bytes,6,opt,name=channel,proto3" json:"channel,omitempty" bson:"channel" rorm:"channel"`
	Avatar     string    `protobuf:"bytes,7,opt,name=avatar,proto3" json:"avatar,omitempty" bson:"avatar" rorm:"avatar"`
	Mobile     string    `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile,omitempty" bson:"mobile"`
	Openid     string    `protobuf:"bytes,9,opt,name=openid,proto3" json:"openid,omitempty" bson:"openid"`
	Ip         string    `protobuf:"bytes,10,opt,name=ip,proto3" json:"ip,omitempty" bson:"ip"`
	Status     int32     `protobuf:"varint,11,opt,name=status,proto3" json:"status,omitempty" bson:"status" rorm:"status"`
	RegTime    time.Time `protobuf:"bytes,12,opt,name=regTime,stdtime" json:"regTime" bson:"regtime"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorModel, []int{0} }

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetSalt() string {
	if m != nil {
		return m.Salt
	}
	return ""
}

func (m *User) GetChanneluid() string {
	if m != nil {
		return m.Channeluid
	}
	return ""
}

func (m *User) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *User) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *User) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

func (m *User) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *User) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *User) GetRegTime() time.Time {
	if m != nil {
		return m.RegTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*User)(nil), "passport.User")
}
func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintModel(dAtA, i, uint64(m.Id))
	}
	if len(m.Username) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintModel(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintModel(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	if len(m.Salt) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintModel(dAtA, i, uint64(len(m.Salt)))
		i += copy(dAtA[i:], m.Salt)
	}
	if len(m.Channeluid) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintModel(dAtA, i, uint64(len(m.Channeluid)))
		i += copy(dAtA[i:], m.Channeluid)
	}
	if len(m.Channel) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintModel(dAtA, i, uint64(len(m.Channel)))
		i += copy(dAtA[i:], m.Channel)
	}
	if len(m.Avatar) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintModel(dAtA, i, uint64(len(m.Avatar)))
		i += copy(dAtA[i:], m.Avatar)
	}
	if len(m.Mobile) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintModel(dAtA, i, uint64(len(m.Mobile)))
		i += copy(dAtA[i:], m.Mobile)
	}
	if len(m.Openid) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintModel(dAtA, i, uint64(len(m.Openid)))
		i += copy(dAtA[i:], m.Openid)
	}
	if len(m.Ip) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintModel(dAtA, i, uint64(len(m.Ip)))
		i += copy(dAtA[i:], m.Ip)
	}
	if m.Status != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintModel(dAtA, i, uint64(m.Status))
	}
	dAtA[i] = 0x62
	i++
	i = encodeVarintModel(dAtA, i, uint64(types.SizeOfStdTime(m.RegTime)))
	n1, err := types.StdTimeMarshalTo(m.RegTime, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func encodeVarintModel(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *User) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovModel(uint64(m.Id))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.Salt)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.Channeluid)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.Avatar)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.Mobile)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.Openid)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.Ip)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovModel(uint64(m.Status))
	}
	l = types.SizeOfStdTime(m.RegTime)
	n += 1 + l + sovModel(uint64(l))
	return n
}

func sovModel(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozModel(x uint64) (n int) {
	return sovModel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModel
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Salt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channeluid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channeluid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Avatar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Avatar = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mobile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mobile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Openid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Openid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ip = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.RegTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthModel
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
func skipModel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModel
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
					return 0, ErrIntOverflowModel
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
					return 0, ErrIntOverflowModel
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
				return 0, ErrInvalidLengthModel
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowModel
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
				next, err := skipModel(dAtA[start:])
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
	ErrInvalidLengthModel = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModel   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("model.proto", fileDescriptorModel) }

var fileDescriptorModel = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcf, 0x6e, 0xd3, 0x30,
	0x1c, 0xc7, 0x97, 0xae, 0x7f, 0x5d, 0x86, 0xc0, 0x17, 0xac, 0x4e, 0x8b, 0x83, 0xc5, 0xa1, 0x08,
	0x91, 0x4a, 0x80, 0x84, 0xd4, 0x9d, 0x28, 0xda, 0x01, 0x21, 0x06, 0xb2, 0xba, 0xf3, 0x94, 0x34,
	0x26, 0xb3, 0x94, 0xc4, 0xc1, 0xb1, 0xd9, 0x6b, 0x70, 0xe4, 0x15, 0x78, 0x93, 0x1d, 0x79, 0x02,
	0x83, 0xca, 0x1b, 0xe4, 0x09, 0x50, 0xec, 0x24, 0x9b, 0x76, 0xf3, 0xf7, 0xe7, 0xcf, 0xe7, 0x5b,
	0xb7, 0xfd, 0x81, 0x79, 0x2e, 0x12, 0x96, 0x85, 0xa5, 0x14, 0x4a, 0xc0, 0x69, 0x19, 0x55, 0x55,
	0x29, 0xa4, 0x5a, 0xbc, 0x4c, 0xb9, 0xba, 0xd2, 0x71, 0xb8, 0x13, 0xf9, 0x2a, 0x15, 0xa9, 0x58,
	0x59, 0x20, 0xd6, 0x5f, 0x6d, 0xb2, 0xc1, 0x9e, 0x9c, 0xb8, 0xc0, 0xa9, 0x10, 0x69, 0xc6, 0x6e,
	0x29, 0xc5, 0x73, 0x56, 0xa9, 0x28, 0x2f, 0x1d, 0x40, 0x7e, 0x8d, 0xc0, 0xf0, 0xa2, 0x62, 0x12,
	0xbe, 0x01, 0x03, 0x9e, 0x20, 0x2f, 0xf0, 0x96, 0x87, 0x9b, 0x67, 0xb5, 0xc1, 0x41, 0x5c, 0x89,
	0x62, 0x4d, 0x2e, 0x79, 0x42, 0x82, 0x54, 0xc8, 0x7c, 0x4d, 0xde, 0x5d, 0x6c, 0x3f, 0x5f, 0x7e,
	0x38, 0x7f, 0x4f, 0xcf, 0x3e, 0x9d, 0x9d, 0x6f, 0x09, 0x1d, 0xf0, 0x04, 0x7e, 0x04, 0x53, 0x5d,
	0x31, 0x59, 0x44, 0x39, 0x43, 0x83, 0xc0, 0x5b, 0xce, 0x36, 0xab, 0xda, 0xe0, 0x17, 0xce, 0xed,
	0x6e, 0x48, 0x20, 0x6d, 0x81, 0x76, 0x21, 0xd0, 0x05, 0xff, 0xa6, 0xd9, 0x9a, 0x28, 0xa9, 0x19,
	0xa1, 0x7d, 0x01, 0x5c, 0x03, 0xfb, 0x3d, 0xaf, 0x85, 0x4c, 0xd0, 0xa1, 0x2d, 0xf3, 0x6b, 0x83,
	0x17, 0xae, 0xac, 0xbb, 0xe9, 0xca, 0xca, 0xeb, 0x84, 0xd0, 0x9e, 0x87, 0x2b, 0x30, 0xac, 0xa2,
	0x4c, 0xa1, 0xa1, 0xf5, 0x8e, 0x6b, 0x83, 0x9f, 0x38, 0xaf, 0x99, 0x76, 0x8e, 0x3d, 0x53, 0x0b,
	0xc2, 0x53, 0x00, 0x76, 0x57, 0x51, 0x51, 0xb0, 0x4c, 0xf3, 0x04, 0x8d, 0xee, 0x6b, 0x3b, 0xcd,
	0xfb, 0x8f, 0xb2, 0x67, 0x7a, 0x07, 0x87, 0xa7, 0x60, 0xd2, 0x26, 0x34, 0xb6, 0xe6, 0xd3, 0xda,
	0xe0, 0x93, 0xd6, 0x74, 0x17, 0xbd, 0xdc, 0x46, 0xda, 0x19, 0xf0, 0x2d, 0x18, 0x47, 0xdf, 0x23,
	0x15, 0x49, 0x34, 0xb1, 0x2e, 0xae, 0x0d, 0x3e, 0x76, 0xae, 0x9b, 0x77, 0x6a, 0x9b, 0x68, 0x8b,
	0xc3, 0xe7, 0x60, 0x9c, 0x8b, 0x98, 0x67, 0x0c, 0x4d, 0xad, 0xf8, 0xb8, 0x36, 0xf8, 0xc8, 0x89,
	0x6e, 0x4e, 0x68, 0x0b, 0x34, 0xa8, 0x28, 0x59, 0xc1, 0x13, 0x34, 0xbb, 0x8f, 0xba, 0x39, 0xa1,
	0x2d, 0x00, 0x4f, 0xc0, 0x80, 0x97, 0x08, 0x58, 0xec, 0xa8, 0x36, 0x78, 0xe6, 0x30, 0x5e, 0x36,
	0xff, 0x70, 0xd9, 0xbc, 0xb6, 0x52, 0x91, 0xd2, 0x15, 0x9a, 0x07, 0xde, 0x72, 0x74, 0xf7, 0xb5,
	0x6e, 0xde, 0xff, 0xb8, 0x2e, 0xd1, 0x16, 0x87, 0x5f, 0xc0, 0x44, 0xb2, 0x74, 0xcb, 0x73, 0x86,
	0x1e, 0x04, 0xde, 0x72, 0xfe, 0x6a, 0x11, 0xba, 0x65, 0x0c, 0xbb, 0x65, 0x0c, 0xb7, 0xdd, 0x32,
	0x6e, 0x16, 0x37, 0x06, 0x1f, 0xd4, 0x06, 0x3f, 0x74, 0xcd, 0x92, 0xa5, 0xcd, 0xa2, 0x92, 0x1f,
	0x7f, 0xb0, 0x47, 0xbb, 0x9a, 0xcd, 0xa3, 0x9b, 0xbd, 0xef, 0xfd, 0xde, 0xfb, 0xde, 0xdf, 0xbd,
	0xef, 0xfd, 0xfc, 0xe7, 0x1f, 0xc4, 0x63, 0x5b, 0xf5, 0xfa, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xc0, 0x62, 0x6a, 0xf7, 0x2d, 0x03, 0x00, 0x00,
}
