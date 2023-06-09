// Code generated by goa v3.11.2, DO NOT EDIT.
//
// Character gRPC client encoders and decoders
//
// Command:
// $ goa gen character/design

package client

import (
	character "character/gen/character"
	characterpb "character/gen/grpc/character/pb"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildCreateCharacterFunc builds the remote method to invoke for "Character"
// service "createCharacter" endpoint.
func BuildCreateCharacterFunc(grpccli characterpb.CharacterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.CreateCharacter(ctx, reqpb.(*characterpb.CreateCharacterRequest), opts...)
		}
		return grpccli.CreateCharacter(ctx, &characterpb.CreateCharacterRequest{}, opts...)
	}
}

// EncodeCreateCharacterRequest encodes requests sent to Character
// createCharacter endpoint.
func EncodeCreateCharacterRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*character.Character)
	if !ok {
		return nil, goagrpc.ErrInvalidType("Character", "createCharacter", "*character.Character", v)
	}
	return NewProtoCreateCharacterRequest(payload), nil
}

// DecodeCreateCharacterResponse decodes responses from the Character
// createCharacter endpoint.
func DecodeCreateCharacterResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*characterpb.CreateCharacterResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("Character", "createCharacter", "*characterpb.CreateCharacterResponse", v)
	}
	if err := ValidateCreateCharacterResponse(message); err != nil {
		return nil, err
	}
	res := NewCreateCharacterResult(message)
	return res, nil
}

// BuildGetCharacterFunc builds the remote method to invoke for "Character"
// service "getCharacter" endpoint.
func BuildGetCharacterFunc(grpccli characterpb.CharacterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.GetCharacter(ctx, reqpb.(*characterpb.GetCharacterRequest), opts...)
		}
		return grpccli.GetCharacter(ctx, &characterpb.GetCharacterRequest{}, opts...)
	}
}

// EncodeGetCharacterRequest encodes requests sent to Character getCharacter
// endpoint.
func EncodeGetCharacterRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*character.GetCharacterPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("Character", "getCharacter", "*character.GetCharacterPayload", v)
	}
	return NewProtoGetCharacterRequest(payload), nil
}

// DecodeGetCharacterResponse decodes responses from the Character getCharacter
// endpoint.
func DecodeGetCharacterResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*characterpb.GetCharacterResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("Character", "getCharacter", "*characterpb.GetCharacterResponse", v)
	}
	if err := ValidateGetCharacterResponse(message); err != nil {
		return nil, err
	}
	res := NewGetCharacterResult(message)
	return res, nil
}

// BuildListCharactersFunc builds the remote method to invoke for "Character"
// service "listCharacters" endpoint.
func BuildListCharactersFunc(grpccli characterpb.CharacterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.ListCharacters(ctx, reqpb.(*characterpb.ListCharactersRequest), opts...)
		}
		return grpccli.ListCharacters(ctx, &characterpb.ListCharactersRequest{}, opts...)
	}
}

// DecodeListCharactersResponse decodes responses from the Character
// listCharacters endpoint.
func DecodeListCharactersResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*characterpb.ListCharactersResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("Character", "listCharacters", "*characterpb.ListCharactersResponse", v)
	}
	if err := ValidateListCharactersResponse(message); err != nil {
		return nil, err
	}
	res := NewListCharactersResult(message)
	return res, nil
}

// BuildUpdateCharacterFunc builds the remote method to invoke for "Character"
// service "updateCharacter" endpoint.
func BuildUpdateCharacterFunc(grpccli characterpb.CharacterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.UpdateCharacter(ctx, reqpb.(*characterpb.UpdateCharacterRequest), opts...)
		}
		return grpccli.UpdateCharacter(ctx, &characterpb.UpdateCharacterRequest{}, opts...)
	}
}

// EncodeUpdateCharacterRequest encodes requests sent to Character
// updateCharacter endpoint.
func EncodeUpdateCharacterRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*character.Character)
	if !ok {
		return nil, goagrpc.ErrInvalidType("Character", "updateCharacter", "*character.Character", v)
	}
	return NewProtoUpdateCharacterRequest(payload), nil
}

// DecodeUpdateCharacterResponse decodes responses from the Character
// updateCharacter endpoint.
func DecodeUpdateCharacterResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*characterpb.UpdateCharacterResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("Character", "updateCharacter", "*characterpb.UpdateCharacterResponse", v)
	}
	if err := ValidateUpdateCharacterResponse(message); err != nil {
		return nil, err
	}
	res := NewUpdateCharacterResult(message)
	return res, nil
}

// BuildDeleteCharacterFunc builds the remote method to invoke for "Character"
// service "deleteCharacter" endpoint.
func BuildDeleteCharacterFunc(grpccli characterpb.CharacterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.DeleteCharacter(ctx, reqpb.(*characterpb.DeleteCharacterRequest), opts...)
		}
		return grpccli.DeleteCharacter(ctx, &characterpb.DeleteCharacterRequest{}, opts...)
	}
}

// EncodeDeleteCharacterRequest encodes requests sent to Character
// deleteCharacter endpoint.
func EncodeDeleteCharacterRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*character.DeleteCharacterPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("Character", "deleteCharacter", "*character.DeleteCharacterPayload", v)
	}
	return NewProtoDeleteCharacterRequest(payload), nil
}
