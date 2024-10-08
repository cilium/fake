// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	"math"
	"testing"

	flowpb "github.com/cilium/cilium/api/v1/flow"
	"github.com/stretchr/testify/assert"
)

func Test_Verdict(t *testing.T) {
	tests := []struct {
		name string
		opts []VerdictOption
		want flowpb.Verdict
	}{
		{
			name: "verdict forwarded default",
		}, {
			name: "verdict forwarded 1/1",
			opts: []VerdictOption{
				WithVerdictForwardedProbability(1.0),
			},
		}, {
			name: "verdict forwarded 0/1",
			opts: []VerdictOption{
				WithVerdictForwardedProbability(0.0),
			},
		}, {
			name: "verdict forwarded 4/5",
			opts: []VerdictOption{
				WithVerdictForwardedProbability(4.0 / 5.0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := verdictOptions{
				forwardProbability: 0.999,
			}
			for _, opt := range tt.opts {
				opt.apply(&opts)
			}

			forwarded, count := 0, 10000
			for range count {
				if v := Verdict(tt.opts...); v == flowpb.Verdict_FORWARDED {
					forwarded++
				}
			}
			ratio := float64(forwarded) / float64(count)
			threshold := 0.02 // that's a good enough threshold for our purpose
			assert.Less(t, math.Abs(ratio-opts.forwardProbability), threshold)
		})
	}
}
