// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/believe-cli/internal/mocktest"
)

func TestBiscuitsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"biscuits", "retrieve",
		"--biscuit-id", "biscuit_id",
	)
}

func TestBiscuitsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"biscuits", "list",
		"--limit", "10",
		"--skip", "0",
	)
}

func TestBiscuitsGetFresh(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"biscuits", "get-fresh",
	)
}
