// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package fake

// K8sLabels implements the Faker interface for faker.
func (f *faker) K8sLabels() []string {
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

// K8sNamespace implements the Faker interface for faker.
func (f *faker) K8sNamespace() string {
	if f.IntN(2) == 0 {
		return namespaces[f.IntN(len(namespaces))]
	}
	return join3(f.App(), "-", f.DeploymentTier())
}

// K8sNodeName implements the Faker interface for faker.
func (f *faker) K8sNodeName() string {
	return join3(f.Adjective(), "-", f.Noun())
}

// K8sPodName implements the Faker interface for faker.
func (f *faker) K8sPodName() string {
	return join3(f.App(), "-", f.AlphaNum(5))
}
