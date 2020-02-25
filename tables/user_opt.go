package tables

import "github.com/GoAdminGroup/go-admin/template/types"

func genderOpt() types.FieldOptions {
	return types.FieldOptions{
		types.FieldOption{
			Text:          "men",
			Value:         "0",
			Selected:      false,
			SelectedLabel: "men",
		},
		types.FieldOption{
			Text:          "women",
			Value:         "1",
			Selected:      false,
			SelectedLabel: "women",
		},
	}
}
