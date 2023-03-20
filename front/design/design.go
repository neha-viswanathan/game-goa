package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("front", func() {
	Title("Front Service")
	Description("An HTTP service for accessing the backend services of a multiplayer game")
	Server("front", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
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

var Item = Type("Item", func() {
	Description("Attributes of the item")

	Field(1, "name", String, "Name, which is a unique identifier", func() {
		MaxLength(32)
		Example("Potion")
	})
	Field(2, "description", String, "Description", func() {
		MaxLength(128)
		Example("Restores health")
	})
	Field(3, "damage", UInt32, "Amount of damage", func() {
		Minimum(0)
		Maximum(1e6)
		Default(0)
	})
	Field(4, "healing", UInt32, "Amount of healing", func() {
		Minimum(0)
		Maximum(1e6)
		Default(0)
	})
	Field(5, "protection", UInt32, "Amount of protection", func() {
		Minimum(0)
		Maximum(1e4)
		Default(0)
	})
	Required("name", "description")
})

var ItemAlreadyExists = Type("ItemAlreadyExists", func() {
	Description("item already exists")
	Field(1, "message", String, "item already exists")
	Field(2, "name", String, "item name")
	Required("message", "name")
})

var ItemNotFound = Type("ItemNotFound", func() {
	Description("item not found")
	Field(1, "message", String, "item not found")
	Field(2, "name", String, "item name")
	Required("message", "name")
})

var _ = Service("Front", func() {
	Description("The front service is the consumer facing API. It proxies the call to the backend services.")

	Error("CharacterAlreadyExists", CharacterAlreadyExists, "character already exists")
	Error("CharacterNotFound", CharacterNotFound, "character not found")
	Error("ItemAlreadyExists", ItemAlreadyExists, "item already exists")
	Error("ItemNotFound", ItemNotFound, "item not found")

	Method("createCharacter", func() {
		Description("Create a new character")
		Payload(Character)
		Result(Character)
		HTTP(func() {
			POST("/api/character")
			Response(StatusCreated)
			Response("CharacterAlreadyExists", StatusConflict)
		})
	})

	Method("getCharacter", func() {
		Description("Get a character by name")
		Payload(func() {
			Field(1, "name", String, "name of character to be retrieved")
			Required("name")
		})
		Result(Character)
		HTTP(func() {
			GET("/api/character/{name}")
			Response(StatusOK)
			Response("CharacterNotFound", StatusNotFound)
		})
	})

	Method("listCharacters", func() {
		Description("List all characters")
		Result(ArrayOf(Character))
		HTTP(func() {
			GET("/api/character")
			Response(StatusOK)
		})
	})

	Method("updateCharacter", func() {
		Description("Update a given character")
		Payload(Character)
		Result(Character)
		HTTP(func() {
			PUT("/api/character/{name}")
			Response(StatusOK)
			Response("CharacterNotFound", StatusNotFound)
		})
	})

	Method("deleteCharacter", func() {
		Description("Delete a given character")
		Payload(func() {
			Field(1, "name", String, "name of character to be deleted")
			Required("name")
		})
		HTTP(func() {
			DELETE("/api/character/{name}")
			Response(StatusOK)
		})
	})

	Method("createItem", func() {
		Description("Create a new item")
		Payload(Item)
		Result(Item)
		HTTP(func() {
			POST("/api/item")
			Response(StatusCreated)
			Response("ItemAlreadyExists", StatusConflict)
		})
	})

	Method("getItem", func() {
		Description("Get an item by name")
		Payload(func() {
			Field(1, "name", String, "name of item to be retrieved")
			Required("name")
		})
		Result(Item)
		HTTP(func() {
			GET("/api/item/{name}")
			Response(StatusOK)
			Response("ItemNotFound", StatusNotFound)
		})
	})

	Method("listItems", func() {
		Description("List all items")
		Result(ArrayOf(Item))
		HTTP(func() {
			GET("/api/item")
			Response(StatusOK)
		})
	})

	Method("updateItem", func() {
		Description("Update a given item")
		Payload(Item)
		Result(Item)
		Error("ItemNotFound", ItemNotFound, "item not found")
		HTTP(func() {
			PUT("/api/item/{name}")
			Response("ItemNotFound", StatusNotFound)
		})
	})

	Method("deleteItem", func() {
		Description("Delete a given item")
		Payload(func() {
			Field(1, "name", String, "name of item to be deleted")
			Required("name")
		})
		HTTP(func() {
			DELETE("/api/item/{name}")
			Response(StatusOK)
		})
	})

	Method("addInventoryItem", func() {
		Description("Add an item to a character's inventory")
		Payload(func() {
			Field(1, "character", String, "character's name")
			Field(2, "item", String, "item's name")
			Required("character", "item")
		})
		HTTP(func() {
			POST("/api/inventory/{character}/{item}")
			Response(StatusOK)
			Response("CharacterNotFound", StatusNotFound)
			Response("ItemNotFound", StatusNotFound)
		})
	})

	Method("removeInventoryItem", func() {
		Description("Remove an item from a character's inventory")
		Payload(func() {
			Field(1, "character", String, "character's name")
			Field(2, "item", String, "item's name")
			Required("character", "item")
		})
		HTTP(func() {
			DELETE("/api/inventory/{character}/{item}")
			Response(StatusOK)
			Response("CharacterNotFound", StatusNotFound)
		})
	})

	Method("getInventory", func() {
		Description("Get a character's inventory")
		Payload(func() {
			Field(1, "character", String, "name of character")
			Required("character")
		})
		Result(ArrayOf(String))
		HTTP(func() {
			GET("/api/inventory/{character}")
			Response(StatusOK)
			Response("CharacterNotFound", StatusNotFound)
		})
	})
})
