// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/networking/v1alpha1/virtual_mesh.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/types"
	_ "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for VirtualMeshSpec
func (this *VirtualMeshSpec) MarshalJSON() ([]byte, error) {
	str, err := VirtualMeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshSpec
func (this *VirtualMeshSpec) UnmarshalJSON(b []byte) error {
	return VirtualMeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualMeshSpec_CertificateAuthority
func (this *VirtualMeshSpec_CertificateAuthority) MarshalJSON() ([]byte, error) {
	str, err := VirtualMeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshSpec_CertificateAuthority
func (this *VirtualMeshSpec_CertificateAuthority) UnmarshalJSON(b []byte) error {
	return VirtualMeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualMeshSpec_CertificateAuthority_Builtin
func (this *VirtualMeshSpec_CertificateAuthority_Builtin) MarshalJSON() ([]byte, error) {
	str, err := VirtualMeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshSpec_CertificateAuthority_Builtin
func (this *VirtualMeshSpec_CertificateAuthority_Builtin) UnmarshalJSON(b []byte) error {
	return VirtualMeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualMeshSpec_CertificateAuthority_Provided
func (this *VirtualMeshSpec_CertificateAuthority_Provided) MarshalJSON() ([]byte, error) {
	str, err := VirtualMeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshSpec_CertificateAuthority_Provided
func (this *VirtualMeshSpec_CertificateAuthority_Provided) UnmarshalJSON(b []byte) error {
	return VirtualMeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualMeshSpec_Federation
func (this *VirtualMeshSpec_Federation) MarshalJSON() ([]byte, error) {
	str, err := VirtualMeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshSpec_Federation
func (this *VirtualMeshSpec_Federation) UnmarshalJSON(b []byte) error {
	return VirtualMeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualMeshSpec_SharedTrust
func (this *VirtualMeshSpec_SharedTrust) MarshalJSON() ([]byte, error) {
	str, err := VirtualMeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshSpec_SharedTrust
func (this *VirtualMeshSpec_SharedTrust) UnmarshalJSON(b []byte) error {
	return VirtualMeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualMeshSpec_LimitedTrust
func (this *VirtualMeshSpec_LimitedTrust) MarshalJSON() ([]byte, error) {
	str, err := VirtualMeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshSpec_LimitedTrust
func (this *VirtualMeshSpec_LimitedTrust) UnmarshalJSON(b []byte) error {
	return VirtualMeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualMeshStatus
func (this *VirtualMeshStatus) MarshalJSON() ([]byte, error) {
	str, err := VirtualMeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshStatus
func (this *VirtualMeshStatus) UnmarshalJSON(b []byte) error {
	return VirtualMeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	VirtualMeshMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	VirtualMeshUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
