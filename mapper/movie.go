package mapper

import (
	"github.com/jackc/pgx/v5/pgtype"
	movieentity "org.idev.koala/backend/domain/movie/entity"
	sqlc_generated "org.idev.koala/backend/generated/sqlc"
)

func MapToDomainMovie(m sqlc_generated.Movie) *movieentity.Movie {
	// TODO: Implement this function
	movie := &movieentity.Movie{}
	movie.MovieID = m.MovieID
	movie.ItemStatus = m.ItemStatus
	movie.Description = m.Description
	movie.ThumbnailUrl = m.ThumbnailUrl.String
	return movie
}

func MapToSQLModelMovie(m movieentity.Movie) *sqlc_generated.Movie {
	// TODO: Implement this function
	movie := &sqlc_generated.Movie{}
	movie.MovieID = m.MovieID
	movie.ItemStatus = m.ItemStatus
	movie.Description = m.Description
	movie.ThumbnailUrl = pgtype.Text{String: m.ThumbnailUrl, Valid: true}
	return movie
}
