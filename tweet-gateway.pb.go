// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tweet-gateway.proto

package tweet_gateway

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

type TweetRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TweetRequest) Reset()         { *m = TweetRequest{} }
func (m *TweetRequest) String() string { return proto.CompactTextString(m) }
func (*TweetRequest) ProtoMessage()    {}
func (*TweetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2c6e576d127c43, []int{0}
}

func (m *TweetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TweetRequest.Unmarshal(m, b)
}
func (m *TweetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TweetRequest.Marshal(b, m, deterministic)
}
func (m *TweetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TweetRequest.Merge(m, src)
}
func (m *TweetRequest) XXX_Size() int {
	return xxx_messageInfo_TweetRequest.Size(m)
}
func (m *TweetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TweetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TweetRequest proto.InternalMessageInfo

func (m *TweetRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *TweetRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type TweetResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TweetResponse) Reset()         { *m = TweetResponse{} }
func (m *TweetResponse) String() string { return proto.CompactTextString(m) }
func (*TweetResponse) ProtoMessage()    {}
func (*TweetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2c6e576d127c43, []int{1}
}

func (m *TweetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TweetResponse.Unmarshal(m, b)
}
func (m *TweetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TweetResponse.Marshal(b, m, deterministic)
}
func (m *TweetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TweetResponse.Merge(m, src)
}
func (m *TweetResponse) XXX_Size() int {
	return xxx_messageInfo_TweetResponse.Size(m)
}
func (m *TweetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TweetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TweetResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TweetRequest)(nil), "TweetRequest")
	proto.RegisterType((*TweetResponse)(nil), "TweetResponse")
}

func init() {
	proto.RegisterFile("tweet-gateway.proto", fileDescriptor_0a2c6e576d127c43)
}

var fileDescriptor_0a2c6e576d127c43 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x29, 0x4f, 0x4d,
	0x2d, 0xd1, 0x4d, 0x4f, 0x2c, 0x49, 0x2d, 0x4f, 0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57,
	0x72, 0xe1, 0xe2, 0x09, 0x01, 0x09, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x71,
	0x71, 0x94, 0x16, 0xa7, 0x16, 0xe5, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0xc1, 0xf9, 0x42, 0x12, 0x5c, 0xec, 0xc9, 0xf9, 0x79, 0x25, 0xa9, 0x79, 0x25, 0x12, 0x4c, 0x60,
	0x29, 0x18, 0x57, 0x89, 0x9f, 0x8b, 0x17, 0x6a, 0x4a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x91,
	0x19, 0xd4, 0x58, 0x77, 0x88, 0x65, 0x42, 0x6a, 0x5c, 0xac, 0x60, 0xdb, 0x85, 0x78, 0xf5, 0x90,
	0xad, 0x93, 0xe2, 0xd3, 0x43, 0xd1, 0x97, 0xc4, 0x06, 0x76, 0x95, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x2d, 0x9e, 0x1c, 0xcd, 0xac, 0x00, 0x00, 0x00,
}
