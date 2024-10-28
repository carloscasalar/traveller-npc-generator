package ui

//go:generate enumer -type=Characteristic -output=characteristic_auto.go
type Characteristic int

const (
	STR Characteristic = iota // Strength
	DEX                       // Dexterity
	END                       // Endurance
	INT                       // Intelligence
	EDU                       // Education
	SOC                       // Social Standing
)
