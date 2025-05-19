package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sanket9162/hotel-reservation/db"
	"github.com/sanket9162/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)

	hotel := types.Hotel{
		Name:     "Bellucia",
		Location: "India",
	}

	room := types.Rooms{
		Type:      types.SingleRoomType,
		BasePrice: 2500,
	}
	_ = room
	insertedHotel, err := hotelStore.InserHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	room.HotelID = insertedHotel.ID

	fmt.Println(insertedHotel)
}
