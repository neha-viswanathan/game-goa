package frontapi

import (
	characterpb "character/gen/grpc/character/pb"
	"context"
	"fmt"
	front "front/gen/front"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	inventorypb "inventory/gen/grpc/inventory/pb"
	itempb "item/gen/grpc/item/pb"
	"log"
	"strings"
)

// Front service example implementation.
// The example methods log the requests and return zero values.
type frontsrvc struct {
	logger          *log.Logger
	characterClient characterpb.CharacterClient
	itemClient      itempb.ItemClient
	inventoryClient inventorypb.InventoryClient
}

// NewFront returns the Front service implementation.
func NewFront(logger *log.Logger, characterClient characterpb.CharacterClient, itemClient itempb.ItemClient, inventoryClient inventorypb.InventoryClient) (front.Service, error) {
	return &frontsrvc{
		logger:          logger,
		characterClient: characterClient,
		itemClient:      itemClient,
		inventoryClient: inventoryClient,
	}, nil
}

// CreateCharacter Create a new character
func (s *frontsrvc) CreateCharacter(ctx context.Context, p *front.Character) (res *front.Character, err error) {
	s.logger.Print("front.create.character")
	createResponse, err := s.characterClient.CreateCharacter(ctx, &characterpb.CreateCharacterRequest{
		Name:        p.Name,
		Description: &p.Description,
		Health:      &p.Health,
		Experience:  &p.Experience,
	})

	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			return nil, fmt.Errorf("unsupported error: %v", err)
		}
		switch status.Code() {
		case codes.AlreadyExists:
			return nil, &front.CharacterAlreadyExists{Message: "character already exists"}
		default:
			return nil, fmt.Errorf("unsupported code: %d, %v", status.Code(), err)
		}
	}

	res = &front.Character{
		Name:        createResponse.Name,
		Description: *createResponse.Description,
		Health:      *createResponse.Health,
		Experience:  *createResponse.Experience,
	}
	return res, nil
}

// GetCharacter Get a character by name
func (s *frontsrvc) GetCharacter(ctx context.Context, p *front.GetCharacterPayload) (res *front.Character, err error) {
	s.logger.Print("front.get.character")
	getResponse, err := s.characterClient.GetCharacter(ctx, &characterpb.GetCharacterRequest{
		Name: p.Name,
	})
	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			return nil, fmt.Errorf("unsupported error: %v", err)
		}
		switch status.Code() {
		case codes.NotFound:
			return nil, &front.CharacterNotFound{Message: "character not found"}
		default:
			return nil, fmt.Errorf("unsupported code: %d, %v", status.Code(), err)
		}
	}
	res = &front.Character{
		Name:        getResponse.Name,
		Description: *getResponse.Description,
		Health:      *getResponse.Health,
		Experience:  *getResponse.Experience,
	}
	return res, nil
}

// ListCharacters List all characters
func (s *frontsrvc) ListCharacters(ctx context.Context) (res []*front.Character, err error) {
	s.logger.Print("front.list.character")
	listResponse, err := s.characterClient.ListCharacters(ctx, &characterpb.ListCharactersRequest{})
	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			return nil, fmt.Errorf("unsupported error: %v", err)
		}
		switch status.Code() {
		default:
			return nil, fmt.Errorf("unsupported code: %d, %v", status.Code(), err)
		}
	}
	for _, entry := range listResponse.Field {
		c := &front.Character{
			Name:        entry.Name,
			Description: *entry.Description,
			Health:      *entry.Health,
			Experience:  *entry.Experience,
		}
		res = append(res, c)
	}
	return res, nil
}

// UpdateCharacter Update a given character
func (s *frontsrvc) UpdateCharacter(ctx context.Context, p *front.Character) (res *front.Character, err error) {
	s.logger.Print("front.update.character")
	updateResponse, err := s.characterClient.UpdateCharacter(ctx, &characterpb.UpdateCharacterRequest{
		Name:        p.Name,
		Description: &p.Description,
		Health:      &p.Health,
		Experience:  &p.Experience,
	})
	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			return nil, fmt.Errorf("unsupported error: %v", err)
		}
		switch status.Code() {
		case codes.NotFound:
			return nil, &front.CharacterNotFound{Message: "character not found"}
		default:
			return nil, fmt.Errorf("unsupported code: %d, %v", status.Code(), err)
		}
	}
	res = &front.Character{
		Name:        updateResponse.Name,
		Description: *updateResponse.Description,
		Health:      *updateResponse.Health,
		Experience:  *updateResponse.Experience,
	}
	return res, nil
}

// DeleteCharacter Delete a given character
func (s *frontsrvc) DeleteCharacter(ctx context.Context, p *front.DeleteCharacterPayload) (err error) {
	s.logger.Print("front.delete.character")
	_, err = s.characterClient.DeleteCharacter(ctx, &characterpb.DeleteCharacterRequest{
		Name: p.Name,
	})
	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			return fmt.Errorf("unsupported error: %v", err)
		}
		switch status.Code() {
		default:
			return fmt.Errorf("unsupported code: %d, %v", status.Code(), err)
		}
	}
	return nil
}

// CreateItem Create a new item
func (s *frontsrvc) CreateItem(ctx context.Context, item *front.Item) (res *front.Item, err error) {
	s.logger.Print("front.create.item")
	createResponse, err := s.itemClient.CreateItem(ctx, &itempb.CreateItemRequest{
		Name:        item.Name,
		Description: item.Description,
		Damage:      &item.Damage,
		Healing:     &item.Healing,
		Protection:  &item.Protection,
	})
	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			return nil, fmt.Errorf("unsupported error: %v", err)
		}
		switch status.Code() {
		case codes.AlreadyExists:
			return nil, &front.ItemAlreadyExists{Message: "item already exists"}
		default:
			return nil, fmt.Errorf("unsupported code: %d, %v", status.Code(), err)
		}
	}
	res = &front.Item{
		Name:        createResponse.Name,
		Description: createResponse.Description,
		Damage:      *createResponse.Damage,
		Healing:     *createResponse.Healing,
		Protection:  *createResponse.Protection,
	}
	return res, nil
}

// GetItem Get an item by name
func (s *frontsrvc) GetItem(ctx context.Context, payload *front.GetItemPayload) (res *front.Item, err error) {
	s.logger.Print("front.get.item")
	getResponse, err := s.itemClient.GetItem(ctx, &itempb.GetItemRequest{
		Name: payload.Name,
	})
	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			return nil, fmt.Errorf("unsupported error: %v", err)
		}
		switch status.Code() {
		case codes.NotFound:
			return nil, &front.ItemNotFound{Message: "item not found"}
		default:
			return nil, fmt.Errorf("unsupported code: %d, %v", status.Code(), err)
		}
	}
	res = &front.Item{
		Name:        getResponse.Name,
		Description: getResponse.Description,
		Damage:      *getResponse.Damage,
		Healing:     *getResponse.Healing,
		Protection:  *getResponse.Protection,
	}
	return res, nil
}

// ListItems List all items
func (s *frontsrvc) ListItems(ctx context.Context) (res []*front.Item, err error) {
	s.logger.Print("front.list.item")
	listResponse, err := s.itemClient.ListItems(ctx, &itempb.ListItemsRequest{})
	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			return nil, fmt.Errorf("unsupported error: %v", err)
		}
		switch status.Code() {
		default:
			return nil, fmt.Errorf("unsupported code: %d, %v", status.Code(), err)
		}
	}
	for _, entry := range listResponse.Field {
		i := &front.Item{
			Name:        entry.Name,
			Description: entry.Description,
			Damage:      *entry.Damage,
			Healing:     *entry.Healing,
			Protection:  *entry.Protection,
		}
		res = append(res, i)
	}
	return res, nil
}

// UpdateItem Update a given item
func (s *frontsrvc) UpdateItem(ctx context.Context, item *front.Item) (res *front.Item, err error) {
	s.logger.Print("front.update.item")
	updateResponse, err := s.itemClient.UpdateItem(ctx, &itempb.UpdateItemRequest{
		Name:        item.Name,
		Description: item.Description,
		Damage:      &item.Damage,
		Healing:     &item.Healing,
		Protection:  &item.Protection,
	})
	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			return nil, fmt.Errorf("unsupported error: %v", err)
		}
		switch status.Code() {
		case codes.NotFound:
			return nil, &front.ItemNotFound{Message: "item not found"}
		default:
			return nil, fmt.Errorf("unsupported code: %d, %v", status.Code(), err)
		}
	}
	res = &front.Item{
		Name:        updateResponse.Name,
		Description: updateResponse.Description,
		Damage:      *updateResponse.Damage,
		Healing:     *updateResponse.Healing,
		Protection:  *updateResponse.Protection,
	}
	return res, nil
}

// DeleteItem Delete a given item
func (s *frontsrvc) DeleteItem(ctx context.Context, payload *front.DeleteItemPayload) (err error) {
	s.logger.Print("front.delete.item")
	_, err = s.itemClient.DeleteItem(ctx, &itempb.DeleteItemRequest{
		Name: payload.Name,
	})
	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			return fmt.Errorf("unsupported error: %v", err)
		}
		switch status.Code() {
		default:
			return fmt.Errorf("unsupported code: %d, %v", status.Code(), err)
		}
	}
	return nil
}

// AddInventoryItem Add an item to a character's inventory
func (s *frontsrvc) AddInventoryItem(ctx context.Context, p *front.AddInventoryItemPayload) (err error) {
	s.logger.Print("front.addItem")
	_, err = s.inventoryClient.AddItem(ctx, &inventorypb.AddItemRequest{
		Character: p.Character,
		Item:      p.Item,
	})
	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			return fmt.Errorf("unsupported error: %v", err)
		}
		switch status.Code() {
		case codes.NotFound:
			switch {
			case strings.Contains(status.Message(), "character"):
				return &front.CharacterNotFound{Message: "character not found"}
			case strings.Contains(status.Message(), "item"):
				return &front.ItemNotFound{Message: "item not found"}
			}
		default:
			return fmt.Errorf("unsupported code: %d, %v", status.Code(), err)
		}
	}
	return nil
}

// RemoveInventoryItem Remove an item from a character's inventory
func (s *frontsrvc) RemoveInventoryItem(ctx context.Context, p *front.RemoveInventoryItemPayload) (err error) {
	s.logger.Print("front.removeItem")
	_, err = s.inventoryClient.RemoveItem(ctx, &inventorypb.RemoveItemRequest{
		Character: p.Character,
		Item:      p.Item,
	})
	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			return fmt.Errorf("unsupported error: %v", err)
		}
		switch status.Code() {
		case codes.NotFound:
			return &front.CharacterNotFound{Message: "character not found"}
		default:
			return fmt.Errorf("unsupported code: %d, %v", status.Code(), err)
		}
	}
	return nil
}

// GetInventory Get a character's inventory
func (s *frontsrvc) GetInventory(ctx context.Context, p *front.GetInventoryPayload) (res []string, err error) {
	s.logger.Print("front.getInventory")
	getResponse, err := s.inventoryClient.GetInventory(ctx, &inventorypb.GetInventoryRequest{
		Character: p.Character,
	})
	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			return nil, fmt.Errorf("unsupported error: %v", err)
		}
		switch status.Code() {
		case codes.NotFound:
			return nil, &front.CharacterNotFound{Message: "character not found"}
		default:
			return nil, fmt.Errorf("unsupported code: %d, %v", status.Code(), err)
		}
	}

	for _, entry := range getResponse.Field {
		res = append(res, entry)
	}

	return res, nil
}
