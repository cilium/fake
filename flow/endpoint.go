// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

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
	workloads map[string]string
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

// WithEndpointWorkloads sets the endpoint workloads. The map's keys are the
// workload names, and values are the workload kind.
func WithEndpointWorkloads(workloads map[string]string) EndpointOption {
	return funcEndpointOption(func(o *endpointOptions) {
		o.workloads = make(map[string]string, len(workloads))
		for name, kind := range workloads {
			o.workloads[name] = kind
		}
	})
}

// Endpoint generates a random Endpoint. Options may be provided to customize
// the endpoint to return.
func Endpoint(options ...EndpointOption) *flowpb.Endpoint {
	opts := endpointOptions{
		namespace: fake.K8sNamespace(),
		podName:   fake.K8sPodName(),
		labels:    fake.K8sLabels(),
		workloads: fakeWorkloads(),
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
		Workloads: workloadsFromMap(opts.workloads),
	}
}

// see https://github.com/cilium/cilium/blob/d6483e1d26052b7f210a746ff85e919b76c9430f/pkg/k8s/utils/workload.go#L32
var workloadKinds []string = []string{
	"CronJob",
	"DaemonSet",
	"Deployment",
	"DeploymentConfig",
	"ReplicaSet",
}

func fakeWorkloads() map[string]string {
	workloads := map[string]string{
		fake.App(): workloadKinds[rand.Intn(len(workloadKinds))],
	}
	if rand.Intn(10) == 0 { // 10% chance of having more than one workload.
		workloads[fake.App()] = workloadKinds[rand.Intn(len(workloadKinds))]
	}
	if rand.Intn(100) == 0 { // 1% chance of having more than two workloads.
		workloads[fake.App()] = workloadKinds[rand.Intn(len(workloadKinds))]
	}
	return workloads
}

func workloadsFromMap(m map[string]string) []*flowpb.Workload {
	workloads := make([]*flowpb.Workload, 0, len(m))
	for name, kind := range m {
		workloads = append(workloads, &flowpb.Workload{
			Name: name,
			Kind: kind,
		})
	}
	return workloads
}
