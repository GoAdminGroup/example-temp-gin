package tables

import (
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "user_info" => http://localhost:9033/admin/info/user_info
//
var Generators = map[string]table.Generator{
	"posts":   GetPostsTable,
	"authors": GetAuthorsTable,
}

func lg(v string) string {
	return language.Get(v)
}
