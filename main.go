package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/johnnypivot/ziplookup/zippopotam"
	log "github.com/sirupsen/logrus"
)

func main() {
	hc := http.Client{Timeout: 10 * time.Second}
	zc := zippopotam.NewClient("http://api.zippopotam.us/us/", &hc)

	r := mux.NewRouter()
	r.HandleFunc("/{zip}", func(w http.ResponseWriter, r *http.Request) {
		zip, ok := mux.Vars(r)["zip"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if len(zip) != 5 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.WithFields(log.Fields{"zip": zip}).Info("Looking up zip")

		result, err := zc.Lookup(zip)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.WithField("place", result).Info("Got result")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	log.Info("Listening on :8080")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
