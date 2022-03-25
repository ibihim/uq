package cmd

import (
	"bufio"
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

var (
	addVerb    = "add"
	removeVerb = "remove"

	// Query represents the add command
	Query = &cobra.Command{
		Use:   "query",
		Short: "Modifies query parameters of a URL",
		Long: `Query has two verbs:

query add key value - adds key value to query
query remove key - removes key from query`,
		Run: queryRun,
	}
)

type conf struct {
	verb  string
	key   string
	value string
}

func makeConf(args []string) (conf, error) {
	var c conf

	if len(args) < 1 {
		return c, fmt.Errorf("`query` is used with `add key value` or `remove key`")
	}

	switch args[0] {
	case addVerb:
		if len(args) < 3 {
			return c, fmt.Errorf("`query add` is used with a `key` and a `value`")
		}

		return conf{
			verb:  addVerb,
			key:   args[1],
			value: args[2],
		}, nil
	case removeVerb:
		if len(args) < 2 {
			return c, fmt.Errorf("`query remove` is used with a `key`")
		}

		return conf{
			verb: removeVerb,
			key:  args[1],
		}, nil
	}

	return c, nil
}

func queryRun(cmd *cobra.Command, args []string) {
	conf, err := makeConf(args)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	urls, err := readURL()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	for i := range urls {
		q := urls[i].Query()

		switch conf.verb {
		case addVerb:
			q.Add(conf.key, conf.value)
		case removeVerb:
			q.Del(conf.key)
		}

		urls[i].RawQuery = q.Encode()

		fmt.Println(urls[i])
	}
}

func readURL() ([]*url.URL, error) {
	input := bufio.NewScanner(os.Stdin)
	ul := make([]*url.URL, 0, 1)

	for input.Scan() {
		u, err := url.Parse(input.Text())
		if err != nil {
			return nil, err
		}

		ul = append(ul, u)
	}

	return ul, nil
}
