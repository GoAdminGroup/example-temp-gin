package demo

import (
	"github.com/GoAdminGroup/example-temp-gin/model/dbglobal"
	"github.com/GoAdminGroup/example-temp-gin/util/timestamp"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	editType "github.com/GoAdminGroup/go-admin/template/types/table"
	"github.com/lexkong/log"
	"strconv"
)

func GetDemoClassTable(ctx *context.Context) table.Table {

	demoClassTable := table.NewDefaultTable(table.
		DefaultConfigWithDriver("sqlite").
		SetExportable(true).
		SetDeletable(false))

	info := demoClassTable.GetInfo()
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
	info.AddField("Class_name", "class_name", db.Varchar)
	info.AddField("Class_desc", "class_desc", db.Varchar)
	info.AddField("Grade", "grade_id", db.Integer).
		FieldDisplay(func(model types.FieldModel) interface{} {
			gradleTable, err := dbglobal.Table("demo_grade")
			if err != nil {
				return model.Value
			}
			find, err := gradleTable.
				Select("grade_name").
				Where("id", "=", model.Value).First()
			if err != nil {
				return model.Value
			}
			return find["grade_name"].(string)
		}).
		FieldEditAble(editType.Select).
		FieldEditOptions(dbFetchDemoGradeFieldOption())
	info.AddField("Class_time_start", "class_time_start", db.Varchar)
	info.AddField("Class_time_end", "class_time_end", db.Varchar)
	info.AddField("Created_at", "created_at", db.Datetime).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange}) // Show filters by creation time
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("demo_class").SetTitle("Demo_class").SetDescription("Demo_class")

	formList := demoClassTable.GetForm()

	formList.AddField("Id", "id", db.Integer, form.Default).FieldNotAllowAdd()
	formList.AddField("Class_name", "class_name", db.Varchar, form.Text).
		FieldHelpMsg("must has class name").
		FieldMust()
	formList.AddField("Class_desc", "class_desc", db.Varchar, form.Text).
		FieldHelpMsg("can add some desc")
	formList.AddField("Grade_id", "grade_id", db.Integer, form.SelectSingle).
		FieldOptionInitFn(func(model types.FieldModel) types.FieldOptions {
			return dbFetchDemoGradeFieldOption()
		}).
		FieldMust()
	formList.AddField("Class_time_start", "class_time_start", db.Varchar, form.Text).
		FieldHelpMsg("class start time like 9:00").
		FieldMust()
	formList.AddField("Class_time_end", "class_time_end", db.Varchar, form.Text).
		FieldHelpMsg("class start time like 9:45").
		FieldMust()
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime).
		FieldNotAllowEdit().
		FieldNotAllowAdd()
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime).
		FieldDefault(timestamp.LocalTimeSecond()).
		FieldValue(timestamp.LocalTimeSecond()).
		FieldNotAllowEdit()

	formList.SetTable("demo_class").SetTitle("Demo_class").SetDescription("Demo_class")

	formList.SetPostValidator(func(model form2.Values) error {
		log.Debugf("SetPostValidator model %v", model)
		return nil
	})

	formList.SetPostHook(func(model form2.Values) error {
		log.Debugf("SetPostHook model %v", model)
		return nil
	})

	return demoClassTable
}

// db fetch demo_grade for show
func dbFetchDemoGradeFieldOption() types.FieldOptions {
	var options types.FieldOptions
	tableGradle, err := dbglobal.Table("demo_grade")
	if err != nil {
		return options
	}
	allGrade, err := tableGradle.All()
	if err != nil {
		return options
	}
	for _, grade := range allGrade {
		var raw types.FieldOption
		raw.Value = strconv.FormatInt(grade["id"].(int64), 10)
		raw.Text = grade["grade_name"].(string)
		options = append(options, raw)
	}
	return options
}
