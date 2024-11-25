// Copyright 2020-2023 Project Capsule Authors.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"context"
	"strings"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	capsulev1beta2 "github.com/projectcapsule/capsule/api/v1beta2"
	"github.com/projectcapsule/capsule/pkg/utils"
)

func IsCapsuleUser(ctx context.Context, req admission.Request, clt client.Client, userGroups []string, excludeUserGroups []string) bool {
	groupList := utils.NewUserGroupList(req.UserInfo.Groups)

	for _, group := range excludeUserGroups {
		if groupList.Find(group) {
			return false
		}
	}
	//nolint:nestif
	if sets.NewString(req.UserInfo.Groups...).Has("system:serviceaccounts") {
		parts := strings.Split(req.UserInfo.Username, ":")

		if len(parts) == 4 {
			targetNamespace := parts[2]

			tl := &capsulev1beta2.TenantList{}
			if err := clt.List(ctx, tl, client.MatchingFieldsSelector{Selector: fields.OneTermEqualSelector(".status.namespaces", targetNamespace)}); err != nil {
				return false
			}

			if len(tl.Items) == 1 {
				return true
			}
		}
	}

	for _, group := range userGroups {
		if groupList.Find(group) {
			return true
		}
	}

	return false
}
