package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func NewConfigLoader(configName string) *ConfigLoader {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}

	return &ConfigLoader{
		v:          viper.New(),
		configName: configName,
		configType: "toml",
		configPath: []string{"."},
		env:        env,
	}
}

func (c *ConfigLoader) SetConfigName(name string) *ConfigLoader {
	c.configName = name
	return c
}

func (c *ConfigLoader) SetConfigType(configType string) *ConfigLoader {
	c.configType = configType
	return c
}

func (c *ConfigLoader) AddConfigPath(path string) *ConfigLoader {
	c.configPath = append(c.configPath, path)
	return c
}

func (c *ConfigLoader) SetEnvironment(env string) *ConfigLoader {
	c.env = env
	return c
}

func (c *ConfigLoader) LoadConfig() (*viper.Viper, error) {
	v := c.v

	configFile := fmt.Sprintf("%s.%s", c.configName, c.env)
	v.SetConfigName(configFile)
	v.SetConfigType(c.configType)

	for _, path := range c.configPath {
		v.AddConfigPath(path)
	}

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			v.SetConfigName(c.configName)
			if err := v.ReadInConfig(); err != nil {
				return nil, fmt.Errorf("error loading configuration: %w", err)
			}
		} else {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	return v, nil
}
