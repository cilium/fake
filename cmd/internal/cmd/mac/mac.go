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
