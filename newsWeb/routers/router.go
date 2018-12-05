package routers

import (
	"newsWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register",&controllers.UserController{},"get:ShowRegister;post:HandleReg")
    beego.Router("/login",&controllers.UserController{},"get:ShowLogin;post:HandleLogin")
    beego.Router("/articleList",&controllers.ArticleController{},"get:ShowArticleList")
    //beego.Router("/addArticle",&controllers.ArticleController{})
	//beego.Router("/addArticle",&controllers.ArticleController{},"get:ShowAddArticle;post:HandeAddArticle")

}
