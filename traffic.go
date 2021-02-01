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

	flowpb "github.com/cilium/cilium/api/v1/flow"
)

// TrafficDirection returns randomly traffic direction INGRESS or EGRESS.
func TrafficDirection() flowpb.TrafficDirection {
	if rand.Intn(2) == 0 { // 50% chance of picking up ingress
		return flowpb.TrafficDirection_INGRESS
	}
	return flowpb.TrafficDirection_EGRESS
}
