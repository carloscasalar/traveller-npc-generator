package main

import (
	"fmt"
	"github.com/carloscasalar/traveller-npc-generator/pkg/generator"
	"os"
)

func main() {
	npcGenerator, err := generator.NewNpcGeneratorBuilder().
		NameGenerator(new(CustomNameGenerator)).
		Build()
	if err != nil {
		fmt.Printf("Error creating NPC: %v", err)
		os.Exit(1)
	}

	for _, gender := range generator.GenderValues() {
		request := generator.NewGenerateCharacterRequestBuilder().
			Category(generator.CategoryExceptional).
			Experience(generator.ExperienceVeteran).
			Role(generator.RoleLeader).
			Gender(gender).
			Build()

		character, err := npcGenerator.Generate(*request)
		if err != nil {
			fmt.Printf("Error generating character: %v", err)
			os.Exit(1)
		}

		fmt.Printf("Generated Character: %v\n", character)
	}
}

type CustomNameGenerator struct {
}

func (c CustomNameGenerator) Generate(gender generator.Gender) (firstName, surname string) {
	switch gender {
	case generator.GenderMale:
		return "Dwayne", "Hicks"
	case generator.GenderFemale:
		return "Hellen", "Ripley"
	default:
		return "Forge", "Jynxori"
	}
}
