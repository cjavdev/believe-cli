// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestBelieveSubmit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"believe", "submit",
		"--api-key", "string",
		"--situation", "I just got passed over for a promotion I've been working toward for two years.",
		"--situation-type", "work_challenge",
		"--context", "I've always tried to be a team player and support my colleagues.",
		"--intensity", "7",
	)
}
