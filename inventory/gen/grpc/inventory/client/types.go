// Code generated by goa v3.11.2, DO NOT EDIT.
//
// Inventory gRPC client types
//
// Command:
// $ goa gen inventory/design

package client

import (
	inventorypb "inventory/gen/grpc/inventory/pb"
	inventory "inventory/gen/inventory"
)

// NewProtoAddItemRequest builds the gRPC request type from the payload of the
// "addItem" endpoint of the "Inventory" service.
func NewProtoAddItemRequest(payload *inventory.AddItemPayload) *inventorypb.AddItemRequest {
	message := &inventorypb.AddItemRequest{
		Character: payload.Character,
		Item:      payload.Item,
		Count:     payload.Count,
	}
	return message
}

// NewAddItemCharacterNotFoundError builds the error type of the "addItem"
// endpoint of the "Inventory" service from the gRPC error response type.
func NewAddItemCharacterNotFoundError(message *inventorypb.AddItemCharacterNotFoundError) *inventory.CharacterNotFound {
	er := &inventory.CharacterNotFound{
		Message: message.Message_,
		Name:    message.Name,
	}
	return er
}

// NewAddItemItemNotFoundError builds the error type of the "addItem" endpoint
// of the "Inventory" service from the gRPC error response type.
func NewAddItemItemNotFoundError(message *inventorypb.AddItemItemNotFoundError) *inventory.ItemNotFound {
	er := &inventory.ItemNotFound{
		Message: message.Message_,
		Name:    message.Name,
	}
	return er
}

// NewProtoRemoveItemRequest builds the gRPC request type from the payload of
// the "removeItem" endpoint of the "Inventory" service.
func NewProtoRemoveItemRequest(payload *inventory.RemoveItemPayload) *inventorypb.RemoveItemRequest {
	message := &inventorypb.RemoveItemRequest{
		Character: payload.Character,
		Item:      payload.Item,
		Count:     payload.Count,
	}
	return message
}

// NewRemoveItemCharacterNotFoundError builds the error type of the
// "removeItem" endpoint of the "Inventory" service from the gRPC error
// response type.
func NewRemoveItemCharacterNotFoundError(message *inventorypb.RemoveItemCharacterNotFoundError) *inventory.CharacterNotFound {
	er := &inventory.CharacterNotFound{
		Message: message.Message_,
		Name:    message.Name,
	}
	return er
}

// NewProtoGetInventoryRequest builds the gRPC request type from the payload of
// the "getInventory" endpoint of the "Inventory" service.
func NewProtoGetInventoryRequest(payload *inventory.GetInventoryPayload) *inventorypb.GetInventoryRequest {
	message := &inventorypb.GetInventoryRequest{
		Character: payload.Character,
	}
	return message
}

// NewGetInventoryResult builds the result type of the "getInventory" endpoint
// of the "Inventory" service from the gRPC response type.
func NewGetInventoryResult(message *inventorypb.GetInventoryResponse) []*inventory.InventoryEntry {
	result := make([]*inventory.InventoryEntry, len(message.Field))
	for i, val := range message.Field {
		result[i] = &inventory.InventoryEntry{
			Item:  val.Item,
			Count: val.Count,
		}
	}
	return result
}

// NewGetInventoryCharacterNotFoundError builds the error type of the
// "getInventory" endpoint of the "Inventory" service from the gRPC error
// response type.
func NewGetInventoryCharacterNotFoundError(message *inventorypb.GetInventoryCharacterNotFoundError) *inventory.CharacterNotFound {
	er := &inventory.CharacterNotFound{
		Message: message.Message_,
		Name:    message.Name,
	}
	return er
}
