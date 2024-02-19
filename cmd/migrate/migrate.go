package main

import (
	"os"

	"github.com/alvinfebriando/project-batman-be/logger"
	"github.com/alvinfebriando/project-batman-be/migration"
	"github.com/alvinfebriando/project-batman-be/repository"
)

func main() {
	logger.SetLogrusLogger()

	_ = os.Setenv("APP_ENV", "debug")
	
	db, err := repository.GetConnection()
	if err != nil {
		logger.Log.Error(err)
	}

	migration.Migrate(db)
}
