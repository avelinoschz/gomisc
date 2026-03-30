package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyMessage(t *testing.T) {
	tests := []struct {
		desc string

		message  string
		expected bool
	}{
		{"happy path", "Hello", true},
		{"wrong msg", "World", false},
		{"empty string", "", false},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			result := VerifyingMessage(test.message)
			assert.Equal(t, test.expected, result)
		})
	}
}
