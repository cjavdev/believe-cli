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

var charactersCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Add a new character to the Ted Lasso universe.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "background",
			Usage:    "Character background and history",
			Required: true,
			BodyPath: "background",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "emotional-stats",
			Usage:    "Emotional intelligence statistics for a character.",
			Required: true,
			BodyPath: "emotional_stats",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Character's full name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "personality-trait",
			Usage:    "Key personality traits",
			Required: true,
			BodyPath: "personality_traits",
		},
		&requestflag.Flag[string]{
			Name:     "role",
			Usage:    "Roles characters can have.",
			Required: true,
			BodyPath: "role",
		},
		&requestflag.Flag[any]{
			Name:     "date-of-birth",
			Usage:    "Character's date of birth",
			BodyPath: "date_of_birth",
		},
		&requestflag.Flag[any]{
			Name:     "email",
			Usage:    "Character's email address",
			BodyPath: "email",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "growth-arc",
			Usage:    "Character development across seasons",
			BodyPath: "growth_arcs",
		},
		&requestflag.Flag[any]{
			Name:     "height-meters",
			Usage:    "Height in meters",
			BodyPath: "height_meters",
		},
		&requestflag.Flag[any]{
			Name:     "profile-image-url",
			Usage:    "URL to character's profile image",
			BodyPath: "profile_image_url",
		},
		&requestflag.Flag[any]{
			Name:     "salary-gbp",
			Usage:    "Annual salary in GBP",
			BodyPath: "salary_gbp",
		},
		&requestflag.Flag[[]string]{
			Name:     "signature-quote",
			Usage:    "Memorable quotes from this character",
			BodyPath: "signature_quotes",
		},
		&requestflag.Flag[any]{
			Name:     "team-id",
			Usage:    "ID of the team they belong to",
			BodyPath: "team_id",
		},
	},
	Action:          handleCharactersCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"emotional-stats": {
		&requestflag.InnerFlag[int64]{
			Name:       "emotional-stats.curiosity",
			Usage:      "Level of curiosity over judgment (0-100)",
			InnerField: "curiosity",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "emotional-stats.empathy",
			Usage:      "Capacity for empathy (0-100)",
			InnerField: "empathy",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "emotional-stats.optimism",
			Usage:      "Level of optimism (0-100)",
			InnerField: "optimism",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "emotional-stats.resilience",
			Usage:      "Bounce-back ability (0-100)",
			InnerField: "resilience",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "emotional-stats.vulnerability",
			Usage:      "Willingness to be vulnerable (0-100)",
			InnerField: "vulnerability",
		},
	},
	"growth-arc": {
		&requestflag.InnerFlag[string]{
			Name:       "growth-arc.breakthrough",
			Usage:      "Key breakthrough moment",
			InnerField: "breakthrough",
		},
		&requestflag.InnerFlag[string]{
			Name:       "growth-arc.challenge",
			Usage:      "Main challenge faced",
			InnerField: "challenge",
		},
		&requestflag.InnerFlag[string]{
			Name:       "growth-arc.ending-point",
			Usage:      "Where the character ends up",
			InnerField: "ending_point",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "growth-arc.season",
			Usage:      "Season number",
			InnerField: "season",
		},
		&requestflag.InnerFlag[string]{
			Name:       "growth-arc.starting-point",
			Usage:      "Where the character starts emotionally",
			InnerField: "starting_point",
		},
	},
})

var charactersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve detailed information about a specific character.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "character-id",
			Required: true,
		},
	},
	Action:          handleCharactersRetrieve,
	HideHelpCommand: true,
}

var charactersUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update specific fields of an existing character.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "character-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "background",
			BodyPath: "background",
		},
		&requestflag.Flag[any]{
			Name:     "date-of-birth",
			BodyPath: "date_of_birth",
		},
		&requestflag.Flag[any]{
			Name:     "email",
			BodyPath: "email",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "emotional-stats",
			Usage:    "Emotional intelligence statistics for a character.",
			BodyPath: "emotional_stats",
		},
		&requestflag.Flag[any]{
			Name:     "growth-arc",
			BodyPath: "growth_arcs",
		},
		&requestflag.Flag[any]{
			Name:     "height-meters",
			BodyPath: "height_meters",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "personality-trait",
			BodyPath: "personality_traits",
		},
		&requestflag.Flag[any]{
			Name:     "profile-image-url",
			BodyPath: "profile_image_url",
		},
		&requestflag.Flag[string]{
			Name:     "role",
			Usage:    "Roles characters can have.",
			BodyPath: "role",
		},
		&requestflag.Flag[any]{
			Name:     "salary-gbp",
			BodyPath: "salary_gbp",
		},
		&requestflag.Flag[any]{
			Name:     "signature-quote",
			BodyPath: "signature_quotes",
		},
		&requestflag.Flag[any]{
			Name:     "team-id",
			BodyPath: "team_id",
		},
	},
	Action:          handleCharactersUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"emotional-stats": {
		&requestflag.InnerFlag[int64]{
			Name:       "emotional-stats.curiosity",
			Usage:      "Level of curiosity over judgment (0-100)",
			InnerField: "curiosity",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "emotional-stats.empathy",
			Usage:      "Capacity for empathy (0-100)",
			InnerField: "empathy",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "emotional-stats.optimism",
			Usage:      "Level of optimism (0-100)",
			InnerField: "optimism",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "emotional-stats.resilience",
			Usage:      "Bounce-back ability (0-100)",
			InnerField: "resilience",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "emotional-stats.vulnerability",
			Usage:      "Willingness to be vulnerable (0-100)",
			InnerField: "vulnerability",
		},
	},
	"growth-arc": {
		&requestflag.InnerFlag[string]{
			Name:       "growth-arc.breakthrough",
			Usage:      "Key breakthrough moment",
			InnerField: "breakthrough",
		},
		&requestflag.InnerFlag[string]{
			Name:       "growth-arc.challenge",
			Usage:      "Main challenge faced",
			InnerField: "challenge",
		},
		&requestflag.InnerFlag[string]{
			Name:       "growth-arc.ending-point",
			Usage:      "Where the character ends up",
			InnerField: "ending_point",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "growth-arc.season",
			Usage:      "Season number",
			InnerField: "season",
		},
		&requestflag.InnerFlag[string]{
			Name:       "growth-arc.starting-point",
			Usage:      "Where the character starts emotionally",
			InnerField: "starting_point",
		},
	},
})

var charactersList = cli.Command{
	Name:    "list",
	Usage:   "Get a paginated list of Ted Lasso characters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return (max: 100)",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "min-optimism",
			Usage:     "Minimum optimism score",
			QueryPath: "min_optimism",
		},
		&requestflag.Flag[string]{
			Name:      "role",
			Usage:     "Roles characters can have.",
			QueryPath: "role",
		},
		&requestflag.Flag[int64]{
			Name:      "skip",
			Usage:     "Number of items to skip (offset)",
			Default:   0,
			QueryPath: "skip",
		},
		&requestflag.Flag[any]{
			Name:      "team-id",
			Usage:     "Filter by team ID",
			QueryPath: "team_id",
		},
	},
	Action:          handleCharactersList,
	HideHelpCommand: true,
}

var charactersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Remove a character from the database.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "character-id",
			Required: true,
		},
	},
	Action:          handleCharactersDelete,
	HideHelpCommand: true,
}

var charactersGetQuotes = cli.Command{
	Name:    "get-quotes",
	Usage:   "Get all signature quotes from a specific character.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "character-id",
			Required: true,
		},
	},
	Action:          handleCharactersGetQuotes,
	HideHelpCommand: true,
}

func handleCharactersCreate(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.CharacterNewParams{}

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
	_, err = client.Characters.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "characters create", obj, format, transform)
}

func handleCharactersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("character-id") && len(unusedArgs) > 0 {
		cmd.Set("character-id", unusedArgs[0])
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
	_, err = client.Characters.Get(ctx, cmd.Value("character-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "characters retrieve", obj, format, transform)
}

func handleCharactersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("character-id") && len(unusedArgs) > 0 {
		cmd.Set("character-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.CharacterUpdateParams{}

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
	_, err = client.Characters.Update(
		ctx,
		cmd.Value("character-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "characters update", obj, format, transform)
}

func handleCharactersList(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.CharacterListParams{}

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
		_, err = client.Characters.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "characters list", obj, format, transform)
	} else {
		iter := client.Characters.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "characters list", iter, format, transform)
	}
}

func handleCharactersDelete(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("character-id") && len(unusedArgs) > 0 {
		cmd.Set("character-id", unusedArgs[0])
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

	return client.Characters.Delete(ctx, cmd.Value("character-id").(string), options...)
}

func handleCharactersGetQuotes(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("character-id") && len(unusedArgs) > 0 {
		cmd.Set("character-id", unusedArgs[0])
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
	_, err = client.Characters.GetQuotes(ctx, cmd.Value("character-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "characters get-quotes", obj, format, transform)
}
