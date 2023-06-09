// Code generated by goa v3.11.2, DO NOT EDIT.
//
// Item gRPC client CLI support package
//
// Command:
// $ goa gen item/design

package client

import (
	"encoding/json"
	"fmt"
	itempb "item/gen/grpc/item/pb"
	item "item/gen/item"
)

// BuildCreateItemPayload builds the payload for the Item createItem endpoint
// from CLI flags.
func BuildCreateItemPayload(itemCreateItemMessage string) (*item.Item, error) {
	var err error
	var message itempb.CreateItemRequest
	{
		if itemCreateItemMessage != "" {
			err = json.Unmarshal([]byte(itemCreateItemMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"damage\": 298971,\n      \"description\": \"Restores health\",\n      \"healing\": 501251,\n      \"name\": \"Potion\",\n      \"protection\": 445\n   }'")
			}
		}
	}
	v := &item.Item{
		Name:        message.Name,
		Description: message.Description,
	}
	if message.Damage != nil {
		v.Damage = *message.Damage
	}
	if message.Healing != nil {
		v.Healing = *message.Healing
	}
	if message.Protection != nil {
		v.Protection = *message.Protection
	}
	if message.Damage == nil {
		v.Damage = 0
	}
	if message.Healing == nil {
		v.Healing = 0
	}
	if message.Protection == nil {
		v.Protection = 0
	}

	return v, nil
}

// BuildGetItemPayload builds the payload for the Item getItem endpoint from
// CLI flags.
func BuildGetItemPayload(itemGetItemMessage string) (*item.GetItemPayload, error) {
	var err error
	var message itempb.GetItemRequest
	{
		if itemGetItemMessage != "" {
			err = json.Unmarshal([]byte(itemGetItemMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Quo qui ut nemo omnis porro eum.\"\n   }'")
			}
		}
	}
	v := &item.GetItemPayload{
		Name: message.Name,
	}

	return v, nil
}

// BuildUpdateItemPayload builds the payload for the Item updateItem endpoint
// from CLI flags.
func BuildUpdateItemPayload(itemUpdateItemMessage string) (*item.Item, error) {
	var err error
	var message itempb.UpdateItemRequest
	{
		if itemUpdateItemMessage != "" {
			err = json.Unmarshal([]byte(itemUpdateItemMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"damage\": 893006,\n      \"description\": \"Restores health\",\n      \"healing\": 421259,\n      \"name\": \"Potion\",\n      \"protection\": 5974\n   }'")
			}
		}
	}
	v := &item.Item{
		Name:        message.Name,
		Description: message.Description,
	}
	if message.Damage != nil {
		v.Damage = *message.Damage
	}
	if message.Healing != nil {
		v.Healing = *message.Healing
	}
	if message.Protection != nil {
		v.Protection = *message.Protection
	}
	if message.Damage == nil {
		v.Damage = 0
	}
	if message.Healing == nil {
		v.Healing = 0
	}
	if message.Protection == nil {
		v.Protection = 0
	}

	return v, nil
}

// BuildDeleteItemPayload builds the payload for the Item deleteItem endpoint
// from CLI flags.
func BuildDeleteItemPayload(itemDeleteItemMessage string) (*item.DeleteItemPayload, error) {
	var err error
	var message itempb.DeleteItemRequest
	{
		if itemDeleteItemMessage != "" {
			err = json.Unmarshal([]byte(itemDeleteItemMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Aliquid consequatur libero.\"\n   }'")
			}
		}
	}
	v := &item.DeleteItemPayload{
		Name: message.Name,
	}

	return v, nil
}
