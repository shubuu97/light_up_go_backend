// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package proto

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

type UserTypes int32

const (
	UserTypes_LIGHTER      UserTypes = 0
	UserTypes_LIGHT_SEEKER UserTypes = 1
	UserTypes_ADMIN        UserTypes = 2
)

var UserTypes_name = map[int32]string{
	0: "LIGHTER",
	1: "LIGHT_SEEKER",
	2: "ADMIN",
}

var UserTypes_value = map[string]int32{
	"LIGHTER":      0,
	"LIGHT_SEEKER": 1,
	"ADMIN":        2,
}

func (x UserTypes) String() string {
	return proto.EnumName(UserTypes_name, int32(x))
}

func (UserTypes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type IdRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdRequest) Reset()         { *m = IdRequest{} }
func (m *IdRequest) String() string { return proto.CompactTextString(m) }
func (*IdRequest) ProtoMessage()    {}
func (*IdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *IdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdRequest.Unmarshal(m, b)
}
func (m *IdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdRequest.Marshal(b, m, deterministic)
}
func (m *IdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdRequest.Merge(m, src)
}
func (m *IdRequest) XXX_Size() int {
	return xxx_messageInfo_IdRequest.Size(m)
}
func (m *IdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdRequest proto.InternalMessageInfo

func (m *IdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type EmailRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailRequest) Reset()         { *m = EmailRequest{} }
func (m *EmailRequest) String() string { return proto.CompactTextString(m) }
func (*EmailRequest) ProtoMessage()    {}
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *EmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailRequest.Unmarshal(m, b)
}
func (m *EmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailRequest.Marshal(b, m, deterministic)
}
func (m *EmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailRequest.Merge(m, src)
}
func (m *EmailRequest) XXX_Size() int {
	return xxx_messageInfo_EmailRequest.Size(m)
}
func (m *EmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmailRequest proto.InternalMessageInfo

func (m *EmailRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type IdsRequest struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdsRequest) Reset()         { *m = IdsRequest{} }
func (m *IdsRequest) String() string { return proto.CompactTextString(m) }
func (*IdsRequest) ProtoMessage()    {}
func (*IdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *IdsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdsRequest.Unmarshal(m, b)
}
func (m *IdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdsRequest.Marshal(b, m, deterministic)
}
func (m *IdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdsRequest.Merge(m, src)
}
func (m *IdsRequest) XXX_Size() int {
	return xxx_messageInfo_IdsRequest.Size(m)
}
func (m *IdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdsRequest proto.InternalMessageInfo

func (m *IdsRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("proto.UserTypes", UserTypes_name, UserTypes_value)
	proto.RegisterType((*IdRequest)(nil), "proto.IdRequest")
	proto.RegisterType((*EmailRequest)(nil), "proto.EmailRequest")
	proto.RegisterType((*IdsRequest)(nil), "proto.IdsRequest")
	proto.RegisterType((*Empty)(nil), "proto.Empty")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xd2, 0x5c, 0x9c,
	0x9e, 0x29, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x4c, 0x99, 0x29, 0x4a, 0x2a, 0x5c, 0x3c, 0xae, 0xb9, 0x89,
	0x99, 0x39, 0x30, 0x79, 0x11, 0x2e, 0xd6, 0x54, 0x10, 0x1f, 0xaa, 0x04, 0xc2, 0x51, 0x92, 0xe3,
	0xe2, 0xf2, 0x4c, 0x29, 0x86, 0xa9, 0x11, 0xe0, 0x62, 0xce, 0x4c, 0x29, 0x96, 0x60, 0x54, 0x60,
	0xd6, 0xe0, 0x0c, 0x02, 0x31, 0x95, 0xd8, 0xb9, 0x58, 0x5d, 0x73, 0x0b, 0x4a, 0x2a, 0xb5, 0x4c,
	0xb9, 0x38, 0x43, 0x8b, 0x53, 0x8b, 0x42, 0x2a, 0x0b, 0x52, 0x8b, 0x85, 0xb8, 0xb9, 0xd8, 0x7d,
	0x3c, 0xdd, 0x3d, 0x42, 0x5c, 0x83, 0x04, 0x18, 0x84, 0x04, 0xb8, 0x78, 0xc0, 0x9c, 0xf8, 0x60,
	0x57, 0x57, 0x6f, 0xd7, 0x20, 0x01, 0x46, 0x21, 0x4e, 0x2e, 0x56, 0x47, 0x17, 0x5f, 0x4f, 0x3f,
	0x01, 0x26, 0x27, 0xf9, 0x28, 0xd9, 0x9c, 0xcc, 0xf4, 0x8c, 0x12, 0xdd, 0xd2, 0x02, 0xdd, 0xa4,
	0xc4, 0xe4, 0xec, 0xd4, 0xbc, 0x14, 0x7d, 0x88, 0x57, 0xf4, 0xc1, 0x7e, 0x48, 0x62, 0x03, 0x53,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x80, 0x00, 0x9f, 0x0c, 0xe1, 0x00, 0x00, 0x00,
}