// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: passport.proto

/*
	Package passport is a generated protocol buffer package.

	It is generated from these files:
		passport.proto
		protocol.proto

	It has these top-level messages:
		LoginRegister
		LoginRegisterRet
		LoginLogin
		LoginLoginRet
		ChannelLogin
		ChannelLoginRet
		PassportRequest
		PassportResponse
*/
package passport

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

type LoginRegisterRetRegister_Result int32

const (
	LoginRegisterRet_registerSuccess LoginRegisterRetRegister_Result = 0
	LoginRegisterRet_userExists      LoginRegisterRetRegister_Result = 1
	LoginRegisterRet_invalidFormat   LoginRegisterRetRegister_Result = 2
)

var LoginRegisterRetRegister_Result_name = map[int32]string{
	0: "registerSuccess",
	1: "userExists",
	2: "invalidFormat",
}
var LoginRegisterRetRegister_Result_value = map[string]int32{
	"registerSuccess": 0,
	"userExists":      1,
	"invalidFormat":   2,
}

func (x LoginRegisterRetRegister_Result) String() string {
	return proto.EnumName(LoginRegisterRetRegister_Result_name, int32(x))
}
func (LoginRegisterRetRegister_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorPassport, []int{1, 0}
}

// 服务端不允许login文件名存在,特改为bblogin
// 登录相关通讯协议
type LoginLoginRetLogin_Result int32

const (
	LoginLoginRet_loginSuccess  LoginLoginRetLogin_Result = 0
	LoginLoginRet_invalidUser   LoginLoginRetLogin_Result = 1
	LoginLoginRet_invalidPwd    LoginLoginRetLogin_Result = 2
	LoginLoginRet_forbiddenUser LoginLoginRetLogin_Result = 3
)

var LoginLoginRetLogin_Result_name = map[int32]string{
	0: "loginSuccess",
	1: "invalidUser",
	2: "invalidPwd",
	3: "forbiddenUser",
}
var LoginLoginRetLogin_Result_value = map[string]int32{
	"loginSuccess":  0,
	"invalidUser":   1,
	"invalidPwd":    2,
	"forbiddenUser": 3,
}

func (x LoginLoginRetLogin_Result) String() string {
	return proto.EnumName(LoginLoginRetLogin_Result_name, int32(x))
}
func (LoginLoginRetLogin_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorPassport, []int{3, 0}
}

// 普通注册账号
type LoginRegister struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Server   int32  `protobuf:"varint,3,opt,name=server,proto3" json:"server,omitempty"`
}

func (m *LoginRegister) Reset()                    { *m = LoginRegister{} }
func (m *LoginRegister) String() string            { return proto.CompactTextString(m) }
func (*LoginRegister) ProtoMessage()               {}
func (*LoginRegister) Descriptor() ([]byte, []int) { return fileDescriptorPassport, []int{0} }

func (m *LoginRegister) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRegister) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRegister) GetServer() int32 {
	if m != nil {
		return m.Server
	}
	return 0
}

// 登录服务器注册账号返回
type LoginRegisterRet struct {
	Result LoginRegisterRetRegister_Result `protobuf:"varint,1,opt,name=result,proto3,enum=passport.LoginRegisterRetRegister_Result" json:"result,omitempty"`
	Uid    int64                           `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Token  string                          `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Msg    string                          `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *LoginRegisterRet) Reset()                    { *m = LoginRegisterRet{} }
func (m *LoginRegisterRet) String() string            { return proto.CompactTextString(m) }
func (*LoginRegisterRet) ProtoMessage()               {}
func (*LoginRegisterRet) Descriptor() ([]byte, []int) { return fileDescriptorPassport, []int{1} }

func (m *LoginRegisterRet) GetResult() LoginRegisterRetRegister_Result {
	if m != nil {
		return m.Result
	}
	return LoginRegisterRet_registerSuccess
}

func (m *LoginRegisterRet) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *LoginRegisterRet) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginRegisterRet) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// 用户名密码登录服务器
type LoginLogin struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *LoginLogin) Reset()                    { *m = LoginLogin{} }
func (m *LoginLogin) String() string            { return proto.CompactTextString(m) }
func (*LoginLogin) ProtoMessage()               {}
func (*LoginLogin) Descriptor() ([]byte, []int) { return fileDescriptorPassport, []int{2} }

func (m *LoginLogin) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginLogin) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// 登录服务器返回
type LoginLoginRet struct {
	Result LoginLoginRetLogin_Result `protobuf:"varint,1,opt,name=result,proto3,enum=passport.LoginLoginRetLogin_Result" json:"result,omitempty"`
	Uid    int64                     `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Token  string                    `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Msg    string                    `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *LoginLoginRet) Reset()                    { *m = LoginLoginRet{} }
func (m *LoginLoginRet) String() string            { return proto.CompactTextString(m) }
func (*LoginLoginRet) ProtoMessage()               {}
func (*LoginLoginRet) Descriptor() ([]byte, []int) { return fileDescriptorPassport, []int{3} }

func (m *LoginLoginRet) GetResult() LoginLoginRetLogin_Result {
	if m != nil {
		return m.Result
	}
	return LoginLoginRet_loginSuccess
}

func (m *LoginLoginRet) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *LoginLoginRet) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginLoginRet) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// 渠道登录
type ChannelLogin struct {
	Channel    string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	ChannelUID string `protobuf:"bytes,2,opt,name=channelUID,proto3" json:"channelUID,omitempty"`
	Sdk        string `protobuf:"bytes,3,opt,name=sdk,proto3" json:"sdk,omitempty"`
	Ip         string `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (m *ChannelLogin) Reset()                    { *m = ChannelLogin{} }
func (m *ChannelLogin) String() string            { return proto.CompactTextString(m) }
func (*ChannelLogin) ProtoMessage()               {}
func (*ChannelLogin) Descriptor() ([]byte, []int) { return fileDescriptorPassport, []int{4} }

func (m *ChannelLogin) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *ChannelLogin) GetChannelUID() string {
	if m != nil {
		return m.ChannelUID
	}
	return ""
}

func (m *ChannelLogin) GetSdk() string {
	if m != nil {
		return m.Sdk
	}
	return ""
}

func (m *ChannelLogin) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

// 渠道登录结果
type ChannelLoginRet struct {
	Uid   int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *ChannelLoginRet) Reset()                    { *m = ChannelLoginRet{} }
func (m *ChannelLoginRet) String() string            { return proto.CompactTextString(m) }
func (*ChannelLoginRet) ProtoMessage()               {}
func (*ChannelLoginRet) Descriptor() ([]byte, []int) { return fileDescriptorPassport, []int{5} }

func (m *ChannelLoginRet) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *ChannelLoginRet) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRegister)(nil), "passport.login_register")
	proto.RegisterType((*LoginRegisterRet)(nil), "passport.login_register_ret")
	proto.RegisterType((*LoginLogin)(nil), "passport.login_login")
	proto.RegisterType((*LoginLoginRet)(nil), "passport.login_login_ret")
	proto.RegisterType((*ChannelLogin)(nil), "passport.channel_login")
	proto.RegisterType((*ChannelLoginRet)(nil), "passport.channel_login_ret")
	proto.RegisterEnum("passport.LoginRegisterRetRegister_Result", LoginRegisterRetRegister_Result_name, LoginRegisterRetRegister_Result_value)
	proto.RegisterEnum("passport.LoginLoginRetLogin_Result", LoginLoginRetLogin_Result_name, LoginLoginRetLogin_Result_value)
}
func (m *LoginRegister) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginRegister) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Username) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPassport(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPassport(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	if m.Server != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintPassport(dAtA, i, uint64(m.Server))
	}
	return i, nil
}

func (m *LoginRegisterRet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginRegisterRet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Result != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPassport(dAtA, i, uint64(m.Result))
	}
	if m.Uid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPassport(dAtA, i, uint64(m.Uid))
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPassport(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintPassport(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	return i, nil
}

func (m *LoginLogin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginLogin) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Username) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPassport(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPassport(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	return i, nil
}

func (m *LoginLoginRet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginLoginRet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Result != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPassport(dAtA, i, uint64(m.Result))
	}
	if m.Uid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPassport(dAtA, i, uint64(m.Uid))
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPassport(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintPassport(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	return i, nil
}

func (m *ChannelLogin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChannelLogin) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Channel) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPassport(dAtA, i, uint64(len(m.Channel)))
		i += copy(dAtA[i:], m.Channel)
	}
	if len(m.ChannelUID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPassport(dAtA, i, uint64(len(m.ChannelUID)))
		i += copy(dAtA[i:], m.ChannelUID)
	}
	if len(m.Sdk) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPassport(dAtA, i, uint64(len(m.Sdk)))
		i += copy(dAtA[i:], m.Sdk)
	}
	if len(m.Ip) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintPassport(dAtA, i, uint64(len(m.Ip)))
		i += copy(dAtA[i:], m.Ip)
	}
	return i, nil
}

func (m *ChannelLoginRet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChannelLoginRet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Uid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPassport(dAtA, i, uint64(m.Uid))
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPassport(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	return i, nil
}

func encodeVarintPassport(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LoginRegister) Size() (n int) {
	var l int
	_ = l
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovPassport(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovPassport(uint64(l))
	}
	if m.Server != 0 {
		n += 1 + sovPassport(uint64(m.Server))
	}
	return n
}

func (m *LoginRegisterRet) Size() (n int) {
	var l int
	_ = l
	if m.Result != 0 {
		n += 1 + sovPassport(uint64(m.Result))
	}
	if m.Uid != 0 {
		n += 1 + sovPassport(uint64(m.Uid))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovPassport(uint64(l))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovPassport(uint64(l))
	}
	return n
}

func (m *LoginLogin) Size() (n int) {
	var l int
	_ = l
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovPassport(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovPassport(uint64(l))
	}
	return n
}

func (m *LoginLoginRet) Size() (n int) {
	var l int
	_ = l
	if m.Result != 0 {
		n += 1 + sovPassport(uint64(m.Result))
	}
	if m.Uid != 0 {
		n += 1 + sovPassport(uint64(m.Uid))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovPassport(uint64(l))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovPassport(uint64(l))
	}
	return n
}

func (m *ChannelLogin) Size() (n int) {
	var l int
	_ = l
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovPassport(uint64(l))
	}
	l = len(m.ChannelUID)
	if l > 0 {
		n += 1 + l + sovPassport(uint64(l))
	}
	l = len(m.Sdk)
	if l > 0 {
		n += 1 + l + sovPassport(uint64(l))
	}
	l = len(m.Ip)
	if l > 0 {
		n += 1 + l + sovPassport(uint64(l))
	}
	return n
}

func (m *ChannelLoginRet) Size() (n int) {
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovPassport(uint64(m.Uid))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovPassport(uint64(l))
	}
	return n
}

func sovPassport(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPassport(x uint64) (n int) {
	return sovPassport(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LoginRegister) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPassport
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
			return fmt.Errorf("proto: login_register: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: login_register: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
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
				return ErrInvalidLengthPassport
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
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
				return ErrInvalidLengthPassport
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Server", wireType)
			}
			m.Server = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Server |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPassport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPassport
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
func (m *LoginRegisterRet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPassport
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
			return fmt.Errorf("proto: login_register_ret: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: login_register_ret: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= (LoginRegisterRetRegister_Result(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
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
				return ErrInvalidLengthPassport
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
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
				return ErrInvalidLengthPassport
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPassport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPassport
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
func (m *LoginLogin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPassport
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
			return fmt.Errorf("proto: login_login: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: login_login: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
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
				return ErrInvalidLengthPassport
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
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
				return ErrInvalidLengthPassport
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPassport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPassport
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
func (m *LoginLoginRet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPassport
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
			return fmt.Errorf("proto: login_login_ret: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: login_login_ret: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= (LoginLoginRetLogin_Result(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
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
				return ErrInvalidLengthPassport
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
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
				return ErrInvalidLengthPassport
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPassport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPassport
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
func (m *ChannelLogin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPassport
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
			return fmt.Errorf("proto: channel_login: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: channel_login: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
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
				return ErrInvalidLengthPassport
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
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
				return ErrInvalidLengthPassport
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sdk", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
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
				return ErrInvalidLengthPassport
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sdk = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
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
				return ErrInvalidLengthPassport
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ip = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPassport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPassport
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
func (m *ChannelLoginRet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPassport
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
			return fmt.Errorf("proto: channel_login_ret: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: channel_login_ret: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
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
				return ErrInvalidLengthPassport
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPassport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPassport
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
func skipPassport(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPassport
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
					return 0, ErrIntOverflowPassport
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
					return 0, ErrIntOverflowPassport
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
				return 0, ErrInvalidLengthPassport
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPassport
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
				next, err := skipPassport(dAtA[start:])
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
	ErrInvalidLengthPassport = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPassport   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("passport.proto", fileDescriptorPassport) }

var fileDescriptorPassport = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x8e, 0xda, 0x30,
	0x10, 0xc6, 0x71, 0x52, 0x28, 0x0c, 0x25, 0x18, 0xb7, 0xaa, 0xa2, 0x1e, 0x22, 0x94, 0x43, 0xc5,
	0xa1, 0xe2, 0xd0, 0x1e, 0x2b, 0xf5, 0x50, 0x41, 0x25, 0x6e, 0x55, 0x5a, 0xce, 0x34, 0x10, 0x97,
	0x5a, 0x84, 0x38, 0xb2, 0x1d, 0xe8, 0xa3, 0xf4, 0x91, 0x7a, 0xdc, 0x37, 0xd8, 0x5d, 0xf6, 0x45,
	0x56, 0x76, 0x4c, 0x08, 0xec, 0x9e, 0x56, 0x7b, 0x89, 0xe6, 0x1b, 0xcf, 0x9f, 0x6f, 0x7e, 0x52,
	0xc0, 0xcb, 0x63, 0x29, 0x73, 0x2e, 0xd4, 0x38, 0x17, 0x5c, 0x71, 0xd2, 0x3e, 0xea, 0xf0, 0x17,
	0x78, 0x29, 0x5f, 0xb3, 0x6c, 0x21, 0xe8, 0x9a, 0x49, 0x45, 0x05, 0x79, 0x07, 0xed, 0x42, 0x52,
	0x91, 0xc5, 0x5b, 0xea, 0xa3, 0x21, 0x1a, 0x75, 0xa2, 0x4a, 0xeb, 0x37, 0xdd, 0xb9, 0xe7, 0x22,
	0xf1, 0x9d, 0xf2, 0xed, 0xa8, 0xc9, 0x5b, 0x68, 0x49, 0x2a, 0x76, 0x54, 0xf8, 0xee, 0x10, 0x8d,
	0x9a, 0x91, 0x55, 0xe1, 0x35, 0x02, 0x72, 0xbe, 0x62, 0x21, 0xa8, 0x22, 0x13, 0x68, 0x09, 0x2a,
	0x8b, 0x54, 0x99, 0x25, 0xde, 0xc7, 0x0f, 0xe3, 0xca, 0xe3, 0xc3, 0xea, 0x71, 0x25, 0x22, 0xd3,
	0x13, 0xd9, 0x5e, 0x82, 0xc1, 0x2d, 0x58, 0xe9, 0xc5, 0x8d, 0x74, 0x48, 0xde, 0x40, 0x53, 0xf1,
	0x0d, 0xcd, 0x8c, 0x8b, 0x4e, 0x54, 0x0a, 0x5d, 0xb7, 0x95, 0x6b, 0xff, 0x85, 0xc9, 0xe9, 0x30,
	0x9c, 0x41, 0xff, 0x62, 0x28, 0x79, 0x7d, 0x4a, 0xfd, 0x28, 0x56, 0x2b, 0x2a, 0x25, 0x6e, 0x10,
	0x0f, 0x40, 0x9f, 0x3f, 0xfd, 0xcb, 0xa4, 0x92, 0x18, 0x91, 0x01, 0xf4, 0x58, 0xb6, 0x8b, 0x53,
	0x96, 0x7c, 0xe3, 0x62, 0x1b, 0x2b, 0xec, 0x84, 0x53, 0xe8, 0x96, 0x96, 0xcd, 0xf7, 0xa9, 0x00,
	0xc3, 0x5b, 0x04, 0xfd, 0xda, 0x1c, 0x43, 0xe9, 0xcb, 0x05, 0xa5, 0xf7, 0x97, 0x94, 0xaa, 0x52,
	0xab, 0x9f, 0x8d, 0xcf, 0x4f, 0x78, 0x55, 0x9f, 0x48, 0xb0, 0xd5, 0x27, 0x32, 0x7d, 0xe8, 0x5a,
	0x12, 0x73, 0x49, 0x05, 0x46, 0x1a, 0x95, 0x4d, 0x7c, 0xdf, 0x27, 0xd8, 0xd1, 0xa8, 0x7e, 0x73,
	0xb1, 0x64, 0x49, 0x42, 0x33, 0x53, 0xe2, 0x86, 0x1b, 0xe8, 0xad, 0xfe, 0xc4, 0x59, 0x46, 0x53,
	0x0b, 0xcb, 0x87, 0x97, 0x36, 0x61, 0x59, 0x1d, 0x25, 0x09, 0x00, 0x6c, 0x38, 0x9f, 0x4d, 0x2c,
	0xac, 0x5a, 0x46, 0x5b, 0x96, 0xc9, 0xc6, 0x9e, 0xa1, 0x43, 0xe2, 0x81, 0xc3, 0x72, 0x7b, 0x83,
	0xc3, 0xf2, 0xf0, 0x33, 0x0c, 0xce, 0x96, 0x19, 0xa2, 0x96, 0x08, 0x7a, 0x84, 0x88, 0x53, 0x23,
	0xf2, 0x15, 0xff, 0x3f, 0x04, 0xe8, 0xea, 0x10, 0xa0, 0x9b, 0x43, 0x80, 0xfe, 0xdd, 0x05, 0x8d,
	0x65, 0xcb, 0xfc, 0x3b, 0x9f, 0xee, 0x03, 0x00, 0x00, 0xff, 0xff, 0x13, 0x41, 0x16, 0xf5, 0x4d,
	0x03, 0x00, 0x00,
}
