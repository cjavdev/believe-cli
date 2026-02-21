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

var matchesCommentaryStream = cli.Command{
	Name:    "stream",
	Usage:   "Stream live match commentary for a specific match. Uses Server-Sent Events (SSE)\nto stream commentary events in real-time.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "match-id",
			Required: true,
		},
	},
	Action:          handleMatchesCommentaryStream,
	HideHelpCommand: true,
}

func handleMatchesCommentaryStream(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Matches.Commentary.Stream(ctx, cmd.Value("match-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "matches:commentary stream", obj, format, transform)
}
