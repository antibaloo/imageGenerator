package configs

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type ConfI interface {
	GetPort() string
}

type Conf struct {
	port string
}

func (c *Conf) GetPort() string {
	return c.port
}

func New(confPath string) (*Conf, error) {
	err := godotenv.Load(confPath)
	if err != nil {
		return nil, err
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		return nil, errors.New("server port is empty string")
	}
	return &Conf{
		port: port,
	}, nil
}
