//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FluxConfig) DeepCopyInto(out *FluxConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Flux != nil {
		in, out := &in.Flux, &out.Flux
		*out = new(FluxInstallation)
		(*in).DeepCopyInto(*out)
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(Source)
		(*in).DeepCopyInto(*out)
	}
	if in.Kustomization != nil {
		in, out := &in.Kustomization, &out.Kustomization
		*out = new(Kustomization)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalSecretResourceNames != nil {
		in, out := &in.AdditionalSecretResourceNames, &out.AdditionalSecretResourceNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluxConfig.
func (in *FluxConfig) DeepCopy() *FluxConfig {
	if in == nil {
		return nil
	}
	out := new(FluxConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FluxConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FluxInstallation) DeepCopyInto(out *FluxInstallation) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluxInstallation.
func (in *FluxInstallation) DeepCopy() *FluxInstallation {
	if in == nil {
		return nil
	}
	out := new(FluxInstallation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kustomization) DeepCopyInto(out *Kustomization) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kustomization.
func (in *Kustomization) DeepCopy() *Kustomization {
	if in == nil {
		return nil
	}
	out := new(Kustomization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	if in.SecretResourceName != nil {
		in, out := &in.SecretResourceName, &out.SecretResourceName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}
