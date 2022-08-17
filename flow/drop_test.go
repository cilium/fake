// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	"math"
	"testing"

	flowpb "github.com/cilium/cilium/api/v1/flow"
	"github.com/stretchr/testify/assert"
)

func Test_DropReason(t *testing.T) {
	tests := []struct {
		name string
		opts []DropReasonOption
	}{
		{
			name: "a non-drop DropReason",
			opts: []DropReasonOption{
				WithDropReasonNonDropProbability(1.0),
			},
		}, {
			name: "a random drop DropReason",
			opts: []DropReasonOption{
				WithDropReasonNonDropProbability(0.0),
			},
		}, {
			name: "a random DropReason from a subset",
			opts: []DropReasonOption{
				WithDropReasonNonDropProbability(0.0),
				WithDropReasonSubSet([]flowpb.DropReason{
					flowpb.DropReason_UNKNOWN_ICMPV6_CODE,
					flowpb.DropReason_UNKNOWN_ICMPV6_TYPE,
					flowpb.DropReason_ERROR_RETRIEVING_TUNNEL_KEY,
					flowpb.DropReason_INVALID_GENEVE_OPTION,
				}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := dropReasonOptions{}
			for _, opt := range tt.opts {
				opt.apply(&opts)
			}

			dropReason := DropReason(tt.opts...)
			if _, ok := flowpb.DropReason_name[int32(dropReason)]; !ok {
				t.Fatalf("drop reason does not exist: %v", dropReason)
			}

			if len(opts.set) > 0 {
				found := false
				for _, reason := range opts.set {
					if dropReason == reason {
						found = true
						break
					}
				}
				assert.True(t, found)
			}

			nonDrop, count := 0, 10000
			for i := 0; i < count; i++ {
				if d := DropReason(tt.opts...); d == flowpb.DropReason_DROP_REASON_UNKNOWN {
					nonDrop++
				}
			}
			ratio := float64(nonDrop) / float64(count)
			threshold := 0.03 // that's a good enough threshold for our purpose
			assert.Less(t, math.Abs(ratio-opts.nonDropProbability), threshold)
		})
	}
}
