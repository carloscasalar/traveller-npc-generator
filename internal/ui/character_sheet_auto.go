// Code generated by gonstructor -type=CharacterSheet -constructorTypes=builder -output=character_sheet_auto.go; DO NOT EDIT.

package ui

type CharacterSheetBuilder struct {
	fullName        string
	role            string
	citizenCategory string
	experience      string
	skills          []string
	characteristics map[Characteristic]int
}

func NewCharacterSheetBuilder() *CharacterSheetBuilder {
	return &CharacterSheetBuilder{}
}

func (b *CharacterSheetBuilder) FullName(fullName string) *CharacterSheetBuilder {
	b.fullName = fullName
	return b
}

func (b *CharacterSheetBuilder) Role(role string) *CharacterSheetBuilder {
	b.role = role
	return b
}

func (b *CharacterSheetBuilder) CitizenCategory(citizenCategory string) *CharacterSheetBuilder {
	b.citizenCategory = citizenCategory
	return b
}

func (b *CharacterSheetBuilder) Experience(experience string) *CharacterSheetBuilder {
	b.experience = experience
	return b
}

func (b *CharacterSheetBuilder) Skills(skills []string) *CharacterSheetBuilder {
	b.skills = skills
	return b
}

func (b *CharacterSheetBuilder) Characteristics(characteristics map[Characteristic]int) *CharacterSheetBuilder {
	b.characteristics = characteristics
	return b
}

func (b *CharacterSheetBuilder) Build() *CharacterSheet {
	return &CharacterSheet{
		fullName:        b.fullName,
		role:            b.role,
		citizenCategory: b.citizenCategory,
		experience:      b.experience,
		skills:          b.skills,
		characteristics: b.characteristics,
	}
}
