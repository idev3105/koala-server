package movierepo

import (
	"context"

	moviedomain "org.idev.koala/backend/domain/movie"
	movieentity "org.idev.koala/backend/domain/movie/entity"
)

type MovieRepo struct {
	sqlRepo *MovieSqlRepo
}

func NewMovieRepo(sqlRepo *MovieSqlRepo) moviedomain.MovieRepository {
	return &MovieRepo{sqlRepo: sqlRepo}
}

func (r *MovieRepo) CreateMovie(ctx context.Context, movie *movieentity.Movie) (*movieentity.Movie, error) {
	return r.sqlRepo.CreateMovie(ctx, movie)
}

func (r *MovieRepo) UpdateMovie(ctx context.Context, movie *movieentity.Movie) (*movieentity.Movie, error) {
	return r.sqlRepo.UpdateMovie(ctx, movie)
}

func (r *MovieRepo) GetMovieById(ctx context.Context, movieId string) (*movieentity.Movie, error) {
	return r.sqlRepo.GetMovieById(ctx, movieId)
}

func (r *MovieRepo) CreateSeason(ctx context.Context, movieId string, season *movieentity.Season) (*movieentity.Season, error) {
	return r.sqlRepo.CreateSeason(ctx, movieId, season)
}

func (r *MovieRepo) UpdateSeason(ctx context.Context, season *movieentity.Season) (*movieentity.Season, error) {
	return r.sqlRepo.UpdateSeason(ctx, season)
}

func (r *MovieRepo) GetSeasonsByMovieId(ctx context.Context, movieId string, seasonId string) ([]movieentity.Season, error) {
	return r.sqlRepo.GetSeasonsByMovieId(ctx, movieId, seasonId)
}

func (r *MovieRepo) CreateEpisode(ctx context.Context, seasonId string, episode *movieentity.Episode) (*movieentity.Episode, error) {
	return r.sqlRepo.CreateEpisode(ctx, seasonId, episode)
}

func (r *MovieRepo) UpdateEpisode(ctx context.Context, episode *movieentity.Episode) (*movieentity.Episode, error) {
	return r.sqlRepo.UpdateEpisode(ctx, episode)
}

func (r *MovieRepo) GetEpisodesBySeasonId(ctx context.Context, seasonId string, episodeId string) ([]movieentity.Episode, error) {
	return r.sqlRepo.GetEpisodesBySeasonId(ctx, seasonId, episodeId)
}
