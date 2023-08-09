package exchangerate

import (
	"context"
	"database/sql"
	"github.com/mateusmatinato/client-server-api/server/domain"
	"log"
	"time"
)

type Repository interface {
	Save(ctx context.Context, exchangeRate domain.ExchangeRateDAO) error
}

type repositoryHandler struct {
	*sql.DB
}

func (e *repositoryHandler) Save(ctx context.Context, exchangeRate domain.ExchangeRateDAO) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	stmt, err := e.DB.PrepareContext(ctx, "INSERT INTO exchange_rate(id, bid, created_at) VALUES (?, ?, ?)")
	if err != nil {
		log.Printf("error preparing statement: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, exchangeRate.ID, exchangeRate.Bid, exchangeRate.CreatedAt)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("timeout on save exchange rate")
			return ctx.Err()
		}

		log.Printf("error executing statement: %s\n", err.Error())
		return err
	}

	log.Printf("exchange rate saved: %+v\n", exchangeRate)
	return nil
}

func NewRepository(db *sql.DB) Repository {
	return &repositoryHandler{DB: db}
}
