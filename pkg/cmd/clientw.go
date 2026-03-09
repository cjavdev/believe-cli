// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/cjavdev/believe-cli/internal/apiquery"
	"github.com/cjavdev/believe-go"
	"github.com/urfave/cli/v3"
)

var clientWsTest = cli.Command{
	Name:            "test",
	Usage:           "Simple WebSocket test endpoint for connectivity testing.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleClientWsTest,
	HideHelpCommand: true,
}

func handleClientWsTest(ctx context.Context, cmd *cli.Command) error {
	client := believe.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	return client.Client.Ws.Test(ctx, options...)
}
