package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// 'videos get' subcommand
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	// inputs for 'videos get' command
	getAll := getCmd.Bool("all", false, "Get all videos")
	getID := getCmd.String("id", "", "YouTube video ID")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addID := addCmd.String("id", "", "YouTube video ID")
	addTitle := addCmd.String("title", "", "YouTube video title")
	addURL := addCmd.String("url", "", "YouTube video URL")
	addImageURL := addCmd.String("imageurl", "", "YouTube video image url")
	addDesc := addCmd.String("desc", "", "YouTube video description")

	if len(os.Args) < 2 {
		fmt.Println("expected a 'get' or 'add' subcommands.")
		os.Exit(1)
	}

	switch os.Args[1] {
		case "get":
			HandleGet(getCmd, getAll, getID)
		case "add":
			HandleAdd(addCmd, addID, addTitle, addURL, addImageURL, addDesc)
		default:
			fmt.Println("Command not found. Try 'get' or 'add'")
	}
}

// HandleGet is the get videos handler. It returns all videos or a single video by id
func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	getCmd.Parse(os.Args[2:])

	if !*all && *id == "" {
		fmt.Println("id is required or specify --all for all videos")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *all {
		// return all videos
		videos := getVideos()

		fmt.Printf("ID \tTitle \tURL \tImageURL \tDescription \n")
		for _, video := range videos {
			fmt.Printf("%v \t%v \t%v \t%v \t%v \n", video.ID, video.Title, video.URL, video.ImageURL, video.Description)
		}
		return
	}

	if *id != "" {
		videos := getVideos()
		id := *id

		for _, video := range videos {
			if id == video.ID {
				fmt.Printf("ID \tTitle \tURL \tImageURL \tDescription \n")
				fmt.Printf("%v \t%v \t%v \t%v \t%v \n", video.ID, video.Title, video.URL, video.ImageURL, video.Description)
			}
		}
	}
}

// ValidateVideo validates that all the required fields for adding a video are present
func ValidateVideo(addCmd *flag.FlagSet, id *string, title *string, url *string, imageurl *string, desc *string) {
	if *id == "" || *title == "" || *url == "" || *imageurl == "" || *desc == "" {
		fmt.Println("all fields are required for adding a video")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
}

// HandleAdd is the add video handler. It adds a video to the list of videos
func HandleAdd(addCmd *flag.FlagSet, id *string, title *string, url *string, imageurl *string, desc *string) {
	ValidateVideo(addCmd, id, title, url, imageurl, desc)

	video := Video {
		ID: *id,
		Title: *title,
		URL: *url,
		ImageURL: *imageurl,
		Description: *desc,
	}

	videos := getVideos()
	videos = append(videos, video)

	saveVideos(videos)
}