package demo

import (
	"github.com/GoAdminGroup/example-temp-gin/pkg/zlog"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	editType "github.com/GoAdminGroup/go-admin/template/types/table"
	"github.com/sinlovgo/timezone"
)

func GetDemoStudentTable(ctx *context.Context) table.Table {

	demoStudentTable := table.NewDefaultTable(table.
		DefaultConfigWithDriver("mysql").
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
		FieldFilterable(types.FilterType{FormType: form.NumberRange})
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
		FieldFilterable(types.FilterType{FormType: form.SelectSingle}). // In-Filter use SelectSingle
		FieldFilterOptions(sexOptionInit()).
		FieldEditAble(editType.Select).
		FieldEditOptions(sexOption())
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
		FieldOptionInitFn(func(model types.FieldModel) types.FieldOptions {
			return sexOptionInit()
		}).
		FieldOptions(sexOption()).
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
		FieldNotAllowEdit().
		FieldNotAllowAdd()
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime).
		FieldDefault(timezone.UTCSecond()).
		FieldValue(timezone.UTCSecond()).
		FieldNotAllowEdit()

	formList.SetTable("demo_student").SetTitle("Demo_student").SetDescription("Demo_student")

	formList.SetPostValidator(func(model form2.Values) error {
		zlog.S().Debugf("SetPostValidator model %v", model)
		return nil
	})

	formList.SetPostHook(func(model form2.Values) error {
		zlog.S().Debugf("SetPostHook model %v", model)
		return nil
	})

	//formList.SetInsertFn(func(model form2.Values) error {
	//	log.Debugf("SetInsertFn model %v", model)
	//	return nil
	//})

	//formList.SetUpdateFn(func(model form2.Values) error {
	//	log.Debugf("SetUpdateFn model %v", model)
	//	return nil
	//})

	return demoStudentTable
}
