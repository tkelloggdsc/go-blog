package main

import (
	"encoding/json"
	"net/http"
)

// PostIndexResponse - structure of post index
type PostIndexResponse struct {
	Posts []Post
}

// PostsIndex - return all posts
func PostsIndex(w http.ResponseWriter, r *http.Request) {
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
