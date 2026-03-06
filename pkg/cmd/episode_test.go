// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestEpisodesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "episodes", "create",
			"--api-key", "string",
			"--air-date", "'2020-10-02'",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"air_date: '2020-10-02'\n" +
			"character_focus:\n" +
			"  - ted-lasso\n" +
			"  - coach-beard\n" +
			"  - higgins\n" +
			"  - nate\n" +
			"director: MJ Delaney\n" +
			"episode_number: 8\n" +
			"main_theme: The power of vulnerability and male friendship\n" +
			"runtime_minutes: 29\n" +
			"season: 1\n" +
			"synopsis: >-\n" +
			"  Ted creates a support group for the coaching staff while Rebecca faces a\n" +
			"  difficult decision about her future.\n" +
			"ted_wisdom: 'There''s two buttons I never like to hit: that''s panic and snooze.'\n" +
			"title: The Diamond Dogs\n" +
			"writer: Jason Sudeikis, Brendan Hunt, Joe Kelly\n" +
			"biscuits_with_boss_moment: Ted and Rebecca have an honest conversation about trust.\n" +
			"memorable_moments:\n" +
			"  - First Diamond Dogs meeting\n" +
			"  - The famous dart scene with Rupert\n" +
			"  - Be curious, not judgmental speech\n" +
			"us_viewers_millions: 1.42\n" +
			"viewer_rating: 9.1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "episodes", "create",
			"--api-key", "string",
		)
	})
}

func TestEpisodesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "episodes", "retrieve",
			"--api-key", "string",
			"--episode-id", "episode_id",
		)
	})
}

func TestEpisodesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "episodes", "update",
			"--api-key", "string",
			"--episode-id", "episode_id",
			"--air-date", "'2019-12-27'",
			"--biscuits-with-boss-moment", "biscuits_with_boss_moment",
			"--character-focus", "[string]",
			"--director", "director",
			"--episode-number", "1",
			"--main-theme", "main_theme",
			"--memorable-moment", "[string]",
			"--runtime-minutes", "20",
			"--season", "1",
			"--synopsis", "synopsis",
			"--ted-wisdom", "ted_wisdom",
			"--title", "x",
			"--us-viewers-millions", "0",
			"--viewer-rating", "0",
			"--writer", "writer",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"air_date: '2019-12-27'\n" +
			"biscuits_with_boss_moment: biscuits_with_boss_moment\n" +
			"character_focus:\n" +
			"  - string\n" +
			"director: director\n" +
			"episode_number: 1\n" +
			"main_theme: main_theme\n" +
			"memorable_moments:\n" +
			"  - string\n" +
			"runtime_minutes: 20\n" +
			"season: 1\n" +
			"synopsis: synopsis\n" +
			"ted_wisdom: ted_wisdom\n" +
			"title: x\n" +
			"us_viewers_millions: 0\n" +
			"viewer_rating: 0\n" +
			"writer: writer\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "episodes", "update",
			"--api-key", "string",
			"--episode-id", "episode_id",
		)
	})
}

func TestEpisodesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "episodes", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--character-focus", "character_focus",
			"--limit", "10",
			"--season", "1",
			"--skip", "0",
		)
	})
}

func TestEpisodesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "episodes", "delete",
			"--api-key", "string",
			"--episode-id", "episode_id",
		)
	})
}

func TestEpisodesGetWisdom(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "episodes", "get-wisdom",
			"--api-key", "string",
			"--episode-id", "episode_id",
		)
	})
}
