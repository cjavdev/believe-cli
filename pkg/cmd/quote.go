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

var quotesCreate = cli.Command{
	Name:    "create",
	Usage:   "Add a new memorable quote to the collection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "character-id",
			Usage:    "ID of the character who said it",
			Required: true,
			BodyPath: "character_id",
		},
		&requestflag.Flag[string]{
			Name:     "context",
			Usage:    "Context in which the quote was said",
			Required: true,
			BodyPath: "context",
		},
		&requestflag.Flag[string]{
			Name:     "moment-type",
			Usage:    "Types of moments when quotes occur.",
			Required: true,
			BodyPath: "moment_type",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "The quote text",
			Required: true,
			BodyPath: "text",
		},
		&requestflag.Flag[string]{
			Name:     "theme",
			Usage:    "Themes that quotes can be categorized under.",
			Required: true,
			BodyPath: "theme",
		},
		&requestflag.Flag[any]{
			Name:     "episode-id",
			Usage:    "Episode where the quote appears",
			BodyPath: "episode_id",
		},
		&requestflag.Flag[bool]{
			Name:     "is-funny",
			Usage:    "Whether this quote is humorous",
			Default:  false,
			BodyPath: "is_funny",
		},
		&requestflag.Flag[bool]{
			Name:     "is-inspirational",
			Usage:    "Whether this quote is inspirational",
			Default:  true,
			BodyPath: "is_inspirational",
		},
		&requestflag.Flag[any]{
			Name:     "popularity-score",
			Usage:    "Popularity/virality score (0-100)",
			BodyPath: "popularity_score",
		},
		&requestflag.Flag[[]string]{
			Name:     "secondary-theme",
			Usage:    "Additional themes",
			BodyPath: "secondary_themes",
		},
		&requestflag.Flag[any]{
			Name:     "times-shared",
			Usage:    "Number of times shared on social media",
			BodyPath: "times_shared",
		},
	},
	Action:          handleQuotesCreate,
	HideHelpCommand: true,
}

var quotesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a specific quote by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "quote-id",
			Required: true,
		},
	},
	Action:          handleQuotesRetrieve,
	HideHelpCommand: true,
}

var quotesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update specific fields of an existing quote.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "quote-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "character-id",
			BodyPath: "character_id",
		},
		&requestflag.Flag[any]{
			Name:     "context",
			BodyPath: "context",
		},
		&requestflag.Flag[any]{
			Name:     "episode-id",
			BodyPath: "episode_id",
		},
		&requestflag.Flag[any]{
			Name:     "is-funny",
			BodyPath: "is_funny",
		},
		&requestflag.Flag[any]{
			Name:     "is-inspirational",
			BodyPath: "is_inspirational",
		},
		&requestflag.Flag[string]{
			Name:     "moment-type",
			Usage:    "Types of moments when quotes occur.",
			BodyPath: "moment_type",
		},
		&requestflag.Flag[any]{
			Name:     "popularity-score",
			BodyPath: "popularity_score",
		},
		&requestflag.Flag[any]{
			Name:     "secondary-theme",
			BodyPath: "secondary_themes",
		},
		&requestflag.Flag[any]{
			Name:     "text",
			BodyPath: "text",
		},
		&requestflag.Flag[string]{
			Name:     "theme",
			Usage:    "Themes that quotes can be categorized under.",
			BodyPath: "theme",
		},
		&requestflag.Flag[any]{
			Name:     "times-shared",
			BodyPath: "times_shared",
		},
	},
	Action:          handleQuotesUpdate,
	HideHelpCommand: true,
}

var quotesList = cli.Command{
	Name:    "list",
	Usage:   "Get a paginated list of all memorable Ted Lasso quotes with optional filtering.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "character-id",
			Usage:     "Filter by character",
			QueryPath: "character_id",
		},
		&requestflag.Flag[any]{
			Name:      "funny",
			Usage:     "Filter funny quotes",
			QueryPath: "funny",
		},
		&requestflag.Flag[any]{
			Name:      "inspirational",
			Usage:     "Filter inspirational quotes",
			QueryPath: "inspirational",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return (max: 100)",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "moment-type",
			Usage:     "Types of moments when quotes occur.",
			QueryPath: "moment_type",
		},
		&requestflag.Flag[int64]{
			Name:      "skip",
			Usage:     "Number of items to skip (offset)",
			Default:   0,
			QueryPath: "skip",
		},
		&requestflag.Flag[string]{
			Name:      "theme",
			Usage:     "Themes that quotes can be categorized under.",
			QueryPath: "theme",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleQuotesList,
	HideHelpCommand: true,
}

var quotesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Remove a quote from the collection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "quote-id",
			Required: true,
		},
	},
	Action:          handleQuotesDelete,
	HideHelpCommand: true,
}

var quotesGetRandom = cli.Command{
	Name:    "get-random",
	Usage:   "Get a random Ted Lasso quote, optionally filtered.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "character-id",
			Usage:     "Filter by character",
			QueryPath: "character_id",
		},
		&requestflag.Flag[any]{
			Name:      "inspirational",
			Usage:     "Filter inspirational quotes",
			QueryPath: "inspirational",
		},
		&requestflag.Flag[string]{
			Name:      "theme",
			Usage:     "Themes that quotes can be categorized under.",
			QueryPath: "theme",
		},
	},
	Action:          handleQuotesGetRandom,
	HideHelpCommand: true,
}

var quotesListByCharacter = cli.Command{
	Name:    "list-by-character",
	Usage:   "Get a paginated list of quotes from a specific character.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "character-id",
			Required: true,
		},
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleQuotesListByCharacter,
	HideHelpCommand: true,
}

var quotesListByTheme = cli.Command{
	Name:    "list-by-theme",
	Usage:   "Get a paginated list of quotes related to a specific theme.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "theme",
			Usage:    "Themes that quotes can be categorized under.",
			Required: true,
		},
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleQuotesListByTheme,
	HideHelpCommand: true,
}

func handleQuotesCreate(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.QuoteNewParams{}

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
	_, err = client.Quotes.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "quotes create", obj, format, transform)
}

func handleQuotesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("quote-id") && len(unusedArgs) > 0 {
		cmd.Set("quote-id", unusedArgs[0])
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
	_, err = client.Quotes.Get(ctx, cmd.Value("quote-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "quotes retrieve", obj, format, transform)
}

func handleQuotesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("quote-id") && len(unusedArgs) > 0 {
		cmd.Set("quote-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.QuoteUpdateParams{}

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
	_, err = client.Quotes.Update(
		ctx,
		cmd.Value("quote-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "quotes update", obj, format, transform)
}

func handleQuotesList(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.QuoteListParams{}

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
		_, err = client.Quotes.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "quotes list", obj, format, transform)
	} else {
		iter := client.Quotes.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "quotes list", iter, format, transform, maxItems)
	}
}

func handleQuotesDelete(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("quote-id") && len(unusedArgs) > 0 {
		cmd.Set("quote-id", unusedArgs[0])
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

	return client.Quotes.Delete(ctx, cmd.Value("quote-id").(string), options...)
}

func handleQuotesGetRandom(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.QuoteGetRandomParams{}

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
	_, err = client.Quotes.GetRandom(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "quotes get-random", obj, format, transform)
}

func handleQuotesListByCharacter(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("character-id") && len(unusedArgs) > 0 {
		cmd.Set("character-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.QuoteListByCharacterParams{}

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
		_, err = client.Quotes.ListByCharacter(
			ctx,
			cmd.Value("character-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "quotes list-by-character", obj, format, transform)
	} else {
		iter := client.Quotes.ListByCharacterAutoPaging(
			ctx,
			cmd.Value("character-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "quotes list-by-character", iter, format, transform, maxItems)
	}
}

func handleQuotesListByTheme(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("theme") && len(unusedArgs) > 0 {
		cmd.Set("theme", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.QuoteListByThemeParams{}

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
		_, err = client.Quotes.ListByTheme(
			ctx,
			believe.QuoteTheme(cmd.Value("theme").(string)),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "quotes list-by-theme", obj, format, transform)
	} else {
		iter := client.Quotes.ListByThemeAutoPaging(
			ctx,
			believe.QuoteTheme(cmd.Value("theme").(string)),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "quotes list-by-theme", iter, format, transform, maxItems)
	}
}
