// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	"math/rand"
	"time"

	flowpb "github.com/cilium/cilium/api/v1/flow"
	"github.com/cilium/fake"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type flowOptions struct {
	time                    time.Time
	verdict                 flowpb.Verdict
	authType                flowpb.AuthType
	dropReason              flowpb.DropReason
	clusterName             string
	nodeName                string
	sourceNames, destNames  []string
	ip                      *flowpb.IP
	ethernet                *flowpb.Ethernet
	l4                      *flowpb.Layer4
	epSource, epDest        *flowpb.Endpoint
	typ                     flowpb.FlowType
	traceContextProbability float64
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

// WithFlowAuthType sets the authentication type field of a flow.
func WithFlowAuthType(t flowpb.AuthType) Option {
	return funcFlowOption(func(o *flowOptions) {
		o.authType = t
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

// WithFlowClusterName sets the cluster name of a flow.
func WithFlowClusterName(name string) Option {
	return funcFlowOption(func(o *flowOptions) {
		o.clusterName = name
	})
}

// WithFlowSourceNames sets the source names field of a flow.
func WithFlowSourceNames(names []string) Option {
	return funcFlowOption(func(o *flowOptions) {
		o.sourceNames = names
	})
}

// WithFlowDestinationNames sets the destination names field of a flow.
func WithFlowDestinationNames(names []string) Option {
	return funcFlowOption(func(o *flowOptions) {
		o.destNames = names
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

// WithTraceContextProbability defines the probability for a flow to contain a
// trace context. The value must be between 0 and 1. Defaults to 0.1.
func WithTraceContextProbability(probability float64) Option {
	return funcFlowOption(func(o *flowOptions) {
		switch {
		case probability < 0:
			o.traceContextProbability = 0.0
		case probability > 1:
			o.traceContextProbability = 1.0
		default:
			o.traceContextProbability = probability
		}
	})
}

// New generates a random flow. Options may be provided to customize the flow.
func New(options ...Option) *flowpb.Flow {
	opts := flowOptions{
		time:                    time.Now().UTC(),
		verdict:                 Verdict(),
		authType:                AuthType(),
		typ:                     flowpb.FlowType_L3_L4,
		nodeName:                fake.K8sNodeName(),
		clusterName:             "default",
		sourceNames:             fake.Names(5),
		destNames:               fake.Names(5),
		epSource:                Endpoint(),
		epDest:                  Endpoint(),
		traceContextProbability: 0.1,
	}
	for _, opt := range options {
		opt.apply(&opts)
	}

	if opts.epSource.ClusterName == "" {
		opts.epSource.ClusterName = opts.clusterName
	}
	if opts.epDest.ClusterName == "" {
		opts.epDest.ClusterName = opts.clusterName
	}

	if opts.typ == flowpb.FlowType_L3_L4 {
		if opts.ip == nil {
			opts.ip = &flowpb.IP{
				Source:       fake.IP(fake.WithIPCIDR("10.0.0.0/8")),
				SourceXlated: "", // TODO
				Destination:  fake.IP(fake.WithIPCIDR("10.0.0.0/8")),
				IpVersion:    flowpb.IPVersion_IPv4,
			}
		}
		if opts.l4 == nil {
			var l4Option Layer4Option
			switch rand.Intn(5) {
			case 0:
				l4Option = WithLayer4TCP()
			case 1:
				l4Option = WithLayer4UDP()
			case 2:
				l4Option = WithLayer4SCTP()
			case 3:
				l4Option = WithLayer4ICMPv4()
			case 4:
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

	// These fields are marked as omitempty, so we want to nil them for empty slices
	if len(opts.sourceNames) == 0 {
		opts.sourceNames = nil
	}
	if len(opts.destNames) == 0 {
		opts.destNames = nil
	}

	// TODO: in theory, we should only add a trace context to L7 flows with the
	// given probability. However, as the library doesn't support L7 flows yet,
	// add trace context unconditionally.
	var tc *flowpb.TraceContext
	if p := rand.Float64(); p < opts.traceContextProbability {
		tc = TraceContext()
	}

	return &flowpb.Flow{
		Time:     timestamppb.New(opts.time),
		Uuid:     uuid.NewString(),
		Verdict:  opts.verdict,
		AuthType: opts.authType,
		// NOTE: don't populate DropReason as it is deprecated.
		Ethernet:         opts.ethernet,
		IP:               opts.ip,
		L4:               opts.l4,
		Source:           opts.epSource,
		Destination:      opts.epDest,
		Type:             opts.typ,
		NodeName:         opts.clusterName + "/" + opts.nodeName,
		NodeLabels:       []string{}, //TODO
		SourceNames:      opts.sourceNames,
		DestinationNames: opts.destNames,
		// TODO: L7
		// NOTE: don't populate Reply as it is deprecated.
		EventType:             EventType(),
		SourceService:         Service(),
		DestinationService:    Service(),
		TrafficDirection:      TrafficDirection(),
		PolicyMatchType:       uint32(rand.Intn(5)), //nolint:gosec
		TraceObservationPoint: TraceObservationPoint(),
		TraceReason:           TraceReason(),
		DropReasonDesc:        opts.dropReason,
		IsReply:               IsReply(),
		TraceContext:          tc,
		SockXlatePoint:        flowpb.SocketTranslationPoint(rand.Intn(len(flowpb.SocketTranslationPoint_name))), //nolint:gosec
		SocketCookie:          rand.Uint64(),
		CgroupId:              rand.Uint64(),
		// NOTE: don't populate Summary as it is deprecated.
		EgressAllowedBy:  []*flowpb.Policy{}, //TODO
		IngressAllowedBy: []*flowpb.Policy{}, //TODO
		EgressDeniedBy:   []*flowpb.Policy{}, //TODO
		IngressDeniedBy:  []*flowpb.Policy{}, //TODO
	}
}
