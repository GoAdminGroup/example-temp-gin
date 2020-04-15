package router

import (
	configProject "github.com/GoAdminGroup/example-temp-gin/config"
	"github.com/GoAdminGroup/go-admin/modules/config"
)

func goAdminConfig(global configProject.Conf) config.Config {
	dbHost := parseEnvStringOrDefault(envKeyMysqlHost, global.GoAdmin.DataBases.Default.Host)
	dbPort := parseEnvStringOrDefault(envKeyMysqlPort, global.GoAdmin.DataBases.Default.Port)
	dbName := parseEnvStringOrDefault(envKeyMysqlDBName, global.GoAdmin.DataBases.Default.Name)
	dbUser := parseEnvStringOrDefault(envKeyMysqlDBUser, global.GoAdmin.DataBases.Default.User)
	dbPwd := parseEnvStringOrDefault(envKeyMysqlDBPWD, global.GoAdmin.DataBases.Default.Pwd)

	cfg := config.Config{
		Debug: configProject.IsDebug(),
		Databases: config.DatabaseList{
			"default": {
				Driver:     global.GoAdmin.DataBases.Default.Driver,
				Host:       dbHost,
				Port:       dbPort,
				Name:       dbName,
				User:       dbUser,
				Pwd:        dbPwd,
				MaxIdleCon: global.GoAdmin.DataBases.Default.MaxIdleCon,
				MaxOpenCon: global.GoAdmin.DataBases.Default.MaxOpenCon,
			},
		},
		IndexUrl:  global.GoAdmin.IndexURL,
		UrlPrefix: global.GoAdmin.URLPrefix, // The url prefix of the website.
		// Store must be set and guaranteed to have write access, otherwise new administrator users cannot be added.
		Store: config.Store{
			Path:   global.GoAdmin.Store.Path,
			Prefix: global.GoAdmin.Store.Prefix,
		},
		Language:    configGoAdminLanguageByYaml(),
		Theme:       global.GoAdmin.Theme,
		ColorScheme: global.GoAdmin.ColorScheme,
		SqlLog:      configProject.IsDebug(),
		Title:       global.GoAdmin.Title,
		LoginTitle:  global.GoAdmin.Title,
	}

	//cfg.LoginLogo = template.HTML(headTemplate)
	//cfg.ErrorLogPath = ``

	//cfg.Animation = config.PageAnimation{
	//	Type:     "fadeInUp",
	//	Duration: 0.9,
	//}

	return cfg
}
