package domain

import (
	"github.com/google/uuid"
	"time"
)

type ExchangeRateDAO struct {
	ID        string    `json:"id"`
	Bid       string    `json:"bid"`
	CreatedAt time.Time `json:"created_at"`
}

func NewExchangeRateDAO(bid string) ExchangeRateDAO {
	id := uuid.New().String()
	return ExchangeRateDAO{
		ID:        id,
		Bid:       bid,
		CreatedAt: time.Now(),
	}
}
