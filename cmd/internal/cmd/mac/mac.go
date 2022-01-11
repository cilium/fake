// Copyright (C) Isovalent, Inc. - All Rights Reserved.
//
// NOTICE: All information contained herein is, and remains the property of
// Isovalent Inc and its suppliers, if any. The intellectual and technical
// concepts contained herein are proprietary to Isovalent Inc and its suppliers
// and may be covered by U.S. and Foreign Patents, patents in process, and are
// protected by trade secret or copyright law.  Dissemination of this information
// or reproduction of this material is strictly forbidden unless prior written
// permission is obtained from Isovalent Inc.

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
	for i := 0; i < opts.count; i++ {
		fmt.Fprintln(cmd.OutOrStdout(), fake.MAC())
	}
	return nil
}
