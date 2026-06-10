// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	flowpb "github.com/cilium/cilium/api/v1/flow"
)

// Policies generates a random slice of network policies (0 to 3 entries).
func (f *FlowFaker) Policies() []*flowpb.Policy {
	n := f.IntN(4)
	policies := make([]*flowpb.Policy, n)
	for i := range n {
		policies[i] = &flowpb.Policy{
			Name:      f.Name(),
			Namespace: f.K8sNamespace(),
			Labels:    f.K8sLabels(),
			Revision:  f.Uint64() % 100,
		}
	}
	return policies
}
