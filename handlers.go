package main

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// IndexHandler - serves client application
func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}

// AllPostsEndpoint - return all posts
func AllPostsEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var posts []Post
	posts, err := dao.FindAllPosts()
	if err != nil {
		handleError(err, InternalServiceError, w)
		return
	}

	response, err := json.Marshal(posts)
	if err != nil {
		handleError(err, InternalServiceError, w)
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}

// CreatePostEndpoint - save a new post
func CreatePostEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var post Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		handleError(err, InternalServiceError, w)
		return
	}

	post.ID = bson.NewObjectId()
	if err := dao.InsertPost(post); err != nil {
		handleError(err, InternalServiceError, w)
		return
	}

	response, err := json.Marshal(post)
	if err != nil {
		handleError(err, InternalServiceError, w)
		return
	}

	w.WriteHeader(201)
	w.Write(response)
}
