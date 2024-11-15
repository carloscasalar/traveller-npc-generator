package generator

type NpcGenerator struct{}

func NewNpcGenerator() *NpcGenerator {
	return &NpcGenerator{}
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
