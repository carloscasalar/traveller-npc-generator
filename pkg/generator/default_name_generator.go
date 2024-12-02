package generator

import (
	"fmt"
	"github.com/carloscasalar/traveller-npc-generator/assets"
	"gopkg.in/yaml.v3"
)

func NewDefaultNameGenerator() (NameGenerator, error) {
	var config nameConfig
	if err := yaml.Unmarshal(assets.EmbedNames, &config); err != nil {
		return nil, fmt.Errorf("error unmarshalling names config file: %v", err)
	}
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("error parsing names config file: %v", err)
	}

	return NewCatalogSourcedNameGenerator(config.Surnames, config.NonGenderedNames, config.FemaNames, config.MaleNames), nil
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
