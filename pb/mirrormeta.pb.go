// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mirrormeta.proto

package pb

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MirrorMeta struct {
	Package              []*MirrorMeta_Package `protobuf:"bytes,1,rep,name=package" json:"package,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MirrorMeta) Reset()         { *m = MirrorMeta{} }
func (m *MirrorMeta) String() string { return proto.CompactTextString(m) }
func (*MirrorMeta) ProtoMessage()    {}
func (*MirrorMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_beb1b007a14d6496, []int{0}
}

func (m *MirrorMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MirrorMeta.Unmarshal(m, b)
}
func (m *MirrorMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MirrorMeta.Marshal(b, m, deterministic)
}
func (m *MirrorMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MirrorMeta.Merge(m, src)
}
func (m *MirrorMeta) XXX_Size() int {
	return xxx_messageInfo_MirrorMeta.Size(m)
}
func (m *MirrorMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_MirrorMeta.DiscardUnknown(m)
}

var xxx_messageInfo_MirrorMeta proto.InternalMessageInfo

func (m *MirrorMeta) GetPackage() []*MirrorMeta_Package {
	if m != nil {
		return m.Package
	}
	return nil
}

type MirrorMeta_Package struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	WellKnownPath        []string `protobuf:"bytes,2,rep,name=well_known_path,json=wellKnownPath" json:"well_known_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MirrorMeta_Package) Reset()         { *m = MirrorMeta_Package{} }
func (m *MirrorMeta_Package) String() string { return proto.CompactTextString(m) }
func (*MirrorMeta_Package) ProtoMessage()    {}
func (*MirrorMeta_Package) Descriptor() ([]byte, []int) {
	return fileDescriptor_beb1b007a14d6496, []int{0, 0}
}

func (m *MirrorMeta_Package) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MirrorMeta_Package.Unmarshal(m, b)
}
func (m *MirrorMeta_Package) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MirrorMeta_Package.Marshal(b, m, deterministic)
}
func (m *MirrorMeta_Package) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MirrorMeta_Package.Merge(m, src)
}
func (m *MirrorMeta_Package) XXX_Size() int {
	return xxx_messageInfo_MirrorMeta_Package.Size(m)
}
func (m *MirrorMeta_Package) XXX_DiscardUnknown() {
	xxx_messageInfo_MirrorMeta_Package.DiscardUnknown(m)
}

var xxx_messageInfo_MirrorMeta_Package proto.InternalMessageInfo

func (m *MirrorMeta_Package) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *MirrorMeta_Package) GetWellKnownPath() []string {
	if m != nil {
		return m.WellKnownPath
	}
	return nil
}

func init() {
	proto.RegisterType((*MirrorMeta)(nil), "pb.MirrorMeta")
	proto.RegisterType((*MirrorMeta_Package)(nil), "pb.MirrorMeta.Package")
}

func init() { proto.RegisterFile("mirrormeta.proto", fileDescriptor_beb1b007a14d6496) }

var fileDescriptor_beb1b007a14d6496 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xcd, 0x2c, 0x2a,
	0xca, 0x2f, 0xca, 0x4d, 0x2d, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
	0x52, 0x6a, 0x65, 0xe4, 0xe2, 0xf2, 0x05, 0x4b, 0xf8, 0xa6, 0x96, 0x24, 0x0a, 0x19, 0x70, 0xb1,
	0x17, 0x24, 0x26, 0x67, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x89, 0xe9,
	0x15, 0x24, 0xe9, 0x21, 0x14, 0xe8, 0x05, 0x40, 0x64, 0x83, 0x60, 0xca, 0xa4, 0x5c, 0xb9, 0xd8,
	0xa1, 0x62, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0x20, 0x9d, 0x8c, 0x1a, 0x9c, 0x41, 0x60,
	0xb6, 0x90, 0x1a, 0x17, 0x7f, 0x79, 0x6a, 0x4e, 0x4e, 0x7c, 0x76, 0x5e, 0x7e, 0x79, 0x5e, 0x7c,
	0x41, 0x62, 0x49, 0x86, 0x04, 0x93, 0x02, 0xb3, 0x06, 0x67, 0x10, 0x2f, 0x48, 0xd8, 0x1b, 0x24,
	0x1a, 0x90, 0x58, 0x92, 0x01, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x3c, 0xf5, 0x87, 0x9e, 0x00,
	0x00, 0x00,
}
