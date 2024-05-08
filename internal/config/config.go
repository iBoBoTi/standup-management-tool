package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

// Config map .env configs to this struct
type Config struct {
	Timezone             string        `mapstructure:"TIMEZONE"`
	DBHost               string        `mapstructure:"DB_HOST"`
	DBUser               string        `mapstructure:"DB_USER"`
	DBPassword           string        `mapstructure:"DB_PASSWORD"`
	DBName               string        `mapstructure:"DB_NAME"`
	DBPort               string        `mapstructure:"DB_PORT"`
	HttpServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFERESH_TOKEN_DURATION"`
	Environment          string        `mapstructure:"ENVIRONMENT"`
}

func Load(p string) (cfg Config, err error) {
	return loader(p, ".env")
}

func LoadWithPath(p string, env string) (cfg Config, err error) {
	return loader(p, env)
}

func loader(p string, env string) (cfg Config, err error) {
	environment := os.Getenv("ENVIRONMENT")
	viper.AddConfigPath(p)
	viper.SetConfigName(env)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)
	log.Println("environment: ",environment)

	if environment == "production" {
		viper.SetDefault("ENVIRONMENT", environment)
		viper.SetDefault("DB_HOST", os.Getenv("DB_HOST"))
		viper.SetDefault("DB_USER", os.Getenv("DB_USER"))
		viper.SetDefault("DB_PASSWORD", os.Getenv("DB_PASSWORD"))
		viper.SetDefault("DB_NAME", os.Getenv("DB_NAME"))
		viper.SetDefault("DB_PORT", os.Getenv("DB_PORT"))
		viper.SetDefault("HTTP_SERVER_ADDRESS", os.Getenv("HTTP_SERVER_ADDRESS"))
		viper.SetDefault("TOKEN_SYMMETRIC_KEY", os.Getenv("TOKEN_SYMMETRIC_KEY"))
		viper.SetDefault("ACCESS_TOKEN_DURATION", os.Getenv("ACCESS_TOKEN_DURATION"))
		viper.SetDefault("REFERESH_TOKEN_DURATION", os.Getenv("REFERESH_TOKEN_DURATION"))
	}

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config")
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}
