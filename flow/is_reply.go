// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

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
