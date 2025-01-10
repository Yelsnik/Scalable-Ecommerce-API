package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver             string `mapstructure:"DB_DRIVER"`
	DBSource             string `mapstructure:"DB_SOURCE"`
	GRPCServerAddress    string `mapstructure:"GRPC_SERVER_ADDRESS"`
	StripeSecretKey      string `mapstructure:"STRIPE_SECRET_KEY"`
	StripePublishableKey string `mapstructure:"STRIPE_PUBLISHABLE_KEY"`
	WebhookSigningKey    string `mapstructure:"WEBHOOK_SIGNING_KEY"`
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
