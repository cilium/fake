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

	"google.golang.org/protobuf/types/known/wrapperspb"
)

// IsReply returns either nil, or a wrapped boolean value reprenting true, or a
// wrapped boolean value reprenting false, with equal probability.
func IsReply() *wrapperspb.BoolValue {
	switch rand.Intn(3) {
	case 0:
		return &wrapperspb.BoolValue{Value: false}
	case 1:
		return &wrapperspb.BoolValue{Value: true}
	default:
		return nil
	}
}
