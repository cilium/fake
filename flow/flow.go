// Copyright (C) Isovalent, Inc. - All Rights Reserved.
//
// NOTICE: All information contained herein is, and remains the property of
// Isovalent Inc and its suppliers, if any. The intellectual and technical
// concepts contained herein are proprietary to Isovalent Inc and its suppliers
// and may be covered by U.S. and Foreign Patents, patents in process, and are
// protected by trade secret or copyright law.  Dissemination of this information
// or reproduction of this material is strictly forbidden unless prior written
// permission is obtained from Isovalent Inc.

package flow

import (
	"math/rand"
	"time"

	flowpb "github.com/cilium/cilium/api/v1/flow"
	"github.com/cilium/fake"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type flowOptions struct {
	time             time.Time
	verdict          flowpb.Verdict
	dropReason       flowpb.DropReason
	nodeName         string
	ip               *flowpb.IP
	ethernet         *flowpb.Ethernet
	l4               *flowpb.Layer4
	epSource, epDest *flowpb.Endpoint
	typ              flowpb.FlowType
}

// Option is an option to use with Flow.
type Option interface {
	apply(o *flowOptions)
}

type funcFlowOption func(*flowOptions)

func (ffo funcFlowOption) apply(fo *flowOptions) {
	ffo(fo)
}

// WithFlowTime sets the time of a flow.
func WithFlowTime(t time.Time) Option {
	return funcFlowOption(func(o *flowOptions) {
		o.time = t
	})
}

// WithFlowVerdict sets the verdict of a flow.
func WithFlowVerdict(v flowpb.Verdict) Option {
	return funcFlowOption(func(o *flowOptions) {
		o.verdict = v
	})
}

// WithFlowDropReason sets the drop reason of a flow.
func WithFlowDropReason(r flowpb.DropReason) Option {
	return funcFlowOption(func(o *flowOptions) {
		o.dropReason = r
	})
}

// WithFlowNodeName sets the node name field of a flow.
func WithFlowNodeName(name string) Option {
	return funcFlowOption(func(o *flowOptions) {
		o.nodeName = name
	})
}

// WithFlowEthernet sets the ethernet field a flow.
func WithFlowEthernet(e *flowpb.Ethernet) Option {
	return funcFlowOption(func(o *flowOptions) {
		o.ethernet = e
	})
}

// WithFlowIP sets the IP field of a flow.
func WithFlowIP(ip *flowpb.IP) Option {
	return funcFlowOption(func(o *flowOptions) {
		o.ip = ip
	})
}

// WithFlowL4 sets the L4 field of a flow.
func WithFlowL4(l4 *flowpb.Layer4) Option {
	return funcFlowOption(func(o *flowOptions) {
		o.l4 = l4
	})
}

// WithFlowEndpoints sets the source and destination endpoints of a flow.
func WithFlowEndpoints(src, dst *flowpb.Endpoint) Option {
	return funcFlowOption(func(o *flowOptions) {
		o.epSource = src
		o.epDest = dst
	})
}

// WithFlowType sets the type of the flow.
func WithFlowType(t flowpb.FlowType) Option {
	return funcFlowOption(func(o *flowOptions) {
		o.typ = t
	})
}

// New generates a random flow. Options may be provided to customize the flow.
func New(options ...Option) *flowpb.Flow {
	opts := flowOptions{
		time:     time.Now().UTC(),
		verdict:  Verdict(),
		typ:      flowpb.FlowType_L3_L4,
		nodeName: fake.NodeName(),
		epSource: Endpoint(),
		epDest:   Endpoint(),
	}
	for _, opt := range options {
		opt.apply(&opts)
	}

	if opts.typ == flowpb.FlowType_L3_L4 {
		if opts.ip == nil {
			opts.ip = &flowpb.IP{
				Source:      fake.IP(fake.WithIPCIDR("10.0.0.0/8")),
				Destination: fake.IP(fake.WithIPCIDR("10.0.0.0/8")),
				IpVersion:   flowpb.IPVersion_IPv4,
			}
		}
		if opts.l4 == nil {
			var l4Option Layer4Option
			switch rand.Intn(4) {
			case 0:
				l4Option = WithLayer4TCP()
			case 1:
				l4Option = WithLayer4UDP()
			case 2:
				l4Option = WithLayer4ICMPv4()
			case 3:
				l4Option = WithLayer4ICMPv6()
			}
			opts.l4 = Layer4(l4Option)
		}
	}

	// If an IP is defined, the Ethernet part shall be as well
	if opts.ip != nil && opts.ethernet == nil {
		opts.ethernet = &flowpb.Ethernet{
			Source:      fake.MAC(),
			Destination: fake.MAC(),
		}
	}

	return &flowpb.Flow{
		Time:    timestamppb.New(opts.time),
		Verdict: opts.verdict,
		// NOTE: don't populate DropReason as it is deprecated.
		Ethernet:         opts.ethernet,
		IP:               opts.ip,
		L4:               opts.l4,
		Source:           opts.epSource,
		Destination:      opts.epDest,
		Type:             opts.typ,
		NodeName:         opts.nodeName,
		SourceNames:      fake.Names(5),
		DestinationNames: fake.Names(5),
		// TODO: L7
		// NOTE: don't populate Reply as it is deprecated.
		EventType:             CiliumEventType(),
		SourceService:         Service(),
		DestinationService:    Service(),
		TrafficDirection:      flowpb.TrafficDirection(rand.Intn(3)),
		PolicyMatchType:       uint32(rand.Intn(5)),
		TraceObservationPoint: TraceObservationPoint(),
		DropReasonDesc:        opts.dropReason,
		IsReply:               IsReply(),
		// NOTE: don't populate Summary as it is deprecated.
	}
}
