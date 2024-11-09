package main

import (
	"fmt"
	"github.com/carloscasalar/traveller-npc-generator/internal/npc"
	"github.com/carloscasalar/traveller-npc-generator/internal/ui"
	"os"
)

func main() {
	opts := readOptionsOrFail()
	nameGenerator := spawnNameGeneratorOrFail()
	debugEnabled := opts.EnableDebug

	if debugEnabled {
		printOptionsRead(opts)
	}

	category := readCitizenCategory(opts)
	experience := readExperience(opts)
	role := readRole(opts)
	gender := readGender(opts)
	firstName, surname := nameGenerator.Generate(gender)
	fullName := fmt.Sprintf("%v %v", firstName, surname)
	characteristic := role.RandomCharacteristic(category)

	if debugEnabled {
		printGeneratedValues(fullName, role, experience, characteristic)
	}

	sheet := ui.NewCharacterSheetBuilder().
		FullName(fullName).
		Role(role.String()).
		CitizenCategory(category.String()).
		Experience(experience.String()).
		Skills(role.Skills(experience)).
		Characteristics(toUICharacteristics(characteristic)).
		Build()
	printSheet(sheet)
}

func printOptionsRead(opts CommandOptions) {
	prompt("Inputs--------------")
	titleValue("CitizenCategory: ", opts.CitizenCategory)
	titleValue("Experience: ", opts.Experience)
	titleValue("CrewRole: ", opts.CrewRole)
}

func printGeneratedValues(fullName string, role npc.Role, experience npc.Experience, characteristic map[npc.Characteristic]int) {
	prompt("Generated ----------")
	titleValue("Name: ", fullName)
	titleValue("Role: ", role.String())
	titleValue("Experience: ", experience.String())
	titleValue("Skills: ", role.Skills(experience))
	titleValue("Characteristics: ", characteristic)
	prompt("--------------------")
}

func toUICharacteristics(characteristic map[npc.Characteristic]int) map[ui.Characteristic]int {
	return map[ui.Characteristic]int{
		ui.STR: characteristic[npc.STR],
		ui.DEX: characteristic[npc.DEX],
		ui.END: characteristic[npc.END],
		ui.INT: characteristic[npc.INT],
		ui.EDU: characteristic[npc.EDU],
		ui.SOC: characteristic[npc.SOC],
	}
}

func prompt(value string) {
	fmt.Println(ui.NewPromptRenderer(value).Render())
}

func printErrorf(template string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, template, a...)
}

func titleValue[T any](title string, value T) {
	fmt.Println(ui.NewTitleValueRenderer(title, fmt.Sprintf("%v", value)).Render())
}

func printSheet(sheet *ui.CharacterSheet) {
	fmt.Println(sheet.Render())
}
