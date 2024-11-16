package generator_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/carloscasalar/traveller-npc-generator/pkg/generator"
	"github.com/stretchr/testify/assert"
)

func Test_resulting_character_should_have_same_category_experience_and_role_as_request(t *testing.T) {
	request := generator.NewGenerateCharacterRequestBuilder().
		Category(generator.CategoryExceptional).
		Experience(generator.ExperienceRecruit).
		Role(generator.RoleDiplomat).
		Build()

	npcGenerator, _ := generator.NewNpcGenerator()
	character, err := npcGenerator.Generate(*request)

	require.NoError(t, err)
	assert.Equal(t, generator.CategoryExceptional, character.Category)
	assert.Equal(t, generator.ExperienceRecruit, character.Experience)
	assert.Equal(t, generator.RoleDiplomat, character.Role)
}

func Test_when_category_is_invalid_it_returns_error(t *testing.T) {
	nonValidCategory := generator.CitizenCategory(99)
	request := generator.NewGenerateCharacterRequestBuilder().
		Category(nonValidCategory).
		Build()

	npcGenerator, _ := generator.NewNpcGenerator()
	_, err := npcGenerator.Generate(*request)

	assert.Error(t, err)
	assert.Equal(t, "invalid category", err.Error())
}

func Test_when_experience_is_invalid_it_returns_error(t *testing.T) {
	nonValidExperience := generator.Experience(99)
	request := generator.NewGenerateCharacterRequestBuilder().
		Experience(nonValidExperience).
		Build()

	npcGenerator, _ := generator.NewNpcGenerator()
	_, err := npcGenerator.Generate(*request)

	assert.Error(t, err)
	assert.Equal(t, "invalid experience", err.Error())
}

func Test_when_role_is_invalid_it_returns_error(t *testing.T) {
	nonValidRole := generator.Role(99)
	request := generator.NewGenerateCharacterRequestBuilder().
		Role(nonValidRole).
		Build()

	npcGenerator, _ := generator.NewNpcGenerator()
	_, err := npcGenerator.Generate(*request)

	assert.Error(t, err)
	assert.Equal(t, "invalid role", err.Error())
}

func Test_when_gender_is_invalid_it_returns_error(t *testing.T) {
	nonValidGender := generator.Gender(99)
	request := generator.NewGenerateCharacterRequestBuilder().
		Gender(nonValidGender).
		Build()

	npcGenerator, _ := generator.NewNpcGenerator()
	_, err := npcGenerator.Generate(*request)

	assert.Error(t, err)
	assert.Equal(t, "invalid gender", err.Error())
}
