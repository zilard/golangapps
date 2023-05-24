package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CatFact struct {
	Fact   string `bson:"fact"`
	Length int    `bson:"length"`
}

type Storer interface {
	GetAll() ([]*CatFact, error)
	Put(*CatFact) error
}

type MongoStore struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewMongoStore() (*MongoStore, error) {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017")) //27017 - default MongoDB port
	if err != nil {
		return nil, err
	}

	return &MongoStore{
		client:     client,
		database:   "catfact",
		collection: "facts",
	}, nil

}

func (store *MongoStore) Put(fact *CatFact) error {
	coll := store.client.Database(store.database).Collection(store.collection)
	_, err := coll.InsertOne(context.TODO(), fact)
	return err
}

func (store *MongoStore) GetAll() ([]*CatFact, error) {
	coll := store.client.Database(store.database).Collection(store.collection) // collection

	// query
	query := bson.M{}
	cursor, err := coll.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}

	results := []*CatFact{} // slice of CatFact
	//check for errors in the conversion
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

type Server struct {
	store Storer
}

func NewServer(s Storer) *Server {
	return &Server{
		store: s,
	}
}

func (s *Server) handleGetAllFacts(w http.ResponseWriter, r *http.Request) {
	catFacts, err := s.store.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(catFacts)

}

// BSON Binary serialization of JSON data, this how MongoDB stores data
type CatFactWorker struct {
	store       Storer
	apiEndpoint string
}

func NewCatFactWorker(store Storer, apiEndpoint string) *CatFactWorker {
	return &CatFactWorker{
		store:       store,
		apiEndpoint: apiEndpoint,
	}
}

func (cfw *CatFactWorker) start() error {

	// we dont want constantly polling that endpoint, hence need for ticker
	ticker := time.NewTicker(2 * time.Second)

	for {
		// this link spits out random facts about Cats
		resp, err := http.Get(cfw.apiEndpoint)
		if err != nil {
			return err
		}
		var catFact CatFact // map[string]any // map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&catFact); err != nil {
			return err
		}

		if err := cfw.store.Put(&catFact); err != nil {
			return err
		}

		<-ticker.C
	}

}

func main() {
	mongoStore, err := NewMongoStore()
	if err != nil {
		log.Fatal(err)
	}

	worker := NewCatFactWorker(mongoStore, "https://catfact.ninja/fact")
	go worker.start()

	server := NewServer(mongoStore)
	// expose a route /facts
	http.HandleFunc("/facts", server.handleGetAllFacts)
	// boot up a server here
	http.ListenAndServe(":3000", nil)
}
