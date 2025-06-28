package config

const EnvPath = ".env"

type AppConfig struct {
	Rest Rest
}

type Rest struct {
	ListenAddress string `envconfig:"PORT" required:"true"`
}
