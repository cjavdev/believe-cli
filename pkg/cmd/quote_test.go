// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestQuotesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "quotes", "create",
			"--api-key", "string",
			"--character-id", "ted-lasso",
			"--context", "Ted's first team meeting, revealing his coaching philosophy",
			"--moment-type", "locker_room",
			"--text", "I believe in believe.",
			"--theme", "belief",
			"--episode-id", "s01e01",
			"--is-funny=false",
			"--is-inspirational=true",
			"--popularity-score", "98.5",
			"--secondary-theme", "leadership",
			"--secondary-theme", "teamwork",
			"--times-shared", "250000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"character_id: ted-lasso\n" +
			"context: Ted's first team meeting, revealing his coaching philosophy\n" +
			"moment_type: locker_room\n" +
			"text: I believe in believe.\n" +
			"theme: belief\n" +
			"episode_id: s01e01\n" +
			"is_funny: false\n" +
			"is_inspirational: true\n" +
			"popularity_score: 98.5\n" +
			"secondary_themes:\n" +
			"  - leadership\n" +
			"  - teamwork\n" +
			"times_shared: 250000\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "quotes", "create",
			"--api-key", "string",
		)
	})
}

func TestQuotesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "quotes", "retrieve",
			"--api-key", "string",
			"--quote-id", "quote_id",
		)
	})
}

func TestQuotesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "quotes", "update",
			"--api-key", "string",
			"--quote-id", "quote_id",
			"--character-id", "character_id",
			"--context", "context",
			"--episode-id", "episode_id",
			"--is-funny=true",
			"--is-inspirational=true",
			"--moment-type", "halftime_speech",
			"--popularity-score", "0",
			"--secondary-theme", "[belief]",
			"--text", "x",
			"--theme", "belief",
			"--times-shared", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"character_id: character_id\n" +
			"context: context\n" +
			"episode_id: episode_id\n" +
			"is_funny: true\n" +
			"is_inspirational: true\n" +
			"moment_type: halftime_speech\n" +
			"popularity_score: 0\n" +
			"secondary_themes:\n" +
			"  - belief\n" +
			"text: x\n" +
			"theme: belief\n" +
			"times_shared: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "quotes", "update",
			"--api-key", "string",
			"--quote-id", "quote_id",
		)
	})
}

func TestQuotesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "quotes", "list",
			"--api-key", "string",
			"--character-id", "character_id",
			"--funny=true",
			"--inspirational=true",
			"--limit", "10",
			"--moment-type", "halftime_speech",
			"--skip", "0",
			"--theme", "belief",
		)
	})
}

func TestQuotesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "quotes", "delete",
			"--api-key", "string",
			"--quote-id", "quote_id",
		)
	})
}

func TestQuotesGetRandom(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "quotes", "get-random",
			"--api-key", "string",
			"--character-id", "character_id",
			"--inspirational=true",
			"--theme", "belief",
		)
	})
}

func TestQuotesListByCharacter(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "quotes", "list-by-character",
			"--api-key", "string",
			"--character-id", "character_id",
			"--limit", "10",
			"--skip", "0",
		)
	})
}

func TestQuotesListByTheme(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "quotes", "list-by-theme",
			"--api-key", "string",
			"--theme", "belief",
			"--limit", "10",
			"--skip", "0",
		)
	})
}
