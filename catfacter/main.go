package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// BSON Binary serialization of JSON data, this how MongoDB stores data

type CatFactWorker struct {
	client *mongo.Client
}

func NewCatFactWorker(c *mongo.Client) *CatFactWorker {
	return &CatFactWorker{
		client: c,
	}
}

func (cfw *CatFactWorker) start() error {
	// coll := cfw.client.Database("catfact").Collection("facts")

	// we dont want constantly polling that endpoint, hence need for ticker
	ticker := time.NewTicker(2 * time.Second)

	for {
		// this link spits out random facts about Cats
		resp, err := http.Get("https://catfact.ninja/fact")
		if err != nil {
			return err
		}
		var catFact bson.M // map[string]any // map[string]interface{}

		if err := json.NewDecoder(resp.Body).Decode(&catFact); err != nil {
			return err
		}

		fmt.Println(catFact)

		<-ticker.C
	}

}

func main() {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	worker := NewCatFactWorker(client)
	worker.start()

}
