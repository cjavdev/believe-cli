// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
	"github.com/cjavdev/believe-cli/internal/requestflag"
)

func TestCharactersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"characters", "create",
		"--api-key", "string",
		"--background", "Legendary midfielder for Chelsea and AFC Richmond, now assistant coach. Known for his gruff exterior hiding a heart of gold.",
		"--emotional-stats", "{curiosity: 40, empathy: 85, optimism: 45, resilience: 95, vulnerability: 60}",
		"--name", "Roy Kent",
		"--personality-trait", "intense",
		"--personality-trait", "loyal",
		"--personality-trait", "secretly caring",
		"--personality-trait", "profane",
		"--role", "coach",
		"--date-of-birth", "'1977-03-15'",
		"--email", "roy.kent@afcrichmond.com",
		"--growth-arc", "{breakthrough: Finding purpose beyond playing, challenge: Accepting his body's limitations, ending_point: Retired but lost, season: 1, starting_point: Aging footballer facing retirement}",
		"--height-meters", "1.78",
		"--profile-image-url", "https://afcrichmond.com/images/roy-kent.jpg",
		"--salary-gbp", "'175000.00'",
		"--signature-quote", "He's here, he's there, he's every-f***ing-where, Roy Kent!",
		"--signature-quote", "Whistle!",
		"--team-id", "afc-richmond",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(charactersCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"characters", "create",
		"--background", "Legendary midfielder for Chelsea and AFC Richmond, now assistant coach. Known for his gruff exterior hiding a heart of gold.",
		"--emotional-stats.curiosity", "40",
		"--emotional-stats.empathy", "85",
		"--emotional-stats.optimism", "45",
		"--emotional-stats.resilience", "95",
		"--emotional-stats.vulnerability", "60",
		"--name", "Roy Kent",
		"--personality-trait", "intense",
		"--personality-trait", "loyal",
		"--personality-trait", "secretly caring",
		"--personality-trait", "profane",
		"--role", "coach",
		"--date-of-birth", "'1977-03-15'",
		"--email", "roy.kent@afcrichmond.com",
		"--growth-arc.breakthrough", "Finding purpose beyond playing",
		"--growth-arc.challenge", "Accepting his body's limitations",
		"--growth-arc.ending-point", "Retired but lost",
		"--growth-arc.season", "1",
		"--growth-arc.starting-point", "Aging footballer facing retirement",
		"--height-meters", "1.78",
		"--profile-image-url", "https://afcrichmond.com/images/roy-kent.jpg",
		"--salary-gbp", "'175000.00'",
		"--signature-quote", "He's here, he's there, he's every-f***ing-where, Roy Kent!",
		"--signature-quote", "Whistle!",
		"--team-id", "afc-richmond",
	)
}

func TestCharactersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"characters", "retrieve",
		"--api-key", "string",
		"--character-id", "character_id",
	)
}

func TestCharactersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"characters", "update",
		"--api-key", "string",
		"--character-id", "character_id",
		"--background", "background",
		"--date-of-birth", "'2019-12-27'",
		"--email", "dev@stainless.com",
		"--emotional-stats", "{curiosity: 99, empathy: 100, optimism: 95, resilience: 90, vulnerability: 80}",
		"--growth-arc", "[{breakthrough: breakthrough, challenge: challenge, ending_point: ending_point, season: 1, starting_point: starting_point}]",
		"--height-meters", "1",
		"--name", "x",
		"--personality-trait", "[string]",
		"--profile-image-url", "https://example.com",
		"--role", "coach",
		"--salary-gbp", "0",
		"--signature-quote", "[string]",
		"--team-id", "team_id",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(charactersUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"characters", "update",
		"--character-id", "character_id",
		"--background", "background",
		"--date-of-birth", "'2019-12-27'",
		"--email", "dev@stainless.com",
		"--emotional-stats.curiosity", "99",
		"--emotional-stats.empathy", "100",
		"--emotional-stats.optimism", "95",
		"--emotional-stats.resilience", "90",
		"--emotional-stats.vulnerability", "80",
		"--growth-arc.breakthrough", "breakthrough",
		"--growth-arc.challenge", "challenge",
		"--growth-arc.ending-point", "ending_point",
		"--growth-arc.season", "1",
		"--growth-arc.starting-point", "starting_point",
		"--height-meters", "1",
		"--name", "x",
		"--personality-trait", "[string]",
		"--profile-image-url", "https://example.com",
		"--role", "coach",
		"--salary-gbp", "0",
		"--signature-quote", "[string]",
		"--team-id", "team_id",
	)
}

func TestCharactersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"characters", "list",
		"--api-key", "string",
		"--limit", "10",
		"--min-optimism", "0",
		"--role", "coach",
		"--skip", "0",
		"--team-id", "team_id",
	)
}

func TestCharactersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"characters", "delete",
		"--api-key", "string",
		"--character-id", "character_id",
	)
}

func TestCharactersGetQuotes(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"characters", "get-quotes",
		"--api-key", "string",
		"--character-id", "character_id",
	)
}
