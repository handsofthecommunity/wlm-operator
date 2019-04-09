// +build !ignore_autogenerated

// Copyright (c) 2019 Sylabs, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmJob) DeepCopyInto(out *SlurmJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmJob.
func (in *SlurmJob) DeepCopy() *SlurmJob {
	if in == nil {
		return nil
	}
	out := new(SlurmJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SlurmJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmJobList) DeepCopyInto(out *SlurmJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SlurmJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmJobList.
func (in *SlurmJobList) DeepCopy() *SlurmJobList {
	if in == nil {
		return nil
	}
	out := new(SlurmJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SlurmJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmJobResults) DeepCopyInto(out *SlurmJobResults) {
	*out = *in
	in.Mount.DeepCopyInto(&out.Mount)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmJobResults.
func (in *SlurmJobResults) DeepCopy() *SlurmJobResults {
	if in == nil {
		return nil
	}
	out := new(SlurmJobResults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmJobSSH) DeepCopyInto(out *SlurmJobSSH) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(v1.EnvVarSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(v1.EnvVarSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmJobSSH.
func (in *SlurmJobSSH) DeepCopy() *SlurmJobSSH {
	if in == nil {
		return nil
	}
	out := new(SlurmJobSSH)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmJobSpec) DeepCopyInto(out *SlurmJobSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(map[string]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SSH != nil {
		in, out := &in.SSH, &out.SSH
		*out = new(SlurmJobSSH)
		(*in).DeepCopyInto(*out)
	}
	if in.Results != nil {
		in, out := &in.Results, &out.Results
		*out = new(SlurmJobResults)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmJobSpec.
func (in *SlurmJobSpec) DeepCopy() *SlurmJobSpec {
	if in == nil {
		return nil
	}
	out := new(SlurmJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmJobStatus) DeepCopyInto(out *SlurmJobStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmJobStatus.
func (in *SlurmJobStatus) DeepCopy() *SlurmJobStatus {
	if in == nil {
		return nil
	}
	out := new(SlurmJobStatus)
	in.DeepCopyInto(out)
	return out
}