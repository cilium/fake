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
	flowpb "github.com/cilium/cilium/api/v1/flow"
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
		namespace: Namespace(),
		name:      Name(),
	}
	for _, opt := range options {
		opt.apply(&opts)
	}

	return &flowpb.Service{
		Namespace: opts.namespace,
		Name:      opts.name,
	}
}
