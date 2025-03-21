// Package config handles application configuration.
// It loads environment variables and provides structured access to them.
package config

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

// Environment stores application configuration loaded from environment variables.
type Environment struct {
	Port string `env:"PORT" envDefault:"8080"`
	DB   *dbEnvironment
}

// dbEnvironment holds database connection parameters.
type dbEnvironment struct {
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	Name     string `env:"DB_NAME,required"`
	Port     string `env:"DB_PORT,required"`
	Host     string `env:"DB_HOST,required"`
}

// ConnString returns the formatted PostgreSQL connection string.
func (db *dbEnvironment) ConnString() string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		db.User,
		db.Password,
		db.Name,
		db.Host,
		db.Port,
	)
}

// Load parses environment variables into an Environment structure.
func Load() (*Environment, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	environment := &Environment{}
	environment.DB = &dbEnvironment{}

	err := env.Parse(environment)

	return environment, err
}
