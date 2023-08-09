package exchangerate

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/mateusmatinato/client-server-api/server/domain"
	"log"
	"net/http"
	"time"
)

const url = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

type Client interface {
	Get(ctx context.Context) (*domain.ExchangeRateClientResponse, error)
}

type clientHandler struct {
	httpClient *http.Client
}

func (e clientHandler) Get(ctx context.Context) (*domain.ExchangeRateClientResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Printf("error on create request: %s\n", err.Error())
		return nil, err
	}

	resp, err := e.httpClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("timeout on get exchange rate")
			return nil, ctx.Err()
		}

		log.Printf("error on get exchange rate: %s\n", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("exchange rate api returned invalid status: %d\n", resp.StatusCode)
		return nil, errors.New("error on get exchange rate")
	}

	var exchangeRateResponse domain.ExchangeRateClientResponse
	if err = json.NewDecoder(resp.Body).Decode(&exchangeRateResponse); err != nil {
		log.Printf("error on decode exchange rate response: %s\n", err.Error())
		return nil, err
	}

	log.Printf("exchange rate response: %+v\n", exchangeRateResponse)
	return &exchangeRateResponse, nil
}

func NewClient() Client {
	return &clientHandler{httpClient: &http.Client{}}
}
