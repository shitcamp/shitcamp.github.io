package config

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type TwitchConfig struct {
	ClientID    string `mapstructure:"client_id"`    // client ID for app
	AccessToken string `mapstructure:"access_token"` // secret app access token to authorize Twitch API requests
}

type CacheConfig struct {
	DefaultExpiryTime      time.Duration `mapstructure:"default_expiry_time"`
	DefaultCleanupInterval time.Duration `mapstructure:"default_cleanup_interval"`
}

type Config struct {
	ServerAddress string       `mapstructure:"address"`
	Twitch        TwitchConfig `mapstructure:"twitch"`
	Cache         CacheConfig  `mapstructure:"cache"`
	Debug         bool         `mapstructure:"debug"`
}

var cfg = Config{
	Cache: CacheConfig{
		DefaultExpiryTime:      2 * time.Minute,
		DefaultCleanupInterval: 5 * time.Minute,
	},
	Debug: false,
}

const (
	EnvAddress   = "ADDRESS"
	EnvLocalPort = "LOCAL_PORT"

	EnvTwitchClientID    = "TWITCH_CLIENT_ID"
	EnvTwitchAccessToken = "TWITCH_ACCESS_TOKEN"

	EnvCacheDefaultExpiryTime      = "CACHE_DEFAULT_EXPIRY_TIME"
	EnvCacheDefaultCleanupInterval = "CACHE_DEFAULT_CLEANUP_INTERVAL"

	EnvDebug = "DEBUG"
)

func loadConfigsFromFile(configFile string) error {
	configFile, err := filepath.Abs(configFile)
	if err != nil {
		panic(err)
	}

	logger.WithFields(logger.Fields{"path": configFile}).Info("loading_config")

	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("json")

	err = v.ReadInConfig()
	if err != nil {
		logger.Fatal("error parsing config file:", err)
		return err
	}

	err = v.Unmarshal(&cfg)
	if err != nil {
		return err
	}

	return nil
}

func loadConfigsFromEnv() {
	if address := os.Getenv(EnvAddress); address != "" {
		cfg.ServerAddress = address
	}
	if port := os.Getenv(EnvLocalPort); port != "" {
		cfg.ServerAddress = "0.0.0.0:" + port
	}

	if twitchClientID := os.Getenv(EnvTwitchClientID); twitchClientID != "" {
		cfg.Twitch.ClientID = twitchClientID
	}
	if twitchAccessToken := os.Getenv(EnvTwitchAccessToken); twitchAccessToken != "" {
		cfg.Twitch.AccessToken = twitchAccessToken
	}

	if expiryTime := os.Getenv(EnvCacheDefaultExpiryTime); expiryTime != "" {
		if d, err := time.ParseDuration(expiryTime); err == nil {
			cfg.Cache.DefaultExpiryTime = d
		}
	}
	if cleanupInterval := os.Getenv(EnvCacheDefaultCleanupInterval); cleanupInterval != "" {
		if d, err := time.ParseDuration(cleanupInterval); err == nil {
			cfg.Cache.DefaultCleanupInterval = d
		}
	}

	if debug := os.Getenv(EnvDebug); debug != "" {
		debugBool, _ := strconv.ParseBool(debug)
		cfg.Debug = debugBool
	}
}

func LoadConfig(configFile string) error {
	if configFile != "" {
		err := loadConfigsFromFile(configFile)
		if err != nil {
			return err
		}
	}

	loadConfigsFromEnv()

	logger.WithField("config", cfg).Info("loaded_config")
	return nil
}

func SetConfig(c *Config) {
	cfg = *c
}

func GetConfig() *Config {
	return &cfg
}
