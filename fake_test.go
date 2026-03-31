// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package fake_test

import (
	"testing"

	"github.com/cilium/fake"
)

func BenchmarkLegacyAPI(b *testing.B) {
	b.Run("Sequential", func(b *testing.B) {
		for range b.N {
			_ = fake.IP()
			_ = fake.MAC()
			_ = fake.Port()
			_ = fake.K8sNamespace()
		}
	})

	b.Run("Parallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = fake.IP()
				_ = fake.MAC()
				_ = fake.Port()
				_ = fake.K8sNamespace()
			}
		})
	})
}

func BenchmarkInstanceAPI(b *testing.B) {
	b.Run("Sequential", func(b *testing.B) {
		f := fake.New()
		for range b.N {
			_ = f.IP()
			_ = f.MAC()
			_ = f.Port()
			_ = f.K8sNamespace()
		}
	})

	b.Run("Parallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			f := fake.New()
			for pb.Next() {
				_ = f.IP()
				_ = f.MAC()
				_ = f.Port()
				_ = f.K8sNamespace()
			}
		})
	})
}
