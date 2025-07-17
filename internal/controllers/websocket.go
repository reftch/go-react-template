package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func (c *Controller) WsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	// Watch for file changes
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()
		watchDirsRecursive(watcher, "./web/static/js")
		for event := range watcher.Events {
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Println("XAXAXA ddd")
				err = conn.WriteMessage(websocket.TextMessage, []byte("reload"))
				if err != nil {
					return
				}
			}
		}
	}()

	// Keep the connection open to handle further messages
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

// Add all directories recursively to the watcher
func watchDirsRecursive(watcher *fsnotify.Watcher, root string) error {
	return filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if err := watcher.Add(path); err != nil {
				log.Printf("failed to watch directory %s: %v", path, err)
			}
		}
		return nil
	})
}
