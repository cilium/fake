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
