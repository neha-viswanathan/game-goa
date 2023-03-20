// Code generated by goa v3.11.2, DO NOT EDIT.
//
// Character client
//
// Command:
// $ goa gen character/design

package character

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "Character" service client.
type Client struct {
	CreateCharacterEndpoint goa.Endpoint
	GetCharacterEndpoint    goa.Endpoint
	ListCharactersEndpoint  goa.Endpoint
	UpdateCharacterEndpoint goa.Endpoint
	DeleteCharacterEndpoint goa.Endpoint
}

// NewClient initializes a "Character" service client given the endpoints.
func NewClient(createCharacter, getCharacter, listCharacters, updateCharacter, deleteCharacter goa.Endpoint) *Client {
	return &Client{
		CreateCharacterEndpoint: createCharacter,
		GetCharacterEndpoint:    getCharacter,
		ListCharactersEndpoint:  listCharacters,
		UpdateCharacterEndpoint: updateCharacter,
		DeleteCharacterEndpoint: deleteCharacter,
	}
}

// CreateCharacter calls the "createCharacter" endpoint of the "Character"
// service.
// CreateCharacter may return the following errors:
//   - "CharacterAlreadyExists" (type *CharacterAlreadyExists): character already exists
//   - "CharacterNotFound" (type *CharacterNotFound)
//   - error: internal error
func (c *Client) CreateCharacter(ctx context.Context, p *Character) (res *Character, err error) {
	var ires interface{}
	ires, err = c.CreateCharacterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Character), nil
}

// GetCharacter calls the "getCharacter" endpoint of the "Character" service.
// GetCharacter may return the following errors:
//   - "CharacterNotFound" (type *CharacterNotFound): character not found
//   - "CharacterAlreadyExists" (type *CharacterAlreadyExists)
//   - error: internal error
func (c *Client) GetCharacter(ctx context.Context, p *GetCharacterPayload) (res *Character, err error) {
	var ires interface{}
	ires, err = c.GetCharacterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Character), nil
}

// ListCharacters calls the "listCharacters" endpoint of the "Character"
// service.
// ListCharacters may return the following errors:
//   - "CharacterAlreadyExists" (type *CharacterAlreadyExists)
//   - "CharacterNotFound" (type *CharacterNotFound)
//   - error: internal error
func (c *Client) ListCharacters(ctx context.Context) (res []*Character, err error) {
	var ires interface{}
	ires, err = c.ListCharactersEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]*Character), nil
}

// UpdateCharacter calls the "updateCharacter" endpoint of the "Character"
// service.
// UpdateCharacter may return the following errors:
//   - "CharacterNotFound" (type *CharacterNotFound): character not found
//   - "CharacterAlreadyExists" (type *CharacterAlreadyExists)
//   - error: internal error
func (c *Client) UpdateCharacter(ctx context.Context, p *Character) (res *Character, err error) {
	var ires interface{}
	ires, err = c.UpdateCharacterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Character), nil
}

// DeleteCharacter calls the "deleteCharacter" endpoint of the "Character"
// service.
// DeleteCharacter may return the following errors:
//   - "CharacterAlreadyExists" (type *CharacterAlreadyExists)
//   - "CharacterNotFound" (type *CharacterNotFound)
//   - error: internal error
func (c *Client) DeleteCharacter(ctx context.Context, p *DeleteCharacterPayload) (err error) {
	_, err = c.DeleteCharacterEndpoint(ctx, p)
	return
}