// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	"math/rand"

	flowpb "github.com/cilium/cilium/api/v1/flow"
	"github.com/cilium/fake"
)

type ipOptions struct {
	sourceNATProbability float64
}

// IPOption is an option to use with IP.
type IPOption interface {
	apply(o *ipOptions)
}

type funcipOption func(*ipOptions)

func (fso funcipOption) apply(io *ipOptions) {
	fso(io)
}

// WithSourceNATProbability defines the probability for a IP to contain a
// SourceXlated addresss. The value must be between 0 and 1.
func WithSourceNATProbability(probability float64) IPOption {
	return funcipOption(func(o *ipOptions) {
		switch {
		case probability < 0:
			o.sourceNATProbability = 0.0
		case probability > 1:
			o.sourceNATProbability = 1.0
		default:
			o.sourceNATProbability = probability
		}
	})
}

func IP(options ...IPOption) *flowpb.IP {
	opts := ipOptions{}
	for _, opt := range options {
		opt.apply(&opts)
	}
	if p := rand.Float64(); p < opts.sourceNATProbability {
		return &flowpb.IP{
			Source:       fake.IP(fake.WithIPCIDR("10.0.0.0/8")),
			SourceXlated: fake.IP(fake.WithIPCIDR("172.16.0.0/12")),
			Destination:  fake.IP(fake.WithIPCIDR("172.16.0.0/12")),
			IpVersion:    flowpb.IPVersion_IPv4,
		}
	} else {
		return &flowpb.IP{
			Source:      fake.IP(fake.WithIPCIDR("10.0.0.0/8")),
			Destination: fake.IP(fake.WithIPCIDR("10.0.0.0/8")),
			IpVersion:   flowpb.IPVersion_IPv4,
		}
	}
}
