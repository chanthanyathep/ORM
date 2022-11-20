package service

type Buyer_order struct {
	Order_id     int    `json:"order_id"`
	Buyer_name   string `json:"buyer_name"`
	Order_status int    `json:"order_status"`
	Order_date   string `json:"order_date"`
	Is_active    string `json:"is_active"`
}

type BuyerService interface {
	GetBuyers() ([]Buyer_order, error)
	GetBuyerById(id int) (b *Buyer_order, err error)
	CreateBuyer(buyer Buyer_order) (b Buyer_order, err error)
	UpdateBuyer(buyer Buyer_order) (b Buyer_order, err error)
	DeleteBuyer(id int) error
}
