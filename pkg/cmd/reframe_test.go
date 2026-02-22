// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestReframeTransformNegativeThoughts(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"reframe", "transform-negative-thoughts",
		"--negative-thought", "I'm not good enough for this job.",
		"--recurring=true",
	)
}
