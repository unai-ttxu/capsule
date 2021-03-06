//+build e2e

/*
Copyright 2020 Clastix Labs.

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

package e2e

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/clastix/capsule/api/v1alpha1"
)

var _ = Describe("creating a Namespace as Tenant owner with custom --capsule-group", func() {
	tnt := &v1alpha1.Tenant{
		ObjectMeta: metav1.ObjectMeta{
			Name: "tenant-assigned-custom-group",
		},
		Spec: v1alpha1.TenantSpec{
			Owner: v1alpha1.OwnerSpec{
				Name: "alice",
				Kind: "User",
			},
		},
	}

	JustBeforeEach(func() {
		EventuallyCreation(func() error {
			tnt.ResourceVersion = ""
			return k8sClient.Create(context.TODO(), tnt)
		}).Should(Succeed())
	})
	JustAfterEach(func() {
		Expect(k8sClient.Delete(context.TODO(), tnt)).Should(Succeed())
	})

	It("should fail using a User non matching the capsule-user-group flag", func() {
		args := append(defaulManagerPodArgs, []string{"--capsule-user-group=test"}...)
		ModifyCapsuleManagerPodArgs(args)
		CapsuleClusterGroupParam(podRecreationTimeoutInterval).Should(BeIdenticalTo("test"))
		ns := NewNamespace("cg-namespace-fail")
		NamespaceCreation(ns, tnt, podRecreationTimeoutInterval).ShouldNot(Succeed())
	})

	It("should succeed and be available in Tenant namespaces list", func() {
		ModifyCapsuleManagerPodArgs(defaulManagerPodArgs)
		CapsuleClusterGroupParam(podRecreationTimeoutInterval).Should(BeIdenticalTo("capsule.clastix.io"))
		ns := NewNamespace("cg-namespace")
		NamespaceCreation(ns, tnt, podRecreationTimeoutInterval).Should(Succeed())
		TenantNamespaceList(tnt, podRecreationTimeoutInterval).Should(ContainElement(ns.GetName()))
	})
})
