// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	flowpb "github.com/cilium/cilium/api/v1/flow"
)

// TrafficDirection generates a random TrafficDirection.
func (f *flowfaker) TrafficDirection() flowpb.TrafficDirection {
	return flowpb.TrafficDirection(f.IntN(len(flowpb.TrafficDirection_name))) //nolint:gosec
}
