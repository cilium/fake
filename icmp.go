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
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

// https://godoc.org/golang.org/x/net/ipv4#ICMPType
var icmpv4Types = []ipv4.ICMPType{
	ipv4.ICMPTypeEchoReply,
	ipv4.ICMPTypeDestinationUnreachable,
	ipv4.ICMPTypeRedirect,
	ipv4.ICMPTypeEcho,
	ipv4.ICMPTypeRouterAdvertisement,
	ipv4.ICMPTypeRouterSolicitation,
	ipv4.ICMPTypeTimeExceeded,
	ipv4.ICMPTypeParameterProblem,
	ipv4.ICMPTypeTimestamp,
	ipv4.ICMPTypeTimestampReply,
	ipv4.ICMPTypePhoturis,
	ipv4.ICMPTypeExtendedEchoRequest,
	ipv4.ICMPTypeExtendedEchoReply,
}

// ICMPv4 generates a random ICMPv4 flow.
func ICMPv4() *flowpb.ICMPv4 {
	t := icmpv4Types[rand.Intn(len(icmpv4Types))]

	// https://tools.ietf.org/html/rfc792
	// https://tools.ietf.org/html/rfc1122
	// https://tools.ietf.org/html/rfc1812
	c := 0
	switch t {
	case ipv4.ICMPTypeDestinationUnreachable:
		c = rand.Intn(16)
	case ipv4.ICMPTypeRedirect:
		c = rand.Intn(4)
	case ipv4.ICMPTypeRouterAdvertisement:
		if rand.Intn(2) == 0 { // 1 in 2 chances
			c = 16
		}
	case ipv4.ICMPTypeTimeExceeded:
		c = rand.Intn(2)
	case ipv4.ICMPTypePhoturis:
		c = rand.Intn(6)
	case ipv4.ICMPTypeParameterProblem:
		c = rand.Intn(3)
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

	// https://tools.ietf.org/html/rfc4443
	// https://tools.ietf.org/html/rfc8335
	c := 0
	switch t {
	case ipv6.ICMPTypeDestinationUnreachable:
		c = rand.Intn(16)
	case ipv6.ICMPTypeTimeExceeded:
		c = rand.Intn(2)
	case ipv6.ICMPTypeParameterProblem:
		c = rand.Intn(3)
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
	case ipv6.ICMPTypeExtendedEchoReply:
		c = rand.Intn(5)
	}

	return &flowpb.ICMPv6{
		Type: uint32(t),
		Code: uint32(c),
	}
}
