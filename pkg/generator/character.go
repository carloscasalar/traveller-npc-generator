package generator

type Character struct {
	FirstName  string
	Surname    string
	Role       Role
	Category   CitizenCategory
	Experience Experience
	Skills     map[string]int
}
