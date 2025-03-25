package main

import (
	"fmt"

	"github.com/jackidu14/pholio/internal/database/connector"
	"github.com/jackidu14/pholio/internal/helpers/cfg"
	"github.com/jackidu14/pholio/internal/server"
)

func main() {
	config := cfg.SetServerConfig()
	connector.Connect()
	defer connector.Disconnect()

	router := server.SetupRouter(config)
	router.Run(fmt.Sprintf(":%s", config.Instance.Port))
}
