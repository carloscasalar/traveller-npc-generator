package generator

import "errors"

type NpcGenerator struct{}

func NewNpcGenerator() *NpcGenerator {
	return &NpcGenerator{}
}

func (g *NpcGenerator) Generate(request GenerateCharacterRequest) (*Character, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return nil, errors.New("not implemented")
}
