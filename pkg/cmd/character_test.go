// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
	"github.com/cjavdev/believe-cli/internal/requestflag"
)

func TestCharactersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "characters", "create",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(charactersCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "characters", "create",
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"background: >-\n" +
			"  Legendary midfielder for Chelsea and AFC Richmond, now assistant coach. Known\n" +
			"  for his gruff exterior hiding a heart of gold.\n" +
			"emotional_stats:\n" +
			"  curiosity: 40\n" +
			"  empathy: 85\n" +
			"  optimism: 45\n" +
			"  resilience: 95\n" +
			"  vulnerability: 60\n" +
			"name: Roy Kent\n" +
			"personality_traits:\n" +
			"  - intense\n" +
			"  - loyal\n" +
			"  - secretly caring\n" +
			"  - profane\n" +
			"role: coach\n" +
			"date_of_birth: '1977-03-15'\n" +
			"email: roy.kent@afcrichmond.com\n" +
			"growth_arcs:\n" +
			"  - breakthrough: Finding purpose beyond playing\n" +
			"    challenge: Accepting his body's limitations\n" +
			"    ending_point: Retired but lost\n" +
			"    season: 1\n" +
			"    starting_point: Aging footballer facing retirement\n" +
			"height_meters: 1.78\n" +
			"profile_image_url: https://afcrichmond.com/images/roy-kent.jpg\n" +
			"salary_gbp: '175000.00'\n" +
			"signature_quotes:\n" +
			"  - He's here, he's there, he's every-f***ing-where, Roy Kent!\n" +
			"  - Whistle!\n" +
			"team_id: afc-richmond\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "characters", "create",
			"--api-key", "string",
		)
	})
}

func TestCharactersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "characters", "retrieve",
			"--api-key", "string",
			"--character-id", "character_id",
		)
	})
}

func TestCharactersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "characters", "update",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(charactersUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "characters", "update",
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"background: background\n" +
			"date_of_birth: '2019-12-27'\n" +
			"email: dev@stainless.com\n" +
			"emotional_stats:\n" +
			"  curiosity: 99\n" +
			"  empathy: 100\n" +
			"  optimism: 95\n" +
			"  resilience: 90\n" +
			"  vulnerability: 80\n" +
			"growth_arcs:\n" +
			"  - breakthrough: breakthrough\n" +
			"    challenge: challenge\n" +
			"    ending_point: ending_point\n" +
			"    season: 1\n" +
			"    starting_point: starting_point\n" +
			"height_meters: 1\n" +
			"name: x\n" +
			"personality_traits:\n" +
			"  - string\n" +
			"profile_image_url: https://example.com\n" +
			"role: coach\n" +
			"salary_gbp: 0\n" +
			"signature_quotes:\n" +
			"  - string\n" +
			"team_id: team_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "characters", "update",
			"--api-key", "string",
			"--character-id", "character_id",
		)
	})
}

func TestCharactersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "characters", "list",
			"--api-key", "string",
			"--limit", "10",
			"--min-optimism", "0",
			"--role", "coach",
			"--skip", "0",
			"--team-id", "team_id",
		)
	})
}

func TestCharactersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "characters", "delete",
			"--api-key", "string",
			"--character-id", "character_id",
		)
	})
}

func TestCharactersGetQuotes(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "characters", "get-quotes",
			"--api-key", "string",
			"--character-id", "character_id",
		)
	})
}
