// Copyright (C) Isovalent, Inc. - All Rights Reserved.
//
// NOTICE: All information contained herein is, and remains the property of
// Isovalent Inc and its suppliers, if any. The intellectual and technical
// concepts contained herein are proprietary to Isovalent Inc and its suppliers
// and may be covered by U.S. and Foreign Patents, patents in process, and are
// protected by trade secret or copyright law.  Dissemination of this information
// or reproduction of this material is strictly forbidden unless prior written
// permission is obtained from Isovalent Inc.

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
		namespace: fake.Namespace(),
		podName:   fake.PodName(),
		labels:    fake.Labels(),
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
