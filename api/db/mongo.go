package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	EnvConnectionURI = "CONNECTION_URI"
)

const (
	DbName = "mantil"
)

var (
	ErrNotFound = fmt.Errorf("not found")
)

type Mdb struct {
	client *mongo.Client
	db     *mongo.Database
}

func connect() (*Mdb, error) {
	conn, ok := os.LookupEnv(EnvConnectionURI)
	if !ok {
		return nil, fmt.Errorf("connection URI for mongo not set")
	}
	clientOptions := options.Client().ApplyURI(conn).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error connecting to mongo db - %v", err)
	}
	return &Mdb{
		client: client,
		db:     client.Database(DbName),
	}, nil
}

func (mdb *Mdb) Disconnect() {
	mdb.client.Disconnect(context.Background())
}

func (mdb *Mdb) Get(col string, id interface{}, o interface{}) error {
	c := mdb.db.Collection(col)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := c.FindOne(ctx, bson.D{{"_id", id}}).Decode(o)
	if err == mongo.ErrNoDocuments {
		return ErrNotFound
	}
	return err
}

func (mdb *Mdb) Create(col string, o interface{}) error {
	c := mdb.db.Collection(col)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := c.InsertOne(ctx, o)
	return err
}

func (mdb *Mdb) Delete(col string, id interface{}) error {
	c := mdb.db.Collection(col)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	dr, err := c.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	if dr.DeletedCount == 0 {
		return ErrNotFound
	}
	return nil
}
