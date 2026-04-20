// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Package main is the CLI of the fake library.
package main

import (
	"fmt"
	"os"

	"github.com/cilium/fake/cmd/root"
)

func main() {
	if err := execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func execute() error {
	return root.New().Execute()
}
