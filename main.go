package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/zhooravell/golang-http-test-example/coins-rate"
)

const (
	defaultApiAddress = ":8080"
)

var (
	apiAddress    string
	httpClient    *http.Client
	rateResources []coins_rate.Resource
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	apiAddress = getVar("API_ADDRESS", defaultApiAddress)
	httpClient = &http.Client{}

	rateResources = make([]coins_rate.Resource, 3)
	rateResources[0] = coins_rate.NewCoinCapResource(httpClient)
	rateResources[1] = coins_rate.NewCryptoCompareResource(httpClient)
	rateResources[2] = coins_rate.NewCryptonatorResource(httpClient)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("run time panic: %v", r)
		}
	}()

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(apiAddress, nil))
}

// Handler to get BitCoin rate
func handler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	var result []float64

	for _, res := range rateResources {
		wg.Add(1)
		go func(res coins_rate.Resource) {
			defer wg.Done()
			rate, err := res.BitCoinToUSDRate(nil)

			if err != nil {
				log.Println(err)
				return
			}

			result = append(result, rate)
		}(res)
	}

	wg.Wait()

	if len(result) == 0 {
		w.WriteHeader(http.StatusNotFound)
		if _, err := fmt.Fprint(w, "There is not result\n"); err != nil {
			log.Println(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprintf(w, "BitCoin to USD rate: %f $\n", avg(result)); err != nil {
		log.Println(err)
	}
}

// Get environment variable or default value
// See https://golang.org/pkg/os/#Getenv
func getVar(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

// average array numbers
func avg(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	var total float64 = 0
	for _, value := range data {
		total += value
	}

	return total / float64(len(data))
}
