// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cjavdev/believe-cli/internal/mocktest"
)

func TestTicketSalesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ticket-sales", "create",
			"--api-key", "string",
			"--buyer-name", "Mae Green",
			"--currency", "GBP",
			"--discount", "9.00",
			"--match-id", "match-001",
			"--purchase-method", "online",
			"--quantity", "2",
			"--subtotal", "90.00",
			"--tax", "16.20",
			"--total", "97.20",
			"--unit-price", "45.00",
			"--buyer-email", "mae.green@example.com",
			"--coupon-code", "BELIEVE10",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"buyer_name: Mae Green\n" +
			"currency: GBP\n" +
			"discount: '9.00'\n" +
			"match_id: match-001\n" +
			"purchase_method: online\n" +
			"quantity: 2\n" +
			"subtotal: '90.00'\n" +
			"tax: '16.20'\n" +
			"total: '97.20'\n" +
			"unit_price: '45.00'\n" +
			"buyer_email: mae.green@example.com\n" +
			"coupon_code: BELIEVE10\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ticket-sales", "create",
			"--api-key", "string",
		)
	})
}

func TestTicketSalesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ticket-sales", "retrieve",
			"--api-key", "string",
			"--ticket-sale-id", "ticket_sale_id",
		)
	})
}

func TestTicketSalesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ticket-sales", "update",
			"--api-key", "string",
			"--ticket-sale-id", "ticket_sale_id",
			"--buyer-email", "dev@stainless.com",
			"--buyer-name", "buyer_name",
			"--coupon-code", "coupon_code",
			"--currency", "currency",
			"--discount", "discount",
			"--match-id", "match_id",
			"--purchase-method", "online",
			"--quantity", "1",
			"--subtotal", "subtotal",
			"--tax", "tax",
			"--total", "total",
			"--unit-price", "unit_price",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"buyer_email: dev@stainless.com\n" +
			"buyer_name: buyer_name\n" +
			"coupon_code: coupon_code\n" +
			"currency: currency\n" +
			"discount: discount\n" +
			"match_id: match_id\n" +
			"purchase_method: online\n" +
			"quantity: 1\n" +
			"subtotal: subtotal\n" +
			"tax: tax\n" +
			"total: total\n" +
			"unit_price: unit_price\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ticket-sales", "update",
			"--api-key", "string",
			"--ticket-sale-id", "ticket_sale_id",
		)
	})
}

func TestTicketSalesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ticket-sales", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--coupon-code", "coupon_code",
			"--currency", "currency",
			"--limit", "10",
			"--match-id", "match_id",
			"--purchase-method", "online",
			"--skip", "0",
		)
	})
}

func TestTicketSalesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ticket-sales", "delete",
			"--api-key", "string",
			"--ticket-sale-id", "ticket_sale_id",
		)
	})
}
