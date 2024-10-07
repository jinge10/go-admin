package gear

import (
	// add gin adapter
	ada "github.com/jinge10/go-admin/adapter/gear"
	"github.com/teambition/gear"

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

	"github.com/jinge10/go-admin/engine"
	"github.com/jinge10/go-admin/modules/config"
	"github.com/jinge10/go-admin/modules/language"
	"github.com/jinge10/go-admin/plugins/admin/modules/table"
	"github.com/jinge10/go-admin/template"
	"github.com/jinge10/go-admin/template/chartjs"
	"github.com/jinge10/go-admin/tests/tables"
)

func internalHandler() http.Handler {
	app := gear.New()

	eng := engine.Default()

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddGenerators(tables.Generators).
		AddGenerator("user", tables.GetUserTable).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	app := gear.New()

	eng := engine.Default()

	template.AddComp(chartjs.NewChart())

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
		AddAdapter(new(ada.Gear)).
		AddGenerators(gens).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app
}
