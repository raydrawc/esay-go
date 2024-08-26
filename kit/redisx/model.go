package redisx

import "github.com/redis/go-redis/v9"

type Client struct {
	redis.Client
}

// NewClient 初始化redis客户端
func NewClient(opt *redis.Options) *Client {
	return &Client{Client: *redis.NewClient(opt)}
}
