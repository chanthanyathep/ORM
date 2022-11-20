package repository

type Buyer_order struct {
	Order_id     int    `gorm:"column:order_id;type:int(11)"`
	Buyer_name   string `gorm:"column:buyer_name;type:varchar(100)"`
	Order_status int    `gorm:"column:order_status;type:varchar(100)"`
	Order_date   string `gorm:"column:order_date;type:date"`
	Is_active    string `gorm:"column:is_active;type:varchar(1)"`
}

type BuyerRepoitory interface {
	GetBuyers() ([]Buyer_order, error)
	GetBuyerById(id int) (b Buyer_order, err error)
	CreateBuyer(buyer Buyer_order) (b Buyer_order, err error)
	UpdateBuyer(id int, name string, status int, date string, active string) (b Buyer_order, err error)
	DeleteBuyer(id int) error
}

func (Buyer_order) TableName() string {
	return "buyer_order"
}

// func mockData(db *gorm.DB) error {
// 	seed := rand.NewSource(time.Now().UnixNano())
// 	random := rand.New(seed)

// 	buyer := []Buyer_order{}
// 	for i := 0; i < 5000; i++ {
// 		buyer = append(buyer, Buyer_order{
// 			Buyer_name:   fmt.Sprintf("Buyer%v", i+1),
// 			Order_status: random.Intn(50),
// 			Order_date:   "2022-11-16",
// 			Is_active:    "N",
// 		})
// 	}
// 	return db.Create(&buyer).Error
// }
