// Copyright 2022 Authors of Cilium
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
	assert.Equal(t, len(names), 0)

	max := 100
	names = fake.Names(max)
	assert.NotNil(t, names)
	assert.LessOrEqual(t, len(names), max)
	for _, n := range names {
		assert.NotEmpty(t, n)
	}
}
