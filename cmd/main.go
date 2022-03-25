package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ibihim/uq/cmd/query"
	"github.com/ibihim/uq/pkg/io"
)

func main() {
	rootCmd.AddCommand(query.Cmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "uq",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if err := io.ActOn(os.Stdin, func(s string) error {
			u, err := url.Parse(
				strings.ReplaceAll(
					s,
					"\"", "",
				),
			)
			if err != nil {
				return err
			}

			io.PrettyURL(u)
			return nil
		}); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}
