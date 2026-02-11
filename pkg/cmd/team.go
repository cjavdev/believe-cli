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

var teamsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Add a new team to the league.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "culture-score",
			Usage:    "Team culture/morale score (0-100)",
			Required: true,
			BodyPath: "culture_score",
		},
		&requestflag.Flag[int64]{
			Name:     "founded-year",
			Usage:    "Year the club was founded",
			Required: true,
			BodyPath: "founded_year",
		},
		&requestflag.Flag[string]{
			Name:     "league",
			Usage:    "Football leagues.",
			Required: true,
			BodyPath: "league",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Team name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "stadium",
			Usage:    "Home stadium name",
			Required: true,
			BodyPath: "stadium",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "values",
			Usage:    "Core values that define a team's culture.",
			Required: true,
			BodyPath: "values",
		},
		&requestflag.Flag[any]{
			Name:     "annual-budget-gbp",
			Usage:    "Annual budget in GBP",
			BodyPath: "annual_budget_gbp",
		},
		&requestflag.Flag[any]{
			Name:     "average-attendance",
			Usage:    "Average match attendance",
			BodyPath: "average_attendance",
		},
		&requestflag.Flag[any]{
			Name:     "contact-email",
			Usage:    "Team contact email",
			BodyPath: "contact_email",
		},
		&requestflag.Flag[bool]{
			Name:     "is-active",
			Usage:    "Whether the team is currently active",
			Default:  true,
			BodyPath: "is_active",
		},
		&requestflag.Flag[any]{
			Name:     "nickname",
			Usage:    "Team nickname",
			BodyPath: "nickname",
		},
		&requestflag.Flag[any]{
			Name:     "primary-color",
			Usage:    "Primary team color (hex)",
			BodyPath: "primary_color",
		},
		&requestflag.Flag[[]string]{
			Name:     "rival-team",
			Usage:    "List of rival team IDs",
			BodyPath: "rival_teams",
		},
		&requestflag.Flag[any]{
			Name:     "secondary-color",
			Usage:    "Secondary team color (hex)",
			BodyPath: "secondary_color",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "stadium-location",
			Usage:    "Geographic coordinates for a location.",
			BodyPath: "stadium_location",
		},
		&requestflag.Flag[any]{
			Name:     "website",
			Usage:    "Official team website",
			BodyPath: "website",
		},
		&requestflag.Flag[any]{
			Name:     "win-percentage",
			Usage:    "Season win percentage",
			BodyPath: "win_percentage",
		},
	},
	Action:          handleTeamsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"values": {
		&requestflag.InnerFlag[string]{
			Name:       "values.primary-value",
			Usage:      "The team's primary guiding value",
			InnerField: "primary_value",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "values.secondary-values",
			Usage:      "Supporting values",
			InnerField: "secondary_values",
		},
		&requestflag.InnerFlag[string]{
			Name:       "values.team-motto",
			Usage:      "Team's motivational motto",
			InnerField: "team_motto",
		},
	},
	"stadium-location": {
		&requestflag.InnerFlag[float64]{
			Name:       "stadium-location.latitude",
			Usage:      "Latitude in degrees",
			InnerField: "latitude",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "stadium-location.longitude",
			Usage:      "Longitude in degrees",
			InnerField: "longitude",
		},
	},
})

var teamsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve detailed information about a specific team.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "team-id",
			Required: true,
		},
	},
	Action:          handleTeamsRetrieve,
	HideHelpCommand: true,
}

var teamsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update specific fields of an existing team.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "team-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "annual-budget-gbp",
			BodyPath: "annual_budget_gbp",
		},
		&requestflag.Flag[any]{
			Name:     "average-attendance",
			BodyPath: "average_attendance",
		},
		&requestflag.Flag[any]{
			Name:     "contact-email",
			BodyPath: "contact_email",
		},
		&requestflag.Flag[any]{
			Name:     "culture-score",
			BodyPath: "culture_score",
		},
		&requestflag.Flag[any]{
			Name:     "founded-year",
			BodyPath: "founded_year",
		},
		&requestflag.Flag[any]{
			Name:     "is-active",
			BodyPath: "is_active",
		},
		&requestflag.Flag[string]{
			Name:     "league",
			Usage:    "Football leagues.",
			BodyPath: "league",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "nickname",
			BodyPath: "nickname",
		},
		&requestflag.Flag[any]{
			Name:     "primary-color",
			BodyPath: "primary_color",
		},
		&requestflag.Flag[any]{
			Name:     "rival-team",
			BodyPath: "rival_teams",
		},
		&requestflag.Flag[any]{
			Name:     "secondary-color",
			BodyPath: "secondary_color",
		},
		&requestflag.Flag[any]{
			Name:     "stadium",
			BodyPath: "stadium",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "stadium-location",
			Usage:    "Geographic coordinates for a location.",
			BodyPath: "stadium_location",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "values",
			Usage:    "Core values that define a team's culture.",
			BodyPath: "values",
		},
		&requestflag.Flag[any]{
			Name:     "website",
			BodyPath: "website",
		},
		&requestflag.Flag[any]{
			Name:     "win-percentage",
			BodyPath: "win_percentage",
		},
	},
	Action:          handleTeamsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"stadium-location": {
		&requestflag.InnerFlag[float64]{
			Name:       "stadium-location.latitude",
			Usage:      "Latitude in degrees",
			InnerField: "latitude",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "stadium-location.longitude",
			Usage:      "Longitude in degrees",
			InnerField: "longitude",
		},
	},
	"values": {
		&requestflag.InnerFlag[string]{
			Name:       "values.primary-value",
			Usage:      "The team's primary guiding value",
			InnerField: "primary_value",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "values.secondary-values",
			Usage:      "Supporting values",
			InnerField: "secondary_values",
		},
		&requestflag.InnerFlag[string]{
			Name:       "values.team-motto",
			Usage:      "Team's motivational motto",
			InnerField: "team_motto",
		},
	},
})

var teamsList = cli.Command{
	Name:    "list",
	Usage:   "Get a paginated list of all teams with optional filtering by league or culture\nscore.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "league",
			Usage:     "Football leagues.",
			QueryPath: "league",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return (max: 100)",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "min-culture-score",
			Usage:     "Minimum culture score",
			QueryPath: "min_culture_score",
		},
		&requestflag.Flag[int64]{
			Name:      "skip",
			Usage:     "Number of items to skip (offset)",
			Default:   0,
			QueryPath: "skip",
		},
	},
	Action:          handleTeamsList,
	HideHelpCommand: true,
}

var teamsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Remove a team from the database (relegation to oblivion).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "team-id",
			Required: true,
		},
	},
	Action:          handleTeamsDelete,
	HideHelpCommand: true,
}

var teamsGetCulture = cli.Command{
	Name:    "get-culture",
	Usage:   "Get detailed culture and values information for a team.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "team-id",
			Required: true,
		},
	},
	Action:          handleTeamsGetCulture,
	HideHelpCommand: true,
}

var teamsGetRivals = cli.Command{
	Name:    "get-rivals",
	Usage:   "Get all rival teams for a specific team.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "team-id",
			Required: true,
		},
	},
	Action:          handleTeamsGetRivals,
	HideHelpCommand: true,
}

var teamsListLogos = cli.Command{
	Name:    "list-logos",
	Usage:   "List all uploaded logos for a team.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "team-id",
			Required: true,
		},
	},
	Action:          handleTeamsListLogos,
	HideHelpCommand: true,
}

func handleTeamsCreate(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.TeamNewParams{}

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
	_, err = client.Teams.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "teams create", obj, format, transform)
}

func handleTeamsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("team-id") && len(unusedArgs) > 0 {
		cmd.Set("team-id", unusedArgs[0])
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
	_, err = client.Teams.Get(ctx, cmd.Value("team-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "teams retrieve", obj, format, transform)
}

func handleTeamsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("team-id") && len(unusedArgs) > 0 {
		cmd.Set("team-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.TeamUpdateParams{}

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
	_, err = client.Teams.Update(
		ctx,
		cmd.Value("team-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "teams update", obj, format, transform)
}

func handleTeamsList(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.TeamListParams{}

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
		_, err = client.Teams.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "teams list", obj, format, transform)
	} else {
		iter := client.Teams.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "teams list", iter, format, transform)
	}
}

func handleTeamsDelete(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("team-id") && len(unusedArgs) > 0 {
		cmd.Set("team-id", unusedArgs[0])
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

	return client.Teams.Delete(ctx, cmd.Value("team-id").(string), options...)
}

func handleTeamsGetCulture(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("team-id") && len(unusedArgs) > 0 {
		cmd.Set("team-id", unusedArgs[0])
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
	_, err = client.Teams.GetCulture(ctx, cmd.Value("team-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "teams get-culture", obj, format, transform)
}

func handleTeamsGetRivals(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("team-id") && len(unusedArgs) > 0 {
		cmd.Set("team-id", unusedArgs[0])
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
	_, err = client.Teams.GetRivals(ctx, cmd.Value("team-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "teams get-rivals", obj, format, transform)
}

func handleTeamsListLogos(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("team-id") && len(unusedArgs) > 0 {
		cmd.Set("team-id", unusedArgs[0])
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
	_, err = client.Teams.ListLogos(ctx, cmd.Value("team-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "teams list-logos", obj, format, transform)
}
