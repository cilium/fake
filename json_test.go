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
	"encoding/json"
	"reflect"
	"testing"

	observerpb "github.com/cilium/cilium/api/v1/observer"
)

func doTestFakeJSON(t *testing.T) {
	flow := Flow()
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

	if !reflect.DeepEqual(resp1.ResponseTypes, resp2.ResponseTypes) {
		t.FailNow()
	}

}

func Test_JSON(t *testing.T) {
	for i := 0; i < 1000; i++ {
		doTestFakeJSON(t)
	}
}
