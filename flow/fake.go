// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	"encoding/binary"
	"io"
	randv2 "math/rand/v2"

	"github.com/cilium/fake"
	"google.golang.org/protobuf/types/known/wrapperspb"

	flowpb "github.com/cilium/cilium/api/v1/flow"
)

// FlowFaker is the main interface exposed to generate fake flow data.
type FlowFaker interface { //nolint:interfacebloat
	// AuthType generates a random AuthType.
	AuthType() flowpb.AuthType
	// DropReason generates a DropReason.
	DropReason(options ...DropReasonOption) flowpb.DropReason
	// Endpoint generates a random Endpoint.
	Endpoint(options ...EndpointOption) *flowpb.Endpoint
	// EventType generates a random EventType.
	EventType() *flowpb.CiliumEventType
	// New generates a random Hubble Flow.
	NewFlow(options ...Option) *flowpb.Flow
	// ICMPv4 generates a random flow ICMPv4 struct.
	ICMPv4() *flowpb.ICMPv4
	// ICMPv6 generates a random flow ICMPv6 struct.
	ICMPv6() *flowpb.ICMPv6
	// IP generates a random flow IP struct.
	IP(options ...IPOption) *flowpb.IP
	// IsReply returns either nil, or a wrapped boolean value reprenting true,
	// or a wrapped boolean value reprenting false, with equal probability.
	IsReply() *wrapperspb.BoolValue
	// Policies generates a list of random policy references.
	Policies() []*flowpb.Policy
	// Layer4 generates a layer 4. If no option is provided, it will be TCP.
	Layer4(options ...Layer4Option) *flowpb.Layer4
	// Service generates a random Service.
	Service(options ...ServiceOption) *flowpb.Service
	// TraceObservationPoint generates a random TraceObservationPoint.
	TraceObservationPoint() flowpb.TraceObservationPoint
	// TraceReason generates a random TraceReason.
	TraceReason() flowpb.TraceReason
	// TraceContext generates a TraceContext.
	TraceContext(options ...TraceContextOption) *flowpb.TraceContext
	// TrafficDirection generates a random TrafficDirection.
	TrafficDirection() flowpb.TrafficDirection
	// Verdict generates a FORWARDED or DROPPPED verdict randomly.
	Verdict(options ...VerdictOption) flowpb.Verdict
}

// NewFaker creates a new Faker using a random seed.
func NewFaker() FlowFaker {
	// NOTE: We seed from the global math/rand/v2 source rather than
	// crypto/rand to avoid errors handling; the faker should not be used for
	// security-sensitive purposes. We use ChaCha8 rather than PCG as ChaCha8
	// implements io.Reader.
	var seed [32]byte
	for i := range 4 {
		binary.LittleEndian.PutUint64(seed[i*8:], randv2.Uint64())
	}
	return NewWithSource(randv2.NewChaCha8(seed))
}

// NewWithSource creates a new Faker using the given random source. Useful to
// control the faker output, e.g. for testing.
func NewWithSource(src fake.RandSourceReader) FlowFaker {
	return &flowfaker{
		Faker:  fake.NewWithSource(src),
		Rand:   randv2.New(src),
		Reader: src,
	}
}

// faker is a private struct implementing Faker.
type flowfaker struct {
	fake.Faker
	*randv2.Rand
	io.Reader
}
