package main

import (
	"os"
	"strconv"
	"strings"
)

//sqlconfig
type Sqlconfigstr struct {
	User   string
	Pass   string
	db     string
	server string
	port   int
}

type GitHubConfig struct {
	Username string
	APIKey   string
}

type TeleConfigstr struct {
	Token    string
	telehigh []string
}

type Config struct {
	TeleConfig TeleConfigstr
	Sqlconfig  Sqlconfigstr
	GitHub     GitHubConfig
	DebugMode  bool
	UserRoles  []string
	MaxUsers   int
	Result     bool
}

// New returns a new Config struct
func New(ttype string) *Config {
	switch ttype {
	case "tele":
		return &Config{
			TeleConfig: TeleConfigstr{
				Token:    getEnv("TELETOKEN", ""),
				telehigh: getEnvAsSlice("telehigh", []string{"0"}, ","),
			},
			DebugMode: getEnvAsBool("DEBUG_MODE_TELE", true),
			Result:    true,
		}

	case "sql":
		return &Config{
			Sqlconfig: Sqlconfigstr{
				User:   getEnv("sqluser", ""),
				Pass:   getEnv("sqlpassword", ""),
				db:     getEnv("sqldb", ""),
				server: getEnv("sqlserver", ""),
				port:   getEnvAsInt("sqlport", 0),
			},
			DebugMode: getEnvAsBool("DEBUG_MODE_SQL", true),
			Result:    true,
		}

	default:
		return &Config{
			Result: false,
		}

	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
