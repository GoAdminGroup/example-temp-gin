package router

import (
	"fmt"
	config2 "github.com/GoAdminGroup/example-temp-gin/config"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"html/template"
)

func goAdminCustom(cfg *config.Config, custom config2.Custom) {
	if custom.LogoUrl == "" {
		publicHead := `<link rel="icon" type="image/png" sizes="32x32" href="//quick.go-admin.cn/official/assets/imgs/icons.ico/favicon-32x32.png">
       <link rel="icon" type="image/png" sizes="96x96" href="//quick.go-admin.cn/official/assets/imgs/icons.ico/favicon-64x64.png">
       <link rel="icon" type="image/png" sizes="16x16" href="//quick.go-admin.cn/official/assets/imgs/icons.ico/favicon-16x16.png">`
		cfg.CustomHeadHtml = template.HTML(publicHead)
	} else {
		headTemplate := fmt.Sprintf(`<link rel="icon" type="image/png" sizes="32x32" href="%v">
      <link rel="icon" type="image/png" sizes="96x96" href="%v">
      <link rel="icon" type="image/png" sizes="16x16" href="%v">`, custom.LogoUrl, custom.LogoUrl, custom.LogoUrl)
		cfg.CustomHeadHtml = template.HTML(headTemplate)
	}

	//	cfg.CustomFootHtml = template.HTML(`<div style="display:none;">
	//    <script type="text/javascript" src="https://s9.cnzz.com/z_stat.php?id=1278156902&web_id=1278156902"></script>
	//</div>`)
}
