package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/lookandhate/course_chat/internal/repository/model"
	"github.com/lookandhate/course_platform_lib/pkg/db"
)

type PostgresRepository struct {
	db db.Client
}

const (
	chatTable       = "chats"
	chatMemberTable = "chat_members"
	messageTable    = "message"

	idColumn        = "id"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"

	userIDColumn   = "user_id"
	chatIDColumn   = "chat_id"
	authorIDColumn = "author_id"
	contentColumn  = "content"
)

// NewPostgresRepository creates PostgresRepository instance.
func NewPostgresRepository(db db.Client) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// CreateChat creates chat.go with chat.go members.
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

	query := db.Query{
		Name:     "PostgresRepository.CreateChat",
		QueryRaw: sql,
	}
	var chatModel model.ChatModel

	err = r.db.DB().ScanOneContext(ctx, &chatModel, query, args...)
	if err != nil {
		return nil, err
	}

	chatModel.UserIDs = request.UserIDs

	err = r.addUsersToChat(ctx, chatModel.ID, request.UserIDs)

	return &chatModel, err
}

// addUsersToChat creates chat.go members with given chatID and userIDSs.
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

	query := db.Query{
		Name:     "PostgresRepository.addUsersToChat",
		QueryRaw: sql,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)

	return err
}

// CreateMessage creates message.
func (r *PostgresRepository) CreateMessage(
	ctx context.Context,
	message *model.CreateMessageModel,
) (*model.MessageModel, error) {
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

	query := db.Query{
		Name:     "PostgresRepository.CreateMessage",
		QueryRaw: sql,
	}

	var createdMessage model.MessageModel
	err = r.db.DB().ScanOneContext(ctx, &createdMessage, query, args...)
	return &createdMessage, err
}

// Delete deletes chat.go.
func (r *PostgresRepository) Delete(ctx context.Context, id int64) error {
	builder := squirrel.Delete(chatTable).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idColumn: id})

	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "PostgresRepository.DeleteChat",
		QueryRaw: sql,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)

	return err
}

// ChatExists checks if chat.go exists.
func (r *PostgresRepository) ChatExists(ctx context.Context, chatID int) (bool, error) {
	var exists bool
	// using Prefix and suffix for EXIST query
	builder := squirrel.Select(
		fmt.Sprintf(
			"EXISTS(SELECT 1 FROM %s WHERE id = %s) AS chat_exists",
			chatTable,
			strconv.Itoa(chatID),
		),
	)

	sql, args, err := builder.ToSql()
	if err != nil {
		return false, err
	}

	query := db.Query{
		Name:     "PostgresRepository.ChatExists",
		QueryRaw: sql,
	}

	err = r.db.DB().ScanOneContext(ctx, &exists, query, args...)

	return exists, err
}
