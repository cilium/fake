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
	return uint32(rand.Intn(opts.max+1-opts.min) + opts.min)
}
