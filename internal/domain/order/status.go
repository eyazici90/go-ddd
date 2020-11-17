package order

type Status int

const (
	Submitted Status = 1
	Paid      Status = 2
	Shipped   Status = 3
	Cancelled Status = 4
)

func ToStatus(i int) Status { return Status(i) }

func FromStatus(status Status) int { return int(status) }
