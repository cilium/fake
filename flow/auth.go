// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	flowpb "github.com/cilium/cilium/api/v1/flow"
)

// AuthType generates a random authentication type.
func (f *FlowFaker) AuthType() flowpb.AuthType {
	return flowpb.AuthType(f.IntN(len(flowpb.AuthType_name))) //nolint:gosec
}
