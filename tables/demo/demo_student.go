package demo

import (
	"github.com/GoAdminGroup/example-temp-gin/util/timestamp"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	editType "github.com/GoAdminGroup/go-admin/template/types/table"
)

func GetDemoStudentTable(ctx *context.Context) table.Table {

	demoStudentTable := table.NewDefaultTable(table.
		DefaultConfigWithDriver("sqlite").
		SetExportable(true).
		SetDeletable(true))

	info := demoStudentTable.GetInfo()
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

	info.AddField("Id", "id", db.Integer)
	info.AddField("Stu_name", "stu_name", db.Varchar).
		FieldDisplay(func(model types.FieldModel) interface{} {
			return "<span class='label' style='color: #fb0000;'>" + model.Value + "</span>"
		}).
		FieldFilterable()
	info.AddField("Stu_age", "stu_age", db.Integer).
		FieldFilterable()
	info.AddField("Stu_sex", "stu_sex", db.Integer).
		FieldDisplay(func(model types.FieldModel) interface{} {
			if model.Value == "1" {
				return "Male"
			}
			if model.Value == "2" {
				return "Female"
			}
			return "unknown"
		}).
		FieldEditAble(editType.Select).
		FieldEditOptions([]map[string]string{ // In-form editing options
			{"value": "1", "text": "Male"},    // value DataRecord 1 text ShowText Male
			{"value": "2", "text": "Female"},  // value DataRecord 2  text ShowText Female
			{"value": "0", "text": "unknown"}, // value DataRecord 0  text ShowText unknown
		}).FieldFilterable(types.FilterType{FormType: form.SelectSingle}). // In-Filter use SelectSingle
		FieldFilterOptions([]map[string]string{ // In-Filter editing options
			{"value": "1", "field": "Male"},
			{"value": "2", "field": "Female"},
			{"value": "0", "field": "unknown"},
		})
	info.AddField("Created_at", "created_at", db.Datetime).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange}) // Show filters by creation time
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("demo_student").SetTitle("Demo_student").SetDescription("Demo_student")

	formList := demoStudentTable.GetForm()

	formList.AddField("Id", "id", db.Integer, form.Default).
		FieldNotAllowAdd()
	formList.AddField("Stu_name", "stu_name", db.Varchar, form.Text).
		FieldHelpMsg("in put student name").
		FieldMust()
	formList.AddField("Stu_age", "stu_age", db.Integer, form.Number).
		FieldHelpMsg("in put student age").
		FieldDefault("10").
		FieldMust()
	formList.AddField("Stu_sex", "stu_sex", db.Integer, form.SelectSingle).
		FieldOptions([]map[string]string{
			{
				"field":    "Male",
				"value":    "1",
				"selected": "selected",
			},
			{
				"field": "Female",
				"value": "2",
			},
		}).
		FieldDefault("1").
		FieldDisplay(func(model types.FieldModel) interface{} {
			switch model.Value {
			default:
				return "unknown"
			case "1":
				return "Male"
			case "2":
				return "Female"
			}
		}).
		FieldMust()
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime).
		FieldNotAllowAdd()
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime).
		FieldDefault(timestamp.LocalTimeSecond()).
		FieldValue(timestamp.LocalTimeSecond()).
		FieldNotAllowEdit()

	formList.SetTable("demo_student").SetTitle("Demo_student").SetDescription("Demo_student")

	return demoStudentTable
}
