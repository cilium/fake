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
