package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sanket9162/hotel-reservation/api"
	"github.com/sanket9162/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://root:example@localhost:27017/"
const dbname = "hotel-reservation"
const userColl = "users"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})

	},
}

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	// handlers initialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	listenAddr := flag.String("listenAddr", "localhost:4000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("sanket")
	})

	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	app.Listen(*listenAddr)
}
