package repository

import (
	"context"
	"github.com/KRUL-marketplace/common-libs/pkg/client/db"
	sq "github.com/Masterminds/squirrel"
	"user-service/internal/repository/model"
)

func (r *repo) Create(ctx context.Context, info *model.UserInfo) (string, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(phoneNumberColumn, nameColumn).
		Values(info.PhoneNumber, info.Name).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id string
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}
