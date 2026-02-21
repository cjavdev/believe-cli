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

var episodesCreate = cli.Command{
	Name:    "create",
	Usage:   "Add a new episode to the series.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "air-date",
			Usage:    "Original air date",
			Required: true,
			BodyPath: "air_date",
		},
		&requestflag.Flag[[]string]{
			Name:     "character-focus",
			Usage:    "Characters with significant development",
			Required: true,
			BodyPath: "character_focus",
		},
		&requestflag.Flag[string]{
			Name:     "director",
			Usage:    "Episode director",
			Required: true,
			BodyPath: "director",
		},
		&requestflag.Flag[int64]{
			Name:     "episode-number",
			Usage:    "Episode number within season",
			Required: true,
			BodyPath: "episode_number",
		},
		&requestflag.Flag[string]{
			Name:     "main-theme",
			Usage:    "Central theme of the episode",
			Required: true,
			BodyPath: "main_theme",
		},
		&requestflag.Flag[int64]{
			Name:     "runtime-minutes",
			Usage:    "Episode runtime in minutes",
			Required: true,
			BodyPath: "runtime_minutes",
		},
		&requestflag.Flag[int64]{
			Name:     "season",
			Usage:    "Season number",
			Required: true,
			BodyPath: "season",
		},
		&requestflag.Flag[string]{
			Name:     "synopsis",
			Usage:    "Brief plot synopsis",
			Required: true,
			BodyPath: "synopsis",
		},
		&requestflag.Flag[string]{
			Name:     "ted-wisdom",
			Usage:    "Key piece of Ted wisdom from the episode",
			Required: true,
			BodyPath: "ted_wisdom",
		},
		&requestflag.Flag[string]{
			Name:     "title",
			Usage:    "Episode title",
			Required: true,
			BodyPath: "title",
		},
		&requestflag.Flag[string]{
			Name:     "writer",
			Usage:    "Episode writer(s)",
			Required: true,
			BodyPath: "writer",
		},
		&requestflag.Flag[any]{
			Name:     "biscuits-with-boss-moment",
			Usage:    "Notable biscuits with the boss scene",
			BodyPath: "biscuits_with_boss_moment",
		},
		&requestflag.Flag[[]string]{
			Name:     "memorable-moment",
			Usage:    "Standout moments from the episode",
			BodyPath: "memorable_moments",
		},
		&requestflag.Flag[any]{
			Name:     "us-viewers-millions",
			Usage:    "US viewership in millions",
			BodyPath: "us_viewers_millions",
		},
		&requestflag.Flag[any]{
			Name:     "viewer-rating",
			Usage:    "Viewer rating out of 10",
			BodyPath: "viewer_rating",
		},
	},
	Action:          handleEpisodesCreate,
	HideHelpCommand: true,
}

var episodesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve detailed information about a specific episode.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "episode-id",
			Required: true,
		},
	},
	Action:          handleEpisodesRetrieve,
	HideHelpCommand: true,
}

var episodesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update specific fields of an existing episode.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "episode-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "air-date",
			BodyPath: "air_date",
		},
		&requestflag.Flag[any]{
			Name:     "biscuits-with-boss-moment",
			BodyPath: "biscuits_with_boss_moment",
		},
		&requestflag.Flag[any]{
			Name:     "character-focus",
			BodyPath: "character_focus",
		},
		&requestflag.Flag[any]{
			Name:     "director",
			BodyPath: "director",
		},
		&requestflag.Flag[any]{
			Name:     "episode-number",
			BodyPath: "episode_number",
		},
		&requestflag.Flag[any]{
			Name:     "main-theme",
			BodyPath: "main_theme",
		},
		&requestflag.Flag[any]{
			Name:     "memorable-moment",
			BodyPath: "memorable_moments",
		},
		&requestflag.Flag[any]{
			Name:     "runtime-minutes",
			BodyPath: "runtime_minutes",
		},
		&requestflag.Flag[any]{
			Name:     "season",
			BodyPath: "season",
		},
		&requestflag.Flag[any]{
			Name:     "synopsis",
			BodyPath: "synopsis",
		},
		&requestflag.Flag[any]{
			Name:     "ted-wisdom",
			BodyPath: "ted_wisdom",
		},
		&requestflag.Flag[any]{
			Name:     "title",
			BodyPath: "title",
		},
		&requestflag.Flag[any]{
			Name:     "us-viewers-millions",
			BodyPath: "us_viewers_millions",
		},
		&requestflag.Flag[any]{
			Name:     "viewer-rating",
			BodyPath: "viewer_rating",
		},
		&requestflag.Flag[any]{
			Name:     "writer",
			BodyPath: "writer",
		},
	},
	Action:          handleEpisodesUpdate,
	HideHelpCommand: true,
}

var episodesList = cli.Command{
	Name:    "list",
	Usage:   "Get a paginated list of all Ted Lasso episodes with optional filtering by\nseason.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "character-focus",
			Usage:     "Filter by character focus (character ID)",
			QueryPath: "character_focus",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return (max: 100)",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "season",
			Usage:     "Filter by season",
			QueryPath: "season",
		},
		&requestflag.Flag[int64]{
			Name:      "skip",
			Usage:     "Number of items to skip (offset)",
			Default:   0,
			QueryPath: "skip",
		},
	},
	Action:          handleEpisodesList,
	HideHelpCommand: true,
}

var episodesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Remove an episode from the database.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "episode-id",
			Required: true,
		},
	},
	Action:          handleEpisodesDelete,
	HideHelpCommand: true,
}

var episodesGetWisdom = cli.Command{
	Name:    "get-wisdom",
	Usage:   "Get Ted's wisdom and memorable moments from a specific episode.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "episode-id",
			Required: true,
		},
	},
	Action:          handleEpisodesGetWisdom,
	HideHelpCommand: true,
}

var episodesListBySeason = cli.Command{
	Name:    "list-by-season",
	Usage:   "Get a paginated list of episodes from a specific season.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "season-number",
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
	},
	Action:          handleEpisodesListBySeason,
	HideHelpCommand: true,
}

func handleEpisodesCreate(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.EpisodeNewParams{}

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
	_, err = client.Episodes.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "episodes create", obj, format, transform)
}

func handleEpisodesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("episode-id") && len(unusedArgs) > 0 {
		cmd.Set("episode-id", unusedArgs[0])
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
	_, err = client.Episodes.Get(ctx, cmd.Value("episode-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "episodes retrieve", obj, format, transform)
}

func handleEpisodesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("episode-id") && len(unusedArgs) > 0 {
		cmd.Set("episode-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.EpisodeUpdateParams{}

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
	_, err = client.Episodes.Update(
		ctx,
		cmd.Value("episode-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "episodes update", obj, format, transform)
}

func handleEpisodesList(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.EpisodeListParams{}

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
		_, err = client.Episodes.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "episodes list", obj, format, transform)
	} else {
		iter := client.Episodes.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "episodes list", iter, format, transform)
	}
}

func handleEpisodesDelete(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("episode-id") && len(unusedArgs) > 0 {
		cmd.Set("episode-id", unusedArgs[0])
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

	return client.Episodes.Delete(ctx, cmd.Value("episode-id").(string), options...)
}

func handleEpisodesGetWisdom(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("episode-id") && len(unusedArgs) > 0 {
		cmd.Set("episode-id", unusedArgs[0])
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
	_, err = client.Episodes.GetWisdom(ctx, cmd.Value("episode-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "episodes get-wisdom", obj, format, transform)
}

func handleEpisodesListBySeason(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("season-number") && len(unusedArgs) > 0 {
		cmd.Set("season-number", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.EpisodeListBySeasonParams{}

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
		_, err = client.Episodes.ListBySeason(
			ctx,
			cmd.Value("season-number").(int64),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "episodes list-by-season", obj, format, transform)
	} else {
		iter := client.Episodes.ListBySeasonAutoPaging(
			ctx,
			cmd.Value("season-number").(int64),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "episodes list-by-season", iter, format, transform)
	}
}
