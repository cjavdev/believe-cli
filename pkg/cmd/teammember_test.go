// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/believe-cli/internal/mocktest"
)

func TestTeamMembersCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"team-members", "create",
		"--member", "{character_id: jamie-tartt, jersey_number: 9, position: forward, team_id: afc-richmond, years_with_team: 3, assists: 23, goals_scored: 47, is_captain: false, member_type: player}",
	)
}

func TestTeamMembersRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"team-members", "retrieve",
		"--member-id", "member_id",
	)
}

func TestTeamMembersUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"team-members", "update",
		"--member-id", "member_id",
		"--updates", "{assists: 0, goals_scored: 0, is_captain: true, jersey_number: 1, position: goalkeeper, team_id: team_id, years_with_team: 0}",
	)
}

func TestTeamMembersList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"team-members", "list",
		"--limit", "10",
		"--member-type", "player",
		"--skip", "0",
		"--team-id", "team_id",
	)
}

func TestTeamMembersDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"team-members", "delete",
		"--member-id", "member_id",
	)
}

func TestTeamMembersListCoaches(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"team-members", "list-coaches",
		"--limit", "10",
		"--skip", "0",
		"--specialty", "head_coach",
		"--team-id", "team_id",
	)
}

func TestTeamMembersListPlayers(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"team-members", "list-players",
		"--limit", "10",
		"--position", "goalkeeper",
		"--skip", "0",
		"--team-id", "team_id",
	)
}

func TestTeamMembersListStaff(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"team-members", "list-staff",
		"--limit", "10",
		"--skip", "0",
		"--team-id", "team_id",
	)
}
