// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: endorsement.proto

package endorsement

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VMLaunchEndorsement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerializedUefiGolden []byte `protobuf:"bytes,1,opt,name=serialized_uefi_golden,json=serializedUefiGolden,proto3" json:"serialized_uefi_golden,omitempty"`
	Signature            []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *VMLaunchEndorsement) Reset() {
	*x = VMLaunchEndorsement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_endorsement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMLaunchEndorsement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMLaunchEndorsement) ProtoMessage() {}

func (x *VMLaunchEndorsement) ProtoReflect() protoreflect.Message {
	mi := &file_endorsement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMLaunchEndorsement.ProtoReflect.Descriptor instead.
func (*VMLaunchEndorsement) Descriptor() ([]byte, []int) {
	return file_endorsement_proto_rawDescGZIP(), []int{0}
}

func (x *VMLaunchEndorsement) GetSerializedUefiGolden() []byte {
	if x != nil {
		return x.SerializedUefiGolden
	}
	return nil
}

func (x *VMLaunchEndorsement) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type VMGoldenMeasurement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The changelist number this UEFI was built from.
	ClSpec uint64 `protobuf:"varint,2,opt,name=cl_spec,json=clSpec,proto3" json:"cl_spec,omitempty"`
	// The commit hash this UEFI was built from.
	Commit []byte `protobuf:"bytes,3,opt,name=commit,proto3" json:"commit,omitempty"`
	// DER format certificate of the key that signed this document.
	Cert []byte `protobuf:"bytes,4,opt,name=cert,proto3" json:"cert,omitempty"`
	// SHA-384 digest of the UEFI binary without TEE-specifics about launch.
	Digest []byte `protobuf:"bytes,5,opt,name=digest,proto3" json:"digest,omitempty"`
	// A sequence of PEM-encoded certificates of keys used in cert in Root ...
	// final intermediate order. The last certificate will have signed cert.
	CaBundle []byte    `protobuf:"bytes,6,opt,name=ca_bundle,json=caBundle,proto3" json:"ca_bundle,omitempty"`
	SevSnp   *VMSevSnp `protobuf:"bytes,7,opt,name=sev_snp,json=sevSnp,proto3" json:"sev_snp,omitempty"`
	Tdx      *VMTdx    `protobuf:"bytes,8,opt,name=tdx,proto3" json:"tdx,omitempty"`
}

func (x *VMGoldenMeasurement) Reset() {
	*x = VMGoldenMeasurement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_endorsement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMGoldenMeasurement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMGoldenMeasurement) ProtoMessage() {}

func (x *VMGoldenMeasurement) ProtoReflect() protoreflect.Message {
	mi := &file_endorsement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMGoldenMeasurement.ProtoReflect.Descriptor instead.
func (*VMGoldenMeasurement) Descriptor() ([]byte, []int) {
	return file_endorsement_proto_rawDescGZIP(), []int{1}
}

func (x *VMGoldenMeasurement) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *VMGoldenMeasurement) GetClSpec() uint64 {
	if x != nil {
		return x.ClSpec
	}
	return 0
}

func (x *VMGoldenMeasurement) GetCommit() []byte {
	if x != nil {
		return x.Commit
	}
	return nil
}

func (x *VMGoldenMeasurement) GetCert() []byte {
	if x != nil {
		return x.Cert
	}
	return nil
}

func (x *VMGoldenMeasurement) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *VMGoldenMeasurement) GetCaBundle() []byte {
	if x != nil {
		return x.CaBundle
	}
	return nil
}

func (x *VMGoldenMeasurement) GetSevSnp() *VMSevSnp {
	if x != nil {
		return x.SevSnp
	}
	return nil
}

func (x *VMGoldenMeasurement) GetTdx() *VMTdx {
	if x != nil {
		return x.Tdx
	}
	return nil
}

type VMSevSnp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Google-reported security version number of this UEFI on SEV-SNP.
	Svn uint32 `protobuf:"varint,1,opt,name=svn,proto3" json:"svn,omitempty"`
	// Expected MEASUREMENT report field values given [key]-many VMSAs at launch.
	Measurements map[uint32][]byte `protobuf:"bytes,2,rep,name=measurements,proto3" json:"measurements,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // bytes size 48
	// A UUID that Google uses for its CVM UEFIs
	FamilyId []byte `protobuf:"bytes,3,opt,name=family_id,json=familyId,proto3" json:"family_id,omitempty"` // size 16
	// A UUID to name this specific release of the UEFI image.
	ImageId []byte `protobuf:"bytes,4,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"` // size 16
	// The launch policy that verifiers should expect with this UEFI.
	Policy uint64 `protobuf:"varint,5,opt,name=policy,proto3" json:"policy,omitempty"`
	// Optional. PEM-encoded certs for Identity..Author..Root. If a singleton,
	// only an Id-key is used.
	CaBundle []byte `protobuf:"bytes,6,opt,name=ca_bundle,json=caBundle,proto3" json:"ca_bundle,omitempty"`
	// Expected SEV-SNP launch measurement as reported by igvmmeasure (and its hex
	// digits are decoded). Use igvmmeasure's --native_zero.
	SvsmMeasurement []byte `protobuf:"bytes,7,opt,name=svsm_measurement,json=svsmMeasurement,proto3" json:"svsm_measurement,omitempty"`
}

func (x *VMSevSnp) Reset() {
	*x = VMSevSnp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_endorsement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMSevSnp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMSevSnp) ProtoMessage() {}

func (x *VMSevSnp) ProtoReflect() protoreflect.Message {
	mi := &file_endorsement_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMSevSnp.ProtoReflect.Descriptor instead.
func (*VMSevSnp) Descriptor() ([]byte, []int) {
	return file_endorsement_proto_rawDescGZIP(), []int{2}
}

func (x *VMSevSnp) GetSvn() uint32 {
	if x != nil {
		return x.Svn
	}
	return 0
}

func (x *VMSevSnp) GetMeasurements() map[uint32][]byte {
	if x != nil {
		return x.Measurements
	}
	return nil
}

func (x *VMSevSnp) GetFamilyId() []byte {
	if x != nil {
		return x.FamilyId
	}
	return nil
}

func (x *VMSevSnp) GetImageId() []byte {
	if x != nil {
		return x.ImageId
	}
	return nil
}

func (x *VMSevSnp) GetPolicy() uint64 {
	if x != nil {
		return x.Policy
	}
	return 0
}

func (x *VMSevSnp) GetCaBundle() []byte {
	if x != nil {
		return x.CaBundle
	}
	return nil
}

func (x *VMSevSnp) GetSvsmMeasurement() []byte {
	if x != nil {
		return x.SvsmMeasurement
	}
	return nil
}

type VMTdx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Google-reported security version number of this UEFI on TDX.
	Svn uint32 `protobuf:"varint,1,opt,name=svn,proto3" json:"svn,omitempty"`
	// Expected MRTD report field values given legal configurations.
	Measurements []*VMTdx_Measurement `protobuf:"bytes,2,rep,name=measurements,proto3" json:"measurements,omitempty"`
}

func (x *VMTdx) Reset() {
	*x = VMTdx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_endorsement_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMTdx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMTdx) ProtoMessage() {}

func (x *VMTdx) ProtoReflect() protoreflect.Message {
	mi := &file_endorsement_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMTdx.ProtoReflect.Descriptor instead.
func (*VMTdx) Descriptor() ([]byte, []int) {
	return file_endorsement_proto_rawDescGZIP(), []int{3}
}

func (x *VMTdx) GetSvn() uint32 {
	if x != nil {
		return x.Svn
	}
	return 0
}

func (x *VMTdx) GetMeasurements() []*VMTdx_Measurement {
	if x != nil {
		return x.Measurements
	}
	return nil
}

type VMTdx_Measurement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The amount of RAM in GiB provided to the VM. This is relevant to the
	// construction of the measured TDHOB page that includes memory regions'
	// resource attributes.
	RamGib uint32 `protobuf:"varint,1,opt,name=ram_gib,json=ramGib,proto3" json:"ram_gib,omitempty"`
	// If true, EFI_UNACCEPTED_MEMORY not presented to guest.
	// All memory is accepted by the firmware. Relevant to the TDHOB page
	// since the resource attribute will include
	// EFI_RESOURCE_ATTRIBUTE_NEEDS_EARLY_ACCEPT.
	EarlyAccept bool `protobuf:"varint,2,opt,name=early_accept,json=earlyAccept,proto3" json:"early_accept,omitempty"`
	// The SHA-384 digest of the measurement operations for the VM at launch.
	Mrtd []byte `protobuf:"bytes,3,opt,name=mrtd,proto3" json:"mrtd,omitempty"`
}

func (x *VMTdx_Measurement) Reset() {
	*x = VMTdx_Measurement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_endorsement_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMTdx_Measurement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMTdx_Measurement) ProtoMessage() {}

func (x *VMTdx_Measurement) ProtoReflect() protoreflect.Message {
	mi := &file_endorsement_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMTdx_Measurement.ProtoReflect.Descriptor instead.
func (*VMTdx_Measurement) Descriptor() ([]byte, []int) {
	return file_endorsement_proto_rawDescGZIP(), []int{3, 0}
}

func (x *VMTdx_Measurement) GetRamGib() uint32 {
	if x != nil {
		return x.RamGib
	}
	return 0
}

func (x *VMTdx_Measurement) GetEarlyAccept() bool {
	if x != nil {
		return x.EarlyAccept
	}
	return false
}

func (x *VMTdx_Measurement) GetMrtd() []byte {
	if x != nil {
		return x.Mrtd
	}
	return nil
}

var File_endorsement_proto protoreflect.FileDescriptor

var file_endorsement_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x76, 0x6d, 0x6d, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x13, 0x56, 0x4d, 0x4c, 0x61, 0x75, 0x6e, 0x63,
	0x68, 0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x16,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x75, 0x65, 0x66, 0x69, 0x5f,
	0x67, 0x6f, 0x6c, 0x64, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x14, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x55, 0x65, 0x66, 0x69, 0x47, 0x6f, 0x6c, 0x64,
	0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x22, 0xa7, 0x02, 0x0a, 0x13, 0x56, 0x4d, 0x47, 0x6f, 0x6c, 0x64, 0x65, 0x6e, 0x4d, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6c, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x61, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x63, 0x61, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x32, 0x0a, 0x07,
	0x73, 0x65, 0x76, 0x5f, 0x73, 0x6e, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x76, 0x6d, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x56, 0x4d, 0x53, 0x65, 0x76, 0x53, 0x6e, 0x70, 0x52, 0x06, 0x73, 0x65, 0x76, 0x53, 0x6e, 0x70,
	0x12, 0x28, 0x0a, 0x03, 0x74, 0x64, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x76, 0x6d, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x56, 0x4d, 0x54, 0x64, 0x78, 0x52, 0x03, 0x74, 0x64, 0x78, 0x22, 0xc6, 0x02, 0x0a, 0x08, 0x56,
	0x4d, 0x53, 0x65, 0x76, 0x53, 0x6e, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x76, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x76, 0x6e, 0x12, 0x4f, 0x0a, 0x0c, 0x6d, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x76, 0x6d, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x56, 0x4d, 0x53, 0x65, 0x76, 0x53, 0x6e, 0x70, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x6d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x66,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61,
	0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63,
	0x61, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x76, 0x73, 0x6d, 0x5f,
	0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0f, 0x73, 0x76, 0x73, 0x6d, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x1a, 0x3f, 0x0a, 0x11, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xc0, 0x01, 0x0a, 0x05, 0x56, 0x4d, 0x54, 0x64, 0x78, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x76, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x76, 0x6e, 0x12,
	0x46, 0x0a, 0x0c, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x76, 0x6d,
	0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x4d, 0x54, 0x64, 0x78, 0x2e, 0x4d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x6d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x5d, 0x0a, 0x0b, 0x4d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x61, 0x6d, 0x5f, 0x67, 0x69,
	0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x6d, 0x47, 0x69, 0x62, 0x12,
	0x21, 0x0a, 0x0c, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x72, 0x74, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x6d, 0x72, 0x74, 0x64, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x67, 0x63, 0x65, 0x2d,
	0x74, 0x63, 0x62, 0x2d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_endorsement_proto_rawDescOnce sync.Once
	file_endorsement_proto_rawDescData = file_endorsement_proto_rawDesc
)

func file_endorsement_proto_rawDescGZIP() []byte {
	file_endorsement_proto_rawDescOnce.Do(func() {
		file_endorsement_proto_rawDescData = protoimpl.X.CompressGZIP(file_endorsement_proto_rawDescData)
	})
	return file_endorsement_proto_rawDescData
}

var file_endorsement_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_endorsement_proto_goTypes = []interface{}{
	(*VMLaunchEndorsement)(nil),   // 0: cloud_vmm_proto.VMLaunchEndorsement
	(*VMGoldenMeasurement)(nil),   // 1: cloud_vmm_proto.VMGoldenMeasurement
	(*VMSevSnp)(nil),              // 2: cloud_vmm_proto.VMSevSnp
	(*VMTdx)(nil),                 // 3: cloud_vmm_proto.VMTdx
	nil,                           // 4: cloud_vmm_proto.VMSevSnp.MeasurementsEntry
	(*VMTdx_Measurement)(nil),     // 5: cloud_vmm_proto.VMTdx.Measurement
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_endorsement_proto_depIdxs = []int32{
	6, // 0: cloud_vmm_proto.VMGoldenMeasurement.timestamp:type_name -> google.protobuf.Timestamp
	2, // 1: cloud_vmm_proto.VMGoldenMeasurement.sev_snp:type_name -> cloud_vmm_proto.VMSevSnp
	3, // 2: cloud_vmm_proto.VMGoldenMeasurement.tdx:type_name -> cloud_vmm_proto.VMTdx
	4, // 3: cloud_vmm_proto.VMSevSnp.measurements:type_name -> cloud_vmm_proto.VMSevSnp.MeasurementsEntry
	5, // 4: cloud_vmm_proto.VMTdx.measurements:type_name -> cloud_vmm_proto.VMTdx.Measurement
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_endorsement_proto_init() }
func file_endorsement_proto_init() {
	if File_endorsement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_endorsement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMLaunchEndorsement); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_endorsement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMGoldenMeasurement); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_endorsement_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMSevSnp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_endorsement_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMTdx); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_endorsement_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMTdx_Measurement); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_endorsement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_endorsement_proto_goTypes,
		DependencyIndexes: file_endorsement_proto_depIdxs,
		MessageInfos:      file_endorsement_proto_msgTypes,
	}.Build()
	File_endorsement_proto = out.File
	file_endorsement_proto_rawDesc = nil
	file_endorsement_proto_goTypes = nil
	file_endorsement_proto_depIdxs = nil
}
