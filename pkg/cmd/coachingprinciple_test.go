// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestCoachingPrinciplesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "coaching:principles", "retrieve",
			"--api-key", "string",
			"--principle-id", "principle_id",
		)
	})
}

func TestCoachingPrinciplesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "coaching:principles", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--limit", "10",
			"--skip", "0",
		)
	})
}

func TestCoachingPrinciplesGetRandom(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "coaching:principles", "get-random",
			"--api-key", "string",
		)
	})
}
