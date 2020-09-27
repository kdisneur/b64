package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kdisneur/b64/internal"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var encoder internal.Encoder
	var showVersion bool

	fs := flag.NewFlagSet("b64", flag.ExitOnError)
	fs.BoolVar(&encoder.ShouldDecode, "d", false, "decode the base 64 input")
	fs.BoolVar(&encoder.URLEncodingFomat, "u", false, "input/output follow base 64 URL encoded format")
	fs.BoolVar(&encoder.WithPadding, "p", true, "input/output are base 64 padded string")
	fs.BoolVar(&showVersion, "v", false, "displays the current version")

	if err := fs.Parse(os.Args[1:]); err != nil {
		return err
	}

	if showVersion {
		fmt.Println(internal.GetVersionInfo().String())
		return nil
	}

	input, err := getInput(os.Stdin, strings.TrimSpace(strings.Join(fs.Args(), " ")))
	if err != nil {
		return fmt.Errorf("can't get input data: %v", err)
	}

	data, err := encoder.Transform(strings.TrimSpace(input))
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
