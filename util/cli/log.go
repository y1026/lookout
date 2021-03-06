package cli

import (
	"encoding/json"

	"gopkg.in/src-d/go-log.v1"
)

// LogOptions defines logging flags. It is meant to be embedded in a
// command struct.
type LogOptions struct {
	LogLevel       string `long:"log-level" env:"LOG_LEVEL" default:"info" description:"Logging level (info, debug, warning or error)"`
	LogFormat      string `long:"log-format" env:"LOG_FORMAT" description:"log format (text or json), defaults to text on a terminal and json otherwise"`
	LogFields      string `long:"log-fields" env:"LOG_FIELDS" description:"default fields for the logger, specified in json"`
	LogForceFormat bool   `long:"log-force-format" env:"LOG_FORCE_FORMAT" description:"ignore if it is running on a terminal or not"`
	Verbose        bool   `long:"verbose" short:"v" description:"enable verbose logging"`
}

var _ initializer = &LogOptions{}

// Init initializes the default logger factory.
func (c *LogOptions) init(app *App) {
	if c.Verbose {
		c.LogLevel = "debug"
	}

	if c.LogFields == "" {
		bytes, err := json.Marshal(log.Fields{"app": app.Name})
		if err != nil {
			panic(err)
		}
		c.LogFields = string(bytes)
	}

	log.DefaultFactory = &log.LoggerFactory{
		Level:       c.LogLevel,
		Format:      c.LogFormat,
		Fields:      c.LogFields,
		ForceFormat: c.LogForceFormat,
	}
	log.DefaultFactory.ApplyToLogrus()

	log.DefaultLogger = log.New(nil)
}
