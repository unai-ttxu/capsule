// +build !ignore_autogenerated

// Copyright 2020-2021 Clastix Labs
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalMetadataSpec) DeepCopyInto(out *AdditionalMetadataSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalMetadataSpec.
func (in *AdditionalMetadataSpec) DeepCopy() *AdditionalMetadataSpec {
	if in == nil {
		return nil
	}
	out := new(AdditionalMetadataSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalRoleBindingsSpec) DeepCopyInto(out *AdditionalRoleBindingsSpec) {
	*out = *in
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]v1.Subject, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalRoleBindingsSpec.
func (in *AdditionalRoleBindingsSpec) DeepCopy() *AdditionalRoleBindingsSpec {
	if in == nil {
		return nil
	}
	out := new(AdditionalRoleBindingsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedListSpec) DeepCopyInto(out *AllowedListSpec) {
	*out = *in
	if in.Exact != nil {
		in, out := &in.Exact, &out.Exact
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedListSpec.
func (in *AllowedListSpec) DeepCopy() *AllowedListSpec {
	if in == nil {
		return nil
	}
	out := new(AllowedListSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedServices) DeepCopyInto(out *AllowedServices) {
	*out = *in
	if in.NodePort != nil {
		in, out := &in.NodePort, &out.NodePort
		*out = new(bool)
		**out = **in
	}
	if in.ExternalName != nil {
		in, out := &in.ExternalName, &out.ExternalName
		*out = new(bool)
		**out = **in
	}
	if in.LoadBalancer != nil {
		in, out := &in.LoadBalancer, &out.LoadBalancer
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedServices.
func (in *AllowedServices) DeepCopy() *AllowedServices {
	if in == nil {
		return nil
	}
	out := new(AllowedServices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ByKindAndName) DeepCopyInto(out *ByKindAndName) {
	{
		in := &in
		*out = make(ByKindAndName, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ByKindAndName.
func (in ByKindAndName) DeepCopy() ByKindAndName {
	if in == nil {
		return nil
	}
	out := new(ByKindAndName)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalServiceIPsSpec) DeepCopyInto(out *ExternalServiceIPsSpec) {
	*out = *in
	if in.Allowed != nil {
		in, out := &in.Allowed, &out.Allowed
		*out = make([]AllowedIP, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalServiceIPsSpec.
func (in *ExternalServiceIPsSpec) DeepCopy() *ExternalServiceIPsSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalServiceIPsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForbiddenListSpec) DeepCopyInto(out *ForbiddenListSpec) {
	*out = *in
	if in.Exact != nil {
		in, out := &in.Exact, &out.Exact
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForbiddenListSpec.
func (in *ForbiddenListSpec) DeepCopy() *ForbiddenListSpec {
	if in == nil {
		return nil
	}
	out := new(ForbiddenListSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressOptions) DeepCopyInto(out *IngressOptions) {
	*out = *in
	if in.AllowedClasses != nil {
		in, out := &in.AllowedClasses, &out.AllowedClasses
		*out = new(AllowedListSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AllowedHostnames != nil {
		in, out := &in.AllowedHostnames, &out.AllowedHostnames
		*out = new(AllowedListSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressOptions.
func (in *IngressOptions) DeepCopy() *IngressOptions {
	if in == nil {
		return nil
	}
	out := new(IngressOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitRangesSpec) DeepCopyInto(out *LimitRangesSpec) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]corev1.LimitRangeSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitRangesSpec.
func (in *LimitRangesSpec) DeepCopy() *LimitRangesSpec {
	if in == nil {
		return nil
	}
	out := new(LimitRangesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceOptions) DeepCopyInto(out *NamespaceOptions) {
	*out = *in
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		*out = new(int32)
		**out = **in
	}
	if in.AdditionalMetadata != nil {
		in, out := &in.AdditionalMetadata, &out.AdditionalMetadata
		*out = new(AdditionalMetadataSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceOptions.
func (in *NamespaceOptions) DeepCopy() *NamespaceOptions {
	if in == nil {
		return nil
	}
	out := new(NamespaceOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicySpec) DeepCopyInto(out *NetworkPolicySpec) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]networkingv1.NetworkPolicySpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicySpec.
func (in *NetworkPolicySpec) DeepCopy() *NetworkPolicySpec {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in OwnerListSpec) DeepCopyInto(out *OwnerListSpec) {
	{
		in := &in
		*out = make(OwnerListSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OwnerListSpec.
func (in OwnerListSpec) DeepCopy() OwnerListSpec {
	if in == nil {
		return nil
	}
	out := new(OwnerListSpec)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OwnerSpec) DeepCopyInto(out *OwnerSpec) {
	*out = *in
	if in.ProxyOperations != nil {
		in, out := &in.ProxyOperations, &out.ProxyOperations
		*out = make([]ProxySettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OwnerSpec.
func (in *OwnerSpec) DeepCopy() *OwnerSpec {
	if in == nil {
		return nil
	}
	out := new(OwnerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxySettings) DeepCopyInto(out *ProxySettings) {
	*out = *in
	if in.Operations != nil {
		in, out := &in.Operations, &out.Operations
		*out = make([]ProxyOperation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxySettings.
func (in *ProxySettings) DeepCopy() *ProxySettings {
	if in == nil {
		return nil
	}
	out := new(ProxySettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaSpec) DeepCopyInto(out *ResourceQuotaSpec) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]corev1.ResourceQuotaSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaSpec.
func (in *ResourceQuotaSpec) DeepCopy() *ResourceQuotaSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceQuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceOptions) DeepCopyInto(out *ServiceOptions) {
	*out = *in
	if in.AdditionalMetadata != nil {
		in, out := &in.AdditionalMetadata, &out.AdditionalMetadata
		*out = new(AdditionalMetadataSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AllowedServices != nil {
		in, out := &in.AllowedServices, &out.AllowedServices
		*out = new(AllowedServices)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalServiceIPs != nil {
		in, out := &in.ExternalServiceIPs, &out.ExternalServiceIPs
		*out = new(ExternalServiceIPsSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceOptions.
func (in *ServiceOptions) DeepCopy() *ServiceOptions {
	if in == nil {
		return nil
	}
	out := new(ServiceOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tenant) DeepCopyInto(out *Tenant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tenant.
func (in *Tenant) DeepCopy() *Tenant {
	if in == nil {
		return nil
	}
	out := new(Tenant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Tenant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantList) DeepCopyInto(out *TenantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Tenant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantList.
func (in *TenantList) DeepCopy() *TenantList {
	if in == nil {
		return nil
	}
	out := new(TenantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantSpec) DeepCopyInto(out *TenantSpec) {
	*out = *in
	if in.Owners != nil {
		in, out := &in.Owners, &out.Owners
		*out = make(OwnerListSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NamespaceOptions != nil {
		in, out := &in.NamespaceOptions, &out.NamespaceOptions
		*out = new(NamespaceOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceOptions != nil {
		in, out := &in.ServiceOptions, &out.ServiceOptions
		*out = new(ServiceOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageClasses != nil {
		in, out := &in.StorageClasses, &out.StorageClasses
		*out = new(AllowedListSpec)
		(*in).DeepCopyInto(*out)
	}
	in.IngressOptions.DeepCopyInto(&out.IngressOptions)
	if in.ContainerRegistries != nil {
		in, out := &in.ContainerRegistries, &out.ContainerRegistries
		*out = new(AllowedListSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NetworkPolicies != nil {
		in, out := &in.NetworkPolicies, &out.NetworkPolicies
		*out = new(NetworkPolicySpec)
		(*in).DeepCopyInto(*out)
	}
	if in.LimitRanges != nil {
		in, out := &in.LimitRanges, &out.LimitRanges
		*out = new(LimitRangesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceQuota != nil {
		in, out := &in.ResourceQuota, &out.ResourceQuota
		*out = new(ResourceQuotaSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalRoleBindings != nil {
		in, out := &in.AdditionalRoleBindings, &out.AdditionalRoleBindings
		*out = make([]AdditionalRoleBindingsSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImagePullPolicies != nil {
		in, out := &in.ImagePullPolicies, &out.ImagePullPolicies
		*out = make([]ImagePullPolicySpec, len(*in))
		copy(*out, *in)
	}
	if in.PriorityClasses != nil {
		in, out := &in.PriorityClasses, &out.PriorityClasses
		*out = new(AllowedListSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantSpec.
func (in *TenantSpec) DeepCopy() *TenantSpec {
	if in == nil {
		return nil
	}
	out := new(TenantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantStatus) DeepCopyInto(out *TenantStatus) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantStatus.
func (in *TenantStatus) DeepCopy() *TenantStatus {
	if in == nil {
		return nil
	}
	out := new(TenantStatus)
	in.DeepCopyInto(out)
	return out
}
