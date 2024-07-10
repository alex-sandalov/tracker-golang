package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	// Database configuration.
	Database Database `env-prefix:"DB_"`

	// CORS configuration.
	CORS CORS `env-prefix:"CORS_"`

	Server Server `env-prefix:"SERVER_"`
}

type Server struct {
	// Port specifies the server port.
	// It is set by the `PORT` environment variable.
	// Default value is "8080".
	Port string `env:"PORT" env-default:"8080"`

	// MaxHeaderBytes specifies the maximum allowed size of the request header.
	// It is set by the `MAX_HEADER_BYTES` environment variable.
	MaxHeaderBytes int `env:"MAX_HEADER_BYTES"`

	// ReadTimeout specifies the maximum duration to read the request.
	// It is set by the `READ_TIMEOUT` environment variable.
	// Default value is 10 seconds.
	ReadTimeout time.Duration `env:"READ_TIMEOUT" env-default:"10s"`

	// WriteTimeout specifies the maximum duration to write the response.
	// It is set by the `WRITE_TIMEOUT` environment variable.
	// Default value is 10 seconds.
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" env-default:"10s"`
}

type Database struct {
	// Host is the hostname or IP address of the database server.
	// Default: "localhost"
	Host string `env:"HOST" env-default:"localhost"`

	// Port is the port number of the database server.
	// Default: "5432"
	Port string `env:"PORT" env-default:"5432"`

	// User is the username used to authenticate with the database.
	// Default: "postgres"
	User string `env:"USER" env-default:"postgres"`

	// Password is the password used to authenticate with the database.
	// Required.
	Password string `env:"PASSWORD" env-required:"true"`

	// Name is the name of the database to connect to.
	// Required.
	Name string `env:"NAME" env-required:"true"`

	// SSLMode is the SSL mode used to connect to the database.
	// Default: "disable"
	SSLMode string `env:"SSL_MODE" env-default:"disable"`
}

// CORS configuration.
type CORS struct {
	// AllowedOrigins is a list of origins a cross-domain request can be
	// executed from.
	// Required.
	AllowedOrigins []string `env:"ALLOWED_ORIGINS" env-required:"true"`

	// AllowedMethods is a list of methods the client is allowed to use with
	// cross-domain requests.
	// Required.
	AllowedMethods []string `env:"ALLOWED_METHODS" env-required:"true"`

	// AllowedHeaders is a list of non simple headers the client is allowed to use
	// with cross-domain requests.
	// Required.
	AllowedHeaders []string `env:"ALLOWED_HEADERS" env-required:"true"`

	// AllowedCredentials indicates whether the request can include user credentials.
	// Required.
	AllowedCredentials bool `env:"ALLOWED_CREDENTIALS" env-required:"true"`

	// MaxAge indicates how long (in seconds) the results of a preflight request
	// can be cached.
	// Required.
	MaxAge time.Duration `env:"MAX_AGE" env-required:"true"`
}

// MustLoad loads the configuration from the specified file path.
// It returns a pointer to the loaded configuration.
//
// If the CONFIG_PATH environment variable is not set, it logs a fatal error and exits.
// If the file specified by CONFIG_PATH does not exist, it logs a fatal error and exits.
// If there is an error reading the configuration file, it logs a fatal error and exits.
func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatalf("CONFIG_PATH IS NOT SET")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("not found %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("%s", err)
	}

	return &cfg
}
