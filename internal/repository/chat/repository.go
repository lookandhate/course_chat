package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lookandhate/microservice-courese/chat/internal/repository/model"
)

type PostgresRepository struct {
	pgx *pgxpool.Pool
}

const (
	chatTable       = "chats"
	chatMemberTable = "chat_members"
	messageTable    = "messages"

	idColumn        = "id"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"

	userIDColumn   = "user_id"
	chatIDColumn   = "chat_id"
	authorIDColumn = "author_id"
	contentColumn  = "content"
)

// NewPostgresRepository creates PostgresRepository instance.
func NewPostgresRepository(connectionPool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pgx: connectionPool}
}

// CreateChat creates chat with chat members.
func (r *PostgresRepository) CreateChat(ctx context.Context, request *model.CreateChatModel) (*model.ChatModel, error) {
	builder := squirrel.Insert(chatTable).
		PlaceholderFormat(squirrel.Dollar).
		Columns(createdAtColumn).
		Values(time.Now()).
		Suffix(fmt.Sprintf("returning %s, %s, %s", idColumn, createdAtColumn, updatedAtColumn))

	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var chatModel model.ChatModel
	err = r.pgx.QueryRow(ctx, sql, args...).Scan(&chatModel.ID, &chatModel.CreatedAt, &chatModel.UpdatedAt)
	chatModel.UserIDs = request.UserIDs
	if err != nil {
		return nil, err
	}
	err = r.addUsersToChat(ctx, chatModel.ID, request.UserIDs)

	return &chatModel, err
}

// addUsersToChat creates chat members with given chatID and userIDSs.
func (r *PostgresRepository) addUsersToChat(ctx context.Context, chatID int, userIDs []int64) error {
	builder := squirrel.
		Insert(chatMemberTable).
		PlaceholderFormat(squirrel.Dollar).
		Columns(userIDColumn, chatIDColumn)

	for _, userID := range userIDs {
		builder = builder.Values(userID, chatID)
	}

	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pgx.Exec(ctx, sql, args...)

	return err
}

// CreateMessage creates message.
func (r *PostgresRepository) CreateMessage(ctx context.Context, message *model.CreateMessageModel) (*model.MessageModel, error) {
	builder := squirrel.Insert(messageTable).
		PlaceholderFormat(squirrel.Dollar).
		Columns(authorIDColumn, contentColumn, chatIDColumn).
		Values(message.AuthorID, message.Content, message.ChatID).
		Suffix(fmt.Sprintf("returning %s, %s, %s, %s, %s, %s",
			idColumn, authorIDColumn, contentColumn, chatIDColumn, createdAtColumn, updatedAtColumn))

	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var createdMessage model.MessageModel
	err = r.pgx.QueryRow(ctx, sql, args...).Scan(
		&createdMessage.ID,
		&createdMessage.Author,
		&createdMessage.Content,
		&createdMessage.ChatID,
		&createdMessage.CreatedAt,
		&createdMessage.UpdatedAt)

	return &createdMessage, err
}

// Delete deletes chat.
func (r *PostgresRepository) Delete(ctx context.Context, id int64) error {
	builder := squirrel.Delete(chatTable).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idColumn: id})

	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	_, err = r.pgx.Exec(ctx, sql, args...)
	return err
}

// ChatExists checks if chat exists.
func (r *PostgresRepository) ChatExists(ctx context.Context, chatID int) (bool, error) {
	var exists bool
	// using Prefix and suffix for EXIST query
	builder := squirrel.Select(fmt.Sprintf("EXISTS(SELECT 1 FROM %s WHERE id = %s) AS chat_exists", chatTable, strconv.Itoa(chatID)))

	sql, args, err := builder.ToSql()

	if err != nil {
		return false, err
	}
	err = r.pgx.QueryRow(ctx, sql, args...).Scan(&exists)
	return exists, err
}
