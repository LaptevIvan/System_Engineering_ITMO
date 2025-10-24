package config

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/viper"
)

const (
	defaultLogValue = true
)

type (
	Config struct {
		GRPC struct {
			Port        string `env:"GRPC_PORT"`
			GatewayPort string `env:"GRPC_GATEWAY_PORT"`
		}

		PG struct {
			URL      string
			Host     string `env:"POSTGRES_HOST"`
			Port     string `env:"POSTGRES_PORT"`
			DB       string `env:"POSTGRES_DB"`
			User     string `env:"POSTGRES_USER"`
			Password string `env:"POSTGRES_PASSWORD"`
			MaxConn  string `env:"POSTGRES_MAX_CONN"`
		}

		Log struct {
			LogController   bool `env:"LOG_CONTROLLER_ENABLED"`
			LogTransactor   bool `env:"LOG_TRANSACTOR_ENABLED"`
			LogUseCase      bool `env:"LOG_USECASE_ENABLED"`
			LogDBRepo       bool `env:"LOG_DB_REPO_ENABLED"`
			LogOutboxWorker bool `env:"LOG_OUTBOX_WORKER_ENABLED"`
		}
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	cfg.GRPC.Port = os.Getenv("GRPC_PORT")
	cfg.GRPC.GatewayPort = os.Getenv("GRPC_GATEWAY_PORT")

	cfg.PG.Host = os.Getenv("POSTGRES_HOST")
	cfg.PG.Port = os.Getenv("POSTGRES_PORT")
	cfg.PG.DB = os.Getenv("POSTGRES_DB")
	cfg.PG.User = os.Getenv("POSTGRES_USER")
	cfg.PG.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.PG.MaxConn = os.Getenv("POSTGRES_MAX_CONN")

	cfg.PG.URL = fmt.Sprintf("postgres://%s:%s@", cfg.PG.User, cfg.PG.Password) +
		net.JoinHostPort(cfg.PG.Host, cfg.PG.Port) + fmt.Sprintf("/%s?sslmode=disable", cfg.PG.DB)

	var err error
	v := viper.New()

	if cfg.Log.LogController, err = parseEnvBool(v, "log_controller", "LOG_CONTROLLER_ENABLED", defaultLogValue); err != nil {
		return nil, err
	}

	if cfg.Log.LogTransactor, err = parseEnvBool(v, "log_transactor", "LOG_TRANSACTOR_ENABLED", defaultLogValue); err != nil {
		return nil, err
	}

	if cfg.Log.LogUseCase, err = parseEnvBool(v, "log_usecase", "LOG_USECASE_ENABLED", defaultLogValue); err != nil {
		return nil, err
	}

	if cfg.Log.LogDBRepo, err = parseEnvBool(v, "log_db", "LOG_DB_REPO_ENABLED", defaultLogValue); err != nil {
		return nil, err
	}

	if cfg.Log.LogOutboxWorker, err = parseEnvBool(v, "log_outbox_worker", "LOG_OUTBOX_WORKER_ENABLED", defaultLogValue); err != nil {
		return nil, err
	}

	return cfg, nil
}

func parseEnvBool(v *viper.Viper, key, envVar string, defaultValue ...bool) (bool, error) {
	err := v.BindEnv(key, envVar)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0], err
		}
		return false, err
	}
	if len(defaultValue) > 0 {
		v.SetDefault(key, defaultValue[0])
	}
	return v.GetBool(key), nil
}
