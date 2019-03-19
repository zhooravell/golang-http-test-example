package main

import (
	"context"
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

	handler(w, r)

	resp := w.Result()
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

	handler(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusNotFound {
		t.Fail()
	}

	if string(body) != "There is not result\n" {
		t.Fail()
	}
}
