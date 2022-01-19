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

package flow

import (
	"math/rand"

	flowpb "github.com/cilium/cilium/api/v1/flow"
	"github.com/cilium/cilium/pkg/monitor/api"
)

var allCiliumEventTypes = []int32{
	api.MessageTypeUnspec,
	api.MessageTypeDrop,
	api.MessageTypeDebug,
	api.MessageTypeCapture,
	api.MessageTypeTrace,
	api.MessageTypePolicyVerdict,
	api.MessageTypeAccessLog,
	api.MessageTypeAgent,
}

// CiliumEventType generates a random CiliumEventType.
func CiliumEventType() *flowpb.CiliumEventType {
	typ := allCiliumEventTypes[rand.Intn(len(allCiliumEventTypes))]
	if typ == api.MessageTypeUnspec {
		return nil
	}
	return &flowpb.CiliumEventType{
		Type: typ,
		// NOTE: AgentNotify* are the most numerous.
		SubType: int32(rand.Intn(13)),
	}
}
