package coins_rate

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCryptoCompareResource_BitCoinToUSDRate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"USD":3992.54}`))
		return
	}))

	resource := cryptoCompareResource{
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

func TestCryptoCompareResource_BitCoinToUSDRateNotOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}))

	resource := cryptoCompareResource{
		httpClient: server.Client(),
		baseUrl:    server.URL,
	}

	result, err := resource.BitCoinToUSDRate(nil)

	if err == nil {
		t.Fail()
	}

	if err.Error() != "[CryptoCompare] not OK status code" {
		t.Fail()
	}

	if result != 0 {
		t.Fail()
	}
}

func TestCryptoCompareResource_BitCoinToUSDRateCheckMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fail()
		}
	}))

	resource := cryptoCompareResource{
		httpClient: server.Client(),
		baseUrl:    server.URL,
	}

	resource.BitCoinToUSDRate(nil)
}

func TestCryptoCompareResource_BitCoinToUSDRateCheckURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if fmt.Sprintf("/data/price?fsym=BTC&tsyms=USD") != r.URL.String() {
			t.Fail()
		}
	}))

	resource := cryptoCompareResource{
		httpClient: server.Client(),
		baseUrl:    server.URL,
	}

	resource.BitCoinToUSDRate(nil)
}

func TestCryptoCompareResource_BitCoinToUSDRateInvalidBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":4010.87`))
		return
	}))

	resource := cryptoCompareResource{
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

func TestCryptoCompareResource_BitCoinToUSDRateTimout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusNoContent)
		return
	}))

	resource := cryptoCompareResource{
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
