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

var webhooksCreate = cli.Command{
	Name:    "create",
	Usage:   "Register a new webhook endpoint to receive event notifications.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "The URL to send webhook events to",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "Optional description for this webhook",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "event-type",
			Usage:    "List of event types to subscribe to. If not provided, subscribes to all events.",
			BodyPath: "event_types",
		},
	},
	Action:          handleWebhooksCreate,
	HideHelpCommand: true,
}

var webhooksRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get details of a specific webhook endpoint.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "webhook-id",
			Required: true,
		},
	},
	Action:          handleWebhooksRetrieve,
	HideHelpCommand: true,
}

var webhooksList = cli.Command{
	Name:            "list",
	Usage:           "Get a list of all registered webhook endpoints.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleWebhooksList,
	HideHelpCommand: true,
}

var webhooksDelete = cli.Command{
	Name:    "delete",
	Usage:   "Unregister a webhook endpoint. It will no longer receive events.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "webhook-id",
			Required: true,
		},
	},
	Action:          handleWebhooksDelete,
	HideHelpCommand: true,
}

var webhooksTriggerEvent = cli.Command{
	Name:    "trigger-event",
	Usage:   "Trigger a webhook event and deliver it to all subscribed endpoints.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "event-type",
			Usage:    "The type of event to trigger",
			Required: true,
			BodyPath: "event_type",
		},
		&requestflag.Flag[any]{
			Name:     "payload",
			Usage:    "Optional event payload. If not provided, a sample payload will be generated.",
			BodyPath: "payload",
		},
	},
	Action:          handleWebhooksTriggerEvent,
	HideHelpCommand: true,
}

func handleWebhooksCreate(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.WebhookNewParams{}

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
	_, err = client.Webhooks.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "webhooks create", obj, format, transform)
}

func handleWebhooksRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("webhook-id") && len(unusedArgs) > 0 {
		cmd.Set("webhook-id", unusedArgs[0])
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
	_, err = client.Webhooks.Get(ctx, cmd.Value("webhook-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "webhooks retrieve", obj, format, transform)
}

func handleWebhooksList(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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
	_, err = client.Webhooks.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "webhooks list", obj, format, transform)
}

func handleWebhooksDelete(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("webhook-id") && len(unusedArgs) > 0 {
		cmd.Set("webhook-id", unusedArgs[0])
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
	_, err = client.Webhooks.Delete(ctx, cmd.Value("webhook-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "webhooks delete", obj, format, transform)
}

func handleWebhooksTriggerEvent(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.WebhookTriggerEventParams{}

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
	_, err = client.Webhooks.TriggerEvent(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "webhooks trigger-event", obj, format, transform)
}
