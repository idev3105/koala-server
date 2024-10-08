package userrepository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	userentity "org.idev.koala/backend/domain/user/entity"
	sqlc_generated "org.idev.koala/backend/generated/sqlc"
	"org.idev.koala/backend/mapper"
)

type UserSqlRepo struct {
	db      sqlc_generated.DBTX
	queries *sqlc_generated.Queries
}

// create new instance of sql repository
func NewSqlRepository(queries *sqlc_generated.Queries) *UserSqlRepo {
	return &UserSqlRepo{queries: queries}
}

func (r *UserSqlRepo) ExistsByUserId(ctx context.Context, userId string) (bool, error) {
	return r.queries.ExistsUserByUserId(ctx, userId)
}

func (r *UserSqlRepo) FindByUserId(ctx context.Context, userId string) (*userentity.User, error) {
	user, err := r.queries.FindUserByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return mapper.MapToDomain(user), nil
}

func (r *UserSqlRepo) Save(ctx context.Context, user *userentity.User) (*userentity.User, error) {
	prams := sqlc_generated.SaveUserParams{
		UserId:   user.UserId,
		Username: pgtype.Text{String: user.Username, Valid: true},
	}
	if user.CreatedBy != nil {
		prams.CreatedBy = pgtype.Text{String: *user.CreatedBy, Valid: true}
	}
	if user.UpdatedBy != nil {
		prams.UpdatedBy = pgtype.Text{String: *user.UpdatedBy, Valid: true}
	}
	_, err := r.queries.SaveUser(ctx, prams)
	if err != nil {
		return nil, err
	}
	return user, nil
}
