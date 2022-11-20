package main

import (
	"fmt"
	"log"
	"orm/handler"
	"orm/repository"
	"orm/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db := initDB()

	buyerRepo := repository.NewBuyerRepositoryDB(db)
	buyerServ := service.NewBuyerServ(buyerRepo)
	buyerHandler := handler.NewBuyerHandler(buyerServ)

	app := fiber.New()
	app.Use(logger.New())

	b, _ := buyerServ.GetBuyerById(6000)
	fmt.Println(b)

	app.Get("/", buyerHandler.GetBuyers)
	app.Get("/:id", buyerHandler.GetBuyerById)
	app.Post("/buyer_order/", buyerHandler.CreateBuyer)
	app.Put("/buyer_order/", buyerHandler.UpdateBuyer)
	app.Delete("/:id", buyerHandler.DeleteBuyer)

	log.Fatal(app.Listen(":8005"))

}
