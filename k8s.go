// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package fake

// K8sLabels generates a random set of Kubernetes labels.
func (f *Faker) K8sLabels() []string {
	// XXX: figure out if this is called often, we could init the slice with a
	// len(labels) / 2 capacity.
	var l []string
	for _, name := range labels {
		if f.IntN(2) == 0 { // 50% chance of picking up this label
			l = append(l, join3(name, "=", f.App()))
		}
	}
	return l
}

// K8sNamespace generates a random Kubernetes namespace name.
func (f *Faker) K8sNamespace() string {
	if f.IntN(2) == 0 {
		return namespaces[f.IntN(len(namespaces))]
	}
	return join3(f.App(), "-", f.DeploymentTier())
}

// K8sNodeName generates a random Kubernetes node name.
func (f *Faker) K8sNodeName() string {
	return join3(f.Adjective(), "-", f.Noun())
}

// K8sPodName generates a random Kubernetes pod name.
func (f *Faker) K8sPodName() string {
	return join3(f.App(), "-", f.AlphaNum(5))
}
