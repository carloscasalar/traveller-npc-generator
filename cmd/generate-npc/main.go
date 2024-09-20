package main

import (
	"fmt"
	"github.com/carloscasalar/traveller-npc-generator/internal/ui"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	opts := readOptionsOrFail()
	prompt("CitizenCategory: " + opts.CitizenCategory)
	prompt("Experience: " + opts.Experience)
	prompt("CrewRole: " + opts.CrewRole)
}

func readOptionsOrFail() CommandOptions {
	var opts CommandOptions
	parser := flags.NewParser(&opts, flags.Default)
	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
	return opts
}

type CommandOptions struct {
	CitizenCategory string `short:"c" default:"1" long:"category" choice:"0" choice:"1" choice:"2" choice:"3" description:"Citizen Category: 0-Below average, 1-Average, 2-Above Average, 3 Exceptional" required:"true"`
	Experience      string `short:"e" default:"3" long:"experience" choice:"0" choice:"1" choice:"2" choice:"3" choice:"4" choice:"5" description:"Experience: 0-Recruit, 1-Rookie, 2-Intermediate, 3-Regular, 4-Veteran, 5-Elite" required:"true"`
	CrewRole        string `short:"r" long:"role" required:"true" choice:"pilot" choice:"navigator" choice:"engineer" choice:"steward" choice:"medic" choice:"marine" choice:"gunner" choice:"scout" choice:"technician" choice:"leader" choice:"diplomat" choice:"entertainer" choice:"trader" choice:"thug" description:"Crew role in a starship"`
}

func prompt(value string) {
	fmt.Println(ui.NewPromptRenderer(value).Render())
}
