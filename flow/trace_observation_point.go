// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	flowpb "github.com/cilium/cilium/api/v1/flow"
)

var allTraceObservationPoints = []flowpb.TraceObservationPoint{
	flowpb.TraceObservationPoint_UNKNOWN_POINT,
	flowpb.TraceObservationPoint_TO_PROXY,
	flowpb.TraceObservationPoint_TO_HOST,
	flowpb.TraceObservationPoint_TO_STACK,
	flowpb.TraceObservationPoint_TO_OVERLAY,
	flowpb.TraceObservationPoint_TO_ENDPOINT,
	flowpb.TraceObservationPoint_FROM_ENDPOINT,
	flowpb.TraceObservationPoint_FROM_PROXY,
	flowpb.TraceObservationPoint_FROM_HOST,
	flowpb.TraceObservationPoint_FROM_STACK,
	flowpb.TraceObservationPoint_FROM_OVERLAY,
	flowpb.TraceObservationPoint_FROM_NETWORK,
	flowpb.TraceObservationPoint_TO_NETWORK,
	flowpb.TraceObservationPoint_FROM_CRYPTO,
	flowpb.TraceObservationPoint_TO_CRYPTO,
}

// TraceObservationPoint implements FlowFaker for flowfaker.
func (f *flowfaker) TraceObservationPoint() flowpb.TraceObservationPoint {
	return allTraceObservationPoints[f.IntN(len(allTraceObservationPoints))]
}
