package name

type FirstName string
type Surname string

type Generator interface {
	Generate(gender Gender) (FirstName, Surname)
}
