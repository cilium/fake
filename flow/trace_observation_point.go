// Copyright 2021-2022 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
