package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	var redisHost = "localhost:6379"
	var redisPassword = ""

	rdb := newRedisClient(redisHost, redisPassword)
	ttl := time.Duration(1) * time.Minute
	set := rdb.Set(context.Background(), "nama_mahasiswa", "Helmi Adi Prasetyo", ttl)

	if err := set.Err(); err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
		return
	}
	opt2 := rdb.Get(context.Background(), "nama_mahasiswa")
	if err := opt2.Err(); err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return
	}

	result, err := opt2.Result()
	if err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return
	}
	fmt.Println("nama mahasiswanya adalah ", result)
	fmt.Println("redis client initialized")
}

func newRedisClient(host string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client
}
