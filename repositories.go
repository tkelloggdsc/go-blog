package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

// DAO - database access object used to interact with the database
type DAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const postsCollection = "post"

// Connect - establish connection with DB through Url
func (d *DAO) Connect() {
	session, err := mgo.Dial(d.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(d.Database)
}

// FindPostByID - finds a post by object id
func (d *DAO) FindPostByID(id string) (Post, error) {
	var post Post
	err := db.C(postsCollection).FindId(id).One(&post)
	return post, err
}

// FindAllPosts - finds all posts
func (d *DAO) FindAllPosts() ([]Post, error) {
	var posts []Post
	err := db.C(postsCollection).Find(nil).All(&posts)
	return posts, err
}

// InsertPost - saves a post
func (d *DAO) InsertPost(post Post) error {
	return db.C(postsCollection).Insert(&post)
}
