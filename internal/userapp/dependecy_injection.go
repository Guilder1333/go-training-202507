package userapp

import (
	"hands_on_go/internal/dal"
	"hands_on_go/internal/logic"
	"hands_on_go/internal/presentation"
	"net/http"

	"github.com/rs/zerolog/log"
)

type application struct {
	rootHandler http.Handler
}

func newApplication() (*application, error) {
	config, err := loadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	mysqlConf := dal.MySQLConfig{
		User:     config.MySQL.User,
		Password: config.MySQL.Password,
		Address:  config.MySQL.Address,
		DBName:   config.MySQL.DBName,
	}

	db, err := dal.NewMySQLDB(mysqlConf)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to MySQL database")
	}

	controller, err := presentation.NewUserController(
		presentation.NewUserValidatorImpl(),
		logic.NewUserServiceImpl(dal.NewMySQLUserRepository(db)),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create user controller")
	}

	return &application{
		rootHandler: presentation.NewHandler(controller),
	}, nil
}

func closeApplication(app *application) {
	// TODO
}
