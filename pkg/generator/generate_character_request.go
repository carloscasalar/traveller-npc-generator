package generator

import "errors"

//go:generate gonstructor -type=GenerateCharacterRequest -constructorTypes=builder -output=generate_character_request_auto.go

type GenerateCharacterRequest struct {
	category   CitizenCategory
	experience Experience
	role       Role
	gender     Gender
}

func (r GenerateCharacterRequest) Validate() error {
	if !r.category.IsACitizenCategory() {
		return errors.New("invalid category")
	}
	if !r.experience.IsAExperience() {
		return errors.New("invalid experience")
	}
	if !r.role.IsARole() {
		return errors.New("invalid role")
	}
	if !r.gender.IsAGender() {
		return errors.New("invalid gender")
	}
	return nil
}
