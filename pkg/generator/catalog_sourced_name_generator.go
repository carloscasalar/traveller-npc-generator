package generator

import "math/rand/v2"

type catalogSourcedNameGenerator struct {
	namesByGender map[Gender][]string
	surnames      []string
}

// NewCatalogSourcedNameGenerator creates a new NameGenerator that generates names from a catalog of names by gender and a list of surnames.
func NewCatalogSourcedNameGenerator(surnames, nonGenderNames, femaleFirstNames, maleFirstNames []string) NameGenerator {
	return &catalogSourcedNameGenerator{
		namesByGender: map[Gender][]string{
			GenderUnspecified: nonGenderNames,
			GenderFemale:      femaleFirstNames,
			GenderMale:        maleFirstNames,
		},
		surnames: surnames,
	}
}

func (c catalogSourcedNameGenerator) Generate(gender Gender) (firstName, surname string) {
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
