package name

type Generator interface {
	Generate(gender Gender) (firstName, surname string)
}
