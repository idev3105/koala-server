package app

import (
	"github.com/spf13/viper"
	"org.idev.koala/backend/api/enum"
)

type Config struct {
	Env  enum.Env `mapstructure:"ENV"`
	Port string   `mapstructure:"PORT"`

	EnableDb bool   `mapstructure:"ENABLE_DB"`
	DbUrl    string `mapstructure:"DB_URL"`

	EnableRedis bool   `mapstructure:"ENABLE_REDIS"`
	RedisUrl    string `mapstructure:"REDIS_URL"`

	EnableKafka bool   `mapstructure:"ENABLE_KAFKA"`
	KafkaHost   string `mapstructure:"KAFKA_HOST"`
	KafkaPort   int32  `mapstructure:"KAFKA_PORT"`

	EnableMongo bool   `mapstructure:"ENABLE_MONGO"`
	MongoUrl    string `mapstructure:"MONGO_URL"`
	MongoDbName string `mapstructure:"MONGO_DB_NAME"`

	EnableStorage bool   `mapstructure:"ENABLE_STORAGE"`
	StorageHost   string `mapstructure:"STORAGE_HOST"`
	StoragePort   int32  `mapstructure:"STORAGE_PORT"`
	StorageId     string `mapstructure:"STORAGE_ID"`
	StorageSecret string `mapstructure:"STORAGE_SECRET"`

	JWKsUrl string `mapstructure:"JWKS_URL"`
}

func LoadConfig() (*Config, error) {

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	appConfig := &Config{}
	if err := viper.Unmarshal(appConfig); err != nil {
		return nil, err
	}

	return appConfig, nil
}
