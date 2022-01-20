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
	"math/rand"

	flowpb "github.com/cilium/cilium/api/v1/flow"
	"github.com/cilium/fake"
)

type endpointOptions struct {
	namespace string
	podName   string
	labels    []string
}

// EndpointOption is an option to use with Endpoint.
type EndpointOption interface {
	apply(o *endpointOptions)
}

type funcEndpointOption func(*endpointOptions)

func (feo funcEndpointOption) apply(eo *endpointOptions) {
	feo(eo)
}

// WithEndpointNamespace sets the namespace of the endpoint.
func WithEndpointNamespace(ns string) EndpointOption {
	return funcEndpointOption(func(o *endpointOptions) {
		o.namespace = ns
	})
}

// WithEndpointPodName sets the pod name of the endpoint.
func WithEndpointPodName(name string) EndpointOption {
	return funcEndpointOption(func(o *endpointOptions) {
		o.podName = name
	})
}

// WithEndpointLabels sets the labels of the endpoint.
func WithEndpointLabels(labels []string) EndpointOption {
	return funcEndpointOption(func(o *endpointOptions) {
		o.labels = make([]string, len(labels))
		copy(o.labels, labels)
	})
}

// Endpoint generates a random Endpoint. Options may be provided to customize
// the endpoint to return.
func Endpoint(options ...EndpointOption) *flowpb.Endpoint {
	opts := endpointOptions{
		namespace: fake.K8sNamespace(),
		podName:   fake.K8sPodName(),
		labels:    fake.K8sLabels(),
	}
	for _, opt := range options {
		opt.apply(&opts)
	}

	return &flowpb.Endpoint{
		ID:        rand.Uint32(),
		Identity:  rand.Uint32(),
		Namespace: opts.namespace,
		PodName:   opts.podName,
		Labels:    opts.labels,
	}
}
