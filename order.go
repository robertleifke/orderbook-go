package main

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type Side string

const (
	Buy  Side = "buy"
	Sell Side = "sell"
)

type Order struct {
	side 		Side
	id   		string
	timestamp 	time.Time
	price 		decimal.Decimal
	quantity 	decimal.Decimal
}

func main() {
	// Create a new order
	order := Order{
		side:      Buy,
		id:        "12345",
		timestamp: time.Now(),
		quantity:  decimal.NewFromFloat(10.5),
		price:     decimal.NewFromFloat(123.45),
	}

	// Print order details
	fmt.Printf("Order ID: %s\n", order.id)
	fmt.Printf("Side: %s\n", order.side)
	fmt.Printf("Timestamp: %s\n", order.timestamp)
	fmt.Printf("Quantity: %s\n", order.quantity)
	fmt.Printf("Price: %s\n", order.price)
}