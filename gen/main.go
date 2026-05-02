package main

import (
	"bootpkg/gen/converter"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	connString  = "root:Lx_Bca0zj63322@(43.198.95.52:3306)/dev_game?charset=utf8&parseTime=true&loc=Local&tls=skip-verify"
	dataBase    = ""            //输出目录
	genDir      = "D:/code/gen" //输出目录
	snowflakeId = true
	isLogin     = true //是否生成登录验证代码
)

func main() {

	tableName := []string{
		"fc_customer_link",
		"fc_merchant_link",
	}
	os.Remove(genDir + "routes.js")
	os.Remove(genDir + "menus.sql")
	for _, table := range tableName {
		create(table)
	}
}
func create(tableName string) {
	//var err error
	t2t := converter.NewTable2Struct()
	// 个性化配置
	t2t.Config(&converter.T2tConfig{
		StructNameToHump: true,
		// 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
		RmTagIfUcFirsted: false,
		// tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
		TagToLower: false,
		// 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
		UcFirstOnly: true,
	})
	t2t.EnableJsonTag(true).
		TagKey("gorm").
		RealNameMethod("TableName").
		Table(tableName).
		Dsn(connString).
		AddColumnPrefix("column:").
		PackageName("dos").
		EnableSnowflakeId(false).
		BaseSQLColumnName(map[string]struct{}{
			"id": struct {
			}{},
		}).
		DateToTime(false)

	//首字母转小写
	//model 保存地址

	structName := t2t.CamelCase(tableName)
	filename := strings.ToLower(structName[0:1]) + structName[1:] //模块名称

	t2t.SavePath(filepath.Join(genDir, "pkg/core/modules/", "/dos/", filename+".go"))
	t2t.Run()

	data, tableNameComment, _ := t2t.GetColumnsTableName(tableName)
	m := &ModulesData{
		PackageName:  filename,
		StructName:   structName,
		Path:         "./template/",
		TableColumns: data[tableName],
	}
	//是否生成 登录Router添加
	if isLogin {
		m.LoginHanle = ".Use(handler.AuthMiddleware())"
		m.LoginPackage = "\r\n\"bootpkg/cmd/web/handler\""
	}

	var err error
	createPkg(m, err)
	log.Println("gen path  >>>>>>>>>>>>>>>>>>>>>> " + genDir)
	log.Println("gen code finish!!!")
	//写入前端路由
	f, err := os.OpenFile(genDir+"routes.js", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	f.WriteString(`
           {
              name: '` + tableNameComment + `',
              path: '/fc/` + m.PackageName + `/list',
              component: './fc/` + m.PackageName + `/list',
			  wrappers: [RouteWatcher],
          },
	`)
	fsql, err := os.OpenFile(genDir+"menus.sql", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	fsql.WriteString("\r\n" +
		"INSERT INTO `menus` (`menu_name`, `icon`, `sort`, `role_flag`, `address`, `parent_id`, `type`, `api_regular`, `show_status`) VALUES ('" + tableNameComment + "', '', 0, '/fc/" + m.PackageName + "/list', './fc/" + m.PackageName + "/list', 141, 2, '(*)', 1);" +
		"")
}

type ModulesData struct {
	PackageName  string
	StructName   string
	TableColumns []converter.Column
	Path         string //生成路径
	LoginPackage string //登录包名
	LoginHanle   string //登录中间件
}
