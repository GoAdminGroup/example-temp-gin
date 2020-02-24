package tableutil

import (
	"github.com/GoAdminGroup/go-admin/template/types"
)

func OptionYesOrNo() []map[string]string {
	return []map[string]string{
		{
			"field":    "yes",
			"value":    "1",
			"selected": "selected",
		},
		{
			"field": "no",
			"value": "0",
		},
	}
}

func OptionNoOrYes() []map[string]string {
	return []map[string]string{
		{
			"field":    "no",
			"value":    "0",
			"selected": "selected",
		},
		{
			"field":    "yes",
			"value":    "1",
			"selected": "",
		},
	}
}

func EditOptionYesOrNo() []map[string]string {
	return []map[string]string{ // 表格内编辑 选项
		{
			"value":    "1",
			"text":     "yes",
			"selected": "checked",
		},
		{
			"value": "0",
			"text":  "no",
		},
	}
}

func DisplayYesOrNoInfo(model types.FieldModel) interface{} {
	if model.Value == "0" {
		return "no"
	} else if model.Value == "1" {
		return "yes"
	} else {
		return "-"
	}
}

func DisplayYesOrNoForm(model types.FieldModel) interface{} {
	switch model.Value {
	default:
		return "-"
	case "1":
		return "yes"
	case "0":
		return "no"
	}
}

func DisplayNoOrYesForm(model types.FieldModel) interface{} {
	switch model.Value {
	default:
		return "-"
	case "0":
		return "no"
	case "1":
		return "yes"
	}
}
