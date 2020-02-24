package plugin

import "github.com/GoAdminGroup/go-admin/plugins/example"

// customize a plugin
func Example() *example.Example {
	return example.NewExample()
}
