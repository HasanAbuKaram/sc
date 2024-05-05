package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

func main() {
	fmt.Printf("Hello, man! Running version %s\n", version)

	// Perform the first update check immediately
	checkForUpdates()

	// Then check for updates every hour
	ticker := time.NewTicker(30 * time.Minute)
	defer ticker.Stop()

	// Create a channel to listen for incoming signals
	sigs := make(chan os.Signal, 1)

	// Register the channel to receive os.Interrupt signals
	signal.Notify(sigs, os.Interrupt)

	go func() {
		for {
			// Wait for an os.Interrupt signal
			sig := <-sigs

			// Ask for user input when an os.Interrupt signal is received
			if sig == os.Interrupt {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Are you sure you want to exit? (y/n): ")
				text, _ := reader.ReadString('\n')
				text = strings.TrimSpace(text) // remove leading and trailing whitespace
				if text == "y" || text == "Y" {
					fmt.Println("Exiting...")
					os.Exit(0)
				} else {
					fmt.Println("Continuing...")
				}
			}
		}
	}()

	// Start a goroutine with an HTTP server
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello this is the first run!")
		})
		http.HandleFunc("GET /time", serverTime)
		http.HandleFunc("GET /login", login)
		//	http.ListenAndServe(":8080", nil)
		err := http.ListenAndServeTLS(":443",
			"./certs/server.crt", "./certs/server.key", nil)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}()

	for {
		<-ticker.C
		checkForUpdates()
	}
}
