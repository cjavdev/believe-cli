// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
	"github.com/cjavdev/believe-cli/internal/requestflag"
)

func TestMatchesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"matches", "create",
		"--away-team-id", "tottenham",
		"--date", "2024-02-20T19:45:00Z",
		"--home-team-id", "afc-richmond",
		"--match-type", "cup",
		"--attendance", "24500",
		"--away-score", "0",
		"--episode-id", "s02e05",
		"--home-score", "0",
		"--lesson-learned", "It's not about the wins and losses, it's about helping these young fellas be the best versions of themselves.",
		"--possession-percentage", "50",
		"--result", "pending",
		"--ted-halftime-speech", "You know what the happiest animal on Earth is? It's a goldfish. You know why? It's got a 10-second memory.",
		"--ticket-revenue-gbp", "735000.00",
		"--turning-point", "{description: description, emotional_impact: Galvanized the team's fighting spirit, minute: 0, character_involved: jamie-tartt}",
		"--weather-temp-celsius", "8.5",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(matchesCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"matches", "create",
		"--away-team-id", "tottenham",
		"--date", "2024-02-20T19:45:00Z",
		"--home-team-id", "afc-richmond",
		"--match-type", "cup",
		"--attendance", "24500",
		"--away-score", "0",
		"--episode-id", "s02e05",
		"--home-score", "0",
		"--lesson-learned", "It's not about the wins and losses, it's about helping these young fellas be the best versions of themselves.",
		"--possession-percentage", "50",
		"--result", "pending",
		"--ted-halftime-speech", "You know what the happiest animal on Earth is? It's a goldfish. You know why? It's got a 10-second memory.",
		"--ticket-revenue-gbp", "735000.00",
		"--turning-point.description", "description",
		"--turning-point.emotional-impact", "Galvanized the team's fighting spirit",
		"--turning-point.minute", "0",
		"--turning-point.character-involved", "jamie-tartt",
		"--weather-temp-celsius", "8.5",
	)
}

func TestMatchesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"matches", "retrieve",
		"--match-id", "match_id",
	)
}

func TestMatchesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"matches", "update",
		"--match-id", "match_id",
		"--attendance", "0",
		"--away-score", "0",
		"--away-team-id", "away_team_id",
		"--date", "2019-12-27T18:11:19.117Z",
		"--episode-id", "episode_id",
		"--home-score", "0",
		"--home-team-id", "home_team_id",
		"--lesson-learned", "lesson_learned",
		"--match-type", "league",
		"--possession-percentage", "0",
		"--result", "win",
		"--ted-halftime-speech", "ted_halftime_speech",
		"--ticket-revenue-gbp", "0",
		"--turning-point", "[{description: description, emotional_impact: Galvanized the team's fighting spirit, minute: 0, character_involved: jamie-tartt}]",
		"--weather-temp-celsius", "-30",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(matchesUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"matches", "update",
		"--match-id", "match_id",
		"--attendance", "0",
		"--away-score", "0",
		"--away-team-id", "away_team_id",
		"--date", "2019-12-27T18:11:19.117Z",
		"--episode-id", "episode_id",
		"--home-score", "0",
		"--home-team-id", "home_team_id",
		"--lesson-learned", "lesson_learned",
		"--match-type", "league",
		"--possession-percentage", "0",
		"--result", "win",
		"--ted-halftime-speech", "ted_halftime_speech",
		"--ticket-revenue-gbp", "0",
		"--turning-point.description", "description",
		"--turning-point.emotional-impact", "Galvanized the team's fighting spirit",
		"--turning-point.minute", "0",
		"--turning-point.character-involved", "jamie-tartt",
		"--weather-temp-celsius", "-30",
	)
}

func TestMatchesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"matches", "list",
		"--limit", "10",
		"--match-type", "league",
		"--result", "win",
		"--skip", "0",
		"--team-id", "team_id",
	)
}

func TestMatchesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"matches", "delete",
		"--match-id", "match_id",
	)
}

func TestMatchesGetLesson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"matches", "get-lesson",
		"--match-id", "match_id",
	)
}

func TestMatchesGetTurningPoints(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"matches", "get-turning-points",
		"--match-id", "match_id",
	)
}

func TestMatchesStreamLive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"matches", "stream-live",
		"--away-team", "away_team",
		"--excitement-level", "1",
		"--home-team", "home_team",
		"--speed", "0.1",
	)
}
