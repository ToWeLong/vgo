package main

import (
	"flag"
	"github.com/towelong/vgo/db"
	"github.com/towelong/vgo/global"
	"gorm.io/gen"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "conf", "../../configs/config.prod.yaml", "config path, eg: -conf config.prod.yaml")
}

func main() {
	//g := gen.NewGenerator(gen.Config{
	//	OutPath: "./biz/query",
	//	Mode:    gen.WithDefaultQuery,
	//})
	//
	//// generate from struct in project
	//g.ApplyBasic(model.Article{})
	//
	//g.Execute()
	flag.Parse()
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/query",
	})
	global.Init(configPath)

	g.UseDB(db.Conn())

	// generate all table from database
	g.ApplyBasic(g.GenerateAllTable()...)
	g.GenerateModel("article",
		gen.FieldType("delete_time", "gorm.DeletedAt"),
		gen.FieldType("id", "uint"),
	)
	g.Execute()
}
