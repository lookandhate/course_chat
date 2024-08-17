package model

import (
	"database/sql"
	"time"
)

// CreateChatModel representation of create chat.go on repository layer.
type CreateChatModel struct {
	UserIDs []int64
}

// ChatModel - representation of a chat.go on repository layer.
type ChatModel struct {
	ID        int
	UserIDs   []int64
	UpdatedAt sql.NullTime
	CreatedAt time.Time
}

// CreateMessageModel - representation of create message on repository layer.
type CreateMessageModel struct {
	Content  string
	AuthorID int
	ChatID   int
}

// MessageModel - representation of message on repository layer.
type MessageModel struct {
	ID        int          `db:"id"`
	Content   string       `db:"content"`
	Author    int          `db:"author_id"`
	ChatID    int          `db:"chat_id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

// DeleteChatModel - representation of delete chat.go model.
type DeleteChatModel struct {
	ID int
}
