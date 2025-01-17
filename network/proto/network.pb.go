// Code generated by protoc-gen-go. DO NOT EDIT.
// source: network.proto

package go_micro_network

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/micro/go-micro/router/proto"
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

// Empty request
type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{0}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

// ListResponse is returned by ListNodes and ListNeighbours
type ListResponse struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{1}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

// Node is network node
type Node struct {
	// node ide
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// node address
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{2}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// Connect is sent when the node connects to the network
type Connect struct {
	// network mode
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connect) Reset()         { *m = Connect{} }
func (m *Connect) String() string { return proto.CompactTextString(m) }
func (*Connect) ProtoMessage()    {}
func (*Connect) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{3}
}

func (m *Connect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connect.Unmarshal(m, b)
}
func (m *Connect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connect.Marshal(b, m, deterministic)
}
func (m *Connect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connect.Merge(m, src)
}
func (m *Connect) XXX_Size() int {
	return xxx_messageInfo_Connect.Size(m)
}
func (m *Connect) XXX_DiscardUnknown() {
	xxx_messageInfo_Connect.DiscardUnknown(m)
}

var xxx_messageInfo_Connect proto.InternalMessageInfo

func (m *Connect) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Close is sent when the node disconnects from the network
type Close struct {
	// network mode
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Close) Reset()         { *m = Close{} }
func (m *Close) String() string { return proto.CompactTextString(m) }
func (*Close) ProtoMessage()    {}
func (*Close) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{4}
}

func (m *Close) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Close.Unmarshal(m, b)
}
func (m *Close) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Close.Marshal(b, m, deterministic)
}
func (m *Close) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Close.Merge(m, src)
}
func (m *Close) XXX_Size() int {
	return xxx_messageInfo_Close.Size(m)
}
func (m *Close) XXX_DiscardUnknown() {
	xxx_messageInfo_Close.DiscardUnknown(m)
}

var xxx_messageInfo_Close proto.InternalMessageInfo

func (m *Close) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Neighbour is used to nnounce node neighbourhood
type Neighbour struct {
	// network mode
	Node *Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// neighbours
	Neighbours           []*Node  `protobuf:"bytes,3,rep,name=neighbours,proto3" json:"neighbours,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Neighbour) Reset()         { *m = Neighbour{} }
func (m *Neighbour) String() string { return proto.CompactTextString(m) }
func (*Neighbour) ProtoMessage()    {}
func (*Neighbour) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{5}
}

func (m *Neighbour) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Neighbour.Unmarshal(m, b)
}
func (m *Neighbour) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Neighbour.Marshal(b, m, deterministic)
}
func (m *Neighbour) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Neighbour.Merge(m, src)
}
func (m *Neighbour) XXX_Size() int {
	return xxx_messageInfo_Neighbour.Size(m)
}
func (m *Neighbour) XXX_DiscardUnknown() {
	xxx_messageInfo_Neighbour.DiscardUnknown(m)
}

var xxx_messageInfo_Neighbour proto.InternalMessageInfo

func (m *Neighbour) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *Neighbour) GetNeighbours() []*Node {
	if m != nil {
		return m.Neighbours
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "go.micro.network.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.network.ListResponse")
	proto.RegisterType((*Node)(nil), "go.micro.network.Node")
	proto.RegisterType((*Connect)(nil), "go.micro.network.Connect")
	proto.RegisterType((*Close)(nil), "go.micro.network.Close")
	proto.RegisterType((*Neighbour)(nil), "go.micro.network.Neighbour")
}

func init() { proto.RegisterFile("network.proto", fileDescriptor_8571034d60397816) }

var fileDescriptor_8571034d60397816 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x5d, 0x4b, 0xc3, 0x30,
	0x14, 0xb5, 0xfb, 0x70, 0xec, 0xce, 0x0d, 0xc9, 0x83, 0x84, 0xc2, 0x64, 0xe4, 0x69, 0x88, 0xa6,
	0xb2, 0xa2, 0x4f, 0xbe, 0xed, 0xc1, 0x97, 0x51, 0xb0, 0xff, 0xc0, 0x36, 0xa1, 0x0b, 0xba, 0xdc,
	0x99, 0xa4, 0xf8, 0xc3, 0xfd, 0x03, 0xd2, 0xb4, 0x9b, 0xc5, 0x31, 0x61, 0xf8, 0x96, 0x73, 0xcf,
	0x3d, 0xe7, 0x5e, 0xee, 0x09, 0x8c, 0xb5, 0x74, 0x9f, 0x68, 0xde, 0xf8, 0xd6, 0xa0, 0x43, 0x72,
	0x59, 0x20, 0xdf, 0xa8, 0xdc, 0x20, 0x6f, 0xea, 0x61, 0x5c, 0x28, 0xb7, 0x2e, 0x33, 0x9e, 0xe3,
	0x26, 0xf2, 0x4c, 0x54, 0xe0, 0x5d, 0xfd, 0x30, 0x58, 0x3a, 0x69, 0x22, 0xaf, 0x6c, 0x40, 0x6d,
	0xc3, 0xc6, 0x30, 0x5a, 0x29, 0xeb, 0x52, 0xf9, 0x51, 0x4a, 0xeb, 0xd8, 0x13, 0x5c, 0xd4, 0xd0,
	0x6e, 0x51, 0x5b, 0x49, 0x6e, 0xa1, 0xaf, 0x51, 0x48, 0x4b, 0x83, 0x59, 0x77, 0x3e, 0x5a, 0x5c,
	0xf1, 0xdf, 0x53, 0x79, 0x82, 0x42, 0xa6, 0x75, 0x13, 0xbb, 0x87, 0x5e, 0x05, 0xc9, 0x04, 0x3a,
	0x4a, 0xd0, 0x60, 0x16, 0xcc, 0x87, 0x69, 0x47, 0x09, 0x42, 0x61, 0xf0, 0x2a, 0x84, 0x91, 0xd6,
	0xd2, 0x8e, 0x2f, 0xee, 0x20, 0x7b, 0x80, 0xc1, 0x12, 0xb5, 0x96, 0xb9, 0x23, 0x37, 0xd0, 0xab,
	0x5c, 0xbc, 0xec, 0xf8, 0x24, 0xdf, 0xc3, 0x62, 0xe8, 0x2f, 0xdf, 0xd1, 0xca, 0x93, 0x44, 0x08,
	0xc3, 0x44, 0xaa, 0x62, 0x9d, 0x61, 0x69, 0x4e, 0x11, 0x92, 0x47, 0x00, 0xbd, 0x13, 0x5a, 0xda,
	0xfd, 0xf3, 0x12, 0xad, 0xce, 0xc5, 0x57, 0x00, 0x83, 0xa4, 0x26, 0xc9, 0x33, 0x80, 0x3f, 0x6c,
	0x75, 0x7b, 0x4b, 0xe8, 0x8f, 0xba, 0x49, 0xa3, 0x09, 0x20, 0x9c, 0x1e, 0x30, 0xed, 0x3c, 0xd8,
	0x19, 0x59, 0xc1, 0xb0, 0xaa, 0x54, 0xc3, 0x2c, 0x99, 0x1e, 0x6e, 0xd1, 0x4a, 0x33, 0xbc, 0x3e,
	0x46, 0xef, 0xdd, 0x5e, 0x60, 0xe2, 0xdd, 0xf6, 0x4b, 0xff, 0xdb, 0x32, 0x3b, 0xf7, 0x1f, 0x2b,
	0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x19, 0x6e, 0xfe, 0x91, 0xb0, 0x02, 0x00, 0x00,
}
