// Code generated by protoc-gen-go. DO NOT EDIT.
// source: template.proto

/*
Package model is a generated protocol buffer package.

It is generated from these files:
	template.proto

It has these top-level messages:
	ResourceTemplate
	NoneTemplate
	InstanceTemplate
	LxcTemplate
*/
package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ResourceTemplate struct {
	// Types that are valid to be assigned to Item:
	//	*ResourceTemplate_Instance
	//	*ResourceTemplate_None
	Item isResourceTemplate_Item `protobuf_oneof:"Item"`
}

func (m *ResourceTemplate) Reset()                    { *m = ResourceTemplate{} }
func (m *ResourceTemplate) String() string            { return proto.CompactTextString(m) }
func (*ResourceTemplate) ProtoMessage()               {}
func (*ResourceTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isResourceTemplate_Item interface {
	isResourceTemplate_Item()
}

type ResourceTemplate_Instance struct {
	Instance *InstanceTemplate `protobuf:"bytes,100,opt,name=instance,oneof"`
}
type ResourceTemplate_None struct {
	None *NoneTemplate `protobuf:"bytes,101,opt,name=none,oneof"`
}

func (*ResourceTemplate_Instance) isResourceTemplate_Item() {}
func (*ResourceTemplate_None) isResourceTemplate_Item()     {}

func (m *ResourceTemplate) GetItem() isResourceTemplate_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *ResourceTemplate) GetInstance() *InstanceTemplate {
	if x, ok := m.GetItem().(*ResourceTemplate_Instance); ok {
		return x.Instance
	}
	return nil
}

func (m *ResourceTemplate) GetNone() *NoneTemplate {
	if x, ok := m.GetItem().(*ResourceTemplate_None); ok {
		return x.None
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ResourceTemplate) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ResourceTemplate_OneofMarshaler, _ResourceTemplate_OneofUnmarshaler, _ResourceTemplate_OneofSizer, []interface{}{
		(*ResourceTemplate_Instance)(nil),
		(*ResourceTemplate_None)(nil),
	}
}

func _ResourceTemplate_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ResourceTemplate)
	// Item
	switch x := m.Item.(type) {
	case *ResourceTemplate_Instance:
		b.EncodeVarint(100<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Instance); err != nil {
			return err
		}
	case *ResourceTemplate_None:
		b.EncodeVarint(101<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.None); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ResourceTemplate.Item has unexpected type %T", x)
	}
	return nil
}

func _ResourceTemplate_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ResourceTemplate)
	switch tag {
	case 100: // Item.instance
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(InstanceTemplate)
		err := b.DecodeMessage(msg)
		m.Item = &ResourceTemplate_Instance{msg}
		return true, err
	case 101: // Item.none
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NoneTemplate)
		err := b.DecodeMessage(msg)
		m.Item = &ResourceTemplate_None{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ResourceTemplate_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ResourceTemplate)
	// Item
	switch x := m.Item.(type) {
	case *ResourceTemplate_Instance:
		s := proto.Size(x.Instance)
		n += proto.SizeVarint(100<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ResourceTemplate_None:
		s := proto.Size(x.None)
		n += proto.SizeVarint(101<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NoneTemplate struct {
}

func (m *NoneTemplate) Reset()                    { *m = NoneTemplate{} }
func (m *NoneTemplate) String() string            { return proto.CompactTextString(m) }
func (*NoneTemplate) ProtoMessage()               {}
func (*NoneTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type InstanceTemplate struct {
	// Types that are valid to be assigned to Item:
	//	*InstanceTemplate_Lxc
	Item isInstanceTemplate_Item `protobuf_oneof:"Item"`
}

func (m *InstanceTemplate) Reset()                    { *m = InstanceTemplate{} }
func (m *InstanceTemplate) String() string            { return proto.CompactTextString(m) }
func (*InstanceTemplate) ProtoMessage()               {}
func (*InstanceTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isInstanceTemplate_Item interface {
	isInstanceTemplate_Item()
}

type InstanceTemplate_Lxc struct {
	Lxc *LxcTemplate `protobuf:"bytes,201,opt,name=lxc,oneof"`
}

func (*InstanceTemplate_Lxc) isInstanceTemplate_Item() {}

func (m *InstanceTemplate) GetItem() isInstanceTemplate_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *InstanceTemplate) GetLxc() *LxcTemplate {
	if x, ok := m.GetItem().(*InstanceTemplate_Lxc); ok {
		return x.Lxc
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*InstanceTemplate) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _InstanceTemplate_OneofMarshaler, _InstanceTemplate_OneofUnmarshaler, _InstanceTemplate_OneofSizer, []interface{}{
		(*InstanceTemplate_Lxc)(nil),
	}
}

func _InstanceTemplate_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*InstanceTemplate)
	// Item
	switch x := m.Item.(type) {
	case *InstanceTemplate_Lxc:
		b.EncodeVarint(201<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Lxc); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("InstanceTemplate.Item has unexpected type %T", x)
	}
	return nil
}

func _InstanceTemplate_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*InstanceTemplate)
	switch tag {
	case 201: // Item.lxc
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LxcTemplate)
		err := b.DecodeMessage(msg)
		m.Item = &InstanceTemplate_Lxc{msg}
		return true, err
	default:
		return false, nil
	}
}

func _InstanceTemplate_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*InstanceTemplate)
	// Item
	switch x := m.Item.(type) {
	case *InstanceTemplate_Lxc:
		s := proto.Size(x.Lxc)
		n += proto.SizeVarint(201<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type LxcTemplate struct {
	Vcpu     int32 `protobuf:"varint,1,opt,name=vcpu" json:"vcpu,omitempty"`
	MemoryGb int32 `protobuf:"varint,2,opt,name=memory_gb" json:"memory_gb,omitempty"`
}

func (m *LxcTemplate) Reset()                    { *m = LxcTemplate{} }
func (m *LxcTemplate) String() string            { return proto.CompactTextString(m) }
func (*LxcTemplate) ProtoMessage()               {}
func (*LxcTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LxcTemplate) GetVcpu() int32 {
	if m != nil {
		return m.Vcpu
	}
	return 0
}

func (m *LxcTemplate) GetMemoryGb() int32 {
	if m != nil {
		return m.MemoryGb
	}
	return 0
}

func init() {
	proto.RegisterType((*ResourceTemplate)(nil), "model.ResourceTemplate")
	proto.RegisterType((*NoneTemplate)(nil), "model.NoneTemplate")
	proto.RegisterType((*InstanceTemplate)(nil), "model.InstanceTemplate")
	proto.RegisterType((*LxcTemplate)(nil), "model.LxcTemplate")
}

func init() { proto.RegisterFile("template.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x49, 0xcd, 0x2d,
	0xc8, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcd, 0x4f, 0x49,
	0xcd, 0x51, 0xaa, 0xe1, 0x12, 0x08, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0x0d, 0x81, 0x2a,
	0x10, 0x32, 0xe5, 0xe2, 0xc8, 0xcc, 0x2b, 0x2e, 0x49, 0xcc, 0x4b, 0x4e, 0x95, 0x48, 0x51, 0x60,
	0xd4, 0xe0, 0x36, 0x12, 0xd7, 0x03, 0xab, 0xd6, 0xf3, 0x84, 0x0a, 0xc3, 0x94, 0x7a, 0x30, 0x04,
	0xc1, 0x95, 0x0a, 0x69, 0x72, 0xb1, 0xe4, 0xe5, 0xe7, 0xa5, 0x4a, 0xa4, 0x82, 0xb5, 0x08, 0x43,
	0xb5, 0xf8, 0xe5, 0xe7, 0x21, 0x2b, 0x07, 0x2b, 0x71, 0x62, 0xe3, 0x62, 0xf1, 0x2c, 0x49, 0xcd,
	0x55, 0xe2, 0xe3, 0xe2, 0x41, 0x96, 0x57, 0x72, 0xe6, 0x12, 0x40, 0xb7, 0x42, 0x48, 0x9d, 0x8b,
	0x39, 0xa7, 0x22, 0x59, 0xe2, 0x24, 0x23, 0xd8, 0x58, 0x21, 0xa8, 0xb1, 0x3e, 0x15, 0xc9, 0x48,
	0xa6, 0x82, 0x54, 0xc0, 0x0d, 0xb5, 0xe7, 0xe2, 0x46, 0x92, 0x15, 0x12, 0xe2, 0x62, 0x29, 0x4b,
	0x2e, 0x28, 0x95, 0x00, 0xe9, 0x67, 0x0d, 0x02, 0xb3, 0x85, 0x64, 0xb8, 0x38, 0x73, 0x53, 0x73,
	0xf3, 0x8b, 0x2a, 0xe3, 0xd3, 0x93, 0x24, 0x98, 0xc0, 0x12, 0x08, 0x01, 0x27, 0xa5, 0x28, 0x85,
	0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xcc, 0x92, 0xfc, 0xd2, 0xa2,
	0x4c, 0xfd, 0x82, 0xd4, 0x92, 0xcc, 0x92, 0x30, 0x17, 0x67, 0x7d, 0xb0, 0xfd, 0x49, 0x6c, 0xe0,
	0x50, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x02, 0xc0, 0x4e, 0xd3, 0x57, 0x01, 0x00, 0x00,
}
