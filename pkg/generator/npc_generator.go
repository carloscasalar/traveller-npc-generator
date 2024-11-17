package generator

import (
	"github.com/carloscasalar/traveller-npc-generator/internal/npc"
)

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
	category := request.category.toNpcCitizenCategory()
	characteristic := role.RandomCharacteristic(category)

	return &Character{
		FirstName:       firstName,
		Surname:         surname,
		Role:            request.role,
		Category:        request.category,
		Experience:      request.experience,
		Skills:          skills,
		Characteristics: toCharacteristics(characteristic),
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

func toCharacteristics(characteristic map[npc.Characteristic]int) map[Characteristic]int {
	return map[Characteristic]int{
		STR: characteristic[npc.STR],
		DEX: characteristic[npc.DEX],
		END: characteristic[npc.END],
		INT: characteristic[npc.INT],
		EDU: characteristic[npc.EDU],
		SOC: characteristic[npc.SOC],
	}
}
