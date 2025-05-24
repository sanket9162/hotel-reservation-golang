package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sanket9162/hotel-reservation/db"
	"github.com/sanket9162/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	userStore  db.UserStore
	hotelStore db.HotelStore
	roomStore  db.RoomStore
	ctx        = context.Background()
)

func seedHotel(name, location string) {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
	}

	insertedHotel, err := hotelStore.InserHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("inserted hotel = ", insertedHotel)
	rooms := []types.Room{
		{
			Size:  "small",
			Price: 2500,
		},
		{
			Size:  "normal",
			Price: 4500,
		},
	}

	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		insertedRoom, err := roomStore.InserRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("inserted room = ", insertedRoom)
	}
}

func seedUser(c context.Context, fname, lname, email, password string) {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		FirstName: fname,
		LastName:  lname,
		Email:     email,
		Password:  password,
	})
	if err != nil {
		log.Fatal(err)
	}
	res, err := userStore.InsertUser(c, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted User = ", res)
}

func init() {
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	userStore = db.NewMongoUserStore(client)
	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)

}

func main() {
	seedHotel("Bellucia", "France")
	seedHotel("The cozy hotel", "The Netherlands")
	seedHotel("Die another day", "UK")
	seedUser(context.Background(), "James", "Foo", "james@foo.com", "supersecurepassword")
	fmt.Println("seeded the db")
}
