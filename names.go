// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package fake

// Adjective implements the Faker interface for faker.
func (f *faker) Adjective() string {
	return adjectives[f.IntN(len(adjectives))]
}

// AlphaNum implements the Faker interface for faker.
func (f *faker) AlphaNum(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = alphanum[f.IntN(len(alphanum))]
	}
	return string(b)
}

// App implements the Faker interface for faker.
func (f *faker) App() string {
	return apps[f.IntN(len(apps))]
}

// Noun implements the Faker interface for faker.
func (f *faker) Noun() string {
	return nouns[f.IntN(len(nouns))]
}

// Name implements the Faker interface for faker.
func (f *faker) Name() string {
	return join3(f.Adjective(), "_", f.Noun())
}

// Names implements the Faker interface for faker.
func (f *faker) Names(n int) []string {
	n = f.IntN(n + 1)
	names := make([]string, n)
	for i := range n {
		names[i] = f.Name()
	}
	return names
}

// DeploymentTier implements the Faker interface for faker.
func (f *faker) DeploymentTier() string {
	return tiers[f.IntN(len(tiers))]
}
