// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/believe-cli/internal/apiquery"
	"github.com/stainless-sdks/believe-cli/internal/requestflag"
	"github.com/stainless-sdks/believe-go"
	"github.com/stainless-sdks/believe-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var coachingPrinciplesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get details about a specific coaching principle.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "principle-id",
			Required: true,
		},
	},
	Action:          handleCoachingPrinciplesRetrieve,
	HideHelpCommand: true,
}

var coachingPrinciplesList = cli.Command{
	Name:    "list",
	Usage:   "Get a paginated list of Ted Lasso's core coaching principles and philosophy.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return (max: 100)",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "skip",
			Usage:     "Number of items to skip (offset)",
			Default:   0,
			QueryPath: "skip",
		},
	},
	Action:          handleCoachingPrinciplesList,
	HideHelpCommand: true,
}

var coachingPrinciplesGetRandom = cli.Command{
	Name:            "get-random",
	Usage:           "Get a random coaching principle to inspire your day.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleCoachingPrinciplesGetRandom,
	HideHelpCommand: true,
}

func handleCoachingPrinciplesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("principle-id") && len(unusedArgs) > 0 {
		cmd.Set("principle-id", unusedArgs[0])
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
	_, err = client.Coaching.Principles.Get(ctx, cmd.Value("principle-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "coaching:principles retrieve", obj, format, transform)
}

func handleCoachingPrinciplesList(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.CoachingPrincipleListParams{}

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
		_, err = client.Coaching.Principles.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "coaching:principles list", obj, format, transform)
	} else {
		iter := client.Coaching.Principles.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "coaching:principles list", iter, format, transform)
	}
}

func handleCoachingPrinciplesGetRandom(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Coaching.Principles.GetRandom(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "coaching:principles get-random", obj, format, transform)
}
