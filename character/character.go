package character

import (
	"character/gen/character"
	"context"
	"log"
)

// character service example implementation.
// The example methods log the requests and return zero values.
type charactersrvc struct {
	logger     *log.Logger
	characters map[string]*character.Character
}

// NewCharacter returns the character service implementation.
func NewCharacter(logger *log.Logger) character.Service {
	return &charactersrvc{
		logger:     logger,
		characters: map[string]*character.Character{},
	}
}

// CreateCharacter Create a new character
func (s *charactersrvc) CreateCharacter(ctx context.Context, p *character.Character) (res *character.Character, err error) {
	s.logger.Print("character.create")
	_, found := s.characters[p.Name]
	if found {
		return nil, &character.CharacterAlreadyExists{Message: "character already exists"}
	}
	res = &character.Character{
		Name:        p.Name,
		Description: p.Description,
		Health:      p.Health,
		Experience:  p.Experience,
	}
	s.characters[p.Name] = res
	return res, nil
}

// GetCharacter Get an existing character
func (s *charactersrvc) GetCharacter(ctx context.Context, p *character.GetCharacterPayload) (res *character.Character, err error) {
	s.logger.Print("character.get", p)
	char, found := s.characters[p.Name]
	if !found {
		return nil, &character.CharacterNotFound{Message: "character not found"}
	}
	return char, nil
}

// ListCharacters List all characters
func (s *charactersrvc) ListCharacters(ctx context.Context) (res []*character.Character, err error) {
	s.logger.Print("character.list")
	for _, character := range s.characters {
		res = append(res, character)
	}
	return res, nil
}

// UpdateCharacter Update an existing character
func (s *charactersrvc) UpdateCharacter(ctx context.Context, p *character.Character) (res *character.Character, err error) {
	s.logger.Print("character.update")
	char, found := s.characters[p.Name]
	if !found {
		return nil, &character.CharacterNotFound{Message: "character not found"}
	}
	char.Description = p.Description
	char.Health = p.Health
	char.Experience = p.Experience
	s.characters[char.Name] = char

	return char, nil
}

// DeleteCharacter Delete character
func (s *charactersrvc) DeleteCharacter(ctx context.Context, p *character.DeleteCharacterPayload) (err error) {
	s.logger.Print("character.delete")
	delete(s.characters, p.Name)
	s.logger.Printf("The character %d has been deleted successfully", p.Name)
	return nil
}
