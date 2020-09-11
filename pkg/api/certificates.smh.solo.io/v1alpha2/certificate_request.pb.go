// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/certificates/certificate_request.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
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

// Possible states in which a CertificateRequest can exist.
type CertificateRequestStatus_State int32

const (
	// The CertificateRequest has yet to be picked up by the issuer.
	CertificateRequestStatus_PENDING CertificateRequestStatus_State = 0
	// The Issuer has replied to the request and the signedCertificate and SigningRootCa
	// status fields will be populated.
	CertificateRequestStatus_FINISHED CertificateRequestStatus_State = 1
	// Processing the certificate workflow failed.
	CertificateRequestStatus_FAILED CertificateRequestStatus_State = 2
)

var CertificateRequestStatus_State_name = map[int32]string{
	0: "PENDING",
	1: "FINISHED",
	2: "FAILED",
}

var CertificateRequestStatus_State_value = map[string]int32{
	"PENDING":  0,
	"FINISHED": 1,
	"FAILED":   2,
}

func (x CertificateRequestStatus_State) String() string {
	return proto.EnumName(CertificateRequestStatus_State_name, int32(x))
}

func (CertificateRequestStatus_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bbb195e6f26dce40, []int{1, 0}
}

//
//CertificateRequests are generated by the CertificateRequesting Agent installed on managed clusters.
//They are used to request a signed certificate from Service Mesh Hub based on a private key
//generated by the Agent (which never leaves the managed cluster).
//
//When Service Mesh Hub creates an IssuedCertificate on a managed cluster, the local CertificateRequesting Agent
//will generate a CertificateRequest corresponding to it.
//
//Service Mesh Hub will then process the Certificate Signing Request contained in the
//CertificateRequestSpec and write the signed SSL certificate back as a secret in the managed cluster,
//and update the CertificateRequest Status to point to that secret.
type CertificateRequestSpec struct {
	// Base64-encoded data for the PKCS#10 Certificate Signing Request issued
	// by the CertificateRequesting Agent deployed in the managed cluster, corresponding
	// to the IssuedRequest received by the CertificateRequesting Agent.
	CertificateSigningRequest []byte   `protobuf:"bytes,1,opt,name=certificate_signing_request,json=certificateSigningRequest,proto3" json:"certificate_signing_request,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *CertificateRequestSpec) Reset()         { *m = CertificateRequestSpec{} }
func (m *CertificateRequestSpec) String() string { return proto.CompactTextString(m) }
func (*CertificateRequestSpec) ProtoMessage()    {}
func (*CertificateRequestSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbb195e6f26dce40, []int{0}
}
func (m *CertificateRequestSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateRequestSpec.Unmarshal(m, b)
}
func (m *CertificateRequestSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateRequestSpec.Marshal(b, m, deterministic)
}
func (m *CertificateRequestSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateRequestSpec.Merge(m, src)
}
func (m *CertificateRequestSpec) XXX_Size() int {
	return xxx_messageInfo_CertificateRequestSpec.Size(m)
}
func (m *CertificateRequestSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateRequestSpec.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateRequestSpec proto.InternalMessageInfo

func (m *CertificateRequestSpec) GetCertificateSigningRequest() []byte {
	if m != nil {
		return m.CertificateSigningRequest
	}
	return nil
}

type CertificateRequestStatus struct {
	// The most recent generation observed in the the CertificateRequest metadata.
	// If the observedGeneration does not match generation, the CA has not processed the most
	// recent version of this request.
	ObservedGeneration uint32 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// Any error observed which prevented the CertificateRequest from being processed.
	// If the error is empty, the request has been processed successfully
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// The current state of the CertificateRequest workflow reported by the Issuer.
	State CertificateRequestStatus_State `protobuf:"varint,3,opt,name=state,proto3,enum=certificates.smh.solo.io.CertificateRequestStatus_State" json:"state,omitempty"`
	// The signed intermediate certificate issued by the CA.
	SignedCertificate []byte `protobuf:"bytes,4,opt,name=signed_certificate,json=signedCertificate,proto3" json:"signed_certificate,omitempty"`
	// The root CA used by the CA to sign the certificate.
	SigningRootCa        []byte   `protobuf:"bytes,5,opt,name=signing_root_ca,json=signingRootCa,proto3" json:"signing_root_ca,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateRequestStatus) Reset()         { *m = CertificateRequestStatus{} }
func (m *CertificateRequestStatus) String() string { return proto.CompactTextString(m) }
func (*CertificateRequestStatus) ProtoMessage()    {}
func (*CertificateRequestStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbb195e6f26dce40, []int{1}
}
func (m *CertificateRequestStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateRequestStatus.Unmarshal(m, b)
}
func (m *CertificateRequestStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateRequestStatus.Marshal(b, m, deterministic)
}
func (m *CertificateRequestStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateRequestStatus.Merge(m, src)
}
func (m *CertificateRequestStatus) XXX_Size() int {
	return xxx_messageInfo_CertificateRequestStatus.Size(m)
}
func (m *CertificateRequestStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateRequestStatus.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateRequestStatus proto.InternalMessageInfo

func (m *CertificateRequestStatus) GetObservedGeneration() uint32 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *CertificateRequestStatus) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CertificateRequestStatus) GetState() CertificateRequestStatus_State {
	if m != nil {
		return m.State
	}
	return CertificateRequestStatus_PENDING
}

func (m *CertificateRequestStatus) GetSignedCertificate() []byte {
	if m != nil {
		return m.SignedCertificate
	}
	return nil
}

func (m *CertificateRequestStatus) GetSigningRootCa() []byte {
	if m != nil {
		return m.SigningRootCa
	}
	return nil
}

func init() {
	proto.RegisterEnum("certificates.smh.solo.io.CertificateRequestStatus_State", CertificateRequestStatus_State_name, CertificateRequestStatus_State_value)
	proto.RegisterType((*CertificateRequestSpec)(nil), "certificates.smh.solo.io.CertificateRequestSpec")
	proto.RegisterType((*CertificateRequestStatus)(nil), "certificates.smh.solo.io.CertificateRequestStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/certificates/certificate_request.proto", fileDescriptor_bbb195e6f26dce40)
}

var fileDescriptor_bbb195e6f26dce40 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x49, 0xa1, 0x03, 0xcc, 0x06, 0xc5, 0x4c, 0x28, 0x6c, 0xd2, 0xa8, 0x7a, 0x81, 0x7a,
	0x53, 0x5b, 0x2b, 0x37, 0x5c, 0x21, 0xc1, 0xda, 0x8d, 0x4a, 0x10, 0xa1, 0xe4, 0x06, 0x71, 0x13,
	0x39, 0xe9, 0x99, 0x63, 0xad, 0xc9, 0x09, 0xb6, 0x93, 0x67, 0xe2, 0x11, 0x78, 0x1e, 0x9e, 0x04,
	0xc5, 0xce, 0x46, 0x04, 0x4c, 0xe2, 0xca, 0x3e, 0x3e, 0x7f, 0xbe, 0xfc, 0xff, 0xb1, 0x49, 0x24,
	0x95, 0x2d, 0x9a, 0x8c, 0xe5, 0x58, 0x72, 0x83, 0x3b, 0x5c, 0x28, 0xe4, 0x06, 0x74, 0xab, 0x72,
	0x58, 0x94, 0x60, 0x8a, 0x45, 0xd1, 0x64, 0x5c, 0xd4, 0x8a, 0xe7, 0xa0, 0xad, 0xba, 0x54, 0xb9,
	0xb0, 0x60, 0x86, 0x45, 0xaa, 0xe1, 0x5b, 0x03, 0xc6, 0xb2, 0x5a, 0xa3, 0x45, 0x1a, 0x0e, 0x75,
	0xcc, 0x94, 0x05, 0xeb, 0xa8, 0x4c, 0xe1, 0xd1, 0xb1, 0xb9, 0x6a, 0x97, 0x1e, 0x85, 0x1a, 0x78,
	0x7b, 0xea, 0x56, 0xff, 0xd9, 0xd1, 0x4b, 0x89, 0x28, 0x77, 0xc0, 0x5d, 0x95, 0x35, 0x97, 0xdc,
	0xaa, 0x12, 0x8c, 0x15, 0x65, 0xdd, 0x0b, 0x4e, 0xfe, 0x14, 0x6c, 0x1b, 0x2d, 0xac, 0xc2, 0xaa,
	0xef, 0x1f, 0x4a, 0x94, 0xe8, 0xb6, 0xbc, 0xdb, 0xf9, 0xd3, 0xd9, 0x17, 0xf2, 0xfc, 0xec, 0xb7,
	0x9f, 0xd8, 0x3b, 0x4d, 0x6a, 0xc8, 0xe9, 0x5b, 0x72, 0x3c, 0x0c, 0x61, 0x94, 0xac, 0x54, 0x25,
	0xaf, 0xc3, 0x84, 0xc1, 0x34, 0x98, 0xef, 0xc7, 0x2f, 0x06, 0x92, 0xc4, 0x2b, 0x7a, 0xc6, 0xec,
	0xc7, 0x88, 0x84, 0xff, 0x40, 0x5b, 0x61, 0x1b, 0x43, 0x39, 0x79, 0x86, 0x59, 0x37, 0x45, 0xd8,
	0xa6, 0x12, 0x2a, 0xf0, 0x4e, 0x1d, 0xf4, 0x20, 0xa6, 0xd7, 0xad, 0x8b, 0x9b, 0x0e, 0x3d, 0x24,
	0x63, 0xd0, 0x1a, 0x75, 0x38, 0x9a, 0x06, 0xf3, 0x87, 0xb1, 0x2f, 0x68, 0x44, 0xc6, 0xc6, 0x0a,
	0x0b, 0xe1, 0xdd, 0x69, 0x30, 0x7f, 0xbc, 0x7c, 0xc3, 0x6e, 0x9b, 0x2d, 0xbb, 0xcd, 0x09, 0xeb,
	0x16, 0x88, 0x3d, 0x86, 0x2e, 0x08, 0xed, 0x72, 0xc2, 0x36, 0x1d, 0x80, 0xc2, 0x7b, 0x2e, 0xea,
	0x53, 0xdf, 0x19, 0x80, 0xe8, 0x2b, 0xf2, 0xe4, 0x66, 0x2c, 0x88, 0x36, 0xcd, 0x45, 0x38, 0x76,
	0xda, 0x83, 0xfe, 0x38, 0x46, 0xb4, 0x67, 0x62, 0xc6, 0xc8, 0xd8, 0xfd, 0x86, 0x3e, 0x22, 0xf7,
	0x3f, 0xaf, 0xa3, 0xd5, 0x26, 0xba, 0x98, 0xdc, 0xa1, 0xfb, 0xe4, 0xc1, 0xf9, 0x26, 0xda, 0x24,
	0x1f, 0xd6, 0xab, 0x49, 0x40, 0x09, 0xd9, 0x3b, 0x7f, 0xb7, 0xf9, 0xb8, 0x5e, 0x4d, 0x46, 0xef,
	0x93, 0xef, 0x3f, 0x4f, 0x82, 0xaf, 0x9f, 0xfe, 0xe7, 0xe1, 0xd5, 0x57, 0xf2, 0xaf, 0xc7, 0x37,
	0x0c, 0xce, 0xdb, 0x53, 0xb1, 0xab, 0x0b, 0xb1, 0xcc, 0xf6, 0xdc, 0x85, 0xbf, 0xfe, 0x15, 0x00,
	0x00, 0xff, 0xff, 0x20, 0x6a, 0x9a, 0xe0, 0xd0, 0x02, 0x00, 0x00,
}

func (this *CertificateRequestSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CertificateRequestSpec)
	if !ok {
		that2, ok := that.(CertificateRequestSpec)
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
	if !bytes.Equal(this.CertificateSigningRequest, that1.CertificateSigningRequest) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CertificateRequestStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CertificateRequestStatus)
	if !ok {
		that2, ok := that.(CertificateRequestStatus)
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
	if this.ObservedGeneration != that1.ObservedGeneration {
		return false
	}
	if this.Error != that1.Error {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if !bytes.Equal(this.SignedCertificate, that1.SignedCertificate) {
		return false
	}
	if !bytes.Equal(this.SigningRootCa, that1.SigningRootCa) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
