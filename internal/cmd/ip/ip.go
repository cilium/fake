// Copyright (C) Isovalent, Inc. - All Rights Reserved.
//
// NOTICE: All information contained herein is, and remains the property of
// Isovalent Inc and its suppliers, if any. The intellectual and technical
// concepts contained herein are proprietary to Isovalent Inc and its suppliers
// and may be covered by U.S. and Foreign Patents, patents in process, and are
// protected by trade secret or copyright law.  Dissemination of this information
// or reproduction of this material is strictly forbidden unless prior written
// permission is obtained from Isovalent Inc.

package ip

import (
	"fmt"
	"net"

	"github.com/isovalent/fake"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var opts struct {
	count int

	v4, v6 bool
	cidr   string
}

// New creates a new ip command.
func New() *cobra.Command {
	ipsCmd := &cobra.Command{
		Use:   "ip",
		Short: "Generate random IP addresses",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runIPs(cmd)
		},
	}

	flags := pflag.NewFlagSet("", pflag.ContinueOnError)
	flags.IntVar(&opts.count, "count", 1, "Number of IP addresses to generate.")
	flags.BoolVarP(&opts.v4, "ipv4", "4", false, "Generate IPv4 addresses.")
	flags.BoolVarP(&opts.v6, "ipv6", "6", false, "Generate IPv6 addresses.")
	flags.StringVar(&opts.cidr, "cidr", "", "Generate IPs in the specified network.")
	ipsCmd.Flags().AddFlagSet(flags)

	return ipsCmd
}

func runIPs(cmd *cobra.Command) error {
	var ipOptions []fake.IPOption
	if opts.v4 {
		ipOptions = append(ipOptions, fake.WithIPv4())
	}
	if opts.v6 {
		ipOptions = append(ipOptions, fake.WithIPv6())
	}
	if opts.cidr != "" {
		if _, _, err := net.ParseCIDR(opts.cidr); err != nil {
			return err
		}
		ipOptions = append(ipOptions, fake.WithIPCIDR(opts.cidr))
	}

	for i := 0; i < opts.count; i++ {
		fmt.Fprintln(cmd.OutOrStdout(), fake.IP(ipOptions...))
	}
	return nil
}
