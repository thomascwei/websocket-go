package main

import (
	"context"
	"fmt"
	selfLogger "github.com/thomascwei/golang_logger"
	"net/http"
)

var Logger = selfLogger.InitLogger("main")

func main() {
	// Create a root ctx and a CancelFunc which can be used to cancel retentionMap goroutine
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	defer cancel()

	setupAPI(ctx)

	// Serve on port :8080, fudge yeah hardcoded port
	Logger.Fatal(http.ListenAndServe(":8080", nil))
}

// setupAPI will start all Routes and their Handlers
func setupAPI(ctx context.Context) {
	// Create a Manager instance used to handle WebSocket Connections
	manager := NewManager(ctx)

	// Serve the ./frontend directory at Route /
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.serveWS)
	http.HandleFunc("/login", manager.loginHandler)
	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, len(manager.clients))
	})
}
