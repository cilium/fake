// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	"github.com/cilium/fake"
)

// FlowFaker implements Hubble flow related data generation on top of the
// networking/name generators provided by the embedded Faker.
type FlowFaker struct {
	*fake.Faker
}

// NewFaker creates a new Faker using a random seed.
func NewFaker() *FlowFaker {
	return &FlowFaker{
		Faker: fake.New(),
	}
}

// NewWithSource creates a new Faker using the given random source. Useful to
// control the faker output, e.g. for testing.
func NewWithSource(src fake.RandSourceReader) *FlowFaker {
	return &FlowFaker{
		Faker: fake.NewWithSource(src),
	}
}
