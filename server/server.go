package server

import (
	"log"
	"net/http"

	"github.com/aweist/go-sponsorblockcast/models"
)

func Serve(chromecastMap models.ChromecastMapIface) {
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/entries", handleEntries(chromecastMap))

	log.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
