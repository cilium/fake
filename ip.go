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

package fake

import (
	"fmt"
	"math/rand"
	"net"
)

type ipOptions struct {
	v4, v6  bool
	network *net.IPNet
}

// IPOption is an option to use with IP.
type IPOption interface {
	apply(o *ipOptions)
}

type funcIPOption func(*ipOptions)

func (fio funcIPOption) apply(io *ipOptions) {
	fio(io)
}

// WithIPv6 specifies that the generated IP address must be IPv6.
// This option is incompatible with WithIPv4 and WithIPCIDR.
func WithIPv6() IPOption {
	return funcIPOption(func(o *ipOptions) {
		o.v6 = true
	})
}

// WithIPv4 specifies that the generated IP address must be IPv4.
// This option is incompatible with WithIPv6 and WithIPCIDR.
func WithIPv4() IPOption {
	return funcIPOption(func(o *ipOptions) {
		o.v4 = true
	})
}

// WithIPCIDR defines a network for generated IPs. This option is incompatible
// with WithIPv4 and WithIPv6. It panics if the given cidr is invalid.
func WithIPCIDR(cidr string) IPOption {
	return funcIPOption(func(o *ipOptions) {
		_, network, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(fmt.Sprintf("invalid CIDR (%s): %s", cidr, err))
		}
		o.network = network
	})
}

// IP generates a random IP address. Options may be provided to specify a
// network for the address or if it should be IPv4 or IPv6.
func IP(options ...IPOption) string {
	opts := ipOptions{}
	for _, opt := range options {
		opt.apply(&opts)
	}

	var size int
	if opts.network == nil {
		switch {
		case opts.v4 == opts.v6:
			sizes := []int{net.IPv4len, net.IPv6len}
			size = sizes[rand.Intn(len(sizes))]
		case opts.v4:
			size = net.IPv4len
		case opts.v6:
			size = net.IPv6len
		}
		ip := make([]byte, size)
		_, _ = rand.Read(ip)
		return net.IP(ip).String()
	}

	size = len(opts.network.Mask)
	raw := make([]byte, size)
	_, _ = rand.Read(raw)
	ip := opts.network.IP
	for i, v := range raw {
		ip[i] = ip[i] + (v &^ opts.network.Mask[i])
	}
	return ip.String()
}
