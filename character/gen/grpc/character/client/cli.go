// Code generated by goa v3.11.2, DO NOT EDIT.
//
// Character gRPC client CLI support package
//
// Command:
// $ goa gen character/design

package client

import (
	character "character/gen/character"
	characterpb "character/gen/grpc/character/pb"
	"encoding/json"
	"fmt"
)

// BuildCreateCharacterPayload builds the payload for the Character
// createCharacter endpoint from CLI flags.
func BuildCreateCharacterPayload(characterCreateCharacterMessage string) (*character.Character, error) {
	var err error
	var message characterpb.CreateCharacterRequest
	{
		if characterCreateCharacterMessage != "" {
			err = json.Unmarshal([]byte(characterCreateCharacterMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"description\": \"Savior of Princess Zelda\",\n      \"experience\": 518303,\n      \"health\": 8319921,\n      \"name\": \"Link\"\n   }'")
			}
		}
	}
	v := &character.Character{
		Name: message.Name,
	}
	if message.Description != nil {
		v.Description = *message.Description
	}
	if message.Health != nil {
		v.Health = *message.Health
	}
	if message.Experience != nil {
		v.Experience = *message.Experience
	}
	if message.Description == nil {
		v.Description = ""
	}
	if message.Health == nil {
		v.Health = 100
	}
	if message.Experience == nil {
		v.Experience = 0
	}

	return v, nil
}

// BuildGetCharacterPayload builds the payload for the Character getCharacter
// endpoint from CLI flags.
func BuildGetCharacterPayload(characterGetCharacterMessage string) (*character.GetCharacterPayload, error) {
	var err error
	var message characterpb.GetCharacterRequest
	{
		if characterGetCharacterMessage != "" {
			err = json.Unmarshal([]byte(characterGetCharacterMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Eius earum quis et.\"\n   }'")
			}
		}
	}
	v := &character.GetCharacterPayload{
		Name: message.Name,
	}

	return v, nil
}

// BuildUpdateCharacterPayload builds the payload for the Character
// updateCharacter endpoint from CLI flags.
func BuildUpdateCharacterPayload(characterUpdateCharacterMessage string) (*character.Character, error) {
	var err error
	var message characterpb.UpdateCharacterRequest
	{
		if characterUpdateCharacterMessage != "" {
			err = json.Unmarshal([]byte(characterUpdateCharacterMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"description\": \"Savior of Princess Zelda\",\n      \"experience\": 4698556,\n      \"health\": 6773187,\n      \"name\": \"Link\"\n   }'")
			}
		}
	}
	v := &character.Character{
		Name: message.Name,
	}
	if message.Description != nil {
		v.Description = *message.Description
	}
	if message.Health != nil {
		v.Health = *message.Health
	}
	if message.Experience != nil {
		v.Experience = *message.Experience
	}
	if message.Description == nil {
		v.Description = ""
	}
	if message.Health == nil {
		v.Health = 100
	}
	if message.Experience == nil {
		v.Experience = 0
	}

	return v, nil
}

// BuildDeleteCharacterPayload builds the payload for the Character
// deleteCharacter endpoint from CLI flags.
func BuildDeleteCharacterPayload(characterDeleteCharacterMessage string) (*character.DeleteCharacterPayload, error) {
	var err error
	var message characterpb.DeleteCharacterRequest
	{
		if characterDeleteCharacterMessage != "" {
			err = json.Unmarshal([]byte(characterDeleteCharacterMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Ut quaerat velit fuga voluptas fuga ea.\"\n   }'")
			}
		}
	}
	v := &character.DeleteCharacterPayload{
		Name: message.Name,
	}

	return v, nil
}
