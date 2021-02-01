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
	"math/rand"
)

type portOptions struct {
	min, max int
}

// PortOption is an option to use with Port.
type PortOption interface {
	apply(o *portOptions)
}

type funcPortOption func(*portOptions)

func (fpo funcPortOption) apply(po *portOptions) {
	fpo(po)
}

// WithPortSystem specifies that the generated port must in the system range.
// If zero is true, the port range is [0,1023], otherwise it's [1,1023].
// This option is incompatible with WithPortUser and WithPortDynamic.
func WithPortSystem(zero bool) PortOption {
	return funcPortOption(func(o *portOptions) {
		o.min = 1
		if zero {
			o.min = 0
		}
		o.max = 1023
	})
}

// WithPortUser specifies that the generated port must be in the user range,
// i.e. [1024, 49151].
// This option is incompatible with WithPortSystem and WithPortDynamic.
func WithPortUser() PortOption {
	return funcPortOption(func(o *portOptions) {
		o.min = 1024
		o.max = 49_151
	})
}

// WithPortDynamic specifies that the generated port must be in the dynamic
// range, i.e. [49152,65535].
// This option is incompatible with WithPortSystem and WithPortUser.
func WithPortDynamic() PortOption {
	return funcPortOption(func(o *portOptions) {
		o.min = 49_152
		o.max = 65_535
	})
}

// Port generates a random port number between 1 and 65535 or in the range
// specified by the given option.
func Port(options ...PortOption) uint32 {
	opts := portOptions{
		min: 1,
		max: 65_535,
	}
	for _, opt := range options {
		opt.apply(&opts)
	}
	return uint32(rand.Intn(opts.max-opts.min) + opts.min)
}
