package storage

import (
	"context"
	"database/sql"

	"github.com/Karanth1r3/ws_chat_learn/internal/models"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
}

type repository struct {
	dbtx DBTX
}

func NewRepository(dbtx DBTX) models.Repository {
	return &repository{dbtx: dbtx}
}

func (r *repository) CreateUser(ctx context.Context, user models.User) (*models.User, error) {

	query := "INSERT INTO "

	r.db.QueryRowContext(ctx, query)
}
