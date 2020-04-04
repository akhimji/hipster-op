// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hipsterkind) DeepCopyInto(out *Hipsterkind) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hipsterkind.
func (in *Hipsterkind) DeepCopy() *Hipsterkind {
	if in == nil {
		return nil
	}
	out := new(Hipsterkind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Hipsterkind) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HipsterkindList) DeepCopyInto(out *HipsterkindList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Hipsterkind, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HipsterkindList.
func (in *HipsterkindList) DeepCopy() *HipsterkindList {
	if in == nil {
		return nil
	}
	out := new(HipsterkindList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HipsterkindList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HipsterkindSpec) DeepCopyInto(out *HipsterkindSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HipsterkindSpec.
func (in *HipsterkindSpec) DeepCopy() *HipsterkindSpec {
	if in == nil {
		return nil
	}
	out := new(HipsterkindSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HipsterkindStatus) DeepCopyInto(out *HipsterkindStatus) {
	*out = *in
	if in.PodNames != nil {
		in, out := &in.PodNames, &out.PodNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HipsterkindStatus.
func (in *HipsterkindStatus) DeepCopy() *HipsterkindStatus {
	if in == nil {
		return nil
	}
	out := new(HipsterkindStatus)
	in.DeepCopyInto(out)
	return out
}
