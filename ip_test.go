// Copyright (C) Isovalent, Inc. - All Rights Reserved.
//
// NOTICE: All information contained herein is, and remains the property of
// Isovalent Inc and its suppliers, if any. The intellectual and technical
// concepts contained herein are proprietary to Isovalent Inc and its suppliers
// and may be covered by U.S. and Foreign Patents, patents in process, and are
// protected by trade secret or copyright law.  Dissemination of this information
// or reproduction of this material is strictly forbidden unless prior written
// permission is obtained from Isovalent Inc.

package fake

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IP(t *testing.T) {
	tests := []struct {
		name string
		opts []IPOption
	}{
		{
			name: "a random IPv4 or IPv6",
		}, {
			name: "a random IPv4 or IPv6 (explicit)",
			opts: []IPOption{
				WithIPv4(),
				WithIPv6(),
			},
		}, {
			name: "a random IPv4",
			opts: []IPOption{
				WithIPv4(),
			},
		}, {
			name: "a random IPv6",
			opts: []IPOption{
				WithIPv6(),
			},
		}, {
			name: "an IP within 192.0.2.1/24",
			opts: []IPOption{
				WithIPCIDR("192.0.2.1/24"),
			},
		}, {
			name: "an IP within 10.0.0.0/7",
			opts: []IPOption{
				WithIPCIDR("10.0.0.0/7"),
			},
		}, {
			name: "an IP within 2001:db8:a0b:12f0::1/32",
			opts: []IPOption{
				WithIPCIDR("2001:db8:a0b:12f0::1/32"),
			},
		}, {
			name: "an IP within fc00::/7",
			opts: []IPOption{
				WithIPCIDR("fc00::/7"),
			},
		}, {
			name: "an IPv4 address mapped within ::ffff:0:0/96",
			opts: []IPOption{
				WithIPCIDR("::ffff:0:0/96"),
			},
		}, {
			name: "an IPv4 address translated within ::ffff:0:0:0/96",
			opts: []IPOption{
				WithIPCIDR("::ffff:0:0:0/96"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := ipOptions{}
			for _, opt := range tt.opts {
				opt.apply(&opts)
			}

			ip := net.ParseIP(IP(tt.opts...))
			assert.NotNil(t, ip)
			if opts.v4 && !opts.v6 {
				assert.NotNil(t, ip.To4())
			}
			if !opts.v4 && opts.v6 {
				assert.Nil(t, ip.To4())
			}
			if opts.network != nil {
				assert.True(t, opts.network.Contains(ip))
			}
		})
	}
}
