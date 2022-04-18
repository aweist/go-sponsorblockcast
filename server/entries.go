package server

import (
	"encoding/json"
	"net/http"

	"github.com/aweist/go-sponsorblockcast/models"
)

func handleEntries(chromecastMap models.ChromecastMapIface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chromecastmap := chromecastMap.Entries()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(chromecastmap)
	}
}
