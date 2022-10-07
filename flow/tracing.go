// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	"encoding/hex"
	"math/rand"

	flowpb "github.com/cilium/cilium/api/v1/flow"
)

const traceIDLen = 16

type traceContextOptions struct {
	traceIDs []string
}

// TraceContextOption is an option to use with TraceContext.
type TraceContextOption interface {
	apply(o *traceContextOptions)
}

type funcTraceContextOption func(*traceContextOptions)

func (ftco funcTraceContextOption) apply(tco *traceContextOptions) {
	ftco(tco)
}

// WithTraceIDs defines a set of trace IDs to select from when generating trace
// contexts.
func WithTraceIDs(traceIDs []string) TraceContextOption {
	return funcTraceContextOption(func(o *traceContextOptions) {
		o.traceIDs = make([]string, len(traceIDs))
		copy(o.traceIDs, traceIDs)
	})
}

// TraceContext generates a TraceContext. Options may be provided to customize
// the returned object.
func TraceContext(options ...TraceContextOption) *flowpb.TraceContext {
	opts := traceContextOptions{}
	for _, opt := range options {
		opt.apply(&opts)
	}
	traceID := fakeTraceID()
	if len(opts.traceIDs) != 0 {
		traceID = opts.traceIDs[rand.Intn(len(opts.traceIDs))]
	}
	return &flowpb.TraceContext{
		Parent: &flowpb.TraceParent{
			TraceId: traceID,
		},
	}
}

// fakeTraceID generates a fake trace ID. See the W3C Trace Context
// specification for details:
// https://www.w3.org/TR/trace-context/#trace-id
func fakeTraceID() string {
	var tid [traceIDLen]byte
	for !isValidTraceID(tid[:]) {
		_, _ = rand.Read(tid[:])
	}
	return hex.EncodeToString(tid[:])
}

func isValidTraceID(tid []byte) bool {
	if len(tid) != traceIDLen {
		return false
	}
	// all bytes as zero is considered an invalid value
	for _, b := range tid {
		if b != 0 {
			return true
		}
	}
	return false
}
