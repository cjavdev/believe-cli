// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestTeamMembersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "team-members", "create",
			"--api-key", "string",
			"--member", "{character_id: jamie-tartt, jersey_number: 9, position: forward, team_id: afc-richmond, years_with_team: 3, assists: 23, goals_scored: 47, is_captain: false, member_type: player}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"character_id: jamie-tartt\n" +
			"jersey_number: 9\n" +
			"position: forward\n" +
			"team_id: afc-richmond\n" +
			"years_with_team: 3\n" +
			"assists: 23\n" +
			"goals_scored: 47\n" +
			"is_captain: false\n" +
			"member_type: player\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "team-members", "create",
			"--api-key", "string",
		)
	})
}

func TestTeamMembersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "team-members", "retrieve",
			"--api-key", "string",
			"--member-id", "member_id",
		)
	})
}

func TestTeamMembersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "team-members", "update",
			"--api-key", "string",
			"--member-id", "member_id",
			"--updates", "{assists: 0, goals_scored: 0, is_captain: true, jersey_number: 1, position: goalkeeper, team_id: team_id, years_with_team: 0}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"assists: 0\n" +
			"goals_scored: 0\n" +
			"is_captain: true\n" +
			"jersey_number: 1\n" +
			"position: goalkeeper\n" +
			"team_id: team_id\n" +
			"years_with_team: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "team-members", "update",
			"--api-key", "string",
			"--member-id", "member_id",
		)
	})
}

func TestTeamMembersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "team-members", "list",
			"--api-key", "string",
			"--limit", "10",
			"--member-type", "player",
			"--skip", "0",
			"--team-id", "team_id",
		)
	})
}

func TestTeamMembersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "team-members", "delete",
			"--api-key", "string",
			"--member-id", "member_id",
		)
	})
}

func TestTeamMembersListCoaches(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "team-members", "list-coaches",
			"--api-key", "string",
			"--limit", "10",
			"--skip", "0",
			"--specialty", "head_coach",
			"--team-id", "team_id",
		)
	})
}

func TestTeamMembersListPlayers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "team-members", "list-players",
			"--api-key", "string",
			"--limit", "10",
			"--position", "goalkeeper",
			"--skip", "0",
			"--team-id", "team_id",
		)
	})
}

func TestTeamMembersListStaff(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "team-members", "list-staff",
			"--api-key", "string",
			"--limit", "10",
			"--skip", "0",
			"--team-id", "team_id",
		)
	})
}
