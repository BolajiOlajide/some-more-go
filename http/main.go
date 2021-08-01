package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleGetVideos)
	http.HandleFunc("/update", HandleUpdateVideos)

	http.ListenAndServe(":8080", nil)
}

// HandleGetVideos handler for fetching all videos
func HandleGetVideos(w http.ResponseWriter, r *http.Request) {
	// for header, value := range r.Header {
	// 	fmt.Printf("Key: %v \tValue: %v\n", header, value)
	// }

	videos := getVideos()

	videoBytes, err := json.Marshal(videos)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(videoBytes)
}

func HandleUpdateVideos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		var videos []Video

		err = json.Unmarshal(body, &videos)

		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Bad Request")
		}

		saveVideos(videos)
	} else {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method not Supported")
	}
}
