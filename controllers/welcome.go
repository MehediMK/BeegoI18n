package controllers

import (
	"fmt"
	"time"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

type WelcomeController struct {
	beego.Controller
	i18n.Locale
}

func CheckReq(lang string, welcomeController *WelcomeController) {
	if lang == "" {
		welcomeController.Ctx.SetCookie("lang_choice", "bn-BD")
	} else {
		welcomeController.Ctx.SetCookie("lang_choice", lang)
	}
	fmt.Println("func call:", lang)

}

func (welcomeController *WelcomeController) Get() {
	// here get value from url get request
	lang := welcomeController.GetString("lang")
	// here get value from header get request
	// lang := welcomeController.Ctx.Request.Header.Get("Accept-Language")

	// here check cookies lang_choice exists or not
	getlangcookie := welcomeController.Ctx.GetCookie("lang_choice")

	if getlangcookie != lang || getlangcookie == "" {
		CheckReq(lang, welcomeController)
		fmt.Println("function called Lang:", lang)
		fmt.Println("function called Cookies:", getlangcookie)
	}

	langcookies := &getlangcookie

	fmt.Println("Last cookies", *langcookies)

	welcomeController.Data["datetime"] = time.Now()
	welcomeController.Data["langTemplateKey"] = getlangcookie

	// welcomeController.Data["langTemplateKey"] = lang
	welcomeController.TplName = "welcome.tpl"
}
