// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package fake

import "sync"

// globalFaker is used by the legacy API of package level functions.
var (
	globalFaker = New()
	globalMu    sync.Mutex
)

// Legacy package-level functions

// Adjective generates a random adjective.
func Adjective() string {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFaker.Adjective()
}

// AlphaNum generates a random alphanumeric string of the given length.
func AlphaNum(length int) string {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFaker.AlphaNum(length)
}

// App generates a random software application name.
func App() string {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFaker.App()
}

// Noun generates a random noun.
func Noun() string {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFaker.Noun()
}

// Name generates a random name.
func Name() string {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFaker.Name()
}

// Names generates a random set of names. It panics if n < 0.
func Names(n int) []string {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFaker.Names(n)
}

// DeploymentTier generates a random software deployment tier such as prod,
// staging, etc.
func DeploymentTier() string {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFaker.DeploymentTier()
}

// MAC generates a random MAC address.
func MAC() string {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFaker.MAC()
}

// IP generates a random IP address. Options may be provided to specify a
// network for the address or if it should be IPv4 or IPv6.
func IP(options ...IPOption) string {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFaker.IP(options...)
}

// Port generates a random port number between 1 and 65535 or in the range
// specified by the given option.
func Port(options ...PortOption) uint32 {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFaker.Port(options...)
}

// K8sLabels generates a random set of Kubernetes labels.
func K8sLabels() []string {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFaker.K8sLabels()
}

// K8sNamespace generates a random Kubernetes namespace name.
func K8sNamespace() string {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFaker.K8sNamespace()
}

// K8sNodeName generates a random Kubernetes node name.
func K8sNodeName() string {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFaker.K8sNodeName()
}

// K8sPodName generates a random Kubernetes pod name.
func K8sPodName() string {
	globalMu.Lock()
	defer globalMu.Unlock()
	return globalFaker.K8sPodName()
}
