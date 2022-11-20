package handler

import (
	"errors"
	"orm/logs"
	"orm/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BuyerHandler interface {
	GetBuyers(c *fiber.Ctx) error
	GetBuyerById(c *fiber.Ctx) error
	CreateBuyer(c *fiber.Ctx) error
	UpdateBuyer(c *fiber.Ctx) error
	DeleteBuyer(c *fiber.Ctx) error
}

type buyerHandler struct {
	buyerSrv service.BuyerService
}

func NewBuyerHandler(buyerSrv service.BuyerService) BuyerHandler {
	return buyerHandler{buyerSrv: buyerSrv}
}

func (h buyerHandler) GetBuyers(c *fiber.Ctx) error {
	buyer, err := h.buyerSrv.GetBuyers()
	if err != nil {
		panic(err)
	}

	return c.JSON(buyer)
}

func (h buyerHandler) GetBuyerById(c *fiber.Ctx) error {
	id := c.Params("id")
	checkid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	b, err := h.buyerSrv.GetBuyerById(checkid)
	if b == nil {
		return errors.New("record not found")
	}
	if err != nil {
		return err
	}
	return c.JSON(b)
}

func (h buyerHandler) CreateBuyer(c *fiber.Ctx) error {
	if c.Is("application/json") {
		logs.Error("request body is not in json format")
		return errors.New("request body is not in json format")
	}
	b := service.Buyer_order{}
	err := c.BodyParser(&b)

	if err != nil {
		return err
	}
	response, err := h.buyerSrv.CreateBuyer(b)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"text": "create success",
		"data": response,
	})
}

func (h buyerHandler) UpdateBuyer(c *fiber.Ctx) error {
	b := service.Buyer_order{}
	err := c.BodyParser(&b)
	if err != nil {
		return err
	}
	bb := service.Buyer_order{
		Order_id:     b.Order_id,
		Buyer_name:   b.Buyer_name,
		Order_status: b.Order_status,
		Order_date:   b.Order_date,
		Is_active:    b.Is_active,
	}
	response, err := h.buyerSrv.UpdateBuyer(bb)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"text": "update success",
		"data": response,
	})
}

func (h buyerHandler) DeleteBuyer(c *fiber.Ctx) error {
	id := c.Params("id")
	checkid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	err = h.buyerSrv.DeleteBuyer(checkid)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"text": "delete success",
	})
}
