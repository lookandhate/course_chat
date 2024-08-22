package chat

import (
	"context"
	"fmt"
	"time"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/lookandhate/course_chat/internal/config"
	"github.com/lookandhate/course_chat/internal/service/convertor"
	serviceModel "github.com/lookandhate/course_chat/internal/service/model"
	"github.com/lookandhate/course_platform_lib/pkg/cache/redis"
)

type RedisCache struct {
	redisClient *redis.Client
}

func NewRedisCache(redisPool *redigo.Pool, redisCfg config.RedisConfig) *RedisCache {
	client := redis.NewClient(redisPool, time.Duration(redisCfg.IdleTimeout))

	return &RedisCache{redisClient: client}
}

func (r RedisCache) CreateChat(ctx context.Context, chat *serviceModel.ChatModel) error {
	chatCache := convertor.ServiceChatModelToCacheChatModel(chat)
	err := r.redisClient.HashSet(ctx, r.chatKey(chatCache.ID), chatCache)
	if err != nil {
		return err
	}

	return nil
}

func (r RedisCache) CreateMessage(ctx context.Context, message *serviceModel.MessageModel) error {
	messageCache := convertor.ServiceMessageModelToCacheMessageModel(message)
	err := r.redisClient.HashSet(ctx, r.messageKey(messageCache.ID), messageCache)
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
