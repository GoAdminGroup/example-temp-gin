package demo

import (
	"github.com/GoAdminGroup/example-temp-gin/util/timestamp"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetDemoStudentClassTable(ctx *context.Context) table.Table {

	demoStudentClassTable := table.NewDefaultTable(table.
		DefaultConfigWithDriver("sqlite").
		SetExportable(false).
		SetDeletable(false))

	info := demoStudentClassTable.GetInfo()
	// Form top layout
	// Create button is not disabled by default
	info.HideNewButton()
	// The export button is not hidden by default, export needs to be set SetExportable(true)
	info.HideExportButton()
	// Don't hide Filter layout by default
	//info.HideFilterArea()
	// Set Filter form layout
	info.SetFilterFormLayout(form.LayoutDefault)
	// Column selection row is not hidden by default
	info.HideRowSelector()
	// Set default sort
	// sort desc
	info.SetSortDesc()
	// sort asc
	//info.SetSortAsc()

	info.AddField("Id", "id", db.Integer).FieldFilterable()
	info.AddField("Class_id", "class_id", db.Integer).FieldFilterable()
	info.AddField("Stu_id", "stu_id", db.Integer).FieldFilterable()
	info.AddField("Created_at", "created_at", db.Datetime).
		FieldHide()
	info.AddField("Updated_at", "updated_at", db.Datetime).
		FieldHide()

	info.SetTable("demo_student_class").SetTitle("Demo_student_class").SetDescription("Demo_student_class")

	formList := demoStudentClassTable.GetForm()

	formList.AddField("Id", "id", db.Integer, form.Default).FieldNotAllowAdd()
	formList.AddField("Class_id", "class_id", db.Integer, form.Number).FieldNotAllowEdit()
	formList.AddField("Stu_id", "stu_id", db.Integer, form.Number).FieldNotAllowEdit()
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime).
		FieldNotAllowEdit().
		FieldNotAllowAdd()
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime).
		FieldDefault(timestamp.LocalTimeSecond()).
		FieldValue(timestamp.LocalTimeSecond()).
		FieldNotAllowEdit()

	formList.SetTable("demo_student_class").SetTitle("Demo_student_class").SetDescription("Demo_student_class")

	return demoStudentClassTable
}
