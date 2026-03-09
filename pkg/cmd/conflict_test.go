// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestConflictsResolve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "conflicts", "resolve",
			"--api-key", "string",
			"--conflict-type", "interpersonal",
			"--description", "Alex keeps taking credit for my ideas in meetings and I'm getting resentful.",
			"--parties-involved", "Me",
			"--parties-involved", "My teammate Alex",
			"--attempts-made", "[Mentioned it casually, Avoided them]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"conflict_type: interpersonal\n" +
			"description: Alex keeps taking credit for my ideas in meetings and I'm getting resentful.\n" +
			"parties_involved:\n" +
			"  - Me\n" +
			"  - My teammate Alex\n" +
			"attempts_made:\n" +
			"  - Mentioned it casually\n" +
			"  - Avoided them\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "conflicts", "resolve",
			"--api-key", "string",
		)
	})
}
