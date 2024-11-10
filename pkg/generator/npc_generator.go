package generator

import "errors"

type NpcGenerator struct{}

func NewNpcGenerator() *NpcGenerator {
	return &NpcGenerator{}
}

func (g *NpcGenerator) Generate(request GenerateCharacterRequest) (*Character, error) {
	if !request.category.IsACitizenCategory() {
		return nil, errors.New("invalid category")
	}
	if !request.experience.IsAExperience() {
		return nil, errors.New("invalid experience")
	}
	if !request.role.IsARole() {
		return nil, errors.New("invalid role")
	}
	if !request.gender.IsAGender() {
		return nil, errors.New("invalid gender")
	}

	return nil, errors.New("not implemented")
}
