package terraplug

import (
	"github.com/hashicorp/hcl/v2/hclsimple"
)

const (
	// ConfigFilename is the name of the terraplug configuration file
	ConfigFilename = "terraplug.hcl"
)

// Config represents the terraplug configuration for the project
type Config struct {
	Plugins []PluginConfig `hcl:"plugin,block"`
}

// PluginConfig represents the configuration required to download a plugin
type PluginConfig struct {
	Type         string             `hcl:"type,label"`
	Name         string             `hcl:"name,label"`
	Version      *string            `hcl:"version"`
	Replacements *map[string]string `hcl:"replacements,attr"`
}

// LoadFile loads the terraplug configuration for a Terraform project directory
func LoadFile(file string) (*Config, error) {
	config := &Config{
		Plugins: make([]PluginConfig, 0),
	}
	err := hclsimple.DecodeFile(file, nil, config)
	return config, err
}
