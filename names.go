// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package fake

import (
	"fmt"
	"math/rand/v2"
)

// Adjective generates a random adjective.
func Adjective() string {
	return adjectives[rand.IntN(len(adjectives))]
}

// AlphaNum generates a random alphanumeric string of the given length.
func AlphaNum(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = alphanum[rand.IntN(len(alphanum))]
	}
	return string(b)
}

// App generates a random software application name.
func App() string {
	return apps[rand.IntN(len(apps))]
}

// Noun generates a random noun.
func Noun() string {
	return nouns[rand.IntN(len(nouns))]
}

// Name generates a random name.
func Name() string {
	return fmt.Sprintf("%s_%s", Adjective(), Noun())
}

// Names generates a random set of names. It panics if n < 0.
func Names(n int) []string {
	n = rand.IntN(n + 1)
	names := make([]string, n)
	for i := range n {
		names[i] = Name()
	}
	return names
}

// DeploymentTier generates a random software deployment tier such as prod,
// staging, etc.
func DeploymentTier() string {
	return tiers[rand.IntN(len(tiers))]
}
