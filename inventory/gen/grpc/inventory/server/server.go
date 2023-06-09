// Code generated by goa v3.11.2, DO NOT EDIT.
//
// Inventory gRPC server
//
// Command:
// $ goa gen inventory/design

package server

import (
	"context"
	"errors"
	inventorypb "inventory/gen/grpc/inventory/pb"
	inventory "inventory/gen/inventory"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the inventorypb.InventoryServer interface.
type Server struct {
	AddItemH      goagrpc.UnaryHandler
	RemoveItemH   goagrpc.UnaryHandler
	GetInventoryH goagrpc.UnaryHandler
	inventorypb.UnimplementedInventoryServer
}

// New instantiates the server struct with the Inventory service endpoints.
func New(e *inventory.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		AddItemH:      NewAddItemHandler(e.AddItem, uh),
		RemoveItemH:   NewRemoveItemHandler(e.RemoveItem, uh),
		GetInventoryH: NewGetInventoryHandler(e.GetInventory, uh),
	}
}

// NewAddItemHandler creates a gRPC handler which serves the "Inventory"
// service "addItem" endpoint.
func NewAddItemHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeAddItemRequest, EncodeAddItemResponse)
	}
	return h
}

// AddItem implements the "AddItem" method in inventorypb.InventoryServer
// interface.
func (s *Server) AddItem(ctx context.Context, message *inventorypb.AddItemRequest) (*inventorypb.AddItemResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "addItem")
	ctx = context.WithValue(ctx, goa.ServiceKey, "Inventory")
	resp, err := s.AddItemH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "CharacterNotFound":
				var er *inventory.CharacterNotFound
				errors.As(err, &er)
				return nil, goagrpc.NewStatusError(codes.NotFound, err, NewAddItemCharacterNotFoundError(er))
			case "ItemNotFound":
				var er *inventory.ItemNotFound
				errors.As(err, &er)
				return nil, goagrpc.NewStatusError(codes.NotFound, err, NewAddItemItemNotFoundError(er))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventorypb.AddItemResponse), nil
}

// NewRemoveItemHandler creates a gRPC handler which serves the "Inventory"
// service "removeItem" endpoint.
func NewRemoveItemHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeRemoveItemRequest, EncodeRemoveItemResponse)
	}
	return h
}

// RemoveItem implements the "RemoveItem" method in inventorypb.InventoryServer
// interface.
func (s *Server) RemoveItem(ctx context.Context, message *inventorypb.RemoveItemRequest) (*inventorypb.RemoveItemResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "removeItem")
	ctx = context.WithValue(ctx, goa.ServiceKey, "Inventory")
	resp, err := s.RemoveItemH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "CharacterNotFound":
				var er *inventory.CharacterNotFound
				errors.As(err, &er)
				return nil, goagrpc.NewStatusError(codes.NotFound, err, NewRemoveItemCharacterNotFoundError(er))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventorypb.RemoveItemResponse), nil
}

// NewGetInventoryHandler creates a gRPC handler which serves the "Inventory"
// service "getInventory" endpoint.
func NewGetInventoryHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeGetInventoryRequest, EncodeGetInventoryResponse)
	}
	return h
}

// GetInventory implements the "GetInventory" method in
// inventorypb.InventoryServer interface.
func (s *Server) GetInventory(ctx context.Context, message *inventorypb.GetInventoryRequest) (*inventorypb.GetInventoryResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "getInventory")
	ctx = context.WithValue(ctx, goa.ServiceKey, "Inventory")
	resp, err := s.GetInventoryH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "CharacterNotFound":
				var er *inventory.CharacterNotFound
				errors.As(err, &er)
				return nil, goagrpc.NewStatusError(codes.NotFound, err, NewGetInventoryCharacterNotFoundError(er))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventorypb.GetInventoryResponse), nil
}
