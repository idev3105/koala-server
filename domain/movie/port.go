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
	GetSeasonsByMovieId(ctx context.Context, movieId string, seasonId string) ([]movieentity.Season, error)

	CreateEpisode(ctx context.Context, seasonId string, episode *movieentity.Episode) (*movieentity.Episode, error)
	UpdateEpisode(ctx context.Context, episode *movieentity.Episode) (*movieentity.Episode, error)
	GetEpisodesBySeasonId(ctx context.Context, seasonId string, episodeId string) ([]movieentity.Episode, error)
}

type MovieUseCase interface {
	CreateMovie(ctx context.Context, movie *movieentity.Movie) (*movieentity.Movie, error)
	UpdateMovie(ctx context.Context, movie *movieentity.Movie) (*movieentity.Movie, error)
	GetMovieById(ctx context.Context, movieId string) (*movieentity.Movie, error)
	UpvoteMovie(ctx context.Context, movieId string, userId string) (*movieentity.Movie, error)
	DownvoteMovie(ctx context.Context, movieId string, userId string) (*movieentity.Movie, error)
	RateMovie(ctx context.Context, movieId string, userId string, rate float32) (*movieentity.Movie, error)

	CreateSeason(ctx context.Context, movieId string, season *movieentity.Season) (*movieentity.Season, error)
	UpdateSeason(ctx context.Context, season *movieentity.Season) (*movieentity.Season, error)
	GetSeasonsByMovieId(ctx context.Context, movieId string, seasonId string) ([]movieentity.Season, error)
	UpvoteSeason(ctx context.Context, seasonId string, userId string) (*movieentity.Season, error)
	DownvoteSeason(ctx context.Context, seasonId string, userId string) (*movieentity.Season, error)
	RateSeason(ctx context.Context, seasonId string, userId string, rate float32) (*movieentity.Season, error)

	CreateEpisode(ctx context.Context, seasonId string, episode *movieentity.Episode) (*movieentity.Episode, error)
	UpdateEpisode(ctx context.Context, episode *movieentity.Episode) (*movieentity.Episode, error)
	GetEpisodesBySeasonId(ctx context.Context, seasonId string, episodeId string) ([]movieentity.Episode, error)
	UpvoteEpisode(ctx context.Context, episodeId string, userId string) (*movieentity.Episode, error)
	DownvoteEpisode(ctx context.Context, episodeId string, userId string) (*movieentity.Episode, error)
	RateEpisode(ctx context.Context, episodeId string, userId string, rate float32) (*movieentity.Episode, error)
}
