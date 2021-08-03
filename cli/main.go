package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// `videos get` subcommand
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// inputs for `videos get` command
	getAll := getCmd.Bool("all", false, "Get all videos")
	getID := getCmd.String("id", "", "YouTube video ID")

	// `videos add` subcommand
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	// inputs for `videos add` command
	addID := addCmd.String("id", "", "Youtube video ID")
	addTitle := addCmd.String("title", "", "Youtube video Title")
	addUrl := addCmd.String("url", "", "Youtube video URL")
	addImageUrl := addCmd.String("imageUrl", "", "Youtube video Image URL")
	addDesc := addCmd.String("desc", "", "Youtube video description")

	if len(os.Args) < 2 {
		fmt.Println("expect `get` or `add` subcommands")
		os.Exit(1)
	}

	// look at 2nd argument's value
	switch os.Args[1] {
	case "get": // if it's the `get` command
		// handle get here
		HandleGet(getCmd, getAll, getID)
	case "add": // if it's the `add` command
		// handle add here
		HandleAdd(addCmd, addID, addTitle, addUrl, addImageUrl, addDesc)
	default: // if we don't understand the input
	}
}

// HandleGet handler for getting videos from the JSON
func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	getCmd.Parse(os.Args[2:])

	if *all == false && *id == "" {
		fmt.Print("id is required or specify --all for al videos")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	// return all videos
	videos := getVideos()

	if *all {
		fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")

		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.ID, video.Title, video.URL, video.ImageURL, video.Description)
		}

		return
	}

	if *id != "" {
		for _, video := range videos {
			if *id == video.ID {
				fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.ID, video.Title, video.URL, video.ImageURL, video.Description)
			}
		}
	}
}

// HandleAdd handler for adding a video to the JSON
func HandleAdd(addCmd *flag.FlagSet, id *string, title *string, url *string, imageUrl *string, desc *string) {
	addCmd.Parse(os.Args[2:])
	ValideVideo(addCmd, id, title, url, imageUrl, desc)

	video := Video{
		ID:          *id,
		Title:       *title,
		Description: *desc,
		ImageURL:    *imageUrl,
		URL:         *url,
	}

	videos := getVideos()
	videos = append(videos, video)

	saveVideos(videos)
}

func ValideVideo(addCmd *flag.FlagSet, id *string, title *string, url *string, imageUrl *string, desc *string) {
	if *id == "" || *title == "" || *url == "" || *imageUrl == "" || *desc == "" {
		fmt.Print("all fields are required for adding a video")
		addCmd.PrintDefaults()
		fmt.Print("Something no dey work", *id == "", id, *title == "", title)
		os.Exit(1)
	}
}
