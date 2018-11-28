// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pbft.proto

package types

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Operation struct {
	Value                *Block   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc19f28ccff0670, []int{0}
}

func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (m *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(m, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

func (m *Operation) GetValue() *Block {
	if m != nil {
		return m.Value
	}
	return nil
}

type Checkpoint struct {
	Sequence             uint32   `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Digest               []byte   `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Checkpoint) Reset()         { *m = Checkpoint{} }
func (m *Checkpoint) String() string { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()    {}
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc19f28ccff0670, []int{1}
}

func (m *Checkpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Checkpoint.Unmarshal(m, b)
}
func (m *Checkpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Checkpoint.Marshal(b, m, deterministic)
}
func (m *Checkpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checkpoint.Merge(m, src)
}
func (m *Checkpoint) XXX_Size() int {
	return xxx_messageInfo_Checkpoint.Size(m)
}
func (m *Checkpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Checkpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Checkpoint proto.InternalMessageInfo

func (m *Checkpoint) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Checkpoint) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

type Entry struct {
	Sequence             uint32   `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Digest               []byte   `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	View                 uint32   `protobuf:"varint,3,opt,name=view,proto3" json:"view,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc19f28ccff0670, []int{2}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Entry) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *Entry) GetView() uint32 {
	if m != nil {
		return m.View
	}
	return 0
}

type ViewChange struct {
	Viewchanger          uint32   `protobuf:"varint,1,opt,name=viewchanger,proto3" json:"viewchanger,omitempty"`
	Digest               []byte   `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewChange) Reset()         { *m = ViewChange{} }
func (m *ViewChange) String() string { return proto.CompactTextString(m) }
func (*ViewChange) ProtoMessage()    {}
func (*ViewChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc19f28ccff0670, []int{3}
}

func (m *ViewChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewChange.Unmarshal(m, b)
}
func (m *ViewChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewChange.Marshal(b, m, deterministic)
}
func (m *ViewChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewChange.Merge(m, src)
}
func (m *ViewChange) XXX_Size() int {
	return xxx_messageInfo_ViewChange.Size(m)
}
func (m *ViewChange) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewChange.DiscardUnknown(m)
}

var xxx_messageInfo_ViewChange proto.InternalMessageInfo

func (m *ViewChange) GetViewchanger() uint32 {
	if m != nil {
		return m.Viewchanger
	}
	return 0
}

func (m *ViewChange) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

type Summary struct {
	Sequence             uint32   `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Digest               []byte   `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Summary) Reset()         { *m = Summary{} }
func (m *Summary) String() string { return proto.CompactTextString(m) }
func (*Summary) ProtoMessage()    {}
func (*Summary) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc19f28ccff0670, []int{4}
}

func (m *Summary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Summary.Unmarshal(m, b)
}
func (m *Summary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Summary.Marshal(b, m, deterministic)
}
func (m *Summary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Summary.Merge(m, src)
}
func (m *Summary) XXX_Size() int {
	return xxx_messageInfo_Summary.Size(m)
}
func (m *Summary) XXX_DiscardUnknown() {
	xxx_messageInfo_Summary.DiscardUnknown(m)
}

var xxx_messageInfo_Summary proto.InternalMessageInfo

func (m *Summary) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Summary) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

type Result struct {
	Value                *Block   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc19f28ccff0670, []int{5}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetValue() *Block {
	if m != nil {
		return m.Value
	}
	return nil
}

type Request struct {
	// Types that are valid to be assigned to Value:
	//	*Request_Client
	//	*Request_Preprepare
	//	*Request_Prepare
	//	*Request_Commit
	//	*Request_Checkpoint
	//	*Request_Viewchange
	//	*Request_Ack
	//	*Request_Newview
	Value                isRequest_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc19f28ccff0670, []int{6}
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

type isRequest_Value interface {
	isRequest_Value()
}

type Request_Client struct {
	Client *RequestClient `protobuf:"bytes,1,opt,name=client,proto3,oneof"`
}

type Request_Preprepare struct {
	Preprepare *RequestPrePrepare `protobuf:"bytes,2,opt,name=preprepare,proto3,oneof"`
}

type Request_Prepare struct {
	Prepare *RequestPrepare `protobuf:"bytes,3,opt,name=prepare,proto3,oneof"`
}

type Request_Commit struct {
	Commit *RequestCommit `protobuf:"bytes,4,opt,name=commit,proto3,oneof"`
}

type Request_Checkpoint struct {
	Checkpoint *RequestCheckpoint `protobuf:"bytes,5,opt,name=checkpoint,proto3,oneof"`
}

type Request_Viewchange struct {
	Viewchange *RequestViewChange `protobuf:"bytes,6,opt,name=viewchange,proto3,oneof"`
}

type Request_Ack struct {
	Ack *RequestAck `protobuf:"bytes,7,opt,name=ack,proto3,oneof"`
}

type Request_Newview struct {
	Newview *RequestNewView `protobuf:"bytes,8,opt,name=newview,proto3,oneof"`
}

func (*Request_Client) isRequest_Value() {}

func (*Request_Preprepare) isRequest_Value() {}

func (*Request_Prepare) isRequest_Value() {}

func (*Request_Commit) isRequest_Value() {}

func (*Request_Checkpoint) isRequest_Value() {}

func (*Request_Viewchange) isRequest_Value() {}

func (*Request_Ack) isRequest_Value() {}

func (*Request_Newview) isRequest_Value() {}

func (m *Request) GetValue() isRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Request) GetClient() *RequestClient {
	if x, ok := m.GetValue().(*Request_Client); ok {
		return x.Client
	}
	return nil
}

func (m *Request) GetPreprepare() *RequestPrePrepare {
	if x, ok := m.GetValue().(*Request_Preprepare); ok {
		return x.Preprepare
	}
	return nil
}

func (m *Request) GetPrepare() *RequestPrepare {
	if x, ok := m.GetValue().(*Request_Prepare); ok {
		return x.Prepare
	}
	return nil
}

func (m *Request) GetCommit() *RequestCommit {
	if x, ok := m.GetValue().(*Request_Commit); ok {
		return x.Commit
	}
	return nil
}

func (m *Request) GetCheckpoint() *RequestCheckpoint {
	if x, ok := m.GetValue().(*Request_Checkpoint); ok {
		return x.Checkpoint
	}
	return nil
}

func (m *Request) GetViewchange() *RequestViewChange {
	if x, ok := m.GetValue().(*Request_Viewchange); ok {
		return x.Viewchange
	}
	return nil
}

func (m *Request) GetAck() *RequestAck {
	if x, ok := m.GetValue().(*Request_Ack); ok {
		return x.Ack
	}
	return nil
}

func (m *Request) GetNewview() *RequestNewView {
	if x, ok := m.GetValue().(*Request_Newview); ok {
		return x.Newview
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Request) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Request_OneofMarshaler, _Request_OneofUnmarshaler, _Request_OneofSizer, []interface{}{
		(*Request_Client)(nil),
		(*Request_Preprepare)(nil),
		(*Request_Prepare)(nil),
		(*Request_Commit)(nil),
		(*Request_Checkpoint)(nil),
		(*Request_Viewchange)(nil),
		(*Request_Ack)(nil),
		(*Request_Newview)(nil),
	}
}

func _Request_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Request)
	// value
	switch x := m.Value.(type) {
	case *Request_Client:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Client); err != nil {
			return err
		}
	case *Request_Preprepare:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Preprepare); err != nil {
			return err
		}
	case *Request_Prepare:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Prepare); err != nil {
			return err
		}
	case *Request_Commit:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Commit); err != nil {
			return err
		}
	case *Request_Checkpoint:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Checkpoint); err != nil {
			return err
		}
	case *Request_Viewchange:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Viewchange); err != nil {
			return err
		}
	case *Request_Ack:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ack); err != nil {
			return err
		}
	case *Request_Newview:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Newview); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Request.Value has unexpected type %T", x)
	}
	return nil
}

func _Request_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Request)
	switch tag {
	case 1: // value.client
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestClient)
		err := b.DecodeMessage(msg)
		m.Value = &Request_Client{msg}
		return true, err
	case 2: // value.preprepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestPrePrepare)
		err := b.DecodeMessage(msg)
		m.Value = &Request_Preprepare{msg}
		return true, err
	case 3: // value.prepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestPrepare)
		err := b.DecodeMessage(msg)
		m.Value = &Request_Prepare{msg}
		return true, err
	case 4: // value.commit
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestCommit)
		err := b.DecodeMessage(msg)
		m.Value = &Request_Commit{msg}
		return true, err
	case 5: // value.checkpoint
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestCheckpoint)
		err := b.DecodeMessage(msg)
		m.Value = &Request_Checkpoint{msg}
		return true, err
	case 6: // value.viewchange
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestViewChange)
		err := b.DecodeMessage(msg)
		m.Value = &Request_Viewchange{msg}
		return true, err
	case 7: // value.ack
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestAck)
		err := b.DecodeMessage(msg)
		m.Value = &Request_Ack{msg}
		return true, err
	case 8: // value.newview
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestNewView)
		err := b.DecodeMessage(msg)
		m.Value = &Request_Newview{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Request_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Request)
	// value
	switch x := m.Value.(type) {
	case *Request_Client:
		s := proto.Size(x.Client)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Request_Preprepare:
		s := proto.Size(x.Preprepare)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Request_Prepare:
		s := proto.Size(x.Prepare)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Request_Commit:
		s := proto.Size(x.Commit)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Request_Checkpoint:
		s := proto.Size(x.Checkpoint)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Request_Viewchange:
		s := proto.Size(x.Viewchange)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Request_Ack:
		s := proto.Size(x.Ack)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Request_Newview:
		s := proto.Size(x.Newview)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type RequestClient struct {
	Op                   *Operation `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	Timestamp            string     `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Client               string     `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RequestClient) Reset()         { *m = RequestClient{} }
func (m *RequestClient) String() string { return proto.CompactTextString(m) }
func (*RequestClient) ProtoMessage()    {}
func (*RequestClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc19f28ccff0670, []int{7}
}

func (m *RequestClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestClient.Unmarshal(m, b)
}
func (m *RequestClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestClient.Marshal(b, m, deterministic)
}
func (m *RequestClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestClient.Merge(m, src)
}
func (m *RequestClient) XXX_Size() int {
	return xxx_messageInfo_RequestClient.Size(m)
}
func (m *RequestClient) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestClient.DiscardUnknown(m)
}

var xxx_messageInfo_RequestClient proto.InternalMessageInfo

func (m *RequestClient) GetOp() *Operation {
	if m != nil {
		return m.Op
	}
	return nil
}

func (m *RequestClient) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *RequestClient) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

type RequestPrePrepare struct {
	View                 uint32   `protobuf:"varint,1,opt,name=view,proto3" json:"view,omitempty"`
	Sequence             uint32   `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Digest               []byte   `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	Replica              uint32   `protobuf:"varint,4,opt,name=replica,proto3" json:"replica,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestPrePrepare) Reset()         { *m = RequestPrePrepare{} }
func (m *RequestPrePrepare) String() string { return proto.CompactTextString(m) }
func (*RequestPrePrepare) ProtoMessage()    {}
func (*RequestPrePrepare) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc19f28ccff0670, []int{8}
}

func (m *RequestPrePrepare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestPrePrepare.Unmarshal(m, b)
}
func (m *RequestPrePrepare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestPrePrepare.Marshal(b, m, deterministic)
}
func (m *RequestPrePrepare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestPrePrepare.Merge(m, src)
}
func (m *RequestPrePrepare) XXX_Size() int {
	return xxx_messageInfo_RequestPrePrepare.Size(m)
}
func (m *RequestPrePrepare) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestPrePrepare.DiscardUnknown(m)
}

var xxx_messageInfo_RequestPrePrepare proto.InternalMessageInfo

func (m *RequestPrePrepare) GetView() uint32 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *RequestPrePrepare) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *RequestPrePrepare) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *RequestPrePrepare) GetReplica() uint32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

type RequestPrepare struct {
	View                 uint32   `protobuf:"varint,1,opt,name=view,proto3" json:"view,omitempty"`
	Sequence             uint32   `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Digest               []byte   `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	Replica              uint32   `protobuf:"varint,4,opt,name=replica,proto3" json:"replica,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestPrepare) Reset()         { *m = RequestPrepare{} }
func (m *RequestPrepare) String() string { return proto.CompactTextString(m) }
func (*RequestPrepare) ProtoMessage()    {}
func (*RequestPrepare) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc19f28ccff0670, []int{9}
}

func (m *RequestPrepare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestPrepare.Unmarshal(m, b)
}
func (m *RequestPrepare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestPrepare.Marshal(b, m, deterministic)
}
func (m *RequestPrepare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestPrepare.Merge(m, src)
}
func (m *RequestPrepare) XXX_Size() int {
	return xxx_messageInfo_RequestPrepare.Size(m)
}
func (m *RequestPrepare) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestPrepare.DiscardUnknown(m)
}

var xxx_messageInfo_RequestPrepare proto.InternalMessageInfo

func (m *RequestPrepare) GetView() uint32 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *RequestPrepare) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *RequestPrepare) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *RequestPrepare) GetReplica() uint32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

type RequestCommit struct {
	View                 uint32   `protobuf:"varint,1,opt,name=view,proto3" json:"view,omitempty"`
	Sequence             uint32   `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Replica              uint32   `protobuf:"varint,3,opt,name=replica,proto3" json:"replica,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestCommit) Reset()         { *m = RequestCommit{} }
func (m *RequestCommit) String() string { return proto.CompactTextString(m) }
func (*RequestCommit) ProtoMessage()    {}
func (*RequestCommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc19f28ccff0670, []int{10}
}

func (m *RequestCommit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestCommit.Unmarshal(m, b)
}
func (m *RequestCommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestCommit.Marshal(b, m, deterministic)
}
func (m *RequestCommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestCommit.Merge(m, src)
}
func (m *RequestCommit) XXX_Size() int {
	return xxx_messageInfo_RequestCommit.Size(m)
}
func (m *RequestCommit) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestCommit.DiscardUnknown(m)
}

var xxx_messageInfo_RequestCommit proto.InternalMessageInfo

func (m *RequestCommit) GetView() uint32 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *RequestCommit) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *RequestCommit) GetReplica() uint32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

type RequestCheckpoint struct {
	Sequence             uint32   `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Digest               []byte   `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	Replica              uint32   `protobuf:"varint,3,opt,name=replica,proto3" json:"replica,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestCheckpoint) Reset()         { *m = RequestCheckpoint{} }
func (m *RequestCheckpoint) String() string { return proto.CompactTextString(m) }
func (*RequestCheckpoint) ProtoMessage()    {}
func (*RequestCheckpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc19f28ccff0670, []int{11}
}

func (m *RequestCheckpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestCheckpoint.Unmarshal(m, b)
}
func (m *RequestCheckpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestCheckpoint.Marshal(b, m, deterministic)
}
func (m *RequestCheckpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestCheckpoint.Merge(m, src)
}
func (m *RequestCheckpoint) XXX_Size() int {
	return xxx_messageInfo_RequestCheckpoint.Size(m)
}
func (m *RequestCheckpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestCheckpoint.DiscardUnknown(m)
}

var xxx_messageInfo_RequestCheckpoint proto.InternalMessageInfo

func (m *RequestCheckpoint) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *RequestCheckpoint) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *RequestCheckpoint) GetReplica() uint32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

type RequestViewChange struct {
	View                 uint32        `protobuf:"varint,1,opt,name=view,proto3" json:"view,omitempty"`
	Sequence             uint32        `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Checkpoints          []*Checkpoint `protobuf:"bytes,3,rep,name=checkpoints,proto3" json:"checkpoints,omitempty"`
	Preps                []*Entry      `protobuf:"bytes,4,rep,name=preps,proto3" json:"preps,omitempty"`
	Prepreps             []*Entry      `protobuf:"bytes,5,rep,name=prepreps,proto3" json:"prepreps,omitempty"`
	Replica              uint32        `protobuf:"varint,6,opt,name=replica,proto3" json:"replica,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RequestViewChange) Reset()         { *m = RequestViewChange{} }
func (m *RequestViewChange) String() string { return proto.CompactTextString(m) }
func (*RequestViewChange) ProtoMessage()    {}
func (*RequestViewChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc19f28ccff0670, []int{12}
}

func (m *RequestViewChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestViewChange.Unmarshal(m, b)
}
func (m *RequestViewChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestViewChange.Marshal(b, m, deterministic)
}
func (m *RequestViewChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestViewChange.Merge(m, src)
}
func (m *RequestViewChange) XXX_Size() int {
	return xxx_messageInfo_RequestViewChange.Size(m)
}
func (m *RequestViewChange) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestViewChange.DiscardUnknown(m)
}

var xxx_messageInfo_RequestViewChange proto.InternalMessageInfo

func (m *RequestViewChange) GetView() uint32 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *RequestViewChange) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *RequestViewChange) GetCheckpoints() []*Checkpoint {
	if m != nil {
		return m.Checkpoints
	}
	return nil
}

func (m *RequestViewChange) GetPreps() []*Entry {
	if m != nil {
		return m.Preps
	}
	return nil
}

func (m *RequestViewChange) GetPrepreps() []*Entry {
	if m != nil {
		return m.Prepreps
	}
	return nil
}

func (m *RequestViewChange) GetReplica() uint32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

type RequestAck struct {
	View                 uint32   `protobuf:"varint,1,opt,name=view,proto3" json:"view,omitempty"`
	Replica              uint32   `protobuf:"varint,2,opt,name=replica,proto3" json:"replica,omitempty"`
	Viewchanger          uint32   `protobuf:"varint,3,opt,name=viewchanger,proto3" json:"viewchanger,omitempty"`
	Digest               []byte   `protobuf:"bytes,4,opt,name=digest,proto3" json:"digest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestAck) Reset()         { *m = RequestAck{} }
func (m *RequestAck) String() string { return proto.CompactTextString(m) }
func (*RequestAck) ProtoMessage()    {}
func (*RequestAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc19f28ccff0670, []int{13}
}

func (m *RequestAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestAck.Unmarshal(m, b)
}
func (m *RequestAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestAck.Marshal(b, m, deterministic)
}
func (m *RequestAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestAck.Merge(m, src)
}
func (m *RequestAck) XXX_Size() int {
	return xxx_messageInfo_RequestAck.Size(m)
}
func (m *RequestAck) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestAck.DiscardUnknown(m)
}

var xxx_messageInfo_RequestAck proto.InternalMessageInfo

func (m *RequestAck) GetView() uint32 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *RequestAck) GetReplica() uint32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

func (m *RequestAck) GetViewchanger() uint32 {
	if m != nil {
		return m.Viewchanger
	}
	return 0
}

func (m *RequestAck) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

type RequestNewView struct {
	View                 uint32        `protobuf:"varint,1,opt,name=view,proto3" json:"view,omitempty"`
	Viewchanges          []*ViewChange `protobuf:"bytes,2,rep,name=viewchanges,proto3" json:"viewchanges,omitempty"`
	Summaries            []*Summary    `protobuf:"bytes,4,rep,name=summaries,proto3" json:"summaries,omitempty"`
	Replica              uint32        `protobuf:"varint,5,opt,name=replica,proto3" json:"replica,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RequestNewView) Reset()         { *m = RequestNewView{} }
func (m *RequestNewView) String() string { return proto.CompactTextString(m) }
func (*RequestNewView) ProtoMessage()    {}
func (*RequestNewView) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc19f28ccff0670, []int{14}
}

func (m *RequestNewView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestNewView.Unmarshal(m, b)
}
func (m *RequestNewView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestNewView.Marshal(b, m, deterministic)
}
func (m *RequestNewView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestNewView.Merge(m, src)
}
func (m *RequestNewView) XXX_Size() int {
	return xxx_messageInfo_RequestNewView.Size(m)
}
func (m *RequestNewView) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestNewView.DiscardUnknown(m)
}

var xxx_messageInfo_RequestNewView proto.InternalMessageInfo

func (m *RequestNewView) GetView() uint32 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *RequestNewView) GetViewchanges() []*ViewChange {
	if m != nil {
		return m.Viewchanges
	}
	return nil
}

func (m *RequestNewView) GetSummaries() []*Summary {
	if m != nil {
		return m.Summaries
	}
	return nil
}

func (m *RequestNewView) GetReplica() uint32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

type ClientReply struct {
	View                 uint32   `protobuf:"varint,1,opt,name=view,proto3" json:"view,omitempty"`
	Timestamp            string   `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Client               string   `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
	Replica              uint32   `protobuf:"varint,4,opt,name=replica,proto3" json:"replica,omitempty"`
	Result               *Result  `protobuf:"bytes,5,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientReply) Reset()         { *m = ClientReply{} }
func (m *ClientReply) String() string { return proto.CompactTextString(m) }
func (*ClientReply) ProtoMessage()    {}
func (*ClientReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc19f28ccff0670, []int{15}
}

func (m *ClientReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientReply.Unmarshal(m, b)
}
func (m *ClientReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientReply.Marshal(b, m, deterministic)
}
func (m *ClientReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientReply.Merge(m, src)
}
func (m *ClientReply) XXX_Size() int {
	return xxx_messageInfo_ClientReply.Size(m)
}
func (m *ClientReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientReply.DiscardUnknown(m)
}

var xxx_messageInfo_ClientReply proto.InternalMessageInfo

func (m *ClientReply) GetView() uint32 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *ClientReply) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *ClientReply) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

func (m *ClientReply) GetReplica() uint32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

func (m *ClientReply) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Operation)(nil), "types.Operation")
	proto.RegisterType((*Checkpoint)(nil), "types.Checkpoint")
	proto.RegisterType((*Entry)(nil), "types.Entry")
	proto.RegisterType((*ViewChange)(nil), "types.ViewChange")
	proto.RegisterType((*Summary)(nil), "types.Summary")
	proto.RegisterType((*Result)(nil), "types.Result")
	proto.RegisterType((*Request)(nil), "types.Request")
	proto.RegisterType((*RequestClient)(nil), "types.RequestClient")
	proto.RegisterType((*RequestPrePrepare)(nil), "types.RequestPrePrepare")
	proto.RegisterType((*RequestPrepare)(nil), "types.RequestPrepare")
	proto.RegisterType((*RequestCommit)(nil), "types.RequestCommit")
	proto.RegisterType((*RequestCheckpoint)(nil), "types.RequestCheckpoint")
	proto.RegisterType((*RequestViewChange)(nil), "types.RequestViewChange")
	proto.RegisterType((*RequestAck)(nil), "types.RequestAck")
	proto.RegisterType((*RequestNewView)(nil), "types.RequestNewView")
	proto.RegisterType((*ClientReply)(nil), "types.ClientReply")
}

func init() { proto.RegisterFile("pbft.proto", fileDescriptor_6cc19f28ccff0670) }

var fileDescriptor_6cc19f28ccff0670 = []byte{
	// 677 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x5f, 0x6b, 0xdb, 0x30,
	0x10, 0xaf, 0xe3, 0xc4, 0x69, 0x2e, 0x4d, 0x69, 0xc5, 0x36, 0x44, 0xd9, 0x58, 0x30, 0x14, 0xfa,
	0x50, 0x12, 0x56, 0xbf, 0x0d, 0x06, 0x5b, 0xcb, 0x46, 0x9f, 0xd6, 0xa2, 0xc1, 0x60, 0x7b, 0x73,
	0x34, 0x2d, 0x11, 0x89, 0xff, 0xcc, 0x92, 0x1b, 0xfa, 0x4d, 0xf6, 0xba, 0xcf, 0xb0, 0x2f, 0xb4,
	0x8f, 0x32, 0x24, 0x2b, 0x96, 0xb5, 0xba, 0x65, 0xed, 0x60, 0x90, 0x07, 0x4b, 0x77, 0x3f, 0xdd,
	0xe9, 0xee, 0x77, 0x3f, 0x05, 0x20, 0x9f, 0x7d, 0x95, 0x93, 0xbc, 0xc8, 0x64, 0x86, 0x7a, 0xf2,
	0x3a, 0x67, 0xe2, 0x60, 0x6f, 0xb6, 0xca, 0xe8, 0x92, 0x2e, 0x62, 0x9e, 0x56, 0x86, 0x70, 0x0a,
	0x83, 0x8b, 0x9c, 0x15, 0xb1, 0xe4, 0x59, 0x8a, 0x42, 0xe8, 0x5d, 0xc5, 0xab, 0x92, 0x61, 0x6f,
	0xec, 0x1d, 0x0d, 0x4f, 0x76, 0x26, 0x1a, 0x35, 0x39, 0x55, 0x20, 0x52, 0x99, 0xc2, 0xd7, 0x00,
	0x67, 0x0b, 0x46, 0x97, 0x79, 0xc6, 0x53, 0x89, 0x0e, 0x60, 0x5b, 0xb0, 0x6f, 0x25, 0x4b, 0x69,
	0x05, 0x1a, 0x91, 0x7a, 0x8d, 0x9e, 0x40, 0xf0, 0x85, 0xcf, 0x99, 0x90, 0xb8, 0x33, 0xf6, 0x8e,
	0x76, 0x88, 0x59, 0x85, 0x17, 0xd0, 0x7b, 0x9b, 0xca, 0xe2, 0xfa, 0x21, 0x60, 0x84, 0xa0, 0x7b,
	0xc5, 0xd9, 0x1a, 0xfb, 0xda, 0x5f, 0x7f, 0x87, 0xef, 0x00, 0x3e, 0x72, 0xb6, 0x3e, 0x5b, 0xc4,
	0xe9, 0x9c, 0xa1, 0x31, 0x0c, 0xd5, 0x2e, 0xd5, 0xab, 0xc2, 0x1c, 0xdc, 0xdc, 0xba, 0x35, 0xb1,
	0x57, 0xd0, 0xff, 0x50, 0x26, 0x49, 0xfc, 0xb0, 0xd4, 0xc2, 0x63, 0x08, 0x08, 0x13, 0xe5, 0x4a,
	0xfe, 0x55, 0x1d, 0x7f, 0xfa, 0xd0, 0x27, 0xea, 0x48, 0x21, 0xd1, 0x04, 0x02, 0xba, 0xe2, 0x2c,
	0x95, 0x06, 0xf0, 0xc8, 0x00, 0x8c, 0xfd, 0x4c, 0xdb, 0xce, 0xb7, 0x88, 0xf1, 0x42, 0x2f, 0x01,
	0xf2, 0x82, 0xa9, 0x5f, 0x5c, 0x30, 0x9d, 0xc5, 0xf0, 0x04, 0xbb, 0x98, 0xcb, 0x82, 0x5d, 0x56,
	0xf6, 0xf3, 0x2d, 0xd2, 0xf0, 0x46, 0x2f, 0xa0, 0xbf, 0x01, 0xfa, 0x1a, 0xf8, 0xf8, 0x06, 0xd0,
	0xa0, 0x36, 0x7e, 0x3a, 0xbd, 0x2c, 0x49, 0xb8, 0xc4, 0xdd, 0xd6, 0xf4, 0xb4, 0x4d, 0xa7, 0xa7,
	0xbf, 0x54, 0x7a, 0xb4, 0xa6, 0x08, 0xee, 0xb5, 0xa5, 0x67, 0x29, 0xa4, 0xd2, 0xb3, 0xde, 0x0a,
	0x6b, 0x5b, 0x85, 0x83, 0x36, 0xac, 0xed, 0xb5, 0xc2, 0x5a, 0x6f, 0x74, 0x08, 0x7e, 0x4c, 0x97,
	0xb8, 0xaf, 0x41, 0xfb, 0x2e, 0xe8, 0x0d, 0x5d, 0x9e, 0x6f, 0x11, 0x65, 0x57, 0x15, 0x48, 0xd9,
	0x5a, 0xb3, 0x68, 0xbb, 0xad, 0x02, 0xef, 0xd9, 0x5a, 0x85, 0x50, 0x15, 0x30, 0x7e, 0xa7, 0x7d,
	0xd3, 0xd0, 0x70, 0x0e, 0x23, 0xa7, 0x29, 0x68, 0x0c, 0x9d, 0x2c, 0x37, 0x6d, 0xdb, 0x33, 0xe7,
	0xd4, 0x03, 0x45, 0x3a, 0x59, 0x8e, 0x9e, 0xc2, 0x40, 0xf2, 0x84, 0x09, 0x19, 0x27, 0xb9, 0xee,
	0xd5, 0x80, 0xd8, 0x0d, 0x45, 0x26, 0xd3, 0x7a, 0x5f, 0x9b, 0xcc, 0x2a, 0x2c, 0x61, 0xff, 0x46,
	0x27, 0x6b, 0xf2, 0x7b, 0x96, 0xfc, 0x0e, 0x53, 0x3b, 0xb7, 0x32, 0xd5, 0x77, 0x86, 0x08, 0x43,
	0xbf, 0x60, 0xf9, 0x8a, 0xd3, 0x58, 0x77, 0x74, 0x44, 0x36, 0xcb, 0xb0, 0x80, 0x5d, 0x97, 0x07,
	0xff, 0x21, 0xe6, 0x27, 0x5b, 0xd3, 0x8a, 0x3f, 0xf7, 0x0d, 0xd9, 0x38, 0xda, 0x77, 0x8f, 0x8e,
	0xeb, 0x2a, 0xfe, 0x9b, 0x66, 0xdd, 0x11, 0xe2, 0x97, 0x57, 0xc7, 0x68, 0x88, 0xd0, 0x7d, 0xaf,
	0x10, 0xc1, 0xd0, 0x0e, 0x81, 0xc0, 0xfe, 0xd8, 0x6f, 0x50, 0xd8, 0xe6, 0x4e, 0x9a, 0x5e, 0x4a,
	0x66, 0xd4, 0x88, 0x0a, 0xdc, 0xd5, 0xee, 0x1b, 0x99, 0xd1, 0xe2, 0x4a, 0x2a, 0x13, 0x3a, 0x82,
	0x6d, 0x33, 0xfc, 0x02, 0xf7, 0x5a, 0xdc, 0x6a, 0x6b, 0xf3, 0x8a, 0x81, 0x7b, 0x45, 0x09, 0x60,
	0xa7, 0xa8, 0xf5, 0x6a, 0x0d, 0x6c, 0xc7, 0xc1, 0xfe, 0xa9, 0xc6, 0xfe, 0x5d, 0x6a, 0xdc, 0x75,
	0xe4, 0xf4, 0x87, 0x57, 0x73, 0xd1, 0x4c, 0x64, 0x6b, 0xe8, 0xa8, 0x19, 0x40, 0xe0, 0x8e, 0x53,
	0x39, 0xdb, 0x91, 0x66, 0x4c, 0x81, 0x8e, 0x61, 0x20, 0xb4, 0xd2, 0x73, 0xb6, 0xa9, 0xde, 0xae,
	0x81, 0x98, 0x17, 0x80, 0x58, 0x87, 0xe6, 0xed, 0x7a, 0x6e, 0x65, 0xbe, 0x7b, 0x30, 0xac, 0x84,
	0x80, 0xb0, 0x7c, 0x75, 0xdd, 0x9a, 0xe0, 0x83, 0xe6, 0xff, 0xf6, 0x71, 0x41, 0x87, 0x10, 0x14,
	0xfa, 0x99, 0x31, 0xca, 0x3a, 0xaa, 0xd5, 0x4b, 0x6d, 0x12, 0x63, 0x3c, 0x7d, 0xfe, 0xf9, 0xd9,
	0x9c, 0xcb, 0x45, 0x39, 0x9b, 0xd0, 0x2c, 0x99, 0x46, 0x11, 0x4d, 0xa7, 0xfa, 0xdd, 0x8f, 0xa2,
	0xa9, 0xf6, 0x9f, 0x05, 0xfa, 0x0f, 0x40, 0xf4, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x03, 0x07, 0x21,
	0x7a, 0x27, 0x08, 0x00, 0x00,
}