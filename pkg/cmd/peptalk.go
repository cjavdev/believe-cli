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

var pepTalkRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a motivational pep talk from Ted Lasso himself. By default returns the\ncomplete pep talk. Add `?stream=true` to get Server-Sent Events (SSE) streaming\nthe pep talk chunk by chunk.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:      "stream",
			Usage:     "If true, returns SSE stream instead of full response",
			Default:   false,
			QueryPath: "stream",
		},
	},
	Action:          handlePepTalkRetrieve,
	HideHelpCommand: true,
}

func handlePepTalkRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.PepTalkGetParams{}

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
	_, err = client.PepTalk.Get(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "pep-talk retrieve", obj, format, transform)
}
