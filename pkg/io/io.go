package io

import (
	"bufio"
	nativeIO "io"
)

func ActOn(i nativeIO.Reader, action func(string) error) error {
	input := bufio.NewScanner(i)

	for input.Scan() {
		if err := action(input.Text()); err != nil {
			return err
		}
	}

	return input.Err()
}
