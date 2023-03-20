package inventory

import (
	"context"
	"log"

	"inventory/gen/inventory"

	characterpb "character/gen/grpc/character/pb"
	itempb "item/gen/grpc/item/pb"
)

// inventory service example implementation.
// The example methods log the requests and return zero values.
type inventorysrvc struct {
	logger          *log.Logger
	inventories     map[string]map[string]bool
	characterClient characterpb.CharacterClient
	itemClient      itempb.ItemClient
}

// NewInventory returns the inventory service implementation.
func NewInventory(logger *log.Logger, characterClient characterpb.CharacterClient, itemClient itempb.ItemClient) (inventory.Service, error) {

	return &inventorysrvc{
		logger:          logger,
		inventories:     map[string]map[string]bool{},
		characterClient: characterClient,
		itemClient:      itemClient,
	}, nil
}

// AddItem Add an item to a character's inventory
func (s *inventorysrvc) AddItem(ctx context.Context, p *inventory.AddItemPayload) (err error) {
	s.logger.Print("inventory.addItem")

	// Check if character exists
	_, err = s.characterClient.GetCharacter(ctx, &characterpb.GetCharacterRequest{Name: p.Character})
	if err != nil {
		s.logger.Print("error:", err)
		return &inventory.CharacterNotFound{Message: "character not found"}
	}

	// Check if item exists
	_, err = s.itemClient.GetItem(ctx, &itempb.GetItemRequest{Name: p.Item})
	if err != nil {
		s.logger.Print("error:", err)
		return &inventory.ItemNotFound{Message: "item not found"}
	}

	// Check if character has an inventory
	_, found := s.inventories[p.Character]
	if !found {
		s.inventories[p.Character] = map[string]bool{}
	}

	s.inventories[p.Character][p.Item] = true

	return
}

// RemoveItem Remove an item from a character's inventory
func (s *inventorysrvc) RemoveItem(ctx context.Context, p *inventory.RemoveItemPayload) (err error) {
	s.logger.Print("inventory.removeItem")

	// Check if character exists
	_, err = s.characterClient.GetCharacter(ctx, &characterpb.GetCharacterRequest{Name: p.Character})
	if err != nil {
		s.logger.Print("error:", err)
		return &inventory.CharacterNotFound{Message: "character not found"}
	}

	delete(s.inventories[p.Character], p.Item)

	return nil
}

// GetInventory Get a character's inventory
func (s *inventorysrvc) GetInventory(ctx context.Context, p *inventory.GetInventoryPayload) (res []string, err error) {
	s.logger.Print("inventory.getInventory")

	// Check if character exists
	_, err = s.characterClient.GetCharacter(ctx, &characterpb.GetCharacterRequest{Name: p.Character})
	if err != nil {
		s.logger.Print("error:", err)
		return nil, &inventory.CharacterNotFound{Message: "character not found"}
	}

	// Build output by iterating over a character's inventory
	for item := range s.inventories[p.Character] {
		res = append(res, item)
	}

	return res, nil
}
