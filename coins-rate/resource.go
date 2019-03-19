package coins_rate

import "context"

// Coins rate resource interface
type Resource interface {
	BitCoinToUSDRate(ctx context.Context) (float64, error)
}
