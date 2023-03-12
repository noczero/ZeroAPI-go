package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV" json:"app_env,omitempty"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS" json:"server_address,omitempty"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT" json:"context_timeout,omitempty"`
	DBHost                 string `mapstructure:"DB_HOST" json:"db_host,omitempty"`
	DBPort                 string `mapstructure:"DB_PORT" json:"db_port,omitempty"`
	DBUser                 string `mapstructure:"DB_USER" json:"db_user,omitempty"`
	DBPass                 string `mapstructure:"DB_PASS" json:"db_pass,omitempty"`
	DBName                 string `mapstructure:"DB_NAME" json:"db_name,omitempty"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR" json:"access_token_expiry_hour,omitempty"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR" json:"refresh_token_expiry_hour,omitempty"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET" json:"access_token_secret,omitempty"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET" json:"refresh_token_secret,omitempty"`
	APIKeyChatGPT          string `mapstructure:"API_KEY_CHAT_GPT" json:"api_key_chat_gpt,omitempty"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
