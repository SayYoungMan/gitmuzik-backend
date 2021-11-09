package receiver

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const (
	dailyPlaylistID string = "PLXzLX2ct6ysab-Gy0b1Xrm9Ka-Pg-yqmR"
)

func GetPlaylist() {
	client := getClient(youtube.YoutubeReadonlyScope)
	ctx := context.Background()

	service, err := youtube.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Error creating Youtube client: %v", err)
	}

	call := service.Playlists.List([]string{"snippet", "contentDetails"})
	call = call.Id(dailyPlaylistID)
	response, err := call.Do()
	handleError(err, "")

	for _, playlist := range response.Items {
		fmt.Println(playlist.Id, ": ", playlist.Snippet.Title)
	}
}
