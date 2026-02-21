// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/cjavdev/believe-cli/internal/autocomplete"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command *cli.Command
)

func init() {
	Command = &cli.Command{
		Name:    "believe",
		Usage:   "CLI for the believe API",
		Suggest: true,
		Version: Version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
		},
		Commands: []*cli.Command{
			&getWelcome,
			{
				Name:     "characters",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&charactersCreate,
					&charactersRetrieve,
					&charactersUpdate,
					&charactersList,
					&charactersDelete,
					&charactersGetQuotes,
				},
			},
			{
				Name:     "teams",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&teamsCreate,
					&teamsRetrieve,
					&teamsUpdate,
					&teamsList,
					&teamsDelete,
					&teamsGetCulture,
					&teamsGetRivals,
					&teamsListLogos,
				},
			},
			{
				Name:     "teams:logo",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&teamsLogoDelete,
					&teamsLogoDownload,
					&teamsLogoUpload,
				},
			},
			{
				Name:     "matches",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&matchesCreate,
					&matchesRetrieve,
					&matchesUpdate,
					&matchesList,
					&matchesDelete,
					&matchesGetLesson,
					&matchesGetTurningPoints,
					&matchesStreamLive,
				},
			},
			{
				Name:     "matches:commentary",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&matchesCommentaryStream,
				},
			},
			{
				Name:     "episodes",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&episodesCreate,
					&episodesRetrieve,
					&episodesUpdate,
					&episodesList,
					&episodesDelete,
					&episodesGetWisdom,
					&episodesListBySeason,
				},
			},
			{
				Name:     "quotes",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&quotesCreate,
					&quotesRetrieve,
					&quotesUpdate,
					&quotesList,
					&quotesDelete,
					&quotesGetRandom,
					&quotesListByCharacter,
					&quotesListByTheme,
				},
			},
			{
				Name:     "believe",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&believeSubmit,
				},
			},
			{
				Name:     "conflicts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&conflictsResolve,
				},
			},
			{
				Name:     "reframe",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&reframeTransformNegativeThoughts,
				},
			},
			{
				Name:     "press",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&pressSimulate,
				},
			},
			{
				Name:     "coaching:principles",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&coachingPrinciplesRetrieve,
					&coachingPrinciplesList,
					&coachingPrinciplesGetRandom,
				},
			},
			{
				Name:     "biscuits",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&biscuitsRetrieve,
					&biscuitsList,
					&biscuitsGetFresh,
				},
			},
			{
				Name:     "pep-talk",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&pepTalkRetrieve,
				},
			},
			{
				Name:     "stream",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&streamTestConnection,
				},
			},
			{
				Name:     "team-members",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&teamMembersCreate,
					&teamMembersRetrieve,
					&teamMembersUpdate,
					&teamMembersList,
					&teamMembersDelete,
					&teamMembersListCoaches,
					&teamMembersListPlayers,
					&teamMembersListStaff,
				},
			},
			{
				Name:     "webhooks",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&webhooksCreate,
					&webhooksRetrieve,
					&webhooksList,
					&webhooksDelete,
					&webhooksTriggerEvent,
				},
			},
			{
				Name:     "health",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&healthCheck,
				},
			},
			{
				Name:     "version",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&versionRetrieve,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "believe @manpages [-o believe.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "believe.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "believe.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
