package util

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/msgpack/v5"
	"log"
)

func SendMessage(client *redis.Client, rooms []string) {
	data := []interface{}{
		999,
		map[string]interface{}{
			"nsp":  "/",
			"type": 2,
			"data": []interface{}{
				"channel_message", // type
				map[string]interface{}{
					"data": "AA",
				},
			},
		},
		map[string]interface{}{
			"rooms": rooms,
		},
	}

	remoteJoinMessage, err := msgpack.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal remoteJoin message with msgpack: %v", err)
	}

	channelName := "socket.io#/#"
	err = client.Publish(context.Background(), channelName, remoteJoinMessage).Err()
	if err != nil {
		log.Fatalf("Failed to publish remoteJoin message: %v", err)
	}
}
