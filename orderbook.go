package orderbook

import (
	"container/list"
	"fmt"
	"sync"
)

type OrderBook struct {
	mu sync.RWMutex
	asks *list.List
	bids *list.List
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		asks: list.New(),
		bids: list.New(),
	}
}
