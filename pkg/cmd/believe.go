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

var believeSubmit = cli.Command{
	Name:    "submit",
	Usage:   "Submit your situation and receive Ted Lasso-style motivational guidance.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "situation",
			Usage:    "Describe your situation",
			Required: true,
			BodyPath: "situation",
		},
		&requestflag.Flag[string]{
			Name:     "situation-type",
			Usage:    "Type of situation",
			Required: true,
			BodyPath: "situation_type",
		},
		&requestflag.Flag[any]{
			Name:     "context",
			Usage:    "Additional context",
			BodyPath: "context",
		},
		&requestflag.Flag[int64]{
			Name:     "intensity",
			Usage:    "How intense is the response needed (1=gentle, 10=full Ted)",
			Default:  5,
			BodyPath: "intensity",
		},
	},
	Action:          handleBelieveSubmit,
	HideHelpCommand: true,
}

func handleBelieveSubmit(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := believe.BelieveSubmitParams{}

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
	_, err = client.Believe.Submit(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "believe submit", obj, format, transform)
}
