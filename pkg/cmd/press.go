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

var pressSimulate = cli.Command{
	Name:    "simulate",
	Usage:   "Get Ted's response to press conference questions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "question",
			Usage:    "The press question to answer",
			Required: true,
			BodyPath: "question",
		},
		&requestflag.Flag[bool]{
			Name:     "hostile",
			Usage:    "Is this a hostile question from Trent Crimm?",
			Default:  false,
			BodyPath: "hostile",
		},
		&requestflag.Flag[any]{
			Name:     "topic",
			Usage:    "Topic category",
			BodyPath: "topic",
		},
	},
	Action:          handlePressSimulate,
	HideHelpCommand: true,
}

func handlePressSimulate(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.PressSimulateParams{}

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
	_, err = client.Press.Simulate(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "press simulate", obj, format, transform)
}
