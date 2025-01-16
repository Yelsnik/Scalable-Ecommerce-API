package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver                   string        `mapstructure:"DB_DRIVER"`
	DBSource                   string        `mapstructure:"DB_SOURCE"`
	HTTPServerAddress          string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress          string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey          string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration        time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration       time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EmailSenderName            string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress         string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword        string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
	PasswordResetTokenDuration time.Duration `mapstructure:"PASSWORD_RESET_TOKEN_DURATION"`
	StripeSecretKey            string        `mapstructure:"STRIPE_SECRET_KEY"`
	RabbitMq                   string        `mapstructure:"RABBITMQ"`
	//PasswordResetTokenDuration time.Duration `mapstructure:"PASSWORD_RESET_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
