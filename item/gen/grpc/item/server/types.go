// Code generated by goa v3.11.2, DO NOT EDIT.
//
// Item gRPC server types
//
// Command:
// $ goa gen item/design

package server

import (
	itempb "item/gen/grpc/item/pb"
	item "item/gen/item"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// NewCreateItemPayload builds the payload of the "createItem" endpoint of the
// "Item" service from the gRPC request type.
func NewCreateItemPayload(message *itempb.CreateItemRequest) *item.Item {
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
	return v
}

// NewProtoCreateItemResponse builds the gRPC response type from the result of
// the "createItem" endpoint of the "Item" service.
func NewProtoCreateItemResponse(result *item.Item) *itempb.CreateItemResponse {
	message := &itempb.CreateItemResponse{
		Name:        result.Name,
		Description: result.Description,
		Damage:      &result.Damage,
		Healing:     &result.Healing,
		Protection:  &result.Protection,
	}
	return message
}

// NewCreateItemItemAlreadyExistsError builds the gRPC error response type from
// the error of the "createItem" endpoint of the "Item" service.
func NewCreateItemItemAlreadyExistsError(er *item.ItemAlreadyExists) *itempb.CreateItemItemAlreadyExistsError {
	message := &itempb.CreateItemItemAlreadyExistsError{
		Message_: er.Message,
		Name:     er.Name,
	}
	return message
}

// NewGetItemPayload builds the payload of the "getItem" endpoint of the "Item"
// service from the gRPC request type.
func NewGetItemPayload(message *itempb.GetItemRequest) *item.GetItemPayload {
	v := &item.GetItemPayload{
		Name: message.Name,
	}
	return v
}

// NewProtoGetItemResponse builds the gRPC response type from the result of the
// "getItem" endpoint of the "Item" service.
func NewProtoGetItemResponse(result *item.Item) *itempb.GetItemResponse {
	message := &itempb.GetItemResponse{
		Name:        result.Name,
		Description: result.Description,
		Damage:      &result.Damage,
		Healing:     &result.Healing,
		Protection:  &result.Protection,
	}
	return message
}

// NewGetItemItemNotFoundError builds the gRPC error response type from the
// error of the "getItem" endpoint of the "Item" service.
func NewGetItemItemNotFoundError(er *item.ItemNotFound) *itempb.GetItemItemNotFoundError {
	message := &itempb.GetItemItemNotFoundError{
		Message_: er.Message,
		Name:     er.Name,
	}
	return message
}

// NewProtoListItemsResponse builds the gRPC response type from the result of
// the "listItems" endpoint of the "Item" service.
func NewProtoListItemsResponse(result []*item.Item) *itempb.ListItemsResponse {
	message := &itempb.ListItemsResponse{}
	message.Field = make([]*itempb.Item2, len(result))
	for i, val := range result {
		message.Field[i] = &itempb.Item2{
			Name:        val.Name,
			Description: val.Description,
			Damage:      &val.Damage,
			Healing:     &val.Healing,
			Protection:  &val.Protection,
		}
	}
	return message
}

// NewUpdateItemPayload builds the payload of the "updateItem" endpoint of the
// "Item" service from the gRPC request type.
func NewUpdateItemPayload(message *itempb.UpdateItemRequest) *item.Item {
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
	return v
}

// NewProtoUpdateItemResponse builds the gRPC response type from the result of
// the "updateItem" endpoint of the "Item" service.
func NewProtoUpdateItemResponse(result *item.Item) *itempb.UpdateItemResponse {
	message := &itempb.UpdateItemResponse{
		Name:        result.Name,
		Description: result.Description,
		Damage:      &result.Damage,
		Healing:     &result.Healing,
		Protection:  &result.Protection,
	}
	return message
}

// NewUpdateItemItemNotFoundError builds the gRPC error response type from the
// error of the "updateItem" endpoint of the "Item" service.
func NewUpdateItemItemNotFoundError(er *item.ItemNotFound) *itempb.UpdateItemItemNotFoundError {
	message := &itempb.UpdateItemItemNotFoundError{
		Message_: er.Message,
		Name:     er.Name,
	}
	return message
}

// NewDeleteItemPayload builds the payload of the "deleteItem" endpoint of the
// "Item" service from the gRPC request type.
func NewDeleteItemPayload(message *itempb.DeleteItemRequest) *item.DeleteItemPayload {
	v := &item.DeleteItemPayload{
		Name: message.Name,
	}
	return v
}

// NewProtoDeleteItemResponse builds the gRPC response type from the result of
// the "deleteItem" endpoint of the "Item" service.
func NewProtoDeleteItemResponse() *itempb.DeleteItemResponse {
	message := &itempb.DeleteItemResponse{}
	return message
}

// ValidateCreateItemRequest runs the validations defined on CreateItemRequest.
func ValidateCreateItemRequest(message *itempb.CreateItemRequest) (err error) {
	if utf8.RuneCountInString(message.Name) > 32 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 32, false))
	}
	if utf8.RuneCountInString(message.Description) > 128 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.description", message.Description, utf8.RuneCountInString(message.Description), 128, false))
	}
	if message.Damage != nil {
		if *message.Damage < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.damage", *message.Damage, 0, true))
		}
	}
	if message.Damage != nil {
		if *message.Damage > 1e+06 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.damage", *message.Damage, 1e+06, false))
		}
	}
	if message.Healing != nil {
		if *message.Healing < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.healing", *message.Healing, 0, true))
		}
	}
	if message.Healing != nil {
		if *message.Healing > 1e+06 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.healing", *message.Healing, 1e+06, false))
		}
	}
	if message.Protection != nil {
		if *message.Protection < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.protection", *message.Protection, 0, true))
		}
	}
	if message.Protection != nil {
		if *message.Protection > 10000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.protection", *message.Protection, 10000, false))
		}
	}
	return
}

// ValidateUpdateItemRequest runs the validations defined on UpdateItemRequest.
func ValidateUpdateItemRequest(message *itempb.UpdateItemRequest) (err error) {
	if utf8.RuneCountInString(message.Name) > 32 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 32, false))
	}
	if utf8.RuneCountInString(message.Description) > 128 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.description", message.Description, utf8.RuneCountInString(message.Description), 128, false))
	}
	if message.Damage != nil {
		if *message.Damage < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.damage", *message.Damage, 0, true))
		}
	}
	if message.Damage != nil {
		if *message.Damage > 1e+06 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.damage", *message.Damage, 1e+06, false))
		}
	}
	if message.Healing != nil {
		if *message.Healing < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.healing", *message.Healing, 0, true))
		}
	}
	if message.Healing != nil {
		if *message.Healing > 1e+06 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.healing", *message.Healing, 1e+06, false))
		}
	}
	if message.Protection != nil {
		if *message.Protection < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.protection", *message.Protection, 0, true))
		}
	}
	if message.Protection != nil {
		if *message.Protection > 10000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.protection", *message.Protection, 10000, false))
		}
	}
	return
}
