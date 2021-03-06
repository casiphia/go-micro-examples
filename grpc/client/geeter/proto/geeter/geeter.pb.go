// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/geeter/geeter.proto

package geeter

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

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" gorm:"-"`
	XXX_unrecognized     []byte   `json:"-" gorm:"-"`
	XXX_sizecache        int32    `json:"-" gorm:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ccd1e22cb6df655, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" gorm:"-"`
	XXX_unrecognized     []byte   `json:"-" gorm:"-"`
	XXX_sizecache        int32    `json:"-" gorm:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ccd1e22cb6df655, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" gorm:"-"`
	XXX_unrecognized     []byte   `json:"-" gorm:"-"`
	XXX_sizecache        int32    `json:"-" gorm:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ccd1e22cb6df655, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" gorm:"-"`
	XXX_unrecognized     []byte   `json:"-" gorm:"-"`
	XXX_sizecache        int32    `json:"-" gorm:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ccd1e22cb6df655, []int{3}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" gorm:"-"`
	XXX_unrecognized     []byte   `json:"-" gorm:"-"`
	XXX_sizecache        int32    `json:"-" gorm:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ccd1e22cb6df655, []int{4}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" gorm:"-"`
	XXX_unrecognized     []byte   `json:"-" gorm:"-"`
	XXX_sizecache        int32    `json:"-" gorm:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ccd1e22cb6df655, []int{5}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" gorm:"-"`
	XXX_unrecognized     []byte   `json:"-" gorm:"-"`
	XXX_sizecache        int32    `json:"-" gorm:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ccd1e22cb6df655, []int{6}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "geeter.Message")
	proto.RegisterType((*Request)(nil), "geeter.Request")
	proto.RegisterType((*Response)(nil), "geeter.Response")
	proto.RegisterType((*StreamingRequest)(nil), "geeter.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "geeter.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "geeter.Ping")
	proto.RegisterType((*Pong)(nil), "geeter.Pong")
}

func init() { proto.RegisterFile("proto/geeter/geeter.proto", fileDescriptor_4ccd1e22cb6df655) }

var fileDescriptor_4ccd1e22cb6df655 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xdd, 0x4a, 0xc3, 0x40,
	0x10, 0x85, 0xbb, 0x34, 0x6e, 0xeb, 0x20, 0x18, 0x07, 0x91, 0x36, 0xfe, 0x20, 0x7b, 0x15, 0x51,
	0x6a, 0xd1, 0x27, 0x10, 0x2f, 0xbc, 0x12, 0x24, 0x3e, 0xc1, 0x2a, 0xc3, 0x52, 0x6c, 0x76, 0x6b,
	0x66, 0x7b, 0xe1, 0x0b, 0xf9, 0x9c, 0x92, 0xfd, 0xa1, 0x22, 0xf1, 0x2a, 0x73, 0xe6, 0x3b, 0x33,
	0x9c, 0xc9, 0xc2, 0x7c, 0xd3, 0x39, 0xef, 0x6e, 0x0d, 0x91, 0xa7, 0x2e, 0x7d, 0x16, 0xa1, 0x87,
	0x32, 0x2a, 0x75, 0x0a, 0x93, 0x67, 0x62, 0xd6, 0x86, 0xb0, 0x84, 0x31, 0xeb, 0xaf, 0x99, 0xb8,
	0x14, 0xf5, 0x7e, 0xd3, 0x97, 0xea, 0x1c, 0x26, 0x0d, 0x7d, 0x6e, 0x89, 0x3d, 0x22, 0x14, 0x56,
	0xb7, 0x94, 0x68, 0xa8, 0xd5, 0x19, 0x4c, 0x1b, 0xe2, 0x8d, 0xb3, 0x1c, 0x86, 0x5b, 0x36, 0x79,
	0xb8, 0x65, 0xa3, 0x6a, 0x28, 0x5f, 0x7d, 0x47, 0xba, 0x5d, 0x59, 0x93, 0xb7, 0x1c, 0xc3, 0xde,
	0xbb, 0xdb, 0x5a, 0x1f, 0x7c, 0xe3, 0x26, 0x0a, 0x75, 0x05, 0x47, 0xbf, 0x9c, 0x69, 0xe1, 0xb0,
	0xf5, 0x02, 0x8a, 0x97, 0x95, 0x35, 0x78, 0x02, 0x92, 0x7d, 0xe7, 0x3e, 0x28, 0xe1, 0xa4, 0x02,
	0x77, 0xff, 0xf3, 0xbb, 0x6f, 0x01, 0xf2, 0x29, 0x5c, 0x8e, 0xd7, 0x50, 0x3c, 0xea, 0xf5, 0x1a,
	0x0f, 0x17, 0xe9, 0xc7, 0xa4, 0x90, 0x55, 0xb9, 0x6b, 0xc4, 0x2c, 0x6a, 0x84, 0x0f, 0x20, 0x63,
	0x44, 0x9c, 0x65, 0xfa, 0xf7, 0xb8, 0x6a, 0x3e, 0x40, 0xf2, 0x82, 0xa5, 0xc0, 0x1b, 0x98, 0xf6,
	0xd1, 0x43, 0xbc, 0x83, 0x6c, 0xed, 0x3b, 0xd5, 0x4e, 0x39, 0x6b, 0xd4, 0xa8, 0x16, 0x4b, 0xf1,
	0x26, 0xc3, 0x33, 0xdd, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xee, 0x7a, 0xbf, 0xc3, 0x01,
	0x00, 0x00,
}
