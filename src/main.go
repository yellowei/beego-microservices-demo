package main

import (
	_ "code2/newsWeb/models"
	_ "code2/newsWeb/routers"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {

	redisHost := beego.AppConfig.String("RedisHost")
	redisPort := beego.AppConfig.String("RedisPort")
	redisPass := beego.AppConfig.String("RedisPass")

	if beego.BConfig.WebConfig.Session.SessionProvider == "redis" {
		beego.BConfig.WebConfig.Session.SessionProviderConfig = redisHost + ":" + redisPort
		if redisPass != "" {
			beego.BConfig.WebConfig.Session.SessionProviderConfig = redisHost + ":" + redisPort + ",0," + redisPass
		}
	}

	//给视图函数建立映射
	beego.AddFuncMap("prePage", PrePage)
	beego.AddFuncMap("nextPage", NextPage)
	beego.Run()
}

func PrePage(pageNum int) int {
	if pageNum <= 1 {
		return 1
	}
	return pageNum - 1
}

func NextPage(pageNum int, pageCount float64) int {
	if pageNum >= int(pageCount) {
		return int(pageCount)
	}
	return pageNum + 1
}
