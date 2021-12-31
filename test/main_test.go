package test

import (
	"bench/parser"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"os"
	"testing"
)

var models []parser.TestModel
var keys []string
var redisClient *redis.Client

const (
	batchSize = 200
)

func TestMain(m *testing.M) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		PoolSize: 200,
	})

	setup()

	defer teardown()

	os.Exit(m.Run())
}

func setup() {
	ctx := context.Background()

	model := parser.TestModel{
		Test1: rand.Int63(),
		Test2: generateRandomString(100),
		Test3: true,
		Test4: rand.Int31(),
		Test5: generateRandomStringSlice(100, 100),
		Test6: generateRandomIntSlice(100),
	}

	models = make([]parser.TestModel, batchSize)
	keys = make([]string, batchSize)

	for i := range models {
		models[i] = model

		jp := parser.NewJsonParser()

		err := jp.Encode(&model)
		if err != nil {
			continue
		}

		key := fmt.Sprintf("%v", i)

		keys[i] = key

		redisClient.Set(ctx, key, jp.Data, -1)
	}
}

func teardown() {
	ctx := context.Background()

	redisClient.FlushAll(ctx)
	if err := redisClient.Close(); err != nil {
		fmt.Println("cannot disconnect from redis connection")
	}
}

func generateRandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func generateRandomStringSlice(n int, m int) []string {
	var res []string
	for i := 0; i < m; i++ {
		res = append(res, generateRandomString(n))
	}
	return res
}

func generateRandomIntSlice(m int) []int64 {
	var res []int64
	for i := 0; i < m; i++ {
		res = append(res, rand.Int63())
	}
	return res
}
