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
	"math/rand"
)

// Labels generates a random set of labels.
func Labels() []string {
	var l []string
	for _, name := range labels {
		if rand.Intn(2) == 0 { // 50% chance of picking up this label
			l = append(l, name+"="+App())
		}
	}
	return l
}
