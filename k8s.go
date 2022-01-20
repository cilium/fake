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

// K8sLabels generates a random set of Kubernetes labels.
func K8sLabels() []string {
	var l []string
	for _, name := range labels {
		if rand.Intn(2) == 0 { // 50% chance of picking up this label
			l = append(l, name+"="+App())
		}
	}
	return l
}

// K8sNamespace generates a random Kubernetes namespace name.
func K8sNamespace() string {
	if rand.Intn(2) == 0 {
		return namespaces[rand.Intn(len(namespaces))]
	}
	return fmt.Sprintf("%s-%s", App(), DeploymentTier())
}

// K8sNodeName generates a random Kubernetes node name.
func K8sNodeName() string {
	return fmt.Sprintf(
		"%s-%s",
		Adjective(),
		Noun(),
	)
}

// K8sPodName generates a random Kubernetes pod name.
func K8sPodName() string {
	return fmt.Sprintf(
		"%s-%s",
		App(),
		AlphaNum(5),
	)
}
