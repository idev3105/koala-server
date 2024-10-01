package moviedomain

import (
	"context"

	movieentity "org.idev.koala/backend/domain/movie/entity"
)

type MovieRepository interface {
	CreateMovie(ctx context.Context, movie *movieentity.Movie) (*movieentity.Movie, error)
	UpdateMovie(ctx context.Context, movie *movieentity.Movie) (*movieentity.Movie, error)
	GetMovieById(ctx context.Context, movieId string) (*movieentity.Movie, error)

	CreateSeason(ctx context.Context, movieId string, season *movieentity.Season) (*movieentity.Season, error)
	UpdateSeason(ctx context.Context, season *movieentity.Season) (*movieentity.Season, error)
	GetSeasonsByMovieId(ctx context.Context, movieId string) ([]movieentity.Season, error)

	CreateEpisode(ctx context.Context, seasonId string, episode *movieentity.Episode) (*movieentity.Episode, error)
	UpdateEpisode(ctx context.Context, episode *movieentity.Episode) (*movieentity.Episode, error)
	GetEpisodesBySeasonId(ctx context.Context, seasonId string) ([]movieentity.Episode, error)
}

type MovieUseCase interface {
	CreateMovie(ctx context.Context, movie *movieentity.Movie) (*movieentity.Movie, error)
	UpdateMovie(ctx context.Context, movie *movieentity.Movie) (*movieentity.Movie, error)
	GetMovieById(ctx context.Context, movieId string) (*movieentity.Movie, error)
	VoteMovie(ctx context.Context, movieId, userId string, isUpvote bool) (*movieentity.Movie, error)
	RateMovie(ctx context.Context, movieId, userId string, rate float32) (*movieentity.Movie, error)

	CreateSeason(ctx context.Context, movieId string, season *movieentity.Season) (*movieentity.Season, error)
	UpdateSeason(ctx context.Context, season *movieentity.Season) (*movieentity.Season, error)
	GetSeasonsByMovieId(ctx context.Context, movieId string) ([]movieentity.Season, error)
	VoteSeason(ctx context.Context, seasonId, userId string, isUpvote bool) (*movieentity.Season, error)
	RateSeason(ctx context.Context, seasonId, userId string, rate float32) (*movieentity.Season, error)

	CreateEpisode(ctx context.Context, seasonId string, episode *movieentity.Episode) (*movieentity.Episode, error)
	UpdateEpisode(ctx context.Context, episode *movieentity.Episode) (*movieentity.Episode, error)
	GetEpisodesBySeasonId(ctx context.Context, seasonId string) ([]movieentity.Episode, error)
	VoteEpisode(ctx context.Context, episodeId, userId string, isUpvote bool) (*movieentity.Episode, error)
	RateEpisode(ctx context.Context, episodeId, userId string, rate float32) (*movieentity.Episode, error)
}
