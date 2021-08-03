package main

import (
	"encoding/json"
	"io/ioutil"
)

// Video represents the motion picture entity of our app
type Video struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
	URL         string `json:"url"`
}

func getVideos() (videos []Video) {
	fileBytes, err := ioutil.ReadFile("./videos.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &videos)

	if err != nil {
		panic(err)
	}

	return videos
}

func saveVideos(videos []Video) {
	videoBytes, err := json.Marshal(videos)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./videos.json", videoBytes, 0644)

	if err != nil {
		panic(err)
	}
}
