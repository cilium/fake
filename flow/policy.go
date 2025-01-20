// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	"math/rand"

	flowpb "github.com/cilium/cilium/api/v1/flow"
	"github.com/cilium/fake"
)

// Policies generates a list of random policy references.
func Policies() []*flowpb.Policy {
	policies := []*flowpb.Policy{}
	for range rand.Intn(4) {
		policies = append(policies, &flowpb.Policy{
			Name:      fake.Name(),
			Namespace: fake.K8sNamespace(),
			Labels:    fake.K8sLabels(),
			Revision:  rand.Uint64() % 100,
		})
	}
	return policies
}
