package postgres

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func (c *Config) CreateDbUrl() string {
	connectionUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", c.User, c.Password, c.Host, strconv.Itoa(c.Port), c.Database)
	fmt.Println(connectionUrl)

	return connectionUrl
}

func readConfigFromEnvAndValidate() (string, int, string, string, string) {
	host := os.Getenv("HOST")

	if host == "" {
		panic("Не передан хост для базы данных")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}

	user := os.Getenv("USER")
	if user == "" {
		panic("Не передан пользователь БД")
	}

	password := os.Getenv("PASSWORD")
	if password == "" {
		panic("Не передан пароль от БД")
	}

	database := os.Getenv("DB")
	if database == "" {
		panic("Не передано имя БД")
	}
	return host, port, user, password, database
}

func NewConfig() *Config {
	host, port, user, password, database := readConfigFromEnvAndValidate()

	return &Config{
		host,
		port,
		user,
		password,
		database,
	}
}
