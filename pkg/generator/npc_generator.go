package generator

//go:generate gonstructor -type=NpcGenerator -constructorTypes=builder -init=init -propagateInitFuncReturns -output=npc_generator_auto.go
type NpcGenerator struct {
	nameGenerator NameGenerator
}

func (g *NpcGenerator) Generate(request GenerateCharacterRequest) (*Character, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	firstName, surname := g.nameGenerator.Generate(request.gender)

	return &Character{
		FirstName:  firstName,
		Surname:    surname,
		Role:       request.role,
		Category:   request.category,
		Experience: request.experience,
	}, nil
}

func (g *NpcGenerator) init() error {
	if g.nameGenerator != nil {
		return nil
	}
	nameGenerator, err := NewDefaultNameGenerator()
	if err != nil {
		return err
	}
	g.nameGenerator = nameGenerator
	return nil
}
