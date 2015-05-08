package main

import (
	"flag"
	"fmt"
	"os"
)

// potential modes to implement
//   encrypt small data
//   encrypt random token
//   show rsa numbers

func mainCommandUsage() {
	// if exitcode < 0, do not exit
	fmt.Fprintf(os.Stderr, "usage: %s <subcommand>\n\nsubcommands:\n\ndecrypt\ntoken\n\n%s <subcommand> -h for more help\n", os.Args[0], os.Args[0])
	os.Exit(1)
}

func main() {
	flag.Usage = mainCommandUsage
	flag.Parse()

	if flag.NArg() < 1 {
		mainCommandUsage()
	}

	switch flag.Arg(0) {
	case "token":
		tokenCommand(flag.Args()[1:])
	case "decrypt":
		decryptCommand(flag.Args()[1:])
	default:
		mainCommandUsage()
	}

}
