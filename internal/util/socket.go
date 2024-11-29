package util

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/msgpack/v5"
	"log"
)

func SendSocketMessage(client *redis.Client, rooms []string, event string, data interface{}) {
	msg := []interface{}{
		999,
		map[string]interface{}{
			"nsp":  "/",
			"type": 2,
			"data": []interface{}{
				event, // type
				data,
			},
		},
		map[string]interface{}{
			"rooms": rooms,
		},
	}

	sendMessage, err := msgpack.Marshal(msg)
	if err != nil {
		log.Fatalf("Failed to marshal remoteJoin message with msgpack: %v", err)
	}

	channelName := "socket.io#/#"
	err = client.Publish(context.Background(), channelName, sendMessage).Err()
	if err != nil {
		log.Fatalf("Failed to publish remoteJoin message: %v", err)
	}
}
