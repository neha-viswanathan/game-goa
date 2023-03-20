package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("item", func() {
	Title("Item Service")
	Description("Service for items")
	Server("item", func() {
		Host("localhost", func() {
			URI("grpc://localhost:8081")
		})
	})
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

var _ = Service("Item", func() {
	Description("The item service performs CRUD operations on a item")

	Error("ItemAlreadyExists", ItemAlreadyExists)
	Error("ItemNotFound", ItemNotFound)

	Method("createItem", func() {
		Description("Create a new item")
		Payload(Item)
		Result(Item)
		Error("ItemAlreadyExists", ItemAlreadyExists, "item already exists")
		GRPC(func() {
			Response(CodeOK)
			Response("ItemAlreadyExists", CodeAlreadyExists)
		})
	})

	Method("getItem", func() {
		Description("Get an item by name")
		Payload(func() {
			Field(1, "name", String, "name of item to be retrieved")
			Required("name")
		})
		Result(Item)
		Error("ItemNotFound", ItemNotFound, "item not found")
		GRPC(func() {
			Response(CodeOK)
			Response("ItemNotFound", CodeNotFound)
		})
	})

	Method("listItems", func() {
		Description("List all items")
		Result(ArrayOf(Item))
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("updateItem", func() {
		Description("Update a given item")
		Payload(Item)
		Result(Item)
		Error("ItemNotFound", ItemNotFound, "item not found")
		GRPC(func() {
			Response(CodeOK)
			Response("ItemNotFound", CodeNotFound)
		})
	})

	Method("deleteItem", func() {
		Description("Delete a given item")
		Payload(func() {
			Field(1, "name", String, "name of item to be deleted")
			Required("name")
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})
})
