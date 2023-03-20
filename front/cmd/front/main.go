package main

import (
	characterpb "character/gen/grpc/character/pb"
	"context"
	"flag"
	"fmt"
	frontapi "front"
	front "front/gen/front"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	inventorypb "inventory/gen/grpc/inventory/pb"
	itempb "item/gen/grpc/item/pb"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF                    = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF                  = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF                = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF                  = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF                     = flag.Bool("debug", false, "Log request and response bodies")
		characterServiceEndpoint = flag.String("character-endpoint", "localhost:8080", "gRPC endpoint for the character service")
		itemServiceEndpoint      = flag.String("item-endpoint", "localhost:8081", "gRPC endpoint for the item service")
		inventoryServiceEndpoint = flag.String("inventory-endpoint", "localhost:8082", "gRPC endpoint for the inventory service")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[frontapi] ", log.Ltime)
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

	// Establish a grpc connection to the inventory service
	inventoryConn, err := grpc.Dial(*inventoryServiceEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("unable to establish a connection to the inventory service:", err)
	}
	defer inventoryConn.Close()

	// Initialize the services.
	frontSvc, err := frontapi.NewFront(logger, characterpb.NewCharacterClient(characterConn), itempb.NewItemClient(itemConn), inventorypb.NewInventoryClient(inventoryConn))
	if err != nil {
		logger.Fatal("unable to instantiate the front service:", err)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		frontEndpoints *front.Endpoints
	)
	{
		frontEndpoints = front.NewEndpoints(frontSvc)
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
			addr := "http://localhost:8000"
			u, err := url.Parse(addr)
			if err != nil {
				logger.Fatalf("invalid URL %#v: %s\n", addr, err)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					logger.Fatalf("invalid URL %#v: %s\n", u.Host, err)
				}
				u.Host = net.JoinHostPort(h, *httpPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, "80")
			}
			handleHTTPServer(ctx, u, frontEndpoints, &wg, errc, logger, *dbgF)
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
