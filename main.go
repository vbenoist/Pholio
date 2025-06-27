package main

import (
	"fmt"

	"github.com/vbenoist/pholio/internal/database/connector"
	"github.com/vbenoist/pholio/internal/helpers/auth"
	"github.com/vbenoist/pholio/internal/helpers/cfg"
	routerserver "github.com/vbenoist/pholio/internal/server"
)

func main() {
	config := cfg.GetServerConfig()
	connector.Connect()
	defer connector.Disconnect()
	auth.InitServerDatabase()

	router := routerserver.SetupRouter(config)
	router.Run(fmt.Sprintf(":%s", config.Instance.Port))
}
