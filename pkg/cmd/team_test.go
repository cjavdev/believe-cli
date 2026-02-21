// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
	"github.com/cjavdev/believe-cli/internal/requestflag"
)

func TestTeamsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"teams", "create",
		"--culture-score", "70",
		"--founded-year", "1895",
		"--league", "Premier League",
		"--name", "West Ham United",
		"--stadium", "London Stadium",
		"--values", "{primary_value: Pride, secondary_values: [History, Community, Passion], team_motto: Forever Blowing Bubbles}",
		"--annual-budget-gbp", "150000000.00",
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(teamsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"teams", "create",
		"--culture-score", "70",
		"--founded-year", "1895",
		"--league", "Premier League",
		"--name", "West Ham United",
		"--stadium", "London Stadium",
		"--values.primary-value", "Pride",
		"--values.secondary-values", "[History, Community, Passion]",
		"--values.team-motto", "Forever Blowing Bubbles",
		"--annual-budget-gbp", "150000000.00",
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
}

func TestTeamsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"teams", "retrieve",
		"--team-id", "team_id",
	)
}

func TestTeamsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"teams", "update",
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
		"--rival-team", "string",
		"--secondary-color", "secondary_color",
		"--stadium", "stadium",
		"--stadium-location", "{latitude: 51.4816, longitude: -0.191}",
		"--values", "{primary_value: Believe, secondary_values: [Family, Resilience, Joy], team_motto: Football is life!}",
		"--website", "https://example.com",
		"--win-percentage", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(teamsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"teams", "update",
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
		"--rival-team", "string",
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
}

func TestTeamsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"teams", "list",
		"--league", "Premier League",
		"--limit", "10",
		"--min-culture-score", "0",
		"--skip", "0",
	)
}

func TestTeamsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"teams", "delete",
		"--team-id", "team_id",
	)
}

func TestTeamsGetCulture(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"teams", "get-culture",
		"--team-id", "team_id",
	)
}

func TestTeamsGetRivals(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"teams", "get-rivals",
		"--team-id", "team_id",
	)
}

func TestTeamsListLogos(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"teams", "list-logos",
		"--team-id", "team_id",
	)
}
