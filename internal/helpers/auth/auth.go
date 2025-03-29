package auth

import (
	"github.com/jackidu14/pholio/internal/helpers/cfg"
	"github.com/jackidu14/pholio/internal/services/auth"
)

func InitServerDatabase() {
	cfg := cfg.GetServerConfig()

	/* Checking if default user admin has been registred */
	registred, err := auth.AdminDatabaseExists()
	if err != nil {
		panic(err)
	}

	/* If not registred in database, it's the first start. So, init default values... */
	if registred {
		return
	}

	err = auth.InitAdminDatabase(cfg.Front.User, cfg.Front.Pass)
	if err != nil {
		panic(err)
	}
}
