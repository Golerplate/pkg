package cache

type Config struct {
	Host string `env:"CACHE_HOST"`
	Port uint16 `env:"CACHE_PORT"`
}
