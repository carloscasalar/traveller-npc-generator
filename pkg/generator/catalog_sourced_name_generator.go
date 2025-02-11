package generator

import "math/rand/v2"

type catalogSourcedNameGenerator struct {
	namesByGender map[Gender][]string
	surnames      []string
}

// NewCatalogSourcedNameGenerator creates a new NameGenerator that generates names from a catalog of names by gender and a list of surnames.
func NewCatalogSourcedNameGenerator(surnames, nonGenderNames, femaleFirstNames, maleFirstNames []string) (NameGenerator, error) {
	if len(surnames) == 0 {
		return nil, newInvalidListError("surnames")
	}
	if len(nonGenderNames) == 0 {
		return nil, newInvalidListError("nonGenderNames")
	}
	if len(femaleFirstNames) == 0 {
		return nil, newInvalidListError("femaleFirstNames")
	}
	if len(maleFirstNames) == 0 {
		return nil, newInvalidListError("maleFirstNames")
	}
	return &catalogSourcedNameGenerator{
		namesByGender: map[Gender][]string{
			GenderUnspecified: nonGenderNames,
			GenderFemale:      femaleFirstNames,
			GenderMale:        maleFirstNames,
		},
		surnames: surnames,
	}, nil
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
