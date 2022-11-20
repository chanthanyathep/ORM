package repository

import (
	"orm/logs"

	"gorm.io/gorm"
)

type buyerRepo struct {
	db *gorm.DB
}

func NewBuyerRepositoryDB(db *gorm.DB) BuyerRepoitory {
	return buyerRepo{db: db}
}

func (r buyerRepo) GetBuyers() (b []Buyer_order, err error) {
	err = r.db.Limit(10).Find(&b).Error
	return b, err
}

func (r buyerRepo) GetBuyerById(id int) (b Buyer_order, err error) {
	tx := r.db.Find(&b, "order_id", id)
	if tx.Error != nil {
		return b, err
	}
	return b, nil
}

func (r buyerRepo) CreateBuyer(buyer Buyer_order) (b Buyer_order, err error) {
	tx := r.db.Create(&buyer)
	b = buyer
	if tx.Error != nil {
		logs.Error(err)
		return b, tx.Error
	}
	return b, nil
}

func (r buyerRepo) UpdateBuyer(id int, name string, status int, date string, active string) (b Buyer_order, err error) {
	b = Buyer_order{
		Order_id:     id,
		Buyer_name:   name,
		Order_status: status,
		Order_date:   date,
		Is_active:    active,
	}
	tx := r.db.Model(&Buyer_order{}).Where("order_id", id).Updates(b)
	if tx.Error != nil {
		logs.Error(err)
		return b, tx.Error
	}
	return b, nil
}

func (r buyerRepo) DeleteBuyer(id int) error {
	tx := r.db.Delete(&Buyer_order{}, id)
	if tx.Error != nil {
		logs.Error(tx.Error)
		return tx.Error
	}
	return nil
}

//name string, status int, date string, active string
