package main

import (
	"encoding/json"
	"io/ioutil"
)

// Video is the type that defines what a video should be
type Video struct {
	ID string
	Title string
	Description string
	ImageURL string
	URL string
}

// getVideos gets and returns the videos from a json file
func getVideos() (videos []Video) {
	fileByes, err := ioutil.ReadFile("./videos.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileByes, &videos)
	if err != nil {
		panic(err)
	}

	return videos
}

// saveVideos saves a video to the json file
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