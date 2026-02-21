// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestCoachingPrinciplesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"coaching:principles", "retrieve",
		"--principle-id", "principle_id",
	)
}

func TestCoachingPrinciplesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"coaching:principles", "list",
		"--limit", "10",
		"--skip", "0",
	)
}

func TestCoachingPrinciplesGetRandom(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"coaching:principles", "get-random",
	)
}
