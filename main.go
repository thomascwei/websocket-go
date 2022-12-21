package main

import (
	selfLogger "github.com/thomascwei/golang_logger"
	"net/http"
)

var Logger = selfLogger.InitLogger("main")

func main() {
	setupAPI()

	// Serve on port :8080, fudge yeah hardcoded port
	Logger.Fatal(http.ListenAndServe(":8080", nil))
}

// setupAPI will start all Routes and their Handlers
func setupAPI() {
	// Create a Manager instance used to handle WebSocket Connections
	manager := NewManager()

	// Serve the ./frontend directory at Route /
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.serveWS)

}
