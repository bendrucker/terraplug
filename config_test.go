package terraplug

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFileSumoLogic(t *testing.T) {
	config, err := testConfig("sumologic")
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, config.Plugins, 1)
	plugin := config.Plugins[0]

	assert.Equal(t, "sumologic", plugin.Name)
	assert.EqualValues(t, PluginTypeProvider, plugin.Type)

	assert.NotNil(t, plugin.Replacements)
	_, ok := (*plugin.Replacements)["amd64"]
	assert.True(t, ok)
}

func TestLoadFileCt(t *testing.T) {
	config, err := testConfig("ct")
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, config.Plugins, 1)
	plugin := config.Plugins[0]

	assert.Equal(t, "ct", plugin.Name)
	assert.EqualValues(t, PluginTypeProvider, plugin.Type)
	assert.Nil(t, plugin.Replacements)
}

func testConfig(name string) (*Config, error) {
	_, filename, _, _ := runtime.Caller(1)
	return LoadFile(filepath.Join(
		filepath.Dir(filename),
		"test",
		name,
		ConfigFilename,
	))
}
