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
	role := request.role.toNpcRole()
	experience := request.experience.toNpcExperience()
	skills := role.Skills(experience)
	category := request.citizenCategory.toNpcCitizenCategory()
	characteristic := role.RandomCharacteristic(category)

	return &Character{
		firstName:       firstName,
		surname:         surname,
		role:            request.role,
		citizenCategory: request.citizenCategory,
		experience:      request.experience,
		skills:          skills,
		characteristics: characteristic,
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
