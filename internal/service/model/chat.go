package model

import (
	"database/sql"
	"time"
)

// ChatModel is service layer chat representation.
type ChatModel struct {
	ID        int
	UserIDs   []int64
	ChatID    int
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

// CreateChat is service layer create chat representation.
type CreateChat struct {
	UserIDs []int64
}

// CreateMessage is service layer message creation representation.
type CreateMessage struct {
	ChatID    int
	AuthorID  int
	Content   string
	Timestamp time.Time
}

// MessageModel is service layer message representation.
type MessageModel struct {
	ID        int
	ChatID    int
	AuthorID  int
	Content   string
	Timestamp time.Time
}
