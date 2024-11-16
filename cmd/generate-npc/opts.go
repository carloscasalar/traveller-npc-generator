package main

import (
	"github.com/carloscasalar/traveller-npc-generator/internal/npc"
	"github.com/carloscasalar/traveller-npc-generator/pkg/generator"
	"github.com/jessevdk/go-flags"
	"os"
)

type CommandOptions struct {
	CitizenCategory int    `short:"c" default:"1" long:"category" choice:"0" choice:"1" choice:"2" choice:"3" description:"Citizen Category: 0-Below average, 1-Average, 2-Above Average, 3-Exceptional" required:"true"`
	Experience      int    `short:"e" default:"3" long:"experience" choice:"0" choice:"1" choice:"2" choice:"3" choice:"4" choice:"5" description:"Experience: 0-Recruit, 1-Rookie, 2-Intermediate, 3-Regular, 4-Veteran, 5-Elite" required:"true"`
	CrewRole        string `short:"r" long:"role" required:"true" choice:"pilot" choice:"navigator" choice:"engineer" choice:"steward" choice:"medic" choice:"marine" choice:"gunner" choice:"scout" choice:"technician" choice:"leader" choice:"diplomat" choice:"entertainer" choice:"trader" choice:"thug" description:"Crew role in a starship"`
	Gender          string `short:"g" long:"gender" default:"unspecified" choice:"female" choice:"male" choice:"unspecified" description:"Gender of the NPC"`
	EnableDebug     bool   `short:"d" long:"debug" description:"Enable debug mode"`
}

func readCitizenCategory(opts CommandOptions) npc.CitizenCategory {
	category := npc.CitizenCategory(opts.CitizenCategory)
	if !category.IsACitizenCategory() {
		printErrorf("unknown citizen category %v\n", opts.CitizenCategory)
		os.Exit(1)
	}
	return category
}

func readExperience(opts CommandOptions) npc.Experience {
	if !npc.Experience(opts.Experience).IsAExperience() {
		printErrorf("unknown experience %v\n", opts.Experience)
		os.Exit(1)
	}

	return npc.Experience(opts.Experience)
}

func readRole(opts CommandOptions) npc.Role {
	role, err := npc.RoleString(opts.CrewRole)
	if err != nil {
		printErrorf("unknown role %v\n", opts.CrewRole)
		os.Exit(1)
	}
	return role
}

func readGender(opts CommandOptions) generator.Gender {
	switch opts.Gender {
	case "female":
		return generator.GenderFemale
	case "male":
		return generator.GenderMale
	default:
		return generator.GenderUnspecified
	}
}

func readOptionsOrFail() CommandOptions {
	var opts CommandOptions
	parser := flags.NewParser(&opts, flags.Default)
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(flags.ErrorType); ok && flagsErr == flags.ErrHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}
	return opts
}
