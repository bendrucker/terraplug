package terraplug

import (
	"fmt"
	"path/filepath"

	"github.com/hashicorp/hcl/v2/hclsimple"
	tfconfig "github.com/hashicorp/terraform/config"
)

const (
	// ConfigFilename is the name of the terraplug configuration file
	ConfigFilename = "terraplug.hcl"
)

// Config represents the terraplug configuration for the project
type Config struct {
	Plugins []*PluginConfig `hcl:"plugin,block"`
}

// PluginConfig represents the configuration required to download a plugin
type PluginConfig struct {
	Type         string             `hcl:"type,label"`
	Name         string             `hcl:"name,label"`
	Versions     *[]string          `hcl:"versions"`
	Replacements *map[string]string `hcl:"replacements,attr"`
}

// LoadFile loads the terraplug configuration from a file
func LoadFile(file string) (*Config, error) {
	config := &Config{}
	err := hclsimple.DecodeFile(file, nil, config)
	return config, err
}

// Load loads the terraplug configuration from a Terraform project directory
func Load(dir string) (*Config, error) {
	config, err := LoadFile(filepath.Join(dir, ConfigFilename))
	if err != nil {
		return nil, err
	}

	tf, err := tfconfig.LoadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, plugin := range config.Plugins {
		if plugin.Type != PluginTypeProvider || plugin.Versions != nil {
			continue
		}

		versions := make([]string, 0)

		for _, provider := range tf.ProviderConfigs {
			if provider.Name == plugin.Name && provider.Version != "" {
				versions = append(versions, provider.Version)
			}
		}

		if len(versions) == 0 {
			return nil, fmt.Errorf(`could not find a Terraform provider configuration named "%s" and no versions were specified`, plugin.Name)
		}

		plugin.Versions = &versions
	}

	return config, err
}
