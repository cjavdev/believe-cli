// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestReframeTransformNegativeThoughts(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "reframe", "transform-negative-thoughts",
			"--api-key", "string",
			"--negative-thought", "I'm not good enough for this job.",
			"--recurring=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"negative_thought: I'm not good enough for this job.\n" +
			"recurring: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "reframe", "transform-negative-thoughts",
			"--api-key", "string",
		)
	})
}
