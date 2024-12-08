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

	return NewCatalogSourcedNameGenerator(config.Surnames, config.NonGenderedNames, config.FemaNames, config.MaleNames)
}

type nameConfig struct {
	Surnames         []string `yaml:",flow"`
	NonGenderedNames []string `yaml:"non_gendered_names,flow"`
	FemaNames        []string `yaml:"female_names,flow"`
	MaleNames        []string `yaml:"male_names,flow"`
}
