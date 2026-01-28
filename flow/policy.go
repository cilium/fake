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
	n := rand.Intn(4)
	policies := make([]*flowpb.Policy, n)
	for i := range n {
		policies[i] = &flowpb.Policy{
			Name:      fake.Name(),
			Namespace: fake.K8sNamespace(),
			Labels:    fake.K8sLabels(),
			Revision:  rand.Uint64() % 100,
		}
	}
	return policies
}
