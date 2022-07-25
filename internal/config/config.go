package config

import (
	"context"
	"io/ioutil"
	"strconv"

	"github.com/garaekz/goshort/pkg/dbcontext"
	"github.com/garaekz/goshort/pkg/log"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-env"
	"gopkg.in/yaml.v2"
)

const (
	defaultServerPort         = 8080
	defaultJWTExpirationHours = 72
)

// Config represents an application configuration.
type Config struct {
	// the server port. Defaults to 8080
	ServerPort int `yaml:"server_port" env:"SERVER_PORT"`
	// the data source name (DSN) for connecting to the database. required.
	DSN string `yaml:"dsn" env:"DSN,secret"`
	// JWT signing key. required.
	JWTSigningKey string `yaml:"jwt_signing_key" env:"JWT_SIGNING_KEY,secret"`
	// JWT expiration in hours. Defaults to 72 hours (3 days)
	JWTExpiration int `yaml:"jwt_expiration" env:"JWT_EXPIRATION"`
	// MaxAPIKeysPerUser is the maximum number of API keys per user. Defaults to 1.
	MaxAPIKeysPerUser int `yaml:"max_api_keys_per_user" env:"MAX_API_KEYS_PER_USER" db:"default_api_keys_quantity"`
}

// Validate validates the application configuration.
func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.DSN, validation.Required),
		validation.Field(&c.JWTSigningKey, validation.Required),
	)
}

// Load returns an application configuration which is populated from the given configuration file and environment variables.
func Load(file string, logger log.Logger) (*Config, error) {
	// default config
	c := Config{
		ServerPort:    defaultServerPort,
		JWTExpiration: defaultJWTExpirationHours,
	}

	// load from YAML config file
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	// load from environment variables prefixed with "APP_"
	if err = env.New("APP_", logger.Infof).Load(&c); err != nil {
		return nil, err
	}

	// validation
	if err = c.Validate(); err != nil {
		return nil, err
	}

	return &c, err
}

func (c *Config) GetConfigFromDB(db *dbcontext.DB, logger log.Logger) error {
	var config []struct {
		Name  string `db:"name"`
		Value string `db:"value"`
	}

	err := db.With(context.Background()).Select().From("configs").All(&config)
	if err != nil {
		return err
	}

	for _, cfg := range config {
		switch cfg.Name {
		case "default_api_keys_quantity":
			val, err := strconv.Atoi(cfg.Value)
			if err != nil {
				return err
			}
			c.MaxAPIKeysPerUser = val
		}
	}

	return nil
}
