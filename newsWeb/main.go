package main

import (
	_ "newsWeb/routers"
	"github.com/astaxie/beego"
	_ "newsWeb/models"
)

func main() {
	beego.Run()
}

