// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestQuotesCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"quotes", "create",
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
}

func TestQuotesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"quotes", "retrieve",
		"--quote-id", "quote_id",
	)
}

func TestQuotesUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"quotes", "update",
		"--quote-id", "quote_id",
		"--character-id", "character_id",
		"--context", "context",
		"--episode-id", "episode_id",
		"--is-funny=true",
		"--is-inspirational=true",
		"--moment-type", "halftime_speech",
		"--popularity-score", "0",
		"--secondary-theme", "belief",
		"--text", "x",
		"--theme", "belief",
		"--times-shared", "0",
	)
}

func TestQuotesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"quotes", "list",
		"--character-id", "character_id",
		"--funny=true",
		"--inspirational=true",
		"--limit", "10",
		"--moment-type", "halftime_speech",
		"--skip", "0",
		"--theme", "belief",
	)
}

func TestQuotesDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"quotes", "delete",
		"--quote-id", "quote_id",
	)
}

func TestQuotesGetRandom(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"quotes", "get-random",
		"--character-id", "character_id",
		"--inspirational=true",
		"--theme", "belief",
	)
}

func TestQuotesListByCharacter(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"quotes", "list-by-character",
		"--character-id", "character_id",
		"--limit", "10",
		"--skip", "0",
	)
}

func TestQuotesListByTheme(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"quotes", "list-by-theme",
		"--theme", "belief",
		"--limit", "10",
		"--skip", "0",
	)
}
