// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow_test

import (
	"runtime"
	"sync"
	"testing"

	fakeflow "github.com/cilium/fake/flow"
)

func BenchmarkFakeFlow(b *testing.B) {
	// run GOMAXPROCS goroutine, allowing for override when benchmarking while
	// still being a sane default.
	p := runtime.GOMAXPROCS(0)
	n := (b.N / p) + 1 // round up

	var wg sync.WaitGroup
	wg.Add(p)
	for range p {
		go func() {
			for _ = range n {
				_ = fakeflow.New()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
