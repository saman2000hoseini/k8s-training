package db

type Config struct {
	Address  string `koanf:"address"`
	Password string `koanf:"password"`
}
