// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: movie.sql

package sqlc_generated

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createEpisode = `-- name: CreateEpisode :one
INSERT INTO d_episodes (
        season_id,
        name,
        description,
        thumbnail_url,
        original_video_url,
        stream_video_url,
        available_status,
        episode_order,
        duration,
        release_date,
        created_by,
        created_at,
        updated_by,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14
    )
RETURNING id, season_id, name, description, thumbnail_url, original_video_url, stream_video_url, available_status, episode_order, rate, total_rate, total_viewer, total_upvote, total_downvote, duration, release_date, created_by, created_at, updated_by, updated_at
`

type CreateEpisodeParams struct {
	SeasonID         string
	Name             string
	Description      pgtype.Text
	ThumbnailUrl     pgtype.Text
	OriginalVideoUrl pgtype.Text
	StreamVideoUrl   pgtype.Text
	AvailableStatus  pgtype.Text
	EpisodeOrder     pgtype.Int4
	Duration         pgtype.Int4
	ReleaseDate      pgtype.Date
	CreatedBy        pgtype.Text
	CreatedAt        pgtype.Timestamp
	UpdatedBy        pgtype.Text
	UpdatedAt        pgtype.Timestamp
}

func (q *Queries) CreateEpisode(ctx context.Context, arg CreateEpisodeParams) (DEpisode, error) {
	row := q.db.QueryRow(ctx, createEpisode,
		arg.SeasonID,
		arg.Name,
		arg.Description,
		arg.ThumbnailUrl,
		arg.OriginalVideoUrl,
		arg.StreamVideoUrl,
		arg.AvailableStatus,
		arg.EpisodeOrder,
		arg.Duration,
		arg.ReleaseDate,
		arg.CreatedBy,
		arg.CreatedAt,
		arg.UpdatedBy,
		arg.UpdatedAt,
	)
	var i DEpisode
	err := row.Scan(
		&i.ID,
		&i.SeasonID,
		&i.Name,
		&i.Description,
		&i.ThumbnailUrl,
		&i.OriginalVideoUrl,
		&i.StreamVideoUrl,
		&i.AvailableStatus,
		&i.EpisodeOrder,
		&i.Rate,
		&i.TotalRate,
		&i.TotalViewer,
		&i.TotalUpvote,
		&i.TotalDownvote,
		&i.Duration,
		&i.ReleaseDate,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedBy,
		&i.UpdatedAt,
	)
	return i, err
}

const createMovie = `-- name: CreateMovie :one
INSERT INTO d_movies (
        movie_id,
        item_status,
        title,
        description,
        thumbnail_url,
        available_status,
        types,
        release_year,
        created_by,
        created_at,
        updated_by,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12
    )
RETURNING id, movie_id, item_status, title, description, thumbnail_url, available_status, types, release_year, total_viewer, total_upvote, total_downvote, total_rate, rate, created_at, updated_at, created_by, updated_by
`

type CreateMovieParams struct {
	MovieID         string
	ItemStatus      string
	Title           string
	Description     string
	ThumbnailUrl    pgtype.Text
	AvailableStatus string
	Types           []string
	ReleaseYear     int32
	CreatedBy       pgtype.Text
	CreatedAt       pgtype.Timestamp
	UpdatedBy       pgtype.Text
	UpdatedAt       pgtype.Timestamp
}

func (q *Queries) CreateMovie(ctx context.Context, arg CreateMovieParams) (Movie, error) {
	row := q.db.QueryRow(ctx, createMovie,
		arg.MovieID,
		arg.ItemStatus,
		arg.Title,
		arg.Description,
		arg.ThumbnailUrl,
		arg.AvailableStatus,
		arg.Types,
		arg.ReleaseYear,
		arg.CreatedBy,
		arg.CreatedAt,
		arg.UpdatedBy,
		arg.UpdatedAt,
	)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.MovieID,
		&i.ItemStatus,
		&i.Title,
		&i.Description,
		&i.ThumbnailUrl,
		&i.AvailableStatus,
		&i.Types,
		&i.ReleaseYear,
		&i.TotalViewer,
		&i.TotalUpvote,
		&i.TotalDownvote,
		&i.TotalRate,
		&i.Rate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const createSeason = `-- name: CreateSeason :one
INSERT INTO d_seasons (
        movie_id,
        name,
        description,
        thumbnail_url,
        progression_status,
        available_status,
        release_year,
        "order",
        total_episode,
        created_by,
        created_at,
        updated_by,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13
    )
RETURNING id, movie_id, name, description, thumbnail_url, progression_status, available_status, release_year, rate, total_rate, total_viewer, total_upvote, total_downvote, "order", total_episode, created_by, created_at, updated_by, updated_at
`

type CreateSeasonParams struct {
	MovieID           string
	Name              string
	Description       pgtype.Text
	ThumbnailUrl      pgtype.Text
	ProgressionStatus pgtype.Text
	AvailableStatus   pgtype.Text
	ReleaseYear       pgtype.Int4
	Order             pgtype.Int4
	TotalEpisode      pgtype.Int4
	CreatedBy         pgtype.Text
	CreatedAt         pgtype.Timestamp
	UpdatedBy         pgtype.Text
	UpdatedAt         pgtype.Timestamp
}

func (q *Queries) CreateSeason(ctx context.Context, arg CreateSeasonParams) (DSeason, error) {
	row := q.db.QueryRow(ctx, createSeason,
		arg.MovieID,
		arg.Name,
		arg.Description,
		arg.ThumbnailUrl,
		arg.ProgressionStatus,
		arg.AvailableStatus,
		arg.ReleaseYear,
		arg.Order,
		arg.TotalEpisode,
		arg.CreatedBy,
		arg.CreatedAt,
		arg.UpdatedBy,
		arg.UpdatedAt,
	)
	var i DSeason
	err := row.Scan(
		&i.ID,
		&i.MovieID,
		&i.Name,
		&i.Description,
		&i.ThumbnailUrl,
		&i.ProgressionStatus,
		&i.AvailableStatus,
		&i.ReleaseYear,
		&i.Rate,
		&i.TotalRate,
		&i.TotalViewer,
		&i.TotalUpvote,
		&i.TotalDownvote,
		&i.Order,
		&i.TotalEpisode,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedBy,
		&i.UpdatedAt,
	)
	return i, err
}

const getEpisodesBySeasonId = `-- name: GetEpisodesBySeasonId :many
SELECT id, season_id, name, description, thumbnail_url, original_video_url, stream_video_url, available_status, episode_order, rate, total_rate, total_viewer, total_upvote, total_downvote, duration, release_date, created_by, created_at, updated_by, updated_at
FROM d_episodes
WHERE season_id = $1
    AND (
        $2::VARCHAR IS NULL
        OR id::VARCHAR = $2
    )
ORDER BY episode_order
`

type GetEpisodesBySeasonIdParams struct {
	SeasonID string
	Column2  string
}

func (q *Queries) GetEpisodesBySeasonId(ctx context.Context, arg GetEpisodesBySeasonIdParams) ([]DEpisode, error) {
	rows, err := q.db.Query(ctx, getEpisodesBySeasonId, arg.SeasonID, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DEpisode
	for rows.Next() {
		var i DEpisode
		if err := rows.Scan(
			&i.ID,
			&i.SeasonID,
			&i.Name,
			&i.Description,
			&i.ThumbnailUrl,
			&i.OriginalVideoUrl,
			&i.StreamVideoUrl,
			&i.AvailableStatus,
			&i.EpisodeOrder,
			&i.Rate,
			&i.TotalRate,
			&i.TotalViewer,
			&i.TotalUpvote,
			&i.TotalDownvote,
			&i.Duration,
			&i.ReleaseDate,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.UpdatedBy,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMovieById = `-- name: GetMovieById :one
SELECT id, movie_id, item_status, title, description, thumbnail_url, available_status, types, release_year, total_viewer, total_upvote, total_downvote, total_rate, rate, created_at, updated_at, created_by, updated_by
FROM d_movies
WHERE movie_id = $1
`

func (q *Queries) GetMovieById(ctx context.Context, movieID string) (Movie, error) {
	row := q.db.QueryRow(ctx, getMovieById, movieID)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.MovieID,
		&i.ItemStatus,
		&i.Title,
		&i.Description,
		&i.ThumbnailUrl,
		&i.AvailableStatus,
		&i.Types,
		&i.ReleaseYear,
		&i.TotalViewer,
		&i.TotalUpvote,
		&i.TotalDownvote,
		&i.TotalRate,
		&i.Rate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const getSeasonsByMovieId = `-- name: GetSeasonsByMovieId :many
SELECT id, movie_id, name, description, thumbnail_url, progression_status, available_status, release_year, rate, total_rate, total_viewer, total_upvote, total_downvote, "order", total_episode, created_by, created_at, updated_by, updated_at
FROM d_seasons
WHERE movie_id = $1
    AND (
        $2::VARCHAR IS NULL
        OR id::VARCHAR = $2
    )
ORDER BY "order"
`

type GetSeasonsByMovieIdParams struct {
	MovieID string
	Column2 string
}

func (q *Queries) GetSeasonsByMovieId(ctx context.Context, arg GetSeasonsByMovieIdParams) ([]DSeason, error) {
	rows, err := q.db.Query(ctx, getSeasonsByMovieId, arg.MovieID, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DSeason
	for rows.Next() {
		var i DSeason
		if err := rows.Scan(
			&i.ID,
			&i.MovieID,
			&i.Name,
			&i.Description,
			&i.ThumbnailUrl,
			&i.ProgressionStatus,
			&i.AvailableStatus,
			&i.ReleaseYear,
			&i.Rate,
			&i.TotalRate,
			&i.TotalViewer,
			&i.TotalUpvote,
			&i.TotalDownvote,
			&i.Order,
			&i.TotalEpisode,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.UpdatedBy,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEpisode = `-- name: UpdateEpisode :one
UPDATE d_episodes
SET name = COALESCE($2, name),
    description = COALESCE($3, description),
    thumbnail_url = COALESCE($4, thumbnail_url),
    original_video_url = COALESCE($5, original_video_url),
    stream_video_url = COALESCE($6, stream_video_url),
    available_status = COALESCE($7, available_status),
    episode_order = COALESCE($8, episode_order),
    duration = COALESCE($9, duration),
    release_date = COALESCE($10, release_date),
    updated_by = COALESCE($11, updated_by),
    updated_at = COALESCE($12, updated_at)
WHERE id = $1
RETURNING id, season_id, name, description, thumbnail_url, original_video_url, stream_video_url, available_status, episode_order, rate, total_rate, total_viewer, total_upvote, total_downvote, duration, release_date, created_by, created_at, updated_by, updated_at
`

type UpdateEpisodeParams struct {
	ID               int32
	Name             string
	Description      pgtype.Text
	ThumbnailUrl     pgtype.Text
	OriginalVideoUrl pgtype.Text
	StreamVideoUrl   pgtype.Text
	AvailableStatus  pgtype.Text
	EpisodeOrder     pgtype.Int4
	Duration         pgtype.Int4
	ReleaseDate      pgtype.Date
	UpdatedBy        pgtype.Text
	UpdatedAt        pgtype.Timestamp
}

func (q *Queries) UpdateEpisode(ctx context.Context, arg UpdateEpisodeParams) (DEpisode, error) {
	row := q.db.QueryRow(ctx, updateEpisode,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.ThumbnailUrl,
		arg.OriginalVideoUrl,
		arg.StreamVideoUrl,
		arg.AvailableStatus,
		arg.EpisodeOrder,
		arg.Duration,
		arg.ReleaseDate,
		arg.UpdatedBy,
		arg.UpdatedAt,
	)
	var i DEpisode
	err := row.Scan(
		&i.ID,
		&i.SeasonID,
		&i.Name,
		&i.Description,
		&i.ThumbnailUrl,
		&i.OriginalVideoUrl,
		&i.StreamVideoUrl,
		&i.AvailableStatus,
		&i.EpisodeOrder,
		&i.Rate,
		&i.TotalRate,
		&i.TotalViewer,
		&i.TotalUpvote,
		&i.TotalDownvote,
		&i.Duration,
		&i.ReleaseDate,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedBy,
		&i.UpdatedAt,
	)
	return i, err
}

const updateMovie = `-- name: UpdateMovie :one
UPDATE d_movies
SET item_status = COALESCE($2, item_status),
    title = COALESCE($3, title),
    description = COALESCE($4, description),
    thumbnail_url = COALESCE($5, thumbnail_url),
    available_status = COALESCE($6, available_status),
    types = COALESCE($7, types),
    release_year = COALESCE($8, release_year),
    updated_by = COALESCE($9, updated_by),
    updated_at = COALESCE($10, updated_at)
WHERE movie_id = $1
RETURNING id, movie_id, item_status, title, description, thumbnail_url, available_status, types, release_year, total_viewer, total_upvote, total_downvote, total_rate, rate, created_at, updated_at, created_by, updated_by
`

type UpdateMovieParams struct {
	MovieID         string
	ItemStatus      string
	Title           string
	Description     string
	ThumbnailUrl    pgtype.Text
	AvailableStatus string
	Types           []string
	ReleaseYear     int32
	UpdatedBy       pgtype.Text
	UpdatedAt       pgtype.Timestamp
}

func (q *Queries) UpdateMovie(ctx context.Context, arg UpdateMovieParams) (Movie, error) {
	row := q.db.QueryRow(ctx, updateMovie,
		arg.MovieID,
		arg.ItemStatus,
		arg.Title,
		arg.Description,
		arg.ThumbnailUrl,
		arg.AvailableStatus,
		arg.Types,
		arg.ReleaseYear,
		arg.UpdatedBy,
		arg.UpdatedAt,
	)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.MovieID,
		&i.ItemStatus,
		&i.Title,
		&i.Description,
		&i.ThumbnailUrl,
		&i.AvailableStatus,
		&i.Types,
		&i.ReleaseYear,
		&i.TotalViewer,
		&i.TotalUpvote,
		&i.TotalDownvote,
		&i.TotalRate,
		&i.Rate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const updateSeason = `-- name: UpdateSeason :one
UPDATE d_seasons
SET name = COALESCE($2, name),
    description = COALESCE($3, description),
    thumbnail_url = COALESCE($4, thumbnail_url),
    progression_status = COALESCE($5, progression_status),
    available_status = COALESCE($6, available_status),
    release_year = COALESCE($7, release_year),
    "order" = COALESCE($8, "order"),
    total_episode = COALESCE($9, total_episode),
    updated_by = COALESCE($10, updated_by),
    updated_at = COALESCE($11, updated_at)
WHERE id = $1
RETURNING id, movie_id, name, description, thumbnail_url, progression_status, available_status, release_year, rate, total_rate, total_viewer, total_upvote, total_downvote, "order", total_episode, created_by, created_at, updated_by, updated_at
`

type UpdateSeasonParams struct {
	ID                int32
	Name              string
	Description       pgtype.Text
	ThumbnailUrl      pgtype.Text
	ProgressionStatus pgtype.Text
	AvailableStatus   pgtype.Text
	ReleaseYear       pgtype.Int4
	Order             pgtype.Int4
	TotalEpisode      pgtype.Int4
	UpdatedBy         pgtype.Text
	UpdatedAt         pgtype.Timestamp
}

func (q *Queries) UpdateSeason(ctx context.Context, arg UpdateSeasonParams) (DSeason, error) {
	row := q.db.QueryRow(ctx, updateSeason,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.ThumbnailUrl,
		arg.ProgressionStatus,
		arg.AvailableStatus,
		arg.ReleaseYear,
		arg.Order,
		arg.TotalEpisode,
		arg.UpdatedBy,
		arg.UpdatedAt,
	)
	var i DSeason
	err := row.Scan(
		&i.ID,
		&i.MovieID,
		&i.Name,
		&i.Description,
		&i.ThumbnailUrl,
		&i.ProgressionStatus,
		&i.AvailableStatus,
		&i.ReleaseYear,
		&i.Rate,
		&i.TotalRate,
		&i.TotalViewer,
		&i.TotalUpvote,
		&i.TotalDownvote,
		&i.Order,
		&i.TotalEpisode,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedBy,
		&i.UpdatedAt,
	)
	return i, err
}