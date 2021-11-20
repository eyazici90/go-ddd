package domain

type Status int

const (
	Unknown Status = iota
	Submitted
	Paid
	Shipped
	Canceled
)
