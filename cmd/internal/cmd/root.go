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

package cmd

import (
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/cilium/fake/cmd/internal/cmd/flow"
	"github.com/cilium/fake/cmd/internal/cmd/ip"
	"github.com/cilium/fake/cmd/internal/cmd/mac"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var opts struct {
	pprofCPU bool
	pprofMem bool
}

// New creates a new root command.
func New() *cobra.Command {
	var pprofCPUFile *os.File
	rootCmd := &cobra.Command{
		Use:           "fake",
		Short:         "Generate fake network data",
		SilenceErrors: true, //this is being handled in main
		SilenceUsage:  true,
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			if opts.pprofCPU {
				var err error
				pprofCPUFile, err = os.Create("cpu.prof")
				if err != nil {
					return err
				}
				if err := pprof.StartCPUProfile(pprofCPUFile); err != nil {
					return err
				}
			}
			return nil
		},
		PersistentPostRunE: func(_ *cobra.Command, _ []string) error {
			if pprofCPUFile != nil {
				pprof.StopCPUProfile()
				pprofCPUFile.Close()
			}
			if opts.pprofMem {
				runtime.GC() // get up-to-date statistics
				f, err := os.Create("mem.prof")
				if err != nil {
					return err
				}
				if err := pprof.WriteHeapProfile(f); err != nil {
					return err
				}
				f.Close()
			}
			return nil
		},
	}

	rootCmd.AddCommand(
		flow.New(),
		mac.New(),
		ip.New(),
	)

	flags := pflag.NewFlagSet("", pflag.ContinueOnError)
	flags.BoolVar(&opts.pprofCPU, "pprof-cpu", false, "Write a pprof CPU profile to `cpu.prof`")
	flags.BoolVar(&opts.pprofMem, "pprof-mem", false, "Write a pprof memory profile to `mem.prof`")
	rootCmd.PersistentFlags().AddFlagSet(flags)

	return rootCmd
}

// Execute creates the root command and executes it.
func Execute() error {
	return New().Execute()
}
