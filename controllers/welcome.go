package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

type WelcomeController struct {
	beego.Controller
	i18n.Locale
}

func (welcomeController *WelcomeController) Get() {
	// here get value from url get request
	// lang := welcomeController.GetString("lang")
	lang := welcomeController.Ctx.Request.Header.Get("Accept-Language")

	welcomeController.Data["langTemplateKey"] = lang
	welcomeController.TplName = "welcome.tpl"

}
