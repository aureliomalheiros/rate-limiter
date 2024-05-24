package limiter

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type RateLimiter struct {
	redisClient *redis.Client
	ipLimit     int
	tokenLimit  int
	blockTime   time.Duration
}

func NewRateLimiter() *RateLimiter {
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	ipLimit, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_IP"))
	tokenLimit, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_TOKEN"))
	blockDuration, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_BLOCK_DURATION"))

	return &RateLimiter{
		redisClient: client,
		ipLimit:     ipLimit,
		tokenLimit:  tokenLimit,
		blockTime:   time.Duration(blockDuration) * time.Second,
	}
}

func (rl *RateLimiter) Allow(ip, token string) bool {
	ctx := context.Background()
	key := ""

	if token != "" {
		key = fmt.Sprintf("token:%s", token)
		limit := rl.tokenLimit
		return rl.allowRequest(ctx, key, limit)
	}

	key = fmt.Sprintf("ip:%s", ip)
	limit := rl.ipLimit
	return rl.allowRequest(ctx, key, limit)
}

func (rl *RateLimiter) allowRequest(ctx context.Context, key string, limit int) bool {
	current, err := rl.redisClient.Get(ctx, key).Int()
	if err == redis.Nil {
		rl.redisClient.Set(ctx, key, 1, rl.blockTime)
		return true
	} else if err != nil {
		return false
	}

	if current < limit {
		rl.redisClient.Incr(ctx, key)
		return true
	}

	return false
}
