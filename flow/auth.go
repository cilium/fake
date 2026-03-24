// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	flowpb "github.com/cilium/cilium/api/v1/flow"
)

// AuthType implements FlowFaker for flowfaker.
func (f *flowfaker) AuthType() flowpb.AuthType {
	return flowpb.AuthType(f.IntN(len(flowpb.AuthType_name))) //nolint:gosec
}
