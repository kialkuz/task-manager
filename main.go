package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kialkuz/task-manager/internal/api"
	"github.com/kialkuz/task-manager/internal/config"
	"github.com/kialkuz/task-manager/internal/delivery/http/handler/task"
	"github.com/kialkuz/task-manager/internal/infrastructure/env"
	repository "github.com/kialkuz/task-manager/internal/infrastructure/repository/db"
	"github.com/kialkuz/task-manager/internal/logger/logrus"
	"github.com/kialkuz/task-manager/internal/server"
	taskService "github.com/kialkuz/task-manager/internal/services/task"
	"github.com/kialkuz/task-manager/internal/validator"
)

func main() {
	env.Load()
	cfgData := config.Load()

	db, err := sql.Open(cfgData.DB.DBType, cfgData.DB.DatabaseURI)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	logger := logrus.NewLogger()

	taskHandler := task.NewHandler(
		validator.NewValidator(),
		validator.NewErrorFormatter(),
		logger,
		taskService.NewService(repository.NewTaskRepository(db)),
	)

	server := server.Handle(config.ServerConfig{
		Logger: logger,
		Routes: []api.RouteRegister{
			taskHandler,
		},
	})
	fmt.Println("Port for start: " + env.GetEnv("TODO_PORT", "7540"))

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}
}
