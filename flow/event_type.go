// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	"math/rand"

	flowpb "github.com/cilium/cilium/api/v1/flow"
	"github.com/cilium/cilium/pkg/monitor/api"
)

var allEventTypes = []int32{
	api.MessageTypeUnspec,
	api.MessageTypeDrop,
	api.MessageTypeDebug,
	api.MessageTypeCapture,
	api.MessageTypeTrace,
	api.MessageTypePolicyVerdict,
	api.MessageTypeAccessLog,
	api.MessageTypeAgent,
}

// EventType generates a random EventType.
func EventType() *flowpb.CiliumEventType {
	typ := allEventTypes[rand.Intn(len(allEventTypes))]
	if typ == api.MessageTypeUnspec {
		return nil
	}
	return &flowpb.CiliumEventType{
		Type: typ,
		// NOTE: AgentNotify* are the most numerous.
		SubType: int32(rand.Intn(13)), //nolint:gosec
	}
}
