package main

import (
	"time"
)

// Post - blog article
type Post struct {
	Author    string
	CreatedAt time.Time
	Title     string
	Body      string
	Tags      []string
}
