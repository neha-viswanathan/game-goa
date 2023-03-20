package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("character", func() {
	Title("Character Service")
	Description("Service for characters")
	Server("character", func() {
		Host("localhost", func() {
			URI("grpc://localhost:8080")
		})
	})
})

var Character = Type("Character", func() {
	Description("Attributes of the character")

	Field(1, "name", String, "Name, which is a unique identifier", func() {
		MaxLength(32)
		Example("Link")
	})
	Field(2, "description", String, "Description", func() {
		MaxLength(128)
		Default("")
		Example("Savior of Princess Zelda")
	})
	Field(3, "health", UInt32, "Current health", func() {
		Minimum(0)
		Maximum(1e7)
		Default(100)
	})
	Field(4, "experience", UInt32, "Amount of experience", func() {
		Minimum(0)
		Maximum(1e7)
		Default(0)
	})

	Required("name")
})

var CharacterAlreadyExists = Type("CharacterAlreadyExists", func() {
	Description("character already exists")
	Field(1, "message", String, "character already exists")
	Field(2, "name", String, "character name")
	Required("message", "name")
})

var CharacterNotFound = Type("CharacterNotFound", func() {
	Description("character not found")
	Field(1, "message", String, "character not found")
	Field(2, "name", String, "character name")
	Required("message", "name")
})

var _ = Service("Character", func() {
	Description("The character service performs CRUD operations on a character")

	Error("CharacterAlreadyExists", CharacterAlreadyExists)
	Error("CharacterNotFound", CharacterNotFound)

	Method("createCharacter", func() {
		Description("Create a new character")
		Payload(Character)
		Result(Character)
		Error("CharacterAlreadyExists", CharacterAlreadyExists, "character already exists")
		GRPC(func() {
			Response(CodeOK)
			Response("CharacterAlreadyExists", CodeAlreadyExists)
		})
	})

	Method("getCharacter", func() {
		Description("Get a character by name")
		Payload(func() {
			Field(1, "name", String, "name of character to be retrieved")
			Required("name")
		})
		Result(Character)
		Error("CharacterNotFound", CharacterNotFound, "character not found")
		GRPC(func() {
			Response(CodeOK)
			Response("CharacterNotFound", CodeNotFound)
		})
	})

	Method("listCharacters", func() {
		Description("List all characters")
		Result(ArrayOf(Character))
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("updateCharacter", func() {
		Description("Update a given character")
		Payload(Character)
		Result(Character)
		Error("CharacterNotFound", CharacterNotFound, "character not found")
		GRPC(func() {
			Response(CodeOK)
			Response("CharacterNotFound", CodeNotFound)
		})
	})

	Method("deleteCharacter", func() {
		Description("Delete a given character")
		Payload(func() {
			Field(1, "name", String, "name of character to be deleted")
			Required("name")
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})
})
