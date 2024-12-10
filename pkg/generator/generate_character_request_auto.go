// Code generated by gonstructor -type=GenerateCharacterRequest -constructorTypes=builder -output=generate_character_request_auto.go; DO NOT EDIT.

package generator

type GenerateCharacterRequestBuilder struct {
	citizenCategory CitizenCategory
	experience      Experience
	role            Role
	gender          Gender
}

func NewGenerateCharacterRequestBuilder() *GenerateCharacterRequestBuilder {
	return &GenerateCharacterRequestBuilder{}
}

func (b *GenerateCharacterRequestBuilder) CitizenCategory(citizenCategory CitizenCategory) *GenerateCharacterRequestBuilder {
	b.citizenCategory = citizenCategory
	return b
}

func (b *GenerateCharacterRequestBuilder) Experience(experience Experience) *GenerateCharacterRequestBuilder {
	b.experience = experience
	return b
}

func (b *GenerateCharacterRequestBuilder) Role(role Role) *GenerateCharacterRequestBuilder {
	b.role = role
	return b
}

func (b *GenerateCharacterRequestBuilder) Gender(gender Gender) *GenerateCharacterRequestBuilder {
	b.gender = gender
	return b
}

func (b *GenerateCharacterRequestBuilder) Build() *GenerateCharacterRequest {
	return &GenerateCharacterRequest{
		citizenCategory: b.citizenCategory,
		experience:      b.experience,
		role:            b.role,
		gender:          b.gender,
	}
}
