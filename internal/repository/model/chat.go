package model

import "time"

// CreateChatModel representation of create chat on repository layer.
type CreateChatModel struct {
	UserIDs []int64
}

// ChatModel - representation of a chat on repository layer.
type ChatModel struct {
	ID        int
	UserIDs   []int64
	UpdatedAt *time.Time
	CreatedAt *time.Time
}

// CreateMessageModel - representation of create message on repository layer.
type CreateMessageModel struct {
	Content  string
	AuthorID int
	ChatID   int
}

// MessageModel - representation of message on repository layer.
type MessageModel struct {
	ID        int
	Content   string
	Author    int
	ChatID    int
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// DeleteChatModel - representation of delete chat model.
type DeleteChatModel struct {
	ID int
}
