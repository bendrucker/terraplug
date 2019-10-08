package terraplug

// Config represents the terraplug configuration for the project
type Config struct {
	Plugins []PluginConfig
}

// PluginConfig represents the configuration required to download a plugin
type PluginConfig struct {
	Plugin

	URLTemplate  string
	Replacements map[string]string
}
