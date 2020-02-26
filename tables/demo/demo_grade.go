package demo

import (
	"github.com/GoAdminGroup/example-temp-gin/util/timestamp"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/lexkong/log"
)

func GetDemoGradeTable(ctx *context.Context) table.Table {

	demoGradeTable := table.NewDefaultTable(table.
		DefaultConfigWithDriver("sqlite").
		SetExportable(true).
		SetDeletable(false))

	info := demoGradeTable.GetInfo()
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
	info.AddField("Grade_name", "grade_name", db.Varchar)
	info.AddField("Grade_desc", "grade_desc", db.Varchar)
	info.AddField("Created_at", "created_at", db.Datetime).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange}). // Show filters by creation time
		FieldDisplay(func(model types.FieldModel) interface{} {
			log.Debugf("model.Value -> %v", model.Value)
			timeStr, err := timestamp.ParseLocation("2006-01-02T15:04:05Z", "2006-01-02 15:04:05", model.Value, "UTC", "Asia/Shanghai")
			if err != nil {
				return model.Value
			}
			return timeStr
		})
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("demo_grade").SetTitle("Demo_grade").SetDescription("Demo_grade")

	formList := demoGradeTable.GetForm()

	formList.AddField("Id", "id", db.Integer, form.Default).FieldNotAllowAdd()
	formList.AddField("Grade_name", "grade_name", db.Varchar, form.Text).
		FieldMust()
	formList.AddField("Grade_desc", "grade_desc", db.Varchar, form.Text).
		FieldMust()
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime).
		FieldNotAllowEdit().
		FieldNotAllowAdd()
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime).
		FieldDefault(timestamp.UTCTimeSecond()).
		FieldValue(timestamp.UTCTimeSecond()).
		FieldNotAllowEdit()

	formList.SetTable("demo_grade").SetTitle("Demo_grade").SetDescription("Demo_grade")

	return demoGradeTable
}
