package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type JWT struct {
	Salt string `mapstructure:"salt"`
}
type AppsConfig struct {
	Port               string `mapstructure:"port"`
	AccountSecretKey   string `mapstructure:"account_secret_key"`
	AuthSecretKey      string `mapstructure:"auth_secret_key"`
	ChildrenSecretKey  string `mapstructure:"children_secret_key"`
	ProfanitySecretKey string `mapstructure:"profanity_secret_key"`
	JWT                JWT    `mapstructure:"jwt"`
}

type MLConfig struct {
	Url string `mapstructure:"url"`
}

type GormConfig struct {
	Dialect string `mapstructure:"dialect"`
}

type PostgresConfig struct {
	Url string `mapstructure:"url"`
}

type RabbitmqConfig struct {
	Url string `mapstructure:"url"`
}

type RedisConfig struct {
	Url      string `mapstructure:"url"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
}

type Config struct {
	ENV            string         `mapstructure:"env"`
	App            AppsConfig     `mapstructure:"apps"`
	MLSevice       MLConfig       `mapstructure:"ml_profanity"`
	GormConfig     GormConfig     `mapstructure:"gorm"`
	PostgresConfig PostgresConfig `mapstructure:"postgres"`
	RabbitmqConfig RabbitmqConfig `mapstructure:"rabbitmq"`
	RedisConfig    RedisConfig    `mapstructure:"redis"`
}

func DefaultConfig() (Config, error) {

	if os.Getenv("ENV") != "DOCKER" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	env := os.Getenv("ENV")

	if env == "" {
		os.Setenv("ENV", "prod")
		env = "prod"
	}

	v := viper.New()
	v.SetConfigName("config." + env)
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("config/yaml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Println(err)
		return Config{}, err
	}

	// Using .env / Docker ENVIROMENT as main env
	for _, key := range v.AllKeys() {
		value := v.GetString(key)
		envOrRaw := replaceEnvInConfig([]byte(value))
		v.Set(key, string(envOrRaw))
	}

	cfg := Config{}
	if err := v.Unmarshal(&cfg); err != nil {
		log.Println(err)
		return Config{}, err
	}

	return cfg, nil
}

func replaceEnvInConfig(body []byte) []byte {
	search := regexp.MustCompile(`\$\{([^{}]+)\}`)
	replacedBody := search.ReplaceAllFunc(body, func(b []byte) []byte {
		group1 := search.ReplaceAllString(string(b), `$1`)

		envValue := os.Getenv(group1)
		if len(envValue) > 0 {
			return []byte(envValue)
		}
		return []byte("")
	})
	return replacedBody
}
