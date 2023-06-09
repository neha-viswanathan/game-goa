// Code generated by goa v3.11.2, DO NOT EDIT.
//
// Front HTTP client CLI support package
//
// Command:
// $ goa gen front/design

package client

import (
	"encoding/json"
	"fmt"
	front "front/gen/front"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// BuildCreateCharacterPayload builds the payload for the Front createCharacter
// endpoint from CLI flags.
func BuildCreateCharacterPayload(frontCreateCharacterBody string) (*front.Character, error) {
	var err error
	var body CreateCharacterRequestBody
	{
		err = json.Unmarshal([]byte(frontCreateCharacterBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"description\": \"Savior of Princess Zelda\",\n      \"experience\": 7974469,\n      \"health\": 5252954,\n      \"name\": \"Link\"\n   }'")
		}
		if utf8.RuneCountInString(body.Name) > 32 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 32, false))
		}
		if utf8.RuneCountInString(body.Description) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", body.Description, utf8.RuneCountInString(body.Description), 128, false))
		}
		if body.Health < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.health", body.Health, 0, true))
		}
		if body.Health > 1e+07 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.health", body.Health, 1e+07, false))
		}
		if body.Experience < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.experience", body.Experience, 0, true))
		}
		if body.Experience > 1e+07 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.experience", body.Experience, 1e+07, false))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &front.Character{
		Name:        body.Name,
		Description: body.Description,
		Health:      body.Health,
		Experience:  body.Experience,
	}
	{
		var zero string
		if v.Description == zero {
			v.Description = ""
		}
	}
	{
		var zero uint32
		if v.Health == zero {
			v.Health = 100
		}
	}
	{
		var zero uint32
		if v.Experience == zero {
			v.Experience = 0
		}
	}

	return v, nil
}

// BuildGetCharacterPayload builds the payload for the Front getCharacter
// endpoint from CLI flags.
func BuildGetCharacterPayload(frontGetCharacterName string) (*front.GetCharacterPayload, error) {
	var name string
	{
		name = frontGetCharacterName
	}
	v := &front.GetCharacterPayload{}
	v.Name = name

	return v, nil
}

// BuildUpdateCharacterPayload builds the payload for the Front updateCharacter
// endpoint from CLI flags.
func BuildUpdateCharacterPayload(frontUpdateCharacterBody string, frontUpdateCharacterName string) (*front.Character, error) {
	var err error
	var body UpdateCharacterRequestBody
	{
		err = json.Unmarshal([]byte(frontUpdateCharacterBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"description\": \"Savior of Princess Zelda\",\n      \"experience\": 2029508,\n      \"health\": 1800536\n   }'")
		}
		if utf8.RuneCountInString(body.Description) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", body.Description, utf8.RuneCountInString(body.Description), 128, false))
		}
		if body.Health < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.health", body.Health, 0, true))
		}
		if body.Health > 1e+07 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.health", body.Health, 1e+07, false))
		}
		if body.Experience < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.experience", body.Experience, 0, true))
		}
		if body.Experience > 1e+07 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.experience", body.Experience, 1e+07, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var name string
	{
		name = frontUpdateCharacterName
		if utf8.RuneCountInString(name) > 32 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("name", name, utf8.RuneCountInString(name), 32, false))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &front.Character{
		Description: body.Description,
		Health:      body.Health,
		Experience:  body.Experience,
	}
	{
		var zero string
		if v.Description == zero {
			v.Description = ""
		}
	}
	{
		var zero uint32
		if v.Health == zero {
			v.Health = 100
		}
	}
	{
		var zero uint32
		if v.Experience == zero {
			v.Experience = 0
		}
	}
	v.Name = name

	return v, nil
}

// BuildDeleteCharacterPayload builds the payload for the Front deleteCharacter
// endpoint from CLI flags.
func BuildDeleteCharacterPayload(frontDeleteCharacterName string) (*front.DeleteCharacterPayload, error) {
	var name string
	{
		name = frontDeleteCharacterName
	}
	v := &front.DeleteCharacterPayload{}
	v.Name = name

	return v, nil
}

// BuildCreateItemPayload builds the payload for the Front createItem endpoint
// from CLI flags.
func BuildCreateItemPayload(frontCreateItemBody string) (*front.Item, error) {
	var err error
	var body CreateItemRequestBody
	{
		err = json.Unmarshal([]byte(frontCreateItemBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"damage\": 344029,\n      \"description\": \"Restores health\",\n      \"healing\": 49233,\n      \"name\": \"Potion\",\n      \"protection\": 6286\n   }'")
		}
		if utf8.RuneCountInString(body.Name) > 32 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 32, false))
		}
		if utf8.RuneCountInString(body.Description) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", body.Description, utf8.RuneCountInString(body.Description), 128, false))
		}
		if body.Damage < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.damage", body.Damage, 0, true))
		}
		if body.Damage > 1e+06 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.damage", body.Damage, 1e+06, false))
		}
		if body.Healing < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.healing", body.Healing, 0, true))
		}
		if body.Healing > 1e+06 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.healing", body.Healing, 1e+06, false))
		}
		if body.Protection < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.protection", body.Protection, 0, true))
		}
		if body.Protection > 10000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.protection", body.Protection, 10000, false))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &front.Item{
		Name:        body.Name,
		Description: body.Description,
		Damage:      body.Damage,
		Healing:     body.Healing,
		Protection:  body.Protection,
	}
	{
		var zero uint32
		if v.Damage == zero {
			v.Damage = 0
		}
	}
	{
		var zero uint32
		if v.Healing == zero {
			v.Healing = 0
		}
	}
	{
		var zero uint32
		if v.Protection == zero {
			v.Protection = 0
		}
	}

	return v, nil
}

// BuildGetItemPayload builds the payload for the Front getItem endpoint from
// CLI flags.
func BuildGetItemPayload(frontGetItemName string) (*front.GetItemPayload, error) {
	var name string
	{
		name = frontGetItemName
	}
	v := &front.GetItemPayload{}
	v.Name = name

	return v, nil
}

// BuildUpdateItemPayload builds the payload for the Front updateItem endpoint
// from CLI flags.
func BuildUpdateItemPayload(frontUpdateItemBody string, frontUpdateItemName string) (*front.Item, error) {
	var err error
	var body UpdateItemRequestBody
	{
		err = json.Unmarshal([]byte(frontUpdateItemBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"damage\": 797739,\n      \"description\": \"Restores health\",\n      \"healing\": 473315,\n      \"protection\": 7220\n   }'")
		}
		if utf8.RuneCountInString(body.Description) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", body.Description, utf8.RuneCountInString(body.Description), 128, false))
		}
		if body.Damage < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.damage", body.Damage, 0, true))
		}
		if body.Damage > 1e+06 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.damage", body.Damage, 1e+06, false))
		}
		if body.Healing < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.healing", body.Healing, 0, true))
		}
		if body.Healing > 1e+06 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.healing", body.Healing, 1e+06, false))
		}
		if body.Protection < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.protection", body.Protection, 0, true))
		}
		if body.Protection > 10000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.protection", body.Protection, 10000, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var name string
	{
		name = frontUpdateItemName
		if utf8.RuneCountInString(name) > 32 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("name", name, utf8.RuneCountInString(name), 32, false))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &front.Item{
		Description: body.Description,
		Damage:      body.Damage,
		Healing:     body.Healing,
		Protection:  body.Protection,
	}
	{
		var zero uint32
		if v.Damage == zero {
			v.Damage = 0
		}
	}
	{
		var zero uint32
		if v.Healing == zero {
			v.Healing = 0
		}
	}
	{
		var zero uint32
		if v.Protection == zero {
			v.Protection = 0
		}
	}
	v.Name = name

	return v, nil
}

// BuildDeleteItemPayload builds the payload for the Front deleteItem endpoint
// from CLI flags.
func BuildDeleteItemPayload(frontDeleteItemName string) (*front.DeleteItemPayload, error) {
	var name string
	{
		name = frontDeleteItemName
	}
	v := &front.DeleteItemPayload{}
	v.Name = name

	return v, nil
}

// BuildAddInventoryItemPayload builds the payload for the Front
// addInventoryItem endpoint from CLI flags.
func BuildAddInventoryItemPayload(frontAddInventoryItemBody string, frontAddInventoryItemCharacter string, frontAddInventoryItemItem string) (*front.AddInventoryItemPayload, error) {
	var err error
	var body AddInventoryItemRequestBody
	{
		err = json.Unmarshal([]byte(frontAddInventoryItemBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"count\": 1357457099\n   }'")
		}
	}
	var character string
	{
		character = frontAddInventoryItemCharacter
	}
	var item string
	{
		item = frontAddInventoryItemItem
	}
	v := &front.AddInventoryItemPayload{
		Count: body.Count,
	}
	v.Character = character
	v.Item = item

	return v, nil
}

// BuildRemoveInventoryItemPayload builds the payload for the Front
// removeInventoryItem endpoint from CLI flags.
func BuildRemoveInventoryItemPayload(frontRemoveInventoryItemBody string, frontRemoveInventoryItemCharacter string, frontRemoveInventoryItemItem string) (*front.RemoveInventoryItemPayload, error) {
	var err error
	var body RemoveInventoryItemRequestBody
	{
		err = json.Unmarshal([]byte(frontRemoveInventoryItemBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"count\": 2451514394\n   }'")
		}
	}
	var character string
	{
		character = frontRemoveInventoryItemCharacter
	}
	var item string
	{
		item = frontRemoveInventoryItemItem
	}
	v := &front.RemoveInventoryItemPayload{
		Count: body.Count,
	}
	v.Character = character
	v.Item = item

	return v, nil
}

// BuildGetInventoryPayload builds the payload for the Front getInventory
// endpoint from CLI flags.
func BuildGetInventoryPayload(frontGetInventoryCharacter string) (*front.GetInventoryPayload, error) {
	var character string
	{
		character = frontGetInventoryCharacter
	}
	v := &front.GetInventoryPayload{}
	v.Character = character

	return v, nil
}
