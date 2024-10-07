package example

import (
	c "github.com/jinge10/go-admin/modules/config"
	"github.com/jinge10/go-admin/modules/service"
	"github.com/jinge10/go-admin/plugins"
)

type Example struct {
	*plugins.Base
}

func NewExample() *Example {
	return &Example{
		Base: &plugins.Base{PlugName: "example"},
	}
}

func (e *Example) InitPlugin(srv service.List) {
	e.InitBase(srv, "example")
	e.App = e.initRouter(c.Prefix(), srv)
}
