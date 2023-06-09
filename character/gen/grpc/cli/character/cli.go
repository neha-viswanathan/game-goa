// Code generated by goa v3.11.2, DO NOT EDIT.
//
// character gRPC client CLI support package
//
// Command:
// $ goa gen character/design

package cli

import (
	characterc "character/gen/grpc/character/client"
	"flag"
	"fmt"
	"os"

	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `character (create-character|get-character|list-characters|update-character|delete-character)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` character create-character --message '{
      "description": "Savior of Princess Zelda",
      "experience": 518303,
      "health": 8319921,
      "name": "Link"
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		characterFlags = flag.NewFlagSet("character", flag.ContinueOnError)

		characterCreateCharacterFlags       = flag.NewFlagSet("create-character", flag.ExitOnError)
		characterCreateCharacterMessageFlag = characterCreateCharacterFlags.String("message", "", "")

		characterGetCharacterFlags       = flag.NewFlagSet("get-character", flag.ExitOnError)
		characterGetCharacterMessageFlag = characterGetCharacterFlags.String("message", "", "")

		characterListCharactersFlags = flag.NewFlagSet("list-characters", flag.ExitOnError)

		characterUpdateCharacterFlags       = flag.NewFlagSet("update-character", flag.ExitOnError)
		characterUpdateCharacterMessageFlag = characterUpdateCharacterFlags.String("message", "", "")

		characterDeleteCharacterFlags       = flag.NewFlagSet("delete-character", flag.ExitOnError)
		characterDeleteCharacterMessageFlag = characterDeleteCharacterFlags.String("message", "", "")
	)
	characterFlags.Usage = characterUsage
	characterCreateCharacterFlags.Usage = characterCreateCharacterUsage
	characterGetCharacterFlags.Usage = characterGetCharacterUsage
	characterListCharactersFlags.Usage = characterListCharactersUsage
	characterUpdateCharacterFlags.Usage = characterUpdateCharacterUsage
	characterDeleteCharacterFlags.Usage = characterDeleteCharacterUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "character":
			svcf = characterFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "character":
			switch epn {
			case "create-character":
				epf = characterCreateCharacterFlags

			case "get-character":
				epf = characterGetCharacterFlags

			case "list-characters":
				epf = characterListCharactersFlags

			case "update-character":
				epf = characterUpdateCharacterFlags

			case "delete-character":
				epf = characterDeleteCharacterFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "character":
			c := characterc.NewClient(cc, opts...)
			switch epn {
			case "create-character":
				endpoint = c.CreateCharacter()
				data, err = characterc.BuildCreateCharacterPayload(*characterCreateCharacterMessageFlag)
			case "get-character":
				endpoint = c.GetCharacter()
				data, err = characterc.BuildGetCharacterPayload(*characterGetCharacterMessageFlag)
			case "list-characters":
				endpoint = c.ListCharacters()
				data = nil
			case "update-character":
				endpoint = c.UpdateCharacter()
				data, err = characterc.BuildUpdateCharacterPayload(*characterUpdateCharacterMessageFlag)
			case "delete-character":
				endpoint = c.DeleteCharacter()
				data, err = characterc.BuildDeleteCharacterPayload(*characterDeleteCharacterMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// characterUsage displays the usage of the character command and its
// subcommands.
func characterUsage() {
	fmt.Fprintf(os.Stderr, `The character service performs CRUD operations on a character
Usage:
    %[1]s [globalflags] character COMMAND [flags]

COMMAND:
    create-character: Create a new character
    get-character: Get a character by name
    list-characters: List all characters
    update-character: Update a given character
    delete-character: Delete a given character

Additional help:
    %[1]s character COMMAND --help
`, os.Args[0])
}
func characterCreateCharacterUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character create-character -message JSON

Create a new character
    -message JSON: 

Example:
    %[1]s character create-character --message '{
      "description": "Savior of Princess Zelda",
      "experience": 518303,
      "health": 8319921,
      "name": "Link"
   }'
`, os.Args[0])
}

func characterGetCharacterUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character get-character -message JSON

Get a character by name
    -message JSON: 

Example:
    %[1]s character get-character --message '{
      "name": "Eius earum quis et."
   }'
`, os.Args[0])
}

func characterListCharactersUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character list-characters

List all characters

Example:
    %[1]s character list-characters
`, os.Args[0])
}

func characterUpdateCharacterUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character update-character -message JSON

Update a given character
    -message JSON: 

Example:
    %[1]s character update-character --message '{
      "description": "Savior of Princess Zelda",
      "experience": 4698556,
      "health": 6773187,
      "name": "Link"
   }'
`, os.Args[0])
}

func characterDeleteCharacterUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character delete-character -message JSON

Delete a given character
    -message JSON: 

Example:
    %[1]s character delete-character --message '{
      "name": "Ut quaerat velit fuga voluptas fuga ea."
   }'
`, os.Args[0])
}
