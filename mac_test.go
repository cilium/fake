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
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MAC(t *testing.T) {
	for i := 0; i < 20; i++ {
		hw := MAC()
		t.Run(hw, func(t *testing.T) {
			_, err := net.ParseMAC(hw)
			assert.Nil(t, err)
		})
	}
}
