// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/compute/v1/disk_type.proto

package compute

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

type DiskType struct {
	// ID of the disk type.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Description of the disk type. 0-256 characters long.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Array of availability zones where the disk type is available.
	ZoneIds              []string `protobuf:"bytes,3,rep,name=zone_ids,json=zoneIds,proto3" json:"zone_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiskType) Reset()         { *m = DiskType{} }
func (m *DiskType) String() string { return proto.CompactTextString(m) }
func (*DiskType) ProtoMessage()    {}
func (*DiskType) Descriptor() ([]byte, []int) {
	return fileDescriptor_1766b4de885b3ba4, []int{0}
}

func (m *DiskType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiskType.Unmarshal(m, b)
}
func (m *DiskType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiskType.Marshal(b, m, deterministic)
}
func (m *DiskType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiskType.Merge(m, src)
}
func (m *DiskType) XXX_Size() int {
	return xxx_messageInfo_DiskType.Size(m)
}
func (m *DiskType) XXX_DiscardUnknown() {
	xxx_messageInfo_DiskType.DiscardUnknown(m)
}

var xxx_messageInfo_DiskType proto.InternalMessageInfo

func (m *DiskType) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DiskType) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DiskType) GetZoneIds() []string {
	if m != nil {
		return m.ZoneIds
	}
	return nil
}

func init() {
	proto.RegisterType((*DiskType)(nil), "yandex.cloud.compute.v1.DiskType")
}

func init() {
	proto.RegisterFile("yandex/cloud/compute/v1/disk_type.proto", fileDescriptor_1766b4de885b3ba4)
}

var fileDescriptor_1766b4de885b3ba4 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xaf, 0x4c, 0xcc, 0x4b,
	0x49, 0xad, 0xd0, 0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0x4f, 0xce, 0xcf, 0x2d, 0x28, 0x2d, 0x49,
	0xd5, 0x2f, 0x33, 0xd4, 0x4f, 0xc9, 0x2c, 0xce, 0x8e, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0x87, 0x28, 0xd4, 0x03, 0x2b, 0xd4, 0x83, 0x2a, 0xd4, 0x2b, 0x33,
	0x54, 0x0a, 0xe7, 0xe2, 0x70, 0xc9, 0x2c, 0xce, 0x0e, 0xa9, 0x2c, 0x48, 0x15, 0xe2, 0xe3, 0x62,
	0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x52, 0xe0, 0xe2,
	0x4e, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0xcc, 0xcf, 0x93, 0x60, 0x02, 0x4b, 0x20,
	0x0b, 0x09, 0x49, 0x72, 0x71, 0x54, 0xe5, 0xe7, 0xa5, 0xc6, 0x67, 0xa6, 0x14, 0x4b, 0x30, 0x2b,
	0x30, 0x6b, 0x70, 0x06, 0xb1, 0x83, 0xf8, 0x9e, 0x29, 0xc5, 0x4e, 0xae, 0x51, 0xce, 0xe9, 0x99,
	0x25, 0x19, 0xa5, 0x49, 0x20, 0xdb, 0xf4, 0x21, 0xd6, 0xeb, 0x42, 0xdc, 0x99, 0x9e, 0xaf, 0x9b,
	0x9e, 0x9a, 0x07, 0x76, 0x98, 0x3e, 0x0e, 0x0f, 0x58, 0x43, 0x99, 0x49, 0x6c, 0x60, 0x65, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0x85, 0x4f, 0x91, 0xea, 0x00, 0x00, 0x00,
}
