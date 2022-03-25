package io

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/fatih/color"
)

var (
	keyColor   = color.New(color.FgCyan)
	valueColor = color.New(color.FgYellow)
)

func keyValueln(key, value string) {
	keyColor.Printf(key)
	fmt.Printf(": ")
	valueColor.Printf(value)
	fmt.Printf("\n")
}

func indent(hasPrinted bool) {
	if !hasPrinted {
		fmt.Printf("- ")
	} else {
		fmt.Printf("  ")
	}
}

func indentDepth(hasPrinted bool, n int) {
	indent := "  "

	for ; n > 1; n-- {
		fmt.Printf(indent)
	}

	if !hasPrinted {
		fmt.Printf("- ")
	} else {
		fmt.Printf(indent)
	}
}

func PrettyURL(u *url.URL) {
	var hasPrinted bool

	if u.Scheme != "" {
		indent(hasPrinted)
		hasPrinted = true

		keyValueln("scheme", u.Scheme)
	}

	if u.Host != "" {
		indent(hasPrinted)
		hasPrinted = true

		keyValueln("hostname", u.Host)
	}

	if u.Path != "" {
		indent(hasPrinted)
		hasPrinted = true

		keyValueln("path", u.Path)
	}

	if u.Fragment != "" {
		indent(hasPrinted)
		hasPrinted = true

		keyValueln("fragment", u.Fragment)
	}

	if q := u.Query(); len(q) > 0 {
		indent(hasPrinted)
		hasPrinted = true

		keyValueln("query", "")
		for k, v := range q {
			keyValueln(fmt.Sprintf("    %s", k), strings.Join(v, ", "))
		}
	}
}
