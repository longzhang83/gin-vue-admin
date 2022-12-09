package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/MRP"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/mrp"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	MrpServiceGroup     MRP.ServiceGroup
	MrpServiceGroup     mrp.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
