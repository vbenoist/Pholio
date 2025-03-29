package main

import (
	"fmt"

	"github.com/jackidu14/pholio/internal/database/connector"
	"github.com/jackidu14/pholio/internal/helpers/auth"
	"github.com/jackidu14/pholio/internal/helpers/cfg"
	routerserver "github.com/jackidu14/pholio/internal/server"
)

func main() {
	config := cfg.GetServerConfig()
	connector.Connect()
	defer connector.Disconnect()
	auth.InitServerDatabase()

	router := routerserver.SetupRouter(config)
	router.Run(fmt.Sprintf(":%s", config.Instance.Port))
}
