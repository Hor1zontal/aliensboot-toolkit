// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: room_model.proto

package protocol

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 游戏结算分为1v1、2v2结算，data.resultDisplay用于客户端结算面板的显示，data.resultData用户后台上报。游戏分胜利，失败，平局，逃跑四种结果。
type GameResult int32

const (
	GameResult_Win    GameResult = 0
	GameResult_Lose   GameResult = 1
	GameResult_Equal  GameResult = 2
	GameResult_Escape GameResult = 3
)

var GameResult_name = map[int32]string{
	0: "Win",
	1: "Lose",
	2: "Equal",
	3: "Escape",
}
var GameResult_value = map[string]int32{
	"Win":    0,
	"Lose":   1,
	"Equal":  2,
	"Escape": 3,
}

func (x GameResult) String() string {
	return proto.EnumName(GameResult_name, int32(x))
}
func (GameResult) EnumDescriptor() ([]byte, []int) { return fileDescriptorRoomModel, []int{0} }

type PlayerResult struct {
	Playerid int64   `protobuf:"varint,1,opt,name=playerid,proto3" json:"playerid,omitempty"`
	Record   *Record `protobuf:"bytes,3,opt,name=record" json:"record,omitempty"`
}

func (m *PlayerResult) Reset()                    { *m = PlayerResult{} }
func (m *PlayerResult) String() string            { return proto.CompactTextString(m) }
func (*PlayerResult) ProtoMessage()               {}
func (*PlayerResult) Descriptor() ([]byte, []int) { return fileDescriptorRoomModel, []int{0} }

func (m *PlayerResult) GetPlayerid() int64 {
	if m != nil {
		return m.Playerid
	}
	return 0
}

func (m *PlayerResult) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type Record struct {
	Result GameResult `protobuf:"varint,1,opt,name=result,proto3,enum=protocol.GameResult" json:"result,omitempty"`
	Score  int32      `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	Unit   string     `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptorRoomModel, []int{1} }

func (m *Record) GetResult() GameResult {
	if m != nil {
		return m.Result
	}
	return GameResult_Win
}

func (m *Record) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Record) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

type Player struct {
	Playerid int64  `protobuf:"varint,1,opt,name=playerid,proto3" json:"playerid,omitempty"`
	GroupId  string `protobuf:"bytes,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Seat     int32  `protobuf:"varint,3,opt,name=seat,proto3" json:"seat,omitempty"`
	Nickname string `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Headurl  string `protobuf:"bytes,5,opt,name=headurl,proto3" json:"headurl,omitempty"`
	Gender   string `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender,omitempty"`
}

func (m *Player) Reset()                    { *m = Player{} }
func (m *Player) String() string            { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()               {}
func (*Player) Descriptor() ([]byte, []int) { return fileDescriptorRoomModel, []int{2} }

func (m *Player) GetPlayerid() int64 {
	if m != nil {
		return m.Playerid
	}
	return 0
}

func (m *Player) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *Player) GetSeat() int32 {
	if m != nil {
		return m.Seat
	}
	return 0
}

func (m *Player) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *Player) GetHeadurl() string {
	if m != nil {
		return m.Headurl
	}
	return ""
}

func (m *Player) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func init() {
	proto.RegisterType((*PlayerResult)(nil), "protocol.PlayerResult")
	proto.RegisterType((*Record)(nil), "protocol.Record")
	proto.RegisterType((*Player)(nil), "protocol.Player")
	proto.RegisterEnum("protocol.GameResult", GameResult_name, GameResult_value)
}
func (m *PlayerResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlayerResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Playerid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.Playerid))
	}
	if m.Record != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.Record.Size()))
		n1, err := m.Record.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *Record) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Result != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.Result))
	}
	if m.Score != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.Score))
	}
	if len(m.Unit) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(len(m.Unit)))
		i += copy(dAtA[i:], m.Unit)
	}
	return i, nil
}

func (m *Player) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Player) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Playerid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.Playerid))
	}
	if len(m.GroupId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(len(m.GroupId)))
		i += copy(dAtA[i:], m.GroupId)
	}
	if m.Seat != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.Seat))
	}
	if len(m.Nickname) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(len(m.Nickname)))
		i += copy(dAtA[i:], m.Nickname)
	}
	if len(m.Headurl) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(len(m.Headurl)))
		i += copy(dAtA[i:], m.Headurl)
	}
	if len(m.Gender) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(len(m.Gender)))
		i += copy(dAtA[i:], m.Gender)
	}
	return i, nil
}

func encodeVarintRoomModel(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PlayerResult) Size() (n int) {
	var l int
	_ = l
	if m.Playerid != 0 {
		n += 1 + sovRoomModel(uint64(m.Playerid))
	}
	if m.Record != nil {
		l = m.Record.Size()
		n += 1 + l + sovRoomModel(uint64(l))
	}
	return n
}

func (m *Record) Size() (n int) {
	var l int
	_ = l
	if m.Result != 0 {
		n += 1 + sovRoomModel(uint64(m.Result))
	}
	if m.Score != 0 {
		n += 1 + sovRoomModel(uint64(m.Score))
	}
	l = len(m.Unit)
	if l > 0 {
		n += 1 + l + sovRoomModel(uint64(l))
	}
	return n
}

func (m *Player) Size() (n int) {
	var l int
	_ = l
	if m.Playerid != 0 {
		n += 1 + sovRoomModel(uint64(m.Playerid))
	}
	l = len(m.GroupId)
	if l > 0 {
		n += 1 + l + sovRoomModel(uint64(l))
	}
	if m.Seat != 0 {
		n += 1 + sovRoomModel(uint64(m.Seat))
	}
	l = len(m.Nickname)
	if l > 0 {
		n += 1 + l + sovRoomModel(uint64(l))
	}
	l = len(m.Headurl)
	if l > 0 {
		n += 1 + l + sovRoomModel(uint64(l))
	}
	l = len(m.Gender)
	if l > 0 {
		n += 1 + l + sovRoomModel(uint64(l))
	}
	return n
}

func sovRoomModel(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRoomModel(x uint64) (n int) {
	return sovRoomModel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PlayerResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoomModel
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
			return fmt.Errorf("proto: PlayerResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlayerResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Playerid", wireType)
			}
			m.Playerid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Playerid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Record", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
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
				return ErrInvalidLengthRoomModel
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Record == nil {
				m.Record = &Record{}
			}
			if err := m.Record.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoomModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoomModel
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
func (m *Record) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoomModel
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
			return fmt.Errorf("proto: Record: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= (GameResult(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			m.Score = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Score |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
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
				return ErrInvalidLengthRoomModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Unit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoomModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoomModel
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
func (m *Player) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoomModel
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
			return fmt.Errorf("proto: Player: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Player: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Playerid", wireType)
			}
			m.Playerid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Playerid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
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
				return ErrInvalidLengthRoomModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seat", wireType)
			}
			m.Seat = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seat |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nickname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
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
				return ErrInvalidLengthRoomModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nickname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headurl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
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
				return ErrInvalidLengthRoomModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headurl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
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
				return ErrInvalidLengthRoomModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoomModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoomModel
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
func skipRoomModel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRoomModel
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
					return 0, ErrIntOverflowRoomModel
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
					return 0, ErrIntOverflowRoomModel
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
				return 0, ErrInvalidLengthRoomModel
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRoomModel
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
				next, err := skipRoomModel(dAtA[start:])
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
	ErrInvalidLengthRoomModel = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRoomModel   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("room_model.proto", fileDescriptorRoomModel) }

var fileDescriptorRoomModel = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x10, 0xc0, 0xbf, 0x6d, 0x9a, 0x6d, 0x3b, 0x5f, 0x91, 0x30, 0x14, 0x59, 0x3c, 0x94, 0xd2, 0x53,
	0x10, 0xe9, 0xa1, 0x82, 0x0f, 0x20, 0x14, 0x11, 0x3c, 0xc8, 0x22, 0x78, 0xd4, 0x35, 0x3b, 0xd4,
	0xe0, 0x26, 0x1b, 0x37, 0xcd, 0xc1, 0xb7, 0xf1, 0x71, 0x3c, 0xfa, 0x08, 0x92, 0x27, 0x91, 0xec,
	0xa6, 0xed, 0xcd, 0x53, 0xe6, 0x37, 0x7f, 0x7e, 0x33, 0x59, 0x48, 0x9c, 0xb5, 0xc5, 0x53, 0x61,
	0x35, 0x99, 0x55, 0xe5, 0xec, 0xce, 0xe2, 0xd8, 0x7f, 0x32, 0x6b, 0x96, 0x0f, 0x30, 0xbd, 0x37,
	0xea, 0x83, 0x9c, 0xa4, 0xba, 0x31, 0x3b, 0x3c, 0x83, 0x71, 0xe5, 0x39, 0xd7, 0x82, 0x2d, 0x58,
	0x1a, 0xc9, 0x03, 0x63, 0x0a, 0xdc, 0x51, 0x66, 0x9d, 0x16, 0xd1, 0x82, 0xa5, 0xff, 0xd7, 0xc9,
	0x6a, 0xaf, 0x59, 0x49, 0x9f, 0x97, 0x7d, 0x7d, 0xf9, 0x0c, 0x3c, 0x64, 0xf0, 0xa2, 0x9b, 0xe9,
	0xcc, 0xde, 0x76, 0xb2, 0x9e, 0x1d, 0x67, 0x6e, 0x54, 0x41, 0x61, 0xab, 0xec, 0x7b, 0x70, 0x06,
	0x71, 0x9d, 0x59, 0x47, 0x62, 0xb0, 0x60, 0x69, 0x2c, 0x03, 0x20, 0xc2, 0xb0, 0x29, 0xf3, 0x9d,
	0xdf, 0x3a, 0x91, 0x3e, 0x5e, 0x7e, 0x32, 0xe0, 0xe1, 0xf0, 0x3f, 0x4f, 0x16, 0x30, 0xda, 0x3a,
	0xdb, 0x54, 0xb7, 0xda, 0x2b, 0x27, 0x72, 0x8f, 0x9d, 0xb4, 0x26, 0x15, 0xa4, 0xb1, 0xf4, 0x71,
	0x67, 0x2a, 0xf3, 0xec, 0xad, 0x54, 0x05, 0x89, 0xa1, 0x6f, 0x3f, 0x70, 0x67, 0x7a, 0x25, 0xa5,
	0x1b, 0x67, 0x44, 0x1c, 0x4c, 0x3d, 0xe2, 0x29, 0xf0, 0x2d, 0x95, 0x9a, 0x9c, 0xe0, 0xbe, 0xd0,
	0xd3, 0xf9, 0x15, 0xc0, 0xf1, 0x17, 0x71, 0x04, 0xd1, 0x63, 0x5e, 0x26, 0xff, 0x70, 0x0c, 0xc3,
	0x3b, 0x5b, 0x53, 0xc2, 0x70, 0x02, 0xf1, 0xe6, 0xbd, 0x51, 0x26, 0x19, 0x20, 0x00, 0xdf, 0xd4,
	0x99, 0xaa, 0x28, 0x89, 0xae, 0xa7, 0x5f, 0xed, 0x9c, 0x7d, 0xb7, 0x73, 0xf6, 0xd3, 0xce, 0xd9,
	0x0b, 0xf7, 0xef, 0x75, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xac, 0xf6, 0xac, 0xc5, 0x01,
	0x00, 0x00,
}
