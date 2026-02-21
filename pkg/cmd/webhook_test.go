// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestWebhooksCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "create",
		"--url", "https://example.com/webhooks",
		"--description", "Production webhook for match notifications",
		"--event-type", "match.completed",
		"--event-type", "team_member.transferred",
	)
}

func TestWebhooksRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "retrieve",
		"--webhook-id", "webhook_id",
	)
}

func TestWebhooksList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "list",
	)
}

func TestWebhooksDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "delete",
		"--webhook-id", "webhook_id",
	)
}

func TestWebhooksTriggerEvent(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "trigger-event",
		"--event-type", "match.completed",
		"--payload", "{data: {away_score: 0, away_team_id: away_team_id, completed_at: '2019-12-27T18:11:19.117Z', home_score: 0, home_team_id: home_team_id, match_id: match_id, match_type: league, result: home_win, ted_post_match_quote: ted_post_match_quote, lesson_learned: lesson_learned, man_of_the_match: man_of_the_match}, event_type: match.completed}",
	)
}
