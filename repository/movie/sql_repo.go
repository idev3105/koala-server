package movierepo

import (
	"context"

	movieentity "org.idev.koala/backend/domain/movie/entity"
	sqlc_generated "org.idev.koala/backend/generated/sqlc"
)

type MovieSqlRepo struct {
	queries *sqlc_generated.Queries
}

func NewMovieSqlRepo(queries *sqlc_generated.Queries) *MovieSqlRepo {
	return &MovieSqlRepo{queries: queries}
}

func (r *MovieSqlRepo) CreateMovie(ctx context.Context, movie *movieentity.Movie) (*movieentity.Movie, error) {
	// TODO: Implement this method
	return nil, nil
}

func (r *MovieSqlRepo) UpdateMovie(ctx context.Context, movie *movieentity.Movie) (*movieentity.Movie, error) {
	// TODO: Implement this method
	return nil, nil
}

func (r *MovieSqlRepo) GetMovieById(ctx context.Context, movieId string) (*movieentity.Movie, error) {
	// TODO: Implement this method
	return nil, nil
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
