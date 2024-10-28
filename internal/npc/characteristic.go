package npc

//go:generate enumer -type=Characteristic -output=characteristics_auto.go
type Characteristic int

const (
	STR Characteristic = iota // Strength
	DEX                       // Dexterity
	END                       // Endurance
	INT                       // Intelligence
	EDU                       // Education
	SOC                       // Social Standing
)

type Characteristics map[Characteristic]int

func (c Characteristics) Stringify() map[string]int {
	result := make(map[string]int)
	for k, v := range c {
		result[k.String()] = v
	}
	return result
}

type RoleCharacteristics struct {
	High   [2]Characteristic
	Medium [2]Characteristic
	Low    [2]Characteristic
}

var characteristicPreferenceByRole = map[Role]RoleCharacteristics{
	RolePilot:       {High: [2]Characteristic{DEX, INT}, Medium: [2]Characteristic{EDU, STR}, Low: [2]Characteristic{END, SOC}},
	RoleNavigator:   {High: [2]Characteristic{INT, EDU}, Medium: [2]Characteristic{DEX, SOC}, Low: [2]Characteristic{STR, END}},
	RoleEngineer:    {High: [2]Characteristic{INT, EDU}, Medium: [2]Characteristic{DEX, END}, Low: [2]Characteristic{STR, SOC}},
	RoleSteward:     {High: [2]Characteristic{INT, SOC}, Medium: [2]Characteristic{DEX, EDU}, Low: [2]Characteristic{STR, END}},
	RoleMedic:       {High: [2]Characteristic{INT, EDU}, Medium: [2]Characteristic{DEX, SOC}, Low: [2]Characteristic{STR, END}},
	RoleMarine:      {High: [2]Characteristic{STR, END}, Medium: [2]Characteristic{DEX, INT}, Low: [2]Characteristic{EDU, SOC}},
	RoleGunner:      {High: [2]Characteristic{DEX, INT}, Medium: [2]Characteristic{END, EDU}, Low: [2]Characteristic{STR, SOC}},
	RoleScout:       {High: [2]Characteristic{DEX, INT}, Medium: [2]Characteristic{END, EDU}, Low: [2]Characteristic{STR, SOC}},
	RoleTechnician:  {High: [2]Characteristic{INT, EDU}, Medium: [2]Characteristic{DEX, END}, Low: [2]Characteristic{STR, SOC}},
	RoleLeader:      {High: [2]Characteristic{INT, SOC}, Medium: [2]Characteristic{EDU, END}, Low: [2]Characteristic{DEX, STR}},
	RoleDiplomat:    {High: [2]Characteristic{INT, SOC}, Medium: [2]Characteristic{EDU, DEX}, Low: [2]Characteristic{STR, END}},
	RoleEntertainer: {High: [2]Characteristic{DEX, SOC}, Medium: [2]Characteristic{INT, EDU}, Low: [2]Characteristic{STR, END}},
	RoleTrader:      {High: [2]Characteristic{INT, SOC}, Medium: [2]Characteristic{EDU, DEX}, Low: [2]Characteristic{STR, END}},
	RoleThug:        {High: [2]Characteristic{STR, END}, Medium: [2]Characteristic{DEX, INT}, Low: [2]Characteristic{EDU, SOC}},
}
