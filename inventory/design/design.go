package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("inventory", func() {
	Title("Inventory Service")
	Description("Service for inventories")
	Server("inventory", func() {
		Host("localhost", func() {
			URI("grpc://localhost:8082")
		})
	})
})

var CharacterNotFound = Type("CharacterNotFound", func() {
	Description("character not found")
	Field(1, "message", String, "character not found")
	Field(2, "name", String, "character name")
	Required("message", "name")
})

var ItemNotFound = Type("ItemNotFound", func() {
	Description("item not found")
	Field(1, "message", String, "item not found")
	Field(2, "name", String, "item name")
	Required("message", "name")
})

var InventoryEntry = Type("InventoryEntry", func() {
	Description("A custom type to keep track of item name and its count")
	Field(1, "item", String, "item name")
	Field(2, "count", UInt32, "item count")
	Required("item", "count")
})

var _ = Service("Inventory", func() {
	Description("The inventory service performs CRUD operations on a character's inventory")

	Error("CharacterNotFound", CharacterNotFound)
	Error("ItemNotFound", ItemNotFound)

	Method("addItem", func() {
		Description("Add an item to a character's inventory")
		Payload(func() {
			Field(1, "character", String, "character's name")
			Field(2, "item", String, "item's name")
			Field(3, "count", UInt32, "item's count", func() {
				Minimum(1)
				Default(1)
			})
			Required("character", "item", "count")
		})
		GRPC(func() {
			Response(CodeOK)
			Response("CharacterNotFound", CodeNotFound)
			Response("ItemNotFound", CodeNotFound)
		})
	})

	Method("removeItem", func() {
		Description("Remove an item from a character's inventory")
		Payload(func() {
			Field(1, "character", String, "character's name")
			Field(2, "item", String, "item's name")
			Field(3, "count", UInt32, "item's count", func() {
				Minimum(1)
				Default(1)
			})
			Required("character", "item", "count")
		})
		GRPC(func() {
			Response(CodeOK)
			Response("CharacterNotFound", CodeNotFound)
		})
	})

	Method("getInventory", func() {
		Description("Get a character's inventory")
		Payload(func() {
			Field(1, "character", String, "name of character")
			Required("character")
		})
		Result(ArrayOf(InventoryEntry))
		GRPC(func() {
			Response(CodeOK)
			Response("CharacterNotFound", CodeNotFound)
		})
	})

})
