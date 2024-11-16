package generator

type NameGenerator interface {
	Generate(gender Gender) (firstName, surname string)
}
