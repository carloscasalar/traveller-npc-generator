package generator_test

import (
	"testing"

	"github.com/carloscasalar/traveller-npc-generator/pkg/generator"
	"github.com/stretchr/testify/assert"
)

func Test_when_category_is_invalid_it_returns_error(t *testing.T) {
	nonValidCategory := generator.CitizenCategory(99)
	request := generator.NewGenerateCharacterRequestBuilder().
		Category(nonValidCategory).
		Build()

	npcGenerator := generator.NewNpcGenerator()
	_, err := npcGenerator.Generate(*request)

	assert.Error(t, err)
	assert.Equal(t, "invalid category", err.Error())
}

func Test_when_experience_is_invalid_it_returns_error(t *testing.T) {
	nonValidExperience := generator.Experience(99)
	request := generator.NewGenerateCharacterRequestBuilder().
		Experience(nonValidExperience).
		Build()

	npcGenerator := generator.NewNpcGenerator()
	_, err := npcGenerator.Generate(*request)

	assert.Error(t, err)
	assert.Equal(t, "invalid experience", err.Error())
}

func Test_when_role_is_invalid_it_returns_error(t *testing.T) {
	nonValidRole := generator.Role(99)
	request := generator.NewGenerateCharacterRequestBuilder().
		Role(nonValidRole).
		Build()

	npcGenerator := generator.NewNpcGenerator()
	_, err := npcGenerator.Generate(*request)

	assert.Error(t, err)
	assert.Equal(t, "invalid role", err.Error())
}

func Test_when_gender_is_invalid_it_returns_error(t *testing.T) {
	nonValidGender := generator.Gender(99)
	request := generator.NewGenerateCharacterRequestBuilder().
		Gender(nonValidGender).
		Build()

	npcGenerator := generator.NewNpcGenerator()
	_, err := npcGenerator.Generate(*request)

	assert.Error(t, err)
	assert.Equal(t, "invalid gender", err.Error())
}
