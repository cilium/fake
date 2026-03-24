// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package flow

import (
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// IsReply implements FlowFaker for flowfaker.
func (f *flowfaker) IsReply() *wrapperspb.BoolValue {
	switch f.IntN(3) {
	case 0:
		return &wrapperspb.BoolValue{Value: false}
	case 1:
		return &wrapperspb.BoolValue{Value: true}
	default:
		return nil
	}
}
