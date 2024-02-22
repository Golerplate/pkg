package pgx

type PGXConfig struct {
	Host     string `env:"DB_HOST"`
	Username string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DBName   string `env:"DB_NAME"`
	Port     uint16 `env:"DB_PORT"`
}
