package orderbook

import (
	"encoding/json"
	"reflect"
)
	
// A Side of the order.
type Side int

const (
	// Sell for asks
	Sell Side = 0

	// Buy for bids
	Buy Side = 1
)

func (s Side) String() string {
	return []string{"sell", "buy"}[s]
}