// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/test66/test66.proto

package go_micro_srv_test66

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
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_145664bf23f309d1, []int{0}
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
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_145664bf23f309d1, []int{1}
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
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_145664bf23f309d1, []int{2}
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
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_145664bf23f309d1, []int{3}
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
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_145664bf23f309d1, []int{4}
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
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_145664bf23f309d1, []int{5}
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
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_145664bf23f309d1, []int{6}
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
	proto.RegisterType((*Message)(nil), "go.micro.srv.test66.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.test66.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.test66.Response")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.test66.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.test66.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "go.micro.srv.test66.Ping")
	proto.RegisterType((*Pong)(nil), "go.micro.srv.test66.Pong")
}

func init() { proto.RegisterFile("proto/test66/test66.proto", fileDescriptor_145664bf23f309d1) }

var fileDescriptor_145664bf23f309d1 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x1b, 0x5a, 0xb7, 0x75, 0x4e, 0x75, 0x14, 0xb1, 0x6b, 0x2b, 0x12, 0x50, 0xd6, 0x4b,
	0x2c, 0x0a, 0x7d, 0x01, 0xd1, 0x9b, 0x20, 0xab, 0x37, 0x4f, 0xb1, 0x84, 0xb0, 0xd8, 0x24, 0x35,
	0x93, 0x0a, 0xbe, 0xb7, 0x0f, 0x20, 0x9b, 0xcd, 0x82, 0xc8, 0x2e, 0x9e, 0x76, 0x66, 0xbf, 0x7f,
	0x86, 0xff, 0x1f, 0x02, 0xb3, 0xad, 0x77, 0xc1, 0x5d, 0x07, 0x45, 0x61, 0xb5, 0x4a, 0x1f, 0x11,
	0xff, 0xe1, 0xa1, 0x76, 0xc2, 0x54, 0x6b, 0xef, 0x04, 0xf9, 0x4f, 0xd1, 0x20, 0x7e, 0x0a, 0xe3,
	0x47, 0x45, 0x24, 0xb5, 0xc2, 0x29, 0x0c, 0x49, 0x7e, 0x9d, 0xb0, 0x73, 0x56, 0xec, 0x97, 0x75,
	0xc9, 0x17, 0x30, 0x2e, 0xd5, 0xc7, 0x4e, 0x51, 0x40, 0x84, 0x91, 0x95, 0x46, 0x25, 0x1a, 0x6b,
	0x3e, 0x87, 0x49, 0xa9, 0x68, 0xeb, 0x2c, 0xc5, 0x61, 0x43, 0xba, 0x1d, 0x36, 0xa4, 0x79, 0x01,
	0xd3, 0xe7, 0xe0, 0x95, 0x34, 0x95, 0xd5, 0xed, 0x96, 0x23, 0xd8, 0x5b, 0xbb, 0x9d, 0x0d, 0x51,
	0x37, 0x2c, 0x9b, 0x86, 0x5f, 0xc1, 0xc1, 0x2f, 0x65, 0x5a, 0xd8, 0x2d, 0x3d, 0x83, 0xd1, 0x53,
	0x65, 0x35, 0x1e, 0x43, 0x46, 0xc1, 0xbb, 0x77, 0x95, 0x70, 0xea, 0x22, 0x77, 0xfd, 0xfc, 0xe6,
	0x9b, 0x41, 0xf6, 0x12, 0x93, 0xe3, 0x3d, 0x8c, 0xee, 0xe4, 0x66, 0x83, 0x73, 0xd1, 0x71, 0x17,
	0x91, 0x1c, 0xe7, 0x8b, 0x1e, 0xda, 0xb8, 0xe4, 0x03, 0x7c, 0x85, 0xac, 0x31, 0x8f, 0x17, 0x9d,
	0xd2, 0xbf, 0x37, 0xc8, 0x2f, 0xff, 0x93, 0xb5, 0xab, 0x97, 0x0c, 0x1f, 0x60, 0x52, 0xc7, 0x8d,
	0x91, 0x66, 0x9d, 0x73, 0x35, 0xce, 0x7b, 0x90, 0xb3, 0x9a, 0x0f, 0x0a, 0xb6, 0x64, 0x6f, 0x59,
	0x7c, 0x01, 0xb7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x05, 0x78, 0x70, 0x69, 0x1e, 0x02, 0x00,
	0x00,
}
