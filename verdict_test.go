// Copyright (C) Isovalent, Inc. - All Rights Reserved.
//
// NOTICE: All information contained herein is, and remains the property of
// Isovalent Inc and its suppliers, if any. The intellectual and technical
// concepts contained herein are proprietary to Isovalent Inc and its suppliers
// and may be covered by U.S. and Foreign Patents, patents in process, and are
// protected by trade secret or copyright law.  Dissemination of this information
// or reproduction of this material is strictly forbidden unless prior written
// permission is obtained from Isovalent Inc.

package fake

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
			for i := 0; i < count; i++ {
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
