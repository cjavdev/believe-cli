// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestEpisodesCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"episodes", "create",
		"--air-date", "2020-10-02",
		"--character-focus", "ted-lasso",
		"--character-focus", "coach-beard",
		"--character-focus", "higgins",
		"--character-focus", "nate",
		"--director", "MJ Delaney",
		"--episode-number", "8",
		"--main-theme", "The power of vulnerability and male friendship",
		"--runtime-minutes", "29",
		"--season", "1",
		"--synopsis", "Ted creates a support group for the coaching staff while Rebecca faces a difficult decision about her future.",
		"--ted-wisdom", "There's two buttons I never like to hit: that's panic and snooze.",
		"--title", "The Diamond Dogs",
		"--writer", "Jason Sudeikis, Brendan Hunt, Joe Kelly",
		"--biscuits-with-boss-moment", "Ted and Rebecca have an honest conversation about trust.",
		"--memorable-moment", "First Diamond Dogs meeting",
		"--memorable-moment", "The famous dart scene with Rupert",
		"--memorable-moment", "Be curious, not judgmental speech",
		"--us-viewers-millions", "1.42",
		"--viewer-rating", "9.1",
	)
}

func TestEpisodesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"episodes", "retrieve",
		"--episode-id", "episode_id",
	)
}

func TestEpisodesUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"episodes", "update",
		"--episode-id", "episode_id",
		"--air-date", "2019-12-27",
		"--biscuits-with-boss-moment", "biscuits_with_boss_moment",
		"--character-focus", "string",
		"--director", "director",
		"--episode-number", "1",
		"--main-theme", "main_theme",
		"--memorable-moment", "string",
		"--runtime-minutes", "20",
		"--season", "1",
		"--synopsis", "synopsis",
		"--ted-wisdom", "ted_wisdom",
		"--title", "x",
		"--us-viewers-millions", "0",
		"--viewer-rating", "0",
		"--writer", "writer",
	)
}

func TestEpisodesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"episodes", "list",
		"--character-focus", "character_focus",
		"--limit", "10",
		"--season", "1",
		"--skip", "0",
	)
}

func TestEpisodesDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"episodes", "delete",
		"--episode-id", "episode_id",
	)
}

func TestEpisodesGetWisdom(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"episodes", "get-wisdom",
		"--episode-id", "episode_id",
	)
}

func TestEpisodesListBySeason(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"episodes", "list-by-season",
		"--season-number", "0",
		"--limit", "10",
		"--skip", "0",
	)
}
