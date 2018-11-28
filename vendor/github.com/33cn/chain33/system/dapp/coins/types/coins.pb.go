// Code generated by protoc-gen-go. DO NOT EDIT.
// source: coins.proto

package types

import (
	fmt "fmt"
	math "math"

	types "github.com/33cn/chain33/types"
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

// message for execs.coins
type CoinsAction struct {
	// Types that are valid to be assigned to Value:
	//	*CoinsAction_Transfer
	//	*CoinsAction_Withdraw
	//	*CoinsAction_Genesis
	//	*CoinsAction_TransferToExec
	Value                isCoinsAction_Value `protobuf_oneof:"value"`
	Ty                   int32               `protobuf:"varint,3,opt,name=ty,proto3" json:"ty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CoinsAction) Reset()         { *m = CoinsAction{} }
func (m *CoinsAction) String() string { return proto.CompactTextString(m) }
func (*CoinsAction) ProtoMessage()    {}
func (*CoinsAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_da4483c99519c66a, []int{0}
}

func (m *CoinsAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoinsAction.Unmarshal(m, b)
}
func (m *CoinsAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoinsAction.Marshal(b, m, deterministic)
}
func (m *CoinsAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoinsAction.Merge(m, src)
}
func (m *CoinsAction) XXX_Size() int {
	return xxx_messageInfo_CoinsAction.Size(m)
}
func (m *CoinsAction) XXX_DiscardUnknown() {
	xxx_messageInfo_CoinsAction.DiscardUnknown(m)
}

var xxx_messageInfo_CoinsAction proto.InternalMessageInfo

type isCoinsAction_Value interface {
	isCoinsAction_Value()
}

type CoinsAction_Transfer struct {
	Transfer *types.AssetsTransfer `protobuf:"bytes,1,opt,name=transfer,proto3,oneof"`
}

type CoinsAction_Withdraw struct {
	Withdraw *types.AssetsWithdraw `protobuf:"bytes,4,opt,name=withdraw,proto3,oneof"`
}

type CoinsAction_Genesis struct {
	Genesis *types.AssetsGenesis `protobuf:"bytes,2,opt,name=genesis,proto3,oneof"`
}

type CoinsAction_TransferToExec struct {
	TransferToExec *types.AssetsTransferToExec `protobuf:"bytes,5,opt,name=transferToExec,proto3,oneof"`
}

func (*CoinsAction_Transfer) isCoinsAction_Value() {}

func (*CoinsAction_Withdraw) isCoinsAction_Value() {}

func (*CoinsAction_Genesis) isCoinsAction_Value() {}

func (*CoinsAction_TransferToExec) isCoinsAction_Value() {}

func (m *CoinsAction) GetValue() isCoinsAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *CoinsAction) GetTransfer() *types.AssetsTransfer {
	if x, ok := m.GetValue().(*CoinsAction_Transfer); ok {
		return x.Transfer
	}
	return nil
}

func (m *CoinsAction) GetWithdraw() *types.AssetsWithdraw {
	if x, ok := m.GetValue().(*CoinsAction_Withdraw); ok {
		return x.Withdraw
	}
	return nil
}

func (m *CoinsAction) GetGenesis() *types.AssetsGenesis {
	if x, ok := m.GetValue().(*CoinsAction_Genesis); ok {
		return x.Genesis
	}
	return nil
}

func (m *CoinsAction) GetTransferToExec() *types.AssetsTransferToExec {
	if x, ok := m.GetValue().(*CoinsAction_TransferToExec); ok {
		return x.TransferToExec
	}
	return nil
}

func (m *CoinsAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CoinsAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CoinsAction_OneofMarshaler, _CoinsAction_OneofUnmarshaler, _CoinsAction_OneofSizer, []interface{}{
		(*CoinsAction_Transfer)(nil),
		(*CoinsAction_Withdraw)(nil),
		(*CoinsAction_Genesis)(nil),
		(*CoinsAction_TransferToExec)(nil),
	}
}

func _CoinsAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CoinsAction)
	// value
	switch x := m.Value.(type) {
	case *CoinsAction_Transfer:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Transfer); err != nil {
			return err
		}
	case *CoinsAction_Withdraw:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Withdraw); err != nil {
			return err
		}
	case *CoinsAction_Genesis:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Genesis); err != nil {
			return err
		}
	case *CoinsAction_TransferToExec:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TransferToExec); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CoinsAction.Value has unexpected type %T", x)
	}
	return nil
}

func _CoinsAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CoinsAction)
	switch tag {
	case 1: // value.transfer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(types.AssetsTransfer)
		err := b.DecodeMessage(msg)
		m.Value = &CoinsAction_Transfer{msg}
		return true, err
	case 4: // value.withdraw
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(types.AssetsWithdraw)
		err := b.DecodeMessage(msg)
		m.Value = &CoinsAction_Withdraw{msg}
		return true, err
	case 2: // value.genesis
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(types.AssetsGenesis)
		err := b.DecodeMessage(msg)
		m.Value = &CoinsAction_Genesis{msg}
		return true, err
	case 5: // value.transferToExec
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(types.AssetsTransferToExec)
		err := b.DecodeMessage(msg)
		m.Value = &CoinsAction_TransferToExec{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CoinsAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CoinsAction)
	// value
	switch x := m.Value.(type) {
	case *CoinsAction_Transfer:
		s := proto.Size(x.Transfer)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CoinsAction_Withdraw:
		s := proto.Size(x.Withdraw)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CoinsAction_Genesis:
		s := proto.Size(x.Genesis)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CoinsAction_TransferToExec:
		s := proto.Size(x.TransferToExec)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*CoinsAction)(nil), "types.CoinsAction")
}

func init() { proto.RegisterFile("coins.proto", fileDescriptor_da4483c99519c66a) }

var fileDescriptor_da4483c99519c66a = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xce, 0xcf, 0xcc,
	0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x96, 0x12,
	0x2c, 0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x83, 0xc8, 0x28, 0x75, 0x32,
	0x71, 0x71, 0x3b, 0x83, 0x54, 0x3a, 0x82, 0x45, 0x85, 0x8c, 0xb9, 0x38, 0xc0, 0x8a, 0xd2, 0x52,
	0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x44, 0xf5, 0xc0, 0x9a, 0xf5, 0x1c, 0x8b, 0x8b,
	0x53, 0x4b, 0x8a, 0x43, 0xa0, 0x92, 0x1e, 0x0c, 0x41, 0x70, 0x85, 0x20, 0x4d, 0xe5, 0x99, 0x25,
	0x19, 0x29, 0x45, 0x89, 0xe5, 0x12, 0x2c, 0x58, 0x34, 0x85, 0x43, 0x25, 0x41, 0x9a, 0x60, 0x0a,
	0x85, 0x0c, 0xb8, 0xd8, 0xd3, 0x53, 0xf3, 0x52, 0x8b, 0x33, 0x8b, 0x25, 0x98, 0xc0, 0x7a, 0x44,
	0x50, 0xf4, 0xb8, 0x43, 0xe4, 0x3c, 0x18, 0x82, 0x60, 0xca, 0x84, 0x5c, 0xb9, 0xf8, 0x60, 0x56,
	0x86, 0xe4, 0xbb, 0x56, 0xa4, 0x26, 0x4b, 0xb0, 0x82, 0x35, 0x4a, 0x63, 0x75, 0x21, 0x44, 0x89,
	0x07, 0x43, 0x10, 0x9a, 0x26, 0x21, 0x3e, 0x2e, 0xa6, 0x92, 0x4a, 0x09, 0x66, 0x05, 0x46, 0x0d,
	0xd6, 0x20, 0xa6, 0x92, 0x4a, 0x27, 0x76, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x27, 0xf6,
	0x28, 0x48, 0x38, 0x25, 0xb1, 0x81, 0xc3, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x27,
	0xc9, 0xb3, 0x44, 0x01, 0x00, 0x00,
}