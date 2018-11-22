// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: resultcode.proto

package protocol

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Code int32

const (
	Code_Success         Code = 0
	Code_InvalidRequest  Code = 1
	Code_ServerException Code = 2
	Code_DBExcetpion     Code = 3
	Code_ConfigException Code = 4
	// game
	Code_ValidateException Code = 9
	// room
	Code_appIDNotFound    Code = 50
	Code_playerNotFound   Code = 51
	Code_roomNotFound     Code = 60
	Code_roomMaxPlayer    Code = 61
	Code_gameNotFound     Code = 101
	Code_gameAlreadyStart Code = 103
	Code_InvalidService   Code = 1000
)

var Code_name = map[int32]string{
	0:    "Success",
	1:    "InvalidRequest",
	2:    "ServerException",
	3:    "DBExcetpion",
	4:    "ConfigException",
	9:    "ValidateException",
	50:   "appIDNotFound",
	51:   "playerNotFound",
	60:   "roomNotFound",
	61:   "roomMaxPlayer",
	101:  "gameNotFound",
	103:  "gameAlreadyStart",
	1000: "InvalidService",
}
var Code_value = map[string]int32{
	"Success":           0,
	"InvalidRequest":    1,
	"ServerException":   2,
	"DBExcetpion":       3,
	"ConfigException":   4,
	"ValidateException": 9,
	"appIDNotFound":     50,
	"playerNotFound":    51,
	"roomNotFound":      60,
	"roomMaxPlayer":     61,
	"gameNotFound":      101,
	"gameAlreadyStart":  103,
	"InvalidService":    1000,
}

func (x Code) String() string {
	return proto.EnumName(Code_name, int32(x))
}
func (Code) EnumDescriptor() ([]byte, []int) { return fileDescriptorResultcode, []int{0} }

func init() {
	proto.RegisterEnum("protocol.Code", Code_name, Code_value)
}

func init() { proto.RegisterFile("resultcode.proto", fileDescriptorResultcode) }

var fileDescriptorResultcode = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xd0, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x80, 0x61, 0x52, 0x2a, 0x0a, 0xd7, 0x42, 0xaf, 0x57, 0x58, 0xf3, 0x00, 0x0c, 0x0c, 0x74,
	0x85, 0x81, 0xb6, 0x20, 0x75, 0x00, 0x21, 0x22, 0xb1, 0x1b, 0xe7, 0x88, 0x22, 0xa5, 0x39, 0xe3,
	0x38, 0x55, 0xbb, 0xf2, 0x74, 0x8c, 0x3c, 0x02, 0xca, 0xc4, 0x63, 0x54, 0x97, 0x21, 0x9d, 0x2c,
	0x7d, 0xfa, 0x6d, 0xeb, 0x0e, 0xd0, 0x73, 0x55, 0x17, 0xc1, 0x4a, 0xca, 0x37, 0xce, 0x4b, 0x10,
	0x3a, 0x6d, 0x0f, 0x2b, 0xc5, 0xf5, 0x77, 0x0f, 0xfa, 0x0b, 0x49, 0x99, 0x86, 0x30, 0x48, 0x6a,
	0x6b, 0xb9, 0xaa, 0xf0, 0x88, 0x08, 0x2e, 0x56, 0xe5, 0xc6, 0x14, 0x79, 0xfa, 0xc6, 0x5f, 0x35,
	0x57, 0x01, 0x23, 0x9a, 0xc2, 0x38, 0x61, 0xbf, 0x61, 0xff, 0xb8, 0xb5, 0xec, 0x42, 0x2e, 0x25,
	0xf6, 0x68, 0x0c, 0xc3, 0xe5, 0x5c, 0x21, 0x38, 0x85, 0x63, 0xad, 0x16, 0x52, 0x7e, 0xe6, 0xd9,
	0xa1, 0xea, 0xd3, 0x15, 0x4c, 0xde, 0xf5, 0x31, 0x13, 0xf8, 0xc0, 0x67, 0x34, 0x81, 0x73, 0xe3,
	0xdc, 0x6a, 0xf9, 0x22, 0xe1, 0x49, 0xea, 0x32, 0xc5, 0x5b, 0xfd, 0xd8, 0x15, 0x66, 0xc7, 0xbe,
	0xb3, 0x19, 0x21, 0x8c, 0xbc, 0xc8, 0xba, 0x93, 0x3b, 0xbd, 0xa8, 0xf2, 0x6c, 0xb6, 0xaf, 0x6d,
	0x8c, 0xf7, 0x1a, 0x65, 0x66, 0xcd, 0x5d, 0xc4, 0x74, 0x09, 0xa8, 0xf2, 0x50, 0x78, 0x36, 0xe9,
	0x2e, 0x09, 0xc6, 0x07, 0xcc, 0x68, 0xda, 0x4d, 0xa6, 0xc3, 0xe4, 0x96, 0xf1, 0x7f, 0x30, 0x1f,
	0xfd, 0x34, 0x71, 0xf4, 0xdb, 0xc4, 0xd1, 0x5f, 0x13, 0x47, 0x1f, 0x27, 0xed, 0x72, 0x66, 0xfb,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xe6, 0xe5, 0x89, 0x37, 0x01, 0x00, 0x00,
}
