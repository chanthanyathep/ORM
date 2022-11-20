package handler

import (
	"orm/service"

	"github.com/gofiber/fiber/v2"
)

type buyerHandler struct {
	buyerSrv service.BuyerService
}

func NewBuyerHandler(buyerSrv service.BuyerService) buyerHandler {
	return buyerHandler{buyerSrv: buyerSrv}
}

func (h buyerHandler) GetBuyers(c *fiber.Ctx) error {
	buyer, err := h.buyerSrv.GetBuyers()
	if err != nil {
		panic(err)
	}
	reponse := fiber.Map{
		"status": "os",
		"buyer":  buyer,
	}
	return c.JSON(reponse)
}
