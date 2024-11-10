package generator

//go:generate gonstructor -type=GenerateCharacterRequest -constructorTypes=builder -output=generate_character_request_auto.go

type GenerateCharacterRequest struct {
	category   CitizenCategory
	experience Experience
	role       Role
	gender     Gender
}
