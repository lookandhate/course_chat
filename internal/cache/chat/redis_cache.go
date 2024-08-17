package chat

import (
	"context"
	"fmt"
	"time"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/lookandhate/course_chat/internal/cache/model"
	"github.com/lookandhate/course_chat/internal/config"
	"github.com/lookandhate/course_platform_lib/pkg/cache/redis"
)

type RedisCache struct {
	redisClient *redis.Client
}

func NewRedisCache(redisPool *redigo.Pool, redisCfg config.RedisConfig) *RedisCache {
	client := redis.NewClient(redisPool, time.Duration(redisCfg.IdleTimeout))

	return &RedisCache{redisClient: client}
}

func (r RedisCache) CreateChat(ctx context.Context, chat *model.ChatModel) error {
	err := r.redisClient.HashSet(ctx, r.chatKey(chat.ID), chat)
	if err != nil {
		return err
	}

	return nil
}

func (r RedisCache) CreateMessage(ctx context.Context, message *model.MessageModel) error {
	err := r.redisClient.HashSet(ctx, r.messageKey(message.ID), message)
	if err != nil {
		return err
	}

	return nil
}

func (r RedisCache) DeleteMessage(ctx context.Context, id int) error {
	err := r.redisClient.Del(ctx, r.messageKey(id))
	if err != nil {
		return err
	}

	return nil
}

func (r RedisCache) DeleteChat(ctx context.Context, id int) error {
	err := r.redisClient.Del(ctx, r.chatKey(id))
	if err != nil {
		return err
	}

	return nil
}

func (r RedisCache) chatKey(id int) string {
	return fmt.Sprintf("chat_id-%d", id)
}

func (r RedisCache) messageKey(id int) string {
	return fmt.Sprintf("message_id-%d", id)
}
