package main

import (
	"blogposts"
	"log"
	"os"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range posts {
		log.Println(p)
	}
}
