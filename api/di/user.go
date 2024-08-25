package di

import (
	"org.idev.koala/backend/component/redis"
	userdomain "org.idev.koala/backend/domain/user"
	sqlc_generated "org.idev.koala/backend/generated/sqlc"
	userrepository "org.idev.koala/backend/repository/user"
)

func NewUserUseCase(queries *sqlc_generated.Queries, redisCli *redis.RedisClient) userdomain.UserUseCase {
	repo := NewUserRepository(queries, redisCli)
	return userdomain.NewUserUseCase(repo)
}

func NewUserRepository(queries *sqlc_generated.Queries, redisCli *redis.RedisClient) userdomain.UserRepository {
	sql := userrepository.NewSqlRepository(queries)
	cache := userrepository.NewCacheRepository(redisCli)
	return userrepository.New(sql, cache)
}
