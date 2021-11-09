package receiver

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func GetPlaylistItems(playlistID string) []*youtube.PlaylistItem {
	client := getClient(youtube.YoutubeReadonlyScope)
	ctx := context.Background()

	service, err := youtube.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Error creating Youtube service: %v", err)
	}
	itemService := youtube.NewPlaylistItemsService(service)

	call := itemService.List([]string{"snippet"})
	call = call.PlaylistId(playlistID)
	call = call.MaxResults(10000)
	response, err := call.Do()
	handleError(err, "")

	return response.Items
}
