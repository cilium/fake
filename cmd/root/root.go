// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Package root exposes the root command for the fake CLI.
package root

import (
	"github.com/cilium/fake/cmd/internal/cmd"
	"github.com/spf13/cobra"
)

// New creates a new root command.
func New() *cobra.Command {
	return cmd.New()
}
