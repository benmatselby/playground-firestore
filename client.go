package main

import (
	"context"
	"log"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const CollectionRoom = "rooms"
const CollectionMessage = "messages"

func NewClient() Client {
	ctx := context.Background()
	sa := option.WithCredentialsFile("/tmp/service-account.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	firestore, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	client := Client{
		firestore: firestore,
		context:   ctx,
	}

	return client
}

type Client struct {
	firestore *firestore.Client
	context   context.Context
}

func (c *Client) CreateRoom(document map[string]interface{}) {
	log.Printf("Creating room: %v", document)
	_, _, err := c.firestore.Collection(CollectionRoom).Add(c.context, document)
	if err != nil {
		log.Fatalf("Failed adding document: %v", err)
	}
}

func (c *Client) SendMessage(room string, document map[string]interface{}) {
	log.Printf("Sending message: %v", document)
	_, _, err := c.firestore.Collection(CollectionRoom).Doc(room).Collection(CollectionMessage).Add(c.context, document)
	if err != nil {
		log.Fatalf("Failed adding document: %v", err)
	}
}

func (c *Client) GetAllMessagesForRoom(room string) []map[string]interface{} {
	iter := c.firestore.Collection(CollectionRoom).Doc(room).Collection(CollectionMessage).Where("category", "==", "welcome").Documents(c.context)
	docs := []map[string]interface{}{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to get a document: %v", err)
		}

		docs = append(docs, doc.Data())
	}

	return docs
}

func (c *Client) Close() {
	c.firestore.Close()
}
