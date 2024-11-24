package generator

import (
	"fmt"
	"github.com/carloscasalar/traveller-npc-generator/internal/npc"
)

type Character struct {
	firstName       string
	surname         string
	role            Role
	category        CitizenCategory
	experience      Experience
	skills          []string
	characteristics map[npc.Characteristic]int
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

func (c Character) Category() CitizenCategory {
	return c.category
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

// String returns a string representation of the character. It is prone to change so you should not rely on it.
func (c Character) String() string {
	return fmt.Sprintf("Character{firstName: %v, surname: %v, role: %v, category: %v, experience: %v, skills: %v, characteristics: %v}",
		c.firstName, c.surname, c.role, c.category, c.experience, c.skills, c.characteristics)
}
