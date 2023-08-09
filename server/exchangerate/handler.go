package exchangerate

import (
	"encoding/json"
	"github.com/mateusmatinato/client-server-api/server/domain"
	"net/http"
)

type Handler struct {
	client     Client
	repository Repository
}

func NewHandler(client Client, repository Repository) *Handler {
	return &Handler{client: client, repository: repository}
}

func (h *Handler) GetExchangeRate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		resp := domain.NewAPIResponse("Method not allowed", http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(resp.ToJSON())
		return
	}

	exchangeRate, err := h.client.Get(r.Context())
	if err != nil {
		resp := domain.NewAPIResponse(err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp.ToJSON())
		return
	}

	err = h.repository.Save(r.Context(), domain.NewExchangeRateDAO(exchangeRate.USDBRL.Bid))
	if err != nil {
		resp := domain.NewAPIResponse(err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp.ToJSON())
		return
	}

	resp := domain.NewExchangeRateResponse(exchangeRate.USDBRL.Bid)
	respJson, _ := json.Marshal(resp)
	w.Write(respJson)
	w.WriteHeader(http.StatusOK)
}
