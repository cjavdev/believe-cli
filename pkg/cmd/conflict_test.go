// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestConflictsResolve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"conflicts", "resolve",
		"--conflict-type", "interpersonal",
		"--description", "Alex keeps taking credit for my ideas in meetings and I'm getting resentful.",
		"--parties-involved", "Me",
		"--parties-involved", "My teammate Alex",
		"--attempts-made", "Mentioned it casually",
		"--attempts-made", "Avoided them",
	)
}
