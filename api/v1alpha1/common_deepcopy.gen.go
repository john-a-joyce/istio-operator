// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1alpha1/common.proto

package v1alpha1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	_ "k8s.io/api/core/v1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using NamespacedName within kubernetes types, where deepcopy-gen is used.
func (in *NamespacedName) DeepCopyInto(out *NamespacedName) {
	p := proto.Clone(in).(*NamespacedName)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedName. Required by controller-gen.
func (in *NamespacedName) DeepCopy() *NamespacedName {
	if in == nil {
		return nil
	}
	out := new(NamespacedName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedName. Required by controller-gen.
func (in *NamespacedName) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using BaseK8SResourceConfigurationWithImage within kubernetes types, where deepcopy-gen is used.
func (in *BaseK8SResourceConfigurationWithImage) DeepCopyInto(out *BaseK8SResourceConfigurationWithImage) {
	p := proto.Clone(in).(*BaseK8SResourceConfigurationWithImage)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseK8SResourceConfigurationWithImage. Required by controller-gen.
func (in *BaseK8SResourceConfigurationWithImage) DeepCopy() *BaseK8SResourceConfigurationWithImage {
	if in == nil {
		return nil
	}
	out := new(BaseK8SResourceConfigurationWithImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new BaseK8SResourceConfigurationWithImage. Required by controller-gen.
func (in *BaseK8SResourceConfigurationWithImage) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using BaseK8SResourceConfigurationWithReplicas within kubernetes types, where deepcopy-gen is used.
func (in *BaseK8SResourceConfigurationWithReplicas) DeepCopyInto(out *BaseK8SResourceConfigurationWithReplicas) {
	p := proto.Clone(in).(*BaseK8SResourceConfigurationWithReplicas)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseK8SResourceConfigurationWithReplicas. Required by controller-gen.
func (in *BaseK8SResourceConfigurationWithReplicas) DeepCopy() *BaseK8SResourceConfigurationWithReplicas {
	if in == nil {
		return nil
	}
	out := new(BaseK8SResourceConfigurationWithReplicas)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new BaseK8SResourceConfigurationWithReplicas. Required by controller-gen.
func (in *BaseK8SResourceConfigurationWithReplicas) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using BaseK8SResourceConfigurationWithHPA within kubernetes types, where deepcopy-gen is used.
func (in *BaseK8SResourceConfigurationWithHPA) DeepCopyInto(out *BaseK8SResourceConfigurationWithHPA) {
	p := proto.Clone(in).(*BaseK8SResourceConfigurationWithHPA)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseK8SResourceConfigurationWithHPA. Required by controller-gen.
func (in *BaseK8SResourceConfigurationWithHPA) DeepCopy() *BaseK8SResourceConfigurationWithHPA {
	if in == nil {
		return nil
	}
	out := new(BaseK8SResourceConfigurationWithHPA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new BaseK8SResourceConfigurationWithHPA. Required by controller-gen.
func (in *BaseK8SResourceConfigurationWithHPA) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using BaseK8SResourceConfigurationWithHPAWithoutImage within kubernetes types, where deepcopy-gen is used.
func (in *BaseK8SResourceConfigurationWithHPAWithoutImage) DeepCopyInto(out *BaseK8SResourceConfigurationWithHPAWithoutImage) {
	p := proto.Clone(in).(*BaseK8SResourceConfigurationWithHPAWithoutImage)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseK8SResourceConfigurationWithHPAWithoutImage. Required by controller-gen.
func (in *BaseK8SResourceConfigurationWithHPAWithoutImage) DeepCopy() *BaseK8SResourceConfigurationWithHPAWithoutImage {
	if in == nil {
		return nil
	}
	out := new(BaseK8SResourceConfigurationWithHPAWithoutImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new BaseK8SResourceConfigurationWithHPAWithoutImage. Required by controller-gen.
func (in *BaseK8SResourceConfigurationWithHPAWithoutImage) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using BaseK8SResourceConfiguration within kubernetes types, where deepcopy-gen is used.
func (in *BaseK8SResourceConfiguration) DeepCopyInto(out *BaseK8SResourceConfiguration) {
	p := proto.Clone(in).(*BaseK8SResourceConfiguration)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseK8SResourceConfiguration. Required by controller-gen.
func (in *BaseK8SResourceConfiguration) DeepCopy() *BaseK8SResourceConfiguration {
	if in == nil {
		return nil
	}
	out := new(BaseK8SResourceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new BaseK8SResourceConfiguration. Required by controller-gen.
func (in *BaseK8SResourceConfiguration) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
