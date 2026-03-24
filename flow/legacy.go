// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	"sync"

	flowpb "github.com/cilium/cilium/api/v1/flow"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// globalFlowFaker is used by the legacy API of package level functions.
var (
	globalFlowFaker = NewFaker()
	globalMu        sync.Mutex
)

// AuthType generates a random AuthType.
func AuthType() flowpb.AuthType {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFlowFaker.AuthType()
}

// DropReason generates a DropReason. Options may be provided to customize the
// drop reasons to return.
func DropReason(options ...DropReasonOption) flowpb.DropReason {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFlowFaker.DropReason(options...)
}

// Endpoint generates a random Endpoint. Options may be provided to customize
// the endpoint to return.
func Endpoint(options ...EndpointOption) *flowpb.Endpoint {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFlowFaker.Endpoint(options...)
}

// EventType generates a random EventType.
func EventType() *flowpb.CiliumEventType {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFlowFaker.EventType()
}

// New generates a random Hubble Flow. Options may be provided to customize the
// flow.
func New(options ...Option) *flowpb.Flow {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFlowFaker.NewFlow(options...)
}

// ICMPv4 generates a random flow ICMPv4 struct.
func ICMPv4() *flowpb.ICMPv4 {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFlowFaker.ICMPv4()
}

// ICMPv6 generates a random flow ICMPv6 struct.
func ICMPv6() *flowpb.ICMPv6 {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFlowFaker.ICMPv6()
}

// IP generates a random flow IP struct.
func IP(options ...IPOption) *flowpb.IP {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFlowFaker.IP(options...)
}

// IsReply returns either nil, or a wrapped boolean value reprenting true, or a
// wrapped boolean value reprenting false, with equal probability.
func IsReply() *wrapperspb.BoolValue {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFlowFaker.IsReply()
}

// Policies generates a list of random policy references.
func Policies() []*flowpb.Policy {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFlowFaker.Policies()
}

// Layer4 generates a layer 4. If no option is provided, it will be TCP.
func Layer4(options ...Layer4Option) *flowpb.Layer4 {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFlowFaker.Layer4(options...)
}

// Service generates a random Service. Options may be provided to customize the
// service to return.
func Service(options ...ServiceOption) *flowpb.Service {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFlowFaker.Service(options...)
}

// TraceObservationPoint generates a random TraceObservationPoint.
func TraceObservationPoint() flowpb.TraceObservationPoint {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFlowFaker.TraceObservationPoint()
}

// TraceReason generates a random TraceReason.
func TraceReason() flowpb.TraceReason {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFlowFaker.TraceReason()
}

// TraceContext generates a TraceContext. Options may be provided to customize
// the returned object.
func TraceContext(options ...TraceContextOption) *flowpb.TraceContext {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFlowFaker.TraceContext(options...)
}

// TrafficDirection generates a random TrafficDirection.
func TrafficDirection() flowpb.TrafficDirection {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFlowFaker.TrafficDirection()
}

// Verdict generates a FORWARDED or DROPPPED verdict randomly. The probability
// of the verdict being FORWARDED can be set using
// WithVerdictForwardedProbability.
func Verdict(options ...VerdictOption) flowpb.Verdict {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFlowFaker.Verdict(options...)
}
