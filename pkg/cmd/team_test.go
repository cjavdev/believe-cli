// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
	"github.com/cjavdev/believe-cli/internal/requestflag"
)

func TestTeamsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "teams", "create",
			"--api-key", "string",
			"--culture-score", "70",
			"--founded-year", "1895",
			"--league", "Premier League",
			"--name", "West Ham United",
			"--stadium", "London Stadium",
			"--values", "{primary_value: Pride, secondary_values: [History, Community, Passion], team_motto: Forever Blowing Bubbles}",
			"--annual-budget-gbp", "'150000000.00'",
			"--average-attendance", "59988",
			"--contact-email", "info@westhamunited.co.uk",
			"--is-active=true",
			"--nickname", "The Hammers",
			"--primary-color", "#7A263A",
			"--rival-team", "afc-richmond",
			"--rival-team", "tottenham",
			"--secondary-color", "#1BB1E7",
			"--stadium-location", "{latitude: 51.5387, longitude: -0.0166}",
			"--website", "https://www.whufc.com",
			"--win-percentage", "52.3",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(teamsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "teams", "create",
			"--api-key", "string",
			"--culture-score", "70",
			"--founded-year", "1895",
			"--league", "Premier League",
			"--name", "West Ham United",
			"--stadium", "London Stadium",
			"--values.primary-value", "Pride",
			"--values.secondary-values", "[History, Community, Passion]",
			"--values.team-motto", "Forever Blowing Bubbles",
			"--annual-budget-gbp", "'150000000.00'",
			"--average-attendance", "59988",
			"--contact-email", "info@westhamunited.co.uk",
			"--is-active=true",
			"--nickname", "The Hammers",
			"--primary-color", "#7A263A",
			"--rival-team", "afc-richmond",
			"--rival-team", "tottenham",
			"--secondary-color", "#1BB1E7",
			"--stadium-location.latitude", "51.5387",
			"--stadium-location.longitude", "-0.0166",
			"--website", "https://www.whufc.com",
			"--win-percentage", "52.3",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"culture_score: 70\n" +
			"founded_year: 1895\n" +
			"league: Premier League\n" +
			"name: West Ham United\n" +
			"stadium: London Stadium\n" +
			"values:\n" +
			"  primary_value: Pride\n" +
			"  secondary_values:\n" +
			"    - History\n" +
			"    - Community\n" +
			"    - Passion\n" +
			"  team_motto: Forever Blowing Bubbles\n" +
			"annual_budget_gbp: '150000000.00'\n" +
			"average_attendance: 59988\n" +
			"contact_email: info@westhamunited.co.uk\n" +
			"is_active: true\n" +
			"nickname: The Hammers\n" +
			"primary_color: '#7A263A'\n" +
			"rival_teams:\n" +
			"  - afc-richmond\n" +
			"  - tottenham\n" +
			"secondary_color: '#1BB1E7'\n" +
			"stadium_location:\n" +
			"  latitude: 51.5387\n" +
			"  longitude: -0.0166\n" +
			"website: https://www.whufc.com\n" +
			"win_percentage: 52.3\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "teams", "create",
			"--api-key", "string",
		)
	})
}

func TestTeamsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "teams", "retrieve",
			"--api-key", "string",
			"--team-id", "team_id",
		)
	})
}

func TestTeamsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "teams", "update",
			"--api-key", "string",
			"--team-id", "team_id",
			"--annual-budget-gbp", "0",
			"--average-attendance", "0",
			"--contact-email", "dev@stainless.com",
			"--culture-score", "0",
			"--founded-year", "1800",
			"--is-active=true",
			"--league", "Premier League",
			"--name", "x",
			"--nickname", "nickname",
			"--primary-color", "primary_color",
			"--rival-team", "[string]",
			"--secondary-color", "secondary_color",
			"--stadium", "stadium",
			"--stadium-location", "{latitude: 51.4816, longitude: -0.191}",
			"--values", "{primary_value: Believe, secondary_values: [Family, Resilience, Joy], team_motto: Football is life!}",
			"--website", "https://example.com",
			"--win-percentage", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(teamsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "teams", "update",
			"--api-key", "string",
			"--team-id", "team_id",
			"--annual-budget-gbp", "0",
			"--average-attendance", "0",
			"--contact-email", "dev@stainless.com",
			"--culture-score", "0",
			"--founded-year", "1800",
			"--is-active=true",
			"--league", "Premier League",
			"--name", "x",
			"--nickname", "nickname",
			"--primary-color", "primary_color",
			"--rival-team", "[string]",
			"--secondary-color", "secondary_color",
			"--stadium", "stadium",
			"--stadium-location.latitude", "51.4816",
			"--stadium-location.longitude", "-0.191",
			"--values.primary-value", "Believe",
			"--values.secondary-values", "[Family, Resilience, Joy]",
			"--values.team-motto", "Football is life!",
			"--website", "https://example.com",
			"--win-percentage", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"annual_budget_gbp: 0\n" +
			"average_attendance: 0\n" +
			"contact_email: dev@stainless.com\n" +
			"culture_score: 0\n" +
			"founded_year: 1800\n" +
			"is_active: true\n" +
			"league: Premier League\n" +
			"name: x\n" +
			"nickname: nickname\n" +
			"primary_color: primary_color\n" +
			"rival_teams:\n" +
			"  - string\n" +
			"secondary_color: secondary_color\n" +
			"stadium: stadium\n" +
			"stadium_location:\n" +
			"  latitude: 51.4816\n" +
			"  longitude: -0.191\n" +
			"values:\n" +
			"  primary_value: Believe\n" +
			"  secondary_values:\n" +
			"    - Family\n" +
			"    - Resilience\n" +
			"    - Joy\n" +
			"  team_motto: Football is life!\n" +
			"website: https://example.com\n" +
			"win_percentage: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "teams", "update",
			"--api-key", "string",
			"--team-id", "team_id",
		)
	})
}

func TestTeamsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "teams", "list",
			"--api-key", "string",
			"--league", "Premier League",
			"--limit", "10",
			"--min-culture-score", "0",
			"--skip", "0",
		)
	})
}

func TestTeamsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "teams", "delete",
			"--api-key", "string",
			"--team-id", "team_id",
		)
	})
}

func TestTeamsGetCulture(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "teams", "get-culture",
			"--api-key", "string",
			"--team-id", "team_id",
		)
	})
}

func TestTeamsGetRivals(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "teams", "get-rivals",
			"--api-key", "string",
			"--team-id", "team_id",
		)
	})
}

func TestTeamsListLogos(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "teams", "list-logos",
			"--api-key", "string",
			"--team-id", "team_id",
		)
	})
}
