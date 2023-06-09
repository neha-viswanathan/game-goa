// Code generated by goa v3.11.2, DO NOT EDIT.
//
// Inventory gRPC client CLI support package
//
// Command:
// $ goa gen inventory/design

package client

import (
	"encoding/json"
	"fmt"
	inventorypb "inventory/gen/grpc/inventory/pb"
	inventory "inventory/gen/inventory"
)

// BuildAddItemPayload builds the payload for the Inventory addItem endpoint
// from CLI flags.
func BuildAddItemPayload(inventoryAddItemMessage string) (*inventory.AddItemPayload, error) {
	var err error
	var message inventorypb.AddItemRequest
	{
		if inventoryAddItemMessage != "" {
			err = json.Unmarshal([]byte(inventoryAddItemMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"character\": \"Laborum sed dolores corporis.\",\n      \"count\": 3513363385,\n      \"item\": \"Eum cum officia aut velit dolorum explicabo.\"\n   }'")
			}
		}
	}
	v := &inventory.AddItemPayload{
		Character: message.Character,
		Item:      message.Item,
		Count:     message.Count,
	}

	return v, nil
}

// BuildRemoveItemPayload builds the payload for the Inventory removeItem
// endpoint from CLI flags.
func BuildRemoveItemPayload(inventoryRemoveItemMessage string) (*inventory.RemoveItemPayload, error) {
	var err error
	var message inventorypb.RemoveItemRequest
	{
		if inventoryRemoveItemMessage != "" {
			err = json.Unmarshal([]byte(inventoryRemoveItemMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"character\": \"Recusandae quis.\",\n      \"count\": 155887092,\n      \"item\": \"Alias omnis esse exercitationem molestiae voluptatem rerum.\"\n   }'")
			}
		}
	}
	v := &inventory.RemoveItemPayload{
		Character: message.Character,
		Item:      message.Item,
		Count:     message.Count,
	}

	return v, nil
}

// BuildGetInventoryPayload builds the payload for the Inventory getInventory
// endpoint from CLI flags.
func BuildGetInventoryPayload(inventoryGetInventoryMessage string) (*inventory.GetInventoryPayload, error) {
	var err error
	var message inventorypb.GetInventoryRequest
	{
		if inventoryGetInventoryMessage != "" {
			err = json.Unmarshal([]byte(inventoryGetInventoryMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"character\": \"Consequuntur et porro sit.\"\n   }'")
			}
		}
	}
	v := &inventory.GetInventoryPayload{
		Character: message.Character,
	}

	return v, nil
}
