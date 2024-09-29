-- name: CreateMovie :one
INSERT INTO d_movies (
        movie_id,
        item_status,
        "name",
        "description",
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
RETURNING *;
-- name: UpdateMovie :one
UPDATE d_movies
SET item_status = COALESCE($2, item_status),
    "name" = COALESCE($3, "name"),
    "description" = COALESCE($4, "description"),
    thumbnail_url = COALESCE($5, thumbnail_url),
    available_status = COALESCE($6, available_status),
    types = COALESCE($7, types),
    release_year = COALESCE($8, release_year),
    updated_by = COALESCE($9, updated_by),
    updated_at = COALESCE($10, updated_at)
WHERE movie_id = $1
RETURNING *;
-- name: GetMovieById :one
SELECT *
FROM d_movies
WHERE movie_id = $1;
-- name: CreateSeason :one
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
RETURNING *;
-- name: UpdateSeason :one
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
RETURNING *;
-- name: GetSeasonsByMovieId :many
SELECT *
FROM d_seasons
WHERE movie_id = $1
    AND (
        $2::VARCHAR IS NULL
        OR id::VARCHAR = $2
    )
ORDER BY "order";
-- name: CreateEpisode :one
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
RETURNING *;
-- name: UpdateEpisode :one
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
RETURNING *;
-- name: GetEpisodesBySeasonId :many
SELECT *
FROM d_episodes
WHERE season_id = $1
    AND (
        $2::VARCHAR IS NULL
        OR id::VARCHAR = $2
    )
ORDER BY episode_order;