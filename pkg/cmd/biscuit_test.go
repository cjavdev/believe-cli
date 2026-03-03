// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestBiscuitsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"biscuits", "retrieve",
		"--api-key", "string",
		"--biscuit-id", "biscuit_id",
	)
}

func TestBiscuitsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"biscuits", "list",
		"--api-key", "string",
		"--limit", "10",
		"--skip", "0",
	)
}

func TestBiscuitsGetFresh(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"biscuits", "get-fresh",
		"--api-key", "string",
	)
}
