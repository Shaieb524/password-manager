package wire

import (
	apController "password-manager/src/controllers/accountPassword"
	apRepo "password-manager/src/models/database/accountPassword"
	dB "password-manager/src/providers/database"
	apService "password-manager/src/services/accountPassword"

	"github.com/google/wire"
)

func InitAccountPasswordAPI() *apController.AccountPasswordController {
	wire.Build(dB.NewDatabaseContext, apRepo.NewAccPasswordRepoModule, apService.NewAccPasswordServiceModule, apController.NewAccPasswordControllerModule)

	return &apController.AccountPasswordController{}
}
