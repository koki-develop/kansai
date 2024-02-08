package util

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/term"
)

func ReadPassword(w io.Writer, prompt string) (string, error) {
	if _, err := fmt.Fprintf(w, "%s: ", prompt); err != nil {
		return "", err
	}

	k, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}

	if _, err := fmt.Fprintln(w); err != nil {
		return "", err
	}

	return string(k), nil
}
