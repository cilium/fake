// Copyright (C) Isovalent, Inc. - All Rights Reserved.
//
// NOTICE: All information contained herein is, and remains the property of
// Isovalent Inc and its suppliers, if any. The intellectual and technical
// concepts contained herein are proprietary to Isovalent Inc and its suppliers
// and may be covered by U.S. and Foreign Patents, patents in process, and are
// protected by trade secret or copyright law.  Dissemination of this information
// or reproduction of this material is strictly forbidden unless prior written
// permission is obtained from Isovalent Inc.

package flow

import (
	"math/rand"

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
}

// TraceObservationPoint generates a random TraceObservationPoint.
func TraceObservationPoint() flowpb.TraceObservationPoint {
	return allTraceObservationPoints[rand.Intn(len(allTraceObservationPoints))]
}
