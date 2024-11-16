package generator

import "math/rand/v2"

type generateCatalogSourcedName struct {
	namesByGender map[Gender][]string
	surnames      []string
}

func NewGenerateCatalogSourcedName(surnames, nonGenderNames, femaleFirstNames, maleFirstNames []string) GenerateName {
	return &generateCatalogSourcedName{
		namesByGender: map[Gender][]string{
			GenderUnspecified: nonGenderNames,
			GenderFemale:      femaleFirstNames,
			GenderMale:        maleFirstNames,
		},
		surnames: surnames,
	}
}

func (c generateCatalogSourcedName) Execute(gender Gender) (firstName, surname string) {
	if !gender.IsAGender() {
		gender = pickRandomItem(GenderValues())
	}

	firstName = pickRandomItem(c.namesByGender[gender])
	surname = pickRandomItem(c.surnames)

	return
}

func pickRandomItem[T any](items []T) T {
	itemIndex := rand.IntN(len(items) - 1)
	return items[itemIndex]
}
