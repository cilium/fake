// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	"math/rand/v2"

	flowpb "github.com/cilium/cilium/api/v1/flow"
)

// AuthType generates a random AuthType.
func AuthType() flowpb.AuthType {
	return flowpb.AuthType(rand.IntN(len(flowpb.AuthType_name))) //nolint:gosec
}
