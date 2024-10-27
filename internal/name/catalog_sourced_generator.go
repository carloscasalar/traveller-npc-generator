package name

import "math/rand/v2"

type catalogSourcedGenerator struct {
	namesByGender map[Gender][]string
	surnames      []string
}

func NewCatalogSourcedGenerator(surnames, nonGenderNames, femaleFirstNames, maleFirstNames []string) Generator {
	return catalogSourcedGenerator{
		namesByGender: map[Gender][]string{
			GenderUnspecified: nonGenderNames,
			GenderFemale:      femaleFirstNames,
			GenderMale:        maleFirstNames,
		},
		surnames: surnames,
	}
}

func (c catalogSourcedGenerator) Generate(gender Gender) (FirstName, Surname) {
	if !gender.IsAGender() {
		gender = pickRandomItem(GenderValues())
	}

	firstName := FirstName(pickRandomItem(c.namesByGender[gender]))
	surname := Surname(pickRandomItem(c.surnames))

	return firstName, surname
}

func pickRandomItem[T any](items []T) T {
	itemIndex := rand.IntN(len(items) - 1)
	return items[itemIndex]
}
