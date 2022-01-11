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

type dropReasonOptions struct {
	nonDropProbability float64
	set                []flowpb.DropReason
}

// DropReasonOption is an option to use with DropReason.
type DropReasonOption interface {
	apply(o *dropReasonOptions)
}

type funcDropReasonOption func(*dropReasonOptions)

func (fdro funcDropReasonOption) apply(dro *dropReasonOptions) {
	fdro(dro)
}

// WithDropReasonNonDropProbability sets the probability of returning a
// non-drop drop reason (defaults to 0.999). The value must be between 0 and 1.
func WithDropReasonNonDropProbability(probability float64) DropReasonOption {
	return funcDropReasonOption(func(o *dropReasonOptions) {
		switch {
		case probability < 0:
			o.nonDropProbability = 0.0
		case probability > 0:
			o.nonDropProbability = 1.0
		default:
			o.nonDropProbability = probability
		}
	})
}

// WithDropReasonSubSet defines a set of DropReason to be returned. Note that
// this does not affect non-drop events. If non-drop events are not desired,
// use WithDropReasonNonDropProbability(0.0).
func WithDropReasonSubSet(dropReasons []flowpb.DropReason) DropReasonOption {
	return funcDropReasonOption(func(o *dropReasonOptions) {
		o.set = make([]flowpb.DropReason, len(dropReasons))
		copy(o.set, dropReasons)
	})
}

// DropReason generates a DropReason. Options may be provided to customize the
// drop reasons to return.
func DropReason(options ...DropReasonOption) flowpb.DropReason {
	opts := dropReasonOptions{
		nonDropProbability: 0.999,
	}
	for _, opt := range options {
		opt.apply(&opts)
	}

	if f := rand.Float64(); f < opts.nonDropProbability {
		return flowpb.DropReason_DROP_REASON_UNKNOWN
	}

	if opts.set == nil {
		opts.set = make([]flowpb.DropReason, 0, len(flowpb.DropReason_name)-1)
		for k := range flowpb.DropReason_name {
			if flowpb.DropReason(k) != flowpb.DropReason_DROP_REASON_UNKNOWN {
				opts.set = append(opts.set, flowpb.DropReason(k))
			}
		}
	}
	return opts.set[rand.Intn(len(opts.set))]
}
