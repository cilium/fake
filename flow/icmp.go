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
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

// https://godoc.org/golang.org/x/net/ipv4#ICMPType
var icmpv4Types = []ipv4.ICMPType{
	ipv4.ICMPTypeEchoReply,              // Echo Reply
	ipv4.ICMPTypeDestinationUnreachable, // Destination Unreachable
	ipv4.ICMPTypeRedirect,               // Redirect
	ipv4.ICMPTypeEcho,                   // Echo
	ipv4.ICMPTypeRouterAdvertisement,    // Router Advertisement
	ipv4.ICMPTypeRouterSolicitation,     // Router Solicitation
	ipv4.ICMPTypeTimeExceeded,           // Time Exceeded
	ipv4.ICMPTypeParameterProblem,       // Parameter Problem
	ipv4.ICMPTypeTimestamp,              // Timestamp
	ipv4.ICMPTypeTimestampReply,         // Timescape Reply
	ipv4.ICMPTypePhoturis,               // Photuris
	ipv4.ICMPTypeExtendedEchoRequest,    // Extended Echo Request
	ipv4.ICMPTypeExtendedEchoReply,      // Extended Echo Reply
}

// ICMPv4 generates a random ICMPv4 flow.
func ICMPv4() *flowpb.ICMPv4 {
	t := icmpv4Types[rand.Intn(len(icmpv4Types))]

	// See https://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml
	// https://tools.ietf.org/html/rfc792
	// https://tools.ietf.org/html/rfc1122
	// https://tools.ietf.org/html/rfc1812
	var c int
	switch t {
	case ipv4.ICMPTypeEchoReply:
	case ipv4.ICMPTypeDestinationUnreachable:
		c = rand.Intn(16)
	case ipv4.ICMPTypeRedirect:
		c = rand.Intn(4)
	case ipv4.ICMPTypeEcho:
	case ipv4.ICMPTypeRouterAdvertisement:
		if rand.Intn(2) == 0 { // 1 in 2 chances
			c = 16
		}
	case ipv4.ICMPTypeRouterSolicitation:
	case ipv4.ICMPTypeTimeExceeded:
		c = rand.Intn(2)
	case ipv4.ICMPTypeParameterProblem:
		c = rand.Intn(3)
	case ipv4.ICMPTypeTimestamp:
	case ipv4.ICMPTypeTimestampReply:
	case ipv4.ICMPTypePhoturis:
		c = rand.Intn(6)
	case ipv4.ICMPTypeExtendedEchoRequest:
		c = rand.Intn(5)
	case ipv4.ICMPTypeExtendedEchoReply:
		c = rand.Intn(5)
	}

	return &flowpb.ICMPv4{
		Type: uint32(t),
		Code: uint32(c),
	}
}

// https://godoc.org/golang.org/x/net/ipv6#ICMPType
var icmpv6Types = []ipv6.ICMPType{
	ipv6.ICMPTypeDestinationUnreachable,
	ipv6.ICMPTypePacketTooBig,
	ipv6.ICMPTypeTimeExceeded,
	ipv6.ICMPTypeParameterProblem,
	ipv6.ICMPTypeEchoRequest,
	ipv6.ICMPTypeEchoReply,
	ipv6.ICMPTypeMulticastListenerQuery,
	ipv6.ICMPTypeMulticastListenerReport,
	ipv6.ICMPTypeMulticastListenerDone,
	ipv6.ICMPTypeRouterSolicitation,
	ipv6.ICMPTypeRouterAdvertisement,
	ipv6.ICMPTypeNeighborSolicitation,
	ipv6.ICMPTypeNeighborAdvertisement,
	ipv6.ICMPTypeRedirect,
	ipv6.ICMPTypeRouterRenumbering,
	ipv6.ICMPTypeNodeInformationQuery,
	ipv6.ICMPTypeNodeInformationResponse,
	ipv6.ICMPTypeInverseNeighborDiscoverySolicitation,
	ipv6.ICMPTypeInverseNeighborDiscoveryAdvertisement,
	ipv6.ICMPTypeVersion2MulticastListenerReport,
	ipv6.ICMPTypeHomeAgentAddressDiscoveryRequest,
	ipv6.ICMPTypeHomeAgentAddressDiscoveryReply,
	ipv6.ICMPTypeMobilePrefixSolicitation,
	ipv6.ICMPTypeMobilePrefixAdvertisement,
	ipv6.ICMPTypeCertificationPathSolicitation,
	ipv6.ICMPTypeCertificationPathAdvertisement,
	ipv6.ICMPTypeMulticastRouterAdvertisement,
	ipv6.ICMPTypeMulticastRouterSolicitation,
	ipv6.ICMPTypeMulticastRouterTermination,
	ipv6.ICMPTypeFMIPv6,
	ipv6.ICMPTypeRPLControl,
	ipv6.ICMPTypeILNPv6LocatorUpdate,
	ipv6.ICMPTypeDuplicateAddressRequest,
	ipv6.ICMPTypeDuplicateAddressConfirmation,
	ipv6.ICMPTypeMPLControl,
	ipv6.ICMPTypeExtendedEchoRequest,
	ipv6.ICMPTypeExtendedEchoReply,
}

// ICMPv6 generates a random ICMPv6 flow.
func ICMPv6() *flowpb.ICMPv6 {
	t := icmpv6Types[rand.Intn(len(icmpv6Types))]

	// See https://www.iana.org/assignments/icmpv6-parameters/icmpv6-parameters.xhtml
	// https://tools.ietf.org/html/rfc4443
	// https://tools.ietf.org/html/rfc8335
	var c int
	switch t {
	case ipv6.ICMPTypeDestinationUnreachable:
		c = rand.Intn(16)
	case ipv6.ICMPTypePacketTooBig:
	case ipv6.ICMPTypeTimeExceeded:
		c = rand.Intn(2)
	case ipv6.ICMPTypeParameterProblem:
		c = rand.Intn(3)
	case ipv6.ICMPTypeEchoRequest:
	case ipv6.ICMPTypeEchoReply:
	case ipv6.ICMPTypeMulticastListenerQuery:
	case ipv6.ICMPTypeMulticastListenerReport:
	case ipv6.ICMPTypeMulticastListenerDone:
	case ipv6.ICMPTypeRouterSolicitation:
	case ipv6.ICMPTypeRouterAdvertisement:
	case ipv6.ICMPTypeNeighborSolicitation:
	case ipv6.ICMPTypeNeighborAdvertisement:
	case ipv6.ICMPTypeRedirect:
		c = rand.Intn(4)
	case ipv6.ICMPTypeRouterRenumbering:
		if rand.Intn(3) == 2 { // 1 in 3 chance
			c = 255
		} else {
			c = rand.Intn(2)
		}
	case ipv6.ICMPTypeNodeInformationQuery:
		c = rand.Intn(3)
	case ipv6.ICMPTypeNodeInformationResponse:
		c = rand.Intn(3)
	case ipv6.ICMPTypeInverseNeighborDiscoverySolicitation:
	case ipv6.ICMPTypeInverseNeighborDiscoveryAdvertisement:
	case ipv6.ICMPTypeVersion2MulticastListenerReport:
	case ipv6.ICMPTypeHomeAgentAddressDiscoveryRequest:
	case ipv6.ICMPTypeHomeAgentAddressDiscoveryReply:
	case ipv6.ICMPTypeMobilePrefixSolicitation:
	case ipv6.ICMPTypeMobilePrefixAdvertisement:
	case ipv6.ICMPTypeCertificationPathSolicitation:
	case ipv6.ICMPTypeCertificationPathAdvertisement:
	case ipv6.ICMPTypeMulticastRouterAdvertisement:
	case ipv6.ICMPTypeMulticastRouterSolicitation:
	case ipv6.ICMPTypeMulticastRouterTermination:
	case ipv6.ICMPTypeFMIPv6:
		c = rand.Intn(6)
	case ipv6.ICMPTypeRPLControl:
	case ipv6.ICMPTypeILNPv6LocatorUpdate:
	case ipv6.ICMPTypeDuplicateAddressRequest:
	case ipv6.ICMPTypeDuplicateAddressConfirmation:
	case ipv6.ICMPTypeMPLControl:
	case ipv6.ICMPTypeExtendedEchoRequest:
	case ipv6.ICMPTypeExtendedEchoReply:
		c = rand.Intn(5)
	}

	return &flowpb.ICMPv6{
		Type: uint32(t),
		Code: uint32(c),
	}
}
