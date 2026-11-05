// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	"bytes"
	"strconv"
	"testing"
	"time"

	flowpb "github.com/cilium/cilium/api/v1/flow"
	"github.com/cilium/cilium/hubble/pkg/printer"
)

func TestRunFlowsRejectsInvalidNodeCounts(t *testing.T) {
	t.Cleanup(resetOpts)

	for _, nodesCount := range []int{0, -1} {
		t.Run("nodes-count="+strconv.Itoa(nodesCount), func(t *testing.T) {
			var out bytes.Buffer
			p := printer.New(printer.JSONPB(), printer.Writer(&out))
			t.Cleanup(func() { p.Close() })

			opts.count = 1
			opts.nodesCount = nodesCount
			opts.since = time.Hour
			opts.until = 0
			opts.sourceCIDR = "10.0.0.0/8"
			opts.destCIDR = "10.0.0.0/8"
			opts.ipVersion = 0

			err := runFlows(p)
			if err == nil || err.Error() != "--nodes-count must be at least 1" {
				t.Fatalf("expected invalid nodes-count error, got %v", err)
			}
			if out.Len() != 0 {
				t.Fatalf("expected no output, got %q", out.String())
			}
		})
	}
}

func TestRunFlowsGeneratesOutput(t *testing.T) {
	t.Cleanup(resetOpts)

	var out bytes.Buffer
	p := printer.New(printer.JSONPB(), printer.Writer(&out))
	t.Cleanup(func() { p.Close() })

	opts.count = 1
	opts.nodesCount = 1
	opts.since = time.Hour
	opts.until = 0
	opts.sourceCIDR = "10.0.0.0/8"
	opts.destCIDR = "10.0.0.0/8"

	err := runFlows(p)
	if err != nil {
		t.Fatalf("runFlows returned error: %v", err)
	}
	if out.Len() == 0 {
		t.Fatal("expected generated flow output")
	}
}

func resetOpts() {
	opts = struct {
		output               string
		count                int
		nodesCount           int
		ipVersion            flowpb.IPVersion
		since                time.Duration
		until                time.Duration
		sourceCIDR, destCIDR string
	}{}
}
