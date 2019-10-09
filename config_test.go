package terraplug

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFileSumoLogic(t *testing.T) {
	config, err := testLoadFile("sumologic")
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
	config, err := testLoadFile("ct")
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, config.Plugins, 1)
	plugin := config.Plugins[0]

	assert.Equal(t, "ct", plugin.Name)
	assert.EqualValues(t, PluginTypeProvider, plugin.Type)
	assert.Nil(t, plugin.Replacements)
}

func TestLoadSumoLogic(t *testing.T) {
	config, err := testLoad("sumologic")
	if !assert.NoError(t, err) {
		return
	}

	plugin := config.Plugins[0]
	assert.Equal(t, []string{"1.x"}, *plugin.Versions)
}

func TestLoadCt(t *testing.T) {
	config, err := testLoad("ct")
	if !assert.NoError(t, err) {
		return
	}

	plugin := config.Plugins[0]
	assert.Equal(t, []string{"0.4"}, *plugin.Versions)
}

func TestLoadNoProviderError(t *testing.T) {
	_, err := testLoad("missing-provider")
	assert.EqualError(t, err, `could not find a Terraform provider configuration named "foo" and no versions were specified`)
}

func testLoadFile(name string) (*Config, error) {
	_, filename, _, _ := runtime.Caller(1)
	return LoadFile(filepath.Join(
		filepath.Dir(filename),
		"test",
		name,
		ConfigFilename,
	))
}

func testLoad(name string) (*Config, error) {
	_, filename, _, _ := runtime.Caller(1)
	return Load(filepath.Join(
		filepath.Dir(filename),
		"test",
		name,
	))
}
