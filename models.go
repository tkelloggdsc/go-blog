package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Post - blog article
type Post struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Author    string        `bson:"author" json:"author"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	Title     string        `bson:"title" json:"title"`
	Body      string        `bson:"body" json:"body"`
	Tags      []string      `bson:"tags" json:"tags"`
}
