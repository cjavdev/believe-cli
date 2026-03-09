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

var teamMembersCreate = cli.Command{
	Name:    "create",
	Usage:   "Add a new team member to a team.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "member",
			Usage:    "A football player on the team.",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleTeamMembersCreate,
	HideHelpCommand: true,
}

var teamMembersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve detailed information about a specific team member.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "member-id",
			Required: true,
		},
	},
	Action:          handleTeamMembersRetrieve,
	HideHelpCommand: true,
}

var teamMembersUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update specific fields of an existing team member. Fields vary by member type.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "member-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "updates",
			Usage:    "Update model for players.",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleTeamMembersUpdate,
	HideHelpCommand: true,
}

var teamMembersList = cli.Command{
	Name:    "list",
	Usage:   "Get a paginated list of all team members.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return (max: 100)",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "member-type",
			Usage:     "Filter by member type",
			QueryPath: "member_type",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleTeamMembersList,
	HideHelpCommand: true,
}

var teamMembersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Remove a team member from the roster.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "member-id",
			Required: true,
		},
	},
	Action:          handleTeamMembersDelete,
	HideHelpCommand: true,
}

var teamMembersListCoaches = cli.Command{
	Name:    "list-coaches",
	Usage:   "Get only coaches (filtered subset of team members).",
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
		&requestflag.Flag[string]{
			Name:      "specialty",
			Usage:     "Coaching specialties.",
			QueryPath: "specialty",
		},
		&requestflag.Flag[any]{
			Name:      "team-id",
			Usage:     "Filter by team ID",
			QueryPath: "team_id",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleTeamMembersListCoaches,
	HideHelpCommand: true,
}

var teamMembersListPlayers = cli.Command{
	Name:    "list-players",
	Usage:   "Get only players (filtered subset of team members).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return (max: 100)",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "position",
			Usage:     "Football positions for players.",
			QueryPath: "position",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleTeamMembersListPlayers,
	HideHelpCommand: true,
}

var teamMembersListStaff = cli.Command{
	Name:    "list-staff",
	Usage:   "Get all staff members (medical staff and equipment managers).",
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
		&requestflag.Flag[any]{
			Name:      "team-id",
			Usage:     "Filter by team ID",
			QueryPath: "team_id",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleTeamMembersListStaff,
	HideHelpCommand: true,
}

func handleTeamMembersCreate(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.TeamMemberNewParams{}

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
	_, err = client.TeamMembers.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "team-members create", obj, format, transform)
}

func handleTeamMembersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("member-id") && len(unusedArgs) > 0 {
		cmd.Set("member-id", unusedArgs[0])
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
	_, err = client.TeamMembers.Get(ctx, cmd.Value("member-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "team-members retrieve", obj, format, transform)
}

func handleTeamMembersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("member-id") && len(unusedArgs) > 0 {
		cmd.Set("member-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.TeamMemberUpdateParams{}

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
	_, err = client.TeamMembers.Update(
		ctx,
		cmd.Value("member-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "team-members update", obj, format, transform)
}

func handleTeamMembersList(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.TeamMemberListParams{}

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
		_, err = client.TeamMembers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "team-members list", obj, format, transform)
	} else {
		iter := client.TeamMembers.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "team-members list", iter, format, transform, maxItems)
	}
}

func handleTeamMembersDelete(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("member-id") && len(unusedArgs) > 0 {
		cmd.Set("member-id", unusedArgs[0])
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

	return client.TeamMembers.Delete(ctx, cmd.Value("member-id").(string), options...)
}

func handleTeamMembersListCoaches(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.TeamMemberListCoachesParams{}

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
		_, err = client.TeamMembers.ListCoaches(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "team-members list-coaches", obj, format, transform)
	} else {
		iter := client.TeamMembers.ListCoachesAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "team-members list-coaches", iter, format, transform, maxItems)
	}
}

func handleTeamMembersListPlayers(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.TeamMemberListPlayersParams{}

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
		_, err = client.TeamMembers.ListPlayers(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "team-members list-players", obj, format, transform)
	} else {
		iter := client.TeamMembers.ListPlayersAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "team-members list-players", iter, format, transform, maxItems)
	}
}

func handleTeamMembersListStaff(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.TeamMemberListStaffParams{}

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
		_, err = client.TeamMembers.ListStaff(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "team-members list-staff", obj, format, transform)
	} else {
		iter := client.TeamMembers.ListStaffAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "team-members list-staff", iter, format, transform, maxItems)
	}
}
