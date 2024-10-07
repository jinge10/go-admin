package beego

import (
	// add beego adapter
	_ "github.com/jinge10/go-admin/adapter/beego"
	"github.com/jinge10/go-admin/modules/config"
	"github.com/jinge10/go-admin/modules/language"
	"github.com/jinge10/go-admin/plugins/admin/modules/table"

	// add mysql driver
	_ "github.com/jinge10/go-admin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "github.com/jinge10/go-admin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "github.com/jinge10/go-admin/modules/db/drivers/sqlite"
	// add mssql driver
	_ "github.com/jinge10/go-admin/modules/db/drivers/mssql"
	// add adminlte ui theme
	"github.com/GoAdminGroup/themes/adminlte"

	"net/http"
	"os"

	"github.com/astaxie/beego"
	"github.com/jinge10/go-admin/engine"
	"github.com/jinge10/go-admin/plugins/admin"
	"github.com/jinge10/go-admin/plugins/example"
	"github.com/jinge10/go-admin/template"
	"github.com/jinge10/go-admin/template/chartjs"
	"github.com/jinge10/go-admin/tests/tables"
)

func internalHandler() http.Handler {

	app := beego.NewApp()

	eng := engine.Default()
	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)

	examplePlugin := example.NewExample()

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(app); err != nil {
		panic(err)
	}

	template.AddComp(chartjs.NewChart())

	eng.HTML("GET", "/admin", tables.GetContent)

	beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	beego.BConfig.Listen.HTTPPort = 9087

	return app.Handlers
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {

	app := beego.NewApp()

	eng := engine.Default()
	adminPlugin := admin.NewAdmin(gens)

	if err := eng.AddConfig(&config.Config{
		Databases: dbs,
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:    language.EN,
		IndexUrl:    "/",
		Debug:       true,
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}).
		AddPlugins(adminPlugin).Use(app); err != nil {
		panic(err)
	}

	template.AddComp(chartjs.NewChart())

	eng.HTML("GET", "/admin", tables.GetContent)

	beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	beego.BConfig.Listen.HTTPPort = 9087

	return app.Handlers
}
