package demo

import (
	"github.com/GoAdminGroup/example-temp-gin/util/timestamp"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetDemoStudentScoreTable(ctx *context.Context) table.Table {

	demoStudentScoreTable := table.NewDefaultTable(table.
		DefaultConfigWithDriver("sqlite").
		SetExportable(true).
		SetDeletable(false))

	info := demoStudentScoreTable.GetInfo()
	// Form top layout
	// Create button is not disabled by default
	//info.HideNewButton()
	// The export button is not hidden by default, export needs to be set SetExportable(true)
	//info.HideExportButton()
	// Don't hide Filter layout by default
	info.HideFilterArea()
	// Set Filter form layout
	info.SetFilterFormLayout(form.LayoutDefault)
	// Column selection row is not hidden by default
	//info.HideRowSelector()
	// Set default sort
	// sort desc
	info.SetSortDesc()
	// sort asc
	//info.SetSortAsc()

	info.AddField("Id", "id", db.Integer).FieldFilterable()
	info.AddField("Stu_class_id", "stu_class_id", db.Varchar).FieldFilterable()
	info.AddField("Stu_score", "stu_score", db.Int)
	info.AddField("Created_at", "created_at", db.Datetime).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange}) // Show filters by creation time
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("demo_student_score").SetTitle("Demo_student_score").SetDescription("Demo_student_score")

	formList := demoStudentScoreTable.GetForm()

	formList.AddField("Id", "id", db.Integer, form.Default).FieldNotAllowAdd()
	formList.AddField("Stu_class_id", "stu_class_id", db.Varchar, form.Text).
		FieldMust()
	formList.AddField("Stu_score", "stu_score", db.Int, form.Number).
		FieldMust()
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime).
		FieldNotAllowAdd()
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime).
		FieldDefault(timestamp.LocalTimeSecond()).
		FieldValue(timestamp.LocalTimeSecond()).
		FieldNotAllowEdit()

	formList.SetTable("demo_student_score").SetTitle("Demo_student_score").SetDescription("Demo_student_score")

	return demoStudentScoreTable
}
