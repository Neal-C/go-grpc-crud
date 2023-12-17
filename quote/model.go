package quote

import (
	"time"

	"github.com/Neal-C/go-grpc-crud/protocodegen"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Quote struct {
	ID         uuid.UUID `json:"id,omitempty" db:"id"`
	Book       string    `json:"book,omitempty" db:"book"`
	Quote      string    `json:"quote,omitempty" db:"quote"`
	CreatedAt  time.Time `json:"insertedAt,omitempty" db:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty" db:"updated_at"`
}

func QuoteToGRPCQuote(quote Quote) protocodegen.Quote {
	return protocodegen.Quote{
		// investigate this
		Id: quote.ID.URN(),
		Book: quote.Book,
		Quote: quote.Quote,
		CreatedAt:timestamppb.New(quote.CreatedAt),
		UpdatedAt: timestamppb.New(quote.UpdatedAt),
	}
}