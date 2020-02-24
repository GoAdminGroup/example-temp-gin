package errdef

import (
	"fmt"
	"github.com/GoAdminGroup/example-temp-gin/config"
	"github.com/GoAdminGroup/example-temp-gin/util/security"
)

func NoRoute404(key, val string) string {
	if key == "" {
		key = "NaN"
	}
	shot := fmt.Sprintf("%v=%v",
		security.URLPathEncode(key),
		"NoRoute",
	)
	why := fmt.Sprintf("%v%v%v",
		"NoRoute: ",
		security.URLPathEncode(config.BaseURL()),
		security.URLPathEncode(val),
	)
	return fmt.Sprintf("%v/404.html?%v=%v&%v=%v",
		config.BaseURL(),
		key, shot,
		"why", why,
	)
}
