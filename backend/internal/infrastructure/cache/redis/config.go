package redis

import (
	"os"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
}

func NewConfig(host, user, password string, port int) *Config {
	return &Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
	}
}

func ReadConfigFromEnvironment() *Config {
	host := os.Getenv("REDIS_HOST")

	if host == "" {
		panic("хост для редиса пустой")
	}

	port, err := strconv.Atoi(os.Getenv("REDIS_PORT"))

	if err != nil {
		panic("порт для редиса пустой или не может быть приведен к числу")
	}

	user := os.Getenv("REDIS_USER")
	//if user == "" {
	//	panic("пользователь для редиса не задан")
	//}

	password := os.Getenv("REDIS_USERPASSWORD")
	//if password == "" {
	//	panic("пользователь для редиса не задан")
	//}

	return NewConfig(host, user, password, port)
}
