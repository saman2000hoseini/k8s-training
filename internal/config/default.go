package config

import "github.com/saman2000hoseini/k8s-training/internal/db"

func Default() Config {
	return Config{
		Redis: db.Config{
			Address:  "localhost:6379",
			Password: "secret",
		},
		Server: Server{Port: 65432},
	}
}
