// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Package mac generate random MAC addresses.
package mac

import (
	"fmt"

	"github.com/cilium/fake"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var opts struct {
	count int
}

// New creates a new mac command.
func New() *cobra.Command {
	macCmd := &cobra.Command{
		Use:   "mac",
		Short: "Generate random MAC addresses",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runMACs(cmd)
		},
	}

	flags := pflag.NewFlagSet("", pflag.ContinueOnError)
	flags.IntVar(&opts.count, "count", 1, "Number of MAC address to generate.")
	macCmd.Flags().AddFlagSet(flags)

	return macCmd
}

func runMACs(cmd *cobra.Command) error {
	for range opts.count {
		fmt.Fprintln(cmd.OutOrStdout(), fake.MAC())
	}
	return nil
}
