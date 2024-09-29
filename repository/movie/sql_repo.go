package movierepo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	movieentity "org.idev.koala/backend/domain/movie/entity"
	sqlc_generated "org.idev.koala/backend/generated/sqlc"
	"org.idev.koala/backend/mapper"
)

type MovieSqlRepo struct {
	queries *sqlc_generated.Queries
}

func NewMovieSqlRepo(queries *sqlc_generated.Queries) *MovieSqlRepo {
	return &MovieSqlRepo{queries: queries}
}

func (r *MovieSqlRepo) CreateMovie(ctx context.Context, movie *movieentity.Movie) (*movieentity.Movie, error) {
	createdMovie, err := r.queries.CreateMovie(ctx, sqlc_generated.CreateMovieParams{
		MovieID:         movie.MovieID,
		ItemStatus:      movie.ItemStatus,
		Name:            movie.Name,
		Description:     movie.Description,
		ThumbnailUrl:    pgtype.Text{String: movie.ThumbnailUrl, Valid: true},
		AvailableStatus: movie.AvailableStatus.String(),
		Types:           movie.Types,
		ReleaseYear:     movie.ReleaseYear,
		CreatedBy:       pgtype.Text{String: *movie.CreatedBy, Valid: movie.CreatedBy != nil},
		CreatedAt:       pgtype.Timestamp{Time: *movie.CreatedAt, Valid: true},
		UpdatedBy:       pgtype.Text{String: *movie.UpdatedBy, Valid: movie.UpdatedBy != nil},
		UpdatedAt:       pgtype.Timestamp{Time: *movie.UpdatedAt, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return mapper.MapToDomainMovie(createdMovie), nil
}

func (r *MovieSqlRepo) UpdateMovie(ctx context.Context, movie *movieentity.Movie) (*movieentity.Movie, error) {
	updatedMovie, err := r.queries.UpdateMovie(ctx, sqlc_generated.UpdateMovieParams{
		MovieID:         movie.MovieID,
		ItemStatus:      movie.ItemStatus,
		Name:            movie.Name,
		Description:     movie.Description,
		ThumbnailUrl:    pgtype.Text{String: movie.ThumbnailUrl, Valid: movie.ThumbnailUrl != ""},
		AvailableStatus: movie.AvailableStatus.String(),
		Types:           movie.Types,
		ReleaseYear:     movie.ReleaseYear,
		UpdatedBy:       pgtype.Text{String: *movie.UpdatedBy, Valid: movie.UpdatedBy != nil},
		UpdatedAt:       pgtype.Timestamp{Time: *movie.UpdatedAt, Valid: movie.UpdatedAt != nil},
	})
	if err != nil {
		return nil, err
	}
	return mapper.MapToDomainMovie(updatedMovie), nil
}

func (r *MovieSqlRepo) GetMovieById(ctx context.Context, movieId string) (*movieentity.Movie, error) {
	movie, err := r.queries.GetMovieById(ctx, movieId)
	if err != nil {
		return nil, err
	}
	return mapper.MapToDomainMovie(movie), nil
}

func (r *MovieSqlRepo) CreateSeason(ctx context.Context, movieId string, season *movieentity.Season) (*movieentity.Season, error) {
	// TODO: Implement this method
	return nil, nil
}

func (r *MovieSqlRepo) UpdateSeason(ctx context.Context, season *movieentity.Season) (*movieentity.Season, error) {
	// TODO: Implement this method
	return nil, nil
}

func (r *MovieSqlRepo) GetSeasonsByMovieId(ctx context.Context, movieId string, seasonId string) ([]movieentity.Season, error) {
	// TODO: Implement this method
	return nil, nil
}

func (r *MovieSqlRepo) CreateEpisode(ctx context.Context, seasonId string, episode *movieentity.Episode) (*movieentity.Episode, error) {
	// TODO: Implement this method
	return nil, nil
}

func (r *MovieSqlRepo) UpdateEpisode(ctx context.Context, episode *movieentity.Episode) (*movieentity.Episode, error) {
	// TODO: Implement this method
	return nil, nil
}

func (r *MovieSqlRepo) GetEpisodesBySeasonId(ctx context.Context, seasonId string, episodeId string) ([]movieentity.Episode, error) {
	// TODO: Implement this method
	return nil, nil
}
