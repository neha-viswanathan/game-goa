// Code generated by goa v3.11.2, DO NOT EDIT.
//
// inventory gRPC client CLI support package
//
// Command:
// $ goa gen inventory/design

package cli

import (
	"flag"
	"fmt"
	inventoryc "inventory/gen/grpc/inventory/client"
	"os"

	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `inventory (add-item|remove-item|get-inventory)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` inventory add-item --message '{
      "character": "Laborum sed dolores corporis.",
      "count": 3513363385,
      "item": "Eum cum officia aut velit dolorum explicabo."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		inventoryFlags = flag.NewFlagSet("inventory", flag.ContinueOnError)

		inventoryAddItemFlags       = flag.NewFlagSet("add-item", flag.ExitOnError)
		inventoryAddItemMessageFlag = inventoryAddItemFlags.String("message", "", "")

		inventoryRemoveItemFlags       = flag.NewFlagSet("remove-item", flag.ExitOnError)
		inventoryRemoveItemMessageFlag = inventoryRemoveItemFlags.String("message", "", "")

		inventoryGetInventoryFlags       = flag.NewFlagSet("get-inventory", flag.ExitOnError)
		inventoryGetInventoryMessageFlag = inventoryGetInventoryFlags.String("message", "", "")
	)
	inventoryFlags.Usage = inventoryUsage
	inventoryAddItemFlags.Usage = inventoryAddItemUsage
	inventoryRemoveItemFlags.Usage = inventoryRemoveItemUsage
	inventoryGetInventoryFlags.Usage = inventoryGetInventoryUsage

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
		case "inventory":
			svcf = inventoryFlags
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
		case "inventory":
			switch epn {
			case "add-item":
				epf = inventoryAddItemFlags

			case "remove-item":
				epf = inventoryRemoveItemFlags

			case "get-inventory":
				epf = inventoryGetInventoryFlags

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
		case "inventory":
			c := inventoryc.NewClient(cc, opts...)
			switch epn {
			case "add-item":
				endpoint = c.AddItem()
				data, err = inventoryc.BuildAddItemPayload(*inventoryAddItemMessageFlag)
			case "remove-item":
				endpoint = c.RemoveItem()
				data, err = inventoryc.BuildRemoveItemPayload(*inventoryRemoveItemMessageFlag)
			case "get-inventory":
				endpoint = c.GetInventory()
				data, err = inventoryc.BuildGetInventoryPayload(*inventoryGetInventoryMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// inventoryUsage displays the usage of the inventory command and its
// subcommands.
func inventoryUsage() {
	fmt.Fprintf(os.Stderr, `The inventory service performs CRUD operations on a character's inventory
Usage:
    %[1]s [globalflags] inventory COMMAND [flags]

COMMAND:
    add-item: Add an item to a character's inventory
    remove-item: Remove an item from a character's inventory
    get-inventory: Get a character's inventory

Additional help:
    %[1]s inventory COMMAND --help
`, os.Args[0])
}
func inventoryAddItemUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory add-item -message JSON

Add an item to a character's inventory
    -message JSON: 

Example:
    %[1]s inventory add-item --message '{
      "character": "Laborum sed dolores corporis.",
      "count": 3513363385,
      "item": "Eum cum officia aut velit dolorum explicabo."
   }'
`, os.Args[0])
}

func inventoryRemoveItemUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory remove-item -message JSON

Remove an item from a character's inventory
    -message JSON: 

Example:
    %[1]s inventory remove-item --message '{
      "character": "Recusandae quis.",
      "count": 155887092,
      "item": "Alias omnis esse exercitationem molestiae voluptatem rerum."
   }'
`, os.Args[0])
}

func inventoryGetInventoryUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory get-inventory -message JSON

Get a character's inventory
    -message JSON: 

Example:
    %[1]s inventory get-inventory --message '{
      "character": "Consequuntur et porro sit."
   }'
`, os.Args[0])
}
