package moviedomain

import (
	"context"
	"errors"

	movieentity "org.idev.koala/backend/domain/movie/entity"
)

type movieUseCase struct {
	repo MovieRepository
}

func NewMovieUseCase(repo MovieRepository) MovieUseCase {
	return &movieUseCase{repo: repo}
}

// Movie methods
func (uc *movieUseCase) CreateMovie(ctx context.Context, movie *movieentity.Movie) (*movieentity.Movie, error) {
	return uc.repo.CreateMovie(ctx, movie)
}

func (uc *movieUseCase) UpdateMovie(ctx context.Context, movie *movieentity.Movie) (*movieentity.Movie, error) {
	return uc.repo.UpdateMovie(ctx, movie)
}

func (uc *movieUseCase) GetMovieById(ctx context.Context, movieId string) (*movieentity.Movie, error) {
	return uc.repo.GetMovieById(ctx, movieId)
}

func (uc *movieUseCase) VoteMovie(ctx context.Context, movieId, userId string, isUpvote bool) (*movieentity.Movie, error) {
	// Implement voting logic here
	// This might involve fetching the movie, updating its vote count, and saving it
	// For now, we'll return nil and an error as it's not implemented in the repository
	return nil, errors.New("not implemented")
}

func (uc *movieUseCase) RateMovie(ctx context.Context, movieId, userId string, rate float32) (*movieentity.Movie, error) {
	// Implement rating logic here
	// This might involve fetching the movie, updating its rating, and saving it
	// For now, we'll return nil and an error as it's not implemented in the repository
	return nil, errors.New("not implemented")
}

// Season methods
func (uc *movieUseCase) CreateSeason(ctx context.Context, movieId string, season *movieentity.Season) (*movieentity.Season, error) {
	return uc.repo.CreateSeason(ctx, movieId, season)
}

func (uc *movieUseCase) UpdateSeason(ctx context.Context, season *movieentity.Season) (*movieentity.Season, error) {
	return uc.repo.UpdateSeason(ctx, season)
}

func (uc *movieUseCase) GetSeasonsByMovieId(ctx context.Context, movieId string) ([]movieentity.Season, error) {
	return uc.repo.GetSeasonsByMovieId(ctx, movieId)
}

func (uc *movieUseCase) VoteSeason(ctx context.Context, seasonId, userId string, isUpvote bool) (*movieentity.Season, error) {
	// Implement voting logic for seasons
	return nil, errors.New("not implemented")
}

func (uc *movieUseCase) RateSeason(ctx context.Context, seasonId, userId string, rate float32) (*movieentity.Season, error) {
	// Implement rating logic for seasons
	return nil, errors.New("not implemented")
}

// Episode methods
func (uc *movieUseCase) CreateEpisode(ctx context.Context, seasonId string, episode *movieentity.Episode) (*movieentity.Episode, error) {
	return uc.repo.CreateEpisode(ctx, seasonId, episode)
}

func (uc *movieUseCase) UpdateEpisode(ctx context.Context, episode *movieentity.Episode) (*movieentity.Episode, error) {
	return uc.repo.UpdateEpisode(ctx, episode)
}

func (uc *movieUseCase) GetEpisodesBySeasonId(ctx context.Context, seasonId string) ([]movieentity.Episode, error) {
	return uc.repo.GetEpisodesBySeasonId(ctx, seasonId)
}

func (uc *movieUseCase) VoteEpisode(ctx context.Context, episodeId, userId string, isUpvote bool) (*movieentity.Episode, error) {
	// Implement voting logic for episodes
	return nil, errors.New("not implemented")
}

func (uc *movieUseCase) RateEpisode(ctx context.Context, episodeId, userId string, rate float32) (*movieentity.Episode, error) {
	// Implement rating logic for episodes
	return nil, errors.New("not implemented")
}
