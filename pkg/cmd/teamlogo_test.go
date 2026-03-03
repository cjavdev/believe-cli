// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestTeamsLogoDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"teams:logo", "delete",
		"--api-key", "string",
		"--team-id", "team_id",
		"--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestTeamsLogoDownload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"teams:logo", "download",
		"--api-key", "string",
		"--team-id", "team_id",
		"--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestTeamsLogoUpload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"teams:logo", "upload",
		"--api-key", "string",
		"--team-id", "team_id",
		"--file", "...",
	)
}
