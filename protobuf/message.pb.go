// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/message.proto

package protobuf

import proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message_MessageType int32

const (
	Message_INVITE Message_MessageType = 0
	Message_JOIN   Message_MessageType = 1
	Message_LEAVE  Message_MessageType = 2
	Message_KICK   Message_MessageType = 3
	Message_ACK    Message_MessageType = 4
	Message_POST   Message_MessageType = 5
)

var Message_MessageType_name = map[int32]string{
	0: "INVITE",
	1: "JOIN",
	2: "LEAVE",
	3: "KICK",
	4: "ACK",
	5: "POST",
}
var Message_MessageType_value = map[string]int32{
	"INVITE": 0,
	"JOIN":   1,
	"LEAVE":  2,
	"KICK":   3,
	"ACK":    4,
	"POST":   5,
}

func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}
func (Message_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_fd151bea13cda030, []int{0, 0}
}

type Message struct {
	MessageType          Message_MessageType `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3,enum=mvds.Message_MessageType" json:"message_type,omitempty"`
	Body                 []byte              `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Signature            []byte              `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_fd151bea13cda030, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMessageType() Message_MessageType {
	if m != nil {
		return m.MessageType
	}
	return Message_INVITE
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Message) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "mvds.Message")
	proto.RegisterEnum("mvds.Message_MessageType", Message_MessageType_name, Message_MessageType_value)
}

func init() { proto.RegisterFile("protobuf/message.proto", fileDescriptor_message_fd151bea13cda030) }

var fileDescriptor_message_fd151bea13cda030 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x03, 0x0b, 0x08,
	0xb1, 0xe4, 0x96, 0xa5, 0x14, 0x2b, 0x1d, 0x63, 0xe4, 0x62, 0xf7, 0x85, 0x88, 0x0b, 0xd9, 0x70,
	0xf1, 0x40, 0x95, 0xc4, 0x97, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19, 0x49,
	0xea, 0x81, 0x14, 0xea, 0x41, 0x15, 0xc1, 0xe8, 0x90, 0xca, 0x82, 0xd4, 0x20, 0xee, 0x5c, 0x04,
	0x47, 0x48, 0x88, 0x8b, 0x25, 0x29, 0x3f, 0xa5, 0x52, 0x82, 0x49, 0x81, 0x51, 0x83, 0x27, 0x08,
	0xcc, 0x16, 0x92, 0xe1, 0xe2, 0x2c, 0xce, 0x4c, 0xcf, 0x4b, 0x2c, 0x29, 0x2d, 0x4a, 0x95, 0x60,
	0x06, 0x4b, 0x20, 0x04, 0x94, 0xbc, 0xb9, 0xb8, 0x91, 0x4c, 0x13, 0xe2, 0xe2, 0x62, 0xf3, 0xf4,
	0x0b, 0xf3, 0x0c, 0x71, 0x15, 0x60, 0x10, 0xe2, 0xe0, 0x62, 0xf1, 0xf2, 0xf7, 0xf4, 0x13, 0x60,
	0x14, 0xe2, 0xe4, 0x62, 0xf5, 0x71, 0x75, 0x0c, 0x73, 0x15, 0x60, 0x02, 0x09, 0x7a, 0x7b, 0x3a,
	0x7b, 0x0b, 0x30, 0x0b, 0xb1, 0x73, 0x31, 0x3b, 0x3a, 0x7b, 0x0b, 0xb0, 0x80, 0x84, 0x02, 0xfc,
	0x83, 0x43, 0x04, 0x58, 0x9d, 0xb8, 0xa2, 0x38, 0x60, 0x1e, 0x4d, 0x62, 0x03, 0xb3, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x89, 0xfa, 0x23, 0xe1, 0xfb, 0x00, 0x00, 0x00,
}
