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
	"fmt"
	"math/rand"
)

// Adjective generates a random adjective.
func Adjective() string {
	return adjectives[rand.Intn(len(adjectives))]
}

// AlphaNum generates a random alphanumeric string of the given length.
func AlphaNum(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = alphanum[rand.Intn(len(alphanum))]
	}
	return string(b)
}

// App generates a random application name.
func App() string {
	return apps[rand.Intn(len(apps))]
}

// Noun generates a random noun.
func Noun() string {
	return nouns[rand.Intn(len(nouns))]
}

// Name generates a random name.
func Name() string {
	return fmt.Sprintf("%s_%s", Adjective(), Noun())
}

// Names generates a random set of names. Panic when max <= 0.
func Names(max int) []string {
	n := rand.Intn(max + 1)
	if n == 0 {
		return nil
	}
	names := make([]string, n)
	for i := 0; i < n; i++ {
		names[i] = Name()
	}
	return names
}

// Namespace generates a random namespace name.
func Namespace() string {
	if rand.Intn(2) == 0 {
		return namespaces[rand.Intn(len(namespaces))]
	}
	return fmt.Sprintf("%s-%s", App(), Tier())
}

// NodeName generates a random node name.
func NodeName() string {
	return fmt.Sprintf(
		"%s-%s",
		Adjective(),
		Noun(),
	)
}

// PodName generates a random pod name.
func PodName() string {
	return fmt.Sprintf(
		"%s-%s",
		App(),
		AlphaNum(5),
	)
}

// Tier generates a random tier such as prod, staging, etc.
func Tier() string {
	return tiers[rand.Intn(len(tiers))]
}
