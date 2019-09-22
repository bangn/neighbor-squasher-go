package main

import (
	"os"

	"github.com/bangn/neighbor-squasher-go/cli"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	run(os.Args[1:], os.Exit)
}

func run(args []string, exit func(int)) {
	app := kingpin.New(`neighbor-squasher`, `Squash repeated char in a string`)

	app.Writer(os.Stdout)
	app.Version(cli.VersionString())
	app.Terminate(exit)

	// --------------------------
	//  global flags
}
