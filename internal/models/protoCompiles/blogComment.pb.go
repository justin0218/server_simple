// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blogComment.proto

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

type BlogComment struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BlogId               int32    `protobuf:"varint,2,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime           string   `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           string   `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlogComment) Reset()         { *m = BlogComment{} }
func (m *BlogComment) String() string { return proto.CompactTextString(m) }
func (*BlogComment) ProtoMessage()    {}
func (*BlogComment) Descriptor() ([]byte, []int) {
	return fileDescriptor_5357f06a6fff0436, []int{0}
}

func (m *BlogComment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogComment.Unmarshal(m, b)
}
func (m *BlogComment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogComment.Marshal(b, m, deterministic)
}
func (m *BlogComment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogComment.Merge(m, src)
}
func (m *BlogComment) XXX_Size() int {
	return xxx_messageInfo_BlogComment.Size(m)
}
func (m *BlogComment) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogComment.DiscardUnknown(m)
}

var xxx_messageInfo_BlogComment proto.InternalMessageInfo

func (m *BlogComment) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BlogComment) GetBlogId() int32 {
	if m != nil {
		return m.BlogId
	}
	return 0
}

func (m *BlogComment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *BlogComment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BlogComment) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *BlogComment) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func init() {
	proto.RegisterType((*BlogComment)(nil), "blog_comment")
}

func init() { proto.RegisterFile("blogComment.proto", fileDescriptor_5357f06a6fff0436) }

var fileDescriptor_5357f06a6fff0436 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0x3d, 0x0e, 0xc2, 0x30,
	0x0c, 0x46, 0x95, 0xd2, 0x1f, 0x61, 0xfe, 0x44, 0x16, 0xb2, 0x51, 0x31, 0x75, 0x62, 0xe1, 0x06,
	0x74, 0x62, 0xad, 0x98, 0x58, 0xaa, 0xb6, 0xb1, 0x50, 0xa4, 0x26, 0xa9, 0x8a, 0x39, 0x13, 0xd7,
	0x44, 0x71, 0x04, 0x9b, 0xdf, 0xfb, 0xde, 0x60, 0xd8, 0xf7, 0xa3, 0x7f, 0xd6, 0xde, 0x5a, 0x74,
	0x74, 0x9e, 0x66, 0x4f, 0xfe, 0xf4, 0x11, 0xb0, 0x0e, 0xb6, 0x1d, 0xa2, 0x96, 0x5b, 0x48, 0x8c,
	0x56, 0xa2, 0x14, 0x55, 0xd6, 0x24, 0x46, 0xcb, 0x03, 0x14, 0xbc, 0x1b, 0xad, 0x12, 0x96, 0x79,
	0xc0, 0x9b, 0x96, 0x0a, 0x8a, 0xc1, 0x3b, 0x42, 0x47, 0x6a, 0x51, 0x8a, 0x6a, 0xd9, 0xfc, 0x50,
	0x4a, 0x48, 0x5d, 0x67, 0x51, 0xa5, 0xac, 0xf9, 0x96, 0x47, 0x58, 0x0d, 0x33, 0x76, 0x84, 0x2d,
	0x19, 0x8b, 0x2a, 0xe3, 0x09, 0xa2, 0xba, 0x9b, 0x18, 0xbc, 0x27, 0xfd, 0x0f, 0xf2, 0x18, 0x44,
	0x15, 0x82, 0xeb, 0xee, 0xb1, 0xe1, 0x97, 0x6b, 0x6f, 0x27, 0x33, 0xe2, 0xab, 0xcf, 0x19, 0x2f,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0xbd, 0xf3, 0x05, 0xd6, 0x00, 0x00, 0x00,
}
