package config

import (
	"github.com/garaekz/goshort/pkg/log"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-env"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	defaultServerPort         = 8080
	defaultJWTExpirationHours = 72
	defaultDBType             = "mysql"
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
	// the data source name (DSN) for connecting to the database. required.
	DBType string `yaml:"db_type" env:"DB_TYPE,secret"`
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
	var port int
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = defaultServerPort
	}

	c := Config{
		ServerPort:    port,
		JWTExpiration: defaultJWTExpirationHours,
		DBType:        defaultDBType,
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
