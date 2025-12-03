// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package fake

import (
	"encoding/binary"
	"io"
	randv2 "math/rand/v2"
	"strings"
)

// Faker is the main interface exposed to generate fake data.
type Faker interface { //nolint:interfacebloat
	// Adjective generates a random adjective.
	Adjective() string
	// AlphaNum generates a random alphanumeric string of the given length.
	AlphaNum(length int) string
	// App generates a random software application name.
	App() string
	// Noun generates a random noun.
	Noun() string
	// Name generates a random name.
	Name() string
	// Names generates a random set of names. It panics if n < 0.
	Names(n int) []string
	// DeploymentTier generates a random software deployment tier such as prod,
	// staging, etc.
	DeploymentTier() string

	// K8sLabels generates a random set of Kubernetes labels.
	K8sLabels() []string
	// K8sNamespace generates a random Kubernetes namespace name.
	K8sNamespace() string
	// K8sNodeName generates a random Kubernetes node name.
	K8sNodeName() string
	// K8sPodName generates a random Kubernetes pod name.
	K8sPodName() string

	// MAC generates a random MAC address.
	MAC() string
	// IP generates a random IP address. Options may be provided to specify a
	// network for the address or if it should be IPv4 or IPv6.
	IP(options ...IPOption) string
	// Port generates a random port number between 1 and 65535 or in the range
	// specified by the given option.
	Port(options ...PortOption) uint32
}

// New creates a new Faker using a random seed.
func New() Faker {
	// NOTE: We seed from the global math/rand/v2 source rather than
	// crypto/rand to avoid errors handling; the faker should not be used for
	// security-sensitive purposes. We use ChaCha8 rather than PCG as ChaCha8
	// implements io.Reader.
	var seed [32]byte
	for i := range 4 {
		binary.LittleEndian.PutUint64(seed[i*8:], randv2.Uint64())
	}
	return NewWithSource(randv2.NewChaCha8(seed))
}

// RandSourceReader is a random source that also implements io.Reader.
type RandSourceReader interface {
	randv2.Source
	io.Reader
}

// NewWithSource creates a new Faker using the given random source. Useful to
// control the faker output, e.g. for testing.
func NewWithSource(src RandSourceReader) Faker {
	return &faker{
		Rand:   randv2.New(src),
		Reader: src,
	}
}

// faker is a private struct implementing Faker.
type faker struct {
	*randv2.Rand
	io.Reader
}

// join3 is a helper to build a string composed of three parts.
func join3(left, middle, right string) string {
	var sb strings.Builder
	sb.Grow(len(left) + len(middle) + len(right))
	sb.WriteString(left)
	sb.WriteString(middle)
	sb.WriteString(right)
	return sb.String()
}
