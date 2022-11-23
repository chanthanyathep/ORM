package service

import (
	"orm/logs"
	"orm/repository"
)

type buyerServiceDB struct {
	buyerRepo repository.BuyerRepoitory
}

func NewBuyerServ(db repository.BuyerRepoitory) BuyerService {
	return buyerServiceDB{buyerRepo: db}
}

func (s buyerServiceDB) GetBuyers() (buyer []Buyer_order, err error) {

	buyerDB, err := s.buyerRepo.GetBuyers()
	if err != nil {
		return nil, err
	}

	for _, b := range buyerDB {
		buyer = append(buyer, Buyer_order{
			Order_id:     b.Order_id,
			Buyer_name:   b.Buyer_name,
			Order_status: b.Order_status,
			Order_date:   b.Order_date,
			Is_active:    b.Is_active,
		})
	}

	return buyer, nil
}

func (s buyerServiceDB) GetBuyerById(id int) (b *Buyer_order, err error) {
	buyer, err := s.buyerRepo.GetBuyerById(id)
	if err != nil || buyer == nil {
		return nil, err
	}
	b = &Buyer_order{
		Order_id:     buyer.Order_id,
		Buyer_name:   buyer.Buyer_name,
		Order_status: buyer.Order_status,
		Order_date:   buyer.Order_date,
		Is_active:    buyer.Is_active,
	}

	return b, nil
}

func (s buyerServiceDB) CreateBuyer(buyer Buyer_order) (b Buyer_order, err error) {
	buy := repository.Buyer_order{
		Buyer_name:   buyer.Buyer_name,
		Order_status: buyer.Order_status,
		Order_date:   buyer.Order_date,
		Is_active:    buyer.Is_active,
	}
	ss, err := s.buyerRepo.CreateBuyer(buy)
	if err != nil {
		return b, err
	}
	b = Buyer_order{
		Order_id:     ss.Order_id,
		Buyer_name:   ss.Buyer_name,
		Order_status: ss.Order_status,
		Order_date:   ss.Order_date,
		Is_active:    ss.Is_active,
	}
	return b, nil
}

func (s buyerServiceDB) UpdateBuyer(buyer Buyer_order) (b *Buyer_order, err error) {
	b = &buyer
	bb, err := s.buyerRepo.UpdateBuyer(buyer.Order_id, buyer.Buyer_name, buyer.Order_status, buyer.Order_date, buyer.Is_active)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	b = &Buyer_order{
		Order_id:     bb.Order_id,
		Buyer_name:   bb.Buyer_name,
		Order_status: bb.Order_status,
		Order_date:   bb.Order_date,
		Is_active:    bb.Is_active,
	}
	return b, nil
}

func (s buyerServiceDB) DeleteBuyer(id int) error {
	err := s.buyerRepo.DeleteBuyer(id)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
