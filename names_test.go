// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package fake_test

import (
	"testing"

	"github.com/cilium/fake"
	"github.com/stretchr/testify/assert"
)

func TestNames(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Names(-1) should panic()")
		}
	}()
	_ = fake.Names(-1)

	names := fake.Names(0)
	assert.NotNil(t, names)
	assert.Empty(t, names)

	max := 100
	names = fake.Names(max)
	assert.NotNil(t, names)
	assert.LessOrEqual(t, len(names), max)
	for _, n := range names {
		assert.NotEmpty(t, n)
	}
}
