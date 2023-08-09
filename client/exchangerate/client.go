package exchangerate

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/mateusmatinato/client-server-api/client/domain"
	"log"
	"net/http"
	"time"
)

const url = "http://localhost:8080/cotacao"

type Client interface {
	Get(ctx context.Context) (*domain.ExchangeRateClientResponse, error)
}

type clientHandler struct {
	httpClient *http.Client
}

func (e clientHandler) Get(ctx context.Context) (*domain.ExchangeRateClientResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Printf("error on create request: %s\n", err.Error())
		return nil, err
	}

	resp, err := e.httpClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("timeout on get exchange rate from server")
			return nil, ctx.Err()
		}

		log.Printf("error on get exchange rate from server: %s\n", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("exchange rate server returned invalid status: %d\n", resp.StatusCode)
		return nil, errors.New("error on get exchange rate from server")
	}

	var exchangeRateResponse domain.ExchangeRateClientResponse
	if err = json.NewDecoder(resp.Body).Decode(&exchangeRateResponse); err != nil {
		log.Printf("error on decode exchange rate response from server: %s\n", err.Error())
		return nil, err
	}

	log.Printf("exchange rate response from server: %+v\n", exchangeRateResponse)
	return &exchangeRateResponse, nil
}

func NewClient() Client {
	return &clientHandler{httpClient: &http.Client{}}
}
