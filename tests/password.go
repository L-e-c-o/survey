//go:build ignore

package main

import (
	"github.com/L-e-c-o/survey/v2"
	TestUtil "github.com/L-e-c-o/survey/v2/tests/util"
)

var value = ""

var table = []TestUtil.TestTableEntry{
	{
		"standard", &survey.Password{Message: "Please type your password:"}, &value, nil,
	},
	{
		"please make sure paste works", &survey.Password{Message: "Please paste your password:"}, &value, nil,
	},
	{
		"no help, send '?'", &survey.Password{Message: "Please type your password:"}, &value, nil,
	},
}

func main() {
	TestUtil.RunTable(table)
}
