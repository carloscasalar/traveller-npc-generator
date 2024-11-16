package generator

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"runtime"
)

func NewDefaultGenerateName() (GenerateName, error) {
	const relativeNamesFilePath = "../../assets/names.yml"
	absoluteNamesFilePath := absolutePath(relativeNamesFilePath)
	namesConfigFile, err := os.ReadFile(absoluteNamesFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading names config file at %v: %v\n", absoluteNamesFilePath, err)
	}
	var config nameConfig
	err = yaml.Unmarshal(namesConfigFile, &config)
	if err != nil {
		return nil, fmt.Errorf("error reading names config file at %v: %v\n", absoluteNamesFilePath, err)
	}
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("error parsing names config file at %v: %v\n", absoluteNamesFilePath, err)
	}

	return NewGenerateCatalogSourcedName(config.Surnames, config.NonGenderedNames, config.FemaNames, config.MaleNames), nil
}

func absolutePath(relativeNamesFilePath string) string {
	_, currentFilePath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFilePath)
	absoluteNamesFilePath := filepath.Join(currentDir, relativeNamesFilePath)
	return absoluteNamesFilePath
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
