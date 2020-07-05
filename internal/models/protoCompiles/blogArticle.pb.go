// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blogArticle.proto

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

type BlogArticle struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Cover                string   `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty"`
	GoodNum              int64    `protobuf:"varint,3,opt,name=good_num,json=goodNum,proto3" json:"good_num,omitempty"`
	View                 int64    `protobuf:"varint,4,opt,name=view,proto3" json:"view,omitempty"`
	Recommended          int64    `protobuf:"varint,5,opt,name=recommended,proto3" json:"recommended,omitempty"`
	Type                 int64    `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
	Preface              string   `protobuf:"bytes,7,opt,name=preface,proto3" json:"preface,omitempty"`
	HtmlTxtUrl           string   `protobuf:"bytes,8,opt,name=html_txt_url,json=htmlTxtUrl,proto3" json:"html_txt_url,omitempty"`
	Name                 string   `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime           string   `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           string   `protobuf:"bytes,11,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlogArticle) Reset()         { *m = BlogArticle{} }
func (m *BlogArticle) String() string { return proto.CompactTextString(m) }
func (*BlogArticle) ProtoMessage()    {}
func (*BlogArticle) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e0525f62164731, []int{0}
}

func (m *BlogArticle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogArticle.Unmarshal(m, b)
}
func (m *BlogArticle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogArticle.Marshal(b, m, deterministic)
}
func (m *BlogArticle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogArticle.Merge(m, src)
}
func (m *BlogArticle) XXX_Size() int {
	return xxx_messageInfo_BlogArticle.Size(m)
}
func (m *BlogArticle) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogArticle.DiscardUnknown(m)
}

var xxx_messageInfo_BlogArticle proto.InternalMessageInfo

func (m *BlogArticle) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BlogArticle) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *BlogArticle) GetGoodNum() int64 {
	if m != nil {
		return m.GoodNum
	}
	return 0
}

func (m *BlogArticle) GetView() int64 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *BlogArticle) GetRecommended() int64 {
	if m != nil {
		return m.Recommended
	}
	return 0
}

func (m *BlogArticle) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *BlogArticle) GetPreface() string {
	if m != nil {
		return m.Preface
	}
	return ""
}

func (m *BlogArticle) GetHtmlTxtUrl() string {
	if m != nil {
		return m.HtmlTxtUrl
	}
	return ""
}

func (m *BlogArticle) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BlogArticle) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *BlogArticle) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

type BlogRoom struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cover                string   `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
	CreateTime           string   `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Olnum                int32    `protobuf:"varint,6,opt,name=olnum,proto3" json:"olnum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlogRoom) Reset()         { *m = BlogRoom{} }
func (m *BlogRoom) String() string { return proto.CompactTextString(m) }
func (*BlogRoom) ProtoMessage()    {}
func (*BlogRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e0525f62164731, []int{1}
}

func (m *BlogRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogRoom.Unmarshal(m, b)
}
func (m *BlogRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogRoom.Marshal(b, m, deterministic)
}
func (m *BlogRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogRoom.Merge(m, src)
}
func (m *BlogRoom) XXX_Size() int {
	return xxx_messageInfo_BlogRoom.Size(m)
}
func (m *BlogRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogRoom.DiscardUnknown(m)
}

var xxx_messageInfo_BlogRoom proto.InternalMessageInfo

func (m *BlogRoom) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BlogRoom) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BlogRoom) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *BlogRoom) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *BlogRoom) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *BlogRoom) GetOlnum() int32 {
	if m != nil {
		return m.Olnum
	}
	return 0
}

func init() {
	proto.RegisterType((*BlogArticle)(nil), "blog_article")
	proto.RegisterType((*BlogRoom)(nil), "blog_room")
}

func init() { proto.RegisterFile("blogArticle.proto", fileDescriptor_28e0525f62164731) }

var fileDescriptor_28e0525f62164731 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x95, 0x34, 0x7f, 0x9a, 0x4b, 0x01, 0x61, 0x31, 0x98, 0x89, 0xa8, 0x53, 0x26, 0x16,
	0x9e, 0x00, 0xd8, 0x19, 0xa2, 0xb2, 0xb0, 0x44, 0x69, 0x7c, 0x14, 0x4b, 0x76, 0x1c, 0x19, 0xa7,
	0x94, 0x87, 0xe1, 0x09, 0x78, 0x49, 0xe4, 0xb3, 0x5a, 0x45, 0x5d, 0xd8, 0xee, 0xfb, 0xdd, 0x77,
	0xbe, 0xd3, 0x67, 0xb8, 0xde, 0x2a, 0xb3, 0x7b, 0xb4, 0x4e, 0xf6, 0x0a, 0xef, 0x47, 0x6b, 0x9c,
	0x59, 0xff, 0xc6, 0xb0, 0xf2, 0xb4, 0xed, 0x02, 0x66, 0x97, 0x10, 0x4b, 0xc1, 0xa3, 0x2a, 0xaa,
	0xd3, 0x26, 0x96, 0x82, 0xdd, 0x40, 0xda, 0x9b, 0x3d, 0x5a, 0x1e, 0x57, 0x51, 0x5d, 0x34, 0x41,
	0xb0, 0x5b, 0x58, 0xee, 0x8c, 0x11, 0xed, 0x30, 0x69, 0xbe, 0xa8, 0xa2, 0x7a, 0xd1, 0xe4, 0x5e,
	0xbf, 0x4c, 0x9a, 0x31, 0x48, 0xf6, 0x12, 0xbf, 0x78, 0x42, 0x98, 0x6a, 0x56, 0x41, 0x69, 0xb1,
	0x37, 0x5a, 0xe3, 0x20, 0x50, 0xf0, 0x94, 0x5a, 0x73, 0xe4, 0xa7, 0xdc, 0xf7, 0x88, 0x3c, 0x0b,
	0x53, 0xbe, 0x66, 0x1c, 0xf2, 0xd1, 0xe2, 0x7b, 0xd7, 0x23, 0xcf, 0x69, 0xf9, 0x51, 0xb2, 0x0a,
	0x56, 0x1f, 0x4e, 0xab, 0xd6, 0x1d, 0x5c, 0x3b, 0x59, 0xc5, 0x97, 0xd4, 0x06, 0xcf, 0x36, 0x07,
	0xf7, 0x6a, 0x95, 0x7f, 0x6f, 0xe8, 0x34, 0xf2, 0x82, 0x3a, 0x54, 0xb3, 0x3b, 0x28, 0x7b, 0x8b,
	0x9d, 0xc3, 0xd6, 0x49, 0x8d, 0x1c, 0xc2, 0x50, 0x40, 0x1b, 0x19, 0x0c, 0xd3, 0x28, 0x4e, 0x86,
	0x32, 0x18, 0x02, 0xf2, 0x86, 0xf5, 0x4f, 0x04, 0x05, 0xa5, 0x65, 0x8d, 0xd1, 0xb3, 0xa8, 0x0a,
	0x8a, 0xea, 0xb8, 0x33, 0x9e, 0xed, 0x3c, 0xc5, 0xb7, 0x98, 0xc7, 0x77, 0x76, 0x49, 0xf2, 0xdf,
	0x25, 0xe9, 0xf9, 0x25, 0xfe, 0x5d, 0xa3, 0x7c, 0xfa, 0x19, 0xfd, 0x54, 0x10, 0x4f, 0x57, 0x6f,
	0x17, 0xf4, 0xad, 0xcf, 0x46, 0x8f, 0x52, 0xe1, 0xe7, 0x36, 0x23, 0xf9, 0xf0, 0x17, 0x00, 0x00,
	0xff, 0xff, 0x56, 0x43, 0x38, 0xeb, 0xfa, 0x01, 0x00, 0x00,
}
