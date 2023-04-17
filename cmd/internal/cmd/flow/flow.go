// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"time"

	flowpb "github.com/cilium/cilium/api/v1/flow"
	observerpb "github.com/cilium/cilium/api/v1/observer"
	"github.com/cilium/fake"
	fakeflow "github.com/cilium/fake/flow"
	"github.com/cilium/hubble/pkg/printer"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	day  = 24 * time.Hour
	week = 7 * day
)

var opts struct {
	output               string
	count                int
	nodesCount           int
	ipVersion            flowpb.IPVersion
	since                time.Duration
	until                time.Duration
	sourceCIDR, destCIDR string
}

// New creates a new flow command.
func New() *cobra.Command {
	flowsCmd := &cobra.Command{
		Use:   "flow",
		Short: "Generate random Hubble flows",
		RunE: func(cmd *cobra.Command, _ []string) error {
			opts.ipVersion = flowpb.IPVersion_IPv4
			srcIP, _, err := net.ParseCIDR(opts.sourceCIDR)
			if err != nil {
				return err
			}
			dstIP, _, err := net.ParseCIDR(opts.destCIDR)
			if err != nil {
				return err
			}
			if ((srcIP.To4() == nil) && (dstIP.To4() != nil)) || ((srcIP.To4() != nil) && (dstIP.To4() == nil)) {
				return errors.New("source and destination CIDR IP version mismatch")
			}
			if srcIP.To4() == nil {
				opts.ipVersion = flowpb.IPVersion_IPv6
			}

			var printerOpts []printer.Option
			switch opts.output {
			case "compact":
				printerOpts = append(printerOpts, printer.Compact())
			case "dict":
				printerOpts = append(printerOpts, printer.Dict())
			case "jsonpb", "json", "JSON":
				printerOpts = append(printerOpts, printer.JSONPB())
			case "":
				fallthrough
			case "tab", "table":
				printerOpts = append(printerOpts, printer.Tab())
			default:
				return fmt.Errorf("invalid output format: %s", opts.output)
			}
			p := printer.New(printerOpts...)
			defer p.Close()
			return runFlows(p)
		},
	}

	flags := pflag.NewFlagSet("", pflag.ContinueOnError)
	flags.IntVar(&opts.count, "count", 1, "Number of flows to generate.")
	flags.IntVar(&opts.nodesCount, "nodes-count", 10, "Defines the number of nodes.")
	flags.DurationVar(&opts.since, "since", week, "Defines the time of the oldest flow.")
	flags.DurationVar(&opts.until, "until", 0, "Defines the time of the newest flow. A value of 0 defaults to the current time.")
	flags.StringVar(&opts.sourceCIDR, "cidr-source", "10.0.0.0/8", "Network for the source IPs.")
	flags.StringVar(&opts.destCIDR, "cidr-dest", "10.0.0.0/8", "Network for the destination IPs.")
	flags.StringVarP(&opts.output, "output", "o", "",
		`Specify the output format, one of:
 compact:  Compact output
 dict:     Each flow is shown as KEY:VALUE pair
 jsonpb:   Output each GetFlowResponse according to proto3's JSON mapping
 json:     Alias to jsonpb
 table:    Tab-aligned columns`)
	flowsCmd.Flags().AddFlagSet(flags)

	return flowsCmd
}

type node struct {
	name string
	ip   *flowpb.IP
}

func runFlows(p *printer.Printer) error {
	now := time.Now().UTC()
	until := now.Add(-opts.until)
	since := now.Add(-opts.since)

	if opts.count < 1 {
		return errors.New("--count must be at least 1")
	}
	if !until.After(since) {
		return errors.New("--since must come before --until")
	}

	timeRange := until.Sub(since)
	winRange := timeRange / time.Duration(opts.count)
	if int64(timeRange) < int64(opts.count) {
		return errors.New("time range it too short for the requested count")
	}

	nodesIPs := make([]node, opts.nodesCount)
	for i := 0; i < len(nodesIPs); i++ {
		nodesIPs[i] = node{
			name: fake.K8sNodeName(),
			ip: &flowpb.IP{
				Source:      fake.IP(fake.WithIPCIDR(opts.sourceCIDR)),
				Destination: fake.IP(fake.WithIPCIDR(opts.destCIDR)),
				IpVersion:   opts.ipVersion,
			},
		}
	}

	for i := 0; i < opts.count; i++ {
		// used to get a random node name
		idx := rand.Intn(len(nodesIPs))
		// add a small amount of jitter to the timestamps so they don't have perfect time gaps
		jitter := rand.Int63n(int64(winRange))
		t := since.Add(time.Duration(i)*winRange + time.Duration(jitter))
		err := p.WriteGetFlowsResponse(&observerpb.GetFlowsResponse{
			NodeName: nodesIPs[idx].name,
			Time:     timestamppb.New(t),
			ResponseTypes: &observerpb.GetFlowsResponse_Flow{
				Flow: fakeflow.New(
					fakeflow.WithFlowTime(t),
					fakeflow.WithFlowNodeName(nodesIPs[idx].name),
					fakeflow.WithFlowIP(nodesIPs[idx].ip),
				),
			},
		})
		if err != nil {
			return err

		}
	}
	return nil
}
