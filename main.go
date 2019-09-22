package main

import (
	"fmt"
	"os"

	"github.com/bangn/neighbor-squasher-go/cli"
	"github.com/fatih/color"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	run(os.Args[1:], os.Exit)
}

func run(args []string, exit func(int)) {
	app := kingpin.New(
		`neighbor-squasher`,
		`Squash repeated char in a string`,
	)

	app.Writer(os.Stdout)
	app.Version(cli.VersionString())
	app.Terminate(exit)

	// --------------------------
	//  global flags
	var (
		squashBy string // TODO: should be rune
		input    string
	)

	app.Flag("squash-by", "Character (UTF-8) to be squashed").
		StringVar(&squashBy)

	app.Flag("input", "input string").
		StringVar(&input)

	// --------------------------
	// run the app, parse args

	if _, err := app.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, color.RedString("🚨 %v\n", err))

		if ec, ok := err.(interface{ ExitCode() int }); ok {
			os.Exit(ec.ExitCode())
		} else {
			os.Exit(1)
		}
	}
}
