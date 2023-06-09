// Code generated by goa v3.11.2, DO NOT EDIT.
//
// Item gRPC client types
//
// Command:
// $ goa gen item/design

package client

import (
	itempb "item/gen/grpc/item/pb"
	item "item/gen/item"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// NewProtoCreateItemRequest builds the gRPC request type from the payload of
// the "createItem" endpoint of the "Item" service.
func NewProtoCreateItemRequest(payload *item.Item) *itempb.CreateItemRequest {
	message := &itempb.CreateItemRequest{
		Name:        payload.Name,
		Description: payload.Description,
		Damage:      &payload.Damage,
		Healing:     &payload.Healing,
		Protection:  &payload.Protection,
	}
	return message
}

// NewCreateItemResult builds the result type of the "createItem" endpoint of
// the "Item" service from the gRPC response type.
func NewCreateItemResult(message *itempb.CreateItemResponse) *item.Item {
	result := &item.Item{
		Name:        message.Name,
		Description: message.Description,
	}
	if message.Damage != nil {
		result.Damage = *message.Damage
	}
	if message.Healing != nil {
		result.Healing = *message.Healing
	}
	if message.Protection != nil {
		result.Protection = *message.Protection
	}
	if message.Damage == nil {
		result.Damage = 0
	}
	if message.Healing == nil {
		result.Healing = 0
	}
	if message.Protection == nil {
		result.Protection = 0
	}
	return result
}

// NewCreateItemItemAlreadyExistsError builds the error type of the
// "createItem" endpoint of the "Item" service from the gRPC error response
// type.
func NewCreateItemItemAlreadyExistsError(message *itempb.CreateItemItemAlreadyExistsError) *item.ItemAlreadyExists {
	er := &item.ItemAlreadyExists{
		Message: message.Message_,
		Name:    message.Name,
	}
	return er
}

// NewProtoGetItemRequest builds the gRPC request type from the payload of the
// "getItem" endpoint of the "Item" service.
func NewProtoGetItemRequest(payload *item.GetItemPayload) *itempb.GetItemRequest {
	message := &itempb.GetItemRequest{
		Name: payload.Name,
	}
	return message
}

// NewGetItemResult builds the result type of the "getItem" endpoint of the
// "Item" service from the gRPC response type.
func NewGetItemResult(message *itempb.GetItemResponse) *item.Item {
	result := &item.Item{
		Name:        message.Name,
		Description: message.Description,
	}
	if message.Damage != nil {
		result.Damage = *message.Damage
	}
	if message.Healing != nil {
		result.Healing = *message.Healing
	}
	if message.Protection != nil {
		result.Protection = *message.Protection
	}
	if message.Damage == nil {
		result.Damage = 0
	}
	if message.Healing == nil {
		result.Healing = 0
	}
	if message.Protection == nil {
		result.Protection = 0
	}
	return result
}

// NewGetItemItemNotFoundError builds the error type of the "getItem" endpoint
// of the "Item" service from the gRPC error response type.
func NewGetItemItemNotFoundError(message *itempb.GetItemItemNotFoundError) *item.ItemNotFound {
	er := &item.ItemNotFound{
		Message: message.Message_,
		Name:    message.Name,
	}
	return er
}

// NewProtoListItemsRequest builds the gRPC request type from the payload of
// the "listItems" endpoint of the "Item" service.
func NewProtoListItemsRequest() *itempb.ListItemsRequest {
	message := &itempb.ListItemsRequest{}
	return message
}

// NewListItemsResult builds the result type of the "listItems" endpoint of the
// "Item" service from the gRPC response type.
func NewListItemsResult(message *itempb.ListItemsResponse) []*item.Item {
	result := make([]*item.Item, len(message.Field))
	for i, val := range message.Field {
		result[i] = &item.Item{
			Name:        val.Name,
			Description: val.Description,
		}
		if val.Damage != nil {
			result[i].Damage = *val.Damage
		}
		if val.Healing != nil {
			result[i].Healing = *val.Healing
		}
		if val.Protection != nil {
			result[i].Protection = *val.Protection
		}
		if val.Damage == nil {
			result[i].Damage = 0
		}
		if val.Healing == nil {
			result[i].Healing = 0
		}
		if val.Protection == nil {
			result[i].Protection = 0
		}
	}
	return result
}

// NewProtoUpdateItemRequest builds the gRPC request type from the payload of
// the "updateItem" endpoint of the "Item" service.
func NewProtoUpdateItemRequest(payload *item.Item) *itempb.UpdateItemRequest {
	message := &itempb.UpdateItemRequest{
		Name:        payload.Name,
		Description: payload.Description,
		Damage:      &payload.Damage,
		Healing:     &payload.Healing,
		Protection:  &payload.Protection,
	}
	return message
}

// NewUpdateItemResult builds the result type of the "updateItem" endpoint of
// the "Item" service from the gRPC response type.
func NewUpdateItemResult(message *itempb.UpdateItemResponse) *item.Item {
	result := &item.Item{
		Name:        message.Name,
		Description: message.Description,
	}
	if message.Damage != nil {
		result.Damage = *message.Damage
	}
	if message.Healing != nil {
		result.Healing = *message.Healing
	}
	if message.Protection != nil {
		result.Protection = *message.Protection
	}
	if message.Damage == nil {
		result.Damage = 0
	}
	if message.Healing == nil {
		result.Healing = 0
	}
	if message.Protection == nil {
		result.Protection = 0
	}
	return result
}

// NewUpdateItemItemNotFoundError builds the error type of the "updateItem"
// endpoint of the "Item" service from the gRPC error response type.
func NewUpdateItemItemNotFoundError(message *itempb.UpdateItemItemNotFoundError) *item.ItemNotFound {
	er := &item.ItemNotFound{
		Message: message.Message_,
		Name:    message.Name,
	}
	return er
}

// NewProtoDeleteItemRequest builds the gRPC request type from the payload of
// the "deleteItem" endpoint of the "Item" service.
func NewProtoDeleteItemRequest(payload *item.DeleteItemPayload) *itempb.DeleteItemRequest {
	message := &itempb.DeleteItemRequest{
		Name: payload.Name,
	}
	return message
}

// ValidateCreateItemResponse runs the validations defined on
// CreateItemResponse.
func ValidateCreateItemResponse(message *itempb.CreateItemResponse) (err error) {
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

// ValidateGetItemResponse runs the validations defined on GetItemResponse.
func ValidateGetItemResponse(message *itempb.GetItemResponse) (err error) {
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

// ValidateListItemsResponse runs the validations defined on ListItemsResponse.
func ValidateListItemsResponse(message *itempb.ListItemsResponse) (err error) {
	for _, e := range message.Field {
		if e != nil {
			if err2 := ValidateItem2(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateItem2 runs the validations defined on Item2.
func ValidateItem2(elem *itempb.Item2) (err error) {
	if utf8.RuneCountInString(elem.Name) > 32 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("elem.name", elem.Name, utf8.RuneCountInString(elem.Name), 32, false))
	}
	if utf8.RuneCountInString(elem.Description) > 128 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("elem.description", elem.Description, utf8.RuneCountInString(elem.Description), 128, false))
	}
	if elem.Damage != nil {
		if *elem.Damage < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("elem.damage", *elem.Damage, 0, true))
		}
	}
	if elem.Damage != nil {
		if *elem.Damage > 1e+06 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("elem.damage", *elem.Damage, 1e+06, false))
		}
	}
	if elem.Healing != nil {
		if *elem.Healing < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("elem.healing", *elem.Healing, 0, true))
		}
	}
	if elem.Healing != nil {
		if *elem.Healing > 1e+06 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("elem.healing", *elem.Healing, 1e+06, false))
		}
	}
	if elem.Protection != nil {
		if *elem.Protection < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("elem.protection", *elem.Protection, 0, true))
		}
	}
	if elem.Protection != nil {
		if *elem.Protection > 10000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("elem.protection", *elem.Protection, 10000, false))
		}
	}
	return
}

// ValidateUpdateItemResponse runs the validations defined on
// UpdateItemResponse.
func ValidateUpdateItemResponse(message *itempb.UpdateItemResponse) (err error) {
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
