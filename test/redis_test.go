package test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
)

func BenchmarkFindMGet(b *testing.B) {
	ctx := context.Background()

	for n := 0; n < b.N; n++ {
		err := redisClient.MGet(ctx, keys...).Err()
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkFindPipeline(b *testing.B) {
	ctx := context.Background()

	for n := 0; n < b.N; n++ {
		_, err := redisClient.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
			var err error

			for i := 0; i < len(keys); i++ {
				err = pipeliner.Get(ctx, keys[i]).Err()
			}

			return err
		})
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkFindGet(b *testing.B) {
	ctx := context.Background()

	for n := 0; n < b.N; n++ {
		var err error

		for i := 0; i < len(keys); i++ {
			err = redisClient.Get(ctx, keys[i]).Err()
		}

		if err != nil {
			b.Fail()
		}
	}
}
