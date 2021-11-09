package receiver

import (
	"fmt"
	"log"

	"google.golang.org/api/youtube/v3"
)

func GetPlaylist() {
	client := getClient(youtube.YoutubeReadonlyScope)

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating Youtube client: %v", err)
	}

	call := service.Playlists.List([]string{"snippet", "contentDetails"})
	call = call.Id("PLXzLX2ct6ysab-Gy0b1Xrm9Ka-Pg-yqmR")
	response, err := call.Do()
	handleError(err, "")

	for _, playlist := range response.Items {
		fmt.Println(playlist.Id, ": ", playlist.Snippet.Title)
	}
}
