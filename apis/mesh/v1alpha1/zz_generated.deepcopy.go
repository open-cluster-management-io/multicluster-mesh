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
func (in *AccessLogging) DeepCopyInto(out *AccessLogging) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessLogging.
func (in *AccessLogging) DeepCopy() *AccessLogging {
	if in == nil {
		return nil
	}
	out := new(AccessLogging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mesh) DeepCopyInto(out *Mesh) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mesh.
func (in *Mesh) DeepCopy() *Mesh {
	if in == nil {
		return nil
	}
	out := new(Mesh)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Mesh) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshConfig) DeepCopyInto(out *MeshConfig) {
	*out = *in
	if in.ProxyConfig != nil {
		in, out := &in.ProxyConfig, &out.ProxyConfig
		*out = new(ProxyConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshConfig.
func (in *MeshConfig) DeepCopy() *MeshConfig {
	if in == nil {
		return nil
	}
	out := new(MeshConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshControlPlane) DeepCopyInto(out *MeshControlPlane) {
	*out = *in
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Peers != nil {
		in, out := &in.Peers, &out.Peers
		*out = make([]Peer, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshControlPlane.
func (in *MeshControlPlane) DeepCopy() *MeshControlPlane {
	if in == nil {
		return nil
	}
	out := new(MeshControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshDeployment) DeepCopyInto(out *MeshDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshDeployment.
func (in *MeshDeployment) DeepCopy() *MeshDeployment {
	if in == nil {
		return nil
	}
	out := new(MeshDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MeshDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshDeploymentList) DeepCopyInto(out *MeshDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MeshDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshDeploymentList.
func (in *MeshDeploymentList) DeepCopy() *MeshDeploymentList {
	if in == nil {
		return nil
	}
	out := new(MeshDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MeshDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshDeploymentSpec) DeepCopyInto(out *MeshDeploymentSpec) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ControlPlane != nil {
		in, out := &in.ControlPlane, &out.ControlPlane
		*out = new(MeshControlPlane)
		(*in).DeepCopyInto(*out)
	}
	if in.MeshConfig != nil {
		in, out := &in.MeshConfig, &out.MeshConfig
		*out = new(MeshConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.MeshMemberRoll != nil {
		in, out := &in.MeshMemberRoll, &out.MeshMemberRoll
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshDeploymentSpec.
func (in *MeshDeploymentSpec) DeepCopy() *MeshDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(MeshDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshDeploymentStatus) DeepCopyInto(out *MeshDeploymentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshDeploymentStatus.
func (in *MeshDeploymentStatus) DeepCopy() *MeshDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(MeshDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshFederation) DeepCopyInto(out *MeshFederation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshFederation.
func (in *MeshFederation) DeepCopy() *MeshFederation {
	if in == nil {
		return nil
	}
	out := new(MeshFederation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MeshFederation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshFederationList) DeepCopyInto(out *MeshFederationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MeshFederation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshFederationList.
func (in *MeshFederationList) DeepCopy() *MeshFederationList {
	if in == nil {
		return nil
	}
	out := new(MeshFederationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MeshFederationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshFederationSpec) DeepCopyInto(out *MeshFederationSpec) {
	*out = *in
	if in.MeshPeers != nil {
		in, out := &in.MeshPeers, &out.MeshPeers
		*out = make([]MeshPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TrustConfig != nil {
		in, out := &in.TrustConfig, &out.TrustConfig
		*out = new(TrustConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshFederationSpec.
func (in *MeshFederationSpec) DeepCopy() *MeshFederationSpec {
	if in == nil {
		return nil
	}
	out := new(MeshFederationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshFederationStatus) DeepCopyInto(out *MeshFederationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshFederationStatus.
func (in *MeshFederationStatus) DeepCopy() *MeshFederationStatus {
	if in == nil {
		return nil
	}
	out := new(MeshFederationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshList) DeepCopyInto(out *MeshList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Mesh, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshList.
func (in *MeshList) DeepCopy() *MeshList {
	if in == nil {
		return nil
	}
	out := new(MeshList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MeshList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshPeer) DeepCopyInto(out *MeshPeer) {
	*out = *in
	if in.Peers != nil {
		in, out := &in.Peers, &out.Peers
		*out = make([]Peer, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshPeer.
func (in *MeshPeer) DeepCopy() *MeshPeer {
	if in == nil {
		return nil
	}
	out := new(MeshPeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshSpec) DeepCopyInto(out *MeshSpec) {
	*out = *in
	if in.ControlPlane != nil {
		in, out := &in.ControlPlane, &out.ControlPlane
		*out = new(MeshControlPlane)
		(*in).DeepCopyInto(*out)
	}
	if in.MeshConfig != nil {
		in, out := &in.MeshConfig, &out.MeshConfig
		*out = new(MeshConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.MeshMemberRoll != nil {
		in, out := &in.MeshMemberRoll, &out.MeshMemberRoll
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshSpec.
func (in *MeshSpec) DeepCopy() *MeshSpec {
	if in == nil {
		return nil
	}
	out := new(MeshSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshStatus) DeepCopyInto(out *MeshStatus) {
	*out = *in
	in.Readiness.DeepCopyInto(&out.Readiness)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshStatus.
func (in *MeshStatus) DeepCopy() *MeshStatus {
	if in == nil {
		return nil
	}
	out := new(MeshStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Peer) DeepCopyInto(out *Peer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Peer.
func (in *Peer) DeepCopy() *Peer {
	if in == nil {
		return nil
	}
	out := new(Peer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyConfig) DeepCopyInto(out *ProxyConfig) {
	*out = *in
	if in.AccessLogging != nil {
		in, out := &in.AccessLogging, &out.AccessLogging
		*out = new(AccessLogging)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyConfig.
func (in *ProxyConfig) DeepCopy() *ProxyConfig {
	if in == nil {
		return nil
	}
	out := new(ProxyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrustConfig) DeepCopyInto(out *TrustConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrustConfig.
func (in *TrustConfig) DeepCopy() *TrustConfig {
	if in == nil {
		return nil
	}
	out := new(TrustConfig)
	in.DeepCopyInto(out)
	return out
}
