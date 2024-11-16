package generator

type GenerateName interface {
	Execute(gender Gender) (firstName, surname string)
}
