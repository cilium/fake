// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow_test

import (
	"testing"

	"github.com/cilium/fake/flow"
)

func BenchmarkLegacyFlowAPI(b *testing.B) {
	b.Run("Sequential", func(b *testing.B) {
		for range b.N {
			_ = flow.New()
		}
	})

	b.Run("Parallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = flow.New()
			}
		})
	})
}

func BenchmarkInstanceFlowAPI(b *testing.B) {
	b.Run("Sequential", func(b *testing.B) {
		f := flow.NewFaker()
		for range b.N {
			_ = f.NewFlow()
		}
	})

	b.Run("Parallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			f := flow.NewFaker()
			for pb.Next() {
				_ = f.NewFlow()
			}
		})
	})
}
