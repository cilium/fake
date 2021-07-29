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

// TCPFlags represents a set TCP flags.
type TCPFlags uint16

// These constants represent possible individual TCP flags.
const (
	TCPFlagFIN TCPFlags = 1 << iota
	TCPFlagSYN
	TCPFlagRST
	TCPFlagPSH
	TCPFlagACK
	TCPFlagURG
	TCPFlagECE
	TCPFlagCWR
	TCPFlagNS
)

type layer4Options struct {
	icmpv4, icmpv6         bool
	tcp                    bool
	tcpFlags               *flowpb.TCPFlags
	tcpSrcPort, tcpDstPort uint32
	udp                    bool
	udpSrcPort, udpDstPort uint32
}

// Layer4Option is an option to use with Layer4.
type Layer4Option interface {
	apply(o *layer4Options)
}

type funcLayer4Option func(*layer4Options)

func (fl4o funcLayer4Option) apply(l4o *layer4Options) {
	fl4o(l4o)
}

// WithLayer4ICMPv4 specifies that the Layer4 should be ICMPv4.
func WithLayer4ICMPv4() Layer4Option {
	return funcLayer4Option(func(o *layer4Options) {
		o.icmpv4 = true
	})
}

// WithLayer4ICMPv6 specifies that the Layer4 should be ICMPv6.
func WithLayer4ICMPv6() Layer4Option {
	return funcLayer4Option(func(o *layer4Options) {
		o.icmpv6 = true
	})
}

// WithLayer4TCP specifies that the Layer4 should be TCP.
func WithLayer4TCP() Layer4Option {
	return funcLayer4Option(func(o *layer4Options) {
		o.tcp = true
	})
}

// WithLayer4TCPFlags specifies the TCP flags for Layer4. Only useful in
// conjunction with WithLayer4TCP.
func WithLayer4TCPFlags(flags TCPFlags) Layer4Option {
	return funcLayer4Option(func(o *layer4Options) {
		o.tcpFlags = &flowpb.TCPFlags{
			FIN: flags&TCPFlagFIN != 0,
			SYN: flags&TCPFlagSYN != 0,
			RST: flags&TCPFlagRST != 0,
			PSH: flags&TCPFlagPSH != 0,
			ACK: flags&TCPFlagACK != 0,
			URG: flags&TCPFlagURG != 0,
			ECE: flags&TCPFlagECE != 0,
			CWR: flags&TCPFlagCWR != 0,
			NS:  flags&TCPFlagNS != 0,
		}
	})
}

// WithLayer4TCPSourcePort specifies the TCP source port for Layer4. Only
// useful in conjunction with WithLayer4TCP.
func WithLayer4TCPSourcePort(port uint32) Layer4Option {
	return funcLayer4Option(func(o *layer4Options) {
		o.tcpSrcPort = port
	})
}

// WithLayer4TCPDestinationPort specifies the TCP destination port for Layer4.
// Only useful in conjunction with WithLayer4TCP.
func WithLayer4TCPDestinationPort(port uint32) Layer4Option {
	return funcLayer4Option(func(o *layer4Options) {
		o.tcpDstPort = port
	})
}

// WithLayer4UDP specifies that the Layer4 should be UDP.
func WithLayer4UDP() Layer4Option {
	return funcLayer4Option(func(o *layer4Options) {
		o.udp = true
	})
}

// WithLayer4UDPSourcePort specifies the UDP source port. Only useful in
// conjunction with WithLayer4UDP.
func WithLayer4UDPSourcePort(port uint32) Layer4Option {
	return funcLayer4Option(func(o *layer4Options) {
		o.udpSrcPort = port
	})
}

// WithLayer4UDPDestinationPort specifies the UDP destination port. Only useful
// in conjunction with WithLayer4UDP.
func WithLayer4UDPDestinationPort(port uint32) Layer4Option {
	return funcLayer4Option(func(o *layer4Options) {
		o.udpDstPort = port
	})
}

// Layer4 generates a layer 4. If no option is provided, it will be TCP.
func Layer4(options ...Layer4Option) *flowpb.Layer4 {
	opts := layer4Options{
		tcpSrcPort: Port(WithPortUser()),
		tcpDstPort: Port(WithPortUser()),
		tcpFlags:   randTCPFlags(),
		udpSrcPort: Port(WithPortUser()),
		udpDstPort: Port(WithPortUser()),
	}
	for _, opt := range options {
		opt.apply(&opts)
	}

	switch {
	default:
		fallthrough
	case opts.tcp:
		return &flowpb.Layer4{
			Protocol: &flowpb.Layer4_TCP{
				TCP: &flowpb.TCP{
					SourcePort:      opts.tcpSrcPort,
					DestinationPort: opts.tcpDstPort,
					Flags:           opts.tcpFlags,
				},
			},
		}
	case opts.udp:
		return &flowpb.Layer4{
			Protocol: &flowpb.Layer4_UDP{
				UDP: &flowpb.UDP{
					SourcePort:      opts.udpSrcPort,
					DestinationPort: opts.udpDstPort,
				},
			},
		}
	case opts.icmpv4:
		return &flowpb.Layer4{
			Protocol: &flowpb.Layer4_ICMPv4{
				ICMPv4: ICMPv4(),
			},
		}
	case opts.icmpv6:
		return &flowpb.Layer4{
			Protocol: &flowpb.Layer4_ICMPv6{
				ICMPv6: ICMPv6(),
			},
		}
	}
}

var tcpFlagsPatterns = []*flowpb.TCPFlags{
	{SYN: true},
	{SYN: true, ACK: true},
	{ACK: true},
	{ACK: true},
	{ACK: true},
	{ACK: true, PSH: true},
	{RST: true},
	{ACK: true, FIN: true},
}

func randTCPFlags() *flowpb.TCPFlags {
	return tcpFlagsPatterns[rand.Intn(len(tcpFlagsPatterns))]
}
