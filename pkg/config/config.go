package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type TwitchConfig struct {
	ClientID    string `mapstructure:"client_id"`    // client ID for app
	AccessToken string `mapstructure:"access_token"` // secret app access token to authorize Twitch API requests
}

type ShitcampConfig struct {
	OldestUploadTime string `mapstructure:"oldest_upload_time"`
	NewestUploadTime string `mapstructure:"newest_upload_time"`
}

func (s *ShitcampConfig) GetOldestUploadTime() time.Time {
	t, err := time.Parse(time.RFC3339, s.OldestUploadTime)
	if err != nil {
		logger.WithField("time", s.OldestUploadTime).Warn("parse_GetOldestUploadTime_error")
		t, _ = time.Parse(time.RFC3339, shitcampStartTime)
		return t
	}

	return t
}

func (s *ShitcampConfig) GetNewestUploadTime() time.Time {
	t, err := time.Parse(time.RFC3339, s.NewestUploadTime)
	if err != nil {
		logger.WithField("time", s.NewestUploadTime).Warn("parse_GetNewestUploadTime_error")
		t, _ = time.Parse(time.RFC3339, shitcampEndTime)
		return t
	}

	return t
}

type CacheConfig struct {
	DefaultExpiryTime      time.Duration `mapstructure:"default_expiry_time"`
	DefaultCleanupInterval time.Duration `mapstructure:"default_cleanup_interval"`
}

type RateLimitConfig struct {
	Tokens   uint64        `mapstructure:"tokens"`
	Interval time.Duration `mapstructure:"interval"`
}

type Config struct {
	ServerAddress string          `mapstructure:"address"`
	Twitch        TwitchConfig    `mapstructure:"twitch"`
	Shitcamp      ShitcampConfig  `mapstructure:"shitcamp"`
	Cache         CacheConfig     `mapstructure:"cache"`
	Auth          gin.Accounts    `mapstructure:"auth"`
	RateLimit     RateLimitConfig `mapstructure:"rate_limit"`
	Debug         bool            `mapstructure:"debug"`
}

const shitcampStartTime = "2021-09-26T19:00:00.00-06:00"
const shitcampEndTime = "2021-09-30T23:00:00.00-06:00"

var cfg = Config{
	Shitcamp: ShitcampConfig{
		OldestUploadTime: shitcampStartTime,
		NewestUploadTime: shitcampEndTime,
	},
	Cache: CacheConfig{
		DefaultExpiryTime:      2 * time.Minute,
		DefaultCleanupInterval: 5 * time.Minute,
	},
	Auth: gin.Accounts{},
	RateLimit: RateLimitConfig{
		Tokens:   249,
		Interval: 290 * time.Second,
	},
	Debug: false,
}

const (
	EnvAddress   = "ADDRESS"
	EnvLocalPort = "LOCAL_PORT"

	EnvTwitchClientID    = "TWITCH_CLIENT_ID"
	EnvTwitchAccessToken = "TWITCH_ACCESS_TOKEN"

	EnvShitcampOldestUploadTime = "SHITCAMP_OLDEST_UPLOAD_TIME"
	EnvShitcampNewestUploadTime = "SHITCAMP_NEWEST_UPLOAD_TIME"

	EnvCacheDefaultExpiryTime      = "CACHE_DEFAULT_EXPIRY_TIME"
	EnvCacheDefaultCleanupInterval = "CACHE_DEFAULT_CLEANUP_INTERVAL"

	EnvAuth = "AUTH"

	EnvRateLimitTokens   = "RATE_LIMIT_TOKENS"
	EnvRateLimitInterval = "RATE_LIMIT_INTERVAL"

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

	if oldestUploadTime := os.Getenv(EnvShitcampOldestUploadTime); oldestUploadTime != "" {
		cfg.Shitcamp.OldestUploadTime = oldestUploadTime
	}
	if newestUploadTime := os.Getenv(EnvShitcampNewestUploadTime); newestUploadTime != "" {
		cfg.Shitcamp.NewestUploadTime = newestUploadTime
	}

	if expiryTime := os.Getenv(EnvCacheDefaultExpiryTime); expiryTime != "" {
		if d, err := time.ParseDuration(expiryTime); err != nil {
			logger.WithField("expiryTime", expiryTime).WithError(err).Error("config_expiryTime_error")
		} else {
			cfg.Cache.DefaultExpiryTime = d
		}
	}
	if cleanupInterval := os.Getenv(EnvCacheDefaultCleanupInterval); cleanupInterval != "" {
		if d, err := time.ParseDuration(cleanupInterval); err != nil {
			logger.WithField("cleanupInterval", cleanupInterval).WithError(err).Error("config_cleanupInterval_error")
		} else {
			cfg.Cache.DefaultCleanupInterval = d
		}
	}

	if authStr := os.Getenv(EnvAuth); authStr != "" {
		var auth gin.Accounts
		err := json.Unmarshal([]byte(authStr), &auth)
		if err != nil {
			logger.WithField("auth", authStr).WithError(err).Error("config_auth_error")
		} else {
			cfg.Auth = auth
		}
	}

	if rateLimitTokens := os.Getenv(EnvRateLimitTokens); rateLimitTokens != "" {
		if tokens, err := strconv.ParseUint(rateLimitTokens, 10, 64); err != nil {
			logger.WithField("tokens", tokens).WithError(err).Error("config_rateLimitTokens_error")
		} else {
			cfg.RateLimit.Tokens = tokens
		}
	}
	if rateLimitInterval := os.Getenv(EnvRateLimitInterval); rateLimitInterval != "" {
		if d, err := time.ParseDuration(rateLimitInterval); err != nil {
			logger.WithField("rateLimitInterval", rateLimitInterval).WithError(err).Error("config_rateLimitInterval_error")
		} else {
			cfg.RateLimit.Interval = d
		}
	}

	if debug := os.Getenv(EnvDebug); debug != "" {
		if debugBool, err := strconv.ParseBool(debug); err != nil {
			logger.WithField("debugBool", debugBool).WithError(err).Error("config_debugBool_error")
		} else {
			cfg.Debug = debugBool
		}
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
