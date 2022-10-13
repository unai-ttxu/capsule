// Copyright 2020-2021 Clastix Labs
// SPDX-License-Identifier: Apache-2.0

package v1beta2

import (
	"fmt"
	"strconv"
	"strings"

	"sigs.k8s.io/controller-runtime/pkg/conversion"

	capsulev1beta1 "github.com/clastix/capsule/api/v1beta1"
)

func (in *Tenant) ConvertFrom(raw conversion.Hub) error {
	src, ok := raw.(*capsulev1beta1.Tenant)
	if !ok {
		return fmt.Errorf("expected *capsulev1beta1.Tenant, got %T", raw)
	}

	annotations := src.GetAnnotations()
	if annotations == nil {
		annotations = map[string]string{}
	}

	in.ObjectMeta = src.ObjectMeta
	in.Spec.Owners = make(OwnerListSpec, 0, len(src.Spec.Owners))

	for index, owner := range src.Spec.Owners {
		proxySettings := make([]ProxySettings, 0, len(owner.ProxyOperations))

		for _, proxyOp := range owner.ProxyOperations {
			ops := make([]ProxyOperation, 0, len(proxyOp.Operations))

			for _, op := range proxyOp.Operations {
				ops = append(ops, ProxyOperation(op))
			}

			proxySettings = append(proxySettings, ProxySettings{
				Kind:       ProxyServiceKind(proxyOp.Kind),
				Operations: ops,
			})
		}

		in.Spec.Owners = append(in.Spec.Owners, OwnerSpec{
			Kind:            OwnerKind(owner.Kind),
			Name:            owner.Name,
			ClusterRoles:    owner.GetRoles(*src, index),
			ProxyOperations: proxySettings,
		})
	}

	if nsOpts := src.Spec.NamespaceOptions; nsOpts != nil {
		in.Spec.NamespaceOptions = &NamespaceOptions{}

		in.Spec.NamespaceOptions.Quota = src.Spec.NamespaceOptions.Quota

		in.Spec.NamespaceOptions.AdditionalMetadata = nsOpts.AdditionalMetadata

		if value, found := annotations[capsulev1beta1.ForbiddenNamespaceLabelsAnnotation]; found {
			in.Spec.NamespaceOptions.ForbiddenLabels.Exact = strings.Split(value, ",")

			delete(annotations, capsulev1beta1.ForbiddenNamespaceLabelsAnnotation)
		}

		if value, found := annotations[capsulev1beta1.ForbiddenNamespaceLabelsRegexpAnnotation]; found {
			in.Spec.NamespaceOptions.ForbiddenLabels.Regex = value

			delete(annotations, capsulev1beta1.ForbiddenNamespaceLabelsRegexpAnnotation)
		}

		if value, found := annotations[capsulev1beta1.ForbiddenNamespaceAnnotationsAnnotation]; found {
			in.Spec.NamespaceOptions.ForbiddenAnnotations.Exact = strings.Split(value, ",")

			delete(annotations, capsulev1beta1.ForbiddenNamespaceAnnotationsAnnotation)
		}

		if value, found := annotations[capsulev1beta1.ForbiddenNamespaceAnnotationsRegexpAnnotation]; found {
			in.Spec.NamespaceOptions.ForbiddenAnnotations.Regex = value

			delete(annotations, capsulev1beta1.ForbiddenNamespaceAnnotationsRegexpAnnotation)
		}
	}

	in.Spec.ServiceOptions = src.Spec.ServiceOptions
	in.Spec.StorageClasses = src.Spec.StorageClasses

	if scope := src.Spec.IngressOptions.HostnameCollisionScope; len(scope) > 0 {
		in.Spec.IngressOptions.HostnameCollisionScope = scope
	}

	v, found := annotations[capsulev1beta1.DenyWildcard]
	if found {
		value, err := strconv.ParseBool(v)
		if err == nil {
			in.Spec.IngressOptions.AllowWildcardHostnames = !value

			delete(annotations, capsulev1beta1.DenyWildcard)
		}
	}

	if ingressClass := src.Spec.IngressOptions.AllowedClasses; ingressClass != nil {
		in.Spec.IngressOptions.AllowedClasses = ingressClass
	}

	if hostnames := src.Spec.IngressOptions.AllowedHostnames; hostnames != nil {
		in.Spec.IngressOptions.AllowedHostnames = hostnames
	}

	in.Spec.ContainerRegistries = src.Spec.ContainerRegistries
	in.Spec.NodeSelector = src.Spec.NodeSelector
	in.Spec.NetworkPolicies = src.Spec.NetworkPolicies
	in.Spec.LimitRanges = src.Spec.LimitRanges
	in.Spec.ResourceQuota = src.Spec.ResourceQuota
	in.Spec.AdditionalRoleBindings = src.Spec.AdditionalRoleBindings
	in.Spec.ImagePullPolicies = src.Spec.ImagePullPolicies
	in.Spec.PriorityClasses = src.Spec.PriorityClasses

	if v, found := annotations["capsule.clastix.io/cordon"]; found {
		value, err := strconv.ParseBool(v)
		if err == nil {
			delete(annotations, "capsule.clastix.io/cordon")
		}

		in.Spec.Cordoned = value
	}

	if _, found := annotations[capsulev1beta1.ProtectedTenantAnnotation]; found {
		in.Spec.PreventDeletion = true

		delete(annotations, capsulev1beta1.ProtectedTenantAnnotation)
	}

	in.SetAnnotations(annotations)

	in.Status.Namespaces = src.Status.Namespaces
	in.Status.Size = src.Status.Size
	in.Status.State = tenantState(src.Status.State)

	return nil
}

func (in *Tenant) ConvertTo(raw conversion.Hub) error {
	dst, ok := raw.(*capsulev1beta1.Tenant)
	if !ok {
		return fmt.Errorf("expected *capsulev1beta1.Tenant, got %T", raw)
	}

	annotations := in.GetAnnotations()
	if annotations == nil {
		annotations = map[string]string{}
	}

	dst.ObjectMeta = in.ObjectMeta
	dst.Spec.Owners = make(capsulev1beta1.OwnerListSpec, 0, len(in.Spec.Owners))

	for index, owner := range in.Spec.Owners {
		proxySettings := make([]capsulev1beta1.ProxySettings, 0, len(owner.ProxyOperations))

		for _, proxyOp := range owner.ProxyOperations {
			ops := make([]capsulev1beta1.ProxyOperation, 0, len(proxyOp.Operations))

			for _, op := range proxyOp.Operations {
				ops = append(ops, capsulev1beta1.ProxyOperation(op))
			}

			proxySettings = append(proxySettings, capsulev1beta1.ProxySettings{
				Kind:       capsulev1beta1.ProxyServiceKind(proxyOp.Kind),
				Operations: ops,
			})
		}

		dst.Spec.Owners = append(dst.Spec.Owners, capsulev1beta1.OwnerSpec{
			Kind:            capsulev1beta1.OwnerKind(owner.Kind),
			Name:            owner.Name,
			ProxyOperations: proxySettings,
		})

		if clusterRoles := owner.ClusterRoles; len(clusterRoles) > 0 {
			annotations[fmt.Sprintf("%s/%d", capsulev1beta1.ClusterRoleNamesAnnotation, index)] = strings.Join(owner.ClusterRoles, ",")
		}
	}

	if nsOpts := in.Spec.NamespaceOptions; nsOpts != nil {
		dst.Spec.NamespaceOptions = &capsulev1beta1.NamespaceOptions{}

		dst.Spec.NamespaceOptions.Quota = nsOpts.Quota
		dst.Spec.NamespaceOptions.AdditionalMetadata = nsOpts.AdditionalMetadata

		if exact := nsOpts.ForbiddenAnnotations.Exact; len(exact) > 0 {
			annotations[capsulev1beta1.ForbiddenNamespaceAnnotationsAnnotation] = strings.Join(exact, ",")
		}

		if regex := nsOpts.ForbiddenAnnotations.Regex; len(regex) > 0 {
			annotations[capsulev1beta1.ForbiddenNamespaceAnnotationsRegexpAnnotation] = regex
		}

		if exact := nsOpts.ForbiddenLabels.Exact; len(exact) > 0 {
			annotations[capsulev1beta1.ForbiddenNamespaceLabelsAnnotation] = strings.Join(exact, ",")
		}

		if regex := nsOpts.ForbiddenLabels.Regex; len(regex) > 0 {
			annotations[capsulev1beta1.ForbiddenNamespaceLabelsRegexpAnnotation] = regex
		}
	}

	dst.Spec.ServiceOptions = in.Spec.ServiceOptions
	dst.Spec.StorageClasses = in.Spec.StorageClasses

	dst.Spec.IngressOptions.HostnameCollisionScope = in.Spec.IngressOptions.HostnameCollisionScope

	if allowed := in.Spec.IngressOptions.AllowedClasses; allowed != nil {
		dst.Spec.IngressOptions.AllowedClasses = allowed
	}

	if allowed := in.Spec.IngressOptions.AllowedHostnames; allowed != nil {
		dst.Spec.IngressOptions.AllowedHostnames = allowed
	}

	annotations[capsulev1beta1.DenyWildcard] = fmt.Sprintf("%t", !in.Spec.IngressOptions.AllowWildcardHostnames)

	if allowed := in.Spec.ContainerRegistries; allowed != nil {
		dst.Spec.ContainerRegistries = allowed
	}

	dst.Spec.NodeSelector = in.Spec.NodeSelector
	dst.Spec.NetworkPolicies = in.Spec.NetworkPolicies
	dst.Spec.LimitRanges = in.Spec.LimitRanges
	dst.Spec.ResourceQuota = in.Spec.ResourceQuota
	dst.Spec.AdditionalRoleBindings = in.Spec.AdditionalRoleBindings
	dst.Spec.ImagePullPolicies = in.Spec.ImagePullPolicies
	dst.Spec.PriorityClasses = in.Spec.PriorityClasses

	if in.Spec.PreventDeletion {
		annotations[capsulev1beta1.ProtectedTenantAnnotation] = "true" //nolint:goconst
	}

	if in.Spec.Cordoned {
		annotations["capsule.clastix.io/cordon"] = "true"
	}

	dst.SetAnnotations(annotations)

	return nil
}