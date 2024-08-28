package server

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	appMiddleware "org.idev.koala/backend/api/middleware"
	"org.idev.koala/backend/api/route"
	"org.idev.koala/backend/app"
	"org.idev.koala/backend/common/logger"
	"org.idev.koala/backend/component/kafka"
	"org.idev.koala/backend/component/mongo"
	"org.idev.koala/backend/component/redis"
)

func Create(ctx context.Context) (*Server, error) {
	e := echo.New()
	log := logger.New("Server", "create server")

	e.Debug = true

	e.Use(middleware.Recover()) // recover from panic, use can custom error handler
	e.Use(middleware.Logger())  // show http request logs

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	appConfig, err := app.LoadConfig()
	if err != nil {
		return nil, err
	}

	log.Info("Connect to redis " + appConfig.RedisUrl)
	redisCli, err := redis.NewClient(ctx, appConfig.RedisUrl)
	if err != nil {
		return nil, err
	}

	log.Info("Connect to database " + appConfig.DbUrl)
	poolCfg, err := pgxpool.ParseConfig(appConfig.DbUrl)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	var mongoCli *mongo.Client
	if appConfig.MongoUrl != "" {
		log.Info("Connect to mongodb " + appConfig.MongoUrl)
		mongoCli, err = mongo.NewMongoClient(ctx, appConfig.MongoUrl, appConfig.MongoDbName)
		if err != nil {
			return nil, err
		}
	} else {
		log.Info("MongoDB is not configured. Skip connecting to MongoDB.")
	}

	var kafkaProducer *kafka.Producer
	if appConfig.KafkaHost != "" {
		log.Info("Connect to kafka cluster " + appConfig.KafkaHost + ":" + fmt.Sprint(appConfig.KafkaPort))
		kafkaProducer, err = kafka.NewProducer(appConfig.KafkaHost, appConfig.KafkaPort)
		if err != nil {
			return nil, err
		}
	} else {
		log.Info("Kafka is not configured. Skip connecting to Kafka.")
	}

	// init app context instance
	AppCtx = &app.AppContext{
		Ctx:           ctx,
		Config:        appConfig,
		Db:            pool,
		RedisCli:      redisCli,
		KafkaProducer: kafkaProducer,
		MongoClient:   mongoCli,
	}

	e.Use(appMiddleware.AuthGuard(AppCtx))

	route.NewExamplePanicErrorRouter(e)

	v1 := e.Group("/api/v1")
	{
		route.NewUserRouter(v1, AppCtx)
		route.NewMovieRouter(v1, AppCtx)
	}

	return &Server{e: e}, nil
}
