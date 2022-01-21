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
	"encoding/json"
	"testing"

	flowpb "github.com/cilium/cilium/api/v1/flow"
	observerpb "github.com/cilium/cilium/api/v1/observer"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func Test_JSON(t *testing.T) {
	for i := 0; i < 1000; i++ {
		flow := New()
		resp1 := observerpb.GetFlowsResponse{
			NodeName: "test-node",
			Time:     flow.Time,
			ResponseTypes: &observerpb.GetFlowsResponse_Flow{
				Flow: flow,
			},
		}

		b, err := json.Marshal(&resp1)
		if err != nil {
			t.Fatal(err)
		}

		var resp2 observerpb.GetFlowsResponse
		if err := json.Unmarshal(b, &resp2); err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(
			resp1.ResponseTypes,
			resp2.ResponseTypes,
			cmpopts.IgnoreUnexported(
				flowpb.CiliumEventType{},
				flowpb.Endpoint{},
				flowpb.Ethernet{},
				flowpb.Flow{},
				flowpb.ICMPv4{},
				flowpb.ICMPv6{},
				flowpb.IP{},
				flowpb.Layer4{},
				flowpb.TCP{},
				flowpb.TCPFlags{},
				flowpb.Service{},
				flowpb.UDP{},
				timestamppb.Timestamp{},
				wrapperspb.BoolValue{},
			),
		); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
}
