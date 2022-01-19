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

package flow

import (
	flowpb "github.com/cilium/cilium/api/v1/flow"
	"github.com/cilium/fake"
)

type serviceOptions struct {
	namespace string
	name      string
}

// ServiceOption is an option to use with Service.
type ServiceOption interface {
	apply(o *serviceOptions)
}

type funcServiceOption func(*serviceOptions)

func (fso funcServiceOption) apply(so *serviceOptions) {
	fso(so)
}

// WithServiceNamespace sets the namespace of the service.
func WithServiceNamespace(ns string) ServiceOption {
	return funcServiceOption(func(o *serviceOptions) {
		o.namespace = ns
	})
}

// WithServiceName sets the name of the service.
func WithServiceName(name string) ServiceOption {
	return funcServiceOption(func(o *serviceOptions) {
		o.name = name
	})
}

// Service generates a random Service. Options may be provided to customize the
// service to return.
func Service(options ...ServiceOption) *flowpb.Service {
	opts := serviceOptions{
		namespace: fake.Namespace(),
		name:      fake.Name(),
	}
	for _, opt := range options {
		opt.apply(&opts)
	}

	return &flowpb.Service{
		Namespace: opts.namespace,
		Name:      opts.name,
	}
}
