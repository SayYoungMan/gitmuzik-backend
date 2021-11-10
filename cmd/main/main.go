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

	processed := receiver.ProcessPlaylistItems(items)

	for _, item := range processed {
		fmt.Println("Title: ", item.Title)
		fmt.Println("Added on: ", item.AddDate)
		fmt.Println("ID: ", item.ID)
		fmt.Println("VidID: ", item.VidID)
		fmt.Println("Owner ID: ", item.OwnerID)
		fmt.Println("Owner: ", item.OwnerTitle)
		fmt.Println("Position: ", item.Position)
	}
}
