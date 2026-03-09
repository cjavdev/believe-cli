// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/cjavdev/believe-cli/internal/apiquery"
	"github.com/cjavdev/believe-cli/internal/requestflag"
	"github.com/cjavdev/believe-go"
	"github.com/cjavdev/believe-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var clientTicketSalesCreate = cli.Command{
	Name:    "create",
	Usage:   "Record a new ticket sale.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "buyer-name",
			Usage:    "Name of the ticket buyer",
			Required: true,
			BodyPath: "buyer_name",
		},
		&requestflag.Flag[string]{
			Name:     "currency",
			Usage:    "Currency code (GBP, USD, or EUR)",
			Required: true,
			BodyPath: "currency",
		},
		&requestflag.Flag[string]{
			Name:     "discount",
			Usage:    "Discount amount applied from coupon",
			Required: true,
			BodyPath: "discount",
		},
		&requestflag.Flag[string]{
			Name:     "match-id",
			Usage:    "ID of the match",
			Required: true,
			BodyPath: "match_id",
		},
		&requestflag.Flag[string]{
			Name:     "purchase-method",
			Usage:    "How the ticket was purchased",
			Required: true,
			BodyPath: "purchase_method",
		},
		&requestflag.Flag[int64]{
			Name:     "quantity",
			Usage:    "Number of tickets purchased",
			Required: true,
			BodyPath: "quantity",
		},
		&requestflag.Flag[string]{
			Name:     "subtotal",
			Usage:    "Subtotal before discount and tax (unit_price * quantity)",
			Required: true,
			BodyPath: "subtotal",
		},
		&requestflag.Flag[string]{
			Name:     "tax",
			Usage:    "Tax amount (20% UK VAT on discounted subtotal)",
			Required: true,
			BodyPath: "tax",
		},
		&requestflag.Flag[string]{
			Name:     "total",
			Usage:    "Final total (subtotal - discount + tax)",
			Required: true,
			BodyPath: "total",
		},
		&requestflag.Flag[string]{
			Name:     "unit-price",
			Usage:    "Price per ticket (decimal string)",
			Required: true,
			BodyPath: "unit_price",
		},
		&requestflag.Flag[any]{
			Name:     "buyer-email",
			Usage:    "Email of the ticket buyer",
			BodyPath: "buyer_email",
		},
		&requestflag.Flag[any]{
			Name:     "coupon-code",
			Usage:    "Coupon code applied, if any",
			BodyPath: "coupon_code",
		},
	},
	Action:          handleClientTicketSalesCreate,
	HideHelpCommand: true,
}

var clientTicketSalesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve detailed information about a specific ticket sale.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "ticket-sale-id",
			Required: true,
		},
	},
	Action:          handleClientTicketSalesRetrieve,
	HideHelpCommand: true,
}

var clientTicketSalesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update specific fields of an existing ticket sale.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "ticket-sale-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "buyer-email",
			BodyPath: "buyer_email",
		},
		&requestflag.Flag[any]{
			Name:     "buyer-name",
			BodyPath: "buyer_name",
		},
		&requestflag.Flag[any]{
			Name:     "coupon-code",
			BodyPath: "coupon_code",
		},
		&requestflag.Flag[any]{
			Name:     "currency",
			BodyPath: "currency",
		},
		&requestflag.Flag[any]{
			Name:     "discount",
			BodyPath: "discount",
		},
		&requestflag.Flag[any]{
			Name:     "match-id",
			BodyPath: "match_id",
		},
		&requestflag.Flag[any]{
			Name:     "purchase-method",
			Usage:    "How the ticket was purchased.",
			BodyPath: "purchase_method",
		},
		&requestflag.Flag[any]{
			Name:     "quantity",
			BodyPath: "quantity",
		},
		&requestflag.Flag[any]{
			Name:     "subtotal",
			BodyPath: "subtotal",
		},
		&requestflag.Flag[any]{
			Name:     "tax",
			BodyPath: "tax",
		},
		&requestflag.Flag[any]{
			Name:     "total",
			BodyPath: "total",
		},
		&requestflag.Flag[any]{
			Name:     "unit-price",
			BodyPath: "unit_price",
		},
	},
	Action:          handleClientTicketSalesUpdate,
	HideHelpCommand: true,
}

var clientTicketSalesList = cli.Command{
	Name:    "list",
	Usage:   "Get a paginated list of all ticket sales with optional filtering. With 300\nrecords, this endpoint is ideal for practicing pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "coupon-code",
			Usage:     "Filter by coupon code (use 'none' for sales without coupons)",
			QueryPath: "coupon_code",
		},
		&requestflag.Flag[any]{
			Name:      "currency",
			Usage:     "Filter by currency (GBP, USD, EUR)",
			QueryPath: "currency",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return (max: 100)",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "match-id",
			Usage:     "Filter by match ID",
			QueryPath: "match_id",
		},
		&requestflag.Flag[any]{
			Name:      "purchase-method",
			Usage:     "Filter by purchase method",
			QueryPath: "purchase_method",
		},
		&requestflag.Flag[int64]{
			Name:      "skip",
			Usage:     "Number of items to skip (offset)",
			Default:   0,
			QueryPath: "skip",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleClientTicketSalesList,
	HideHelpCommand: true,
}

var clientTicketSalesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Remove a ticket sale from the database.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "ticket-sale-id",
			Required: true,
		},
	},
	Action:          handleClientTicketSalesDelete,
	HideHelpCommand: true,
}

func handleClientTicketSalesCreate(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.ClientTicketSaleNewParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Client.TicketSales.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "client:ticket-sales create", obj, format, transform)
}

func handleClientTicketSalesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ticket-sale-id") && len(unusedArgs) > 0 {
		cmd.Set("ticket-sale-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Client.TicketSales.Get(ctx, cmd.Value("ticket-sale-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "client:ticket-sales retrieve", obj, format, transform)
}

func handleClientTicketSalesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ticket-sale-id") && len(unusedArgs) > 0 {
		cmd.Set("ticket-sale-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.ClientTicketSaleUpdateParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Client.TicketSales.Update(
		ctx,
		cmd.Value("ticket-sale-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "client:ticket-sales update", obj, format, transform)
}

func handleClientTicketSalesList(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.ClientTicketSaleListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Client.TicketSales.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "client:ticket-sales list", obj, format, transform)
	} else {
		iter := client.Client.TicketSales.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "client:ticket-sales list", iter, format, transform, maxItems)
	}
}

func handleClientTicketSalesDelete(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ticket-sale-id") && len(unusedArgs) > 0 {
		cmd.Set("ticket-sale-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	return client.Client.TicketSales.Delete(ctx, cmd.Value("ticket-sale-id").(string), options...)
}
