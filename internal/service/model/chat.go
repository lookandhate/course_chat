package model

import "time"

// CreateChatRequest is service layer create chat representation.
type CreateChatRequest struct {
	UserIDs []int64
}

// Chat is service layer chat representation.
type Chat struct {
	UserIDs []int
	ChatID  int
}

// SendMessageRequest is service layer message representation.
type SendMessageRequest struct {
	ChatID    int
	AuthorID  int
	Content   string
	Timestamp time.Time
}

// DeleteChatRequest is service layer message for chat deletion.
type DeleteChatRequest struct {
	ChatID int
}
