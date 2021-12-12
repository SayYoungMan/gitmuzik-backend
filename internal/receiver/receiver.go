package receiver

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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

type itemEntrySlice []itemEntry

func ReceiveAndSavePlaylistItems(ctx context.Context, playlistID string, jsonpath string) {
	items := GetPlaylistItems(ctx, playlistID)
	processedItems := processPlaylistItems(items)
	jsonbytes := processedItems.makeItemEntryJSON(ctx)
	err := ioutil.WriteFile(jsonpath, jsonbytes, 0644)
	if err != nil {
		log.Fatalf("Saving JSON file failed: %v", err)
	}
}

func MoveFileToS3(
	ctx context.Context,
	uploader *s3manager.Uploader,
	filepath string,
	bucketName string,
	key string,
	removeAfter bool,
) error {
	// Open the file in filepath
	f, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", filepath, err)
	}

	// Upload the file to s3
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: &bucketName,
		Key:    &key,
		Body:   f,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	if removeAfter {
		if err := os.Remove(filepath); err != nil {
			return fmt.Errorf("failed to remove file, %v", err)
		}
	}
	return nil
}

func GetPlaylistItems(ctx context.Context, playlistID string) []*youtube.PlaylistItem {
	apiKey := getAPIKey()

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
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

func processPlaylistItems(playlistItems []*youtube.PlaylistItem) itemEntrySlice {
	var (
		tmp itemEntry
		rv  itemEntrySlice
	)

	for _, item := range playlistItems {
		tmp.extract(item)
		rv = append(rv, tmp)
	}

	return rv
}

func (items *itemEntrySlice) makeItemEntryJSON(ctx context.Context) []byte {
	b, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		log.Fatalf("JSON Marshalling failed: %v\n", err)
		return nil
	} else {
		return b
	}
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
