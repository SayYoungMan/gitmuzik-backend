package main

import (
	"github.com/SayYoungMan/gitmuzik-backend/internal/receiver"
)

const (
	dailyPlaylistID string = "PLXzLX2ct6ysab-Gy0b1Xrm9Ka-Pg-yqmR"
)

func main() {
	receiver.ReceiveAndSavePlaylistItems(dailyPlaylistID, "test.json")
}
