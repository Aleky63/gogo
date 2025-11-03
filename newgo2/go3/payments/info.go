package payments

type PaymentInfo struct {
	ID          int
	Description string
	Usd         int
	Cancelled   bool
}
