//lint:file-ignore ST1006
package quote

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Neal-C/go-grpc-crud/protocodegen"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	quoteRepository *QuoteRepository
	protocodegen.UnimplementedQuoteApiServer
}

func NewServer(quoteRepository *QuoteRepository) *Server {
	return &Server{
		quoteRepository: quoteRepository,
	}
}

func (self *Server) Create(ctx context.Context, request *protocodegen.QuoteRequest) (*protocodegen.Quote, error) {
	if request.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot be empty")
	}

	if request.Book == ""{
		return nil, status.Errorf(codes.InvalidArgument,"book cannot be empty" )
	}
	if request.Quote == ""{
		return nil, status.Errorf(codes.InvalidArgument,"quote cannot be empty" )
	}

	quoteID, err := uuid.Parse(request.Id);

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id isn't valid UUID")
	}

	now := time.Now()

	quote := Quote{
		ID: quoteID,
		Book: request.Book,
		Quote: request.Quote,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := self.quoteRepository.Insert(ctx, quote); err != nil {
		log.Println("Failed to insert quote", "error : %w", err)
		return nil, fmt.Errorf("failed to insert quote")
	}

	grpcQuote := QuoteToGRPCQuote(quote) 

	return &grpcQuote, nil
}

func (self *Server) Read(ctx context.Context, request *protocodegen.QuoteRequest)(*protocodegen.QuoteList, error){
	quotes, err := self.quoteRepository.FindAll(ctx)

	if err != nil {
		log.Println("Failed to select all from quote", "error : %w", err)
		return nil, fmt.Errorf("couldn't find all quotes")
	}

	response := make([]*protocodegen.Quote, len(quotes))

	for index, item := range quotes {
		quote := QuoteToGRPCQuote(item)

		response[index] = &quote
	}

	return &protocodegen.QuoteList{
		Quotes: response,
	}, nil
}

func (self *Server) ReadOne(ctx context.Context, request *protocodegen.QuoteID) (*protocodegen.Quote, error){

	quote, err := self.quoteRepository.FindByID(ctx, request.Id)

	if err != nil {
		log.Println("Failed to find quote by id ", "error: %w", err)
		return nil, status.Errorf(codes.NotFound, "failed to find quote by id")
	}

	response := QuoteToGRPCQuote(quote)

	return &response, nil
}

func (self *Server) Update(ctx context.Context, request *protocodegen.QuoteRequest)(*protocodegen.Quote, error){
	quote, err := self.quoteRepository.FindByID(ctx, request.Id)

	if err != nil {
		log.Println("Failed to find quote by id ", "error: %w", err)
		return nil, status.Errorf(codes.NotFound, "failed to find quote by id")
	}

	quote.UpdatedAt = time.Now()
	quote.Book = request.Book
	quote.Quote = request.Quote

	err = self.quoteRepository.Update(ctx, quote)

	if errors.Is(err, ErrNotFound){
		log.Println("Failed to update quote", "error: %w", err)
		return nil, status.Errorf(codes.NotFound, "failed to update quote")
	}

	response := QuoteToGRPCQuote(quote)

	return &response, nil
}