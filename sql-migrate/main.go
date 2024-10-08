package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	os.Exit(realMain())
}

var ui cli.Ui

func realMain() int {
	ui = &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr}

	cli := &cli.CLI{
		Args: os.Args[1:],
		Commands: map[string]cli.CommandFactory{
			"up": func() (cli.Command, error) {
				return &UpCommand{}, nil
			},
			"down": func() (cli.Command, error) {
				return &DownCommand{}, nil
			},
			"redo": func() (cli.Command, error) {
				return &RedoCommand{}, nil
			},
			"status": func() (cli.Command, error) {
				return &StatusCommand{}, nil
			},
			"new": func() (cli.Command, error) {
				return &NewCommand{}, nil
			},
			"skip": func() (cli.Command, error) {
				return &SkipCommand{}, nil
			},
			"list-sql-drivers": func() (cli.Command, error) {
				return &ListSqlDriversCommand{}, nil
			},
		},
		HelpFunc:    cli.BasicHelpFunc("sql-migrate"),
		HelpWriter:  os.Stdout,
		ErrorWriter: os.Stderr,
		Version:     GetVersion(),
	}

	exitCode, err := cli.Run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}
