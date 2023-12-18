//lint:file-ignore ST1006 
package quote

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("stuff not found")

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

func (self *QuoteRepository) Update(ctx context.Context, quote Quote) error {
	res, err := self.postgreSQLPool.Exec(ctx, "UPDATE quote SET (book, quote) = ($2 , $3) WHERE id = $1 ", quote.ID, quote.Book, quote.Quote)

	if err != nil {
		log.Println("Failed to update the quote: %w", err)
		return fmt.Errorf("Failed to update the quote")
	}

	if res.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}

func (self *QuoteRepository) DeleteByID(ctx context.Context, id string) error {
	res, err := self.postgreSQLPool.Exec(ctx, "DELETE FROM quote WHERE id = $1", id)

	if err != nil {
		log.Println("Failed to delete the quote: %w", err)
		return fmt.Errorf("Failed to update the quote")
	}

	if res.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
