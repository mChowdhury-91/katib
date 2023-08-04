/*
Copyright 2022 The Kubeflow Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
)

// SourceSpecApplyConfiguration represents an declarative configuration of the SourceSpec type for use
// with apply.
type SourceSpecApplyConfiguration struct {
	HttpGet        *v1.HTTPGetAction                 `json:"httpGet,omitempty"`
	FileSystemPath *FileSystemPathApplyConfiguration `json:"fileSystemPath,omitempty"`
	Filter         *FilterSpecApplyConfiguration     `json:"filter,omitempty"`
}

// SourceSpecApplyConfiguration constructs an declarative configuration of the SourceSpec type for use with
// apply.
func SourceSpec() *SourceSpecApplyConfiguration {
	return &SourceSpecApplyConfiguration{}
}

// WithHttpGet sets the HttpGet field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HttpGet field is set to the value of the last call.
func (b *SourceSpecApplyConfiguration) WithHttpGet(value v1.HTTPGetAction) *SourceSpecApplyConfiguration {
	b.HttpGet = &value
	return b
}

// WithFileSystemPath sets the FileSystemPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FileSystemPath field is set to the value of the last call.
func (b *SourceSpecApplyConfiguration) WithFileSystemPath(value *FileSystemPathApplyConfiguration) *SourceSpecApplyConfiguration {
	b.FileSystemPath = value
	return b
}

// WithFilter sets the Filter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Filter field is set to the value of the last call.
func (b *SourceSpecApplyConfiguration) WithFilter(value *FilterSpecApplyConfiguration) *SourceSpecApplyConfiguration {
	b.Filter = value
	return b
}