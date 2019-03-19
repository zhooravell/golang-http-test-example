package coins_rate

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

// https://coincap.io/
// https://api.coincap.io/v2/assets
// https://api.coincap.io/v2/rates/bitcoin
// https://docs.coincap.io/#c7925509-73f4-4b11-a602-74d1fee44dba

type coinCapResource struct {
	httpClient *http.Client
	baseUrl    string
}

type coinCapResponse struct {
	Data struct {
		RateUsd float64 `json:"rateUsd,string"` // The "string" option signals that a field is stored as JSON inside a JSON-encoded string
	} `json:"data"`
}

// return current BitCoin to USD rate on https://coincap.io/
func (rcv *coinCapResource) BitCoinToUSDRate(ctx context.Context) (float64, error) {
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/v2/rates/bitcoin", rcv.baseUrl), nil)

	if err != nil {
		return 0, errors.Wrap(err, "[CoinCap]")
	}

	r.Header.Set("Accept", "application/json")

	if ctx != nil {
		r = r.WithContext(ctx)
	}

	res, err := rcv.httpClient.Do(r)

	if err != nil {
		return 0, errors.Wrap(err, "[CoinCap]")
	}

	if res.StatusCode != http.StatusOK {
		return 0, errors.New("[CoinCap] not OK status code")
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	var data coinCapResponse

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return 0, errors.Wrap(err, "[CoinCap]")
	}

	return data.Data.RateUsd, nil
}

// Constructor for CoinCap resource
func NewCoinCapResource(httpClient *http.Client) Resource {
	return &coinCapResource{httpClient: httpClient, baseUrl: "https://api.coincap.io"}
}
