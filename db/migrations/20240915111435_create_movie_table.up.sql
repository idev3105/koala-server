CREATE TABLE d_movies (
    id BIGSERIAL PRIMARY KEY,
    movie_id VARCHAR(255) NOT NULL,
    item_status VARCHAR(50) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL,
    thumbnail_url TEXT,
    available_status VARCHAR(50) NOT NULL,
    types VARCHAR(255) [],
    release_year INTEGER NOT NULL,
    total_viewer INTEGER DEFAULT 0,
    total_upvote INTEGER DEFAULT 0,
    total_downvote INTEGER DEFAULT 0,
    total_rate INTEGER DEFAULT 0,
    rate NUMERIC(3, 2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    updated_by VARCHAR(255)
);
ALTER TABLE d_movies
ADD CONSTRAINT unique_movie_id UNIQUE (movie_id);