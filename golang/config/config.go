package config

import (
	"os"
	"fmt"
	"flag"
	"io/ioutil"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// Config is a struct that holds the configuration for the application.
type Config struct {
	// Port is the port to listen on.
	Port int `json:"port"`
	// GinMode is the mode for the Gin framework.
	GinMode string `json:"gin_mode"`
	// GinLogFile is the log file for the Gin framework.
	GinLogFile string `json:"gin_log_file"`
}

// NewConfig returns a new Config struct.
func NewConfig() *Config {
	return &Config{}
}

// Load loads the configuration from the given file.
func (c *Config) Load(file string) error {
	// Open the configuration file.
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	// Read the configuration file.
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	// Unmarshal the configuration.
	err = json.Unmarshal(b, c)
	if err != nil {
		return err
	}

	return nil
}

// Parse parses the command line arguments.
func (c *Config) Parse() error {
	// Parse the command line arguments.
	flag.IntVar(&c.Port, "port", 8080, "Port to listen on")
	flag.StringVar(&c.GinMode, "gin-mode", gin.ReleaseMode, "Gin mode")
	flag.StringVar(&c.GinLogFile, "gin-log-file", "", "Gin log file")
	flag.Parse()

	// Set the Gin mode.
	gin.SetMode(c.GinMode)

	// Set the Gin log file.
	if c.GinLogFile != "" {
		f, err := os.OpenFile(c.GinLogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		gin.DefaultWriter = f
	}

	return nil
}

// String returns a string representation of the Config struct.
func (c *Config) String() string {
	return fmt.Sprintf("Port: %d, GinMode: %s, GinLogFile: %s", c.Port, c.GinMode, c.GinLogFile)
}
