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
		fmt.Println(item.Snippet.Title)
	}
}
