// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestPressSimulate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"press", "simulate",
		"--api-key", "string",
		"--question", "Ted, your team just lost 5-0. How do you explain this embarrassing defeat?",
		"--hostile=true",
		"--topic", "match_result",
	)
}
