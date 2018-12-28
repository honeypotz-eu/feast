// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/specs/FeatureSpec.proto

package specs // import "github.com/gojektech/feast/protos/generated/go/feast/specs"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/gojektech/feast/protos/generated/go/feast/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FeatureSpec struct {
	Id                   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner                string                 `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Description          string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Uri                  string                 `protobuf:"bytes,5,opt,name=uri,proto3" json:"uri,omitempty"`
	Granularity          types.Granularity_Enum `protobuf:"varint,6,opt,name=granularity,proto3,enum=feast.types.Granularity_Enum" json:"granularity,omitempty"`
	ValueType            types.ValueType_Enum   `protobuf:"varint,7,opt,name=valueType,proto3,enum=feast.types.ValueType_Enum" json:"valueType,omitempty"`
	Entity               string                 `protobuf:"bytes,8,opt,name=entity,proto3" json:"entity,omitempty"`
	Group                string                 `protobuf:"bytes,9,opt,name=group,proto3" json:"group,omitempty"`
	Tags                 []string               `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	Options              map[string]string      `protobuf:"bytes,11,rep,name=options,proto3" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DataStores           *DataStores            `protobuf:"bytes,12,opt,name=dataStores,proto3" json:"dataStores,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *FeatureSpec) Reset()         { *m = FeatureSpec{} }
func (m *FeatureSpec) String() string { return proto.CompactTextString(m) }
func (*FeatureSpec) ProtoMessage()    {}
func (*FeatureSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_FeatureSpec_0c79e2ccc174df52, []int{0}
}
func (m *FeatureSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureSpec.Unmarshal(m, b)
}
func (m *FeatureSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureSpec.Marshal(b, m, deterministic)
}
func (dst *FeatureSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureSpec.Merge(dst, src)
}
func (m *FeatureSpec) XXX_Size() int {
	return xxx_messageInfo_FeatureSpec.Size(m)
}
func (m *FeatureSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureSpec proto.InternalMessageInfo

func (m *FeatureSpec) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FeatureSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FeatureSpec) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *FeatureSpec) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *FeatureSpec) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *FeatureSpec) GetGranularity() types.Granularity_Enum {
	if m != nil {
		return m.Granularity
	}
	return types.Granularity_NONE
}

func (m *FeatureSpec) GetValueType() types.ValueType_Enum {
	if m != nil {
		return m.ValueType
	}
	return types.ValueType_UNKNOWN
}

func (m *FeatureSpec) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *FeatureSpec) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *FeatureSpec) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *FeatureSpec) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *FeatureSpec) GetDataStores() *DataStores {
	if m != nil {
		return m.DataStores
	}
	return nil
}

type DataStores struct {
	Serving              *DataStore `protobuf:"bytes,1,opt,name=serving,proto3" json:"serving,omitempty"`
	Warehouse            *DataStore `protobuf:"bytes,2,opt,name=warehouse,proto3" json:"warehouse,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DataStores) Reset()         { *m = DataStores{} }
func (m *DataStores) String() string { return proto.CompactTextString(m) }
func (*DataStores) ProtoMessage()    {}
func (*DataStores) Descriptor() ([]byte, []int) {
	return fileDescriptor_FeatureSpec_0c79e2ccc174df52, []int{1}
}
func (m *DataStores) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataStores.Unmarshal(m, b)
}
func (m *DataStores) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataStores.Marshal(b, m, deterministic)
}
func (dst *DataStores) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataStores.Merge(dst, src)
}
func (m *DataStores) XXX_Size() int {
	return xxx_messageInfo_DataStores.Size(m)
}
func (m *DataStores) XXX_DiscardUnknown() {
	xxx_messageInfo_DataStores.DiscardUnknown(m)
}

var xxx_messageInfo_DataStores proto.InternalMessageInfo

func (m *DataStores) GetServing() *DataStore {
	if m != nil {
		return m.Serving
	}
	return nil
}

func (m *DataStores) GetWarehouse() *DataStore {
	if m != nil {
		return m.Warehouse
	}
	return nil
}

type DataStore struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Options              map[string]string `protobuf:"bytes,2,rep,name=options,proto3" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DataStore) Reset()         { *m = DataStore{} }
func (m *DataStore) String() string { return proto.CompactTextString(m) }
func (*DataStore) ProtoMessage()    {}
func (*DataStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_FeatureSpec_0c79e2ccc174df52, []int{2}
}
func (m *DataStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataStore.Unmarshal(m, b)
}
func (m *DataStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataStore.Marshal(b, m, deterministic)
}
func (dst *DataStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataStore.Merge(dst, src)
}
func (m *DataStore) XXX_Size() int {
	return xxx_messageInfo_DataStore.Size(m)
}
func (m *DataStore) XXX_DiscardUnknown() {
	xxx_messageInfo_DataStore.DiscardUnknown(m)
}

var xxx_messageInfo_DataStore proto.InternalMessageInfo

func (m *DataStore) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataStore) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

func init() {
	proto.RegisterType((*FeatureSpec)(nil), "feast.specs.FeatureSpec")
	proto.RegisterMapType((map[string]string)(nil), "feast.specs.FeatureSpec.OptionsEntry")
	proto.RegisterType((*DataStores)(nil), "feast.specs.DataStores")
	proto.RegisterType((*DataStore)(nil), "feast.specs.DataStore")
	proto.RegisterMapType((map[string]string)(nil), "feast.specs.DataStore.OptionsEntry")
}

func init() {
	proto.RegisterFile("feast/specs/FeatureSpec.proto", fileDescriptor_FeatureSpec_0c79e2ccc174df52)
}

var fileDescriptor_FeatureSpec_0c79e2ccc174df52 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x51, 0x6b, 0xd4, 0x40,
	0x10, 0x26, 0x49, 0xdb, 0x33, 0x93, 0x52, 0xca, 0x22, 0xed, 0x72, 0x5a, 0x08, 0x27, 0xc2, 0x3d,
	0x25, 0x72, 0x0a, 0xea, 0x81, 0x14, 0x8a, 0xa7, 0x8f, 0x4a, 0x2a, 0x7d, 0xd0, 0xa7, 0x6d, 0x32,
	0xe6, 0x62, 0x7b, 0xd9, 0xb0, 0xbb, 0x69, 0xc9, 0x1f, 0xf1, 0x6f, 0xfa, 0x17, 0x64, 0x27, 0x77,
	0xcd, 0x1e, 0x3d, 0x9f, 0x7c, 0x9b, 0xd9, 0xef, 0x9b, 0xc9, 0x7c, 0xdf, 0x64, 0xe0, 0xec, 0x27,
	0x0a, 0x6d, 0x52, 0xdd, 0x60, 0xae, 0xd3, 0x4f, 0x28, 0x4c, 0xab, 0xf0, 0xb2, 0xc1, 0x3c, 0x69,
	0x94, 0x34, 0x92, 0x45, 0x04, 0x27, 0x04, 0x8f, 0x9f, 0xbb, 0xdc, 0x45, 0x6d, 0x2a, 0xd3, 0x0d,
	0xd4, 0xf1, 0x56, 0xa7, 0x4b, 0x23, 0x95, 0x28, 0xf1, 0x31, 0x6c, 0xba, 0x06, 0x75, 0xfa, 0x59,
	0x89, 0xba, 0xbd, 0x15, 0xaa, 0x32, 0xdd, 0x1a, 0x3e, 0x75, 0xe1, 0x2b, 0x71, 0xdb, 0x62, 0x0f,
	0x4c, 0xfe, 0x04, 0x10, 0x39, 0x73, 0xb1, 0x23, 0xf0, 0xab, 0x82, 0x7b, 0xb1, 0x37, 0x0d, 0x33,
	0xbf, 0x2a, 0x18, 0x83, 0xbd, 0x5a, 0xac, 0x90, 0xfb, 0xf4, 0x42, 0x31, 0x7b, 0x0a, 0xfb, 0xf2,
	0xbe, 0x46, 0xc5, 0x03, 0x7a, 0xec, 0x13, 0x16, 0x43, 0x54, 0xa0, 0xce, 0x55, 0xd5, 0x98, 0x4a,
	0xd6, 0x7c, 0x8f, 0x30, 0xf7, 0x89, 0x1d, 0x43, 0xd0, 0xaa, 0x8a, 0xef, 0x13, 0x62, 0x43, 0x76,
	0x0e, 0x51, 0x39, 0xcc, 0xca, 0x0f, 0x62, 0x6f, 0x7a, 0x34, 0x3b, 0x4b, 0x7a, 0x57, 0x68, 0xd8,
	0xc4, 0xd5, 0xb2, 0xa8, 0xdb, 0x55, 0xe6, 0x56, 0xb0, 0xf7, 0x10, 0xde, 0x59, 0x35, 0xdf, 0xba,
	0x06, 0xf9, 0x88, 0xca, 0x9f, 0x6d, 0x95, 0x5f, 0x6d, 0xd0, 0xbe, 0x78, 0x60, 0xb3, 0x13, 0x38,
	0x40, 0x32, 0x99, 0x3f, 0xa1, 0x81, 0xd6, 0x99, 0x55, 0x57, 0x2a, 0xd9, 0x36, 0x3c, 0xec, 0xd5,
	0x51, 0x62, 0x7d, 0x30, 0xa2, 0xd4, 0x1c, 0xe2, 0xc0, 0xfa, 0x60, 0x63, 0x76, 0x0e, 0x23, 0x49,
	0xca, 0x34, 0x8f, 0xe2, 0x60, 0x1a, 0xcd, 0x5e, 0x26, 0xce, 0x3e, 0x13, 0x77, 0xdd, 0x5f, 0x7a,
	0xde, 0xa2, 0x36, 0xaa, 0xcb, 0x36, 0x55, 0xec, 0x2d, 0x40, 0x21, 0x8c, 0xb0, 0xdb, 0x44, 0xcd,
	0x0f, 0x63, 0x6f, 0x1a, 0xcd, 0x4e, 0xb7, 0x7a, 0x7c, 0x7c, 0x80, 0x33, 0x87, 0x3a, 0x9e, 0xc3,
	0xa1, 0xdb, 0xd1, 0x3a, 0x7b, 0x83, 0xdd, 0x7a, 0x6d, 0x36, 0xb4, 0x2a, 0x48, 0xea, 0x7a, 0x71,
	0x7d, 0x32, 0xf7, 0xdf, 0x79, 0x13, 0x03, 0x30, 0x74, 0x65, 0xaf, 0x60, 0xa4, 0x51, 0xdd, 0x55,
	0x75, 0x49, 0xd5, 0xd1, 0xec, 0x64, 0xf7, 0xf7, 0xb3, 0x0d, 0x8d, 0xbd, 0x81, 0xf0, 0x5e, 0x28,
	0x5c, 0xca, 0x56, 0xf7, 0xdd, 0xff, 0x5d, 0x33, 0x10, 0x27, 0xbf, 0x3d, 0x08, 0x1f, 0x80, 0x47,
	0x7f, 0xd9, 0x87, 0xc1, 0x49, 0x9f, 0x9c, 0x7c, 0xb1, 0xbb, 0xe3, 0x6e, 0x1f, 0xff, 0xc7, 0x8e,
	0x8b, 0x1f, 0xe0, 0x1e, 0xe1, 0xc5, 0xb1, 0xb3, 0xb5, 0xaf, 0xf6, 0x42, 0xbe, 0xcf, 0xcb, 0xca,
	0x2c, 0xdb, 0xeb, 0x24, 0x97, 0xab, 0xb4, 0x94, 0xbf, 0xf0, 0xc6, 0x60, 0xbe, 0x4c, 0xfb, 0x7b,
	0xa2, 0x1b, 0xd2, 0x69, 0x89, 0x35, 0x2a, 0x61, 0xb0, 0x48, 0x4b, 0x99, 0x3a, 0x77, 0x7a, 0x7d,
	0x40, 0x84, 0xd7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x29, 0x17, 0xe6, 0xc4, 0x07, 0x04, 0x00,
	0x00,
}