package terraplug

import (
	"fmt"
	"path/filepath"
)

// Supported plugin types (https://www.terraform.io/docs/extend/plugin-types.html)
const (
	PluginTypeProvider    = "provider"
	PluginTypeProvisioner = "provisioner"
)

// Plugin describes a Terraform plugin
type Plugin struct {
	// Name is the short name of the plugin. This is the name of the provider or provisioner.
	Name string
	Type string
}

// PluginVersion describes a Terraform plugin at a specific version
type PluginVersion struct {
	Plugin

	// Version is a desired plugin version, specified in x.y.z form
	Version string
}

// Filename returns the plugin filename as specified by Terraform (https://www.terraform.io/docs/configuration/providers.html#plugin-names-and-versions)
func (p *PluginVersion) Filename() string {
	return fmt.Sprintf("terraform-%s-%s_v%s", p.Type, p.Name, p.Version)
}

// Path returns the relative path within the plugin directory where the plugin will be stored for the target OS/architecture
func (p *PluginVersion) Path(target Target) string {
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
