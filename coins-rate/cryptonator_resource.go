package coins_rate

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

// https://www.cryptonator.com
// https://www.cryptonator.com/api
// https://api.cryptonator.com/api/ticker/USD-BTC

type cryptonatorResource struct {
	httpClient *http.Client
	baseUrl    string
}

type cryptonatorResponse struct {
	Ticker struct {
		Price float64 `json:"price,string"` // The "string" option signals that a field is stored as JSON inside a JSON-encoded string
	} `json:"ticker"`
}

// return current BitCoin to USD rate on https://www.cryptonator.com
func (rcv *cryptonatorResource) BitCoinToUSDRate(ctx context.Context) (float64, error) {
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/api/ticker/BTC-USD", rcv.baseUrl), nil)

	if err != nil {
		return 0, errors.Wrap(err, "[Cryptonator]")
	}

	r.Header.Set("Accept", "application/json")

	if ctx != nil {
		r = r.WithContext(ctx)
	}

	res, err := rcv.httpClient.Do(r)

	if err != nil {
		return 0, errors.Wrap(err, "[Cryptonator]")
	}

	if res.StatusCode != http.StatusOK {
		return 0, errors.New("[Cryptonator] not OK status code")
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	var data cryptonatorResponse

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return 0, errors.Wrap(err, "[Cryptonator]")
	}

	return data.Ticker.Price, nil
}

// Constructor for Cryptonator resource
func NewCryptonatorResource(httpClient *http.Client) Resource {
	return &cryptonatorResource{httpClient: httpClient, baseUrl: "https://api.cryptonator.com"}
}
