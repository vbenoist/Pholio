package common_test

import (
	"github.com/jackidu14/pholio/internal/database/connector"
)

func Init() {
	// config := cfg.SetServerConfig()
	connector.Connect()
	defer connector.Disconnect()

	// router := server.SetupRouter(config)
	// router.Run(fmt.Sprintf(":%s", config.Instance.Port))
}
