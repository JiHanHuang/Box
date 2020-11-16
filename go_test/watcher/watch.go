package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	log.Println("Process start...")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/data/conf/config.sys")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
