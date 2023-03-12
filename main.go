package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/LintaoAmons/nric-go/nric"
	"github.com/atotto/clipboard"
	"github.com/urfave/cli/v2"
)

var NricAction = func(cCtx *cli.Context) error {
	var i int
	i, err := strconv.Atoi(cCtx.Args().Get(0))
	if err != nil {
		i = 1
	}
	count := i
	nricsString := nric.GenerateNricMultiple(count)
	clipboard.WriteAll(nricsString)
	fmt.Println(nricsString)
	return nil
}

var NricCommand = &cli.Command{
	Name:            "nric",
	Description:     "A tool to generate testing Singapore NRIC number, and it will automaticlly copy into you clipboard",
	HideHelpCommand: true,
	UsageText:       "nric [count] -- count is set to 1 if not passed",
	Usage:           "nric [count]",
	Action:          NricAction,
}

func main() {
	app := &cli.App{
		Name:            "nric",
		Description:     "A tool to generate testing Singapore NRIC number, and it will automaticlly copy into you clipboard",
		HideHelpCommand: true,
		UsageText:       "nric [count] -- count is set to 1 if not passed",
		Usage:           "nric [count]",
		Action:          NricAction,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
