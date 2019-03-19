package coins_rate

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCoinCapResource_BitCoinToUSDRate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":{"id":"bitcoin","symbol":"BTC","currencySymbol":"₿","type":"crypto","rateUsd":"4010.8714336221081818"},"timestamp":1552990697033}`))
		return
	}))

	resource := coinCapResource{
		httpClient: server.Client(),
		baseUrl:    server.URL,
	}

	result, err := resource.BitCoinToUSDRate(nil)

	if err != nil {
		t.Error(err)
	}

	if result == 0 {
		t.Fail()
	}
}

func TestCoinCapResource_BitCoinToUSDRateNotOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}))

	resource := coinCapResource{
		httpClient: server.Client(),
		baseUrl:    server.URL,
	}

	result, err := resource.BitCoinToUSDRate(nil)

	if err == nil {
		t.Fail()
	}

	if err.Error() != "[CoinCap] not OK status code" {
		t.Fail()
	}

	if result != 0 {
		t.Fail()
	}
}

func TestCoinCapResource_BitCoinToUSDRateCheckMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fail()
		}
	}))

	resource := coinCapResource{
		httpClient: server.Client(),
		baseUrl:    server.URL,
	}

	resource.BitCoinToUSDRate(nil)
}

func TestCoinCapResource_BitCoinToUSDRateCheckURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if fmt.Sprintf("/v2/rates/bitcoin") != r.URL.String() {
			t.Fail()
		}
	}))

	resource := coinCapResource{
		httpClient: server.Client(),
		baseUrl:    server.URL,
	}

	resource.BitCoinToUSDRate(nil)
}

func TestCoinCapResource_BitCoinToUSDRateInvalidBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":4010.87`))
		return
	}))

	resource := coinCapResource{
		httpClient: server.Client(),
		baseUrl:    server.URL,
	}

	result, err := resource.BitCoinToUSDRate(nil)

	if err == nil {
		t.Fail()
	}

	if result != 0 {
		t.Fail()
	}
}

func TestCoinCapResource_BitCoinToUSDRateTimout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusNoContent)
		return
	}))

	resource := coinCapResource{
		httpClient: server.Client(),
		baseUrl:    server.URL,
	}

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 10*time.Millisecond)

	result, err := resource.BitCoinToUSDRate(ctx)

	if err == nil {
		t.Fail()
	}

	if result != 0 {
		t.Fail()
	}
}
