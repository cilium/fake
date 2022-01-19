// Copyright 2021-2022 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
