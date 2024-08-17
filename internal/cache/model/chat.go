package model

type MessageModel struct {
	ID          int    `redis:"id"`
	Content     string `redis:"content"`
	Author      int    `redis:"author_id"`
	ChatID      int    `redis:"chat_id"`
	TimestampNS int64  `redis:"timestamp_ns"`
}

type ChatModel struct {
	ID          int     `redis:"id"`
	UserIDs     []int64 `redis:"user_ids"`
	CreatedAtNs int64   `redis:"created_at_ns"`
	UpdatedAtNs int64   `redis:"updated_at_ns"`
}
