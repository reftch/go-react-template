package controllers

import (
	"log"
	"net/http"

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
		watcher.Add("./web/static/js/index.js") // Watch the specific file
		for event := range watcher.Events {
			if event.Op&fsnotify.Write == fsnotify.Write {
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
