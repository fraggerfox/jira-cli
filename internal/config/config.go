package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

const (
	// BaseDir is the configuration base directory.
	BaseDir = ".config/jira"
	// Name is the configuration name.
	Name = "config"
)

var jsonConfig *viper.Viper

func init() {
	jsonConfig = viper.New()
	jsonConfig.SetConfigType("json")
}

// FileUsed returns the file used as a config.
func FileUsed() string {
	return viper.ConfigFileUsed()
}

// GetBaseDir returns the config directory base path.
func GetBaseDir() string {
	path := strings.Split(FileUsed(), "/")

	return strings.Join(path[0:len(path)-1], "/")
}

// GetServer returns server configured in used config.
func GetServer() string {
	return viper.GetString("server")
}

// GetLogin returns login configured in used config.
func GetLogin() string {
	return viper.GetString("login")
}

// GetProject returns project configured in used config.
func GetProject() string {
	return viper.GetString("project")
}

// GetBoardID returns board configured in used config.
func GetBoardID() int {
	return viper.GetInt("board")
}

// GetEpicField returns epic field configured in used config.
func GetEpicField() string {
	return viper.GetString("epic.field")
}

// GetIssueTypes fetches issue types from cache.
func GetIssueTypes() interface{} {
	jsonConfig.AddConfigPath(getCacheDir())
	jsonConfig.SetConfigName("_issuetypes")
	if err := jsonConfig.ReadInConfig(); err != nil {
		return nil
	}
	return jsonConfig.Get("data")
}

// GetBoardName returns board name from the configured board.
func GetBoardName() string {
	jsonConfig.AddConfigPath(getCacheDir())
	jsonConfig.SetConfigName("_boards")
	if err := jsonConfig.ReadInConfig(); err != nil {
		return "UNKNOWN"
	}
	return jsonConfig.GetString("data.name")
}

func getCacheDir() string {
	return fmt.Sprintf("%s/%s/%d", GetBaseDir(), strings.ToLower(GetProject()), GetBoardID())
}
