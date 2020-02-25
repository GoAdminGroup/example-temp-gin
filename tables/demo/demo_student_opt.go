package demo

import "github.com/GoAdminGroup/go-admin/template/types"

func sexOptionInit() types.FieldOptions {
	return types.FieldOptions{
		types.FieldOption{
			Text:          "Male",
			Value:         "1",
			Selected:      true,
			SelectedLabel: "Male",
		},
		types.FieldOption{
			Text:          "Female",
			Value:         "2",
			Selected:      false,
			SelectedLabel: "Female",
		},
		types.FieldOption{
			Text:          "unknown",
			Value:         "0",
			Selected:      false,
			SelectedLabel: "unknown",
		},
	}
}

func sexOption() types.FieldOptions {
	return types.FieldOptions{
		types.FieldOption{
			Text:  "Male",
			Value: "1",
		},
		types.FieldOption{
			Text:  "Female",
			Value: "2",
		},
		types.FieldOption{
			Text:  "unknown",
			Value: "0",
		},
	}
}
