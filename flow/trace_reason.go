// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	flowpb "github.com/cilium/cilium/api/v1/flow"
)

var allTraceReasons = []flowpb.TraceReason{
	flowpb.TraceReason_TRACE_REASON_UNKNOWN,
	flowpb.TraceReason_NEW,
	flowpb.TraceReason_ESTABLISHED,
	flowpb.TraceReason_REPLY,
	flowpb.TraceReason_RELATED,
	// flowpb.TraceReason_REOPENED, -- Deprecated
	flowpb.TraceReason_SRV6_ENCAP,
	flowpb.TraceReason_SRV6_DECAP,
	// flowpb.TraceReason_ENCRYPT_OVERLAY, -- Deprecated
}

// TraceReason implements FlowFaker for flowfaker.
func (f *flowfaker) TraceReason() flowpb.TraceReason {
	return allTraceReasons[f.IntN(len(allTraceReasons))]
}
