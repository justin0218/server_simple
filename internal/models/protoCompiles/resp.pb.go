// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resp.proto

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

type BlogList struct {
	Code                 ErrorCodes     `protobuf:"varint,1,opt,name=code,proto3,enum=ErrorCodes" json:"code,omitempty"`
	Msg                  string         `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	BlogArticleList      []*BlogArticle `protobuf:"bytes,3,rep,name=blog_article_list,json=blogArticleList,proto3" json:"blog_article_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BlogList) Reset()         { *m = BlogList{} }
func (m *BlogList) String() string { return proto.CompactTextString(m) }
func (*BlogList) ProtoMessage()    {}
func (*BlogList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5365792f61ddff, []int{0}
}

func (m *BlogList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogList.Unmarshal(m, b)
}
func (m *BlogList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogList.Marshal(b, m, deterministic)
}
func (m *BlogList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogList.Merge(m, src)
}
func (m *BlogList) XXX_Size() int {
	return xxx_messageInfo_BlogList.Size(m)
}
func (m *BlogList) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogList.DiscardUnknown(m)
}

var xxx_messageInfo_BlogList proto.InternalMessageInfo

func (m *BlogList) GetCode() ErrorCodes {
	if m != nil {
		return m.Code
	}
	return ErrorCodes_F
}

func (m *BlogList) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *BlogList) GetBlogArticleList() []*BlogArticle {
	if m != nil {
		return m.BlogArticleList
	}
	return nil
}

type BlogDetail struct {
	Code                 ErrorCodes   `protobuf:"varint,1,opt,name=code,proto3,enum=ErrorCodes" json:"code,omitempty"`
	Msg                  string       `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	CurrentArticle       *BlogArticle `protobuf:"bytes,3,opt,name=current_article,json=currentArticle,proto3" json:"current_article,omitempty"`
	NextArticle          *BlogArticle `protobuf:"bytes,4,opt,name=next_article,json=nextArticle,proto3" json:"next_article,omitempty"`
	PrevArticle          *BlogArticle `protobuf:"bytes,5,opt,name=prev_article,json=prevArticle,proto3" json:"prev_article,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BlogDetail) Reset()         { *m = BlogDetail{} }
func (m *BlogDetail) String() string { return proto.CompactTextString(m) }
func (*BlogDetail) ProtoMessage()    {}
func (*BlogDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5365792f61ddff, []int{1}
}

func (m *BlogDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogDetail.Unmarshal(m, b)
}
func (m *BlogDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogDetail.Marshal(b, m, deterministic)
}
func (m *BlogDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogDetail.Merge(m, src)
}
func (m *BlogDetail) XXX_Size() int {
	return xxx_messageInfo_BlogDetail.Size(m)
}
func (m *BlogDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogDetail.DiscardUnknown(m)
}

var xxx_messageInfo_BlogDetail proto.InternalMessageInfo

func (m *BlogDetail) GetCode() ErrorCodes {
	if m != nil {
		return m.Code
	}
	return ErrorCodes_F
}

func (m *BlogDetail) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *BlogDetail) GetCurrentArticle() *BlogArticle {
	if m != nil {
		return m.CurrentArticle
	}
	return nil
}

func (m *BlogDetail) GetNextArticle() *BlogArticle {
	if m != nil {
		return m.NextArticle
	}
	return nil
}

func (m *BlogDetail) GetPrevArticle() *BlogArticle {
	if m != nil {
		return m.PrevArticle
	}
	return nil
}

type BlogTypes struct {
	Code                 ErrorCodes  `protobuf:"varint,1,opt,name=code,proto3,enum=ErrorCodes" json:"code,omitempty"`
	Msg                  string      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	BlogTypeList         []*BlogType `protobuf:"bytes,3,rep,name=blog_type_list,json=blogTypeList,proto3" json:"blog_type_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BlogTypes) Reset()         { *m = BlogTypes{} }
func (m *BlogTypes) String() string { return proto.CompactTextString(m) }
func (*BlogTypes) ProtoMessage()    {}
func (*BlogTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5365792f61ddff, []int{2}
}

func (m *BlogTypes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogTypes.Unmarshal(m, b)
}
func (m *BlogTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogTypes.Marshal(b, m, deterministic)
}
func (m *BlogTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogTypes.Merge(m, src)
}
func (m *BlogTypes) XXX_Size() int {
	return xxx_messageInfo_BlogTypes.Size(m)
}
func (m *BlogTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogTypes.DiscardUnknown(m)
}

var xxx_messageInfo_BlogTypes proto.InternalMessageInfo

func (m *BlogTypes) GetCode() ErrorCodes {
	if m != nil {
		return m.Code
	}
	return ErrorCodes_F
}

func (m *BlogTypes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *BlogTypes) GetBlogTypeList() []*BlogType {
	if m != nil {
		return m.BlogTypeList
	}
	return nil
}

type BlogComments struct {
	Code                 ErrorCodes     `protobuf:"varint,1,opt,name=code,proto3,enum=ErrorCodes" json:"code,omitempty"`
	Msg                  string         `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	BlogCommentList      []*BlogComment `protobuf:"bytes,3,rep,name=blog_comment_list,json=blogCommentList,proto3" json:"blog_comment_list,omitempty"`
	Total                int32          `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BlogComments) Reset()         { *m = BlogComments{} }
func (m *BlogComments) String() string { return proto.CompactTextString(m) }
func (*BlogComments) ProtoMessage()    {}
func (*BlogComments) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5365792f61ddff, []int{3}
}

func (m *BlogComments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogComments.Unmarshal(m, b)
}
func (m *BlogComments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogComments.Marshal(b, m, deterministic)
}
func (m *BlogComments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogComments.Merge(m, src)
}
func (m *BlogComments) XXX_Size() int {
	return xxx_messageInfo_BlogComments.Size(m)
}
func (m *BlogComments) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogComments.DiscardUnknown(m)
}

var xxx_messageInfo_BlogComments proto.InternalMessageInfo

func (m *BlogComments) GetCode() ErrorCodes {
	if m != nil {
		return m.Code
	}
	return ErrorCodes_F
}

func (m *BlogComments) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *BlogComments) GetBlogCommentList() []*BlogComment {
	if m != nil {
		return m.BlogCommentList
	}
	return nil
}

func (m *BlogComments) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type FileContent struct {
	Code                 ErrorCodes `protobuf:"varint,1,opt,name=code,proto3,enum=ErrorCodes" json:"code,omitempty"`
	Msg                  string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Txt                  string     `protobuf:"bytes,3,opt,name=txt,proto3" json:"txt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FileContent) Reset()         { *m = FileContent{} }
func (m *FileContent) String() string { return proto.CompactTextString(m) }
func (*FileContent) ProtoMessage()    {}
func (*FileContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5365792f61ddff, []int{4}
}

func (m *FileContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileContent.Unmarshal(m, b)
}
func (m *FileContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileContent.Marshal(b, m, deterministic)
}
func (m *FileContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileContent.Merge(m, src)
}
func (m *FileContent) XXX_Size() int {
	return xxx_messageInfo_FileContent.Size(m)
}
func (m *FileContent) XXX_DiscardUnknown() {
	xxx_messageInfo_FileContent.DiscardUnknown(m)
}

var xxx_messageInfo_FileContent proto.InternalMessageInfo

func (m *FileContent) GetCode() ErrorCodes {
	if m != nil {
		return m.Code
	}
	return ErrorCodes_F
}

func (m *FileContent) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *FileContent) GetTxt() string {
	if m != nil {
		return m.Txt
	}
	return ""
}

type BlogChatRecords struct {
	Code                 ErrorCodes        `protobuf:"varint,1,opt,name=code,proto3,enum=ErrorCodes" json:"code,omitempty"`
	Msg                  string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	BlogChatRecordList   []*BlogChatRecord `protobuf:"bytes,3,rep,name=blog_chat_record_list,json=blogChatRecordList,proto3" json:"blog_chat_record_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BlogChatRecords) Reset()         { *m = BlogChatRecords{} }
func (m *BlogChatRecords) String() string { return proto.CompactTextString(m) }
func (*BlogChatRecords) ProtoMessage()    {}
func (*BlogChatRecords) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5365792f61ddff, []int{5}
}

func (m *BlogChatRecords) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogChatRecords.Unmarshal(m, b)
}
func (m *BlogChatRecords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogChatRecords.Marshal(b, m, deterministic)
}
func (m *BlogChatRecords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogChatRecords.Merge(m, src)
}
func (m *BlogChatRecords) XXX_Size() int {
	return xxx_messageInfo_BlogChatRecords.Size(m)
}
func (m *BlogChatRecords) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogChatRecords.DiscardUnknown(m)
}

var xxx_messageInfo_BlogChatRecords proto.InternalMessageInfo

func (m *BlogChatRecords) GetCode() ErrorCodes {
	if m != nil {
		return m.Code
	}
	return ErrorCodes_F
}

func (m *BlogChatRecords) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *BlogChatRecords) GetBlogChatRecordList() []*BlogChatRecord {
	if m != nil {
		return m.BlogChatRecordList
	}
	return nil
}

type BlogRooms struct {
	Code                 ErrorCodes  `protobuf:"varint,1,opt,name=code,proto3,enum=ErrorCodes" json:"code,omitempty"`
	Msg                  string      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	BlogRoomList         []*BlogRoom `protobuf:"bytes,3,rep,name=blog_room_list,json=blogRoomList,proto3" json:"blog_room_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BlogRooms) Reset()         { *m = BlogRooms{} }
func (m *BlogRooms) String() string { return proto.CompactTextString(m) }
func (*BlogRooms) ProtoMessage()    {}
func (*BlogRooms) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5365792f61ddff, []int{6}
}

func (m *BlogRooms) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogRooms.Unmarshal(m, b)
}
func (m *BlogRooms) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogRooms.Marshal(b, m, deterministic)
}
func (m *BlogRooms) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogRooms.Merge(m, src)
}
func (m *BlogRooms) XXX_Size() int {
	return xxx_messageInfo_BlogRooms.Size(m)
}
func (m *BlogRooms) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogRooms.DiscardUnknown(m)
}

var xxx_messageInfo_BlogRooms proto.InternalMessageInfo

func (m *BlogRooms) GetCode() ErrorCodes {
	if m != nil {
		return m.Code
	}
	return ErrorCodes_F
}

func (m *BlogRooms) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *BlogRooms) GetBlogRoomList() []*BlogRoom {
	if m != nil {
		return m.BlogRoomList
	}
	return nil
}

func init() {
	proto.RegisterType((*BlogList)(nil), "blog_list")
	proto.RegisterType((*BlogDetail)(nil), "blog_detail")
	proto.RegisterType((*BlogTypes)(nil), "blog_types")
	proto.RegisterType((*BlogComments)(nil), "blog_comments")
	proto.RegisterType((*FileContent)(nil), "file_content")
	proto.RegisterType((*BlogChatRecords)(nil), "blog_chat_records")
	proto.RegisterType((*BlogRooms)(nil), "blog_rooms")
}

func init() { proto.RegisterFile("resp.proto", fileDescriptor_3c5365792f61ddff) }

var fileDescriptor_3c5365792f61ddff = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xb1, 0x6e, 0xe2, 0x40,
	0x10, 0x95, 0xcf, 0x70, 0x12, 0x83, 0x31, 0x60, 0x71, 0x92, 0x75, 0x95, 0xe5, 0x8a, 0xca, 0x42,
	0x9c, 0x74, 0xd2, 0x95, 0x77, 0x5c, 0x99, 0x6a, 0x45, 0x95, 0xc6, 0x32, 0x66, 0x03, 0x96, 0x6c,
	0xaf, 0xb5, 0x3b, 0x89, 0x20, 0xff, 0x90, 0x2e, 0xbf, 0x97, 0x7f, 0x89, 0x76, 0xbc, 0x26, 0x0e,
	0x71, 0x65, 0xba, 0xdd, 0x37, 0xf3, 0xe6, 0xed, 0xcc, 0x9b, 0x05, 0x90, 0x5c, 0x55, 0x51, 0x25,
	0x05, 0x8a, 0x9f, 0xf3, 0x5d, 0x2e, 0x0e, 0x7f, 0x25, 0x66, 0x69, 0xce, 0x0d, 0xe4, 0x70, 0x29,
	0x85, 0x54, 0xe6, 0xe6, 0xea, 0x84, 0xed, 0xb9, 0xe2, 0x6d, 0xc2, 0x46, 0x14, 0x05, 0x2f, 0xd1,
	0x40, 0x0b, 0x82, 0x8e, 0x09, 0x32, 0x9e, 0x0a, 0xb9, 0xaf, 0xd1, 0xf0, 0x19, 0x46, 0x1a, 0x8f,
	0xf3, 0x4c, 0xa1, 0x17, 0xc0, 0x20, 0x15, 0x7b, 0xee, 0x5b, 0x81, 0xb5, 0x74, 0xd7, 0x4e, 0x44,
	0x12, 0xb1, 0x86, 0x14, 0xa3, 0x88, 0x37, 0x03, 0xbb, 0x50, 0x07, 0xff, 0x5b, 0x60, 0x2d, 0x47,
	0x4c, 0x1f, 0xbd, 0x3f, 0x40, 0x5a, 0x71, 0x52, 0xbf, 0x8e, 0x0a, 0xf9, 0x76, 0x60, 0x2f, 0xc7,
	0xeb, 0x49, 0xd4, 0x8e, 0xb0, 0x69, 0xab, 0x89, 0xbb, 0x4c, 0x61, 0xf8, 0x66, 0xc1, 0x98, 0x32,
	0xf6, 0x1c, 0x93, 0x2c, 0xef, 0x25, 0xff, 0x1b, 0xa6, 0xe9, 0xa3, 0x94, 0xbc, 0xc4, 0x46, 0xc7,
	0xb7, 0x03, 0xeb, 0xab, 0xb8, 0x6b, 0xb2, 0x8c, 0xbe, 0xb7, 0x02, 0xa7, 0xe4, 0xa7, 0x0f, 0xd2,
	0xa0, 0x8b, 0x34, 0xd6, 0x29, 0x2d, 0x46, 0x25, 0xf9, 0xd3, 0x85, 0x31, 0xec, 0x64, 0xe8, 0x14,
	0xc3, 0x08, 0x25, 0x00, 0x05, 0xf1, 0x5c, 0x71, 0xd5, 0xab, 0xbb, 0x15, 0xb8, 0x97, 0x0a, 0xed,
	0xc9, 0x42, 0x74, 0x81, 0x99, 0xd3, 0x58, 0x4f, 0x33, 0x7d, 0xb5, 0x60, 0x42, 0xb1, 0xb4, 0x36,
	0x5f, 0xdd, 0x64, 0xaa, 0x29, 0xd2, 0x61, 0xaa, 0x89, 0xd4, 0xa6, 0x9a, 0x45, 0xd3, 0x0f, 0xf0,
	0x16, 0x30, 0x44, 0x81, 0x49, 0x4e, 0x13, 0x1d, 0xb2, 0xfa, 0x12, 0x6e, 0xc1, 0x79, 0xc8, 0x72,
	0x1e, 0xa7, 0xa2, 0x44, 0x5e, 0xf6, 0xdb, 0xb4, 0x19, 0xd8, 0x78, 0x42, 0xb2, 0x77, 0xc4, 0xf4,
	0x31, 0x7c, 0xb1, 0x9a, 0x77, 0x1e, 0x13, 0x8c, 0x25, 0xed, 0x75, 0xbf, 0x86, 0xff, 0xc3, 0x8f,
	0xeb, 0x42, 0xed, 0xa6, 0xe7, 0xd1, 0x75, 0x94, 0x79, 0x9f, 0xbf, 0x13, 0x0d, 0xbf, 0x31, 0x5c,
	0x0a, 0x51, 0xdc, 0x66, 0xb8, 0xae, 0xd0, 0x61, 0xb8, 0x86, 0x6b, 0xc3, 0x99, 0x10, 0x85, 0xd6,
	0xfc, 0x37, 0xbd, 0x9f, 0xd0, 0x4f, 0xde, 0x88, 0xa2, 0xca, 0x72, 0xae, 0x76, 0xdf, 0xe9, 0xfa,
	0xeb, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x73, 0xef, 0x78, 0xa5, 0x40, 0x04, 0x00, 0x00,
}
