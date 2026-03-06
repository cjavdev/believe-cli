// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestPressSimulate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "press", "simulate",
			"--api-key", "string",
			"--question", "Ted, your team just lost 5-0. How do you explain this embarrassing defeat?",
			"--hostile=true",
			"--topic", "match_result",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"question: Ted, your team just lost 5-0. How do you explain this embarrassing defeat?\n" +
			"hostile: true\n" +
			"topic: match_result\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "press", "simulate",
			"--api-key", "string",
		)
	})
}
