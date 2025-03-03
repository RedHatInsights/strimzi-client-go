//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StrimziPodSet) DeepCopyInto(out *StrimziPodSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(StrimziPodSetSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(StrimziPodSetStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StrimziPodSet.
func (in *StrimziPodSet) DeepCopy() *StrimziPodSet {
	if in == nil {
		return nil
	}
	out := new(StrimziPodSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StrimziPodSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StrimziPodSetList) DeepCopyInto(out *StrimziPodSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StrimziPodSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StrimziPodSetList.
func (in *StrimziPodSetList) DeepCopy() *StrimziPodSetList {
	if in == nil {
		return nil
	}
	out := new(StrimziPodSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StrimziPodSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StrimziPodSetSpec) DeepCopyInto(out *StrimziPodSetSpec) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]v1.Pod, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Selector.DeepCopyInto(&out.Selector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StrimziPodSetSpec.
func (in *StrimziPodSetSpec) DeepCopy() *StrimziPodSetSpec {
	if in == nil {
		return nil
	}
	out := new(StrimziPodSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StrimziPodSetSpecSelector) DeepCopyInto(out *StrimziPodSetSpecSelector) {
	*out = *in
	if in.MatchExpressions != nil {
		in, out := &in.MatchExpressions, &out.MatchExpressions
		*out = make([]StrimziPodSetSpecSelectorMatchExpressionsElem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(StrimziPodSetSpecSelectorMatchLabels, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StrimziPodSetSpecSelector.
func (in *StrimziPodSetSpecSelector) DeepCopy() *StrimziPodSetSpecSelector {
	if in == nil {
		return nil
	}
	out := new(StrimziPodSetSpecSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StrimziPodSetSpecSelectorMatchExpressionsElem) DeepCopyInto(out *StrimziPodSetSpecSelectorMatchExpressionsElem) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StrimziPodSetSpecSelectorMatchExpressionsElem.
func (in *StrimziPodSetSpecSelectorMatchExpressionsElem) DeepCopy() *StrimziPodSetSpecSelectorMatchExpressionsElem {
	if in == nil {
		return nil
	}
	out := new(StrimziPodSetSpecSelectorMatchExpressionsElem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in StrimziPodSetSpecSelectorMatchLabels) DeepCopyInto(out *StrimziPodSetSpecSelectorMatchLabels) {
	{
		in := &in
		*out = make(StrimziPodSetSpecSelectorMatchLabels, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StrimziPodSetSpecSelectorMatchLabels.
func (in StrimziPodSetSpecSelectorMatchLabels) DeepCopy() StrimziPodSetSpecSelectorMatchLabels {
	if in == nil {
		return nil
	}
	out := new(StrimziPodSetSpecSelectorMatchLabels)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StrimziPodSetStatus) DeepCopyInto(out *StrimziPodSetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]StrimziPodSetStatusConditionsElem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CurrentPods != nil {
		in, out := &in.CurrentPods, &out.CurrentPods
		*out = new(int32)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int32)
		**out = **in
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = new(int32)
		**out = **in
	}
	if in.ReadyPods != nil {
		in, out := &in.ReadyPods, &out.ReadyPods
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StrimziPodSetStatus.
func (in *StrimziPodSetStatus) DeepCopy() *StrimziPodSetStatus {
	if in == nil {
		return nil
	}
	out := new(StrimziPodSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StrimziPodSetStatusConditionsElem) DeepCopyInto(out *StrimziPodSetStatusConditionsElem) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StrimziPodSetStatusConditionsElem.
func (in *StrimziPodSetStatusConditionsElem) DeepCopy() *StrimziPodSetStatusConditionsElem {
	if in == nil {
		return nil
	}
	out := new(StrimziPodSetStatusConditionsElem)
	in.DeepCopyInto(out)
	return out
}
