package generator_test

import (
	"fmt"
	"github.com/carloscasalar/traveller-npc-generator/pkg/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCatalogSourcedGenerator_Generate(t *testing.T) {
	surnames := []string{"Surname 1", "Surname 2", "Surname 3"}
	nonGenderNames := []string{"Non Gendered Name 1", "Non Gendered Name 2", "Non Gendered Name 3"}
	femaleFirstNames := []string{"Female Name 1", "Female Name 2", "Female Name 3"}
	maleFirstNames := []string{"Male Name 1", "Male Name 2", "Male Name 3"}

	nameGenerator := generator.NewCatalogSourcedNameGenerator(surnames, nonGenderNames, femaleFirstNames, maleFirstNames)

	tests := []struct {
		gender                 generator.Gender
		expectedToBeChosenFrom []string
	}{
		{gender: generator.GenderUnspecified, expectedToBeChosenFrom: nonGenderNames},
		{gender: generator.GenderFemale, expectedToBeChosenFrom: femaleFirstNames},
		{gender: generator.GenderMale, expectedToBeChosenFrom: maleFirstNames},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Given gender is %v", tt.gender.String()), func(t *testing.T) {
			firstName, surname := nameGenerator.Generate(tt.gender)

			assert.Contains(t, surnames, string(surname))
			assert.Contains(t, tt.expectedToBeChosenFrom, string(firstName))
		})
	}
}
