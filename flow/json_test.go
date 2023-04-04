// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

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
				flowpb.Workload{},
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
				flowpb.SCTP{},
				flowpb.TraceContext{},
				flowpb.TraceParent{},
				timestamppb.Timestamp{},
				wrapperspb.BoolValue{},
			),
		); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
}
