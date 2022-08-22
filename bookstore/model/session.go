package model

type Session struct {
	SessionID string
	UserName  string
	UserID    int
	Cart      *Cart
	OrderID   string
	Orders    []*Order
}
