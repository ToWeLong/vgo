package main

import (
	"flag"
	"fmt"
	"github.com/towelong/vgo/db"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/towelong/vgo"
	"github.com/towelong/vgo/api/v1"
	"github.com/towelong/vgo/global"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "conf", "../../configs/config.prod.yaml", "config path, eg: -conf config.prod.yaml")
}

func initApp(srv *gin.Engine) {
	app := vgo.New(
		vgo.Name(global.Config.APP.Name),
		vgo.Version(global.Config.APP.Version),
		vgo.Addr(global.Config.APP.Addr),
		vgo.Timeout(time.Duration(global.Config.APP.Timeout)*time.Second),
		vgo.Server(srv),
	)
	printAppInfo(app)
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func printAppInfo(app *vgo.App) {
	global.Logger.Info(fmt.Sprintf("Name: %s", app.Name()))
	global.Logger.Info(fmt.Sprintf("Version: %s", app.Version()))
	cstZone := time.FixedZone("GMT", 8*3600)
	global.Logger.Info(fmt.Sprintf("Build: %s", time.Now().In(cstZone).Format("2006-01-02 15:04:05")))
}

func main() {
	flag.Parse()
	global.Init(configPath)
	db.Conn()
	initApp(v1.NewServer())
}
