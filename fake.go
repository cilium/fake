// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package fake

import (
	"encoding/binary"
	"io"
	randv2 "math/rand/v2"
	"strings"
)

// RandSourceReader is a random source that also implements io.Reader.
type RandSourceReader interface {
	randv2.Source
	io.Reader
}

// Faker provide function to generate fake networking data and names.
type Faker struct {
	*randv2.Rand
	io.Reader
}

// New creates a new Faker using a random seed.
func New() *Faker {
	// NOTE: We seed from the global math/rand/v2 source rather than
	// crypto/rand to avoid errors handling; the Faker should not be used for
	// security-sensitive purposes. We use ChaCha8 rather than PCG as ChaCha8
	// implements io.Reader.
	var seed [32]byte
	for i := range 4 {
		binary.LittleEndian.PutUint64(seed[i*8:], randv2.Uint64())
	}
	return NewWithSource(randv2.NewChaCha8(seed))
}

// NewWithSource creates a new Faker using the given random source. Useful to
// control the Faker output, e.g. for testing.
func NewWithSource(src RandSourceReader) *Faker {
	return &Faker{
		Rand:   randv2.New(src),
		Reader: src,
	}
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
