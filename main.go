package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"database/sql"

	_ "github.com/lib/pq"
)



var (
	DB *sql.DB
)

func main()  {
	err := OpenDB()
	if err!= nil{
		fmt.Errorf("database connection failed %v :",err)
	}

	defer DB.Close()

	if err := CreateTable(); err != nil {
		log.Fatal(err)
	}



	// Routes and Server config
	app := fiber.New()
	app.Use(DBMiddleware)

	app.Get("/expenses",GetExpenses)
	app.Get("/expense/:id", GetExpense)
	app.Post("/expense",CreateExpense)
	app.Delete("/expense/:id",DeleteExpense)
	app.Put("/expense/:id",UpdateExpense)

	log.Fatal(app.Listen(":3000"))
}

func DBMiddleware(c *fiber.Ctx) error {
	c.Locals("DB",DB)
	return c.Next()
}







