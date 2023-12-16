package quote

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
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
		log.Println("Failed to insert quote", "error", err)
		return fmt.Errorf("Could not insert quote")
	}

	return nil
}

func (self *QuoteRepository) FindAll(ctx context.Context) ([]Quote, error) {
	rows, err := self.postgreSQLPool.Query(ctx, "SELECT id, book, quote, created_at, updated_at FROM quote")

	if err != nil {
		return nil, fmt.Errorf("failed to find all quotes: %w", err)
	}

	defer rows.Close()

	quotes, err := pgx.CollectRows[Quote](rows, pgx.RowToStructByName[Quote])

	if err != nil {
		log.Println("Failed to map RowToStructByName", "error", err)
		return nil, fmt.Errorf("Failed to scan rows")
	}

	return quotes, nil
}

func (self *QuoteRepository) FindByID(ctx context.Context, id string) (Quote, error) {
	var response Quote

	row := self.postgreSQLPool.QueryRow(ctx, "SELECT id, book, quote, created_at, updated_at FROM quote WHERE id = $1", id)

	err := row.Scan(&response.ID, &response.Book, &response.Quote, &response.CreatedAt, &response.UpdatedAt)

	if err != nil {
		log.Println("Failed to scan quote", "error", err)
		return Quote{}, fmt.Errorf("Could not scan quote")
	}

	return response, nil

}
