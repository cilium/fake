// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package fake

// Adjective generates a random adjective.
func (f *Faker) Adjective() string {
	return adjectives[f.IntN(len(adjectives))]
}

// AlphaNum generates a random alphanumeric string of the given length.
func (f *Faker) AlphaNum(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = alphanum[f.IntN(len(alphanum))]
	}
	return string(b)
}

// App generates a random software application name.
func (f *Faker) App() string {
	return apps[f.IntN(len(apps))]
}

// Noun generates a random noun.
func (f *Faker) Noun() string {
	return nouns[f.IntN(len(nouns))]
}

// Name generates a random name composed of an adjective and a noun.
func (f *Faker) Name() string {
	return join3(f.Adjective(), "_", f.Noun())
}

// Names generates a random number of names, up to n.
func (f *Faker) Names(n int) []string {
	n = f.IntN(n + 1)
	names := make([]string, n)
	for i := range n {
		names[i] = f.Name()
	}
	return names
}

// DeploymentTier generates a random deployment tier such as prod or staging.
func (f *Faker) DeploymentTier() string {
	return tiers[f.IntN(len(tiers))]
}
