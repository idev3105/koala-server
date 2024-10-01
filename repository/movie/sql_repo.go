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
	createdSeason, err := r.queries.CreateSeason(ctx, sqlc_generated.CreateSeasonParams{
		MovieID:           movieId,
		Name:              season.Name,
		Description:       pgtype.Text{String: season.Description, Valid: season.Description != ""},
		ThumbnailUrl:      pgtype.Text{String: season.ThumbnailUrl, Valid: season.ThumbnailUrl != ""},
		ProgressionStatus: pgtype.Text{String: season.ProgressionStatus.String(), Valid: season.ProgressionStatus != ""},
		AvailableStatus:   pgtype.Text{String: season.AvailableStatus.String(), Valid: season.AvailableStatus != ""},
		ReleaseYear:       pgtype.Int4{Int32: season.ReleaseYear, Valid: season.ReleaseYear != 0},
		Order:             pgtype.Int4{Int32: season.Order, Valid: season.Order != 0},
		TotalEpisode:      pgtype.Int4{Int32: season.TotalEpisode, Valid: season.TotalEpisode != 0},
		CreatedBy:         pgtype.Text{String: *season.CreatedBy, Valid: season.CreatedBy != nil},
		CreatedAt:         pgtype.Timestamp{Time: *season.CreatedAt, Valid: season.CreatedAt != nil},
	})
	if err != nil {
		return nil, err
	}
	return mapper.MapToDomainSeason(createdSeason), nil
}

func (r *MovieSqlRepo) UpdateSeason(ctx context.Context, season *movieentity.Season) (*movieentity.Season, error) {
	updatedSeason, err := r.queries.UpdateSeason(ctx, sqlc_generated.UpdateSeasonParams{
		ID:                season.Id,
		Name:              season.Name,
		Description:       pgtype.Text{String: season.Description, Valid: season.Description != ""},
		ThumbnailUrl:      pgtype.Text{String: season.ThumbnailUrl, Valid: season.ThumbnailUrl != ""},
		ProgressionStatus: pgtype.Text{String: season.ProgressionStatus.String(), Valid: season.ProgressionStatus != ""},
		AvailableStatus:   pgtype.Text{String: season.AvailableStatus.String(), Valid: season.AvailableStatus != ""},
		ReleaseYear:       pgtype.Int4{Int32: season.ReleaseYear, Valid: season.ReleaseYear != 0},
		Order:             pgtype.Int4{Int32: season.Order, Valid: season.Order != 0},
		TotalEpisode:      pgtype.Int4{Int32: season.TotalEpisode, Valid: season.TotalEpisode != 0},
		UpdatedBy:         pgtype.Text{String: *season.UpdatedBy, Valid: season.UpdatedBy != nil},
		UpdatedAt:         pgtype.Timestamp{Time: *season.UpdatedAt, Valid: season.UpdatedAt != nil},
	})
	if err != nil {
		return nil, err
	}
	return mapper.MapToDomainSeason(updatedSeason), nil
}

func (r *MovieSqlRepo) GetSeasonsByMovieId(ctx context.Context, movieId string) ([]movieentity.Season, error) {
	seasons, err := r.queries.GetSeasonsByMovieId(ctx, sqlc_generated.GetSeasonsByMovieIdParams{
		MovieID: movieId,
	})
	if err != nil {
		return nil, err
	}
	seasonsDomain := make([]movieentity.Season, 0)
	for _, season := range seasons {
		seasonsDomain = append(seasonsDomain, *mapper.MapToDomainSeason(season))
	}
	return seasonsDomain, nil
}

func (r *MovieSqlRepo) CreateEpisode(ctx context.Context, seasonId string, episode *movieentity.Episode) (*movieentity.Episode, error) {
	createdEpisode, err := r.queries.CreateEpisode(ctx, sqlc_generated.CreateEpisodeParams{
		SeasonID:         seasonId,
		Name:             episode.Name,
		OriginalVideoUrl: pgtype.Text{String: episode.OriginalVideoUrl, Valid: episode.OriginalVideoUrl != ""},
		CreatedBy:        pgtype.Text{String: *episode.CreatedBy, Valid: episode.CreatedBy != nil},
		CreatedAt:        pgtype.Timestamp{Time: *episode.CreatedAt, Valid: episode.CreatedAt != nil},
	})
	if err != nil {
		return nil, err
	}
	return mapper.MapToDomainEpisode(createdEpisode), nil
}

func (r *MovieSqlRepo) UpdateEpisode(ctx context.Context, episode *movieentity.Episode) (*movieentity.Episode, error) {
	updatedEpisode, err := r.queries.UpdateEpisode(ctx, sqlc_generated.UpdateEpisodeParams{
		ID:               episode.Id,
		Name:             episode.Name,
		OriginalVideoUrl: pgtype.Text{String: episode.OriginalVideoUrl, Valid: episode.OriginalVideoUrl != ""},
		StreamVideoUrl:   pgtype.Text{String: episode.StreamVideoUrl, Valid: episode.StreamVideoUrl != ""},
	})
	if err != nil {
		return nil, err
	}
	return mapper.MapToDomainEpisode(updatedEpisode), nil
}

func (r *MovieSqlRepo) GetEpisodesBySeasonId(ctx context.Context, seasonId string) ([]movieentity.Episode, error) {
	episodes, err := r.queries.GetEpisodesBySeasonId(ctx, sqlc_generated.GetEpisodesBySeasonIdParams{
		SeasonID: seasonId,
	})
	if err != nil {
		return nil, err
	}
	episodesDomain := make([]movieentity.Episode, 0)
	for _, episode := range episodes {
		episodesDomain = append(episodesDomain, *mapper.MapToDomainEpisode(episode))
	}
	return episodesDomain, nil
}
