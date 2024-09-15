package di

import (
	moviedomain "org.idev.koala/backend/domain/movie"
	sqlc_generated "org.idev.koala/backend/generated/sqlc"
	movierepo "org.idev.koala/backend/repository/movie"
)

func NewMovieUseCase(queries *sqlc_generated.Queries) moviedomain.MovieUseCase {
	movieSqlRepo := movierepo.NewMovieSqlRepo(queries)
	movieRepo := movierepo.NewMovieRepo(movieSqlRepo)
	return moviedomain.NewMovieUseCase(movieRepo)
}
