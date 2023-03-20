// Code generated by goa v3.11.2, DO NOT EDIT.
//
// Front endpoints
//
// Command:
// $ goa gen front/design

package front

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "Front" service endpoints.
type Endpoints struct {
	CreateCharacter     goa.Endpoint
	GetCharacter        goa.Endpoint
	ListCharacters      goa.Endpoint
	UpdateCharacter     goa.Endpoint
	DeleteCharacter     goa.Endpoint
	CreateItem          goa.Endpoint
	GetItem             goa.Endpoint
	ListItems           goa.Endpoint
	UpdateItem          goa.Endpoint
	DeleteItem          goa.Endpoint
	AddInventoryItem    goa.Endpoint
	RemoveInventoryItem goa.Endpoint
	GetInventory        goa.Endpoint
}

// NewEndpoints wraps the methods of the "Front" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		CreateCharacter:     NewCreateCharacterEndpoint(s),
		GetCharacter:        NewGetCharacterEndpoint(s),
		ListCharacters:      NewListCharactersEndpoint(s),
		UpdateCharacter:     NewUpdateCharacterEndpoint(s),
		DeleteCharacter:     NewDeleteCharacterEndpoint(s),
		CreateItem:          NewCreateItemEndpoint(s),
		GetItem:             NewGetItemEndpoint(s),
		ListItems:           NewListItemsEndpoint(s),
		UpdateItem:          NewUpdateItemEndpoint(s),
		DeleteItem:          NewDeleteItemEndpoint(s),
		AddInventoryItem:    NewAddInventoryItemEndpoint(s),
		RemoveInventoryItem: NewRemoveInventoryItemEndpoint(s),
		GetInventory:        NewGetInventoryEndpoint(s),
	}
}

// Use applies the given middleware to all the "Front" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.CreateCharacter = m(e.CreateCharacter)
	e.GetCharacter = m(e.GetCharacter)
	e.ListCharacters = m(e.ListCharacters)
	e.UpdateCharacter = m(e.UpdateCharacter)
	e.DeleteCharacter = m(e.DeleteCharacter)
	e.CreateItem = m(e.CreateItem)
	e.GetItem = m(e.GetItem)
	e.ListItems = m(e.ListItems)
	e.UpdateItem = m(e.UpdateItem)
	e.DeleteItem = m(e.DeleteItem)
	e.AddInventoryItem = m(e.AddInventoryItem)
	e.RemoveInventoryItem = m(e.RemoveInventoryItem)
	e.GetInventory = m(e.GetInventory)
}

// NewCreateCharacterEndpoint returns an endpoint function that calls the
// method "createCharacter" of service "Front".
func NewCreateCharacterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Character)
		return s.CreateCharacter(ctx, p)
	}
}

// NewGetCharacterEndpoint returns an endpoint function that calls the method
// "getCharacter" of service "Front".
func NewGetCharacterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetCharacterPayload)
		return s.GetCharacter(ctx, p)
	}
}

// NewListCharactersEndpoint returns an endpoint function that calls the method
// "listCharacters" of service "Front".
func NewListCharactersEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListCharacters(ctx)
	}
}

// NewUpdateCharacterEndpoint returns an endpoint function that calls the
// method "updateCharacter" of service "Front".
func NewUpdateCharacterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Character)
		return s.UpdateCharacter(ctx, p)
	}
}

// NewDeleteCharacterEndpoint returns an endpoint function that calls the
// method "deleteCharacter" of service "Front".
func NewDeleteCharacterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteCharacterPayload)
		return nil, s.DeleteCharacter(ctx, p)
	}
}

// NewCreateItemEndpoint returns an endpoint function that calls the method
// "createItem" of service "Front".
func NewCreateItemEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Item)
		return s.CreateItem(ctx, p)
	}
}

// NewGetItemEndpoint returns an endpoint function that calls the method
// "getItem" of service "Front".
func NewGetItemEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetItemPayload)
		return s.GetItem(ctx, p)
	}
}

// NewListItemsEndpoint returns an endpoint function that calls the method
// "listItems" of service "Front".
func NewListItemsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListItems(ctx)
	}
}

// NewUpdateItemEndpoint returns an endpoint function that calls the method
// "updateItem" of service "Front".
func NewUpdateItemEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Item)
		return s.UpdateItem(ctx, p)
	}
}

// NewDeleteItemEndpoint returns an endpoint function that calls the method
// "deleteItem" of service "Front".
func NewDeleteItemEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteItemPayload)
		return nil, s.DeleteItem(ctx, p)
	}
}

// NewAddInventoryItemEndpoint returns an endpoint function that calls the
// method "addInventoryItem" of service "Front".
func NewAddInventoryItemEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AddInventoryItemPayload)
		return nil, s.AddInventoryItem(ctx, p)
	}
}

// NewRemoveInventoryItemEndpoint returns an endpoint function that calls the
// method "removeInventoryItem" of service "Front".
func NewRemoveInventoryItemEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RemoveInventoryItemPayload)
		return nil, s.RemoveInventoryItem(ctx, p)
	}
}

// NewGetInventoryEndpoint returns an endpoint function that calls the method
// "getInventory" of service "Front".
func NewGetInventoryEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetInventoryPayload)
		return s.GetInventory(ctx, p)
	}
}