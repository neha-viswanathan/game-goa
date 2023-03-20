package item

import (
	"context"
	"item/gen/item"
	"log"
)

// item service example implementation.
// The example methods log the requests and return zero values.
type itemsrvc struct {
	logger *log.Logger
	items  map[string]*item.Item
}

// NewItem returns the item service implementation.
func NewItem(logger *log.Logger) item.Service {
	return &itemsrvc{
		logger: logger,
		items:  map[string]*item.Item{},
	}
}

// Create a new item
func (s *itemsrvc) CreateItem(ctx context.Context, p *item.Item) (res *item.Item, err error) {
	s.logger.Print("item.create")
	for name := range s.items {
		if name == p.Name {
			return nil, &item.ItemAlreadyExists{Message: "item already exists"}
		}
	}
	res = &item.Item{
		Name:        p.Name,
		Description: p.Description,
		Damage:      p.Damage,
		Healing:     p.Healing,
		Protection:  p.Protection,
	}
	s.items[p.Name] = res
	return res, nil
}

// GetItem Get an existing item
func (s *itemsrvc) GetItem(ctx context.Context, p *item.GetItemPayload) (res *item.Item, err error) {
	s.logger.Print("item.get", p)
	char, found := s.items[p.Name]
	if !found {
		return nil, &item.ItemNotFound{Message: "item not found"}
	}
	return char, nil
}

// ListItems List all items
func (s *itemsrvc) ListItems(ctx context.Context) (res []*item.Item, err error) {
	s.logger.Print("item.list")
	for _, item := range s.items {
		res = append(res, item)
	}
	return res, nil
}

// UpdateItem Update an existing item
func (s *itemsrvc) UpdateItem(ctx context.Context, p *item.Item) (res *item.Item, err error) {
	s.logger.Print("item.update")
	char, found := s.items[p.Name]
	if !found {
		return nil, &item.ItemNotFound{Message: "item not found"}
	}
	char.Description = p.Description
	char.Damage = p.Damage
	char.Healing = p.Healing
	char.Protection = p.Protection
	s.items[char.Name] = char

	return char, nil
}

// DeleteItem Delete item
func (s *itemsrvc) DeleteItem(ctx context.Context, p *item.DeleteItemPayload) (err error) {
	s.logger.Print("item.delete")
	delete(s.items, p.Name)
	s.logger.Printf("The item %d has been deleted successfully", p.Name)
	return nil
}
