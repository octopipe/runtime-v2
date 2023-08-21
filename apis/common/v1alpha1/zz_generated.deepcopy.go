//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSProviderConfig) DeepCopyInto(out *AWSProviderConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSProviderConfig.
func (in *AWSProviderConfig) DeepCopy() *AWSProviderConfig {
	if in == nil {
		return nil
	}
	out := new(AWSProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Error) DeepCopyInto(out *Error) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Error.
func (in *Error) DeepCopy() *Error {
	if in == nil {
		return nil
	}
	out := new(Error)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutionStatus) DeepCopyInto(out *ExecutionStatus) {
	*out = *in
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]TaskExecutionStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Error = in.Error
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutionStatus.
func (in *ExecutionStatus) DeepCopy() *ExecutionStatus {
	if in == nil {
		return nil
	}
	out := new(ExecutionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Infra) DeepCopyInto(out *Infra) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Infra.
func (in *Infra) DeepCopy() *Infra {
	if in == nil {
		return nil
	}
	out := new(Infra)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Infra) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraList) DeepCopyInto(out *InfraList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Infra, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraList.
func (in *InfraList) DeepCopy() *InfraList {
	if in == nil {
		return nil
	}
	out := new(InfraList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfraList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraRunnerConfig) DeepCopyInto(out *InfraRunnerConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraRunnerConfig.
func (in *InfraRunnerConfig) DeepCopy() *InfraRunnerConfig {
	if in == nil {
		return nil
	}
	out := new(InfraRunnerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraSpec) DeepCopyInto(out *InfraSpec) {
	*out = *in
	out.ProviderConfigRef = in.ProviderConfigRef
	out.RunnerConfig = in.RunnerConfig
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]InfraTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraSpec.
func (in *InfraSpec) DeepCopy() *InfraSpec {
	if in == nil {
		return nil
	}
	out := new(InfraSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraStatus) DeepCopyInto(out *InfraStatus) {
	*out = *in
	in.LastExecution.DeepCopyInto(&out.LastExecution)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraStatus.
func (in *InfraStatus) DeepCopy() *InfraStatus {
	if in == nil {
		return nil
	}
	out := new(InfraStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraTask) DeepCopyInto(out *InfraTask) {
	*out = *in
	if in.Depends != nil {
		in, out := &in.Depends, &out.Depends
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Terraform = in.Terraform
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make([]InfraTaskInput, len(*in))
		copy(*out, *in)
	}
	if in.TaskOutputs != nil {
		in, out := &in.TaskOutputs, &out.TaskOutputs
		*out = make([]InfraTaskOutput, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make([]InfraTaskOutputItem, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraTask.
func (in *InfraTask) DeepCopy() *InfraTask {
	if in == nil {
		return nil
	}
	out := new(InfraTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraTaskInput) DeepCopyInto(out *InfraTaskInput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraTaskInput.
func (in *InfraTaskInput) DeepCopy() *InfraTaskInput {
	if in == nil {
		return nil
	}
	out := new(InfraTaskInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraTaskOutput) DeepCopyInto(out *InfraTaskOutput) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InfraTaskOutputItem, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraTaskOutput.
func (in *InfraTaskOutput) DeepCopy() *InfraTaskOutput {
	if in == nil {
		return nil
	}
	out := new(InfraTaskOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraTaskOutputItem) DeepCopyInto(out *InfraTaskOutputItem) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraTaskOutputItem.
func (in *InfraTaskOutputItem) DeepCopy() *InfraTaskOutputItem {
	if in == nil {
		return nil
	}
	out := new(InfraTaskOutputItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraTaskRef) DeepCopyInto(out *InfraTaskRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraTaskRef.
func (in *InfraTaskRef) DeepCopy() *InfraTaskRef {
	if in == nil {
		return nil
	}
	out := new(InfraTaskRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderConfig) DeepCopyInto(out *ProviderConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderConfig.
func (in *ProviderConfig) DeepCopy() *ProviderConfig {
	if in == nil {
		return nil
	}
	out := new(ProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderConfigList) DeepCopyInto(out *ProviderConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProviderConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderConfigList.
func (in *ProviderConfigList) DeepCopy() *ProviderConfigList {
	if in == nil {
		return nil
	}
	out := new(ProviderConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderConfigSpec) DeepCopyInto(out *ProviderConfigSpec) {
	*out = *in
	out.AWSConfig = in.AWSConfig
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderConfigSpec.
func (in *ProviderConfigSpec) DeepCopy() *ProviderConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ProviderConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderConfigStatus) DeepCopyInto(out *ProviderConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderConfigStatus.
func (in *ProviderConfigStatus) DeepCopy() *ProviderConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ProviderConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ref) DeepCopyInto(out *Ref) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ref.
func (in *Ref) DeepCopy() *Ref {
	if in == nil {
		return nil
	}
	out := new(Ref)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskExecutionStatus) DeepCopyInto(out *TaskExecutionStatus) {
	*out = *in
	if in.Depends != nil {
		in, out := &in.Depends, &out.Depends
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make([]InfraTaskInput, len(*in))
		copy(*out, *in)
	}
	out.Task = in.Task
	if in.TaskOutputs != nil {
		in, out := &in.TaskOutputs, &out.TaskOutputs
		*out = make([]InfraTaskOutput, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Error = in.Error
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskExecutionStatus.
func (in *TaskExecutionStatus) DeepCopy() *TaskExecutionStatus {
	if in == nil {
		return nil
	}
	out := new(TaskExecutionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskOutput) DeepCopyInto(out *TaskOutput) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskOutput.
func (in *TaskOutput) DeepCopy() *TaskOutput {
	if in == nil {
		return nil
	}
	out := new(TaskOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TaskOutput) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskOutputList) DeepCopyInto(out *TaskOutputList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TaskOutput, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskOutputList.
func (in *TaskOutputList) DeepCopy() *TaskOutputList {
	if in == nil {
		return nil
	}
	out := new(TaskOutputList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TaskOutputList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskOutputSpec) DeepCopyInto(out *TaskOutputSpec) {
	*out = *in
	out.Infra = in.Infra
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make([]TaskOutputSpecItem, len(*in))
		copy(*out, *in)
	}
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskOutputSpec.
func (in *TaskOutputSpec) DeepCopy() *TaskOutputSpec {
	if in == nil {
		return nil
	}
	out := new(TaskOutputSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskOutputSpecItem) DeepCopyInto(out *TaskOutputSpecItem) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskOutputSpecItem.
func (in *TaskOutputSpecItem) DeepCopy() *TaskOutputSpecItem {
	if in == nil {
		return nil
	}
	out := new(TaskOutputSpecItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskOutputSpecItemSecret) DeepCopyInto(out *TaskOutputSpecItemSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskOutputSpecItemSecret.
func (in *TaskOutputSpecItemSecret) DeepCopy() *TaskOutputSpecItemSecret {
	if in == nil {
		return nil
	}
	out := new(TaskOutputSpecItemSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskOutputStatus) DeepCopyInto(out *TaskOutputStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskOutputStatus.
func (in *TaskOutputStatus) DeepCopy() *TaskOutputStatus {
	if in == nil {
		return nil
	}
	out := new(TaskOutputStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskStatus) DeepCopyInto(out *TaskStatus) {
	*out = *in
	out.Terraform = in.Terraform
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskStatus.
func (in *TaskStatus) DeepCopy() *TaskStatus {
	if in == nil {
		return nil
	}
	out := new(TaskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Terraform) DeepCopyInto(out *Terraform) {
	*out = *in
	out.CredentialsRef = in.CredentialsRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Terraform.
func (in *Terraform) DeepCopy() *Terraform {
	if in == nil {
		return nil
	}
	out := new(Terraform)
	in.DeepCopyInto(out)
	return out
}
