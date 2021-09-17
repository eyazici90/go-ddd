package order

type Status int

const (
	Unknown Status = iota
	Submitted
	Paid
	Shipped
	Canceled
)
