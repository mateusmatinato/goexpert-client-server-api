package exchangerate

import (
	"fmt"
	"github.com/mateusmatinato/client-server-api/client/domain"
	"os"
)

type Repository interface {
	Save(resp *domain.ExchangeRateClientResponse) error
}

type repositoryHandler struct {
	filename string
}

func (f repositoryHandler) Save(resp *domain.ExchangeRateClientResponse) error {
	file, err := os.OpenFile(f.filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(fmt.Sprintf("Dólar: %s", resp.Bid)))
	if err != nil {
		fmt.Printf("error writing file: %s", err.Error())
		return err
	}

	return nil
}

func NewRepository(filename string) Repository {
	return &repositoryHandler{filename: filename}
}
