package coins_rate

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

// https://www.cryptocompare.com/
// https://www.cryptocompare.com/api/#-api-data-price-
// https://min-api.cryptocompare.com/data/price?fsym=USD&tsyms=BTC,LTC,DASH,ETH

type cryptoCompareResource struct {
	httpClient *http.Client
	baseUrl    string
}

type cryptoCompareResponse struct {
	USD float64 `json:"USD"`
}

// return current BitCoin to USD rate on https://www.cryptocompare.com/
func (rcv *cryptoCompareResource) BitCoinToUSDRate(ctx context.Context) (float64, error) {
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/data/price?fsym=BTC&tsyms=USD", rcv.baseUrl), nil)

	if err != nil {
		return 0, errors.Wrap(err, "[CryptoCompare]")
	}

	r.Header.Set("Accept", "application/json")

	if ctx != nil {
		r = r.WithContext(ctx)
	}

	res, err := rcv.httpClient.Do(r)

	if err != nil {
		return 0, errors.Wrap(err, "[CryptoCompare]")
	}

	if res.StatusCode != http.StatusOK {
		return 0, errors.New("[CryptoCompare] not OK status code")
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	var data cryptoCompareResponse

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return 0, errors.Wrap(err, "[CryptoCompare]")
	}

	return data.USD, nil
}

// Constructor for CryptoCompare resource
func NewCryptoCompareResource(httpClient *http.Client) Resource {
	return &cryptoCompareResource{httpClient: httpClient, baseUrl: "https://min-api.cryptocompare.com"}
}
