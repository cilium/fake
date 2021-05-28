// Copyright (C) Isovalent, Inc. - All Rights Reserved.
//
// NOTICE: All information contained herein is, and remains the property of
// Isovalent Inc and its suppliers, if any. The intellectual and technical
// concepts contained herein are proprietary to Isovalent Inc and its suppliers
// and may be covered by U.S. and Foreign Patents, patents in process, and are
// protected by trade secret or copyright law.  Dissemination of this information
// or reproduction of this material is strictly forbidden unless prior written
// permission is obtained from Isovalent Inc.

package fake

import (
	"math/rand"

	flowpb "github.com/cilium/cilium/api/v1/flow"
	"github.com/cilium/cilium/pkg/monitor/api"
)

var allCiliumEventTypes = []int32{
	api.MessageTypeUnspec,
	api.MessageTypeDrop,
	api.MessageTypeDebug,
	api.MessageTypeCapture,
	api.MessageTypeTrace,
	api.MessageTypePolicyVerdict,
	api.MessageTypeAccessLog,
	api.MessageTypeAgent,
}

// CiliumEventType generates a random CiliumEventType.
func CiliumEventType() *flowpb.CiliumEventType {
	typ := allCiliumEventTypes[rand.Intn(len(allCiliumEventTypes))]
	if typ == api.MessageTypeUnspec {
		return nil
	}
	return &flowpb.CiliumEventType{
		Type: typ,
		// NOTE: AgentNotify* are the most numerous.
		SubType: int32(rand.Intn(13)),
	}
}
