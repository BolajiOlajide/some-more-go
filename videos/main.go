package main

import "fmt"

func main() {
	videos := getVideos()

	for idx := range videos {
		// the `video` (second arg) variable returns a copy of the item
		// in the array so it's not mutable per se
		videos[idx].Description = videos[idx].Description + "\nRemember to Like & Subscribe"
	}

	fmt.Println(videos)

	saveVideos(videos)
}
