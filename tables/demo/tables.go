package demo

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "demo_class" => http://localhost:9033/admin/info/demo_class
// "demo_student" => http://localhost:9033/admin/info/demo_student
// "demo_student_class" => http://localhost:9033/admin/info/demo_student_class
// "demo_student_score" => http://localhost:9033/admin/info/demo_student_score
//
var Generators = map[string]table.Generator{
	"demo_class": GetDemoClassTable,
	"demo_student": GetDemoStudentTable,
	"demo_student_class": GetDemoStudentClassTable,
	"demo_student_score": GetDemoStudentScoreTable,
}
