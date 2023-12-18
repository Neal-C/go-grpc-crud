package quote

import (
	"context"
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

func (self *Server) Create(ctx context.Context, request protocodegen.QuoteRequest) (*protocodegen.Quote, error) {
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
