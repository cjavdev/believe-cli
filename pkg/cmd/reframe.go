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

var reframeTransformNegativeThoughts = cli.Command{
	Name:    "transform-negative-thoughts",
	Usage:   "Transform negative thoughts into positive perspectives with Ted's help.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "negative-thought",
			Usage:    "The negative thought to reframe",
			Required: true,
			BodyPath: "negative_thought",
		},
		&requestflag.Flag[bool]{
			Name:     "recurring",
			Usage:    "Is this a recurring thought?",
			Default:  false,
			BodyPath: "recurring",
		},
	},
	Action:          handleReframeTransformNegativeThoughts,
	HideHelpCommand: true,
}

func handleReframeTransformNegativeThoughts(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.ReframeTransformNegativeThoughtsParams{}

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
	_, err = client.Reframe.TransformNegativeThoughts(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "reframe transform-negative-thoughts", obj, format, transform)
}
