package generator_test

import (
	"fmt"
	"github.com/carloscasalar/traveller-npc-generator/pkg/generator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCatalogSourcedGenerator_Generate_should_generate_names_using_the_provided_lists(t *testing.T) {
	surnames := []string{"Surname 1", "Surname 2", "Surname 3"}
	nonGenderNames := []string{"Non Gendered Name 1", "Non Gendered Name 2", "Non Gendered Name 3"}
	femaleFirstNames := []string{"Female Name 1", "Female Name 2", "Female Name 3"}
	maleFirstNames := []string{"Male Name 1", "Male Name 2", "Male Name 3"}

	nameGenerator, _ := generator.NewCatalogSourcedNameGenerator(surnames, nonGenderNames, femaleFirstNames, maleFirstNames)

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

func TestCatalogSourcedGenerator_Generate_should_fail_when(t *testing.T) {
	var emptyList []string = nil
	nonEmptyList := []string{"Ripley", "Hicks", "Newt"}
	testCases := map[string]struct {
		surnames         []string
		nonGenderNames   []string
		femaleFirstNames []string
		maleFirstNames   []string
		expectedErrorMsg string
	}{
		"surnames list is empty":         {emptyList, nonEmptyList, nonEmptyList, nonEmptyList, "invalid list surnames, a non-empty list must be provided"},
		"nonGenderNames list is empty":   {nonEmptyList, emptyList, nonEmptyList, nonEmptyList, "invalid list nonGenderNames, a non-empty list must be provided"},
		"femaleFirstNames list is empty": {nonEmptyList, nonEmptyList, emptyList, nonEmptyList, "invalid list femaleFirstNames, a non-empty list must be provided"},
		"maleFirstNames list is empty":   {nonEmptyList, nonEmptyList, nonEmptyList, emptyList, "invalid list maleFirstNames, a non-empty list must be provided"},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			_, err := generator.NewCatalogSourcedNameGenerator(tc.surnames, tc.nonGenderNames, tc.femaleFirstNames, tc.maleFirstNames)

			require.Error(t, err)
			assert.EqualError(t, err, tc.expectedErrorMsg)
			assert.ErrorAs(t, err, &generator.InvalidListError{})
		})
	}
}
