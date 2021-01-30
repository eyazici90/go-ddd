package order

type Status int

const (
	Submitted Status = iota
	Paid
	Shipped
	Cancelled
)

func ToStatus(i int) Status { return Status(i) }

func FromStatus(s Status) int { return int(s) }
