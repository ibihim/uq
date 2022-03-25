package main

import (
	"fmt"
	"os"

	"github.com/ibihim/uq/cmd"
)

func main() {
	cmd.Root.AddCommand(cmd.Query)

	if err := cmd.Root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
