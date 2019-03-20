package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zhooravell/golang-http-test-example/coins-rate"
)

type testResource struct {
	result float64
}

func (rcv *testResource) BitCoinToUSDRate(ctx context.Context) (float64, error) {
	return rcv.result, nil
}

func TestGetRates(t *testing.T) {

	rateResources = make([]coins_rate.Resource, 2)
	rateResources[0] = &testResource{result: 10.5}
	rateResources[1] = &testResource{result: 20.5}

	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	getBitcoinRateHandler(w, r)

	resp := w.Result()
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Fail()
	}

	if string(body) != "BitCoin to USD rate: 15.500000 $\n" {
		t.Fail()
	}
}

func TestGetRates_NotFound(t *testing.T) {
	rateResources = make([]coins_rate.Resource, 0)

	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	getBitcoinRateHandler(w, r)

	resp := w.Result()
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusNotFound {
		t.Fail()
	}

	if string(body) != "There is not result\n" {
		t.Fail()
	}
}

func TestRouting_RateBTC(t *testing.T) {
	rateResources = make([]coins_rate.Resource, 2)
	rateResources[0] = &testResource{result: 10.5}
	rateResources[1] = &testResource{result: 20.5}

	srv := httptest.NewServer(handlers())
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/rate/btc", srv.URL))

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "BitCoin to USD rate: 15.500000 $\n" {
		t.Fail()
	}
}
