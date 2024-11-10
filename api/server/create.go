package server

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	echoSwagger "github.com/swaggo/echo-swagger"
	appMiddleware "org.idev.koala/backend/api/middleware"
	"org.idev.koala/backend/api/route"
	"org.idev.koala/backend/app"
	"org.idev.koala/backend/common/errors"
	"org.idev.koala/backend/common/logger"
	"org.idev.koala/backend/component/kafka"
	"org.idev.koala/backend/component/mongo"
	"org.idev.koala/backend/component/redis"
	"org.idev.koala/backend/component/storage"
)

func Create(ctx context.Context) (*Server, error) {
	log := logger.New("Server", "create server")

	appConfig, err := app.LoadConfig()
	if err != nil {
		return nil, errors.Wrap(err, "failed to load config")
	}

	e := setupEcho()

	AppCtx = &app.AppContext{
		Ctx:    ctx,
		Config: appConfig,
	}

	if err := setupComponents(ctx, AppCtx, appConfig, log); err != nil {
		return nil, errors.Wrap(err, "failed to setup components")
	}

	e.Use(appMiddleware.AuthGuard(AppCtx))

	setupRoutes(e, AppCtx)

	return &Server{e: e}, nil
}

func setupEcho() *echo.Echo {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return e
}

func setupComponents(ctx context.Context, appCtx *app.AppContext, config *app.Config, log *logger.Logger) error {
	var err error

	if config.EnableRedis {
		appCtx.Redis, err = setupRedis(ctx, config.RedisUrl, log)
		if err != nil {
			return errors.Wrap(err, "failed to setup Redis")
		}
	}

	if config.EnableDb {
		appCtx.Db, err = setupDatabase(ctx, config.DbUrl, log)
		if err != nil {
			return errors.Wrap(err, "failed to setup database")
		}
	}

	if config.EnableMongo {
		appCtx.Mongo, err = setupMongo(ctx, config.MongoUrl, config.MongoDbName, log)
		if err != nil {
			return errors.Wrap(err, "failed to setup MongoDB")
		}
	}

	if config.EnableKafka {
		appCtx.KafkaProducer, err = setupKafka(config.KafkaHost, config.KafkaPort, log)
		if err != nil {
			return errors.Wrap(err, "failed to setup Kafka")
		}
	}

	if config.EnableStorage {
		appCtx.StorageCli, err = setupStorage(config.StorageHost, config.StoragePort, config.StorageId, config.StorageSecret)
		if err != nil {
			return errors.Wrap(err, "failed to setup storage")
		}
	}

	return nil
}

func setupRedis(ctx context.Context, url string, log *logger.Logger) (*redis.Client, error) {
	log.Info("Connecting to Redis: " + url)
	return redis.NewClient(ctx, url)
}

func setupDatabase(ctx context.Context, url string, log *logger.Logger) (*pgxpool.Pool, error) {
	log.Info("Connecting to database: " + url)
	poolCfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse database config")
	}
	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create database pool")
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to ping database")
	}
	return pool, nil
}

func setupMongo(ctx context.Context, url, dbName string, log *logger.Logger) (*mongo.Client, error) {
	log.Info("Connecting to MongoDB: " + url)
	return mongo.NewMongoClient(ctx, url, dbName)
}

func setupKafka(host string, port int32, log *logger.Logger) (*kafka.Producer, error) {
	log.Info(fmt.Sprintf("Connecting to Kafka cluster: %s:%d", host, port))
	return kafka.NewProducer(host, port)
}

func setupStorage(host string, port int32, id string, secret string) (*storage.StorageClient, error) {
	log.Info(fmt.Sprintf("Connecting to Storage: %s:%d", host, port))
	minioClient, err := minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4(id, secret, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}
	return storage.NewStorageClient(minioClient)
}

func setupRoutes(e *echo.Echo, appCtx *app.AppContext) {
	route.NewExamplePanicErrorRouter(e)

	v1 := e.Group("/api/v1")
	route.NewUserRouter(v1, appCtx)
	route.NewMovieRouter(v1, appCtx)

	storage := e.Group("/storage")
	route.NewMovieStorageRouter(storage, appCtx)
}
