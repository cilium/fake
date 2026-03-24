// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
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
		case probability > 1:
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

// DropReason implements FlowFaker for flowfaker.
func (f *flowfaker) DropReason(options ...DropReasonOption) flowpb.DropReason {
	// FIXME: evaluating all the options here for a single drop reason returned
	// feels like we're paying a ton of overhead as soon as an option is given.
	// Maybe consider moving this to the fake constructor, with maybe override
	// possible on generation for special cases only?
	opts := dropReasonOptions{
		nonDropProbability: 0.999,
	}
	for _, opt := range options {
		opt.apply(&opts)
	}

	if r := f.Float64(); r < opts.nonDropProbability {
		return flowpb.DropReason_DROP_REASON_UNKNOWN
	}

	// FIXME: extract the static default set to be computed only once.
	if opts.set == nil {
		opts.set = make([]flowpb.DropReason, 0, len(flowpb.DropReason_name)-1)
		for k := range flowpb.DropReason_name {
			if flowpb.DropReason(k) != flowpb.DropReason_DROP_REASON_UNKNOWN {
				opts.set = append(opts.set, flowpb.DropReason(k))
			}
		}
	}
	return opts.set[f.IntN(len(opts.set))]
}
