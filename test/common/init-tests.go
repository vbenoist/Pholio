package common_test

import (
	"github.com/vbenoist/pholio/internal/database/connector"
)

func Init() {
	// config := cfg.GetServerConfig()
	connector.Connect()
	defer connector.Disconnect()

	// router := server.SetupRouter(config)
	// router.Run(fmt.Sprintf(":%s", config.Instance.Port))
}
