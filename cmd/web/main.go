package main

import (
	"fmt"

	"github.com/alonegrowing/insight/pkg/config"
	"github.com/alonegrowing/insight/pkg/web"
)

func main() {
	route := web.NewRouter()
	_ = route.Run(fmt.Sprintf(":%d", config.ServiceConfig.Service.WEBPort))
}
