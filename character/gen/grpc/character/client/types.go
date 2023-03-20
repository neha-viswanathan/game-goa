// Code generated by goa v3.11.2, DO NOT EDIT.
//
// Character gRPC client types
//
// Command:
// $ goa gen character/design

package client

import (
	character "character/gen/character"
	characterpb "character/gen/grpc/character/pb"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// NewProtoCreateCharacterRequest builds the gRPC request type from the payload
// of the "createCharacter" endpoint of the "Character" service.
func NewProtoCreateCharacterRequest(payload *character.Character) *characterpb.CreateCharacterRequest {
	message := &characterpb.CreateCharacterRequest{
		Name:        payload.Name,
		Description: &payload.Description,
		Health:      &payload.Health,
		Experience:  &payload.Experience,
	}
	return message
}

// NewCreateCharacterResult builds the result type of the "createCharacter"
// endpoint of the "Character" service from the gRPC response type.
func NewCreateCharacterResult(message *characterpb.CreateCharacterResponse) *character.Character {
	result := &character.Character{
		Name: message.Name,
	}
	if message.Description != nil {
		result.Description = *message.Description
	}
	if message.Health != nil {
		result.Health = *message.Health
	}
	if message.Experience != nil {
		result.Experience = *message.Experience
	}
	if message.Description == nil {
		result.Description = ""
	}
	if message.Health == nil {
		result.Health = 100
	}
	if message.Experience == nil {
		result.Experience = 0
	}
	return result
}

// NewCreateCharacterCharacterAlreadyExistsError builds the error type of the
// "createCharacter" endpoint of the "Character" service from the gRPC error
// response type.
func NewCreateCharacterCharacterAlreadyExistsError(message *characterpb.CreateCharacterCharacterAlreadyExistsError) *character.CharacterAlreadyExists {
	er := &character.CharacterAlreadyExists{
		Message: message.Message_,
		Name:    message.Name,
	}
	return er
}

// NewProtoGetCharacterRequest builds the gRPC request type from the payload of
// the "getCharacter" endpoint of the "Character" service.
func NewProtoGetCharacterRequest(payload *character.GetCharacterPayload) *characterpb.GetCharacterRequest {
	message := &characterpb.GetCharacterRequest{
		Name: payload.Name,
	}
	return message
}

// NewGetCharacterResult builds the result type of the "getCharacter" endpoint
// of the "Character" service from the gRPC response type.
func NewGetCharacterResult(message *characterpb.GetCharacterResponse) *character.Character {
	result := &character.Character{
		Name: message.Name,
	}
	if message.Description != nil {
		result.Description = *message.Description
	}
	if message.Health != nil {
		result.Health = *message.Health
	}
	if message.Experience != nil {
		result.Experience = *message.Experience
	}
	if message.Description == nil {
		result.Description = ""
	}
	if message.Health == nil {
		result.Health = 100
	}
	if message.Experience == nil {
		result.Experience = 0
	}
	return result
}

// NewGetCharacterCharacterNotFoundError builds the error type of the
// "getCharacter" endpoint of the "Character" service from the gRPC error
// response type.
func NewGetCharacterCharacterNotFoundError(message *characterpb.GetCharacterCharacterNotFoundError) *character.CharacterNotFound {
	er := &character.CharacterNotFound{
		Message: message.Message_,
		Name:    message.Name,
	}
	return er
}

// NewProtoListCharactersRequest builds the gRPC request type from the payload
// of the "listCharacters" endpoint of the "Character" service.
func NewProtoListCharactersRequest() *characterpb.ListCharactersRequest {
	message := &characterpb.ListCharactersRequest{}
	return message
}

// NewListCharactersResult builds the result type of the "listCharacters"
// endpoint of the "Character" service from the gRPC response type.
func NewListCharactersResult(message *characterpb.ListCharactersResponse) []*character.Character {
	result := make([]*character.Character, len(message.Field))
	for i, val := range message.Field {
		result[i] = &character.Character{
			Name: val.Name,
		}
		if val.Description != nil {
			result[i].Description = *val.Description
		}
		if val.Health != nil {
			result[i].Health = *val.Health
		}
		if val.Experience != nil {
			result[i].Experience = *val.Experience
		}
		if val.Description == nil {
			result[i].Description = ""
		}
		if val.Health == nil {
			result[i].Health = 100
		}
		if val.Experience == nil {
			result[i].Experience = 0
		}
	}
	return result
}

// NewProtoUpdateCharacterRequest builds the gRPC request type from the payload
// of the "updateCharacter" endpoint of the "Character" service.
func NewProtoUpdateCharacterRequest(payload *character.Character) *characterpb.UpdateCharacterRequest {
	message := &characterpb.UpdateCharacterRequest{
		Name:        payload.Name,
		Description: &payload.Description,
		Health:      &payload.Health,
		Experience:  &payload.Experience,
	}
	return message
}

// NewUpdateCharacterResult builds the result type of the "updateCharacter"
// endpoint of the "Character" service from the gRPC response type.
func NewUpdateCharacterResult(message *characterpb.UpdateCharacterResponse) *character.Character {
	result := &character.Character{
		Name: message.Name,
	}
	if message.Description != nil {
		result.Description = *message.Description
	}
	if message.Health != nil {
		result.Health = *message.Health
	}
	if message.Experience != nil {
		result.Experience = *message.Experience
	}
	if message.Description == nil {
		result.Description = ""
	}
	if message.Health == nil {
		result.Health = 100
	}
	if message.Experience == nil {
		result.Experience = 0
	}
	return result
}

// NewUpdateCharacterCharacterNotFoundError builds the error type of the
// "updateCharacter" endpoint of the "Character" service from the gRPC error
// response type.
func NewUpdateCharacterCharacterNotFoundError(message *characterpb.UpdateCharacterCharacterNotFoundError) *character.CharacterNotFound {
	er := &character.CharacterNotFound{
		Message: message.Message_,
		Name:    message.Name,
	}
	return er
}

// NewProtoDeleteCharacterRequest builds the gRPC request type from the payload
// of the "deleteCharacter" endpoint of the "Character" service.
func NewProtoDeleteCharacterRequest(payload *character.DeleteCharacterPayload) *characterpb.DeleteCharacterRequest {
	message := &characterpb.DeleteCharacterRequest{
		Name: payload.Name,
	}
	return message
}

// ValidateCreateCharacterResponse runs the validations defined on
// CreateCharacterResponse.
func ValidateCreateCharacterResponse(message *characterpb.CreateCharacterResponse) (err error) {
	if utf8.RuneCountInString(message.Name) > 32 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 32, false))
	}
	if message.Description != nil {
		if utf8.RuneCountInString(*message.Description) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("message.description", *message.Description, utf8.RuneCountInString(*message.Description), 128, false))
		}
	}
	if message.Health != nil {
		if *message.Health < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.health", *message.Health, 0, true))
		}
	}
	if message.Health != nil {
		if *message.Health > 1e+07 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.health", *message.Health, 1e+07, false))
		}
	}
	if message.Experience != nil {
		if *message.Experience < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.experience", *message.Experience, 0, true))
		}
	}
	if message.Experience != nil {
		if *message.Experience > 1e+07 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.experience", *message.Experience, 1e+07, false))
		}
	}
	return
}

// ValidateGetCharacterResponse runs the validations defined on
// GetCharacterResponse.
func ValidateGetCharacterResponse(message *characterpb.GetCharacterResponse) (err error) {
	if utf8.RuneCountInString(message.Name) > 32 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 32, false))
	}
	if message.Description != nil {
		if utf8.RuneCountInString(*message.Description) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("message.description", *message.Description, utf8.RuneCountInString(*message.Description), 128, false))
		}
	}
	if message.Health != nil {
		if *message.Health < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.health", *message.Health, 0, true))
		}
	}
	if message.Health != nil {
		if *message.Health > 1e+07 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.health", *message.Health, 1e+07, false))
		}
	}
	if message.Experience != nil {
		if *message.Experience < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.experience", *message.Experience, 0, true))
		}
	}
	if message.Experience != nil {
		if *message.Experience > 1e+07 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.experience", *message.Experience, 1e+07, false))
		}
	}
	return
}

// ValidateListCharactersResponse runs the validations defined on
// ListCharactersResponse.
func ValidateListCharactersResponse(message *characterpb.ListCharactersResponse) (err error) {
	for _, e := range message.Field {
		if e != nil {
			if err2 := ValidateCharacter2(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateCharacter2 runs the validations defined on Character2.
func ValidateCharacter2(elem *characterpb.Character2) (err error) {
	if utf8.RuneCountInString(elem.Name) > 32 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("elem.name", elem.Name, utf8.RuneCountInString(elem.Name), 32, false))
	}
	if elem.Description != nil {
		if utf8.RuneCountInString(*elem.Description) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("elem.description", *elem.Description, utf8.RuneCountInString(*elem.Description), 128, false))
		}
	}
	if elem.Health != nil {
		if *elem.Health < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("elem.health", *elem.Health, 0, true))
		}
	}
	if elem.Health != nil {
		if *elem.Health > 1e+07 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("elem.health", *elem.Health, 1e+07, false))
		}
	}
	if elem.Experience != nil {
		if *elem.Experience < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("elem.experience", *elem.Experience, 0, true))
		}
	}
	if elem.Experience != nil {
		if *elem.Experience > 1e+07 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("elem.experience", *elem.Experience, 1e+07, false))
		}
	}
	return
}

// ValidateUpdateCharacterResponse runs the validations defined on
// UpdateCharacterResponse.
func ValidateUpdateCharacterResponse(message *characterpb.UpdateCharacterResponse) (err error) {
	if utf8.RuneCountInString(message.Name) > 32 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 32, false))
	}
	if message.Description != nil {
		if utf8.RuneCountInString(*message.Description) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("message.description", *message.Description, utf8.RuneCountInString(*message.Description), 128, false))
		}
	}
	if message.Health != nil {
		if *message.Health < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.health", *message.Health, 0, true))
		}
	}
	if message.Health != nil {
		if *message.Health > 1e+07 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.health", *message.Health, 1e+07, false))
		}
	}
	if message.Experience != nil {
		if *message.Experience < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.experience", *message.Experience, 0, true))
		}
	}
	if message.Experience != nil {
		if *message.Experience > 1e+07 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.experience", *message.Experience, 1e+07, false))
		}
	}
	return
}