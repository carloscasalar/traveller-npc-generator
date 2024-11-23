package main

import (
	"fmt"
	"github.com/carloscasalar/traveller-npc-generator/internal/ui"
	"github.com/carloscasalar/traveller-npc-generator/pkg/generator"
	"os"
)

func main() {
	opts := readOptionsOrFail()
	npcGenerator := spawnNpcGeneratorOrFail()
	debugEnabled := opts.EnableDebug

	if debugEnabled {
		printOptionsRead(opts)
	}

	category := readCitizenCategory(opts)
	experience := readExperience(opts)
	role := readRole(opts)
	gender := readGender(opts)
	newCharacterRequest := generator.NewGenerateCharacterRequestBuilder().
		Category(category).
		Experience(experience).
		Role(role).
		Gender(gender).
		Build()

	character, err := npcGenerator.Generate(*newCharacterRequest)
	if err != nil {
		printError(err)
		os.Exit(1)
	}

	if debugEnabled {
		printGeneratedValues(*character)
	}

	sheet := ui.NewCharacterSheetBuilder().
		FullName(character.FullName()).
		Role(role.String()).
		CitizenCategory(category.String()).
		Experience(experience.String()).
		Skills(character.Skills()).
		Characteristics(toUICharacteristics(character.Characteristics())).
		Build()
	printSheet(sheet)
}

func printOptionsRead(opts CommandOptions) {
	prompt("Inputs--------------")
	titleValue("CitizenCategory: ", opts.CitizenCategory)
	titleValue("Experience: ", opts.Experience)
	titleValue("CrewRole: ", opts.CrewRole)
}

func printGeneratedValues(character generator.Character) {
	prompt("Generated ----------")
	titleValue("Name: ", character.FullName())
	titleValue("Role: ", character.Role().String())
	titleValue("Experience: ", character.Experience().String())
	titleValue("Skills: ", character.Skills())
	titleValue("Characteristics: ", character.Characteristics())
	prompt("--------------------")
}

func toUICharacteristics(characteristic map[generator.Characteristic]int) map[ui.Characteristic]int {
	return map[ui.Characteristic]int{
		ui.STR: characteristic[generator.STR],
		ui.DEX: characteristic[generator.DEX],
		ui.END: characteristic[generator.END],
		ui.INT: characteristic[generator.INT],
		ui.EDU: characteristic[generator.EDU],
		ui.SOC: characteristic[generator.SOC],
	}
}

func spawnGenerateNameOrFail() generator.NameGenerator {
	nameGenerator, err := generator.NewDefaultNameGenerator()
	if err != nil {
		printError(err)
		os.Exit(1)
	}

	return nameGenerator
}

func spawnNpcGeneratorOrFail() *generator.NpcGenerator {
	nameGenerator := spawnGenerateNameOrFail()
	npcGenerator, err := generator.NewNpcGeneratorBuilder().NameGenerator(nameGenerator).Build()
	if err != nil {
		printError(err)
		os.Exit(1)
	}
	return npcGenerator
}

func prompt(value string) {
	fmt.Println(ui.NewPromptRenderer(value).Render())
}

func printErrorf(template string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, template, a...)
}

func printError(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err.Error())
}

func titleValue[T any](title string, value T) {
	fmt.Println(ui.NewTitleValueRenderer(title, fmt.Sprintf("%v", value)).Render())
}

func printSheet(sheet *ui.CharacterSheet) {
	fmt.Println(sheet.Render())
}
