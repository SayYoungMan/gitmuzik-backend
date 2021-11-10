package main

import (
	"fmt"

	"github.com/SayYoungMan/gitmuzik-backend/internal/receiver"
)

const (
	dailyPlaylistID string = "PLXzLX2ct6ysab-Gy0b1Xrm9Ka-Pg-yqmR"
)

func main() {
	items := receiver.GetPlaylistItems(dailyPlaylistID)
	for _, item := range items {
		fmt.Println("Title: ", item.Snippet.Title)
		fmt.Println("Added on: ", item.Snippet.PublishedAt)
		fmt.Println("ID: ", item.Id)
		fmt.Println("VidID: ", item.Snippet.ResourceId.VideoId)
		fmt.Println("Owner ID: ", item.Snippet.VideoOwnerChannelId)
		fmt.Println("Owner: ", item.Snippet.VideoOwnerChannelTitle)
		fmt.Println("Position: ", item.Snippet.Position)
		fmt.Println("Description: ", item.Snippet.Description)
	}
}
