// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/core/v1alpha1/registry.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//The MeshService is an abstraction for a service which we have discovered to be, or are told, is part of a
//given mesh. The Mesh object has references to the MeshServices which belong to it.
type MeshServiceSpec struct {
	Service *ResourceRef `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// The mesh with which this service is associated
	Mesh                 *ResourceRef `protobuf:"bytes,2,opt,name=mesh,proto3" json:"mesh,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MeshServiceSpec) Reset()         { *m = MeshServiceSpec{} }
func (m *MeshServiceSpec) String() string { return proto.CompactTextString(m) }
func (*MeshServiceSpec) ProtoMessage()    {}
func (*MeshServiceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_083743b1ed1d6012, []int{0}
}
func (m *MeshServiceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshServiceSpec.Unmarshal(m, b)
}
func (m *MeshServiceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshServiceSpec.Marshal(b, m, deterministic)
}
func (m *MeshServiceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshServiceSpec.Merge(m, src)
}
func (m *MeshServiceSpec) XXX_Size() int {
	return xxx_messageInfo_MeshServiceSpec.Size(m)
}
func (m *MeshServiceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshServiceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_MeshServiceSpec proto.InternalMessageInfo

func (m *MeshServiceSpec) GetService() *ResourceRef {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *MeshServiceSpec) GetMesh() *ResourceRef {
	if m != nil {
		return m.Mesh
	}
	return nil
}

type MeshServiceStatus struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshServiceStatus) Reset()         { *m = MeshServiceStatus{} }
func (m *MeshServiceStatus) String() string { return proto.CompactTextString(m) }
func (*MeshServiceStatus) ProtoMessage()    {}
func (*MeshServiceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_083743b1ed1d6012, []int{1}
}
func (m *MeshServiceStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshServiceStatus.Unmarshal(m, b)
}
func (m *MeshServiceStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshServiceStatus.Marshal(b, m, deterministic)
}
func (m *MeshServiceStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshServiceStatus.Merge(m, src)
}
func (m *MeshServiceStatus) XXX_Size() int {
	return xxx_messageInfo_MeshServiceStatus.Size(m)
}
func (m *MeshServiceStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshServiceStatus.DiscardUnknown(m)
}

var xxx_messageInfo_MeshServiceStatus proto.InternalMessageInfo

//
//The MeshWorkload is an abstraction for a workload/client which we have discovered to be, or are told, is part of a
//given mesh. The Mesh object has references to the MeshWorkloads which belong to it.
type MeshWorkloadSpec struct {
	//
	//Resource ref to the underlying kubernetes controller which is managing the pods associated with the workloads.
	//It has the generic name kube_controller as it can represent either a deployment or a daemonset. Or potentially
	//any other kubernetes object which creates injected pods.
	//
	//The type is specified on the ResourceRef.APIGroup and ResourceRef.Kind fields
	KubeController *ResourceRef `protobuf:"bytes,1,opt,name=kube_controller,json=kubeController,proto3" json:"kube_controller,omitempty"`
	// The mesh with which this workload is associated
	Mesh                 *ResourceRef `protobuf:"bytes,2,opt,name=mesh,proto3" json:"mesh,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MeshWorkloadSpec) Reset()         { *m = MeshWorkloadSpec{} }
func (m *MeshWorkloadSpec) String() string { return proto.CompactTextString(m) }
func (*MeshWorkloadSpec) ProtoMessage()    {}
func (*MeshWorkloadSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_083743b1ed1d6012, []int{2}
}
func (m *MeshWorkloadSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshWorkloadSpec.Unmarshal(m, b)
}
func (m *MeshWorkloadSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshWorkloadSpec.Marshal(b, m, deterministic)
}
func (m *MeshWorkloadSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshWorkloadSpec.Merge(m, src)
}
func (m *MeshWorkloadSpec) XXX_Size() int {
	return xxx_messageInfo_MeshWorkloadSpec.Size(m)
}
func (m *MeshWorkloadSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshWorkloadSpec.DiscardUnknown(m)
}

var xxx_messageInfo_MeshWorkloadSpec proto.InternalMessageInfo

func (m *MeshWorkloadSpec) GetKubeController() *ResourceRef {
	if m != nil {
		return m.KubeController
	}
	return nil
}

func (m *MeshWorkloadSpec) GetMesh() *ResourceRef {
	if m != nil {
		return m.Mesh
	}
	return nil
}

type MeshWorkloadStatus struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshWorkloadStatus) Reset()         { *m = MeshWorkloadStatus{} }
func (m *MeshWorkloadStatus) String() string { return proto.CompactTextString(m) }
func (*MeshWorkloadStatus) ProtoMessage()    {}
func (*MeshWorkloadStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_083743b1ed1d6012, []int{3}
}
func (m *MeshWorkloadStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshWorkloadStatus.Unmarshal(m, b)
}
func (m *MeshWorkloadStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshWorkloadStatus.Marshal(b, m, deterministic)
}
func (m *MeshWorkloadStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshWorkloadStatus.Merge(m, src)
}
func (m *MeshWorkloadStatus) XXX_Size() int {
	return xxx_messageInfo_MeshWorkloadStatus.Size(m)
}
func (m *MeshWorkloadStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshWorkloadStatus.DiscardUnknown(m)
}

var xxx_messageInfo_MeshWorkloadStatus proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MeshServiceSpec)(nil), "core.zephyr.solo.io.MeshServiceSpec")
	proto.RegisterType((*MeshServiceStatus)(nil), "core.zephyr.solo.io.MeshServiceStatus")
	proto.RegisterType((*MeshWorkloadSpec)(nil), "core.zephyr.solo.io.MeshWorkloadSpec")
	proto.RegisterType((*MeshWorkloadStatus)(nil), "core.zephyr.solo.io.MeshWorkloadStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/mesh-projects/api/core/v1alpha1/registry.proto", fileDescriptor_083743b1ed1d6012)
}

var fileDescriptor_083743b1ed1d6012 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x87, 0x59, 0x11, 0x85, 0x08, 0x56, 0xb7, 0x3d, 0x94, 0x1e, 0xa4, 0xec, 0x49, 0x0f, 0x4d,
	0xa8, 0x7a, 0xf2, 0xa6, 0x9e, 0x44, 0x3c, 0xb8, 0x3d, 0x08, 0x5e, 0x64, 0x37, 0x4e, 0x77, 0xe3,
	0xa6, 0x4e, 0x98, 0x64, 0x0b, 0xeb, 0xd5, 0x47, 0xf0, 0x25, 0x7c, 0x2e, 0x9f, 0x44, 0x92, 0xc5,
	0xfa, 0x07, 0xa1, 0xa5, 0xb7, 0xc9, 0x64, 0xbe, 0x99, 0x0f, 0x7e, 0xec, 0xbc, 0x50, 0xae, 0xac,
	0x73, 0x2e, 0x71, 0x26, 0x2c, 0x6a, 0x1c, 0x29, 0x14, 0x33, 0xb0, 0xe5, 0xc8, 0x10, 0x3e, 0x81,
	0x74, 0x56, 0x64, 0x46, 0x09, 0x89, 0x04, 0x62, 0x3e, 0xce, 0xb4, 0x29, 0xb3, 0xb1, 0x20, 0x28,
	0x94, 0x75, 0xd4, 0x70, 0x43, 0xe8, 0x30, 0xee, 0xfa, 0x5f, 0xfe, 0x02, 0xa6, 0x6c, 0x88, 0xfb,
	0x1d, 0x5c, 0xe1, 0xe0, 0x68, 0xf9, 0x92, 0x69, 0xcb, 0x0f, 0x7a, 0x05, 0x16, 0x18, 0x4a, 0xe1,
	0xab, 0xb6, 0x9b, 0xbc, 0x46, 0xac, 0x73, 0x03, 0xb6, 0x9c, 0x00, 0xcd, 0x95, 0x84, 0x89, 0x01,
	0x19, 0x9f, 0xb1, 0x6d, 0xdb, 0x3e, 0xfb, 0xd1, 0x30, 0x3a, 0xdc, 0x39, 0x1e, 0xf2, 0x7f, 0x6e,
	0xf3, 0x14, 0x2c, 0xd6, 0x24, 0x21, 0x85, 0x69, 0xfa, 0x05, 0xc4, 0xa7, 0x6c, 0xd3, 0x2b, 0xf5,
	0x37, 0x56, 0x04, 0xc3, 0x74, 0xd2, 0x65, 0xfb, 0x3f, 0x25, 0x5c, 0xe6, 0x6a, 0x9b, 0xbc, 0x45,
	0x6c, 0xcf, 0x77, 0xef, 0x90, 0x2a, 0x8d, 0xd9, 0x63, 0x70, 0xbb, 0x62, 0x9d, 0xaa, 0xce, 0xe1,
	0x41, 0xe2, 0xb3, 0x23, 0xd4, 0x1a, 0x68, 0x65, 0xc7, 0x5d, 0x0f, 0x5e, 0x2e, 0xb8, 0x35, 0x55,
	0x7b, 0x2c, 0xfe, 0x25, 0x15, 0x5c, 0x2f, 0x6e, 0xdf, 0x3f, 0x0e, 0xa2, 0xfb, 0xeb, 0xa5, 0x29,
	0x9b, 0xaa, 0x58, 0x84, 0xf4, 0xe7, 0xd6, 0x77, 0x66, 0xae, 0x31, 0x60, 0xf3, 0xad, 0x10, 0xd0,
	0xc9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xa7, 0xf9, 0xc5, 0x3b, 0x02, 0x00, 0x00,
}

func (this *MeshServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshServiceSpec)
	if !ok {
		that2, ok := that.(MeshServiceSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Service.Equal(that1.Service) {
		return false
	}
	if !this.Mesh.Equal(that1.Mesh) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshServiceStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshServiceStatus)
	if !ok {
		that2, ok := that.(MeshServiceStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshWorkloadSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshWorkloadSpec)
	if !ok {
		that2, ok := that.(MeshWorkloadSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.KubeController.Equal(that1.KubeController) {
		return false
	}
	if !this.Mesh.Equal(that1.Mesh) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshWorkloadStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshWorkloadStatus)
	if !ok {
		that2, ok := that.(MeshWorkloadStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
