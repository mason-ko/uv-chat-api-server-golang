package util

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestSocket(t *testing.T) {
	// Redis 클라이언트 생성
	c := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 서버 주소
	})
	x := c.Ping(context.Background())
	fmt.Println(x.Result())
	//SendSocketMessage(c, []string{"AA"})
}
