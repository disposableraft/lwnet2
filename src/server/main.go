package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"lancewakeling.net/website/src/gatsby"
)

func main() {
	gatsby.BuildSite()

	fileServer := http.FileServer(http.Dir("./docs"))
	http.Handle("/", fileServer)

	// ---

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				gatsby.BuildSite()

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("templates")
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Add("src/content/writing/blog")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting server at port http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	<-done
}
