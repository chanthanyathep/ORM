package main

import (
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
	app.Get("/", buyerHandler.GetBuyers)

	log.Fatal(app.Listen(":8005"))

}
