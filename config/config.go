package config

type Config struct {
	ConnectionString string
	SecretKey        string
}

func New() *Config {
	return &Config{
		ConnectionString: "postgres://postgres:postgres@localhost:5432/morgan?sslmode=disable",
		SecretKey:        "1498b5467a63dffa2dc9d9e069caf075d16fc33fdd4c3b01bfadae6433767d93",
	}
}
