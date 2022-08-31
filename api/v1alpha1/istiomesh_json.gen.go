// Code generated by protoc-gen-jsonshim. DO NOT EDIT.
package v1alpha1

import (
	bytes "bytes"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

// MarshalJSON is a custom marshaler for IstioMeshSpec
func (this *IstioMeshSpec) MarshalJSON() ([]byte, error) {
	str, err := IstiomeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for IstioMeshSpec
func (this *IstioMeshSpec) UnmarshalJSON(b []byte) error {
	return IstiomeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for IstioMeshStatus
func (this *IstioMeshStatus) MarshalJSON() ([]byte, error) {
	str, err := IstiomeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for IstioMeshStatus
func (this *IstioMeshStatus) UnmarshalJSON(b []byte) error {
	return IstiomeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	IstiomeshMarshaler   = &jsonpb.Marshaler{}
	IstiomeshUnmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}
)
