package main

import (
	"fmt"
	"github.com/carloscasalar/traveller-npc-generator/pkg/generator"
	"math/rand/v2"
	"os"
)

func main() {
	femaleNames := []string{"Hellen", "Jane", "Alice"}
	maleNames := []string{"Dwayne", "John", "Bob"}
	nonBinaryNames := []string{"Forge", "Jynxori", "Alex"}
	surnames := []string{"Hicks", "Doe", "Smith"}

	catalogNameGenerator, err := generator.NewCatalogSourcedNameGenerator(surnames, nonBinaryNames, femaleNames, maleNames)
	if err != nil {
		fmt.Printf("Error creating catalog sourced name generator: %v", err)
		os.Exit(1)
	}
	npcGenerator, err := generator.NewNpcGeneratorBuilder().NameGenerator(catalogNameGenerator).Build()
	if err != nil {
		fmt.Printf("Error creating NPC generator: %v", err)
		os.Exit(1)
	}

	for _, gender := range generator.GenderValues() {
		category := pickRandomItem(generator.CitizenCategoryValues())
		experience := pickRandomItem(generator.ExperienceValues())
		role := pickRandomItem(generator.RoleValues())

		request := generator.NewGenerateCharacterRequestBuilder().
			CitizenCategory(category).
			Experience(experience).
			Role(role).
			Gender(gender).
			Build()

		character, err := npcGenerator.Generate(*request)
		if err != nil {
			fmt.Printf("Error generating character: %v", err)
			os.Exit(1)
		}

		fmt.Printf("Generated character: %v\n", character)
	}
}

func pickRandomItem[T any](items []T) T {
	itemIndex := rand.IntN(len(items) - 1)
	return items[itemIndex]
}
