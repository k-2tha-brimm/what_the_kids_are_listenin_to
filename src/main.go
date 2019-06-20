package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")

	var port string = ":3000"
	fmt.Println("Listening on port " + port)
	http.ListenAndServe(port, r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	spanishArtists := ArtistsAPIResponse{}
	getJson("http://ws.audioscrobbler.com/2.0/?method=geo.gettopartists&country=japan&limit=10&api_key=5cdf39c88d18d6dd486af4a7036787b7&format=json", &spanishArtists)
	artistsJson, err := json.Marshal(spanishArtists)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(artistsJson)
	fmt.Println("Your JSON has rendered, sir.")
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type ArtistsAPIResponse struct {
	TopArtists struct {
		Attr struct {
			Country string `json:"country"`
			Page    string `json:"page"`
			PerPage string `json:"perPage"`
			Total   string `json:"total"`
		} `json:"@attr"`
		Artist []struct {
			Listeners string `json:"listeners"`
			MBID      string `json:"mbid"`
			Name      string `json:"name"`
			URL       string `json:"url"`
		} `json:"artist"`
	} `json:"topArtists"`
}
