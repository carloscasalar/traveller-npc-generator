package main

import (
	"fmt"
	"github.com/carloscasalar/traveller-npc-generator/pkg/generator"
	"os"
)

func main() {
	npcGenerator, err := generator.NewNpcGeneratorBuilder().Build()
	if err != nil {
		fmt.Printf("Error creating NPC generator: %v", err)
		os.Exit(1)
	}

	request := generator.NewGenerateCharacterRequestBuilder().
		Category(generator.CategoryAboveAverage).
		Experience(generator.ExperienceRookie).
		Role(generator.RolePilot).
		Gender(generator.GenderUnspecified).
		Build()

	character, err := npcGenerator.Generate(*request)
	if err != nil {
		fmt.Printf("Error generating character: %v", err)
		os.Exit(1)
	}

	fmt.Println("Generated Character:", character)
}
