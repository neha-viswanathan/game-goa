// Code generated by goa v3.11.2, DO NOT EDIT.
//
// Inventory client
//
// Command:
// $ goa gen inventory/design

package inventory

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "Inventory" service client.
type Client struct {
	AddItemEndpoint      goa.Endpoint
	RemoveItemEndpoint   goa.Endpoint
	GetInventoryEndpoint goa.Endpoint
}

// NewClient initializes a "Inventory" service client given the endpoints.
func NewClient(addItem, removeItem, getInventory goa.Endpoint) *Client {
	return &Client{
		AddItemEndpoint:      addItem,
		RemoveItemEndpoint:   removeItem,
		GetInventoryEndpoint: getInventory,
	}
}

// AddItem calls the "addItem" endpoint of the "Inventory" service.
// AddItem may return the following errors:
//   - "CharacterNotFound" (type *CharacterNotFound)
//   - "ItemNotFound" (type *ItemNotFound)
//   - error: internal error
func (c *Client) AddItem(ctx context.Context, p *AddItemPayload) (err error) {
	_, err = c.AddItemEndpoint(ctx, p)
	return
}

// RemoveItem calls the "removeItem" endpoint of the "Inventory" service.
// RemoveItem may return the following errors:
//   - "CharacterNotFound" (type *CharacterNotFound)
//   - "ItemNotFound" (type *ItemNotFound)
//   - error: internal error
func (c *Client) RemoveItem(ctx context.Context, p *RemoveItemPayload) (err error) {
	_, err = c.RemoveItemEndpoint(ctx, p)
	return
}

// GetInventory calls the "getInventory" endpoint of the "Inventory" service.
// GetInventory may return the following errors:
//   - "CharacterNotFound" (type *CharacterNotFound)
//   - "ItemNotFound" (type *ItemNotFound)
//   - error: internal error
func (c *Client) GetInventory(ctx context.Context, p *GetInventoryPayload) (res []string, err error) {
	var ires interface{}
	ires, err = c.GetInventoryEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]string), nil
}
