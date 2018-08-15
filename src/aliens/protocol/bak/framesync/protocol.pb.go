// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protocol.proto

package framesync

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// request
type Request struct {
	Auth *Auth `protobuf:"bytes,4,opt,name=auth" json:"auth,omitempty"`
	// -----------------帧同步服务接口---------------
	Command          *Command          `protobuf:"bytes,5,opt,name=command" json:"command,omitempty"`
	RequestLostFrame *RequestLostFrame `protobuf:"bytes,6,opt,name=requestLostFrame" json:"requestLostFrame,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{0} }

func (m *Request) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *Request) GetCommand() *Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *Request) GetRequestLostFrame() *RequestLostFrame {
	if m != nil {
		return m.RequestLostFrame
	}
	return nil
}

// response
type Response struct {
	AuthRet             bool                 `protobuf:"varint,4,opt,name=authRet,proto3" json:"authRet,omitempty"`
	RequestLostFrameRet *RequestLostFrameRet `protobuf:"bytes,6,opt,name=requestLostFrameRet" json:"requestLostFrameRet,omitempty"`
	SyncFrame           *Frame               `protobuf:"bytes,10,opt,name=syncFrame" json:"syncFrame,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{1} }

func (m *Response) GetAuthRet() bool {
	if m != nil {
		return m.AuthRet
	}
	return false
}

func (m *Response) GetRequestLostFrameRet() *RequestLostFrameRet {
	if m != nil {
		return m.RequestLostFrameRet
	}
	return nil
}

func (m *Response) GetSyncFrame() *Frame {
	if m != nil {
		return m.SyncFrame
	}
	return nil
}

type Auth struct {
	RoomID string `protobuf:"bytes,1,opt,name=roomID,proto3" json:"roomID,omitempty"`
	SeatID uint32 `protobuf:"varint,2,opt,name=seatID,proto3" json:"seatID,omitempty"`
	Token  string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *Auth) Reset()                    { *m = Auth{} }
func (m *Auth) String() string            { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()               {}
func (*Auth) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{2} }

func (m *Auth) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

func (m *Auth) GetSeatID() uint32 {
	if m != nil {
		return m.SeatID
	}
	return 0
}

func (m *Auth) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "framesync.Request")
	proto.RegisterType((*Response)(nil), "framesync.Response")
	proto.RegisterType((*Auth)(nil), "framesync.auth")
}
func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Auth != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.Auth.Size()))
		n1, err := m.Auth.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Command != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.Command.Size()))
		n2, err := m.Command.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.RequestLostFrame != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.RequestLostFrame.Size()))
		n3, err := m.RequestLostFrame.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AuthRet {
		dAtA[i] = 0x20
		i++
		if m.AuthRet {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.RequestLostFrameRet != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.RequestLostFrameRet.Size()))
		n4, err := m.RequestLostFrameRet.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.SyncFrame != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.SyncFrame.Size()))
		n5, err := m.SyncFrame.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}

func (m *Auth) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Auth) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RoomID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.RoomID)))
		i += copy(dAtA[i:], m.RoomID)
	}
	if m.SeatID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.SeatID))
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	return i, nil
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
func (m *Request) Size() (n int) {
	var l int
	_ = l
	if m.Auth != nil {
		l = m.Auth.Size()
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.Command != nil {
		l = m.Command.Size()
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.RequestLostFrame != nil {
		l = m.RequestLostFrame.Size()
		n += 1 + l + sovProtocol(uint64(l))
	}
	return n
}

func (m *Response) Size() (n int) {
	var l int
	_ = l
	if m.AuthRet {
		n += 2
	}
	if m.RequestLostFrameRet != nil {
		l = m.RequestLostFrameRet.Size()
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.SyncFrame != nil {
		l = m.SyncFrame.Size()
		n += 1 + l + sovProtocol(uint64(l))
	}
	return n
}

func (m *Auth) Size() (n int) {
	var l int
	_ = l
	l = len(m.RoomID)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.SeatID != 0 {
		n += 1 + sovProtocol(uint64(m.SeatID))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
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
func (m *Request) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field auth", wireType)
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
			if m.Auth == nil {
				m.Auth = &Auth{}
			}
			if err := m.Auth.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Command", wireType)
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
			if m.Command == nil {
				m.Command = &Command{}
			}
			if err := m.Command.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestLostFrame", wireType)
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
			if m.RequestLostFrame == nil {
				m.RequestLostFrame = &RequestLostFrame{}
			}
			if err := m.RequestLostFrame.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthRet", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AuthRet = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestLostFrameRet", wireType)
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
			if m.RequestLostFrameRet == nil {
				m.RequestLostFrameRet = &RequestLostFrameRet{}
			}
			if err := m.RequestLostFrameRet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SyncFrame", wireType)
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
			if m.SyncFrame == nil {
				m.SyncFrame = &Frame{}
			}
			if err := m.SyncFrame.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Auth) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: auth: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: auth: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
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
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoomID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeatID", wireType)
			}
			m.SeatID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SeatID |= (uint32(b) & 0x7F) << shift
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
					return ErrIntOverflowProtocol
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
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
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
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x4a, 0xf4, 0x40,
	0x10, 0xfd, 0xfa, 0x73, 0x26, 0x3f, 0x25, 0x3a, 0x43, 0x8f, 0x68, 0xe3, 0x22, 0x84, 0xe8, 0x22,
	0xab, 0x2c, 0xf4, 0x00, 0x82, 0x0c, 0x42, 0x60, 0x56, 0x75, 0x81, 0x10, 0x63, 0x8b, 0xe0, 0x24,
	0x35, 0xa6, 0x7b, 0x16, 0x9e, 0x43, 0x04, 0x8f, 0xe4, 0xd2, 0x23, 0x48, 0xbc, 0x88, 0xa4, 0x3a,
	0xe3, 0x80, 0xce, 0xae, 0x5e, 0xbd, 0x57, 0xef, 0x55, 0x57, 0xc3, 0xe1, 0xaa, 0x25, 0x4b, 0x15,
	0x2d, 0x33, 0x2e, 0x64, 0x78, 0xdf, 0x96, 0xb5, 0x36, 0xcf, 0x4d, 0x75, 0x3a, 0xf9, 0x29, 0x1d,
	0x97, 0xbc, 0x08, 0xf0, 0x51, 0x3f, 0xad, 0xb5, 0xb1, 0xf2, 0x0c, 0x46, 0xe5, 0xda, 0x3e, 0xa8,
	0x51, 0x2c, 0xd2, 0xfd, 0x8b, 0x49, 0xb6, 0xd5, 0xf6, 0x6d, 0x64, 0x52, 0x26, 0xe0, 0x57, 0x54,
	0xd7, 0x65, 0x73, 0xa7, 0xc6, 0xac, 0x0b, 0xb2, 0x01, 0xe3, 0x86, 0x90, 0x57, 0x30, 0x6d, 0x9d,
	0xe7, 0x82, 0x8c, 0xbd, 0xe9, 0x6d, 0x94, 0xc7, 0xe2, 0x59, 0x36, 0x10, 0xc5, 0x92, 0x8c, 0x2d,
	0x38, 0x01, 0xff, 0x88, 0x93, 0x57, 0x01, 0x01, 0x6a, 0xb3, 0xa2, 0xc6, 0x68, 0xa9, 0xc0, 0xe7,
	0x7c, 0x6d, 0x79, 0xb3, 0x00, 0x37, 0x50, 0xe6, 0x30, 0xfb, 0x3d, 0xda, 0xab, 0x5c, 0xd4, 0xc9,
	0x8e, 0xa8, 0xa2, 0xd5, 0x16, 0x77, 0xcd, 0xc8, 0x73, 0x08, 0xfb, 0x97, 0xba, 0x5d, 0x81, 0x0d,
	0x3c, 0x77, 0x00, 0xdc, 0x12, 0xc9, 0xc2, 0x5d, 0x48, 0x1e, 0x83, 0xd7, 0x12, 0xd5, 0xf9, 0x5c,
	0x89, 0x58, 0xa4, 0x21, 0x0e, 0xa8, 0xef, 0x1b, 0x5d, 0xda, 0x7c, 0xae, 0xfe, 0xc7, 0x22, 0x3d,
	0xc0, 0x01, 0xc9, 0x23, 0x18, 0x5b, 0x7a, 0xd4, 0x8d, 0xda, 0x63, 0xb9, 0x03, 0xd7, 0xd3, 0xf7,
	0x2e, 0x12, 0x1f, 0x5d, 0x24, 0x3e, 0xbb, 0x48, 0xbc, 0x7d, 0x45, 0xff, 0x6e, 0x3d, 0xfe, 0x94,
	0xcb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xc1, 0x3f, 0x71, 0xc2, 0x01, 0x00, 0x00,
}