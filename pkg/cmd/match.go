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

var matchesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Schedule a new match.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "away-team-id",
			Usage:    "Away team ID",
			Required: true,
			BodyPath: "away_team_id",
		},
		&requestflag.Flag[any]{
			Name:     "date",
			Usage:    "Match date and time",
			Required: true,
			BodyPath: "date",
		},
		&requestflag.Flag[string]{
			Name:     "home-team-id",
			Usage:    "Home team ID",
			Required: true,
			BodyPath: "home_team_id",
		},
		&requestflag.Flag[string]{
			Name:     "match-type",
			Usage:    "Types of matches.",
			Required: true,
			BodyPath: "match_type",
		},
		&requestflag.Flag[any]{
			Name:     "attendance",
			Usage:    "Match attendance",
			BodyPath: "attendance",
		},
		&requestflag.Flag[int64]{
			Name:     "away-score",
			Usage:    "Away team score",
			Default:  0,
			BodyPath: "away_score",
		},
		&requestflag.Flag[any]{
			Name:     "episode-id",
			Usage:    "Episode ID where this match is featured",
			BodyPath: "episode_id",
		},
		&requestflag.Flag[int64]{
			Name:     "home-score",
			Usage:    "Home team score",
			Default:  0,
			BodyPath: "home_score",
		},
		&requestflag.Flag[any]{
			Name:     "lesson-learned",
			Usage:    "The life lesson learned from this match",
			BodyPath: "lesson_learned",
		},
		&requestflag.Flag[any]{
			Name:     "possession-percentage",
			Usage:    "Home team possession percentage",
			BodyPath: "possession_percentage",
		},
		&requestflag.Flag[string]{
			Name:     "result",
			Usage:    "Match result types.",
			BodyPath: "result",
		},
		&requestflag.Flag[any]{
			Name:     "ted-halftime-speech",
			Usage:    "Ted's inspirational halftime speech",
			BodyPath: "ted_halftime_speech",
		},
		&requestflag.Flag[any]{
			Name:     "ticket-revenue-gbp",
			Usage:    "Total ticket revenue in GBP",
			BodyPath: "ticket_revenue_gbp",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "turning-point",
			Usage:    "Key moments that changed the match",
			BodyPath: "turning_points",
		},
		&requestflag.Flag[any]{
			Name:     "weather-temp-celsius",
			Usage:    "Temperature at kickoff in Celsius",
			BodyPath: "weather_temp_celsius",
		},
	},
	Action:          handleMatchesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"turning-point": {
		&requestflag.InnerFlag[string]{
			Name:       "turning-point.description",
			Usage:      "What happened",
			InnerField: "description",
		},
		&requestflag.InnerFlag[string]{
			Name:       "turning-point.emotional-impact",
			Usage:      "How this affected the team emotionally",
			InnerField: "emotional_impact",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "turning-point.minute",
			Usage:      "Minute of the match",
			InnerField: "minute",
		},
		&requestflag.InnerFlag[any]{
			Name:       "turning-point.character-involved",
			Usage:      "Character ID who was central to this moment",
			InnerField: "character_involved",
		},
	},
})

var matchesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve detailed information about a specific match.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "match-id",
			Required: true,
		},
	},
	Action:          handleMatchesRetrieve,
	HideHelpCommand: true,
}

var matchesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update specific fields of an existing match (e.g., update score).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "match-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "attendance",
			BodyPath: "attendance",
		},
		&requestflag.Flag[any]{
			Name:     "away-score",
			BodyPath: "away_score",
		},
		&requestflag.Flag[any]{
			Name:     "away-team-id",
			BodyPath: "away_team_id",
		},
		&requestflag.Flag[any]{
			Name:     "date",
			BodyPath: "date",
		},
		&requestflag.Flag[any]{
			Name:     "episode-id",
			BodyPath: "episode_id",
		},
		&requestflag.Flag[any]{
			Name:     "home-score",
			BodyPath: "home_score",
		},
		&requestflag.Flag[any]{
			Name:     "home-team-id",
			BodyPath: "home_team_id",
		},
		&requestflag.Flag[any]{
			Name:     "lesson-learned",
			BodyPath: "lesson_learned",
		},
		&requestflag.Flag[string]{
			Name:     "match-type",
			Usage:    "Types of matches.",
			BodyPath: "match_type",
		},
		&requestflag.Flag[any]{
			Name:     "possession-percentage",
			BodyPath: "possession_percentage",
		},
		&requestflag.Flag[string]{
			Name:     "result",
			Usage:    "Match result types.",
			BodyPath: "result",
		},
		&requestflag.Flag[any]{
			Name:     "ted-halftime-speech",
			BodyPath: "ted_halftime_speech",
		},
		&requestflag.Flag[any]{
			Name:     "ticket-revenue-gbp",
			BodyPath: "ticket_revenue_gbp",
		},
		&requestflag.Flag[any]{
			Name:     "turning-point",
			BodyPath: "turning_points",
		},
		&requestflag.Flag[any]{
			Name:     "weather-temp-celsius",
			BodyPath: "weather_temp_celsius",
		},
	},
	Action:          handleMatchesUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"turning-point": {
		&requestflag.InnerFlag[string]{
			Name:       "turning-point.description",
			Usage:      "What happened",
			InnerField: "description",
		},
		&requestflag.InnerFlag[string]{
			Name:       "turning-point.emotional-impact",
			Usage:      "How this affected the team emotionally",
			InnerField: "emotional_impact",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "turning-point.minute",
			Usage:      "Minute of the match",
			InnerField: "minute",
		},
		&requestflag.InnerFlag[any]{
			Name:       "turning-point.character-involved",
			Usage:      "Character ID who was central to this moment",
			InnerField: "character_involved",
		},
	},
})

var matchesList = cli.Command{
	Name:    "list",
	Usage:   "Get a paginated list of all matches with optional filtering.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return (max: 100)",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "match-type",
			Usage:     "Types of matches.",
			QueryPath: "match_type",
		},
		&requestflag.Flag[string]{
			Name:      "result",
			Usage:     "Match result types.",
			QueryPath: "result",
		},
		&requestflag.Flag[int64]{
			Name:      "skip",
			Usage:     "Number of items to skip (offset)",
			Default:   0,
			QueryPath: "skip",
		},
		&requestflag.Flag[any]{
			Name:      "team-id",
			Usage:     "Filter by team (home or away)",
			QueryPath: "team_id",
		},
	},
	Action:          handleMatchesList,
	HideHelpCommand: true,
}

var matchesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Remove a match from the database.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "match-id",
			Required: true,
		},
	},
	Action:          handleMatchesDelete,
	HideHelpCommand: true,
}

var matchesGetLesson = cli.Command{
	Name:    "get-lesson",
	Usage:   "Get the life lesson learned from a specific match.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "match-id",
			Required: true,
		},
	},
	Action:          handleMatchesGetLesson,
	HideHelpCommand: true,
}

var matchesGetTurningPoints = cli.Command{
	Name:    "get-turning-points",
	Usage:   "Get all turning points from a specific match.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "match-id",
			Required: true,
		},
	},
	Action:          handleMatchesGetTurningPoints,
	HideHelpCommand: true,
}

var matchesStreamLive = cli.Command{
	Name:    "stream-live",
	Usage:   "WebSocket endpoint for real-time live match simulation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "away-team",
			Usage:     "Away team name",
			Default:   "West Ham United",
			QueryPath: "away_team",
		},
		&requestflag.Flag[int64]{
			Name:      "excitement-level",
			Usage:     "How eventful the match should be (1=boring, 10=chaos)",
			Default:   5,
			QueryPath: "excitement_level",
		},
		&requestflag.Flag[string]{
			Name:      "home-team",
			Usage:     "Home team name",
			Default:   "AFC Richmond",
			QueryPath: "home_team",
		},
		&requestflag.Flag[float64]{
			Name:      "speed",
			Usage:     "Simulation speed multiplier (1.0 = real-time)",
			Default:   1,
			QueryPath: "speed",
		},
	},
	Action:          handleMatchesStreamLive,
	HideHelpCommand: true,
}

func handleMatchesCreate(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.MatchNewParams{}

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
	_, err = client.Matches.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "matches create", obj, format, transform)
}

func handleMatchesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("match-id") && len(unusedArgs) > 0 {
		cmd.Set("match-id", unusedArgs[0])
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
	_, err = client.Matches.Get(ctx, cmd.Value("match-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "matches retrieve", obj, format, transform)
}

func handleMatchesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("match-id") && len(unusedArgs) > 0 {
		cmd.Set("match-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.MatchUpdateParams{}

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
	_, err = client.Matches.Update(
		ctx,
		cmd.Value("match-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "matches update", obj, format, transform)
}

func handleMatchesList(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.MatchListParams{}

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
		_, err = client.Matches.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "matches list", obj, format, transform)
	} else {
		iter := client.Matches.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "matches list", iter, format, transform)
	}
}

func handleMatchesDelete(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("match-id") && len(unusedArgs) > 0 {
		cmd.Set("match-id", unusedArgs[0])
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

	return client.Matches.Delete(ctx, cmd.Value("match-id").(string), options...)
}

func handleMatchesGetLesson(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("match-id") && len(unusedArgs) > 0 {
		cmd.Set("match-id", unusedArgs[0])
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
	_, err = client.Matches.GetLesson(ctx, cmd.Value("match-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "matches get-lesson", obj, format, transform)
}

func handleMatchesGetTurningPoints(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("match-id") && len(unusedArgs) > 0 {
		cmd.Set("match-id", unusedArgs[0])
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
	_, err = client.Matches.GetTurningPoints(ctx, cmd.Value("match-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "matches get-turning-points", obj, format, transform)
}

func handleMatchesStreamLive(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.MatchStreamLiveParams{}

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

	return client.Matches.StreamLive(ctx, params, options...)
}
