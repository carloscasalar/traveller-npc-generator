package generator

import (
	"fmt"
	"github.com/carloscasalar/traveller-npc-generator/internal/name"
	"gopkg.in/yaml.v3"
	"os"
)

type NameGenerator interface {
	Generate(gender Gender) (firstName, surname string)
}

type defaultNameGenerator struct {
	nameGenerator name.Generator
}

func (d *defaultNameGenerator) Generate(gender Gender) (firstName, surname string) {
	return d.nameGenerator.Generate(gender.toNpcGender())
}

func NewDefaultNameGenerator() (NameGenerator, error) {
	const namesFilePath = "../../assets/names.yml"
	namesConfigFile, err := os.ReadFile(namesFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading names config file at %v: %v\n", namesFilePath, err)
	}
	var config nameConfig
	err = yaml.Unmarshal(namesConfigFile, &config)
	if err != nil {
		return nil, fmt.Errorf("error reading names config file at %v: %v\n", namesFilePath, err)
	}
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("error parsing names config file at %v: %v\n", namesFilePath, err)
	}

	return &defaultNameGenerator{
		name.NewCatalogSourcedGenerator(config.Surnames, config.NonGenderedNames, config.FemaNames, config.MaleNames),
	}, nil
}

type nameConfig struct {
	Surnames         []string `yaml:",flow"`
	NonGenderedNames []string `yaml:"non_gendered_names,flow"`
	FemaNames        []string `yaml:"female_names,flow"`
	MaleNames        []string `yaml:"male_names,flow"`
}

func (c nameConfig) Validate() error {
	if len(c.Surnames) == 0 {
		return fmt.Errorf("surnames cannot be empty")
	}
	if len(c.NonGenderedNames) == 0 {
		return fmt.Errorf("non_gendered_names cannot be empty")
	}
	if len(c.FemaNames) == 0 {
		return fmt.Errorf("female_names cannot be empty")
	}
	if len(c.MaleNames) == 0 {
		return fmt.Errorf("male_names cannot be empty")
	}

	return nil
}
