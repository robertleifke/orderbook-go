package orderbookgo
import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	ID        string `json:"id"`
	Quantity int    `json:"quantity"`
	Price     int    `json:"price"`
}
