// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ws.proto

package protoCompiles

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Events int32

const (
	Events_CHAT_CONTENT     Events = 0
	Events_CHAT_RECORDS     Events = 1
	Events_GDRAW_CONTENT    Events = 2
	Events_GDRAW_ROOMCREATE Events = 3
)

var Events_name = map[int32]string{
	0: "CHAT_CONTENT",
	1: "CHAT_RECORDS",
	2: "GDRAW_CONTENT",
	3: "GDRAW_ROOMCREATE",
}

var Events_value = map[string]int32{
	"CHAT_CONTENT":     0,
	"CHAT_RECORDS":     1,
	"GDRAW_CONTENT":    2,
	"GDRAW_ROOMCREATE": 3,
}

func (x Events) String() string {
	return proto.EnumName(Events_name, int32(x))
}

func (Events) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7ff87931dac4ca82, []int{0}
}

type WsMsgBase struct {
	Event                Events   `protobuf:"varint,1,opt,name=event,proto3,enum=Events" json:"event,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WsMsgBase) Reset()         { *m = WsMsgBase{} }
func (m *WsMsgBase) String() string { return proto.CompactTextString(m) }
func (*WsMsgBase) ProtoMessage()    {}
func (*WsMsgBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ff87931dac4ca82, []int{0}
}

func (m *WsMsgBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WsMsgBase.Unmarshal(m, b)
}
func (m *WsMsgBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WsMsgBase.Marshal(b, m, deterministic)
}
func (m *WsMsgBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WsMsgBase.Merge(m, src)
}
func (m *WsMsgBase) XXX_Size() int {
	return xxx_messageInfo_WsMsgBase.Size(m)
}
func (m *WsMsgBase) XXX_DiscardUnknown() {
	xxx_messageInfo_WsMsgBase.DiscardUnknown(m)
}

var xxx_messageInfo_WsMsgBase proto.InternalMessageInfo

func (m *WsMsgBase) GetEvent() Events {
	if m != nil {
		return m.Event
	}
	return Events_CHAT_CONTENT
}

func (m *WsMsgBase) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ChatContent struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	TimeString           string   `protobuf:"bytes,2,opt,name=time_string,json=timeString,proto3" json:"time_string,omitempty"`
	Timer                int32    `protobuf:"varint,3,opt,name=timer,proto3" json:"timer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatContent) Reset()         { *m = ChatContent{} }
func (m *ChatContent) String() string { return proto.CompactTextString(m) }
func (*ChatContent) ProtoMessage()    {}
func (*ChatContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ff87931dac4ca82, []int{1}
}

func (m *ChatContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatContent.Unmarshal(m, b)
}
func (m *ChatContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatContent.Marshal(b, m, deterministic)
}
func (m *ChatContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatContent.Merge(m, src)
}
func (m *ChatContent) XXX_Size() int {
	return xxx_messageInfo_ChatContent.Size(m)
}
func (m *ChatContent) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatContent.DiscardUnknown(m)
}

var xxx_messageInfo_ChatContent proto.InternalMessageInfo

func (m *ChatContent) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ChatContent) GetTimeString() string {
	if m != nil {
		return m.TimeString
	}
	return ""
}

func (m *ChatContent) GetTimer() int32 {
	if m != nil {
		return m.Timer
	}
	return 0
}

type GdrawContent struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Color                string   `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	Size                 int32    `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Uid                  int64    `protobuf:"varint,5,opt,name=uid,proto3" json:"uid,omitempty"`
	RoomId               string   `protobuf:"bytes,6,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GdrawContent) Reset()         { *m = GdrawContent{} }
func (m *GdrawContent) String() string { return proto.CompactTextString(m) }
func (*GdrawContent) ProtoMessage()    {}
func (*GdrawContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ff87931dac4ca82, []int{2}
}

func (m *GdrawContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GdrawContent.Unmarshal(m, b)
}
func (m *GdrawContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GdrawContent.Marshal(b, m, deterministic)
}
func (m *GdrawContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GdrawContent.Merge(m, src)
}
func (m *GdrawContent) XXX_Size() int {
	return xxx_messageInfo_GdrawContent.Size(m)
}
func (m *GdrawContent) XXX_DiscardUnknown() {
	xxx_messageInfo_GdrawContent.DiscardUnknown(m)
}

var xxx_messageInfo_GdrawContent proto.InternalMessageInfo

func (m *GdrawContent) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *GdrawContent) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *GdrawContent) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *GdrawContent) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *GdrawContent) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *GdrawContent) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func init() {
	proto.RegisterEnum("Events", Events_name, Events_value)
	proto.RegisterType((*WsMsgBase)(nil), "ws_msg_base")
	proto.RegisterType((*ChatContent)(nil), "chat_content")
	proto.RegisterType((*GdrawContent)(nil), "gdraw_content")
}

func init() { proto.RegisterFile("ws.proto", fileDescriptor_7ff87931dac4ca82) }

var fileDescriptor_7ff87931dac4ca82 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0xbf, 0x6d, 0x9a, 0xf4, 0xeb, 0x34, 0xd5, 0xb8, 0x14, 0xcc, 0x45, 0x0c, 0x3d, 0x05,
	0x0f, 0x3d, 0xe8, 0x1f, 0xb0, 0xa6, 0x41, 0x3d, 0xd8, 0xc0, 0x34, 0x52, 0xf0, 0x12, 0xd2, 0x26,
	0xc4, 0x40, 0xd3, 0x2d, 0xd9, 0xd5, 0xb6, 0x9e, 0xfc, 0xe9, 0xb2, 0xb3, 0x90, 0xdb, 0xfb, 0xbc,
	0xcc, 0x3c, 0x0c, 0x03, 0xff, 0x8f, 0x72, 0x76, 0x68, 0x85, 0x12, 0xd3, 0x47, 0x18, 0x1d, 0x65,
	0xd6, 0xc8, 0x2a, 0xdb, 0xe4, 0xb2, 0xe4, 0x37, 0x60, 0x97, 0xdf, 0xe5, 0x5e, 0xf9, 0x2c, 0x60,
	0xe1, 0xc5, 0xfd, 0x60, 0x46, 0x24, 0xd1, 0xb4, 0x9c, 0x43, 0xbf, 0xc8, 0x55, 0xee, 0xf7, 0x02,
	0x16, 0xba, 0x48, 0x79, 0xba, 0x06, 0x77, 0xfb, 0x99, 0xab, 0x6c, 0x2b, 0xf6, 0x4a, 0xcf, 0x78,
	0x60, 0x35, 0xb2, 0x22, 0xc1, 0x10, 0x75, 0xe4, 0xb7, 0x30, 0x52, 0x75, 0x53, 0x66, 0x52, 0xb5,
	0xf5, 0xbe, 0xa2, 0xe5, 0x21, 0x82, 0xae, 0x56, 0xd4, 0xf0, 0x09, 0xd8, 0x9a, 0x5a, 0xdf, 0x0a,
	0x58, 0x68, 0xa3, 0x81, 0xe9, 0x2f, 0x83, 0x71, 0x55, 0xb4, 0xf9, 0xb1, 0x53, 0xbb, 0xc0, 0x4e,
	0x24, 0xb6, 0x91, 0x9d, 0x34, 0x9d, 0x49, 0x66, 0x23, 0x3b, 0x6b, 0xc7, 0x56, 0xec, 0x84, 0x71,
	0x0c, 0xd1, 0x80, 0x3e, 0x58, 0xd6, 0x3f, 0xa5, 0xdf, 0xa7, 0x31, 0xca, 0xfa, 0xc0, 0xaf, 0xba,
	0xf0, 0xed, 0x80, 0x85, 0x16, 0xea, 0xc8, 0xaf, 0x61, 0xd0, 0x0a, 0xd1, 0x64, 0x75, 0xe1, 0x3b,
	0xb4, 0xed, 0x68, 0x7c, 0x2d, 0xee, 0xde, 0xc1, 0x31, 0x0f, 0xe0, 0x1e, 0xb8, 0xd1, 0xcb, 0x3c,
	0xcd, 0xa2, 0x64, 0x99, 0xc6, 0xcb, 0xd4, 0xfb, 0xd7, 0x35, 0x18, 0x47, 0x09, 0x2e, 0x56, 0x1e,
	0xe3, 0x57, 0x30, 0x7e, 0x5e, 0xe0, 0x7c, 0xdd, 0x0d, 0xf5, 0xf8, 0x04, 0x3c, 0x53, 0x61, 0x92,
	0xbc, 0x45, 0x18, 0xcf, 0xd3, 0xd8, 0xb3, 0x9e, 0x2e, 0x3f, 0xc6, 0xf4, 0xfd, 0x48, 0x34, 0x87,
	0x7a, 0x57, 0xca, 0x8d, 0x43, 0xf8, 0xf0, 0x17, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x3d, 0xf7, 0x9b,
	0x98, 0x01, 0x00, 0x00,
}
