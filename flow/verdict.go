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

type verdictOptions struct {
	forwardProbability float64
}

// VerdictOption is an option to use with Verdict.
type VerdictOption interface {
	apply(o *verdictOptions)
}

type funcVerdictOption func(*verdictOptions)

func (fvo funcVerdictOption) apply(vo *verdictOptions) {
	fvo(vo)
}

// WithVerdictForwardedProbability sets the probability of returning a
// forwarded verdict (defaults to 0.999). The value must be between 0 and 1.
func WithVerdictForwardedProbability(probability float64) VerdictOption {
	return funcVerdictOption(func(o *verdictOptions) {
		switch {
		case probability < 0:
			o.forwardProbability = 0.0
		case probability > 1:
			o.forwardProbability = 1.0
		default:
			o.forwardProbability = probability
		}
	})
}

// Verdict generates a FORWARDED or DROPPPED verdict randomly. The probability
// of the verdict being FORWARDED can be set using
// WithVerdictForwardedProbability.
func Verdict(options ...VerdictOption) flowpb.Verdict {
	opts := verdictOptions{
		forwardProbability: 0.999,
	}
	for _, opt := range options {
		opt.apply(&opts)
	}

	if f := rand.Float64(); f < opts.forwardProbability {
		return flowpb.Verdict_FORWARDED
	}
	//TODO: return other verdict types? With which probability?
	return flowpb.Verdict_DROPPED
}
