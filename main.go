package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sanket9162/hotel-reservation/api"
	"github.com/sanket9162/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const dburi = "mongodb://root:example@localhost:27017/"
const dbname = "hotel-reservation"
const userColl = "users"

func main() {

	client, err := mongo.Connect(options.Client().ApplyURI(dburi))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()
	coll := client.Database(dbname).Collection(userColl)
	user := types.User{
		FirstName: "Sanket",
		LastName:  "GondhaliS",
	}

	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(res)

	listenAddr := flag.String("listenAddr", "localhost:5000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("sanket")
	})

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddr)
}
