// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

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
