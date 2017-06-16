package main

import (
	_ "github.com/Unknwon/goconfig"
	"github.com/astaxie/beego"
	_ "github.com/codegangsta/cli"
	log "github.com/sirupsen/logrus"
	"zplay.com/whao/mygopm/hello"
)

func main() {
	println("Beego version:", beego.VERSION)
	log.Info(".....=====.....")
	hello.Say()
}
