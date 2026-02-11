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

var conflictsResolve = cli.Command{
	Name:    "resolve",
	Usage:   "Get Ted Lasso-style advice for resolving conflicts.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "conflict-type",
			Usage:    "Type of conflict",
			Required: true,
			BodyPath: "conflict_type",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Describe the conflict",
			Required: true,
			BodyPath: "description",
		},
		&requestflag.Flag[[]string]{
			Name:     "parties-involved",
			Usage:    "Who is involved in the conflict",
			Required: true,
			BodyPath: "parties_involved",
		},
		&requestflag.Flag[any]{
			Name:     "attempts-made",
			Usage:    "What you've already tried",
			BodyPath: "attempts_made",
		},
	},
	Action:          handleConflictsResolve,
	HideHelpCommand: true,
}

func handleConflictsResolve(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.ConflictResolveParams{}

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
	_, err = client.Conflicts.Resolve(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conflicts resolve", obj, format, transform)
}
