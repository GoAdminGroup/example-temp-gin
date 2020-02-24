package plugin

import (
	"github.com/GoAdminGroup/example-temp-gin/tables"
	"github.com/GoAdminGroup/example-temp-gin/tables/demo"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
)

// add generator, first parameter is the url prefix of table when visit.
// example:
//
// "user" => http://localhost:9033/admin/info/user
//
func AdminPlugin() *admin.Admin {
	adminPlugin := admin.NewAdmin(tables.Generators, demo.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)
	return adminPlugin
}
