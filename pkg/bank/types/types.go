package types

// Money
type Money int64

// Category (auto, chemist, shops, restaurants, etc...)
type Category string

// Currency
type Currency string

// Currency code
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN
type PAN string

// Card type
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

type Payment struct {
	ID       int
	Amount   Money
	Category Category
}

type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
