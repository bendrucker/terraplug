package terraplug

import (
	"fmt"
	"path/filepath"
)

// Supported plugin types
const (
	PluginTypeProvider    PluginType = "provider"
	PluginTypeProvisioner            = "provisioner"
)

// PluginType describes a Terraform plugin type (https://www.terraform.io/docs/extend/plugin-types.html)
type PluginType string

// Plugin describes a third party Terraform plugin
type Plugin struct {
	// Name is the short name of the plugin. This is the name of the provider or provisioner.
	Name string
	Type PluginType

	// Version is a desired plugin version, specified in x.y.z form
	Version string
}

// Filename returns the plugin filename as specified by Terraform (https://www.terraform.io/docs/configuration/providers.html#plugin-names-and-versions)
func (p *Plugin) Filename() string {
	return fmt.Sprintf("terraform-%s-%s_v%s", p.Type, p.Name, p.Version)
}

// Path returns the relative path within the plugin directory where the plugin will be stored for the target OS/architecture
func (p *Plugin) Path(target Target) string {
	return filepath.Join(target.String(), p.Filename())
}

// Target is an operating system/architecture target for a plugin binary
type Target struct {
	OS   string
	Arch string
}

func (t *Target) String() string {
	return t.OS + "_" + t.Arch
}
