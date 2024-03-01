package entities

import "time"

type Employee struct {
	ID    string
	Name  string
	Email string
	Admin bool
}

type Order struct {
	OrderID    int
	UserID     string
	OrderDate  time.Time
	Status     string
	OrderItems []OrderItem
}
type OrderItem struct {
	OrderItemID  int
	OrderID      int
	ItemID       string
	Quantity     int
	PricePerUnit float64
}

type Invoice struct {
	OrderID       string
	UserID        string
	TotalAmount   float64
	InvoiceDate   time.Time
	Status        string
	PaymentMethod string
	PayeeName     string
	PayeeAddress  string
	PayeeEmail    string
}
