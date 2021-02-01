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
	"net"
)

// a MAC address have different sizes. For the purpose of simplicity, only
// support 6 bytes MAC.
const macLen = 6

// MAC generates a random MAC address.
func MAC() string {
	hw := make(net.HardwareAddr, macLen)
	_, _ = rand.Read(hw)
	return hw.String()
}
