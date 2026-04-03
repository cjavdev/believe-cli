// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestBelieveSubmit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"believe", "submit",
			"--situation", "I just got passed over for a promotion I've been working toward for two years.",
			"--situation-type", "work_challenge",
			"--context", "I've always tried to be a team player and support my colleagues.",
			"--intensity", "7",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"situation: I just got passed over for a promotion I've been working toward for two years.\n" +
			"situation_type: work_challenge\n" +
			"context: I've always tried to be a team player and support my colleagues.\n" +
			"intensity: 7\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"believe", "submit",
		)
	})
}
