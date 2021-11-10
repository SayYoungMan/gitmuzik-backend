package receiver

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const (
	maxResult int64 = 10000
)

type itemEntry struct {
	Title       string
	AddDate     string
	ID          string
	VidID       string
	OwnerID     string
	OwnerTitle  string
	Position    int64
	Description string
}

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
	call = call.MaxResults(maxResult)
	response, err := call.Do()
	handleError(err, "")

	return response.Items
}

func ProcessPlaylistItems(playlistItems []*youtube.PlaylistItem) []itemEntry {
	var (
		tmp itemEntry
		rv  []itemEntry
	)

	for _, item := range playlistItems {
		tmp.extract(item)
		rv = append(rv, tmp)
	}

	return rv
}

func (tmp *itemEntry) extract(item *youtube.PlaylistItem) {
	tmp.Title = item.Snippet.Title
	tmp.AddDate = item.Snippet.PublishedAt
	tmp.ID = item.Id
	tmp.VidID = item.Snippet.ResourceId.VideoId
	tmp.OwnerID = item.Snippet.VideoOwnerChannelId
	tmp.OwnerTitle = item.Snippet.VideoOwnerChannelTitle
	tmp.Position = item.Snippet.Position
	tmp.Description = item.Snippet.Description
}
