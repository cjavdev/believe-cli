// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
  "strings"
  "testing"

  "github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestTeamsLogoDelete(t *testing.T) {
  t.Skip("Mock server tests are disabled")
  t.Run("regular flags", func(t *testing.T) {
    mocktest.TestRunMockTestWithFlags(
      t,
      "--api-key", "string",
      "teams:logo", "delete",
      "--team-id", "team_id",
      "--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
    )
  })
}

func TestTeamsLogoDownload(t *testing.T) {
  t.Skip("Mock server tests are disabled")
  t.Run("regular flags", func(t *testing.T) {
    mocktest.TestRunMockTestWithFlags(
      t,
      "--api-key", "string",
      "teams:logo", "download",
      "--team-id", "team_id",
      "--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
    )
  })
}

func TestTeamsLogoUpload(t *testing.T) {
  t.Skip("Mock server tests are disabled")
  t.Run("regular flags", func(t *testing.T) {
    mocktest.TestRunMockTestWithFlags(
      t,
      "--api-key", "string",
      "teams:logo", "upload",
      "--team-id", "team_id",
      "--file", mocktest.TestFile(t, "Example data"),
    )
  })

  t.Run("piping data", func(t *testing.T) {
    testFile := mocktest.TestFile(t, "Example data")
    // Test piping YAML data over stdin
    pipeDataStr := "file: Example data"
                  pipeDataStr = strings.ReplaceAll(pipeDataStr, "Example data", testFile)
                  pipeData := []byte(pipeDataStr)
    mocktest.TestRunMockTestWithPipeAndFlags(
      t, pipeData,
      "--api-key", "string",
      "teams:logo", "upload",
      "--team-id", "team_id",
    )
  })
}
