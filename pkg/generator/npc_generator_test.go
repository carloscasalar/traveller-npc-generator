package generator_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"slices"
	"testing"

	"github.com/carloscasalar/traveller-npc-generator/pkg/generator"
	"github.com/stretchr/testify/assert"
)

func Test_resulting_character_should_have_same_category_experience_and_role_as_request(t *testing.T) {
	request := generator.NewGenerateCharacterRequestBuilder().
		CitizenCategory(generator.CitizenCategoryExceptional).
		Experience(generator.ExperienceRecruit).
		Role(generator.RoleDiplomat).
		Build()

	npcGenerator, _ := newGenerator()
	character, err := npcGenerator.Generate(*request)

	require.NoError(t, err)
	assert.Equal(t, generator.CitizenCategoryExceptional, character.CitizenCategory())
	assert.Equal(t, generator.ExperienceRecruit, character.Experience())
	assert.Equal(t, generator.RoleDiplomat, character.Role())
}

func Test_should_generate_an_npc_with_a_name(t *testing.T) {
	request := exceptionalRecruitDiplomatRequest()

	npcGenerator, _ := generator.NewNpcGeneratorBuilder().
		NameGenerator(NewFixedGenerateName("John", "Doe")).
		Build()
	character, err := npcGenerator.Generate(*request)

	require.NoError(t, err)
	assert.Equal(t, "John", character.FirstName())
	assert.Equal(t, "Doe", character.Surname())
}

func Test_should_generate_an_npc_name_even_when_no_name_generator_is_provided(t *testing.T) {
	request := exceptionalRecruitDiplomatRequest()

	npcGenerator, _ := newGenerator()
	character, err := npcGenerator.Generate(*request)

	require.NoError(t, err)
	assert.NotEmpty(t, character.FirstName())
	assert.NotEmpty(t, character.Surname())
}

func Test_should_generate_an_npc_with_skills(t *testing.T) {
	request := exceptionalRecruitDiplomatRequest()

	npcGenerator, _ := newGenerator()
	character, err := npcGenerator.Generate(*request)

	require.NoError(t, err)
	assert.NotEmpty(t, character.Skills())
}

func Test_number_of_skills_generated_for(t *testing.T) {
	allRoles := generator.RoleValues()
	expectedNumberOfSkillsByExperience := map[generator.Experience]int{
		generator.ExperienceRecruit:      4,
		generator.ExperienceRookie:       6,
		generator.ExperienceIntermediate: 7,
		generator.ExperienceRegular:      9,
		generator.ExperienceVeteran:      10,
		generator.ExperienceElite:        12,
	}
	for _, role := range allRoles {
		for _, experience := range generator.ExperienceValues() {
			t.Run(fmt.Sprintf("a %v %v should be %d", experience, role, expectedNumberOfSkillsByExperience[experience]), func(t *testing.T) {
				request := generator.NewGenerateCharacterRequestBuilder().
					Experience(experience).
					Role(role).
					Build()

				npcGenerator, _ := newGenerator()
				character, err := npcGenerator.Generate(*request)

				require.NoError(t, err)
				assert.Len(t, character.Skills(), expectedNumberOfSkillsByExperience[experience])
			})
		}
	}
}

func Test_characteristic_array_for_citizen_category(t *testing.T) {
	allCategories := generator.CitizenCategoryValues()
	expectedCharacteristicArray := map[generator.CitizenCategory][]int{
		generator.CitizenCategoryBelowAverage: {8, 7, 6, 6, 5, 4},
		generator.CitizenCategoryAverage:      {9, 8, 7, 7, 6, 5},
		generator.CitizenCategoryAboveAverage: {10, 9, 8, 8, 7, 6},
		generator.CitizenCategoryExceptional:  {11, 10, 9, 9, 8, 7},
	}

	for _, category := range allCategories {
		t.Run(fmt.Sprintf("%v should be %v", category, expectedCharacteristicArray[category]), func(t *testing.T) {
			request := generator.NewGenerateCharacterRequestBuilder().
				CitizenCategory(category).
				Build()

			npcGenerator, _ := newGenerator()
			character, err := npcGenerator.Generate(*request)

			require.NoError(t, err)
			assert.Equal(t, expectedCharacteristicArray[category], getSortedArray(character.Characteristics()))
		})
	}
}

func Test_when_category_is_invalid_it_returns_error(t *testing.T) {
	nonValidCategory := generator.CitizenCategory(99)
	request := generator.NewGenerateCharacterRequestBuilder().
		CitizenCategory(nonValidCategory).
		Build()

	npcGenerator, _ := newGenerator()
	_, err := npcGenerator.Generate(*request)

	assert.Error(t, err)
	assert.Equal(t, "invalid citizen category", err.Error())
}

func Test_when_experience_is_invalid_it_returns_error(t *testing.T) {
	nonValidExperience := generator.Experience(99)
	request := generator.NewGenerateCharacterRequestBuilder().
		Experience(nonValidExperience).
		Build()

	npcGenerator, _ := newGenerator()
	_, err := npcGenerator.Generate(*request)

	assert.Error(t, err)
	assert.Equal(t, "invalid experience", err.Error())
}

func Test_when_role_is_invalid_it_returns_error(t *testing.T) {
	nonValidRole := generator.Role(99)
	request := generator.NewGenerateCharacterRequestBuilder().
		Role(nonValidRole).
		Build()

	npcGenerator, _ := newGenerator()
	_, err := npcGenerator.Generate(*request)

	assert.Error(t, err)
	assert.Equal(t, "invalid role", err.Error())
}

func Test_when_gender_is_invalid_it_returns_error(t *testing.T) {
	nonValidGender := generator.Gender(99)
	request := generator.NewGenerateCharacterRequestBuilder().
		Gender(nonValidGender).
		Build()

	npcGenerator, _ := newGenerator()
	_, err := npcGenerator.Generate(*request)

	assert.Error(t, err)
	assert.Equal(t, "invalid gender", err.Error())
}

func getSortedArray(characteristics map[generator.Characteristic]int) []int {
	charArray := []int{
		characteristics[generator.STR],
		characteristics[generator.DEX],
		characteristics[generator.END],
		characteristics[generator.INT],
		characteristics[generator.EDU],
		characteristics[generator.SOC],
	}
	slices.Sort(charArray)
	slices.Reverse(charArray)
	return charArray
}

func exceptionalRecruitDiplomatRequest() *generator.GenerateCharacterRequest {
	return generator.NewGenerateCharacterRequestBuilder().
		CitizenCategory(generator.CitizenCategoryExceptional).
		Experience(generator.ExperienceRecruit).
		Role(generator.RoleDiplomat).
		Build()
}

func newGenerator() (*generator.NpcGenerator, error) {
	return generator.NewNpcGeneratorBuilder().Build()
}

type FixedNameGenerator struct {
	firstName string
	surname   string
}

func NewFixedGenerateName(firstName, surname string) *FixedNameGenerator {
	return &FixedNameGenerator{
		firstName: firstName,
		surname:   surname,
	}
}

func (g FixedNameGenerator) Generate(generator.Gender) (string, string) {
	return g.firstName, g.surname
}
