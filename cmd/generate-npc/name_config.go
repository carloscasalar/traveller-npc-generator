package main

import (
	"fmt"
	"github.com/carloscasalar/traveller-npc-generator/internal/name"
	"gopkg.in/yaml.v3"
	"os"
)

type NameConfig struct {
	Surnames         []string `yaml:",flow"`
	NonGenderedNames []string `yaml:"non_gendered_names,flow"`
	FemaNames        []string `yaml:"female_names,flow"`
	MaleNames        []string `yaml:"male_names,flow"`
}

func (c NameConfig) Validate() error {
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

func spawnNameGeneratorOrFail() name.Generator {
	namesConfigFile, err := os.ReadFile("assets/names.yml")
	if err != nil {
		printErrorf("error reading names config file: %v\n", err)
		os.Exit(1)
	}

	var config NameConfig
	err = yaml.Unmarshal(namesConfigFile, &config)
	if err != nil {
		printErrorf("error reading names config: %v\n", err)
		os.Exit(1)
	}

	if err := config.Validate(); err != nil {
		printErrorf("error parsing names config: %v\n", err)
		os.Exit(1)
	}

	return name.NewCatalogSourcedGenerator(config.Surnames, config.NonGenderedNames, config.FemaNames, config.MaleNames)
}
