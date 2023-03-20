package main

import (
	"context"
	"flag"
	"fmt"
	"inventory/gen/inventory"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	characterpb "character/gen/grpc/character/pb"
	inventoryapi "inventory"
	itempb "item/gen/grpc/item/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF                    = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF                  = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		grpcPortF                = flag.String("grpc-port", "", "gRPC port (overrides host gRPC port specified in service design)")
		secureF                  = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF                     = flag.Bool("debug", false, "Log request and response bodies")
		characterServiceEndpoint = flag.String("character-endpoint", "localhost:8080", "gRPC endpoint for the character service")
		itemServiceEndpoint      = flag.String("item-endpoint", "localhost:8081", "gRPC endpoint for the item service")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[inventoryapi] ", log.Ltime)
	}

	// Establish a grpc connection to the character service
	characterConn, err := grpc.Dial(*characterServiceEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("unable to establish a connection to the character service:", err)
	}
	defer characterConn.Close()

	// Establish a grpc connection to the item service
	itemConn, err := grpc.Dial(*itemServiceEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("unable to establish a connection to the item service:", err)
	}
	defer itemConn.Close()

	// Initialize the services
	inventorySvc, err := inventoryapi.NewInventory(logger, characterpb.NewCharacterClient(characterConn), itempb.NewItemClient(itemConn))
	if err != nil {
		logger.Fatal("unable to instantiate the inventory service:", err)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		inventoryEndpoints *inventory.Endpoints
	)
	{
		inventoryEndpoints = inventory.NewEndpoints(inventorySvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "grpc://localhost:8082"
			u, err := url.Parse(addr)
			if err != nil {
				logger.Fatalf("invalid URL %#v: %s\n", addr, err)
			}
			if *secureF {
				u.Scheme = "grpcs"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *grpcPortF != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					logger.Fatalf("invalid URL %#v: %s\n", u.Host, err)
				}
				u.Host = net.JoinHostPort(h, *grpcPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, "8080")
			}
			handleGRPCServer(ctx, u, inventoryEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		logger.Fatalf("invalid host argument: %q (valid hosts: localhost)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
