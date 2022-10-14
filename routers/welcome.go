package routers

import (
	"BeegoI18n/controllers"
	"fmt"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

func init() {
	beego.Router("/welcome", &controllers.WelcomeController{})

	// here get langs variable value from conf directory
	langs, err := beego.AppConfig.String("langs")
	if err != nil {
		fmt.Println("Failed to load languages from the app.conf")
		return
	}

	// here split langs value
	langsArr := strings.Split(langs, "|")

	for _, lang := range langsArr {
		if err := i18n.SetMessage(lang, "conf/"+lang+".ini"); err != nil {
			fmt.Println("Failed to set message file for l10n")
			return
		}
	}
}
