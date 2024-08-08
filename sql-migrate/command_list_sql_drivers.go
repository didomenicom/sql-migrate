package main

import (
	"database/sql"
	"fmt"
	"strings"
)

type ListSqlDriversCommand struct{}

func (*ListSqlDriversCommand) Help() string {
	helpText := `
Usage: sql-migrate list-sql-drivers ...

  Show the available sql connection drivers

Options:

  -config=dbconfig.yml   Configuration file to use.
  -env="development"     Environment.

`
	return strings.TrimSpace(helpText)
}

func (*ListSqlDriversCommand) Synopsis() string {
	return "Show available sql connection drivers"
}

func (c *ListSqlDriversCommand) Run(args []string) int {
	fmt.Printf(":: SQL Drivers ::\n%s\n", sql.Drivers())

	return 0
}
