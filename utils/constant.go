package utils

type UserRole int

type OrderStatus int

const (
	Buyer UserRole = iota
	Seller
)

const (
	Pending OrderStatus = iota
	Accepted
)

func (r UserRole) String() string {
	switch r {
	case Buyer:
		return "buyer"
	case Seller:
		return "seller"
	}
	return "N/A"
}

func (o OrderStatus) String() string {
	switch o {
	case Pending:
		return "Pending"
	case Accepted:
		return "Accepted"
	}
	return "N/A"
}
