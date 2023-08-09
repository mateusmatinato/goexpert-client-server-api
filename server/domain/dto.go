package domain

type ExchangeRateClientResponse struct {
	USDBRL ExchangeRateClientDetails `json:"usdbrl"`
}

type ExchangeRateClientDetails struct {
	Code       string `json:"code"`
	CodeIn     string `json:"code_in"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type ExchangeRateResponse struct {
	Bid string `json:"bid"`
}

func NewExchangeRateResponse(big string) ExchangeRateResponse {
	return ExchangeRateResponse{
		Bid: big,
	}
}
