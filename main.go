package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type encoder struct {
	WithPadding      bool
	URLEncodingFomat bool
	ShouldDecode     bool
}

func (e encoder) GetEncoder() *base64.Encoding {
	if e.URLEncodingFomat {
		if e.WithPadding {
			return base64.URLEncoding
		}
		return base64.RawURLEncoding
	}

	if e.WithPadding {
		return base64.StdEncoding
	}

	return base64.RawStdEncoding
}

func (e encoder) Transform(input string) (string, error) {
	if e.ShouldDecode {
		data, err := e.GetEncoder().DecodeString(input)
		if err != nil {
			return "", fmt.Errorf("can't decode input: %v", err)
		}

		return string(data), nil
	}

	return e.GetEncoder().EncodeToString([]byte(input)), nil
}

func run() error {
	var e encoder

	fs := flag.NewFlagSet("b64", flag.ExitOnError)
	fs.BoolVar(&e.ShouldDecode, "d", false, "decode the base 64 input")
	fs.BoolVar(&e.URLEncodingFomat, "u", false, "input/output follow base 64 URL encoded format")
	fs.BoolVar(&e.WithPadding, "p", true, "input/output are base 64 padded string")

	if err := fs.Parse(os.Args[1:]); err != nil {
		return err
	}

	input, err := getInput(os.Stdin, strings.TrimSpace(strings.Join(fs.Args(), " ")))
	if err != nil {
		return fmt.Errorf("can't get input data: %v", err)
	}

	data, err := e.Transform(strings.TrimSpace(input))
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, data)

	return nil
}

func getInput(f *os.File, arg string) (string, error) {
	fi, err := f.Stat()
	if err != nil {
		if arg == "" {
			return "", fmt.Errorf("args are empty and STDIN is not readabale: %v", err)
		}
		return arg, nil
	}

	size := fi.Size()
	if size == 0 && fi.Mode() == os.ModeCharDevice {
		if arg == "" {
			return "", fmt.Errorf("args and STDIN are empty")
		}

		return arg, nil
	}

	input, err := ioutil.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("can't read from STDIN: %v", err)
	}

	if arg != "" {
		return "", fmt.Errorf("args and STDIN are both set: choose only one")
	}

	return string(input), nil
}
