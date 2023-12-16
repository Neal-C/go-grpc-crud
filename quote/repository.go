package quote

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type QuoteRepository struct {
	postgreSQLPool *pgxpool.Pool
}

func NewQuoteRepository(connectionPool *pgxpool.Pool) *QuoteRepository {
	return &QuoteRepository{
		postgreSQLPool: connectionPool,
	}
}

func (self *QuoteRepository) Insert(ctx context.Context, quote Quote) error {

	_, err := self.postgreSQLPool.Exec(ctx, "INSERT INTO quote VALUES ($1, $2, $3, $4, $5)", quote.ID, quote.Book, quote.Quote, quote.CreatedAt, quote.UpdatedAt)

	if err != nil {
		return fmt.Errorf("Could not insert quote: %w", err)
	}

	return nil;
}
