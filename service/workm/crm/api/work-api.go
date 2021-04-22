package main

import (
	"dark-work/common/configz"
	"dark-work/common/errorz"
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"

	"dark-work/service/workm/crm/api/internal/config"
	"dark-work/service/workm/crm/api/internal/handler"
	"dark-work/service/workm/crm/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/work-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(configz.MustMergeFile("assets/etcs/common.yaml", *configFile), &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)
	httpx.SetErrorHandler(errorz.ErrorHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
