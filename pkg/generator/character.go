package generator

import (
	"fmt"
	"github.com/carloscasalar/traveller-npc-generator/internal/npc"
)

type Character struct {
	firstName       string
	surname         string
	role            Role
	citizenCategory CitizenCategory
	experience      Experience
	skills          []string
	characteristics map[npc.Characteristic]int
	equipment       EquipmentSet
}

func (c Character) FirstName() string {
	return c.firstName
}

func (c Character) Surname() string {
	return c.surname
}

func (c Character) FullName() string {
	return fmt.Sprintf("%v %v", c.firstName, c.surname)
}

func (c Character) Role() Role {
	return c.role
}

func (c Character) CitizenCategory() CitizenCategory {
	return c.citizenCategory
}

func (c Character) Experience() Experience {
	return c.experience
}

func (c Character) Skills() []string {
	return c.skills
}

func (c Character) Characteristics() map[Characteristic]int {
	return map[Characteristic]int{
		STR: c.characteristics[npc.STR],
		DEX: c.characteristics[npc.DEX],
		END: c.characteristics[npc.END],
		INT: c.characteristics[npc.INT],
		EDU: c.characteristics[npc.EDU],
		SOC: c.characteristics[npc.SOC],
	}
}

func (c Character) Equipment() EquipmentSet {
	return c.equipment
}

// String returns a string representation of the character. It is prone to change so you should not rely on it.
func (c Character) String() string {
	return fmt.Sprintf("Character{firstName: %v, surname: %v, role: %v, citizenCategory: %v, experience: %v, skills: %v, characteristics: %v, equipment: %+v}",
		c.firstName, c.surname, c.role, c.citizenCategory, c.experience, c.skills, c.characteristics, c.equipment)
}
