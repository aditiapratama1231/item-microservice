package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/aditiapratama1231/item-service/pkg/cmd"
	transport "github.com/aditiapratama1231/item-service/pkg/transport"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	httpPort := os.Getenv("ITEM_HTTP_PORT")
	// grpcPort := os.Getenv("ITEM_GRPC_PORT")

	var (
		httpAddr = flag.String("http", ":"+httpPort, "http listen address")
		// grpcAddr = flag.String("grpc", ":"+grpcPort, "gRPC listen address")
	)

	flag.Parse()
	ctx := context.Background()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// Run HTTP Server
	go func() {
		log.Println("Item Service (http) is listening on port", *httpAddr)
		handler := transport.NewHTTPServer(ctx, cmd.Endpoints)
		handler = handlers.LoggingHandler(os.Stdout, handler)
		handler = handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"}),
			handlers.AllowedOrigins([]string{"*"}))(handler)
		if err := http.ListenAndServe(*httpAddr, handler); err != nil {
			errChan <- err
		}
	}()

	log.Fatalln(<-errChan)
}
