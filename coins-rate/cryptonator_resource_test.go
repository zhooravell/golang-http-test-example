package coins_rate

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestCryptonatorResource_BitCoinToUSDRate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ticker":{"base":"BTC","target":"USD","price":"4007.93579679","volume":"32581.54482720","change":"2.52534325"},"timestamp":1552993382,"success":true,"error":""}`))
		return
	}))

	resource := cryptonatorResource{
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

func TestCryptonatorResource_BitCoinToUSDRateNotOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}))

	resource := cryptonatorResource{
		httpClient: server.Client(),
		baseUrl:    server.URL,
	}

	result, err := resource.BitCoinToUSDRate(nil)

	if err == nil {
		t.Fail()
	}

	if err.Error() != "[Cryptonator] not OK status code" {
		t.Fail()
	}

	if result != 0 {
		t.Fail()
	}
}

func TestCryptonatorResource_BitCoinToUSDRateCheckMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fail()
		}
	}))

	resource := cryptonatorResource{
		httpClient: server.Client(),
		baseUrl:    server.URL,
	}

	resource.BitCoinToUSDRate(nil)
}

func TestCryptonatorResource_BitCoinToUSDRateCheckURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if fmt.Sprintf("/api/ticker/BTC-USD") != r.URL.String() {
			t.Fail()
		}
	}))

	resource := cryptonatorResource{
		httpClient: server.Client(),
		baseUrl:    server.URL,
	}

	resource.BitCoinToUSDRate(nil)
}

func TestCryptonatorResource_BitCoinToUSDRateInvalidBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":4010.87`))
		return
	}))

	resource := cryptonatorResource{
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

func TestCryptonatorResource_BitCoinToUSDRateTimout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusNoContent)
		return
	}))

	resource := cryptonatorResource{
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

func TestCryptonatorResource_BitCoinToUSDRateInvalidBaseURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	resource := cryptonatorResource{
		httpClient: server.Client(),
		baseUrl:    "%%2",
	}

	result, err := resource.BitCoinToUSDRate(nil)

	if err == nil {
		t.Error(err)
	}

	if result != 0 {
		t.Fail()
	}
}

func TestNewCryptonatorResource(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	r := NewCryptonatorResource(server.Client())

	if reflect.TypeOf(r).String() != "*coins_rate.cryptonatorResource" {
		t.Fail()
	}
}
