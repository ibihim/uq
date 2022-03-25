package cmd

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ibihim/uq/io"
)

// Root represents the base command when called without any subcommands
var Root = &cobra.Command{
	Use:   "uq",
	Short: "Modifies URL queries and pretty prints URLs",
	Long: `Pretty prints a URL, when piped. Has additional commands:

query add key value - adds key value to query
query remove key - removes key from query`,
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
