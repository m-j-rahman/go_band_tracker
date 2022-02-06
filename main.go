package main

import (
	"fmt"
	tracker "tracker/groupie"
)

func main() {
	fmt.Printf("Starting server...\n")
	fmt.Printf("Listening on port 8080!\n")
	tracker.OpenBrowser("http://localhost:8080")
	tracker.MainHandler()
}
