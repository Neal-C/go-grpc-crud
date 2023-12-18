//lint:file-ignore ST1006 
package main

import (
	"context"
	"log"
	"net"

	"github.com/Neal-C/go-grpc-crud/database"
	"github.com/Neal-C/go-grpc-crud/protocodegen"
	"github.com/Neal-C/go-grpc-crud/quote"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("failed to create listener:", err)
	}
	defer listener.Close();

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pool, err := database.Connect(context.Background())
	if err != nil {
		log.Fatalln("failed to connect to database:", err)
	}
	defer pool.Close();

	quoteRepository := quote.NewQuoteRepository(pool)
	quoteApiServer := quote.NewServer(quoteRepository)

	protocodegen.RegisterQuoteApiServer(grpcServer,quoteApiServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalln("Failed to serve grpcServer", err)
	}
}
