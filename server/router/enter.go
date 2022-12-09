package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/MRP"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/mrp"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Mrp     MRP.RouterGroup
	Mrp     mrp.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
