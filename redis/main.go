package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func newRedisClient(host string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client

}

func main() {
	redisHost := "localhost:6379"
	redisPassword := ""

	rdb := newRedisClient(redisHost, redisPassword)

	fmt.Println("redis Client Installed")

	key := "key-1"
	data := "Hello Redis"
	ttl := time.Duration(5) * time.Second

	op1 := rdb.Set(context.Background(), key, data, ttl)

	if err := op1.Err(); err != nil {
		fmt.Println("gagal")
	}

	log.Println("Berhasil set operasi")

	time.Sleep(time.Duration(4)*time.Second)

	op2 := rdb.Get(context.Background(), key)
	if err := op2.Err(); err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return
	}
	res, err := op2.Result()
	if err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return
	}
	log.Println("get operation success. result:", res)
}


