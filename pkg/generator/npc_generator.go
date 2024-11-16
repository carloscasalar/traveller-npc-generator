package generator

//go:generate gonstructor -type=NpcGenerator -constructorTypes=builder -init=init -propagateInitFuncReturns -output=npc_generator_auto.go
type NpcGenerator struct {
	generateName GenerateName
}

func (g *NpcGenerator) Generate(request GenerateCharacterRequest) (*Character, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return &Character{
		Role:       request.role,
		Category:   request.category,
		Experience: request.experience,
	}, nil
}

func (g *NpcGenerator) init() error {
	if g.generateName != nil {
		return nil
	}
	generateName, err := NewDefaultGenerateName()
	if err != nil {
		return err
	}
	g.generateName = generateName
	return nil
}
