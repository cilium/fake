// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	"math/rand"

	flowpb "github.com/cilium/cilium/api/v1/flow"
)

// AuthType generates a random AuthType.
func AuthType() flowpb.AuthType {
	return flowpb.AuthType(rand.Intn(len(flowpb.AuthType_name) + 1))
}
