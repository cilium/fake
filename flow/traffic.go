// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	"math/rand"

	flowpb "github.com/cilium/cilium/api/v1/flow"
)

// TrafficDirection generates a random TrafficDirection.
func TrafficDirection() flowpb.TrafficDirection {
	return flowpb.TrafficDirection(rand.Intn(len(flowpb.TrafficDirection_name))) //nolint:gosec
}
