package config

import (
	"os"
	"time"

	"github.com/rapatao/go-injector"
	"github.com/rs/zerolog/log"
)

const (
	defaultApiKey   = "mock.luPzMYLzMTWJqMLCO37ZJRmPllQ7ct78"
	defaultDuration = 24 * time.Hour
)

type Config struct {
	ApiKey         string
	Event          bool
	EventEndpoint  string
	History        time.Duration
	WebStaticFiles string
	StorageFile    string
	BlockDeleteAll bool
	MessageDelay   time.Duration
}

func (c *Config) Initialize(_ *injector.Container) error {
	c.apiKey()
	c.events()
	c.history()
	c.webStaticFiles()
	c.storageDir()
	c.blockDeleteAll()
	c.messageDelay()

	return nil
}

func (c *Config) apiKey() {
	env := os.Getenv("API_KEY")
	if env == "" {
		log.Info().Msgf("No API_KEY environment variable set, using default %s", defaultApiKey)

		c.ApiKey = defaultApiKey

		return
	}

	c.ApiKey = env
}

func (c *Config) events() {
	env := os.Getenv("EVENT_DELIVERY_URL")
	if env == "" {
		log.Info().Msg("no EVENT_DELIVERY_URL environment variable set")

		c.Event = false

		return
	}

	log.Info().Msgf("event enabled using %s", env)

	c.Event = true
	c.EventEndpoint = env
}

func (c *Config) history() {
	duration, err := time.ParseDuration(os.Getenv("MAIL_HISTORY_DURATION"))
	if err == nil {
		log.Info().Msgf("history enabled using %s", duration.String())

		c.History = duration

		return
	}

	log.Error().Err(err).Msgf("failed to parse MAIL_HISTORY_DURATION, using default %s", defaultDuration.String())

	c.History = defaultDuration

	return
}

func (c *Config) webStaticFiles() {
	static := os.Getenv("WEB_STATIC_FILES")

	if static == "" {
		static = "./web/"
	}

	c.WebStaticFiles = static
}

func (c *Config) storageDir() {
	storage := os.Getenv("STORAGE_FILE")
	if storage == "" {
		file, err := os.CreateTemp(os.TempDir(), "sendgrid-mock.")
		if err != nil {
			log.Fatal().Err(err).Msgf("failed to create temp file")
		}

		storage = file.Name()
	}

	log.Info().Msgf("using %s to store messages", storage)

	c.StorageFile = storage
}

func (c *Config) messageDelay() {
	duration, err := time.ParseDuration(os.Getenv("MESSAGE_DELAY"))
	if err != nil {
		log.Error().Err(err).Msgf("failed to parse MESSAGE_DELAY, using default 0")

		c.MessageDelay = 0

		return
	}

	c.MessageDelay = duration

	return
}

func (c *Config) blockDeleteAll() {
	env := os.Getenv("BLOCK_DELETE_ALL")
	if env == "true" {
		c.BlockDeleteAll = true
	}

	c.BlockDeleteAll = false

	log.Info().Msgf("block deleting all messages is %t", c.BlockDeleteAll)
}

var _ injector.Injectable = (*Config)(nil)
